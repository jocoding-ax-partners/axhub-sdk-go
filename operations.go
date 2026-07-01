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


func (x *AppsClient) AppsGetApiV1AdminTemplates(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "appsGetApiV1AdminTemplates", params)
}
func (x *AppsClient) AppsPostApiV1AdminTemplates(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "appsPostApiV1AdminTemplates", params)
}
func (x *AppsClient) AppsGetApiV1AdminTemplatesByTemplateID(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "appsGetApiV1AdminTemplatesByTemplateID", params)
}
func (x *AppsClient) AppsPatchApiV1AdminTemplatesByTemplateID(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "appsPatchApiV1AdminTemplatesByTemplateID", params)
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
func (x *AppsClient) AppsDeleteApiV1AppsByAppIDEnvVarsByKey(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "appsDeleteApiV1AppsByAppIDEnvVarsByKey", params)
}
func (x *AppsClient) AppsPutApiV1AppsByAppIDEnvVarsByKeyStagingValue(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "appsPutApiV1AppsByAppIDEnvVarsByKeyStagingValue", params)
}
func (x *AppsClient) AppsPostApiV1AppsByAppIDIconDarkUploadUrl(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "appsPostApiV1AppsByAppIDIconDarkUploadUrl", params)
}
func (x *AppsClient) AppsPostApiV1AppsByAppIDIconUploadUrl(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "appsPostApiV1AppsByAppIDIconUploadUrl", params)
}
func (x *AppsClient) AppsPostApiV1AppsByAppIDInvitations(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "appsPostApiV1AppsByAppIDInvitations", params)
}
func (x *AppsClient) AppsDeleteApiV1AppsByAppIDInvitationsByUserID(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "appsDeleteApiV1AppsByAppIDInvitationsByUserID", params)
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
func (x *AppsClient) AppsGetApiV1AppsByAppIDMembers(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "appsGetApiV1AppsByAppIDMembers", params)
}
func (x *AppsClient) AppsDeleteApiV1AppsByAppIDPermanent(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "appsDeleteApiV1AppsByAppIDPermanent", params)
}
func (x *AppsClient) AppsDeleteApiV1AppsByAppIDRawDb(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "appsDeleteApiV1AppsByAppIDRawDb", params)
}
func (x *AppsClient) AppsPostApiV1AppsByAppIDRawDb(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "appsPostApiV1AppsByAppIDRawDb", params)
}
func (x *AppsClient) AppsPostApiV1AppsByAppIDReactivate(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "appsPostApiV1AppsByAppIDReactivate", params)
}
func (x *AppsClient) AppsPostApiV1AppsByAppIDResume(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "appsPostApiV1AppsByAppIDResume", params)
}
func (x *AppsClient) AppsGetApiV1AppsByAppIDReviewRequests(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "appsGetApiV1AppsByAppIDReviewRequests", params)
}
func (x *AppsClient) AppsPostApiV1AppsByAppIDReviewRequests(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "appsPostApiV1AppsByAppIDReviewRequests", params)
}
func (x *AppsClient) AppsPostApiV1AppsByAppIDSuspend(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "appsPostApiV1AppsByAppIDSuspend", params)
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
func (x *AppsClient) AppsGetApiV1ResourcePresets(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "appsGetApiV1ResourcePresets", params)
}
func (x *AppsClient) AppsGetApiV1ReviewRequestsByRrID(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "appsGetApiV1ReviewRequestsByRrID", params)
}
func (x *AppsClient) AppsPostApiV1ReviewRequestsByRrIDApprove(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "appsPostApiV1ReviewRequestsByRrIDApprove", params)
}
func (x *AppsClient) AppsPostApiV1ReviewRequestsByRrIDReject(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "appsPostApiV1ReviewRequestsByRrIDReject", params)
}
func (x *AppsClient) AppsGetApiV1ReviewRequestsHistory(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "appsGetApiV1ReviewRequestsHistory", params)
}
func (x *AppsClient) AppsGetApiV1ReviewRequestsPending(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "appsGetApiV1ReviewRequestsPending", params)
}
func (x *AppsClient) AppsGetApiV1StaticAuthStart(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "appsGetApiV1StaticAuthStart", params)
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
func (x *AppsClient) AppsGetApiV1TenantsByTenantIDAppsByAppIDStaticReleases(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "appsGetApiV1TenantsByTenantIDAppsByAppIDStaticReleases", params)
}
func (x *AppsClient) AppsPostApiV1TenantsByTenantIDAppsByAppIDStaticReleases(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "appsPostApiV1TenantsByTenantIDAppsByAppIDStaticReleases", params)
}
func (x *AppsClient) AppsDeleteApiV1TenantsByTenantIDAppsByAppIDStaticReleasesByReleaseID(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "appsDeleteApiV1TenantsByTenantIDAppsByAppIDStaticReleasesByReleaseID", params)
}
func (x *AppsClient) AppsPostApiV1TenantsByTenantIDAppsByAppIDStaticReleasesByReleaseIDActivate(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "appsPostApiV1TenantsByTenantIDAppsByAppIDStaticReleasesByReleaseIDActivate", params)
}
func (x *AppsClient) AppsPostApiV1TenantsByTenantIDAppsByAppIDStaticReleasesByReleaseIDFinalize(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "appsPostApiV1TenantsByTenantIDAppsByAppIDStaticReleasesByReleaseIDFinalize", params)
}
func (x *AppsClient) AppsPostApiV1TenantsByTenantIDAppsByAppIDStaticReleasesByReleaseIDPromoteApprove(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "appsPostApiV1TenantsByTenantIDAppsByAppIDStaticReleasesByReleaseIDPromoteApprove", params)
}
func (x *AppsClient) AppsPostApiV1TenantsByTenantIDAppsByAppIDStaticReleasesByReleaseIDPromoteReject(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "appsPostApiV1TenantsByTenantIDAppsByAppIDStaticReleasesByReleaseIDPromoteReject", params)
}
func (x *AppsClient) AppsPostApiV1TenantsByTenantIDAppsByAppIDStaticReleasesByReleaseIDPromoteRequest(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "appsPostApiV1TenantsByTenantIDAppsByAppIDStaticReleasesByReleaseIDPromoteRequest", params)
}
func (x *AppsClient) AppsPostApiV1TenantsByTenantIDAppsByAppIDStaticReleasesByReleaseIDRollback(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "appsPostApiV1TenantsByTenantIDAppsByAppIDStaticReleasesByReleaseIDRollback", params)
}
func (x *AppsClient) AppsPostApiV1TenantsByTenantIDAppsByAppIDStaticReleasesByReleaseIDStage(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "appsPostApiV1TenantsByTenantIDAppsByAppIDStaticReleasesByReleaseIDStage", params)
}
func (x *AppsClient) AppsGetApiV1TenantsByTenantIDAppsByAppIDStaticSite(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "appsGetApiV1TenantsByTenantIDAppsByAppIDStaticSite", params)
}
func (x *AppsClient) AppsPatchApiV1TenantsByTenantIDAppsByAppIDStaticSite(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "appsPatchApiV1TenantsByTenantIDAppsByAppIDStaticSite", params)
}
func (x *AppsClient) AppsPostApiV1TenantsByTenantIDAppsByAppIDStaticSiteStagingDisable(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "appsPostApiV1TenantsByTenantIDAppsByAppIDStaticSiteStagingDisable", params)
}
func (x *AppsClient) AppsPostApiV1TenantsByTenantIDAppsByAppIDStaticSiteStagingEnable(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "appsPostApiV1TenantsByTenantIDAppsByAppIDStaticSiteStagingEnable", params)
}
func (x *AppsClient) AppsPostApiV1TenantsByTenantIDAppsByAppIDStaticSiteUnpublish(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "appsPostApiV1TenantsByTenantIDAppsByAppIDStaticSiteUnpublish", params)
}
func (x *AppsClient) AppsGetApiV1TenantsByTenantIDAppsCheckAvailability(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "appsGetApiV1TenantsByTenantIDAppsCheckAvailability", params)
}
func (x *AppsClient) AppsPostApiV1TenantsByTenantIDAppsIconUploadUrl(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "appsPostApiV1TenantsByTenantIDAppsIconUploadUrl", params)
}
func (x *AppsClient) AppsGetApiV1TenantsByTenantIDCategories(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "appsGetApiV1TenantsByTenantIDCategories", params)
}
func (x *AppsClient) AppsPostApiV1TenantsByTenantIDCategories(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "appsPostApiV1TenantsByTenantIDCategories", params)
}
func (x *AppsClient) AppsDeleteApiV1TenantsByTenantIDCategoriesByCategoryID(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "appsDeleteApiV1TenantsByTenantIDCategoriesByCategoryID", params)
}
func (x *AppsClient) AppsGetApiV1TenantsByTenantIDCategoriesByCategoryID(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "appsGetApiV1TenantsByTenantIDCategoriesByCategoryID", params)
}
func (x *AppsClient) AppsPatchApiV1TenantsByTenantIDCategoriesByCategoryID(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "appsPatchApiV1TenantsByTenantIDCategoriesByCategoryID", params)
}
func (x *AppsClient) AppsGetApiV1TenantsByTenantIDDiscoverApps(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "appsGetApiV1TenantsByTenantIDDiscoverApps", params)
}
func (x *AppsClient) AppsGetApiV1UsersMeApps(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "appsGetApiV1UsersMeApps", params)
}
func (x *AppsClient) AppsGetInternalAppAccess(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "appsGetInternalAppAccess", params)
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
func (x *IdentityClient) AuthPostApiV1AdminUsersByUidRevokeAll(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "authPostApiV1AdminUsersByUidRevokeAll", params)
}
func (x *IdentityClient) AuthPostApiV1AppsByAppIDOauthClients(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "authPostApiV1AppsByAppIDOauthClients", params)
}
func (x *IdentityClient) AuthGetApiV1Me(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "authGetApiV1Me", params)
}
func (x *IdentityClient) AuthPostApiV1MeInvitationsByInvitationIDAccept(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "authPostApiV1MeInvitationsByInvitationIDAccept", params)
}
func (x *IdentityClient) AuthGetApiV1OauthClientsByClientID(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "authGetApiV1OauthClientsByClientID", params)
}
func (x *IdentityClient) AuthDeleteApiV1OauthClientsByClientIDGrantsMe(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "authDeleteApiV1OauthClientsByClientIDGrantsMe", params)
}
func (x *IdentityClient) AuthGetApiV1TenantsByTenantIDIdentityProviders(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "authGetApiV1TenantsByTenantIDIdentityProviders", params)
}
func (x *IdentityClient) AuthPostApiV1TenantsByTenantIDIdentityProviders(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "authPostApiV1TenantsByTenantIDIdentityProviders", params)
}
func (x *IdentityClient) AuthPostApiV1TenantsByTenantIDIdentityProvidersByProviderIDDisable(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "authPostApiV1TenantsByTenantIDIdentityProvidersByProviderIDDisable", params)
}
func (x *IdentityClient) AuthPostApiV1TenantsByTenantIDIdentityProvidersByProviderIDEnable(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "authPostApiV1TenantsByTenantIDIdentityProvidersByProviderIDEnable", params)
}
func (x *IdentityClient) AuthGetAuthByProviderIDStart(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "authGetAuthByProviderIDStart", params)
}
func (x *IdentityClient) IdentityGetAuthGithub(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "identityGetAuthGithub", params)
}
func (x *IdentityClient) IdentityGetAuthGithubCallback(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "identityGetAuthGithubCallback", params)
}
func (x *IdentityClient) AuthGetAuthGoogleOauth2Callback(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "authGetAuthGoogleOauth2Callback", params)
}
func (x *IdentityClient) AuthGetAuthGoogleOauth2Start(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "authGetAuthGoogleOauth2Start", params)
}
func (x *IdentityClient) AuthPostAuthLogout(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "authPostAuthLogout", params)
}
func (x *IdentityClient) AuthGetAuthOidcCallback(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "authGetAuthOidcCallback", params)
}
func (x *IdentityClient) AuthGetAuthProviders(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "authGetAuthProviders", params)
}
func (x *IdentityClient) AuthPostAuthRefresh(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "authPostAuthRefresh", params)
}
func (x *IdentityClient) AuthGetAuthSilentCallback(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "authGetAuthSilentCallback", params)
}
func (x *IdentityClient) AuthGetAuthSilentStart(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "authGetAuthSilentStart", params)
}
func (x *IdentityClient) AuthGetOauthAuthorize(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "authGetOauthAuthorize", params)
}
func (x *IdentityClient) AuthPostOauthAuthorizeTenant(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "authPostOauthAuthorizeTenant", params)
}
func (x *IdentityClient) AuthPostOauthDeviceAuthorization(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "authPostOauthDeviceAuthorization", params)
}
func (x *IdentityClient) AuthPostOauthDeviceAuthorize(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "authPostOauthDeviceAuthorize", params)
}
func (x *IdentityClient) AuthGetOauthDeviceLookup(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "authGetOauthDeviceLookup", params)
}
func (x *IdentityClient) AuthPostOauthRegister(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "authPostOauthRegister", params)
}
func (x *IdentityClient) AuthPostOauthRevoke(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "authPostOauthRevoke", params)
}
func (x *IdentityClient) AuthPostOauthToken(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "authPostOauthToken", params)
}
func (x *IdentityClient) AuthGetOauthUserinfo(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "authGetOauthUserinfo", params)
}

type TenantsClient struct{ client *Client }
func (c *Client) Tenants() *TenantsClient { return &TenantsClient{client: c} }

func (x *TenantsClient) TenantsGetApiV1InviteLinksByToken(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "tenantsGetApiV1InviteLinksByToken", params)
}
func (x *TenantsClient) TenantsPostApiV1InviteLinksByTokenAccept(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "tenantsPostApiV1InviteLinksByTokenAccept", params)
}
func (x *TenantsClient) TenantsGetApiV1Tenants(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "tenantsGetApiV1Tenants", params)
}
func (x *TenantsClient) TenantsPostApiV1Tenants(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "tenantsPostApiV1Tenants", params)
}
func (x *TenantsClient) TenantsDeleteApiV1TenantsByTenantID(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "tenantsDeleteApiV1TenantsByTenantID", params)
}
func (x *TenantsClient) TenantsGetApiV1TenantsByTenantID(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "tenantsGetApiV1TenantsByTenantID", params)
}
func (x *TenantsClient) TenantsPatchApiV1TenantsByTenantID(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "tenantsPatchApiV1TenantsByTenantID", params)
}
func (x *TenantsClient) TenantsGetApiV1TenantsByTenantIDEmailDomains(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "tenantsGetApiV1TenantsByTenantIDEmailDomains", params)
}
func (x *TenantsClient) TenantsPostApiV1TenantsByTenantIDEmailDomains(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "tenantsPostApiV1TenantsByTenantIDEmailDomains", params)
}
func (x *TenantsClient) TenantsDeleteApiV1TenantsByTenantIDEmailDomainsByDomain(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "tenantsDeleteApiV1TenantsByTenantIDEmailDomainsByDomain", params)
}
func (x *TenantsClient) TenantsGetApiV1TenantsByTenantIDGroups(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "tenantsGetApiV1TenantsByTenantIDGroups", params)
}
func (x *TenantsClient) TenantsPostApiV1TenantsByTenantIDGroups(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "tenantsPostApiV1TenantsByTenantIDGroups", params)
}
func (x *TenantsClient) TenantsDeleteApiV1TenantsByTenantIDGroupsByGroupID(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "tenantsDeleteApiV1TenantsByTenantIDGroupsByGroupID", params)
}
func (x *TenantsClient) TenantsPatchApiV1TenantsByTenantIDGroupsByGroupID(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "tenantsPatchApiV1TenantsByTenantIDGroupsByGroupID", params)
}
func (x *TenantsClient) TenantsDeleteApiV1TenantsByTenantIDIcon(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "tenantsDeleteApiV1TenantsByTenantIDIcon", params)
}
func (x *TenantsClient) TenantsPostApiV1TenantsByTenantIDIconUploadUrl(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "tenantsPostApiV1TenantsByTenantIDIconUploadUrl", params)
}
func (x *TenantsClient) TenantsGetApiV1TenantsByTenantIDInvitations(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "tenantsGetApiV1TenantsByTenantIDInvitations", params)
}
func (x *TenantsClient) TenantsPostApiV1TenantsByTenantIDInvitations(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "tenantsPostApiV1TenantsByTenantIDInvitations", params)
}
func (x *TenantsClient) TenantsDeleteApiV1TenantsByTenantIDInvitationsByInvitationID(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "tenantsDeleteApiV1TenantsByTenantIDInvitationsByInvitationID", params)
}
func (x *TenantsClient) TenantsPostApiV1TenantsByTenantIDInvitationsBulk(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "tenantsPostApiV1TenantsByTenantIDInvitationsBulk", params)
}
func (x *TenantsClient) TenantsGetApiV1TenantsByTenantIDInviteLinks(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "tenantsGetApiV1TenantsByTenantIDInviteLinks", params)
}
func (x *TenantsClient) TenantsPostApiV1TenantsByTenantIDInviteLinks(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "tenantsPostApiV1TenantsByTenantIDInviteLinks", params)
}
func (x *TenantsClient) TenantsDeleteApiV1TenantsByTenantIDInviteLinksByLinkID(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "tenantsDeleteApiV1TenantsByTenantIDInviteLinksByLinkID", params)
}
func (x *TenantsClient) TenantsGetApiV1TenantsByTenantIDMembers(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "tenantsGetApiV1TenantsByTenantIDMembers", params)
}
func (x *TenantsClient) TenantsPatchApiV1TenantsByTenantIDMembersByMembershipID(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "tenantsPatchApiV1TenantsByTenantIDMembersByMembershipID", params)
}
func (x *TenantsClient) TenantsPostApiV1TenantsByTenantIDMembersByMembershipIDDeactivate(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "tenantsPostApiV1TenantsByTenantIDMembersByMembershipIDDeactivate", params)
}
func (x *TenantsClient) TenantsPatchApiV1TenantsByTenantIDMembersByMembershipIDGroup(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "tenantsPatchApiV1TenantsByTenantIDMembersByMembershipIDGroup", params)
}
func (x *TenantsClient) TenantsPostApiV1TenantsByTenantIDMembersByMembershipIDReactivate(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "tenantsPostApiV1TenantsByTenantIDMembersByMembershipIDReactivate", params)
}
func (x *TenantsClient) TenantsPostApiV1TenantsByTenantIDMembersByMembershipIDRestoreScim(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "tenantsPostApiV1TenantsByTenantIDMembersByMembershipIDRestoreScim", params)
}
func (x *TenantsClient) TenantsDeleteApiV1TenantsByTenantIDMembersByUserIDSeat(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "tenantsDeleteApiV1TenantsByTenantIDMembersByUserIDSeat", params)
}
func (x *TenantsClient) TenantsPostApiV1TenantsByTenantIDMembersByUserIDSeat(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "tenantsPostApiV1TenantsByTenantIDMembersByUserIDSeat", params)
}
func (x *TenantsClient) TenantsGetApiV1TenantsByTenantIDMembersDirectory(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "tenantsGetApiV1TenantsByTenantIDMembersDirectory", params)
}
func (x *TenantsClient) TenantsGetApiV1TenantsByTenantIDScimConnection(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "tenantsGetApiV1TenantsByTenantIDScimConnection", params)
}
func (x *TenantsClient) TenantsDeleteApiV1TenantsByTenantIDScimToken(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "tenantsDeleteApiV1TenantsByTenantIDScimToken", params)
}
func (x *TenantsClient) TenantsPostApiV1TenantsByTenantIDScimToken(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "tenantsPostApiV1TenantsByTenantIDScimToken", params)
}

type AuthzClient struct{ client *Client }
func (c *Client) Authz() *AuthzClient { return &AuthzClient{client: c} }

func (x *AuthzClient) AuthorizationGetApiV1TenantsByTenantIDGrants(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "authorizationGetApiV1TenantsByTenantIDGrants", params)
}
func (x *AuthzClient) AuthorizationPostApiV1TenantsByTenantIDGrants(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "authorizationPostApiV1TenantsByTenantIDGrants", params)
}
func (x *AuthzClient) AuthorizationDeleteApiV1TenantsByTenantIDGrantsByGrantID(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "authorizationDeleteApiV1TenantsByTenantIDGrantsByGrantID", params)
}
func (x *AuthzClient) AuthorizationGetApiV1TenantsByTenantIDGrantsByGrantID(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "authorizationGetApiV1TenantsByTenantIDGrantsByGrantID", params)
}
func (x *AuthzClient) AuthorizationGetApiV1TenantsByTenantIDMeGrants(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "authorizationGetApiV1TenantsByTenantIDMeGrants", params)
}
func (x *AuthzClient) AuthorizationGetApiV1TenantsByTenantIDPresets(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "authorizationGetApiV1TenantsByTenantIDPresets", params)
}
func (x *AuthzClient) AuthorizationPostApiV1TenantsByTenantIDPresets(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "authorizationPostApiV1TenantsByTenantIDPresets", params)
}
func (x *AuthzClient) AuthorizationDeleteApiV1TenantsByTenantIDPresetsByPresetID(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "authorizationDeleteApiV1TenantsByTenantIDPresetsByPresetID", params)
}
func (x *AuthzClient) AuthorizationGetApiV1TenantsByTenantIDPresetsByPresetID(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "authorizationGetApiV1TenantsByTenantIDPresetsByPresetID", params)
}
func (x *AuthzClient) AuthorizationPatchApiV1TenantsByTenantIDPresetsByPresetID(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "authorizationPatchApiV1TenantsByTenantIDPresetsByPresetID", params)
}
func (x *AuthzClient) AuthorizationGetApiV1TenantsByTenantIDSubjects(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "authorizationGetApiV1TenantsByTenantIDSubjects", params)
}
func (x *AuthzClient) AuthorizationPostApiV1TenantsByTenantIDSubjects(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "authorizationPostApiV1TenantsByTenantIDSubjects", params)
}
func (x *AuthzClient) AuthorizationGetApiV1TenantsByTenantIDSubjectsBySubjectID(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "authorizationGetApiV1TenantsByTenantIDSubjectsBySubjectID", params)
}

type AuditClient struct{ client *Client }
func (c *Client) Audit() *AuditClient { return &AuditClient{client: c} }

func (x *AuditClient) AuditGetApiV1TenantsByTenantIDAuditEvents(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "auditGetApiV1TenantsByTenantIDAuditEvents", params)
}
func (x *AuditClient) AuditGetApiV1TenantsByTenantIDAuditEventsByEventID(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "auditGetApiV1TenantsByTenantIDAuditEventsByEventID", params)
}
func (x *AuditClient) AuditPostApiV1TenantsByTenantIDAuditEventsAnonymize(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "auditPostApiV1TenantsByTenantIDAuditEventsAnonymize", params)
}
func (x *AuditClient) AuditGetApiV1TenantsByTenantIDAuditEventsIntegrityCheck(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "auditGetApiV1TenantsByTenantIDAuditEventsIntegrityCheck", params)
}

type GatewayClient struct{ client *Client }
func (c *Client) Gateway() *GatewayClient { return &GatewayClient{client: c} }

func (x *GatewayClient) GatewayGetApiV1TenantsByTenantIDConnectors(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "gatewayGetApiV1TenantsByTenantIDConnectors", params)
}
func (x *GatewayClient) GatewayPostApiV1TenantsByTenantIDConnectors(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "gatewayPostApiV1TenantsByTenantIDConnectors", params)
}
func (x *GatewayClient) GatewayDeleteApiV1TenantsByTenantIDConnectorsByConnectorID(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "gatewayDeleteApiV1TenantsByTenantIDConnectorsByConnectorID", params)
}
func (x *GatewayClient) GatewayGetApiV1TenantsByTenantIDConnectorsByConnectorID(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "gatewayGetApiV1TenantsByTenantIDConnectorsByConnectorID", params)
}
func (x *GatewayClient) GatewayPatchApiV1TenantsByTenantIDConnectorsByConnectorID(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "gatewayPatchApiV1TenantsByTenantIDConnectorsByConnectorID", params)
}
func (x *GatewayClient) GatewayPostApiV1TenantsByTenantIDConnectorsByConnectorIDDiscover(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "gatewayPostApiV1TenantsByTenantIDConnectorsByConnectorIDDiscover", params)
}
func (x *GatewayClient) GatewayGetApiV1TenantsByTenantIDConnectorsByConnectorIDResources(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "gatewayGetApiV1TenantsByTenantIDConnectorsByConnectorIDResources", params)
}
func (x *GatewayClient) GatewayPostApiV1TenantsByTenantIDConnectorsByConnectorIDTestConnection(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "gatewayPostApiV1TenantsByTenantIDConnectorsByConnectorIDTestConnection", params)
}
func (x *GatewayClient) GatewayPostApiV1TenantsByTenantIDConnectorsOauthGoogleFinalize(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "gatewayPostApiV1TenantsByTenantIDConnectorsOauthGoogleFinalize", params)
}
func (x *GatewayClient) GatewayPostApiV1TenantsByTenantIDConnectorsOauthGoogleStart(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "gatewayPostApiV1TenantsByTenantIDConnectorsOauthGoogleStart", params)
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

type CostClient struct{ client *Client }
func (c *Client) Cost() *CostClient { return &CostClient{client: c} }

func (x *CostClient) CostGetApiV1TenantsByTenantIDCostByApp(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "costGetApiV1TenantsByTenantIDCostByApp", params)
}
func (x *CostClient) CostGetApiV1TenantsByTenantIDCostByCostCenter(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "costGetApiV1TenantsByTenantIDCostByCostCenter", params)
}
func (x *CostClient) CostGetApiV1TenantsByTenantIDCostExport(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "costGetApiV1TenantsByTenantIDCostExport", params)
}
func (x *CostClient) CostGetApiV1TenantsByTenantIDCostMonths(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "costGetApiV1TenantsByTenantIDCostMonths", params)
}
func (x *CostClient) CostGetApiV1TenantsByTenantIDCostSummary(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "costGetApiV1TenantsByTenantIDCostSummary", params)
}
func (x *CostClient) CostGetApiV1TenantsByTenantIDCostTimeseries(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "costGetApiV1TenantsByTenantIDCostTimeseries", params)
}
func (x *CostClient) CostGetApiV1TenantsByTenantIDInfraAppsByAppIDUsageSeries(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "costGetApiV1TenantsByTenantIDInfraAppsByAppIDUsageSeries", params)
}
func (x *CostClient) CostGetApiV1TenantsByTenantIDInfraUsage(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "costGetApiV1TenantsByTenantIDInfraUsage", params)
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
func (x *DeploymentsClient) DeployPostWebhooksGithub(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "deployPostWebhooksGithub", params)
}

