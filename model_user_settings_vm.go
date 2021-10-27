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
	ProfileVisibility *ProfileVisibility `json:"profileVisibility,omitempty"`
	DefaultSubscription NullableString `json:"defaultSubscription,omitempty"`
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
func (o *UserSettingsVM) GetProfileVisibility() ProfileVisibility {
	if o == nil || o.ProfileVisibility == nil {
		var ret ProfileVisibility
		return ret
	}
	return *o.ProfileVisibility
}

// GetProfileVisibilityOk returns a tuple with the ProfileVisibility field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSettingsVM) GetProfileVisibilityOk() (*ProfileVisibility, bool) {
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

// SetProfileVisibility gets a reference to the given ProfileVisibility and assigns it to the ProfileVisibility field.
func (o *UserSettingsVM) SetProfileVisibility(v ProfileVisibility) {
	o.ProfileVisibility = &v
}

// GetDefaultSubscription returns the DefaultSubscription field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserSettingsVM) GetDefaultSubscription() string {
	if o == nil || o.DefaultSubscription.Get() == nil {
		var ret string
		return ret
	}
	return *o.DefaultSubscription.Get()
}

// GetDefaultSubscriptionOk returns a tuple with the DefaultSubscription field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserSettingsVM) GetDefaultSubscriptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DefaultSubscription.Get(), o.DefaultSubscription.IsSet()
}

// HasDefaultSubscription returns a boolean if a field has been set.
func (o *UserSettingsVM) HasDefaultSubscription() bool {
	if o != nil && o.DefaultSubscription.IsSet() {
		return true
	}

	return false
}

// SetDefaultSubscription gets a reference to the given NullableString and assigns it to the DefaultSubscription field.
func (o *UserSettingsVM) SetDefaultSubscription(v string) {
	o.DefaultSubscription.Set(&v)
}
// SetDefaultSubscriptionNil sets the value for DefaultSubscription to be an explicit nil
func (o *UserSettingsVM) SetDefaultSubscriptionNil() {
	o.DefaultSubscription.Set(nil)
}

// UnsetDefaultSubscription ensures that no value is present for DefaultSubscription, not even an explicit nil
func (o *UserSettingsVM) UnsetDefaultSubscription() {
	o.DefaultSubscription.Unset()
}

func (o UserSettingsVM) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ProfileVisibility != nil {
		toSerialize["profileVisibility"] = o.ProfileVisibility
	}
	if o.DefaultSubscription.IsSet() {
		toSerialize["defaultSubscription"] = o.DefaultSubscription.Get()
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


