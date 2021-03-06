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

// UpdateSubscriptionPermissionsVM struct for UpdateSubscriptionPermissionsVM
type UpdateSubscriptionPermissionsVM struct {
	NewPermissions SubscriptionPermissions `json:"newPermissions"`
	Administrate SubscriptionAdministrate `json:"administrate"`
}

// NewUpdateSubscriptionPermissionsVM instantiates a new UpdateSubscriptionPermissionsVM object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateSubscriptionPermissionsVM(newPermissions SubscriptionPermissions, administrate SubscriptionAdministrate) *UpdateSubscriptionPermissionsVM {
	this := UpdateSubscriptionPermissionsVM{}
	this.NewPermissions = newPermissions
	this.Administrate = administrate
	return &this
}

// NewUpdateSubscriptionPermissionsVMWithDefaults instantiates a new UpdateSubscriptionPermissionsVM object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateSubscriptionPermissionsVMWithDefaults() *UpdateSubscriptionPermissionsVM {
	this := UpdateSubscriptionPermissionsVM{}
	return &this
}

// GetNewPermissions returns the NewPermissions field value
func (o *UpdateSubscriptionPermissionsVM) GetNewPermissions() SubscriptionPermissions {
	if o == nil {
		var ret SubscriptionPermissions
		return ret
	}

	return o.NewPermissions
}

// GetNewPermissionsOk returns a tuple with the NewPermissions field value
// and a boolean to check if the value has been set.
func (o *UpdateSubscriptionPermissionsVM) GetNewPermissionsOk() (*SubscriptionPermissions, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.NewPermissions, true
}

// SetNewPermissions sets field value
func (o *UpdateSubscriptionPermissionsVM) SetNewPermissions(v SubscriptionPermissions) {
	o.NewPermissions = v
}

// GetAdministrate returns the Administrate field value
func (o *UpdateSubscriptionPermissionsVM) GetAdministrate() SubscriptionAdministrate {
	if o == nil {
		var ret SubscriptionAdministrate
		return ret
	}

	return o.Administrate
}

// GetAdministrateOk returns a tuple with the Administrate field value
// and a boolean to check if the value has been set.
func (o *UpdateSubscriptionPermissionsVM) GetAdministrateOk() (*SubscriptionAdministrate, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Administrate, true
}

// SetAdministrate sets field value
func (o *UpdateSubscriptionPermissionsVM) SetAdministrate(v SubscriptionAdministrate) {
	o.Administrate = v
}

func (o UpdateSubscriptionPermissionsVM) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["newPermissions"] = o.NewPermissions
	}
	if true {
		toSerialize["administrate"] = o.Administrate
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateSubscriptionPermissionsVM struct {
	value *UpdateSubscriptionPermissionsVM
	isSet bool
}

func (v NullableUpdateSubscriptionPermissionsVM) Get() *UpdateSubscriptionPermissionsVM {
	return v.value
}

func (v *NullableUpdateSubscriptionPermissionsVM) Set(val *UpdateSubscriptionPermissionsVM) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateSubscriptionPermissionsVM) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateSubscriptionPermissionsVM) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateSubscriptionPermissionsVM(val *UpdateSubscriptionPermissionsVM) *NullableUpdateSubscriptionPermissionsVM {
	return &NullableUpdateSubscriptionPermissionsVM{value: val, isSet: true}
}

func (v NullableUpdateSubscriptionPermissionsVM) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateSubscriptionPermissionsVM) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


