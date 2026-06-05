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
func (x *TenantsClient) TenantsGetApiV1TenantsByTenantIDMembers(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "tenantsGetApiV1TenantsByTenantIDMembers", params)
}
func (x *TenantsClient) TenantsPatchApiV1TenantsByTenantIDMembersByMembershipID(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "tenantsPatchApiV1TenantsByTenantIDMembersByMembershipID", params)
}
func (x *TenantsClient) TenantsPostApiV1TenantsByTenantIDMembersByMembershipIDDeactivate(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "tenantsPostApiV1TenantsByTenantIDMembersByMembershipIDDeactivate", params)
}
func (x *TenantsClient) TenantsPostApiV1TenantsByTenantIDMembersByMembershipIDReactivate(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "tenantsPostApiV1TenantsByTenantIDMembersByMembershipIDReactivate", params)
}
func (x *TenantsClient) TenantsGetTenantsByTenantIDEmailDomains(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "tenantsGetTenantsByTenantIDEmailDomains", params)
}
func (x *TenantsClient) TenantsPostTenantsByTenantIDEmailDomains(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "tenantsPostTenantsByTenantIDEmailDomains", params)
}
func (x *TenantsClient) TenantsDeleteTenantsByTenantIDEmailDomainsByDomain(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "tenantsDeleteTenantsByTenantIDEmailDomainsByDomain", params)
}

type AuthzClient struct{ client *Client }
func (c *Client) Authz() *AuthzClient { return &AuthzClient{client: c} }

func (x *AuthzClient) AuthorizationGetApiV1TenantsByTenantIDGrants(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "authorizationGetApiV1TenantsByTenantIDGrants", params)
}
func (x *AuthzClient) AuthorizationPostApiV1TenantsByTenantIDGrantsByGrantIDGrant(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "authorizationPostApiV1TenantsByTenantIDGrantsByGrantIDGrant", params)
}
func (x *AuthzClient) AuthorizationPostApiV1TenantsByTenantIDGrantsByGrantIDRevoke(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "authorizationPostApiV1TenantsByTenantIDGrantsByGrantIDRevoke", params)
}
func (x *AuthzClient) AuthorizationGetApiV1TenantsByTenantIDSubjects(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "authorizationGetApiV1TenantsByTenantIDSubjects", params)
}
func (x *AuthzClient) AuthorizationPostApiV1TenantsByTenantIDSubjects(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "authorizationPostApiV1TenantsByTenantIDSubjects", params)
}
func (x *AuthzClient) AuthorizationDeleteApiV1TenantsByTenantIDSubjectsBySubjectID(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "authorizationDeleteApiV1TenantsByTenantIDSubjectsBySubjectID", params)
}
func (x *AuthzClient) AuthorizationPatchApiV1TenantsByTenantIDSubjectsBySubjectID(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "authorizationPatchApiV1TenantsByTenantIDSubjectsBySubjectID", params)
}
func (x *AuthzClient) AuthorizationPostApiV1TenantsByTenantIDSubjectsBySubjectIDMove(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "authorizationPostApiV1TenantsByTenantIDSubjectsBySubjectIDMove", params)
}
func (x *AuthzClient) AuthorizationPostApiV1TenantsByTenantIDSubjectsBySubjectIDTags(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "authorizationPostApiV1TenantsByTenantIDSubjectsBySubjectIDTags", params)
}
func (x *AuthzClient) AuthorizationDeleteApiV1TenantsByTenantIDSubjectsBySubjectIDTagsByTagID(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "authorizationDeleteApiV1TenantsByTenantIDSubjectsBySubjectIDTagsByTagID", params)
}
func (x *AuthzClient) AuthorizationGetApiV1TenantsByTenantIDTags(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "authorizationGetApiV1TenantsByTenantIDTags", params)
}
func (x *AuthzClient) AuthorizationPostApiV1TenantsByTenantIDTags(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "authorizationPostApiV1TenantsByTenantIDTags", params)
}
func (x *AuthzClient) AuthorizationDeleteApiV1TenantsByTenantIDTagsByTagID(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "authorizationDeleteApiV1TenantsByTenantIDTagsByTagID", params)
}
func (x *AuthzClient) AuthorizationPatchApiV1TenantsByTenantIDTagsByTagID(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "authorizationPatchApiV1TenantsByTenantIDTagsByTagID", params)
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

func (x *GatewayClient) GatewayGetApiV1CatalogKinds(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "gatewayGetApiV1CatalogKinds", params)
}
func (x *GatewayClient) GatewayGetApiV1Engines(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "gatewayGetApiV1Engines", params)
}
func (x *GatewayClient) GatewayGetApiV1TenantsByTenantIDCatalogConnectors(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "gatewayGetApiV1TenantsByTenantIDCatalogConnectors", params)
}
func (x *GatewayClient) GatewayGetApiV1TenantsByTenantIDCatalogResources(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "gatewayGetApiV1TenantsByTenantIDCatalogResources", params)
}
func (x *GatewayClient) GatewayGetApiV1TenantsByTenantIDCatalogResourcesByConnectorByPath(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "gatewayGetApiV1TenantsByTenantIDCatalogResourcesByConnectorByPath", params)
}
func (x *GatewayClient) GatewayPostApiV1TenantsByTenantIDCatalogResourcesByConnectorByPath(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "gatewayPostApiV1TenantsByTenantIDCatalogResourcesByConnectorByPath", params)
}
func (x *GatewayClient) GatewayGetApiV1TenantsByTenantIDConnectors(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "gatewayGetApiV1TenantsByTenantIDConnectors", params)
}
func (x *GatewayClient) GatewayPostApiV1TenantsByTenantIDConnectors(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "gatewayPostApiV1TenantsByTenantIDConnectors", params)
}
func (x *GatewayClient) GatewayDeleteApiV1TenantsByTenantIDConnectorsByConnectorID(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "gatewayDeleteApiV1TenantsByTenantIDConnectorsByConnectorID", params)
}
func (x *GatewayClient) GatewayPatchApiV1TenantsByTenantIDConnectorsByConnectorID(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "gatewayPatchApiV1TenantsByTenantIDConnectorsByConnectorID", params)
}
func (x *GatewayClient) GatewayPostApiV1TenantsByTenantIDConnectorsByConnectorIDCredentials(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "gatewayPostApiV1TenantsByTenantIDConnectorsByConnectorIDCredentials", params)
}
func (x *GatewayClient) GatewayGetApiV1TenantsByTenantIDConnectorsByConnectorIDDiscover(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "gatewayGetApiV1TenantsByTenantIDConnectorsByConnectorIDDiscover", params)
}
func (x *GatewayClient) GatewayPostApiV1TenantsByTenantIDGatewayQuery(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "gatewayPostApiV1TenantsByTenantIDGatewayQuery", params)
}
func (x *GatewayClient) GatewayGetApiV1TenantsByTenantIDResources(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "gatewayGetApiV1TenantsByTenantIDResources", params)
}
func (x *GatewayClient) GatewayDeleteApiV1TenantsByTenantIDResourcesByResourceID(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "gatewayDeleteApiV1TenantsByTenantIDResourcesByResourceID", params)
}
func (x *GatewayClient) GatewayPatchApiV1TenantsByTenantIDResourcesByResourceID(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "gatewayPatchApiV1TenantsByTenantIDResourcesByResourceID", params)
}
func (x *GatewayClient) GatewayPostApiV1TenantsByTenantIDResourcesByResourceIDMove(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "gatewayPostApiV1TenantsByTenantIDResourcesByResourceIDMove", params)
}
func (x *GatewayClient) GatewayPostApiV1TenantsByTenantIDResourcesByResourceIDTags(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "gatewayPostApiV1TenantsByTenantIDResourcesByResourceIDTags", params)
}
func (x *GatewayClient) GatewayDeleteApiV1TenantsByTenantIDResourcesByResourceIDTagsByTagID(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "gatewayDeleteApiV1TenantsByTenantIDResourcesByResourceIDTagsByTagID", params)
}
func (x *GatewayClient) GatewayPostApiV1TenantsByTenantIDResourcesBulk(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "gatewayPostApiV1TenantsByTenantIDResourcesBulk", params)
}
func (x *GatewayClient) GatewayPostApiV1TenantsByTenantIDResourcesNamespaces(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "gatewayPostApiV1TenantsByTenantIDResourcesNamespaces", params)
}
func (x *GatewayClient) ConfigGetConfigPublic(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "configGetConfigPublic", params)
}

type DataClient struct{ client *Client }
func (c *Client) Data() *DataClient { return &DataClient{client: c} }

func (x *DataClient) SchemaGetApiV1AppsByAppIDTables(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "schemaGetApiV1AppsByAppIDTables", params)
}
func (x *DataClient) SchemaPostApiV1AppsByAppIDTables(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "schemaPostApiV1AppsByAppIDTables", params)
}
func (x *DataClient) SchemaDeleteApiV1AppsByAppIDTablesByTableName(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "schemaDeleteApiV1AppsByAppIDTablesByTableName", params)
}
func (x *DataClient) SchemaGetApiV1AppsByAppIDTablesByTableName(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "schemaGetApiV1AppsByAppIDTablesByTableName", params)
}
func (x *DataClient) SchemaPostApiV1AppsByAppIDTablesByTableNameColumns(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "schemaPostApiV1AppsByAppIDTablesByTableNameColumns", params)
}
func (x *DataClient) SchemaDeleteApiV1AppsByAppIDTablesByTableNameColumnsByColumnName(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "schemaDeleteApiV1AppsByAppIDTablesByTableNameColumnsByColumnName", params)
}
func (x *DataClient) SchemaGetApiV1AppsByAppIDTablesByTableNameGrants(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "schemaGetApiV1AppsByAppIDTablesByTableNameGrants", params)
}
func (x *DataClient) SchemaPostApiV1AppsByAppIDTablesByTableNameGrants(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "schemaPostApiV1AppsByAppIDTablesByTableNameGrants", params)
}
func (x *DataClient) SchemaDeleteApiV1AppsByAppIDTablesByTableNameGrantsByGrantID(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "schemaDeleteApiV1AppsByAppIDTablesByTableNameGrantsByGrantID", params)
}
func (x *DataClient) SchemaGetApiV1AppsByAppIDTablesByTableNameRows(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "schemaGetApiV1AppsByAppIDTablesByTableNameRows", params)
}
func (x *DataClient) SchemaGetApiV1AppsByAppIDTablesCheckAvailability(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "schemaGetApiV1AppsByAppIDTablesCheckAvailability", params)
}
func (x *DataClient) SchemaGetApiV1AppsByAppIDTablesColumnTypes(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "schemaGetApiV1AppsByAppIDTablesColumnTypes", params)
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
func (x *DataClient) SchemaGetDataByTenantSlugByAppSlugByTable(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "schemaGetDataByTenantSlugByAppSlugByTable", params)
}
func (x *DataClient) SchemaPostDataByTenantSlugByAppSlugByTable(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "schemaPostDataByTenantSlugByAppSlugByTable", params)
}
func (x *DataClient) SchemaGetDataByTenantSlugByAppSlugByTableCount(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "schemaGetDataByTenantSlugByAppSlugByTableCount", params)
}
func (x *DataClient) SchemaDeleteDataByTenantSlugByAppSlugByTableById(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "schemaDeleteDataByTenantSlugByAppSlugByTableById", params)
}
func (x *DataClient) SchemaGetDataByTenantSlugByAppSlugByTableById(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "schemaGetDataByTenantSlugByAppSlugByTableById", params)
}
func (x *DataClient) SchemaPatchDataByTenantSlugByAppSlugByTableById(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "schemaPatchDataByTenantSlugByAppSlugByTableById", params)
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
func (x *DeploymentsClient) DeployGetApiV1AppsByAppIDGitGithubInstallStart(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "deployGetApiV1AppsByAppIDGitGithubInstallStart", params)
}
func (x *DeploymentsClient) DeployGetApiV1AppsByAppIDLogs(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "deployGetApiV1AppsByAppIDLogs", params)
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
func (x *DeploymentsClient) DeployPostWebhooksGithub(ctx context.Context, params OperationParams) (map[string]any, error) {
	return x.client.Operation(ctx, "deployPostWebhooksGithub", params)
}

