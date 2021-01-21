package rpc

// GitpodServerInterface wraps the Gitpod server
type APIInterface interface {
	GetLoggedInUser(ctx context.Context) (res *User, err error)
	GetTerms(ctx context.Context) (res *Terms, err error)
	UpdateLoggedInUser(ctx context.Context, user *User) (res *User, err error)
	GetAuthProviders(ctx context.Context) (res []*AuthProviderInfo, err error)
	GetOwnAuthProviders(ctx context.Context) (res []*AuthProviderEntry, err error)
	UpdateOwnAuthProvider(ctx context.Context, params *GitpodServerUpdateOwnAuthProviderParams) (err error)
	DeleteOwnAuthProvider(ctx context.Context, params *GitpodServerDeleteOwnAuthProviderParams) (err error)
	GetBranding(ctx context.Context) (res *Branding, err error)
	GetConfiguration(ctx context.Context) (res *Configuration, err error)
	GetToken(ctx context.Context, query *GitpodServerGetTokenSearchOptions) (res *Token, err error)
	GetPortAuthenticationToken(ctx context.Context, workspaceId string) (res *Token, err error)
	DeleteAccount(ctx context.Context) (err error)
	GetClientRegion(ctx context.Context) (res string, err error)
	HasPermission(ctx context.Context, permission *PermissionName) (res bool, err error)
	GetWorkspaces(ctx context.Context, options *GitpodServerGetWorkspacesOptions) (res []*WorkspaceInfo, err error)
	GetWorkspaceOwner(ctx context.Context, workspaceId string) (res *UserInfo, err error)
	GetWorkspaceUsers(ctx context.Context, workspaceId string) (res []*WorkspaceInstanceUser, err error)
	GetFeaturedRepositories(ctx context.Context) (res []*WhitelistedRepository, err error)
	GetWorkspace(ctx context.Context, id string) (res *WorkspaceInfo, err error)
	IsWorkspaceOwner(ctx context.Context, workspaceId string) (res bool, err error)
	CreateWorkspace(ctx context.Context, options *GitpodServerCreateWorkspaceOptions) (res *WorkspaceCreationResult, err error)
	StartWorkspace(ctx context.Context, id string, options *GitpodServerStartWorkspaceOptions) (res *StartWorkspaceResult, err error)
	StopWorkspace(ctx context.Context, id string) (err error)
	DeleteWorkspace(ctx context.Context, id string) (err error)
	SetWorkspaceDescription(ctx context.Context, id string, desc string) (err error)
	ControlAdmission(ctx context.Context, id string, level *GitpodServerAdmissionLevel) (err error)
	UpdateWorkspaceUserPin(ctx context.Context, id string, action *GitpodServerPinAction) (err error)
	SendHeartBeat(ctx context.Context, options *GitpodServerSendHeartBeatOptions) (err error)
	WatchWorkspaceImageBuildLogs(ctx context.Context, workspaceId string) (err error)
	WatchHeadlessWorkspaceLogs(ctx context.Context, workspaceId string) (err error)
	IsPrebuildDone(ctx context.Context, pwsid string) (res bool, err error)
	SetWorkspaceTimeout(ctx context.Context, workspaceId string, duration *WorkspaceTimeoutDuration) (res *SetWorkspaceTimeoutResult, err error)
	GetWorkspaceTimeout(ctx context.Context, workspaceId string) (res *GetWorkspaceTimeoutResult, err error)
	SendHeartBeat(ctx context.Context, options *GitpodServerSendHeartBeatOptions) (err error)
	UpdateWorkspaceUserPin(ctx context.Context, id string, action *GitpodServerPinAction) (err error)
	GetOpenPorts(ctx context.Context, workspaceId string) (res []*WorkspaceInstancePort, err error)
	OpenPort(ctx context.Context, workspaceId string, port *WorkspaceInstancePort) (res *WorkspaceInstancePort, err error)
	ClosePort(ctx context.Context, workspaceId string, port float32) (err error)
	GetUserMessages(ctx context.Context, options *GitpodServerGetUserMessagesOptions) (res []*UserMessage, err error)
	UpdateUserMessages(ctx context.Context, options *GitpodServerUpdateUserMessagesOptions) (err error)
	GetUserStorageResource(ctx context.Context, options *GitpodServerGetUserStorageResourceOptions) (res string, err error)
	UpdateUserStorageResource(ctx context.Context, options *GitpodServerUpdateUserStorageResourceOptions) (err error)
	GetEnvVars(ctx context.Context) (res []*UserEnvVarValue, err error)
	SetEnvVar(ctx context.Context, variable *UserEnvVarValue) (err error)
	DeleteEnvVar(ctx context.Context, variable *UserEnvVarValue) (err error)
	GetContentBlobUploadUrl(ctx context.Context, name string, mediaType string, timeToLive string) (res string, err error)
	GetContentBlobDownloadUrl(ctx context.Context, name string, mediaType string, timeToLive string) (res string, err error)
	GetGitpodTokens(ctx context.Context) (res []*GitpodToken, err error)
	GenerateNewGitpodToken(ctx context.Context, options *GitpodServerGenerateNewGitpodTokenOptions) (res string, err error)
	DeleteGitpodToken(ctx context.Context, tokenHash string) (err error)
	SendFeedback(ctx context.Context, feedback string) (res string, err error)
	RegisterGithubApp(ctx context.Context, installationId string) (err error)
	TakeSnapshot(ctx context.Context, options *GitpodServerTakeSnapshotOptions) (res string, err error)
	GetSnapshots(ctx context.Context, workspaceID string) (res []*string, err error)
	StoreLayout(ctx context.Context, workspaceId string, layoutData string) (err error)
	GetLayout(ctx context.Context, workspaceId string) (res string, err error)
	PreparePluginUpload(ctx context.Context, params *PreparePluginUploadParams) (res string, err error)
	ResolvePlugins(ctx context.Context, workspaceId string, params *ResolvePluginsParams) (res *ResolvedPlugins, err error)
	InstallUserPlugins(ctx context.Context, params *InstallPluginsParams) (res bool, err error)
	UninstallUserPlugin(ctx context.Context, params *UninstallPluginParams) (res bool, err error)
}

// FunctionName is the name of an RPC function
type FunctionName string

const (
	// FunctionGetLoggedInUser is the name of the getLoggedInUser function
	FunctionGetLoggedInUser FunctionName = "getLoggedInUser"
	// FunctionGetTerms is the name of the getTerms function
	FunctionGetTerms FunctionName = "getTerms"
	// FunctionUpdateLoggedInUser is the name of the updateLoggedInUser function
	FunctionUpdateLoggedInUser FunctionName = "updateLoggedInUser"
	// FunctionGetAuthProviders is the name of the getAuthProviders function
	FunctionGetAuthProviders FunctionName = "getAuthProviders"
	// FunctionGetOwnAuthProviders is the name of the getOwnAuthProviders function
	FunctionGetOwnAuthProviders FunctionName = "getOwnAuthProviders"
	// FunctionUpdateOwnAuthProvider is the name of the updateOwnAuthProvider function
	FunctionUpdateOwnAuthProvider FunctionName = "updateOwnAuthProvider"
	// FunctionDeleteOwnAuthProvider is the name of the deleteOwnAuthProvider function
	FunctionDeleteOwnAuthProvider FunctionName = "deleteOwnAuthProvider"
	// FunctionGetBranding is the name of the getBranding function
	FunctionGetBranding FunctionName = "getBranding"
	// FunctionGetConfiguration is the name of the getConfiguration function
	FunctionGetConfiguration FunctionName = "getConfiguration"
	// FunctionGetToken is the name of the getToken function
	FunctionGetToken FunctionName = "getToken"
	// FunctionGetPortAuthenticationToken is the name of the getPortAuthenticationToken function
	FunctionGetPortAuthenticationToken FunctionName = "getPortAuthenticationToken"
	// FunctionDeleteAccount is the name of the deleteAccount function
	FunctionDeleteAccount FunctionName = "deleteAccount"
	// FunctionGetClientRegion is the name of the getClientRegion function
	FunctionGetClientRegion FunctionName = "getClientRegion"
	// FunctionHasPermission is the name of the hasPermission function
	FunctionHasPermission FunctionName = "hasPermission"
	// FunctionGetWorkspaces is the name of the getWorkspaces function
	FunctionGetWorkspaces FunctionName = "getWorkspaces"
	// FunctionGetWorkspaceOwner is the name of the getWorkspaceOwner function
	FunctionGetWorkspaceOwner FunctionName = "getWorkspaceOwner"
	// FunctionGetWorkspaceUsers is the name of the getWorkspaceUsers function
	FunctionGetWorkspaceUsers FunctionName = "getWorkspaceUsers"
	// FunctionGetFeaturedRepositories is the name of the getFeaturedRepositories function
	FunctionGetFeaturedRepositories FunctionName = "getFeaturedRepositories"
	// FunctionGetWorkspace is the name of the getWorkspace function
	FunctionGetWorkspace FunctionName = "getWorkspace"
	// FunctionIsWorkspaceOwner is the name of the isWorkspaceOwner function
	FunctionIsWorkspaceOwner FunctionName = "isWorkspaceOwner"
	// FunctionCreateWorkspace is the name of the createWorkspace function
	FunctionCreateWorkspace FunctionName = "createWorkspace"
	// FunctionStartWorkspace is the name of the startWorkspace function
	FunctionStartWorkspace FunctionName = "startWorkspace"
	// FunctionStopWorkspace is the name of the stopWorkspace function
	FunctionStopWorkspace FunctionName = "stopWorkspace"
	// FunctionDeleteWorkspace is the name of the deleteWorkspace function
	FunctionDeleteWorkspace FunctionName = "deleteWorkspace"
	// FunctionSetWorkspaceDescription is the name of the setWorkspaceDescription function
	FunctionSetWorkspaceDescription FunctionName = "setWorkspaceDescription"
	// FunctionControlAdmission is the name of the controlAdmission function
	FunctionControlAdmission FunctionName = "controlAdmission"
	// FunctionUpdateWorkspaceUserPin is the name of the updateWorkspaceUserPin function
	FunctionUpdateWorkspaceUserPin FunctionName = "updateWorkspaceUserPin"
	// FunctionSendHeartBeat is the name of the sendHeartBeat function
	FunctionSendHeartBeat FunctionName = "sendHeartBeat"
	// FunctionWatchWorkspaceImageBuildLogs is the name of the watchWorkspaceImageBuildLogs function
	FunctionWatchWorkspaceImageBuildLogs FunctionName = "watchWorkspaceImageBuildLogs"
	// FunctionWatchHeadlessWorkspaceLogs is the name of the watchHeadlessWorkspaceLogs function
	FunctionWatchHeadlessWorkspaceLogs FunctionName = "watchHeadlessWorkspaceLogs"
	// FunctionIsPrebuildDone is the name of the isPrebuildDone function
	FunctionIsPrebuildDone FunctionName = "isPrebuildDone"
	// FunctionSetWorkspaceTimeout is the name of the setWorkspaceTimeout function
	FunctionSetWorkspaceTimeout FunctionName = "setWorkspaceTimeout"
	// FunctionGetWorkspaceTimeout is the name of the getWorkspaceTimeout function
	FunctionGetWorkspaceTimeout FunctionName = "getWorkspaceTimeout"
	// FunctionSendHeartBeat is the name of the sendHeartBeat function
	FunctionSendHeartBeat FunctionName = "sendHeartBeat"
	// FunctionUpdateWorkspaceUserPin is the name of the updateWorkspaceUserPin function
	FunctionUpdateWorkspaceUserPin FunctionName = "updateWorkspaceUserPin"
	// FunctionGetOpenPorts is the name of the getOpenPorts function
	FunctionGetOpenPorts FunctionName = "getOpenPorts"
	// FunctionOpenPort is the name of the openPort function
	FunctionOpenPort FunctionName = "openPort"
	// FunctionClosePort is the name of the closePort function
	FunctionClosePort FunctionName = "closePort"
	// FunctionGetUserMessages is the name of the getUserMessages function
	FunctionGetUserMessages FunctionName = "getUserMessages"
	// FunctionUpdateUserMessages is the name of the updateUserMessages function
	FunctionUpdateUserMessages FunctionName = "updateUserMessages"
	// FunctionGetUserStorageResource is the name of the getUserStorageResource function
	FunctionGetUserStorageResource FunctionName = "getUserStorageResource"
	// FunctionUpdateUserStorageResource is the name of the updateUserStorageResource function
	FunctionUpdateUserStorageResource FunctionName = "updateUserStorageResource"
	// FunctionGetEnvVars is the name of the getEnvVars function
	FunctionGetEnvVars FunctionName = "getEnvVars"
	// FunctionSetEnvVar is the name of the setEnvVar function
	FunctionSetEnvVar FunctionName = "setEnvVar"
	// FunctionDeleteEnvVar is the name of the deleteEnvVar function
	FunctionDeleteEnvVar FunctionName = "deleteEnvVar"
	// FunctionGetContentBlobUploadUrl is the name of the getContentBlobUploadUrl function
	FunctionGetContentBlobUploadUrl FunctionName = "getContentBlobUploadUrl"
	// FunctionGetContentBlobDownloadUrl is the name of the getContentBlobDownloadUrl function
	FunctionGetContentBlobDownloadUrl FunctionName = "getContentBlobDownloadUrl"
	// FunctionGetGitpodTokens is the name of the getGitpodTokens function
	FunctionGetGitpodTokens FunctionName = "getGitpodTokens"
	// FunctionGenerateNewGitpodToken is the name of the generateNewGitpodToken function
	FunctionGenerateNewGitpodToken FunctionName = "generateNewGitpodToken"
	// FunctionDeleteGitpodToken is the name of the deleteGitpodToken function
	FunctionDeleteGitpodToken FunctionName = "deleteGitpodToken"
	// FunctionSendFeedback is the name of the sendFeedback function
	FunctionSendFeedback FunctionName = "sendFeedback"
	// FunctionRegisterGithubApp is the name of the registerGithubApp function
	FunctionRegisterGithubApp FunctionName = "registerGithubApp"
	// FunctionTakeSnapshot is the name of the takeSnapshot function
	FunctionTakeSnapshot FunctionName = "takeSnapshot"
	// FunctionGetSnapshots is the name of the getSnapshots function
	FunctionGetSnapshots FunctionName = "getSnapshots"
	// FunctionStoreLayout is the name of the storeLayout function
	FunctionStoreLayout FunctionName = "storeLayout"
	// FunctionGetLayout is the name of the getLayout function
	FunctionGetLayout FunctionName = "getLayout"
	// FunctionPreparePluginUpload is the name of the preparePluginUpload function
	FunctionPreparePluginUpload FunctionName = "preparePluginUpload"
	// FunctionResolvePlugins is the name of the resolvePlugins function
	FunctionResolvePlugins FunctionName = "resolvePlugins"
	// FunctionInstallUserPlugins is the name of the installUserPlugins function
	FunctionInstallUserPlugins FunctionName = "installUserPlugins"
	// FunctionUninstallUserPlugin is the name of the uninstallUserPlugin function
	FunctionUninstallUserPlugin FunctionName = "uninstallUserPlugin"
)

// APIoverJSONRPC makes JSON RPC calls to the Gitpod server
type APIoverJSONRPC struct {
	C *jsonrpc2.Conn
}

// GetLoggedInUser calls getLoggedInUser on the server
func (gp *APIoverJSONRPC) GetLoggedInUser(ctx context.Context) (res *User, err error) {
	var _params []interface{}

	var result User
	err = gp.C.Call(ctx, "getLoggedInUser", _params, &result)
	if err != nil {
		return
	}
	res = &result

	return
}

// GetTerms calls getTerms on the server
func (gp *APIoverJSONRPC) GetTerms(ctx context.Context) (res *Terms, err error) {
	var _params []interface{}

	var result Terms
	err = gp.C.Call(ctx, "getTerms", _params, &result)
	if err != nil {
		return
	}
	res = &result

	return
}

// UpdateLoggedInUser calls updateLoggedInUser on the server
func (gp *APIoverJSONRPC) UpdateLoggedInUser(ctx context.Context, user *User) (res *User, err error) {
	var _params []interface{}

	_params = append(_params, user)

	var result User
	err = gp.C.Call(ctx, "updateLoggedInUser", _params, &result)
	if err != nil {
		return
	}
	res = &result

	return
}

// GetAuthProviders calls getAuthProviders on the server
func (gp *APIoverJSONRPC) GetAuthProviders(ctx context.Context) (res []*AuthProviderInfo, err error) {
	var _params []interface{}

	var result []*AuthProviderInfo
	err = gp.C.Call(ctx, "getAuthProviders", _params, &result)
	if err != nil {
		return
	}
	res = result

	return
}

// GetOwnAuthProviders calls getOwnAuthProviders on the server
func (gp *APIoverJSONRPC) GetOwnAuthProviders(ctx context.Context) (res []*AuthProviderEntry, err error) {
	var _params []interface{}

	var result []*AuthProviderEntry
	err = gp.C.Call(ctx, "getOwnAuthProviders", _params, &result)
	if err != nil {
		return
	}
	res = result

	return
}

// UpdateOwnAuthProvider calls updateOwnAuthProvider on the server
func (gp *APIoverJSONRPC) UpdateOwnAuthProvider(ctx context.Context, params *GitpodServerUpdateOwnAuthProviderParams) (err error) {
	var _params []interface{}

	_params = append(_params, params)

	err = gp.C.Call(ctx, "updateOwnAuthProvider", _params, nil)
	if err != nil {
		return
	}

	return
}

// DeleteOwnAuthProvider calls deleteOwnAuthProvider on the server
func (gp *APIoverJSONRPC) DeleteOwnAuthProvider(ctx context.Context, params *GitpodServerDeleteOwnAuthProviderParams) (err error) {
	var _params []interface{}

	_params = append(_params, params)

	err = gp.C.Call(ctx, "deleteOwnAuthProvider", _params, nil)
	if err != nil {
		return
	}

	return
}

// GetBranding calls getBranding on the server
func (gp *APIoverJSONRPC) GetBranding(ctx context.Context) (res *Branding, err error) {
	var _params []interface{}

	var result Branding
	err = gp.C.Call(ctx, "getBranding", _params, &result)
	if err != nil {
		return
	}
	res = &result

	return
}

// GetConfiguration calls getConfiguration on the server
func (gp *APIoverJSONRPC) GetConfiguration(ctx context.Context) (res *Configuration, err error) {
	var _params []interface{}

	var result Configuration
	err = gp.C.Call(ctx, "getConfiguration", _params, &result)
	if err != nil {
		return
	}
	res = &result

	return
}

// GetToken calls getToken on the server
func (gp *APIoverJSONRPC) GetToken(ctx context.Context, query *GitpodServerGetTokenSearchOptions) (res *Token, err error) {
	var _params []interface{}

	_params = append(_params, query)

	var result Token
	err = gp.C.Call(ctx, "getToken", _params, &result)
	if err != nil {
		return
	}
	res = &result

	return
}

// GetPortAuthenticationToken calls getPortAuthenticationToken on the server
func (gp *APIoverJSONRPC) GetPortAuthenticationToken(ctx context.Context, workspaceId string) (res *Token, err error) {
	var _params []interface{}

	_params = append(_params, workspaceId)

	var result Token
	err = gp.C.Call(ctx, "getPortAuthenticationToken", _params, &result)
	if err != nil {
		return
	}
	res = &result

	return
}

// DeleteAccount calls deleteAccount on the server
func (gp *APIoverJSONRPC) DeleteAccount(ctx context.Context) (err error) {
	var _params []interface{}

	err = gp.C.Call(ctx, "deleteAccount", _params, nil)
	if err != nil {
		return
	}

	return
}

// GetClientRegion calls getClientRegion on the server
func (gp *APIoverJSONRPC) GetClientRegion(ctx context.Context) (res string, err error) {
	var _params []interface{}

	var result string
	err = gp.C.Call(ctx, "getClientRegion", _params, &result)
	if err != nil {
		return
	}
	res = result

	return
}

// HasPermission calls hasPermission on the server
func (gp *APIoverJSONRPC) HasPermission(ctx context.Context, permission *PermissionName) (res bool, err error) {
	var _params []interface{}

	_params = append(_params, permission)

	var result bool
	err = gp.C.Call(ctx, "hasPermission", _params, &result)
	if err != nil {
		return
	}
	res = &result

	return
}

// GetWorkspaces calls getWorkspaces on the server
func (gp *APIoverJSONRPC) GetWorkspaces(ctx context.Context, options *GitpodServerGetWorkspacesOptions) (res []*WorkspaceInfo, err error) {
	var _params []interface{}

	_params = append(_params, options)

	var result []*WorkspaceInfo
	err = gp.C.Call(ctx, "getWorkspaces", _params, &result)
	if err != nil {
		return
	}
	res = result

	return
}

// GetWorkspaceOwner calls getWorkspaceOwner on the server
func (gp *APIoverJSONRPC) GetWorkspaceOwner(ctx context.Context, workspaceId string) (res *UserInfo, err error) {
	var _params []interface{}

	_params = append(_params, workspaceId)

	var result UserInfo
	err = gp.C.Call(ctx, "getWorkspaceOwner", _params, &result)
	if err != nil {
		return
	}
	res = &result

	return
}

// GetWorkspaceUsers calls getWorkspaceUsers on the server
func (gp *APIoverJSONRPC) GetWorkspaceUsers(ctx context.Context, workspaceId string) (res []*WorkspaceInstanceUser, err error) {
	var _params []interface{}

	_params = append(_params, workspaceId)

	var result []*WorkspaceInstanceUser
	err = gp.C.Call(ctx, "getWorkspaceUsers", _params, &result)
	if err != nil {
		return
	}
	res = result

	return
}

// GetFeaturedRepositories calls getFeaturedRepositories on the server
func (gp *APIoverJSONRPC) GetFeaturedRepositories(ctx context.Context) (res []*WhitelistedRepository, err error) {
	var _params []interface{}

	var result []*WhitelistedRepository
	err = gp.C.Call(ctx, "getFeaturedRepositories", _params, &result)
	if err != nil {
		return
	}
	res = result

	return
}

// GetWorkspace calls getWorkspace on the server
func (gp *APIoverJSONRPC) GetWorkspace(ctx context.Context, id string) (res *WorkspaceInfo, err error) {
	var _params []interface{}

	_params = append(_params, id)

	var result WorkspaceInfo
	err = gp.C.Call(ctx, "getWorkspace", _params, &result)
	if err != nil {
		return
	}
	res = &result

	return
}

// IsWorkspaceOwner calls isWorkspaceOwner on the server
func (gp *APIoverJSONRPC) IsWorkspaceOwner(ctx context.Context, workspaceId string) (res bool, err error) {
	var _params []interface{}

	_params = append(_params, workspaceId)

	var result bool
	err = gp.C.Call(ctx, "isWorkspaceOwner", _params, &result)
	if err != nil {
		return
	}
	res = &result

	return
}

// CreateWorkspace calls createWorkspace on the server
func (gp *APIoverJSONRPC) CreateWorkspace(ctx context.Context, options *GitpodServerCreateWorkspaceOptions) (res *WorkspaceCreationResult, err error) {
	var _params []interface{}

	_params = append(_params, options)

	var result WorkspaceCreationResult
	err = gp.C.Call(ctx, "createWorkspace", _params, &result)
	if err != nil {
		return
	}
	res = &result

	return
}

// StartWorkspace calls startWorkspace on the server
func (gp *APIoverJSONRPC) StartWorkspace(ctx context.Context, id string, options *GitpodServerStartWorkspaceOptions) (res *StartWorkspaceResult, err error) {
	var _params []interface{}

	_params = append(_params, id)
	_params = append(_params, options)

	var result StartWorkspaceResult
	err = gp.C.Call(ctx, "startWorkspace", _params, &result)
	if err != nil {
		return
	}
	res = &result

	return
}

// StopWorkspace calls stopWorkspace on the server
func (gp *APIoverJSONRPC) StopWorkspace(ctx context.Context, id string) (err error) {
	var _params []interface{}

	_params = append(_params, id)

	err = gp.C.Call(ctx, "stopWorkspace", _params, nil)
	if err != nil {
		return
	}

	return
}

// DeleteWorkspace calls deleteWorkspace on the server
func (gp *APIoverJSONRPC) DeleteWorkspace(ctx context.Context, id string) (err error) {
	var _params []interface{}

	_params = append(_params, id)

	err = gp.C.Call(ctx, "deleteWorkspace", _params, nil)
	if err != nil {
		return
	}

	return
}

// SetWorkspaceDescription calls setWorkspaceDescription on the server
func (gp *APIoverJSONRPC) SetWorkspaceDescription(ctx context.Context, id string, desc string) (err error) {
	var _params []interface{}

	_params = append(_params, id)
	_params = append(_params, desc)

	err = gp.C.Call(ctx, "setWorkspaceDescription", _params, nil)
	if err != nil {
		return
	}

	return
}

// ControlAdmission calls controlAdmission on the server
func (gp *APIoverJSONRPC) ControlAdmission(ctx context.Context, id string, level *GitpodServerAdmissionLevel) (err error) {
	var _params []interface{}

	_params = append(_params, id)
	_params = append(_params, level)

	err = gp.C.Call(ctx, "controlAdmission", _params, nil)
	if err != nil {
		return
	}

	return
}

// UpdateWorkspaceUserPin calls updateWorkspaceUserPin on the server
func (gp *APIoverJSONRPC) UpdateWorkspaceUserPin(ctx context.Context, id string, action *GitpodServerPinAction) (err error) {
	var _params []interface{}

	_params = append(_params, id)
	_params = append(_params, action)

	err = gp.C.Call(ctx, "updateWorkspaceUserPin", _params, nil)
	if err != nil {
		return
	}

	return
}

// SendHeartBeat calls sendHeartBeat on the server
func (gp *APIoverJSONRPC) SendHeartBeat(ctx context.Context, options *GitpodServerSendHeartBeatOptions) (err error) {
	var _params []interface{}

	_params = append(_params, options)

	err = gp.C.Call(ctx, "sendHeartBeat", _params, nil)
	if err != nil {
		return
	}

	return
}

// WatchWorkspaceImageBuildLogs calls watchWorkspaceImageBuildLogs on the server
func (gp *APIoverJSONRPC) WatchWorkspaceImageBuildLogs(ctx context.Context, workspaceId string) (err error) {
	var _params []interface{}

	_params = append(_params, workspaceId)

	err = gp.C.Call(ctx, "watchWorkspaceImageBuildLogs", _params, nil)
	if err != nil {
		return
	}

	return
}

// WatchHeadlessWorkspaceLogs calls watchHeadlessWorkspaceLogs on the server
func (gp *APIoverJSONRPC) WatchHeadlessWorkspaceLogs(ctx context.Context, workspaceId string) (err error) {
	var _params []interface{}

	_params = append(_params, workspaceId)

	err = gp.C.Call(ctx, "watchHeadlessWorkspaceLogs", _params, nil)
	if err != nil {
		return
	}

	return
}

// IsPrebuildDone calls isPrebuildDone on the server
func (gp *APIoverJSONRPC) IsPrebuildDone(ctx context.Context, pwsid string) (res bool, err error) {
	var _params []interface{}

	_params = append(_params, pwsid)

	var result bool
	err = gp.C.Call(ctx, "isPrebuildDone", _params, &result)
	if err != nil {
		return
	}
	res = &result

	return
}

// SetWorkspaceTimeout calls setWorkspaceTimeout on the server
func (gp *APIoverJSONRPC) SetWorkspaceTimeout(ctx context.Context, workspaceId string, duration *WorkspaceTimeoutDuration) (res *SetWorkspaceTimeoutResult, err error) {
	var _params []interface{}

	_params = append(_params, workspaceId)
	_params = append(_params, duration)

	var result SetWorkspaceTimeoutResult
	err = gp.C.Call(ctx, "setWorkspaceTimeout", _params, &result)
	if err != nil {
		return
	}
	res = &result

	return
}

// GetWorkspaceTimeout calls getWorkspaceTimeout on the server
func (gp *APIoverJSONRPC) GetWorkspaceTimeout(ctx context.Context, workspaceId string) (res *GetWorkspaceTimeoutResult, err error) {
	var _params []interface{}

	_params = append(_params, workspaceId)

	var result GetWorkspaceTimeoutResult
	err = gp.C.Call(ctx, "getWorkspaceTimeout", _params, &result)
	if err != nil {
		return
	}
	res = &result

	return
}

// SendHeartBeat calls sendHeartBeat on the server
func (gp *APIoverJSONRPC) SendHeartBeat(ctx context.Context, options *GitpodServerSendHeartBeatOptions) (err error) {
	var _params []interface{}

	_params = append(_params, options)

	err = gp.C.Call(ctx, "sendHeartBeat", _params, nil)
	if err != nil {
		return
	}

	return
}

// UpdateWorkspaceUserPin calls updateWorkspaceUserPin on the server
func (gp *APIoverJSONRPC) UpdateWorkspaceUserPin(ctx context.Context, id string, action *GitpodServerPinAction) (err error) {
	var _params []interface{}

	_params = append(_params, id)
	_params = append(_params, action)

	err = gp.C.Call(ctx, "updateWorkspaceUserPin", _params, nil)
	if err != nil {
		return
	}

	return
}

// GetOpenPorts calls getOpenPorts on the server
func (gp *APIoverJSONRPC) GetOpenPorts(ctx context.Context, workspaceId string) (res []*WorkspaceInstancePort, err error) {
	var _params []interface{}

	_params = append(_params, workspaceId)

	var result []*WorkspaceInstancePort
	err = gp.C.Call(ctx, "getOpenPorts", _params, &result)
	if err != nil {
		return
	}
	res = result

	return
}

// OpenPort calls openPort on the server
func (gp *APIoverJSONRPC) OpenPort(ctx context.Context, workspaceId string, port *WorkspaceInstancePort) (res *WorkspaceInstancePort, err error) {
	var _params []interface{}

	_params = append(_params, workspaceId)
	_params = append(_params, port)

	var result WorkspaceInstancePort
	err = gp.C.Call(ctx, "openPort", _params, &result)
	if err != nil {
		return
	}
	res = &result

	return
}

// ClosePort calls closePort on the server
func (gp *APIoverJSONRPC) ClosePort(ctx context.Context, workspaceId string, port float32) (err error) {
	var _params []interface{}

	_params = append(_params, workspaceId)
	_params = append(_params, port)

	err = gp.C.Call(ctx, "closePort", _params, nil)
	if err != nil {
		return
	}

	return
}

// GetUserMessages calls getUserMessages on the server
func (gp *APIoverJSONRPC) GetUserMessages(ctx context.Context, options *GitpodServerGetUserMessagesOptions) (res []*UserMessage, err error) {
	var _params []interface{}

	_params = append(_params, options)

	var result []*UserMessage
	err = gp.C.Call(ctx, "getUserMessages", _params, &result)
	if err != nil {
		return
	}
	res = result

	return
}

// UpdateUserMessages calls updateUserMessages on the server
func (gp *APIoverJSONRPC) UpdateUserMessages(ctx context.Context, options *GitpodServerUpdateUserMessagesOptions) (err error) {
	var _params []interface{}

	_params = append(_params, options)

	err = gp.C.Call(ctx, "updateUserMessages", _params, nil)
	if err != nil {
		return
	}

	return
}

// GetUserStorageResource calls getUserStorageResource on the server
func (gp *APIoverJSONRPC) GetUserStorageResource(ctx context.Context, options *GitpodServerGetUserStorageResourceOptions) (res string, err error) {
	var _params []interface{}

	_params = append(_params, options)

	var result string
	err = gp.C.Call(ctx, "getUserStorageResource", _params, &result)
	if err != nil {
		return
	}
	res = result

	return
}

// UpdateUserStorageResource calls updateUserStorageResource on the server
func (gp *APIoverJSONRPC) UpdateUserStorageResource(ctx context.Context, options *GitpodServerUpdateUserStorageResourceOptions) (err error) {
	var _params []interface{}

	_params = append(_params, options)

	err = gp.C.Call(ctx, "updateUserStorageResource", _params, nil)
	if err != nil {
		return
	}

	return
}

// GetEnvVars calls getEnvVars on the server
func (gp *APIoverJSONRPC) GetEnvVars(ctx context.Context) (res []*UserEnvVarValue, err error) {
	var _params []interface{}

	var result []*UserEnvVarValue
	err = gp.C.Call(ctx, "getEnvVars", _params, &result)
	if err != nil {
		return
	}
	res = result

	return
}

// SetEnvVar calls setEnvVar on the server
func (gp *APIoverJSONRPC) SetEnvVar(ctx context.Context, variable *UserEnvVarValue) (err error) {
	var _params []interface{}

	_params = append(_params, variable)

	err = gp.C.Call(ctx, "setEnvVar", _params, nil)
	if err != nil {
		return
	}

	return
}

// DeleteEnvVar calls deleteEnvVar on the server
func (gp *APIoverJSONRPC) DeleteEnvVar(ctx context.Context, variable *UserEnvVarValue) (err error) {
	var _params []interface{}

	_params = append(_params, variable)

	err = gp.C.Call(ctx, "deleteEnvVar", _params, nil)
	if err != nil {
		return
	}

	return
}

// GetContentBlobUploadUrl calls getContentBlobUploadUrl on the server
func (gp *APIoverJSONRPC) GetContentBlobUploadUrl(ctx context.Context, name string, mediaType string, timeToLive string) (res string, err error) {
	var _params []interface{}

	_params = append(_params, name)
	_params = append(_params, mediaType)
	_params = append(_params, timeToLive)

	var result string
	err = gp.C.Call(ctx, "getContentBlobUploadUrl", _params, &result)
	if err != nil {
		return
	}
	res = result

	return
}

// GetContentBlobDownloadUrl calls getContentBlobDownloadUrl on the server
func (gp *APIoverJSONRPC) GetContentBlobDownloadUrl(ctx context.Context, name string) (res string, err error) {
	var _params []interface{}

	_params = append(_params, name)

	var result string
	err = gp.C.Call(ctx, "getContentBlobDownloadUrl", _params, &result)
	if err != nil {
		return
	}
	res = result

	return
}

// GetGitpodTokens calls getGitpodTokens on the server
func (gp *APIoverJSONRPC) GetGitpodTokens(ctx context.Context) (res []*GitpodToken, err error) {
	var _params []interface{}

	var result []*GitpodToken
	err = gp.C.Call(ctx, "getGitpodTokens", _params, &result)
	if err != nil {
		return
	}
	res = result

	return
}

// GenerateNewGitpodToken calls generateNewGitpodToken on the server
func (gp *APIoverJSONRPC) GenerateNewGitpodToken(ctx context.Context, options *GitpodServerGenerateNewGitpodTokenOptions) (res string, err error) {
	var _params []interface{}

	_params = append(_params, options)

	var result string
	err = gp.C.Call(ctx, "generateNewGitpodToken", _params, &result)
	if err != nil {
		return
	}
	res = result

	return
}

// DeleteGitpodToken calls deleteGitpodToken on the server
func (gp *APIoverJSONRPC) DeleteGitpodToken(ctx context.Context, tokenHash string) (err error) {
	var _params []interface{}

	_params = append(_params, tokenHash)

	err = gp.C.Call(ctx, "deleteGitpodToken", _params, nil)
	if err != nil {
		return
	}

	return
}

// SendFeedback calls sendFeedback on the server
func (gp *APIoverJSONRPC) SendFeedback(ctx context.Context, feedback string) (res string, err error) {
	var _params []interface{}

	_params = append(_params, feedback)

	var result string
	err = gp.C.Call(ctx, "sendFeedback", _params, &result)
	if err != nil {
		return
	}
	res = result

	return
}

// RegisterGithubApp calls registerGithubApp on the server
func (gp *APIoverJSONRPC) RegisterGithubApp(ctx context.Context, installationId string) (err error) {
	var _params []interface{}

	_params = append(_params, installationId)

	err = gp.C.Call(ctx, "registerGithubApp", _params, nil)
	if err != nil {
		return
	}

	return
}

// TakeSnapshot calls takeSnapshot on the server
func (gp *APIoverJSONRPC) TakeSnapshot(ctx context.Context, options *GitpodServerTakeSnapshotOptions) (res string, err error) {
	var _params []interface{}

	_params = append(_params, options)

	var result string
	err = gp.C.Call(ctx, "takeSnapshot", _params, &result)
	if err != nil {
		return
	}
	res = result

	return
}

// GetSnapshots calls getSnapshots on the server
func (gp *APIoverJSONRPC) GetSnapshots(ctx context.Context, workspaceID string) (res []*string, err error) {
	var _params []interface{}

	_params = append(_params, workspaceID)

	var result []*string
	err = gp.C.Call(ctx, "getSnapshots", _params, &result)
	if err != nil {
		return
	}
	res = result

	return
}

// StoreLayout calls storeLayout on the server
func (gp *APIoverJSONRPC) StoreLayout(ctx context.Context, workspaceId string, layoutData string) (err error) {
	var _params []interface{}

	_params = append(_params, workspaceId)
	_params = append(_params, layoutData)

	err = gp.C.Call(ctx, "storeLayout", _params, nil)
	if err != nil {
		return
	}

	return
}

// GetLayout calls getLayout on the server
func (gp *APIoverJSONRPC) GetLayout(ctx context.Context, workspaceId string) (res string, err error) {
	var _params []interface{}

	_params = append(_params, workspaceId)

	var result string
	err = gp.C.Call(ctx, "getLayout", _params, &result)
	if err != nil {
		return
	}
	res = result

	return
}

// PreparePluginUpload calls preparePluginUpload on the server
func (gp *APIoverJSONRPC) PreparePluginUpload(ctx context.Context, params *PreparePluginUploadParams) (res string, err error) {
	var _params []interface{}

	_params = append(_params, params)

	var result string
	err = gp.C.Call(ctx, "preparePluginUpload", _params, &result)
	if err != nil {
		return
	}
	res = result

	return
}

// ResolvePlugins calls resolvePlugins on the server
func (gp *APIoverJSONRPC) ResolvePlugins(ctx context.Context, workspaceId string, params *ResolvePluginsParams) (res *ResolvedPlugins, err error) {
	var _params []interface{}

	_params = append(_params, workspaceId)
	_params = append(_params, params)

	var result ResolvedPlugins
	err = gp.C.Call(ctx, "resolvePlugins", _params, &result)
	if err != nil {
		return
	}
	res = &result

	return
}

// InstallUserPlugins calls installUserPlugins on the server
func (gp *APIoverJSONRPC) InstallUserPlugins(ctx context.Context, params *InstallPluginsParams) (res bool, err error) {
	var _params []interface{}

	_params = append(_params, params)

	var result bool
	err = gp.C.Call(ctx, "installUserPlugins", _params, &result)
	if err != nil {
		return
	}
	res = &result

	return
}

// UninstallUserPlugin calls uninstallUserPlugin on the server
func (gp *APIoverJSONRPC) UninstallUserPlugin(ctx context.Context, params *UninstallPluginParams) (res bool, err error) {
	var _params []interface{}

	_params = append(_params, params)

	var result bool
	err = gp.C.Call(ctx, "uninstallUserPlugin", _params, &result)
	if err != nil {
		return
	}
	res = &result

	return
}
