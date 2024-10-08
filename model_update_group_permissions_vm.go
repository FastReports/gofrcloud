/*
FastReport Cloud

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gofrcloud

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the UpdateGroupPermissionsVM type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateGroupPermissionsVM{}

// UpdateGroupPermissionsVM struct for UpdateGroupPermissionsVM
type UpdateGroupPermissionsVM struct {
	CloudBaseVM
	NewPermissions GroupPermissionsCRUDVM `json:"newPermissions"`
	Administrate GroupAdministrate `json:"administrate"`
	T string `json:"$t"`
}

type _UpdateGroupPermissionsVM UpdateGroupPermissionsVM

// NewUpdateGroupPermissionsVM instantiates a new UpdateGroupPermissionsVM object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateGroupPermissionsVM(newPermissions GroupPermissionsCRUDVM, administrate GroupAdministrate, t string) *UpdateGroupPermissionsVM {
	this := UpdateGroupPermissionsVM{}
	this.T = t
	return &this
}

// NewUpdateGroupPermissionsVMWithDefaults instantiates a new UpdateGroupPermissionsVM object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateGroupPermissionsVMWithDefaults() *UpdateGroupPermissionsVM {
	this := UpdateGroupPermissionsVM{}
	return &this
}

// GetNewPermissions returns the NewPermissions field value
func (o *UpdateGroupPermissionsVM) GetNewPermissions() GroupPermissionsCRUDVM {
	if o == nil {
		var ret GroupPermissionsCRUDVM
		return ret
	}

	return o.NewPermissions
}

// GetNewPermissionsOk returns a tuple with the NewPermissions field value
// and a boolean to check if the value has been set.
func (o *UpdateGroupPermissionsVM) GetNewPermissionsOk() (*GroupPermissionsCRUDVM, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NewPermissions, true
}

// SetNewPermissions sets field value
func (o *UpdateGroupPermissionsVM) SetNewPermissions(v GroupPermissionsCRUDVM) {
	o.NewPermissions = v
}

// GetAdministrate returns the Administrate field value
func (o *UpdateGroupPermissionsVM) GetAdministrate() GroupAdministrate {
	if o == nil {
		var ret GroupAdministrate
		return ret
	}

	return o.Administrate
}

// GetAdministrateOk returns a tuple with the Administrate field value
// and a boolean to check if the value has been set.
func (o *UpdateGroupPermissionsVM) GetAdministrateOk() (*GroupAdministrate, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Administrate, true
}

// SetAdministrate sets field value
func (o *UpdateGroupPermissionsVM) SetAdministrate(v GroupAdministrate) {
	o.Administrate = v
}

// GetT returns the T field value
func (o *UpdateGroupPermissionsVM) GetT() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.T
}

// GetTOk returns a tuple with the T field value
// and a boolean to check if the value has been set.
func (o *UpdateGroupPermissionsVM) GetTOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.T, true
}

// SetT sets field value
func (o *UpdateGroupPermissionsVM) SetT(v string) {
	o.T = v
}

func (o UpdateGroupPermissionsVM) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateGroupPermissionsVM) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedCloudBaseVM, errCloudBaseVM := json.Marshal(o.CloudBaseVM)
	if errCloudBaseVM != nil {
		return map[string]interface{}{}, errCloudBaseVM
	}
	errCloudBaseVM = json.Unmarshal([]byte(serializedCloudBaseVM), &toSerialize)
	if errCloudBaseVM != nil {
		return map[string]interface{}{}, errCloudBaseVM
	}
	toSerialize["newPermissions"] = o.NewPermissions
	toSerialize["administrate"] = o.Administrate
	toSerialize["$t"] = o.T
	return toSerialize, nil
}

func (o *UpdateGroupPermissionsVM) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"newPermissions",
		"administrate",
		"$t",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varUpdateGroupPermissionsVM := _UpdateGroupPermissionsVM{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUpdateGroupPermissionsVM)

	if err != nil {
		return err
	}

	*o = UpdateGroupPermissionsVM(varUpdateGroupPermissionsVM)

	return err
}

type NullableUpdateGroupPermissionsVM struct {
	value *UpdateGroupPermissionsVM
	isSet bool
}

func (v NullableUpdateGroupPermissionsVM) Get() *UpdateGroupPermissionsVM {
	return v.value
}

func (v *NullableUpdateGroupPermissionsVM) Set(val *UpdateGroupPermissionsVM) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateGroupPermissionsVM) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateGroupPermissionsVM) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateGroupPermissionsVM(val *UpdateGroupPermissionsVM) *NullableUpdateGroupPermissionsVM {
	return &NullableUpdateGroupPermissionsVM{value: val, isSet: true}
}

func (v NullableUpdateGroupPermissionsVM) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateGroupPermissionsVM) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


