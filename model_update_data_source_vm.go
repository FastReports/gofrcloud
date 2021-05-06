/*
 * FastReport Cloud
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package FastReport.Cloud.SDK

import (
	"encoding/json"
)

// UpdateDataSourceVM struct for UpdateDataSourceVM
type UpdateDataSourceVM struct {
	Name *string `json:"name,omitempty"`
	SubscriptionId *string `json:"subscriptionId,omitempty"`
	ConnectionString *string `json:"connectionString,omitempty"`
}

// NewUpdateDataSourceVM instantiates a new UpdateDataSourceVM object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateDataSourceVM() *UpdateDataSourceVM {
	this := UpdateDataSourceVM{}
	return &this
}

// NewUpdateDataSourceVMWithDefaults instantiates a new UpdateDataSourceVM object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateDataSourceVMWithDefaults() *UpdateDataSourceVM {
	this := UpdateDataSourceVM{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UpdateDataSourceVM) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDataSourceVM) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UpdateDataSourceVM) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UpdateDataSourceVM) SetName(v string) {
	o.Name = &v
}

// GetSubscriptionId returns the SubscriptionId field value if set, zero value otherwise.
func (o *UpdateDataSourceVM) GetSubscriptionId() string {
	if o == nil || o.SubscriptionId == nil {
		var ret string
		return ret
	}
	return *o.SubscriptionId
}

// GetSubscriptionIdOk returns a tuple with the SubscriptionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDataSourceVM) GetSubscriptionIdOk() (*string, bool) {
	if o == nil || o.SubscriptionId == nil {
		return nil, false
	}
	return o.SubscriptionId, true
}

// HasSubscriptionId returns a boolean if a field has been set.
func (o *UpdateDataSourceVM) HasSubscriptionId() bool {
	if o != nil && o.SubscriptionId != nil {
		return true
	}

	return false
}

// SetSubscriptionId gets a reference to the given string and assigns it to the SubscriptionId field.
func (o *UpdateDataSourceVM) SetSubscriptionId(v string) {
	o.SubscriptionId = &v
}

// GetConnectionString returns the ConnectionString field value if set, zero value otherwise.
func (o *UpdateDataSourceVM) GetConnectionString() string {
	if o == nil || o.ConnectionString == nil {
		var ret string
		return ret
	}
	return *o.ConnectionString
}

// GetConnectionStringOk returns a tuple with the ConnectionString field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDataSourceVM) GetConnectionStringOk() (*string, bool) {
	if o == nil || o.ConnectionString == nil {
		return nil, false
	}
	return o.ConnectionString, true
}

// HasConnectionString returns a boolean if a field has been set.
func (o *UpdateDataSourceVM) HasConnectionString() bool {
	if o != nil && o.ConnectionString != nil {
		return true
	}

	return false
}

// SetConnectionString gets a reference to the given string and assigns it to the ConnectionString field.
func (o *UpdateDataSourceVM) SetConnectionString(v string) {
	o.ConnectionString = &v
}

func (o UpdateDataSourceVM) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.SubscriptionId != nil {
		toSerialize["subscriptionId"] = o.SubscriptionId
	}
	if o.ConnectionString != nil {
		toSerialize["connectionString"] = o.ConnectionString
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateDataSourceVM struct {
	value *UpdateDataSourceVM
	isSet bool
}

func (v NullableUpdateDataSourceVM) Get() *UpdateDataSourceVM {
	return v.value
}

func (v *NullableUpdateDataSourceVM) Set(val *UpdateDataSourceVM) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateDataSourceVM) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateDataSourceVM) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateDataSourceVM(val *UpdateDataSourceVM) *NullableUpdateDataSourceVM {
	return &NullableUpdateDataSourceVM{value: val, isSet: true}
}

func (v NullableUpdateDataSourceVM) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateDataSourceVM) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

