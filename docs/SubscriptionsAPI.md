# \SubscriptionsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SubscriptionsGetDefaultPermissions**](SubscriptionsAPI.md#SubscriptionsGetDefaultPermissions) | **Get** /api/manage/v1/Subscriptions/{subscriptionId}/defaultPermissions | Get subscription&#39;s default permissions for new entities
[**SubscriptionsGetMyPermissions**](SubscriptionsAPI.md#SubscriptionsGetMyPermissions) | **Get** /api/manage/v1/Subscriptions/{subId}/mypermissions | Get user&#39;s permissions for a subscription by id
[**SubscriptionsGetPermissions**](SubscriptionsAPI.md#SubscriptionsGetPermissions) | **Get** /api/manage/v1/Subscriptions/{id}/permissions | Get permissions for a subscription by id
[**SubscriptionsGetSubscription**](SubscriptionsAPI.md#SubscriptionsGetSubscription) | **Get** /api/manage/v1/Subscriptions/{id} | Returns the subscription by id
[**SubscriptionsGetSubscriptions**](SubscriptionsAPI.md#SubscriptionsGetSubscriptions) | **Get** /api/manage/v1/Subscriptions | Returns a list of all subscriptions of current user
[**SubscriptionsRenameSubscription**](SubscriptionsAPI.md#SubscriptionsRenameSubscription) | **Put** /api/manage/v1/Subscriptions/{subscriptionId}/rename | Rename subscription
[**SubscriptionsUpdateDefaultPermissions**](SubscriptionsAPI.md#SubscriptionsUpdateDefaultPermissions) | **Put** /api/manage/v1/Subscriptions/{subscriptionId}/defaultPermissions | Change subscription&#39;s default permissions for new entities
[**SubscriptionsUpdateLocale**](SubscriptionsAPI.md#SubscriptionsUpdateLocale) | **Put** /api/manage/v1/Subscriptions/{subscriptionId}/Locale | Update subscription&#39;s default locale
[**SubscriptionsUpdatePermissions**](SubscriptionsAPI.md#SubscriptionsUpdatePermissions) | **Post** /api/manage/v1/Subscriptions/{id}/permissions | Update permissions



## SubscriptionsGetDefaultPermissions

> DefaultPermissionsVM SubscriptionsGetDefaultPermissions(ctx, subscriptionId).Execute()

Get subscription's default permissions for new entities

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
	subscriptionId := "subscriptionId_example" // string | id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubscriptionsAPI.SubscriptionsGetDefaultPermissions(context.Background(), subscriptionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionsAPI.SubscriptionsGetDefaultPermissions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubscriptionsGetDefaultPermissions`: DefaultPermissionsVM
	fmt.Fprintf(os.Stdout, "Response from `SubscriptionsAPI.SubscriptionsGetDefaultPermissions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionsGetDefaultPermissionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DefaultPermissionsVM**](DefaultPermissionsVM.md)

### Authorization

[ApiKey](../README.md#ApiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscriptionsGetMyPermissions

> MyPermissionsVM SubscriptionsGetMyPermissions(ctx, subId).Execute()

Get user's permissions for a subscription by id

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
	subId := "subId_example" // string | subscription id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubscriptionsAPI.SubscriptionsGetMyPermissions(context.Background(), subId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionsAPI.SubscriptionsGetMyPermissions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubscriptionsGetMyPermissions`: MyPermissionsVM
	fmt.Fprintf(os.Stdout, "Response from `SubscriptionsAPI.SubscriptionsGetMyPermissions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subId** | **string** | subscription id | 

### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionsGetMyPermissionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MyPermissionsVM**](MyPermissionsVM.md)

### Authorization

[ApiKey](../README.md#ApiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscriptionsGetPermissions

> SubscriptionPermissionsVM SubscriptionsGetPermissions(ctx, id).Execute()

Get permissions for a subscription by id

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubscriptionsAPI.SubscriptionsGetPermissions(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionsAPI.SubscriptionsGetPermissions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubscriptionsGetPermissions`: SubscriptionPermissionsVM
	fmt.Fprintf(os.Stdout, "Response from `SubscriptionsAPI.SubscriptionsGetPermissions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionsGetPermissionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SubscriptionPermissionsVM**](SubscriptionPermissionsVM.md)

### Authorization

[ApiKey](../README.md#ApiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscriptionsGetSubscription

> SubscriptionVM SubscriptionsGetSubscription(ctx, id).Execute()

Returns the subscription by id

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
	id := "id_example" // string | Identifier of subscription

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubscriptionsAPI.SubscriptionsGetSubscription(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionsAPI.SubscriptionsGetSubscription``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubscriptionsGetSubscription`: SubscriptionVM
	fmt.Fprintf(os.Stdout, "Response from `SubscriptionsAPI.SubscriptionsGetSubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of subscription | 

### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionsGetSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SubscriptionVM**](SubscriptionVM.md)

### Authorization

[ApiKey](../README.md#ApiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscriptionsGetSubscriptions

> SubscriptionsVM SubscriptionsGetSubscriptions(ctx).Skip(skip).Take(take).Execute()

Returns a list of all subscriptions of current user

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
	skip := int32(56) // int32 | Variable for pagination, defautl value is 0 (optional) (default to 0)
	take := int32(56) // int32 | Variable for pagination, default value is 10 (optional) (default to 10)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubscriptionsAPI.SubscriptionsGetSubscriptions(context.Background()).Skip(skip).Take(take).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionsAPI.SubscriptionsGetSubscriptions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubscriptionsGetSubscriptions`: SubscriptionsVM
	fmt.Fprintf(os.Stdout, "Response from `SubscriptionsAPI.SubscriptionsGetSubscriptions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionsGetSubscriptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **skip** | **int32** | Variable for pagination, defautl value is 0 | [default to 0]
 **take** | **int32** | Variable for pagination, default value is 10 | [default to 10]

### Return type

[**SubscriptionsVM**](SubscriptionsVM.md)

### Authorization

[ApiKey](../README.md#ApiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscriptionsRenameSubscription

> SubscriptionVM SubscriptionsRenameSubscription(ctx, subscriptionId).RenameSubscriptionVM(renameSubscriptionVM).Execute()

Rename subscription

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
	subscriptionId := "subscriptionId_example" // string | id
	renameSubscriptionVM := *openapiclient.NewRenameSubscriptionVM("Name_example", "T_example") // RenameSubscriptionVM | rename VM

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubscriptionsAPI.SubscriptionsRenameSubscription(context.Background(), subscriptionId).RenameSubscriptionVM(renameSubscriptionVM).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionsAPI.SubscriptionsRenameSubscription``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubscriptionsRenameSubscription`: SubscriptionVM
	fmt.Fprintf(os.Stdout, "Response from `SubscriptionsAPI.SubscriptionsRenameSubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionsRenameSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **renameSubscriptionVM** | [**RenameSubscriptionVM**](RenameSubscriptionVM.md) | rename VM | 

### Return type

[**SubscriptionVM**](SubscriptionVM.md)

### Authorization

[ApiKey](../README.md#ApiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscriptionsUpdateDefaultPermissions

> DefaultPermissionsVM SubscriptionsUpdateDefaultPermissions(ctx, subscriptionId).UpdateDefaultPermissionsVM(updateDefaultPermissionsVM).Execute()

Change subscription's default permissions for new entities

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
	subscriptionId := "subscriptionId_example" // string | id
	updateDefaultPermissionsVM := *openapiclient.NewUpdateDefaultPermissionsVM("T_example") // UpdateDefaultPermissionsVM | update default permissions VM

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubscriptionsAPI.SubscriptionsUpdateDefaultPermissions(context.Background(), subscriptionId).UpdateDefaultPermissionsVM(updateDefaultPermissionsVM).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionsAPI.SubscriptionsUpdateDefaultPermissions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubscriptionsUpdateDefaultPermissions`: DefaultPermissionsVM
	fmt.Fprintf(os.Stdout, "Response from `SubscriptionsAPI.SubscriptionsUpdateDefaultPermissions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionsUpdateDefaultPermissionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateDefaultPermissionsVM** | [**UpdateDefaultPermissionsVM**](UpdateDefaultPermissionsVM.md) | update default permissions VM | 

### Return type

[**DefaultPermissionsVM**](DefaultPermissionsVM.md)

### Authorization

[ApiKey](../README.md#ApiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscriptionsUpdateLocale

> SubscriptionVM SubscriptionsUpdateLocale(ctx, subscriptionId).UpdateSubscriptionLocaleVM(updateSubscriptionLocaleVM).Execute()

Update subscription's default locale

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
	subscriptionId := "subscriptionId_example" // string | id
	updateSubscriptionLocaleVM := *openapiclient.NewUpdateSubscriptionLocaleVM("Locale_example", "T_example") // UpdateSubscriptionLocaleVM | update VM

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubscriptionsAPI.SubscriptionsUpdateLocale(context.Background(), subscriptionId).UpdateSubscriptionLocaleVM(updateSubscriptionLocaleVM).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionsAPI.SubscriptionsUpdateLocale``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubscriptionsUpdateLocale`: SubscriptionVM
	fmt.Fprintf(os.Stdout, "Response from `SubscriptionsAPI.SubscriptionsUpdateLocale`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionsUpdateLocaleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateSubscriptionLocaleVM** | [**UpdateSubscriptionLocaleVM**](UpdateSubscriptionLocaleVM.md) | update VM | 

### Return type

[**SubscriptionVM**](SubscriptionVM.md)

### Authorization

[ApiKey](../README.md#ApiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscriptionsUpdatePermissions

> SubscriptionsUpdatePermissions(ctx, id).UpdateSubscriptionPermissionsVM(updateSubscriptionPermissionsVM).Execute()

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
	updateSubscriptionPermissionsVM := *openapiclient.NewUpdateSubscriptionPermissionsVM(*openapiclient.NewSubscriptionPermissionsCRUDVM("T_example"), openapiclient.SubscriptionAdministrate(0), "T_example") // UpdateSubscriptionPermissionsVM |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SubscriptionsAPI.SubscriptionsUpdatePermissions(context.Background(), id).UpdateSubscriptionPermissionsVM(updateSubscriptionPermissionsVM).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionsAPI.SubscriptionsUpdatePermissions``: %v\n", err)
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

Other parameters are passed through a pointer to a apiSubscriptionsUpdatePermissionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateSubscriptionPermissionsVM** | [**UpdateSubscriptionPermissionsVM**](UpdateSubscriptionPermissionsVM.md) |  | 

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

