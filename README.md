# Go API client for [FastReport Cloud](https://fastreport.cloud/)

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: v1
- Package version: 2021.2.13
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import sw "./gofrcloud"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```golang
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

```golang
	config := sw.NewConfiguration()
	config.Servers = sw.ServerConfigurations{
		{
			URL:         "server_url",
			Description: "",
		},
	}
```

### Select Server Configuration

For using other server than the one defined on index 0 set context value `sw.ContextServerIndex` of type `int`.

```golang
ctx := context.WithValue(context.Background(), sw.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), sw.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identifield by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```
ctx := context.WithValue(context.Background(), sw.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), sw.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *http://localhost*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*ApiKeysApi* | [**ApiKeysCreateApiKey**](docs/ApiKeysApi.md#apikeyscreateapikey) | **Post** /api/manage/v1/ApiKeys | Create a new apikey, 5 apikeys for user. Hardcoded for ddos.
*ApiKeysApi* | [**ApiKeysDeleteApiKey**](docs/ApiKeysApi.md#apikeysdeleteapikey) | **Delete** /api/manage/v1/ApiKeys | Delete an apikey
*ApiKeysApi* | [**ApiKeysGetApiKeys**](docs/ApiKeysApi.md#apikeysgetapikeys) | **Get** /api/manage/v1/ApiKeys | Returns list with all api keys of current user
*DataSourcesApi* | [**DataSourcesCreateDataSource**](docs/DataSourcesApi.md#datasourcescreatedatasource) | **Post** /api/data/v1/DataSources | Create new data source
*DataSourcesApi* | [**DataSourcesDeleteDataSource**](docs/DataSourcesApi.md#datasourcesdeletedatasource) | **Delete** /api/data/v1/DataSources/{id} | Delete data source by id
*DataSourcesApi* | [**DataSourcesFetchData**](docs/DataSourcesApi.md#datasourcesfetchdata) | **Get** /api/data/v1/DataSources/{id}/fetch | This should connect to a database and set data structure
*DataSourcesApi* | [**DataSourcesGetAvailableDataSources**](docs/DataSourcesApi.md#datasourcesgetavailabledatasources) | **Get** /api/data/v1/DataSources | Returns all of the data sources, that current user have permission for in a subscription  if subscription id is null, returns all data sources, that current user have permission for
*DataSourcesApi* | [**DataSourcesGetDataSource**](docs/DataSourcesApi.md#datasourcesgetdatasource) | **Get** /api/data/v1/DataSources/{id} | Get data source by id
*DataSourcesApi* | [**DataSourcesGetPermissions**](docs/DataSourcesApi.md#datasourcesgetpermissions) | **Get** /api/data/v1/DataSources/{id}/permissions | Get all Data source permissions
*DataSourcesApi* | [**DataSourcesRenameDataSource**](docs/DataSourcesApi.md#datasourcesrenamedatasource) | **Put** /api/data/v1/DataSources/{id}/rename | Rename data source by id
*DataSourcesApi* | [**DataSourcesUpdateConnectionString**](docs/DataSourcesApi.md#datasourcesupdateconnectionstring) | **Put** /api/data/v1/DataSources/{id}/ConnectionString | Update data source&#39;s connection string by id
*DataSourcesApi* | [**DataSourcesUpdatePermissions**](docs/DataSourcesApi.md#datasourcesupdatepermissions) | **Post** /api/data/v1/DataSources/{id}/permissions | Update permissions
*DataSourcesApi* | [**DataSourcesUpdateSubscriptionDataSource**](docs/DataSourcesApi.md#datasourcesupdatesubscriptiondatasource) | **Put** /api/data/v1/DataSources/{id}/updateSubscription | Update data source&#39;s subscription
*DownloadApi* | [**DownloadGetExport**](docs/DownloadApi.md#downloadgetexport) | **Get** /download/e/{id} | Returns a export file with specified id
*DownloadApi* | [**DownloadGetExportThumbnail**](docs/DownloadApi.md#downloadgetexportthumbnail) | **Get** /download/e/{id}/thumbnail | Returns export&#39;s thumbnail
*DownloadApi* | [**DownloadGetExports**](docs/DownloadApi.md#downloadgetexports) | **Get** /download/es/{archiveName} | Returns a zip archive with selected ids
*DownloadApi* | [**DownloadGetReport**](docs/DownloadApi.md#downloadgetreport) | **Get** /download/r/{id} | Returns a prepared file with specified id
*DownloadApi* | [**DownloadGetReportThumbnail**](docs/DownloadApi.md#downloadgetreportthumbnail) | **Get** /download/r/{id}/thumbnail | Returns report&#39;s thumbnail
*DownloadApi* | [**DownloadGetReports**](docs/DownloadApi.md#downloadgetreports) | **Get** /download/rs/{archiveName} | Returns a zip archive with selected files
*DownloadApi* | [**DownloadGetTemplate**](docs/DownloadApi.md#downloadgettemplate) | **Get** /download/t/{id} | Returns a report file with specified id
*DownloadApi* | [**DownloadGetTemplates**](docs/DownloadApi.md#downloadgettemplates) | **Get** /download/ts/{archiveName} | Returns a zip archive with selected files
*ExportsApi* | [**ExportFolderAndFileGetCount**](docs/ExportsApi.md#exportfolderandfilegetcount) | **Get** /api/rp/v1/Exports/Folder/{id}/CountFolderAndFiles | Get count of files and folders what contains in a specified folder
*ExportsApi* | [**ExportFolderAndFileGetFoldersAndFiles**](docs/ExportsApi.md#exportfolderandfilegetfoldersandfiles) | **Get** /api/rp/v1/Exports/Folder/{id}/ListFolderAndFiles | Get all folders and files from specified folder
*ExportsApi* | [**ExportFoldersCopyFolder**](docs/ExportsApi.md#exportfolderscopyfolder) | **Post** /api/rp/v1/Exports/Folder/{id}/Copy/{folderId} | Move folder to a specified folder
*ExportsApi* | [**ExportFoldersDeleteFolder**](docs/ExportsApi.md#exportfoldersdeletefolder) | **Delete** /api/rp/v1/Exports/Folder/{id} | Delete specified folder
*ExportsApi* | [**ExportFoldersGetBreadcrumbs**](docs/ExportsApi.md#exportfoldersgetbreadcrumbs) | **Get** /api/rp/v1/Exports/Folder/{id}/Breadcrumbs | Get specified folder breadcrumbs
*ExportsApi* | [**ExportFoldersGetFolder**](docs/ExportsApi.md#exportfoldersgetfolder) | **Get** /api/rp/v1/Exports/Folder/{id} | Get specified folder
*ExportsApi* | [**ExportFoldersGetFolders**](docs/ExportsApi.md#exportfoldersgetfolders) | **Get** /api/rp/v1/Exports/Folder/{id}/ListFolders | Get all folders from specified folder
*ExportsApi* | [**ExportFoldersGetFoldersCount**](docs/ExportsApi.md#exportfoldersgetfolderscount) | **Get** /api/rp/v1/Exports/Folder/{id}/CountFolders | Get count of folders what contains in a specified folder
*ExportsApi* | [**ExportFoldersGetPermissions**](docs/ExportsApi.md#exportfoldersgetpermissions) | **Get** /api/rp/v1/Exports/Folder/{id}/permissions | Get all folder permissions
*ExportsApi* | [**ExportFoldersGetRootFolder**](docs/ExportsApi.md#exportfoldersgetrootfolder) | **Get** /api/rp/v1/Exports/Root | Get user&#39;s root folder (without parents)
*ExportsApi* | [**ExportFoldersMoveFolder**](docs/ExportsApi.md#exportfoldersmovefolder) | **Post** /api/rp/v1/Exports/Folder/{id}/Move/{folderId} | Move folder to a specified folder
*ExportsApi* | [**ExportFoldersPostFolder**](docs/ExportsApi.md#exportfolderspostfolder) | **Post** /api/rp/v1/Exports/Folder/{id}/Folder | Create folder
*ExportsApi* | [**ExportFoldersRenameFolder**](docs/ExportsApi.md#exportfoldersrenamefolder) | **Put** /api/rp/v1/Exports/Folder/{id}/Rename | Rename a folder
*ExportsApi* | [**ExportFoldersUpdateIcon**](docs/ExportsApi.md#exportfoldersupdateicon) | **Put** /api/rp/v1/Exports/Folder/{id}/Icon | Update a folder&#39;s icon
*ExportsApi* | [**ExportFoldersUpdatePermissions**](docs/ExportsApi.md#exportfoldersupdatepermissions) | **Post** /api/rp/v1/Exports/{id}/permissions | Update permissions
*ExportsApi* | [**ExportFoldersUpdateTags**](docs/ExportsApi.md#exportfoldersupdatetags) | **Put** /api/rp/v1/Exports/Folder/{id}/UpdateTags | Update tags
*ExportsApi* | [**ExportsCopyFile**](docs/ExportsApi.md#exportscopyfile) | **Post** /api/rp/v1/Exports/File/{id}/Copy/{folderId} | Copy file to a specified folder
*ExportsApi* | [**ExportsDeleteFile**](docs/ExportsApi.md#exportsdeletefile) | **Delete** /api/rp/v1/Exports/File/{id} | Delete specified file
*ExportsApi* | [**ExportsGetFile**](docs/ExportsApi.md#exportsgetfile) | **Get** /api/rp/v1/Exports/File/{id} | Get specified file
*ExportsApi* | [**ExportsGetFilesCount**](docs/ExportsApi.md#exportsgetfilescount) | **Get** /api/rp/v1/Exports/Folder/{id}/CountFiles | Get count of files what contains in a specified folder
*ExportsApi* | [**ExportsGetFilesList**](docs/ExportsApi.md#exportsgetfileslist) | **Get** /api/rp/v1/Exports/Folder/{id}/ListFiles | Get all files from specified folder
*ExportsApi* | [**ExportsGetPermissions**](docs/ExportsApi.md#exportsgetpermissions) | **Get** /api/rp/v1/Exports/File/{id}/permissions | Get all file permissions
*ExportsApi* | [**ExportsMoveFile**](docs/ExportsApi.md#exportsmovefile) | **Post** /api/rp/v1/Exports/File/{id}/Move/{folderId} | Move file to a specified folder
*ExportsApi* | [**ExportsRenameFile**](docs/ExportsApi.md#exportsrenamefile) | **Put** /api/rp/v1/Exports/File/{id}/Rename | Rename a file
*ExportsApi* | [**ExportsUpdateIcon**](docs/ExportsApi.md#exportsupdateicon) | **Put** /api/rp/v1/Exports/File/{id}/Icon | Update a files&#39;s icon
*ExportsApi* | [**ExportsUpdatePermissions**](docs/ExportsApi.md#exportsupdatepermissions) | **Post** /api/rp/v1/Exports/File/{id}/permissions | Update permissions
*ExportsApi* | [**ExportsUpdateTags**](docs/ExportsApi.md#exportsupdatetags) | **Put** /api/rp/v1/Exports/File/{id}/UpdateTags | Update tags
*GroupUsersApi* | [**GroupUsersAddUserToGroup**](docs/GroupUsersApi.md#groupusersaddusertogroup) | **Put** /api/manage/v1/Groups/{id}/Users/{userId} | Add user to the group by identifier
*GroupUsersApi* | [**GroupUsersGetUsersInGroup**](docs/GroupUsersApi.md#groupusersgetusersingroup) | **Get** /api/manage/v1/Groups/{id}/Users | Returns users in the group by identifier
*GroupUsersApi* | [**GroupUsersLeaveFromGroup**](docs/GroupUsersApi.md#groupusersleavefromgroup) | **Delete** /api/manage/v1/Groups/{id}/leave | Leave from the group
*GroupUsersApi* | [**GroupUsersRemoveFromGroup**](docs/GroupUsersApi.md#groupusersremovefromgroup) | **Delete** /api/manage/v1/Groups/{id}/Users/{userId} | Remove user from the group by identifier
*GroupsApi* | [**GroupsCreateGroup**](docs/GroupsApi.md#groupscreategroup) | **Post** /api/manage/v1/Groups | Create a new user group
*GroupsApi* | [**GroupsDeleteGroup**](docs/GroupsApi.md#groupsdeletegroup) | **Delete** /api/manage/v1/Groups/{id} | Delete group by identifier
*GroupsApi* | [**GroupsGetGroup**](docs/GroupsApi.md#groupsgetgroup) | **Get** /api/manage/v1/Groups/{id} | Gets group by identifier
*GroupsApi* | [**GroupsGetGroupList**](docs/GroupsApi.md#groupsgetgrouplist) | **Get** /api/manage/v1/Groups | Gets list of user groups
*GroupsApi* | [**GroupsGetPermissions**](docs/GroupsApi.md#groupsgetpermissions) | **Get** /api/manage/v1/Groups/{id}/permissions | Gets group permissions by identifier
*GroupsApi* | [**GroupsRenameGroup**](docs/GroupsApi.md#groupsrenamegroup) | **Put** /api/manage/v1/Groups/{id}/rename | Rename group by identifier
*GroupsApi* | [**GroupsUpdatePermissions**](docs/GroupsApi.md#groupsupdatepermissions) | **Post** /api/manage/v1/Groups/{id}/permissions | Update permissions
*HealthCheckApi* | [**HealthCheckDataGet**](docs/HealthCheckApi.md#healthcheckdataget) | **Get** /api/backend/v1/HealthCheck | healthcheck
*ReportsApi* | [**ReportFolderAndFileGetCount**](docs/ReportsApi.md#reportfolderandfilegetcount) | **Get** /api/rp/v1/Reports/Folder/{id}/CountFolderAndFiles | Get count of files and folders what contains in a specified folder
*ReportsApi* | [**ReportFolderAndFileGetFoldersAndFiles**](docs/ReportsApi.md#reportfolderandfilegetfoldersandfiles) | **Get** /api/rp/v1/Reports/Folder/{id}/ListFolderAndFiles | Get all folders and files from specified folder
*ReportsApi* | [**ReportFoldersCopyFolder**](docs/ReportsApi.md#reportfolderscopyfolder) | **Post** /api/rp/v1/Reports/Folder/{id}/Copy/{folderId} | Move folder to a specified folder
*ReportsApi* | [**ReportFoldersDeleteFolder**](docs/ReportsApi.md#reportfoldersdeletefolder) | **Delete** /api/rp/v1/Reports/Folder/{id} | Delete specified folder
*ReportsApi* | [**ReportFoldersGetBreadcrumbs**](docs/ReportsApi.md#reportfoldersgetbreadcrumbs) | **Get** /api/rp/v1/Reports/Folder/{id}/Breadcrumbs | Get specified folder breadcrumbs
*ReportsApi* | [**ReportFoldersGetFolder**](docs/ReportsApi.md#reportfoldersgetfolder) | **Get** /api/rp/v1/Reports/Folder/{id} | Get specified folder
*ReportsApi* | [**ReportFoldersGetFolders**](docs/ReportsApi.md#reportfoldersgetfolders) | **Get** /api/rp/v1/Reports/Folder/{id}/ListFolders | Get all folders from specified folder
*ReportsApi* | [**ReportFoldersGetFoldersCount**](docs/ReportsApi.md#reportfoldersgetfolderscount) | **Get** /api/rp/v1/Reports/Folder/{id}/CountFolders | Get count of folders what contains in a specified folder
*ReportsApi* | [**ReportFoldersGetPermissions**](docs/ReportsApi.md#reportfoldersgetpermissions) | **Get** /api/rp/v1/Reports/Folder/{id}/permissions | Get all folder permissions
*ReportsApi* | [**ReportFoldersGetRootFolder**](docs/ReportsApi.md#reportfoldersgetrootfolder) | **Get** /api/rp/v1/Reports/Root | Get user&#39;s root folder (without parents)
*ReportsApi* | [**ReportFoldersMoveFolder**](docs/ReportsApi.md#reportfoldersmovefolder) | **Post** /api/rp/v1/Reports/Folder/{id}/Move/{folderId} | Move folder to a specified folder
*ReportsApi* | [**ReportFoldersPostFolder**](docs/ReportsApi.md#reportfolderspostfolder) | **Post** /api/rp/v1/Reports/Folder/{id}/Folder | Create folder
*ReportsApi* | [**ReportFoldersRenameFolder**](docs/ReportsApi.md#reportfoldersrenamefolder) | **Put** /api/rp/v1/Reports/Folder/{id}/Rename | Rename a folder
*ReportsApi* | [**ReportFoldersUpdateIcon**](docs/ReportsApi.md#reportfoldersupdateicon) | **Put** /api/rp/v1/Reports/Folder/{id}/Icon | Update a folder&#39;s icon
*ReportsApi* | [**ReportFoldersUpdatePermissions**](docs/ReportsApi.md#reportfoldersupdatepermissions) | **Post** /api/rp/v1/Reports/{id}/permissions | Update permissions
*ReportsApi* | [**ReportFoldersUpdateTags**](docs/ReportsApi.md#reportfoldersupdatetags) | **Put** /api/rp/v1/Reports/Folder/{id}/UpdateTags | Update tags
*ReportsApi* | [**ReportsCopyFile**](docs/ReportsApi.md#reportscopyfile) | **Post** /api/rp/v1/Reports/File/{id}/Copy/{folderId} | Copy file to a specified folder
*ReportsApi* | [**ReportsDeleteFile**](docs/ReportsApi.md#reportsdeletefile) | **Delete** /api/rp/v1/Reports/File/{id} | Delete specified file
*ReportsApi* | [**ReportsExport**](docs/ReportsApi.md#reportsexport) | **Post** /api/rp/v1/Reports/File/{id}/Export | Export specified report to a specified format
*ReportsApi* | [**ReportsGetFile**](docs/ReportsApi.md#reportsgetfile) | **Get** /api/rp/v1/Reports/File/{id} | Get specified file
*ReportsApi* | [**ReportsGetFilesCount**](docs/ReportsApi.md#reportsgetfilescount) | **Get** /api/rp/v1/Reports/Folder/{id}/CountFiles | Get count of files what contains in a specified folder
*ReportsApi* | [**ReportsGetFilesList**](docs/ReportsApi.md#reportsgetfileslist) | **Get** /api/rp/v1/Reports/Folder/{id}/ListFiles | Get all files from specified folder
*ReportsApi* | [**ReportsGetPermissions**](docs/ReportsApi.md#reportsgetpermissions) | **Get** /api/rp/v1/Reports/File/{id}/permissions | Get all file permissions
*ReportsApi* | [**ReportsMoveFile**](docs/ReportsApi.md#reportsmovefile) | **Post** /api/rp/v1/Reports/File/{id}/Move/{folderId} | Move file to a specified folder
*ReportsApi* | [**ReportsRenameFile**](docs/ReportsApi.md#reportsrenamefile) | **Put** /api/rp/v1/Reports/File/{id}/Rename | Rename a file
*ReportsApi* | [**ReportsUpdateIcon**](docs/ReportsApi.md#reportsupdateicon) | **Put** /api/rp/v1/Reports/File/{id}/Icon | Update a files&#39;s icon
*ReportsApi* | [**ReportsUpdatePermissions**](docs/ReportsApi.md#reportsupdatepermissions) | **Post** /api/rp/v1/Reports/File/{id}/permissions | Update permissions
*ReportsApi* | [**ReportsUpdateTags**](docs/ReportsApi.md#reportsupdatetags) | **Put** /api/rp/v1/Reports/File/{id}/UpdateTags | Update tags
*ReportsApi* | [**ReportsUploadFile**](docs/ReportsApi.md#reportsuploadfile) | **Post** /api/rp/v1/Reports/Folder/{id}/File | Allows to upload reports into specified folder
*SubscriptionGroupsApi* | [**SubscriptionGroupsGetGroupsList**](docs/SubscriptionGroupsApi.md#subscriptiongroupsgetgroupslist) | **Get** /api/manage/v1/Subscriptions/{subscriptionId}/groups | returns groups of the subscription or subscription user
*SubscriptionInvitesApi* | [**SubscriptionInvitesAcceptInvite**](docs/SubscriptionInvitesApi.md#subscriptioninvitesacceptinvite) | **Get** /api/manage/v1/Subscriptions/{subscriptionId}/invite/{accessToken}/accept | Add a user to the subscription using invite,  the added users will be displayed in the list of users of the subscription,  and these users will also have an active subscription.
*SubscriptionInvitesApi* | [**SubscriptionInvitesCreateInvite**](docs/SubscriptionInvitesApi.md#subscriptioninvitescreateinvite) | **Post** /api/manage/v1/Subscriptions/{subscriptionId}/invite | Create invite to subscription
*SubscriptionInvitesApi* | [**SubscriptionInvitesDeleteInvite**](docs/SubscriptionInvitesApi.md#subscriptioninvitesdeleteinvite) | **Delete** /api/manage/v1/Subscriptions/{subscriptionId}/invite/{accesstoken} | Rename subscription
*SubscriptionInvitesApi* | [**SubscriptionInvitesGetInvites**](docs/SubscriptionInvitesApi.md#subscriptioninvitesgetinvites) | **Get** /api/manage/v1/Subscriptions/{subscriptionId}/invites | Get list of invites in a subscription,  the added users will be displayed in the list of users of the subscription,  and these users will also have an active subscription.
*SubscriptionPlansApi* | [**SubscriptionPlansGetSubscriptionPlan**](docs/SubscriptionPlansApi.md#subscriptionplansgetsubscriptionplan) | **Get** /api/manage/v1/SubscriptionPlans/{id} | Returns a subscription plan. Not all subscriptions can be issued for customer.
*SubscriptionPlansApi* | [**SubscriptionPlansGetSubscriptionPlans**](docs/SubscriptionPlansApi.md#subscriptionplansgetsubscriptionplans) | **Get** /api/manage/v1/SubscriptionPlans | Returns a list of active subscription plans that can be issued to the user.
*SubscriptionUsersApi* | [**SubscriptionUsersAddUser**](docs/SubscriptionUsersApi.md#subscriptionusersadduser) | **Put** /api/manage/v1/Subscriptions/{subscriptionId}/users/{userId} | Add a user to the subscription,  the added users will be displayed in the list of users of the subscription,  and these users will also have an active subscription.
*SubscriptionUsersApi* | [**SubscriptionUsersGetUsers**](docs/SubscriptionUsersApi.md#subscriptionusersgetusers) | **Get** /api/manage/v1/Subscriptions/{subscriptionId}/users | Returns all users of subscription
*SubscriptionUsersApi* | [**SubscriptionUsersLeaveSubscripiton**](docs/SubscriptionUsersApi.md#subscriptionusersleavesubscripiton) | **Delete** /api/manage/v1/Subscriptions/{subscriptionId}/leave | Allows user to leave subscription,.
*SubscriptionUsersApi* | [**SubscriptionUsersRemoveUser**](docs/SubscriptionUsersApi.md#subscriptionusersremoveuser) | **Delete** /api/manage/v1/Subscriptions/{subscriptionId}/users/{userId} | Delete a user from the subscription,  the added users will be displayed in the list of users of the subscription,  and these users will also have an active subscription.
*SubscriptionsApi* | [**SubscriptionsGetDefaultPermissions**](docs/SubscriptionsApi.md#subscriptionsgetdefaultpermissions) | **Get** /api/manage/v1/Subscriptions/{subscriptionId}/defaultPermissions | Get subscription&#39;s default permissions for new entities
*SubscriptionsApi* | [**SubscriptionsGetPermissions**](docs/SubscriptionsApi.md#subscriptionsgetpermissions) | **Get** /api/manage/v1/Subscriptions/{id}/permissions | Get permissions for a subscription by id
*SubscriptionsApi* | [**SubscriptionsGetSubscription**](docs/SubscriptionsApi.md#subscriptionsgetsubscription) | **Get** /api/manage/v1/Subscriptions/{id} | Returns the subscription by id
*SubscriptionsApi* | [**SubscriptionsGetSubscriptions**](docs/SubscriptionsApi.md#subscriptionsgetsubscriptions) | **Get** /api/manage/v1/Subscriptions | Returns a list of all subscriptions of current user
*SubscriptionsApi* | [**SubscriptionsRenameSubscription**](docs/SubscriptionsApi.md#subscriptionsrenamesubscription) | **Put** /api/manage/v1/Subscriptions/{subscriptionId}/rename | Rename subscription
*SubscriptionsApi* | [**SubscriptionsUpdateDefaultPermissions**](docs/SubscriptionsApi.md#subscriptionsupdatedefaultpermissions) | **Put** /api/manage/v1/Subscriptions/{subscriptionId}/defaultPermissions | Change subscription&#39;s default permissions for new entities
*SubscriptionsApi* | [**SubscriptionsUpdateLocale**](docs/SubscriptionsApi.md#subscriptionsupdatelocale) | **Put** /api/manage/v1/Subscriptions/{subscriptionId}/Locale | Update subscription&#39;s default locale
*SubscriptionsApi* | [**SubscriptionsUpdatePermissions**](docs/SubscriptionsApi.md#subscriptionsupdatepermissions) | **Post** /api/manage/v1/Subscriptions/{id}/permissions | Update permissions
*TemplatesApi* | [**TemplateFolderAndFileGetCount**](docs/TemplatesApi.md#templatefolderandfilegetcount) | **Get** /api/rp/v1/Templates/Folder/{id}/CountFolderAndFiles | Get count of files and folders what contains in a specified folder
*TemplatesApi* | [**TemplateFolderAndFileGetFoldersAndFiles**](docs/TemplatesApi.md#templatefolderandfilegetfoldersandfiles) | **Get** /api/rp/v1/Templates/Folder/{id}/ListFolderAndFiles | Get all folders and files from specified folder
*TemplatesApi* | [**TemplateFoldersCopyFolder**](docs/TemplatesApi.md#templatefolderscopyfolder) | **Post** /api/rp/v1/Templates/Folder/{id}/Copy/{folderId} | Move folder to a specified folder
*TemplatesApi* | [**TemplateFoldersDeleteFolder**](docs/TemplatesApi.md#templatefoldersdeletefolder) | **Delete** /api/rp/v1/Templates/Folder/{id} | Delete specified folder
*TemplatesApi* | [**TemplateFoldersGetBreadcrumbs**](docs/TemplatesApi.md#templatefoldersgetbreadcrumbs) | **Get** /api/rp/v1/Templates/Folder/{id}/Breadcrumbs | Get specified folder breadcrumbs
*TemplatesApi* | [**TemplateFoldersGetFolder**](docs/TemplatesApi.md#templatefoldersgetfolder) | **Get** /api/rp/v1/Templates/Folder/{id} | Get specified folder
*TemplatesApi* | [**TemplateFoldersGetFolders**](docs/TemplatesApi.md#templatefoldersgetfolders) | **Get** /api/rp/v1/Templates/Folder/{id}/ListFolders | Get all folders from specified folder
*TemplatesApi* | [**TemplateFoldersGetFoldersCount**](docs/TemplatesApi.md#templatefoldersgetfolderscount) | **Get** /api/rp/v1/Templates/Folder/{id}/CountFolders | Get count of folders what contains in a specified folder
*TemplatesApi* | [**TemplateFoldersGetPermissions**](docs/TemplatesApi.md#templatefoldersgetpermissions) | **Get** /api/rp/v1/Templates/Folder/{id}/permissions | Get all folder permissions
*TemplatesApi* | [**TemplateFoldersGetRootFolder**](docs/TemplatesApi.md#templatefoldersgetrootfolder) | **Get** /api/rp/v1/Templates/Root | Get user&#39;s root folder (without parents)
*TemplatesApi* | [**TemplateFoldersMoveFolder**](docs/TemplatesApi.md#templatefoldersmovefolder) | **Post** /api/rp/v1/Templates/Folder/{id}/Move/{folderId} | Move folder to a specified folder
*TemplatesApi* | [**TemplateFoldersPostFolder**](docs/TemplatesApi.md#templatefolderspostfolder) | **Post** /api/rp/v1/Templates/Folder/{id}/Folder | Create folder
*TemplatesApi* | [**TemplateFoldersRenameFolder**](docs/TemplatesApi.md#templatefoldersrenamefolder) | **Put** /api/rp/v1/Templates/Folder/{id}/Rename | Rename a folder
*TemplatesApi* | [**TemplateFoldersUpdateIcon**](docs/TemplatesApi.md#templatefoldersupdateicon) | **Put** /api/rp/v1/Templates/Folder/{id}/Icon | Update a folder&#39;s icon
*TemplatesApi* | [**TemplateFoldersUpdatePermissions**](docs/TemplatesApi.md#templatefoldersupdatepermissions) | **Post** /api/rp/v1/Templates/{id}/permissions | Update permissions
*TemplatesApi* | [**TemplateFoldersUpdateTags**](docs/TemplatesApi.md#templatefoldersupdatetags) | **Put** /api/rp/v1/Templates/Folder/{id}/UpdateTags | Update tags
*TemplatesApi* | [**TemplatesCopyFile**](docs/TemplatesApi.md#templatescopyfile) | **Post** /api/rp/v1/Templates/File/{id}/Copy/{folderId} | Copy file to a specified folder
*TemplatesApi* | [**TemplatesDeleteFile**](docs/TemplatesApi.md#templatesdeletefile) | **Delete** /api/rp/v1/Templates/File/{id} | Delete specified file
*TemplatesApi* | [**TemplatesExport**](docs/TemplatesApi.md#templatesexport) | **Post** /api/rp/v1/Templates/File/{id}/Export | Export specified report template to a specified format
*TemplatesApi* | [**TemplatesGetFile**](docs/TemplatesApi.md#templatesgetfile) | **Get** /api/rp/v1/Templates/File/{id} | Get specified file
*TemplatesApi* | [**TemplatesGetFilesCount**](docs/TemplatesApi.md#templatesgetfilescount) | **Get** /api/rp/v1/Templates/Folder/{id}/CountFiles | Get count of files what contains in a specified folder
*TemplatesApi* | [**TemplatesGetFilesList**](docs/TemplatesApi.md#templatesgetfileslist) | **Get** /api/rp/v1/Templates/Folder/{id}/ListFiles | Get all files from specified folder
*TemplatesApi* | [**TemplatesGetPermissions**](docs/TemplatesApi.md#templatesgetpermissions) | **Get** /api/rp/v1/Templates/File/{id}/permissions | Get all file permissions
*TemplatesApi* | [**TemplatesMoveFile**](docs/TemplatesApi.md#templatesmovefile) | **Post** /api/rp/v1/Templates/File/{id}/Move/{folderId} | Move file to a specified folder
*TemplatesApi* | [**TemplatesPrepare**](docs/TemplatesApi.md#templatesprepare) | **Post** /api/rp/v1/Templates/File/{id}/Prepare | Prepare specified template to report
*TemplatesApi* | [**TemplatesRenameFile**](docs/TemplatesApi.md#templatesrenamefile) | **Put** /api/rp/v1/Templates/File/{id}/Rename | Rename a file
*TemplatesApi* | [**TemplatesUpdateIcon**](docs/TemplatesApi.md#templatesupdateicon) | **Put** /api/rp/v1/Templates/File/{id}/Icon | Update a files&#39;s icon
*TemplatesApi* | [**TemplatesUpdatePermissions**](docs/TemplatesApi.md#templatesupdatepermissions) | **Post** /api/rp/v1/Templates/File/{id}/permissions | Update permissions
*TemplatesApi* | [**TemplatesUpdateTags**](docs/TemplatesApi.md#templatesupdatetags) | **Put** /api/rp/v1/Templates/File/{id}/UpdateTags | Update tags
*TemplatesApi* | [**TemplatesUploadFile**](docs/TemplatesApi.md#templatesuploadfile) | **Post** /api/rp/v1/Templates/Folder/{id}/File | Upload a file to the specified folder  !
*UserProfileApi* | [**UserProfileGetMyProfile**](docs/UserProfileApi.md#userprofilegetmyprofile) | **Get** /api/manage/v1/UserProfile | Return current profile of the current user
*UserProfileApi* | [**UserProfileGetUserProfile**](docs/UserProfileApi.md#userprofilegetuserprofile) | **Get** /api/manage/v1/UserProfile/{userId} | Return user profile by user identifier.  If the user did not provide information about himself or blocked, then the endpoint will return an empty model. (only id)
*UserProfileApi* | [**UserProfileUpdateMyProfile**](docs/UserProfileApi.md#userprofileupdatemyprofile) | **Put** /api/manage/v1/UserProfile | Update profile of the current user
*UserSettingsApi* | [**UserSettingsGetCurrentUserSettings**](docs/UserSettingsApi.md#usersettingsgetcurrentusersettings) | **Get** /api/manage/v1/UserSettings | Return current user settings.
*UserSettingsApi* | [**UserSettingsUpdateMySettings**](docs/UserSettingsApi.md#usersettingsupdatemysettings) | **Put** /api/manage/v1/UserSettings | Update settings of the current user


## Documentation For Models

 - [ApiKeyVM](docs/ApiKeyVM.md)
 - [ApiKeysVM](docs/ApiKeysVM.md)
 - [BreadcrumbsModel](docs/BreadcrumbsModel.md)
 - [BreadcrumbsVM](docs/BreadcrumbsVM.md)
 - [CountVM](docs/CountVM.md)
 - [CreateApiKeyVM](docs/CreateApiKeyVM.md)
 - [CreateDataSourceVM](docs/CreateDataSourceVM.md)
 - [CreateGroupVM](docs/CreateGroupVM.md)
 - [CreateSubscriptionInviteVM](docs/CreateSubscriptionInviteVM.md)
 - [DataSourcePermission](docs/DataSourcePermission.md)
 - [DataSourcePermissions](docs/DataSourcePermissions.md)
 - [DataSourcePermissionsVM](docs/DataSourcePermissionsVM.md)
 - [DataSourceVM](docs/DataSourceVM.md)
 - [DataSourcesVM](docs/DataSourcesVM.md)
 - [DefaultPermissions](docs/DefaultPermissions.md)
 - [DefaultPermissionsVM](docs/DefaultPermissionsVM.md)
 - [DeleteApiKeyVM](docs/DeleteApiKeyVM.md)
 - [ExportFolderCreateVM](docs/ExportFolderCreateVM.md)
 - [ExportReportTaskVM](docs/ExportReportTaskVM.md)
 - [ExportTemplateTaskVM](docs/ExportTemplateTaskVM.md)
 - [ExportVM](docs/ExportVM.md)
 - [ExportsVM](docs/ExportsVM.md)
 - [FileIconVM](docs/FileIconVM.md)
 - [FilePermission](docs/FilePermission.md)
 - [FilePermissions](docs/FilePermissions.md)
 - [FilePermissionsVM](docs/FilePermissionsVM.md)
 - [FileRenameVM](docs/FileRenameVM.md)
 - [FileTagsUpdateVM](docs/FileTagsUpdateVM.md)
 - [FileVM](docs/FileVM.md)
 - [FilesVM](docs/FilesVM.md)
 - [FolderIconVM](docs/FolderIconVM.md)
 - [FolderRenameVM](docs/FolderRenameVM.md)
 - [FolderTagsUpdateVM](docs/FolderTagsUpdateVM.md)
 - [GroupPermission](docs/GroupPermission.md)
 - [GroupPermissions](docs/GroupPermissions.md)
 - [GroupPermissionsVM](docs/GroupPermissionsVM.md)
 - [GroupUserVM](docs/GroupUserVM.md)
 - [GroupUsersVM](docs/GroupUsersVM.md)
 - [GroupVM](docs/GroupVM.md)
 - [GroupsVM](docs/GroupsVM.md)
 - [InvitedUser](docs/InvitedUser.md)
 - [PrepareTemplateTaskVM](docs/PrepareTemplateTaskVM.md)
 - [ProblemDetails](docs/ProblemDetails.md)
 - [RenameDataSourceVM](docs/RenameDataSourceVM.md)
 - [RenameGroupVM](docs/RenameGroupVM.md)
 - [RenameSubscriptionVM](docs/RenameSubscriptionVM.md)
 - [ReportCreateVM](docs/ReportCreateVM.md)
 - [ReportFolderCreateVM](docs/ReportFolderCreateVM.md)
 - [ReportInfo](docs/ReportInfo.md)
 - [ReportVM](docs/ReportVM.md)
 - [ReportsVM](docs/ReportsVM.md)
 - [SubscriptionFolder](docs/SubscriptionFolder.md)
 - [SubscriptionInviteVM](docs/SubscriptionInviteVM.md)
 - [SubscriptionInvitesVM](docs/SubscriptionInvitesVM.md)
 - [SubscriptionPeriodVM](docs/SubscriptionPeriodVM.md)
 - [SubscriptionPermission](docs/SubscriptionPermission.md)
 - [SubscriptionPermissions](docs/SubscriptionPermissions.md)
 - [SubscriptionPermissionsVM](docs/SubscriptionPermissionsVM.md)
 - [SubscriptionPlanVM](docs/SubscriptionPlanVM.md)
 - [SubscriptionPlansVM](docs/SubscriptionPlansVM.md)
 - [SubscriptionUserVM](docs/SubscriptionUserVM.md)
 - [SubscriptionUsersVM](docs/SubscriptionUsersVM.md)
 - [SubscriptionVM](docs/SubscriptionVM.md)
 - [SubscriptionsVM](docs/SubscriptionsVM.md)
 - [TemplateCreateVM](docs/TemplateCreateVM.md)
 - [TemplateFolderCreateVM](docs/TemplateFolderCreateVM.md)
 - [TemplateVM](docs/TemplateVM.md)
 - [TemplatesVM](docs/TemplatesVM.md)
 - [UpdateDataSourceConnectionStringVM](docs/UpdateDataSourceConnectionStringVM.md)
 - [UpdateDataSourcePermissionsVM](docs/UpdateDataSourcePermissionsVM.md)
 - [UpdateDataSourceSubscriptionVM](docs/UpdateDataSourceSubscriptionVM.md)
 - [UpdateDefaultPermissionsVM](docs/UpdateDefaultPermissionsVM.md)
 - [UpdateFilePermissionsVM](docs/UpdateFilePermissionsVM.md)
 - [UpdateGroupPermissionsVM](docs/UpdateGroupPermissionsVM.md)
 - [UpdateSubscriptionLocaleVM](docs/UpdateSubscriptionLocaleVM.md)
 - [UpdateSubscriptionPermissionsVM](docs/UpdateSubscriptionPermissionsVM.md)
 - [UpdateUserProfileVM](docs/UpdateUserProfileVM.md)
 - [UpdateUserSettingsVM](docs/UpdateUserSettingsVM.md)
 - [UserProfileVM](docs/UserProfileVM.md)
 - [UserSettingsVM](docs/UserSettingsVM.md)


## Documentation For Authorization



### ApiKey

- **Type**: HTTP basic authentication

Example

```golang
auth := context.WithValue(context.Background(), sw.ContextBasicAuth, sw.BasicAuth{
    UserName: "username",
    Password: "password",
})
r, err := client.Service.Operation(auth, args)
```


### JWT

- **Type**: API key
- **API key parameter name**: Authorization
- **Location**: HTTP header

Note, each API key must be added to a map of `map[string]APIKey` where the key is: Authorization and passed in as the auth context for each request.


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Examples

[Here you can find an example of how to use this module](https://github.com/FastReports/FastReport-Cloud/tree/main/Golang%20Demos).

## Authors

Fast Reports team https://www.fast-report.com/en/
