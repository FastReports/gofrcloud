# UpdateSubscriptionPermissionsVM

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NewPermissions** | [**SubscriptionPermissionsCRUDVM**](SubscriptionPermissionsCRUDVM.md) |  | 
**Administrate** | [**SubscriptionAdministrate**](SubscriptionAdministrate.md) |  | 
**T** | **string** |  | 

## Methods

### NewUpdateSubscriptionPermissionsVM

`func NewUpdateSubscriptionPermissionsVM(newPermissions SubscriptionPermissionsCRUDVM, administrate SubscriptionAdministrate, t string, ) *UpdateSubscriptionPermissionsVM`

NewUpdateSubscriptionPermissionsVM instantiates a new UpdateSubscriptionPermissionsVM object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateSubscriptionPermissionsVMWithDefaults

`func NewUpdateSubscriptionPermissionsVMWithDefaults() *UpdateSubscriptionPermissionsVM`

NewUpdateSubscriptionPermissionsVMWithDefaults instantiates a new UpdateSubscriptionPermissionsVM object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNewPermissions

`func (o *UpdateSubscriptionPermissionsVM) GetNewPermissions() SubscriptionPermissionsCRUDVM`

GetNewPermissions returns the NewPermissions field if non-nil, zero value otherwise.

### GetNewPermissionsOk

`func (o *UpdateSubscriptionPermissionsVM) GetNewPermissionsOk() (*SubscriptionPermissionsCRUDVM, bool)`

GetNewPermissionsOk returns a tuple with the NewPermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewPermissions

`func (o *UpdateSubscriptionPermissionsVM) SetNewPermissions(v SubscriptionPermissionsCRUDVM)`

SetNewPermissions sets NewPermissions field to given value.


### GetAdministrate

`func (o *UpdateSubscriptionPermissionsVM) GetAdministrate() SubscriptionAdministrate`

GetAdministrate returns the Administrate field if non-nil, zero value otherwise.

### GetAdministrateOk

`func (o *UpdateSubscriptionPermissionsVM) GetAdministrateOk() (*SubscriptionAdministrate, bool)`

GetAdministrateOk returns a tuple with the Administrate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdministrate

`func (o *UpdateSubscriptionPermissionsVM) SetAdministrate(v SubscriptionAdministrate)`

SetAdministrate sets Administrate field to given value.


### GetT

`func (o *UpdateSubscriptionPermissionsVM) GetT() string`

GetT returns the T field if non-nil, zero value otherwise.

### GetTOk

`func (o *UpdateSubscriptionPermissionsVM) GetTOk() (*string, bool)`

GetTOk returns a tuple with the T field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetT

`func (o *UpdateSubscriptionPermissionsVM) SetT(v string)`

SetT sets T field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


