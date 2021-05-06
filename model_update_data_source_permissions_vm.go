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

// UpdateDataSourcePermissionsVM struct for UpdateDataSourcePermissionsVM
type UpdateDataSourcePermissionsVM struct {
	NewPermissions DataSourcePermissions `json:"newPermissions"`
	Administrate int32 `json:"administrate"`
}

// NewUpdateDataSourcePermissionsVM instantiates a new UpdateDataSourcePermissionsVM object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateDataSourcePermissionsVM(newPermissions DataSourcePermissions, administrate int32) *UpdateDataSourcePermissionsVM {
	this := UpdateDataSourcePermissionsVM{}
	this.NewPermissions = newPermissions
	this.Administrate = administrate
	return &this
}

// NewUpdateDataSourcePermissionsVMWithDefaults instantiates a new UpdateDataSourcePermissionsVM object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateDataSourcePermissionsVMWithDefaults() *UpdateDataSourcePermissionsVM {
	this := UpdateDataSourcePermissionsVM{}
	return &this
}

// GetNewPermissions returns the NewPermissions field value
func (o *UpdateDataSourcePermissionsVM) GetNewPermissions() DataSourcePermissions {
	if o == nil {
		var ret DataSourcePermissions
		return ret
	}

	return o.NewPermissions
}

// GetNewPermissionsOk returns a tuple with the NewPermissions field value
// and a boolean to check if the value has been set.
func (o *UpdateDataSourcePermissionsVM) GetNewPermissionsOk() (*DataSourcePermissions, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.NewPermissions, true
}

// SetNewPermissions sets field value
func (o *UpdateDataSourcePermissionsVM) SetNewPermissions(v DataSourcePermissions) {
	o.NewPermissions = v
}

// GetAdministrate returns the Administrate field value
func (o *UpdateDataSourcePermissionsVM) GetAdministrate() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Administrate
}

// GetAdministrateOk returns a tuple with the Administrate field value
// and a boolean to check if the value has been set.
func (o *UpdateDataSourcePermissionsVM) GetAdministrateOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Administrate, true
}

// SetAdministrate sets field value
func (o *UpdateDataSourcePermissionsVM) SetAdministrate(v int32) {
	o.Administrate = v
}

func (o UpdateDataSourcePermissionsVM) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["newPermissions"] = o.NewPermissions
	}
	if true {
		toSerialize["administrate"] = o.Administrate
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateDataSourcePermissionsVM struct {
	value *UpdateDataSourcePermissionsVM
	isSet bool
}

func (v NullableUpdateDataSourcePermissionsVM) Get() *UpdateDataSourcePermissionsVM {
	return v.value
}

func (v *NullableUpdateDataSourcePermissionsVM) Set(val *UpdateDataSourcePermissionsVM) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateDataSourcePermissionsVM) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateDataSourcePermissionsVM) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateDataSourcePermissionsVM(val *UpdateDataSourcePermissionsVM) *NullableUpdateDataSourcePermissionsVM {
	return &NullableUpdateDataSourcePermissionsVM{value: val, isSet: true}
}

func (v NullableUpdateDataSourcePermissionsVM) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateDataSourcePermissionsVM) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

