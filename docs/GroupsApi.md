# \GroupsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GroupsCreateGroup**](GroupsApi.md#GroupsCreateGroup) | **Post** /api/manage/v1/Groups | Create a new user group
[**GroupsDeleteGroup**](GroupsApi.md#GroupsDeleteGroup) | **Delete** /api/manage/v1/Groups/{id} | Delete group by identifier
[**GroupsGetGroup**](GroupsApi.md#GroupsGetGroup) | **Get** /api/manage/v1/Groups/{id} | Gets group by identifier
[**GroupsGetGroupList**](GroupsApi.md#GroupsGetGroupList) | **Get** /api/manage/v1/Groups | Returns a list of current user&#39;s groups&lt;br /&gt;  This method will return following data about groups : &lt;br /&gt;  Id, Name, Created time (UTC), Edited time (UTC), creator id, &lt;br /&gt;  editor id, subscription id
[**GroupsGetPermissions**](GroupsApi.md#GroupsGetPermissions) | **Get** /api/manage/v1/Groups/{id}/permissions | Gets group permissions by identifier
[**GroupsRenameGroup**](GroupsApi.md#GroupsRenameGroup) | **Put** /api/manage/v1/Groups/{id}/rename | Rename group by identifier
[**GroupsUpdatePermissions**](GroupsApi.md#GroupsUpdatePermissions) | **Post** /api/manage/v1/Groups/{id}/permissions | Update permissions



## GroupsCreateGroup

> GroupVM GroupsCreateGroup(ctx).CreateGroupVM(createGroupVM).Execute()

Create a new user group

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/fastreports/gofrcloud"
)

func main() {
    createGroupVM := *openapiclient.NewCreateGroupVM("Name_example") // CreateGroupVM | Model for creating (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GroupsApi.GroupsCreateGroup(context.Background()).CreateGroupVM(createGroupVM).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsApi.GroupsCreateGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GroupsCreateGroup`: GroupVM
    fmt.Fprintf(os.Stdout, "Response from `GroupsApi.GroupsCreateGroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGroupsCreateGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createGroupVM** | [**CreateGroupVM**](CreateGroupVM.md) | Model for creating | 

### Return type

[**GroupVM**](GroupVM.md)

### Authorization

[ApiKey](../README.md#ApiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GroupsDeleteGroup

> GroupsDeleteGroup(ctx, id).Execute()

Delete group by identifier

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/fastreports/gofrcloud"
)

func main() {
    id := "id_example" // string | Identifier of group

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.GroupsApi.GroupsDeleteGroup(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsApi.GroupsDeleteGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of group | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupsDeleteGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[ApiKey](../README.md#ApiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GroupsGetGroup

> GroupVM GroupsGetGroup(ctx, id).Execute()

Gets group by identifier

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/fastreports/gofrcloud"
)

func main() {
    id := "id_example" // string | Identifier of group

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GroupsApi.GroupsGetGroup(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsApi.GroupsGetGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GroupsGetGroup`: GroupVM
    fmt.Fprintf(os.Stdout, "Response from `GroupsApi.GroupsGetGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of group | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupsGetGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GroupVM**](GroupVM.md)

### Authorization

[ApiKey](../README.md#ApiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GroupsGetGroupList

> GroupsVM GroupsGetGroupList(ctx).Skip(skip).Take(take).Execute()

Returns a list of current user's groups<br />  This method will return following data about groups : <br />  Id, Name, Created time (UTC), Edited time (UTC), creator id, <br />  editor id, subscription id

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/fastreports/gofrcloud"
)

func main() {
    skip := int32(56) // int32 | How many groups need to skip (optional) (default to 0)
    take := int32(56) // int32 | How many groups need to take (optional) (default to 10)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GroupsApi.GroupsGetGroupList(context.Background()).Skip(skip).Take(take).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsApi.GroupsGetGroupList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GroupsGetGroupList`: GroupsVM
    fmt.Fprintf(os.Stdout, "Response from `GroupsApi.GroupsGetGroupList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGroupsGetGroupListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **skip** | **int32** | How many groups need to skip | [default to 0]
 **take** | **int32** | How many groups need to take | [default to 10]

### Return type

[**GroupsVM**](GroupsVM.md)

### Authorization

[ApiKey](../README.md#ApiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GroupsGetPermissions

> GroupPermissionsVM GroupsGetPermissions(ctx, id).Execute()

Gets group permissions by identifier

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/fastreports/gofrcloud"
)

func main() {
    id := "id_example" // string | Identifier of group

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GroupsApi.GroupsGetPermissions(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsApi.GroupsGetPermissions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GroupsGetPermissions`: GroupPermissionsVM
    fmt.Fprintf(os.Stdout, "Response from `GroupsApi.GroupsGetPermissions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of group | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupsGetPermissionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GroupPermissionsVM**](GroupPermissionsVM.md)

### Authorization

[ApiKey](../README.md#ApiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GroupsRenameGroup

> GroupVM GroupsRenameGroup(ctx, id).RenameGroupVM(renameGroupVM).Execute()

Rename group by identifier

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/fastreports/gofrcloud"
)

func main() {
    id := "id_example" // string | Identifier of group
    renameGroupVM := *openapiclient.NewRenameGroupVM("Name_example") // RenameGroupVM | Model for renaming

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GroupsApi.GroupsRenameGroup(context.Background(), id).RenameGroupVM(renameGroupVM).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsApi.GroupsRenameGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GroupsRenameGroup`: GroupVM
    fmt.Fprintf(os.Stdout, "Response from `GroupsApi.GroupsRenameGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of group | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupsRenameGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **renameGroupVM** | [**RenameGroupVM**](RenameGroupVM.md) | Model for renaming | 

### Return type

[**GroupVM**](GroupVM.md)

### Authorization

[ApiKey](../README.md#ApiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GroupsUpdatePermissions

> GroupsUpdatePermissions(ctx, id).UpdateGroupPermissionsVM(updateGroupPermissionsVM).Execute()

Update permissions

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/fastreports/gofrcloud"
)

func main() {
    id := "id_example" // string | 
    updateGroupPermissionsVM := *openapiclient.NewUpdateGroupPermissionsVM(*openapiclient.NewGroupPermissions(), openapiclient.GroupAdministrate(0)) // UpdateGroupPermissionsVM |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.GroupsApi.GroupsUpdatePermissions(context.Background(), id).UpdateGroupPermissionsVM(updateGroupPermissionsVM).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsApi.GroupsUpdatePermissions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupsUpdatePermissionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateGroupPermissionsVM** | [**UpdateGroupPermissionsVM**](UpdateGroupPermissionsVM.md) |  | 

### Return type

 (empty response body)

### Authorization

[ApiKey](../README.md#ApiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

