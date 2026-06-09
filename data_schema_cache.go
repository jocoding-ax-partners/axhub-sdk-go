package axhub

import (
	"container/list"
	"sync"
	"time"
)

// Per-client schema cache for runtime Data().Discover() (mirrors node
// schema-cache.ts). LRU eviction via a container/list ordering + map index, plus
// a negative-TTL stale-while-error window: a transient 5xx during refresh keeps
// the previous entry alive briefly instead of evicting it.
//
// Go-specific: guarded by a sync.Mutex (Python's dict rode the GIL). The node
// in-flight de-dup map is intentionally omitted (see gap-matrix) — the mutex is
// map safety, not request coalescing.

const (
	DefaultSchemaCacheTTLMS         = 5 * 60_000
	DefaultSchemaCacheMaxEntries    = 1000
	DefaultSchemaCacheNegativeTTLMS = 30_000
)

// SchemaCacheOptions configures a SchemaCache.
type SchemaCacheOptions struct {
	MaxEntries    int
	TTLMS         int
	NegativeTTLMS int
}

type cacheEntry struct {
	key       string
	schema    *DataTableSchema
	expiresAt time.Time
}

// SchemaCache is an LRU + TTL cache of discovered schemas.
type SchemaCache struct {
	mu            sync.Mutex
	order         *list.List               // front = LRU, back = MRU
	index         map[string]*list.Element // key -> element holding *cacheEntry
	maxEntries    int
	ttl           time.Duration
	negativeTTL   time.Duration
	nowFn         func() time.Time
}

// NewSchemaCache builds a SchemaCache with the given options (zero-value
// options use the defaults).
func NewSchemaCache(opts SchemaCacheOptions) *SchemaCache {
	maxEntries := opts.MaxEntries
	if maxEntries < 1 {
		maxEntries = DefaultSchemaCacheMaxEntries
	}
	ttl := opts.TTLMS
	if ttl < 1 {
		ttl = DefaultSchemaCacheTTLMS
	}
	negTTL := opts.NegativeTTLMS
	if negTTL < 0 {
		negTTL = DefaultSchemaCacheNegativeTTLMS
	}
	return &SchemaCache{
		order:       list.New(),
		index:       map[string]*list.Element{},
		maxEntries:  maxEntries,
		ttl:         time.Duration(ttl) * time.Millisecond,
		negativeTTL: time.Duration(negTTL) * time.Millisecond,
		nowFn:       time.Now,
	}
}

// Size returns the number of cached entries.
func (c *SchemaCache) Size() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.order.Len()
}

// Get returns a cached schema or nil. Refreshes recency for LRU.
func (c *SchemaCache) Get(key string) *DataTableSchema {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.getLocked(key)
}

func (c *SchemaCache) getLocked(key string) *DataTableSchema {
	el, ok := c.index[key]
	if !ok {
		return nil
	}
	entry := el.Value.(*cacheEntry)
	if !entry.expiresAt.After(c.nowFn()) {
		c.order.Remove(el)
		delete(c.index, key)
		return nil
	}
	c.order.MoveToBack(el)
	return entry.schema
}

// Set caches a schema with an optional TTL override (0 uses the default).
func (c *SchemaCache) Set(key string, schema *DataTableSchema, ttlMS int) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.setLocked(key, schema, ttlMS)
}

func (c *SchemaCache) setLocked(key string, schema *DataTableSchema, ttlMS int) {
	if el, ok := c.index[key]; ok {
		c.order.Remove(el)
		delete(c.index, key)
	}
	ttl := c.ttl
	if ttlMS > 0 {
		ttl = time.Duration(ttlMS) * time.Millisecond
	}
	entry := &cacheEntry{key: key, schema: schema, expiresAt: c.nowFn().Add(ttl)}
	el := c.order.PushBack(entry)
	c.index[key] = el
	c.evictOverflowLocked()
}

// Invalidate drops one key, or all keys when key == "".
func (c *SchemaCache) Invalidate(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if key == "" {
		c.order.Init()
		c.index = map[string]*list.Element{}
		return
	}
	if el, ok := c.index[key]; ok {
		c.order.Remove(el)
		delete(c.index, key)
	}
}

// getOrSet returns a cached schema or loads, caches, and returns it. On a
// transient 5xx the previous entry is kept alive for negativeTTL.
func (c *SchemaCache) getOrSet(key string, loader func() (*DataTableSchema, error), fresh bool, ttlMS int) (*DataTableSchema, error) {
	if !fresh {
		if cached := c.Get(key); cached != nil {
			return cached, nil
		}
	}
	c.mu.Lock()
	var previous *DataTableSchema
	if el, ok := c.index[key]; ok {
		previous = el.Value.(*cacheEntry).schema
	}
	c.mu.Unlock()

	schema, err := loader()
	if err != nil {
		if previous != nil && c.negativeTTL > 0 && isTransientServerError(err) {
			c.mu.Lock()
			if el, ok := c.index[key]; ok {
				c.order.Remove(el)
				delete(c.index, key)
			}
			entry := &cacheEntry{key: key, schema: previous, expiresAt: c.nowFn().Add(c.negativeTTL)}
			c.index[key] = c.order.PushBack(entry)
			c.mu.Unlock()
		}
		return nil, err
	}
	c.Set(key, schema, ttlMS)
	return schema, nil
}

func (c *SchemaCache) evictOverflowLocked() {
	for c.order.Len() > c.maxEntries {
		front := c.order.Front()
		if front == nil {
			return
		}
		entry := front.Value.(*cacheEntry)
		c.order.Remove(front)
		delete(c.index, entry.key)
	}
}

func isTransientServerError(err error) bool {
	if ax, ok := err.(*AxHubError); ok {
		return ax.Status >= 500
	}
	return false
}

func schemaCacheKey(tenantSlug, appSlug, table string) string {
	return tenantSlug + "/" + appSlug + "/" + table
}
