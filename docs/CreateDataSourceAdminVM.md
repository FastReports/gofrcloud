# CreateDataSourceAdminVM

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OwnerId** | **string** |  | 
**Name** | Pointer to **NullableString** |  | [optional] 
**ConnectionString** | **string** |  | 
**SubscriptionId** | **string** |  | 
**ConnectionType** | Pointer to [**DataSourceConnectionType**](DataSourceConnectionType.md) |  | [optional] 

## Methods

### NewCreateDataSourceAdminVM

`func NewCreateDataSourceAdminVM(ownerId string, connectionString string, subscriptionId string, ) *CreateDataSourceAdminVM`

NewCreateDataSourceAdminVM instantiates a new CreateDataSourceAdminVM object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateDataSourceAdminVMWithDefaults

`func NewCreateDataSourceAdminVMWithDefaults() *CreateDataSourceAdminVM`

NewCreateDataSourceAdminVMWithDefaults instantiates a new CreateDataSourceAdminVM object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOwnerId

`func (o *CreateDataSourceAdminVM) GetOwnerId() string`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *CreateDataSourceAdminVM) GetOwnerIdOk() (*string, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *CreateDataSourceAdminVM) SetOwnerId(v string)`

SetOwnerId sets OwnerId field to given value.


### GetName

`func (o *CreateDataSourceAdminVM) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateDataSourceAdminVM) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateDataSourceAdminVM) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateDataSourceAdminVM) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *CreateDataSourceAdminVM) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CreateDataSourceAdminVM) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetConnectionString

`func (o *CreateDataSourceAdminVM) GetConnectionString() string`

GetConnectionString returns the ConnectionString field if non-nil, zero value otherwise.

### GetConnectionStringOk

`func (o *CreateDataSourceAdminVM) GetConnectionStringOk() (*string, bool)`

GetConnectionStringOk returns a tuple with the ConnectionString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionString

`func (o *CreateDataSourceAdminVM) SetConnectionString(v string)`

SetConnectionString sets ConnectionString field to given value.


### GetSubscriptionId

`func (o *CreateDataSourceAdminVM) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *CreateDataSourceAdminVM) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *CreateDataSourceAdminVM) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.


### GetConnectionType

`func (o *CreateDataSourceAdminVM) GetConnectionType() DataSourceConnectionType`

GetConnectionType returns the ConnectionType field if non-nil, zero value otherwise.

### GetConnectionTypeOk

`func (o *CreateDataSourceAdminVM) GetConnectionTypeOk() (*DataSourceConnectionType, bool)`

GetConnectionTypeOk returns a tuple with the ConnectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionType

`func (o *CreateDataSourceAdminVM) SetConnectionType(v DataSourceConnectionType)`

SetConnectionType sets ConnectionType field to given value.

### HasConnectionType

`func (o *CreateDataSourceAdminVM) HasConnectionType() bool`

HasConnectionType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


