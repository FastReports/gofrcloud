/*
 * FastReport Cloud
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gofrcloud

import (
	"encoding/json"
)

// UserSettingsVM struct for UserSettingsVM
type UserSettingsVM struct {
	ProfileVisibility *int32 `json:"profileVisibility,omitempty"`
	DefaultSubscription *string `json:"defaultSubscription,omitempty"`
}

// NewUserSettingsVM instantiates a new UserSettingsVM object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserSettingsVM() *UserSettingsVM {
	this := UserSettingsVM{}
	return &this
}

// NewUserSettingsVMWithDefaults instantiates a new UserSettingsVM object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserSettingsVMWithDefaults() *UserSettingsVM {
	this := UserSettingsVM{}
	return &this
}

// GetProfileVisibility returns the ProfileVisibility field value if set, zero value otherwise.
func (o *UserSettingsVM) GetProfileVisibility() int32 {
	if o == nil || o.ProfileVisibility == nil {
		var ret int32
		return ret
	}
	return *o.ProfileVisibility
}

// GetProfileVisibilityOk returns a tuple with the ProfileVisibility field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSettingsVM) GetProfileVisibilityOk() (*int32, bool) {
	if o == nil || o.ProfileVisibility == nil {
		return nil, false
	}
	return o.ProfileVisibility, true
}

// HasProfileVisibility returns a boolean if a field has been set.
func (o *UserSettingsVM) HasProfileVisibility() bool {
	if o != nil && o.ProfileVisibility != nil {
		return true
	}

	return false
}

// SetProfileVisibility gets a reference to the given int32 and assigns it to the ProfileVisibility field.
func (o *UserSettingsVM) SetProfileVisibility(v int32) {
	o.ProfileVisibility = &v
}

// GetDefaultSubscription returns the DefaultSubscription field value if set, zero value otherwise.
func (o *UserSettingsVM) GetDefaultSubscription() string {
	if o == nil || o.DefaultSubscription == nil {
		var ret string
		return ret
	}
	return *o.DefaultSubscription
}

// GetDefaultSubscriptionOk returns a tuple with the DefaultSubscription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSettingsVM) GetDefaultSubscriptionOk() (*string, bool) {
	if o == nil || o.DefaultSubscription == nil {
		return nil, false
	}
	return o.DefaultSubscription, true
}

// HasDefaultSubscription returns a boolean if a field has been set.
func (o *UserSettingsVM) HasDefaultSubscription() bool {
	if o != nil && o.DefaultSubscription != nil {
		return true
	}

	return false
}

// SetDefaultSubscription gets a reference to the given string and assigns it to the DefaultSubscription field.
func (o *UserSettingsVM) SetDefaultSubscription(v string) {
	o.DefaultSubscription = &v
}

func (o UserSettingsVM) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ProfileVisibility != nil {
		toSerialize["profileVisibility"] = o.ProfileVisibility
	}
	if o.DefaultSubscription != nil {
		toSerialize["defaultSubscription"] = o.DefaultSubscription
	}
	return json.Marshal(toSerialize)
}

type NullableUserSettingsVM struct {
	value *UserSettingsVM
	isSet bool
}

func (v NullableUserSettingsVM) Get() *UserSettingsVM {
	return v.value
}

func (v *NullableUserSettingsVM) Set(val *UserSettingsVM) {
	v.value = val
	v.isSet = true
}

func (v NullableUserSettingsVM) IsSet() bool {
	return v.isSet
}

func (v *NullableUserSettingsVM) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserSettingsVM(val *UserSettingsVM) *NullableUserSettingsVM {
	return &NullableUserSettingsVM{value: val, isSet: true}
}

func (v NullableUserSettingsVM) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserSettingsVM) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


