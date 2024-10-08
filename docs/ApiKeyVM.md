# ApiKeyVM

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Expired** | Pointer to **time.Time** |  | [optional] 
**T** | **string** |  | 

## Methods

### NewApiKeyVM

`func NewApiKeyVM(t string, ) *ApiKeyVM`

NewApiKeyVM instantiates a new ApiKeyVM object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiKeyVMWithDefaults

`func NewApiKeyVMWithDefaults() *ApiKeyVM`

NewApiKeyVMWithDefaults instantiates a new ApiKeyVM object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *ApiKeyVM) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ApiKeyVM) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ApiKeyVM) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *ApiKeyVM) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *ApiKeyVM) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *ApiKeyVM) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetDescription

`func (o *ApiKeyVM) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ApiKeyVM) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ApiKeyVM) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ApiKeyVM) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ApiKeyVM) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ApiKeyVM) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetExpired

`func (o *ApiKeyVM) GetExpired() time.Time`

GetExpired returns the Expired field if non-nil, zero value otherwise.

### GetExpiredOk

`func (o *ApiKeyVM) GetExpiredOk() (*time.Time, bool)`

GetExpiredOk returns a tuple with the Expired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpired

`func (o *ApiKeyVM) SetExpired(v time.Time)`

SetExpired sets Expired field to given value.

### HasExpired

`func (o *ApiKeyVM) HasExpired() bool`

HasExpired returns a boolean if a field has been set.

### GetT

`func (o *ApiKeyVM) GetT() string`

GetT returns the T field if non-nil, zero value otherwise.

### GetTOk

`func (o *ApiKeyVM) GetTOk() (*string, bool)`

GetTOk returns a tuple with the T field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetT

`func (o *ApiKeyVM) SetT(v string)`

SetT sets T field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


