# CreateFetchTaskVM

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** |  | [optional] 
**SubscriptionId** | Pointer to **NullableString** |  | [optional] 
**Type** | Pointer to [**TaskType**](TaskType.md) |  | [optional] 

## Methods

### NewCreateFetchTaskVM

`func NewCreateFetchTaskVM() *CreateFetchTaskVM`

NewCreateFetchTaskVM instantiates a new CreateFetchTaskVM object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateFetchTaskVMWithDefaults

`func NewCreateFetchTaskVMWithDefaults() *CreateFetchTaskVM`

NewCreateFetchTaskVMWithDefaults instantiates a new CreateFetchTaskVM object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateFetchTaskVM) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateFetchTaskVM) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateFetchTaskVM) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateFetchTaskVM) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *CreateFetchTaskVM) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CreateFetchTaskVM) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetSubscriptionId

`func (o *CreateFetchTaskVM) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *CreateFetchTaskVM) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *CreateFetchTaskVM) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *CreateFetchTaskVM) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.

### SetSubscriptionIdNil

`func (o *CreateFetchTaskVM) SetSubscriptionIdNil(b bool)`

 SetSubscriptionIdNil sets the value for SubscriptionId to be an explicit nil

### UnsetSubscriptionId
`func (o *CreateFetchTaskVM) UnsetSubscriptionId()`

UnsetSubscriptionId ensures that no value is present for SubscriptionId, not even an explicit nil
### GetType

`func (o *CreateFetchTaskVM) GetType() TaskType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateFetchTaskVM) GetTypeOk() (*TaskType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateFetchTaskVM) SetType(v TaskType)`

SetType sets Type field to given value.

### HasType

`func (o *CreateFetchTaskVM) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


