# RunFetchTaskVM

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SubscriptionId** | Pointer to **NullableString** |  | [optional] 
**Type** | Pointer to [**TaskType**](TaskType.md) |  | [optional] 

## Methods

### NewRunFetchTaskVM

`func NewRunFetchTaskVM() *RunFetchTaskVM`

NewRunFetchTaskVM instantiates a new RunFetchTaskVM object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRunFetchTaskVMWithDefaults

`func NewRunFetchTaskVMWithDefaults() *RunFetchTaskVM`

NewRunFetchTaskVMWithDefaults instantiates a new RunFetchTaskVM object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubscriptionId

`func (o *RunFetchTaskVM) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *RunFetchTaskVM) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *RunFetchTaskVM) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *RunFetchTaskVM) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.

### SetSubscriptionIdNil

`func (o *RunFetchTaskVM) SetSubscriptionIdNil(b bool)`

 SetSubscriptionIdNil sets the value for SubscriptionId to be an explicit nil

### UnsetSubscriptionId
`func (o *RunFetchTaskVM) UnsetSubscriptionId()`

UnsetSubscriptionId ensures that no value is present for SubscriptionId, not even an explicit nil
### GetType

`func (o *RunFetchTaskVM) GetType() TaskType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RunFetchTaskVM) GetTypeOk() (*TaskType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RunFetchTaskVM) SetType(v TaskType)`

SetType sets Type field to given value.

### HasType

`func (o *RunFetchTaskVM) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


