package axhub

import "context"

// OperationParams carries low-level path, query, and JSON body values for generated route facades.
type OperationParams struct {
	PathParams map[string]string
	Query map[string]string
	Body any
}

func (c *Client) Operation(ctx context.Context, operationID string, params OperationParams) (map[string]any, error) {
	return c.Request(ctx, operationID, params.PathParams, params.Query, params.Body)
}


func (x *AppsClient) AppsDeleteApiV1AppsByAppID(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "appsDeleteApiV1AppsByAppID", params)
}
func (x *AppsClient) AppsGetApiV1AppsByAppID(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "appsGetApiV1AppsByAppID", params)
}
func (x *AppsClient) AppsPatchApiV1AppsByAppID(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "appsPatchApiV1AppsByAppID", params)
}
func (x *AppsClient) AppsDeleteApiV1AppsByAppIDAccess(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "appsDeleteApiV1AppsByAppIDAccess", params)
}
func (x *AppsClient) AppsPostApiV1AppsByAppIDAccess(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "appsPostApiV1AppsByAppIDAccess", params)
}
func (x *AppsClient) AppsGetApiV1AppsByAppIDAccessMe(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "appsGetApiV1AppsByAppIDAccessMe", params)
}
func (x *AppsClient) AppsPostApiV1AppsByAppIDArchive(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "appsPostApiV1AppsByAppIDArchive", params)
}
func (x *AppsClient) AppsGetApiV1AppsByAppIDComments(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "appsGetApiV1AppsByAppIDComments", params)
}
func (x *AppsClient) AppsPostApiV1AppsByAppIDComments(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "appsPostApiV1AppsByAppIDComments", params)
}
func (x *AppsClient) AppsGetApiV1AppsByAppIDEnvVars(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "appsGetApiV1AppsByAppIDEnvVars", params)
}
func (x *AppsClient) AppsPostApiV1AppsByAppIDEnvVars(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "appsPostApiV1AppsByAppIDEnvVars", params)
}
func (x *AppsClient) AppsPutApiV1AppsByAppIDEnvVarsByKeyStagingValue(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "appsPutApiV1AppsByAppIDEnvVarsByKeyStagingValue", params)
}
func (x *AppsClient) AppsDeleteApiV1AppsByAppIDLikes(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "appsDeleteApiV1AppsByAppIDLikes", params)
}
func (x *AppsClient) AppsPostApiV1AppsByAppIDLikes(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "appsPostApiV1AppsByAppIDLikes", params)
}
func (x *AppsClient) AppsGetApiV1AppsByAppIDLikesMe(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "appsGetApiV1AppsByAppIDLikesMe", params)
}
func (x *AppsClient) AppsDeleteApiV1AppsByAppIDRawDb(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "appsDeleteApiV1AppsByAppIDRawDb", params)
}
func (x *AppsClient) AppsPostApiV1AppsByAppIDRawDb(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "appsPostApiV1AppsByAppIDRawDb", params)
}
func (x *AppsClient) AppsGetApiV1AppsDiscover(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "appsGetApiV1AppsDiscover", params)
}
func (x *AppsClient) AppsGetApiV1AppsSearch(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "appsGetApiV1AppsSearch", params)
}
func (x *AppsClient) AppsDeleteApiV1CommentsByCommentID(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "appsDeleteApiV1CommentsByCommentID", params)
}
func (x *AppsClient) AppsGetApiV1MeAppsOwned(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "appsGetApiV1MeAppsOwned", params)
}
func (x *AppsClient) AppsGetApiV1MeAppsReceived(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "appsGetApiV1MeAppsReceived", params)
}
func (x *AppsClient) AppsGetApiV1MeAppsWorkspace(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "appsGetApiV1MeAppsWorkspace", params)
}
func (x *AppsClient) AppsGetApiV1Templates(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "appsGetApiV1Templates", params)
}
func (x *AppsClient) AppsGetApiV1TenantsByTenantIDApps(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "appsGetApiV1TenantsByTenantIDApps", params)
}
func (x *AppsClient) AppsPostApiV1TenantsByTenantIDApps(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "appsPostApiV1TenantsByTenantIDApps", params)
}
func (x *AppsClient) AppsGetApiV1TenantsByTenantIDAppsCheckAvailability(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "appsGetApiV1TenantsByTenantIDAppsCheckAvailability", params)
}
func (x *AppsClient) AppsGetApiV1TenantsByTenantIDDiscoverApps(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "appsGetApiV1TenantsByTenantIDDiscoverApps", params)
}
func (x *AppsClient) AppsGetApiV1UsersMeApps(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "appsGetApiV1UsersMeApps", params)
}

type IdentityClient struct{ client *Client }
func (c *Client) Identity() *IdentityClient { return &IdentityClient{client: c} }

func (x *IdentityClient) AuthGetWellKnownJwksJson(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "authGetWellKnownJwksJson", params)
}
func (x *IdentityClient) AuthGetWellKnownOauthAuthorizationServer(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "authGetWellKnownOauthAuthorizationServer", params)
}
func (x *IdentityClient) AuthGetWellKnownOpenidConfiguration(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "authGetWellKnownOpenidConfiguration", params)
}
func (x *IdentityClient) AuthPostApiV1AppsByAppIDOauthClients(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "authPostApiV1AppsByAppIDOauthClients", params)
}
func (x *IdentityClient) AuthGetApiV1Me(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "authGetApiV1Me", params)
}
func (x *IdentityClient) AuthGetApiV1OauthClientsByClientID(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "authGetApiV1OauthClientsByClientID", params)
}
func (x *IdentityClient) AuthDeleteApiV1OauthClientsByClientIDGrantsMe(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "authDeleteApiV1OauthClientsByClientIDGrantsMe", params)
}
func (x *IdentityClient) AuthPostAuthLogout(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "authPostAuthLogout", params)
}
func (x *IdentityClient) AuthPostAuthRefresh(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "authPostAuthRefresh", params)
}
func (x *IdentityClient) AuthGetOauthUserinfo(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "authGetOauthUserinfo", params)
}

type TenantsClient struct{ client *Client }
func (c *Client) Tenants() *TenantsClient { return &TenantsClient{client: c} }

func (x *TenantsClient) TenantsGetApiV1Tenants(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "tenantsGetApiV1Tenants", params)
}
func (x *TenantsClient) TenantsGetApiV1TenantsByTenantID(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "tenantsGetApiV1TenantsByTenantID", params)
}

type AuthzClient struct{ client *Client }
func (c *Client) Authz() *AuthzClient { return &AuthzClient{client: c} }

func (x *AuthzClient) AuthorizationGetApiV1TenantsByTenantIDMeGrants(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "authorizationGetApiV1TenantsByTenantIDMeGrants", params)
}

type AuditClient struct{ client *Client }
func (c *Client) Audit() *AuditClient { return &AuditClient{client: c} }


type GatewayClient struct{ client *Client }
func (c *Client) Gateway() *GatewayClient { return &GatewayClient{client: c} }

func (x *GatewayClient) GatewayGetApiV1TenantsByTenantIDConnectors(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "gatewayGetApiV1TenantsByTenantIDConnectors", params)
}
func (x *GatewayClient) GatewayGetApiV1TenantsByTenantIDConnectorsByConnectorID(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "gatewayGetApiV1TenantsByTenantIDConnectorsByConnectorID", params)
}
func (x *GatewayClient) GatewayPostApiV1TenantsByTenantIDConnectorsByConnectorIDDiscover(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "gatewayPostApiV1TenantsByTenantIDConnectorsByConnectorIDDiscover", params)
}
func (x *GatewayClient) GatewayGetApiV1TenantsByTenantIDConnectorsByConnectorIDResources(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "gatewayGetApiV1TenantsByTenantIDConnectorsByConnectorIDResources", params)
}
func (x *GatewayClient) GatewayPostApiV1TenantsByTenantIDGatewayDocumentInvoke(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "gatewayPostApiV1TenantsByTenantIDGatewayDocumentInvoke", params)
}
func (x *GatewayClient) GatewayPostApiV1TenantsByTenantIDGatewayFileInvoke(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "gatewayPostApiV1TenantsByTenantIDGatewayFileInvoke", params)
}
func (x *GatewayClient) GatewayPostApiV1TenantsByTenantIDGatewayInvoke(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "gatewayPostApiV1TenantsByTenantIDGatewayInvoke", params)
}
func (x *GatewayClient) GatewayPostApiV1TenantsByTenantIDGatewayQuery(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "gatewayPostApiV1TenantsByTenantIDGatewayQuery", params)
}
func (x *GatewayClient) GatewayPostApiV1TenantsByTenantIDGatewaySessions(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "gatewayPostApiV1TenantsByTenantIDGatewaySessions", params)
}
func (x *GatewayClient) GatewayDeleteApiV1TenantsByTenantIDGatewaySessionsBySessionID(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "gatewayDeleteApiV1TenantsByTenantIDGatewaySessionsBySessionID", params)
}
func (x *GatewayClient) GatewayGetApiV1TenantsByTenantIDMeConnectors(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "gatewayGetApiV1TenantsByTenantIDMeConnectors", params)
}
func (x *GatewayClient) GatewayGetApiV1TenantsByTenantIDMeConnectorsByConnectorIDResources(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "gatewayGetApiV1TenantsByTenantIDMeConnectorsByConnectorIDResources", params)
}
func (x *GatewayClient) ConfigGetConfigPublic(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "configGetConfigPublic", params)
}
func (x *GatewayClient) GatewayGetOauthGoogleCallback(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "gatewayGetOauthGoogleCallback", params)
}

type DataClient struct{ client *Client }
func (c *Client) Data() *DataClient { return &DataClient{client: c} }

func (x *DataClient) SchemaGetApiV1AppsByAppIDDbTables(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "schemaGetApiV1AppsByAppIDDbTables", params)
}
func (x *DataClient) SchemaGetApiV1AppsByAppIDDbTablesByTableRows(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "schemaGetApiV1AppsByAppIDDbTablesByTableRows", params)
}
func (x *DataClient) SchemaGetApiV1MePersonalAccessTokens(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "schemaGetApiV1MePersonalAccessTokens", params)
}
func (x *DataClient) SchemaPostApiV1MePersonalAccessTokens(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "schemaPostApiV1MePersonalAccessTokens", params)
}
func (x *DataClient) SchemaDeleteApiV1MePersonalAccessTokensByPatID(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "schemaDeleteApiV1MePersonalAccessTokensByPatID", params)
}

type DeploymentsClient struct{ client *Client }
func (c *Client) Deployments() *DeploymentsClient { return &DeploymentsClient{client: c} }

func (x *DeploymentsClient) DeployGetApiV1AppsByAppIDDeployments(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "deployGetApiV1AppsByAppIDDeployments", params)
}
func (x *DeploymentsClient) DeployPostApiV1AppsByAppIDDeployments(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "deployPostApiV1AppsByAppIDDeployments", params)
}
func (x *DeploymentsClient) DeployGetApiV1AppsByAppIDDeploymentsByDid(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "deployGetApiV1AppsByAppIDDeploymentsByDid", params)
}
func (x *DeploymentsClient) DeployPostApiV1AppsByAppIDDeploymentsByDidCancel(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "deployPostApiV1AppsByAppIDDeploymentsByDidCancel", params)
}
func (x *DeploymentsClient) DeployPostApiV1AppsByAppIDDeploymentsByDidRollback(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "deployPostApiV1AppsByAppIDDeploymentsByDidRollback", params)
}
func (x *DeploymentsClient) DeployGetApiV1AppsByAppIDDiagnosis(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "deployGetApiV1AppsByAppIDDiagnosis", params)
}
func (x *DeploymentsClient) DeployDeleteApiV1AppsByAppIDGitConnection(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "deployDeleteApiV1AppsByAppIDGitConnection", params)
}
func (x *DeploymentsClient) DeployGetApiV1AppsByAppIDGitConnection(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "deployGetApiV1AppsByAppIDGitConnection", params)
}
func (x *DeploymentsClient) DeployPatchApiV1AppsByAppIDGitConnection(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "deployPatchApiV1AppsByAppIDGitConnection", params)
}
func (x *DeploymentsClient) DeployPostApiV1AppsByAppIDGitConnection(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "deployPostApiV1AppsByAppIDGitConnection", params)
}
func (x *DeploymentsClient) DeployPostApiV1AppsByAppIDGitGithubConnect(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "deployPostApiV1AppsByAppIDGitGithubConnect", params)
}
func (x *DeploymentsClient) DeployGetApiV1AppsByAppIDGitGithubInstallStart(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "deployGetApiV1AppsByAppIDGitGithubInstallStart", params)
}
func (x *DeploymentsClient) DeployGetApiV1AppsByAppIDLogs(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "deployGetApiV1AppsByAppIDLogs", params)
}
func (x *DeploymentsClient) DeployPostApiV1AppsByAppIDPromotionsByRequestIDRetry(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "deployPostApiV1AppsByAppIDPromotionsByRequestIDRetry", params)
}
func (x *DeploymentsClient) DeployGetApiV1AppsByAppIDReleases(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "deployGetApiV1AppsByAppIDReleases", params)
}
func (x *DeploymentsClient) DeployGetApiV1AppsByAppIDReleasesPromotePreflight(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "deployGetApiV1AppsByAppIDReleasesPromotePreflight", params)
}
func (x *DeploymentsClient) DeployDeleteApiV1AppsByAppIDStaging(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "deployDeleteApiV1AppsByAppIDStaging", params)
}
func (x *DeploymentsClient) DeployPutApiV1AppsByAppIDStaging(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "deployPutApiV1AppsByAppIDStaging", params)
}
func (x *DeploymentsClient) DeployPostApiV1GitGithubComplete(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "deployPostApiV1GitGithubComplete", params)
}
func (x *DeploymentsClient) DeployGetApiV1GithubAccounts(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "deployGetApiV1GithubAccounts", params)
}
func (x *DeploymentsClient) DeployGetApiV1GithubInstallationsByInstallationIDRepositories(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "deployGetApiV1GithubInstallationsByInstallationIDRepositories", params)
}
func (x *DeploymentsClient) DeployPostApiV1TenantsByTenantIDAppBootstraps(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "deployPostApiV1TenantsByTenantIDAppBootstraps", params)
}
func (x *DeploymentsClient) DeployGetApiV1TenantsByTenantIDAppBootstrapsByBootstrapID(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "deployGetApiV1TenantsByTenantIDAppBootstrapsByBootstrapID", params)
}
func (x *DeploymentsClient) DeployGetApiV1TenantsByTenantIDDeployments(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "deployGetApiV1TenantsByTenantIDDeployments", params)
}

