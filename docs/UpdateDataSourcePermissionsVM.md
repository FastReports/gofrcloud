# UpdateDataSourcePermissionsVM

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NewPermissions** | [**DataSourcePermissionsCRUDVM**](DataSourcePermissionsCRUDVM.md) |  | 
**Administrate** | [**DataSourceAdministrate**](DataSourceAdministrate.md) |  | 
**T** | **string** |  | 

## Methods

### NewUpdateDataSourcePermissionsVM

`func NewUpdateDataSourcePermissionsVM(newPermissions DataSourcePermissionsCRUDVM, administrate DataSourceAdministrate, t string, ) *UpdateDataSourcePermissionsVM`

NewUpdateDataSourcePermissionsVM instantiates a new UpdateDataSourcePermissionsVM object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateDataSourcePermissionsVMWithDefaults

`func NewUpdateDataSourcePermissionsVMWithDefaults() *UpdateDataSourcePermissionsVM`

NewUpdateDataSourcePermissionsVMWithDefaults instantiates a new UpdateDataSourcePermissionsVM object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNewPermissions

`func (o *UpdateDataSourcePermissionsVM) GetNewPermissions() DataSourcePermissionsCRUDVM`

GetNewPermissions returns the NewPermissions field if non-nil, zero value otherwise.

### GetNewPermissionsOk

`func (o *UpdateDataSourcePermissionsVM) GetNewPermissionsOk() (*DataSourcePermissionsCRUDVM, bool)`

GetNewPermissionsOk returns a tuple with the NewPermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewPermissions

`func (o *UpdateDataSourcePermissionsVM) SetNewPermissions(v DataSourcePermissionsCRUDVM)`

SetNewPermissions sets NewPermissions field to given value.


### GetAdministrate

`func (o *UpdateDataSourcePermissionsVM) GetAdministrate() DataSourceAdministrate`

GetAdministrate returns the Administrate field if non-nil, zero value otherwise.

### GetAdministrateOk

`func (o *UpdateDataSourcePermissionsVM) GetAdministrateOk() (*DataSourceAdministrate, bool)`

GetAdministrateOk returns a tuple with the Administrate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdministrate

`func (o *UpdateDataSourcePermissionsVM) SetAdministrate(v DataSourceAdministrate)`

SetAdministrate sets Administrate field to given value.


### GetT

`func (o *UpdateDataSourcePermissionsVM) GetT() string`

GetT returns the T field if non-nil, zero value otherwise.

### GetTOk

`func (o *UpdateDataSourcePermissionsVM) GetTOk() (*string, bool)`

GetTOk returns a tuple with the T field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetT

`func (o *UpdateDataSourcePermissionsVM) SetT(v string)`

SetT sets T field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


