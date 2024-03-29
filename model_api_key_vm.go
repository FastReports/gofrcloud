/*
FastReport Cloud

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gofrcloud

import (
	"encoding/json"
	"time"
)

// checks if the ApiKeyVM type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiKeyVM{}

// ApiKeyVM struct for ApiKeyVM
type ApiKeyVM struct {
	Value NullableString `json:"value,omitempty"`
	Description NullableString `json:"description,omitempty"`
	Expired *time.Time `json:"expired,omitempty"`
}

// NewApiKeyVM instantiates a new ApiKeyVM object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiKeyVM() *ApiKeyVM {
	this := ApiKeyVM{}
	return &this
}

// NewApiKeyVMWithDefaults instantiates a new ApiKeyVM object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiKeyVMWithDefaults() *ApiKeyVM {
	this := ApiKeyVM{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiKeyVM) GetValue() string {
	if o == nil || IsNil(o.Value.Get()) {
		var ret string
		return ret
	}
	return *o.Value.Get()
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiKeyVM) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Value.Get(), o.Value.IsSet()
}

// HasValue returns a boolean if a field has been set.
func (o *ApiKeyVM) HasValue() bool {
	if o != nil && o.Value.IsSet() {
		return true
	}

	return false
}

// SetValue gets a reference to the given NullableString and assigns it to the Value field.
func (o *ApiKeyVM) SetValue(v string) {
	o.Value.Set(&v)
}
// SetValueNil sets the value for Value to be an explicit nil
func (o *ApiKeyVM) SetValueNil() {
	o.Value.Set(nil)
}

// UnsetValue ensures that no value is present for Value, not even an explicit nil
func (o *ApiKeyVM) UnsetValue() {
	o.Value.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiKeyVM) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiKeyVM) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *ApiKeyVM) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *ApiKeyVM) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *ApiKeyVM) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *ApiKeyVM) UnsetDescription() {
	o.Description.Unset()
}

// GetExpired returns the Expired field value if set, zero value otherwise.
func (o *ApiKeyVM) GetExpired() time.Time {
	if o == nil || IsNil(o.Expired) {
		var ret time.Time
		return ret
	}
	return *o.Expired
}

// GetExpiredOk returns a tuple with the Expired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiKeyVM) GetExpiredOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Expired) {
		return nil, false
	}
	return o.Expired, true
}

// HasExpired returns a boolean if a field has been set.
func (o *ApiKeyVM) HasExpired() bool {
	if o != nil && !IsNil(o.Expired) {
		return true
	}

	return false
}

// SetExpired gets a reference to the given time.Time and assigns it to the Expired field.
func (o *ApiKeyVM) SetExpired(v time.Time) {
	o.Expired = &v
}

func (o ApiKeyVM) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiKeyVM) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Value.IsSet() {
		toSerialize["value"] = o.Value.Get()
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if !IsNil(o.Expired) {
		toSerialize["expired"] = o.Expired
	}
	return toSerialize, nil
}

type NullableApiKeyVM struct {
	value *ApiKeyVM
	isSet bool
}

func (v NullableApiKeyVM) Get() *ApiKeyVM {
	return v.value
}

func (v *NullableApiKeyVM) Set(val *ApiKeyVM) {
	v.value = val
	v.isSet = true
}

func (v NullableApiKeyVM) IsSet() bool {
	return v.isSet
}

func (v *NullableApiKeyVM) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiKeyVM(val *ApiKeyVM) *NullableApiKeyVM {
	return &NullableApiKeyVM{value: val, isSet: true}
}

func (v NullableApiKeyVM) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiKeyVM) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


