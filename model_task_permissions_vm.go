/*
FastReport Cloud

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gofrcloud

import (
	"encoding/json"
)

// checks if the TaskPermissionsVM type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TaskPermissionsVM{}

// TaskPermissionsVM struct for TaskPermissionsVM
type TaskPermissionsVM struct {
	Permissions *TaskPermissions `json:"permissions,omitempty"`
}

// NewTaskPermissionsVM instantiates a new TaskPermissionsVM object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTaskPermissionsVM() *TaskPermissionsVM {
	this := TaskPermissionsVM{}
	return &this
}

// NewTaskPermissionsVMWithDefaults instantiates a new TaskPermissionsVM object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTaskPermissionsVMWithDefaults() *TaskPermissionsVM {
	this := TaskPermissionsVM{}
	return &this
}

// GetPermissions returns the Permissions field value if set, zero value otherwise.
func (o *TaskPermissionsVM) GetPermissions() TaskPermissions {
	if o == nil || IsNil(o.Permissions) {
		var ret TaskPermissions
		return ret
	}
	return *o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskPermissionsVM) GetPermissionsOk() (*TaskPermissions, bool) {
	if o == nil || IsNil(o.Permissions) {
		return nil, false
	}
	return o.Permissions, true
}

// HasPermissions returns a boolean if a field has been set.
func (o *TaskPermissionsVM) HasPermissions() bool {
	if o != nil && !IsNil(o.Permissions) {
		return true
	}

	return false
}

// SetPermissions gets a reference to the given TaskPermissions and assigns it to the Permissions field.
func (o *TaskPermissionsVM) SetPermissions(v TaskPermissions) {
	o.Permissions = &v
}

func (o TaskPermissionsVM) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TaskPermissionsVM) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Permissions) {
		toSerialize["permissions"] = o.Permissions
	}
	return toSerialize, nil
}

type NullableTaskPermissionsVM struct {
	value *TaskPermissionsVM
	isSet bool
}

func (v NullableTaskPermissionsVM) Get() *TaskPermissionsVM {
	return v.value
}

func (v *NullableTaskPermissionsVM) Set(val *TaskPermissionsVM) {
	v.value = val
	v.isSet = true
}

func (v NullableTaskPermissionsVM) IsSet() bool {
	return v.isSet
}

func (v *NullableTaskPermissionsVM) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTaskPermissionsVM(val *TaskPermissionsVM) *NullableTaskPermissionsVM {
	return &NullableTaskPermissionsVM{value: val, isSet: true}
}

func (v NullableTaskPermissionsVM) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTaskPermissionsVM) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


