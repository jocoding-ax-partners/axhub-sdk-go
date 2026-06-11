package axhub

var ContextRoutes = map[string][]Route{
	"apps": {},
	"identity": {},
	"tenants": {},
	"authz": {},
	"audit": {},
	"gateway": {},
	"cost": {},
	"data": {},
	"deployments": {},
}

func init() {
	for _, route := range Routes {
		name := contextName(route)
		ContextRoutes[name] = append(ContextRoutes[name], route)
	}
}

func contextName(route Route) string {
	switch route.Tag {
	case "Apps":
		return "apps"
	case "Auth", "identity":
		return "identity"
	case "Tenants":
		return "tenants"
	case "Authorization":
		return "authz"
	case "Audit":
		return "audit"
	case "Gateway", "Config":
		return "gateway"
	case "Cost":
		return "cost"
	case "Schema":
		return "data"
	case "Deploy", "deploy":
		return "deployments"
	default:
		panic("unmapped route tag: " + route.Tag)
	}
}
