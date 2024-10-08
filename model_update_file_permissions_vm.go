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

// checks if the UpdateFilePermissionsVM type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateFilePermissionsVM{}

// UpdateFilePermissionsVM struct for UpdateFilePermissionsVM
type UpdateFilePermissionsVM struct {
	CloudBaseVM
	NewPermissions FilePermissionsCRUDVM `json:"newPermissions"`
	Administrate FileAdministrate `json:"administrate"`
	T string `json:"$t"`
}

type _UpdateFilePermissionsVM UpdateFilePermissionsVM

// NewUpdateFilePermissionsVM instantiates a new UpdateFilePermissionsVM object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateFilePermissionsVM(newPermissions FilePermissionsCRUDVM, administrate FileAdministrate, t string) *UpdateFilePermissionsVM {
	this := UpdateFilePermissionsVM{}
	this.T = t
	return &this
}

// NewUpdateFilePermissionsVMWithDefaults instantiates a new UpdateFilePermissionsVM object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateFilePermissionsVMWithDefaults() *UpdateFilePermissionsVM {
	this := UpdateFilePermissionsVM{}
	return &this
}

// GetNewPermissions returns the NewPermissions field value
func (o *UpdateFilePermissionsVM) GetNewPermissions() FilePermissionsCRUDVM {
	if o == nil {
		var ret FilePermissionsCRUDVM
		return ret
	}

	return o.NewPermissions
}

// GetNewPermissionsOk returns a tuple with the NewPermissions field value
// and a boolean to check if the value has been set.
func (o *UpdateFilePermissionsVM) GetNewPermissionsOk() (*FilePermissionsCRUDVM, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NewPermissions, true
}

// SetNewPermissions sets field value
func (o *UpdateFilePermissionsVM) SetNewPermissions(v FilePermissionsCRUDVM) {
	o.NewPermissions = v
}

// GetAdministrate returns the Administrate field value
func (o *UpdateFilePermissionsVM) GetAdministrate() FileAdministrate {
	if o == nil {
		var ret FileAdministrate
		return ret
	}

	return o.Administrate
}

// GetAdministrateOk returns a tuple with the Administrate field value
// and a boolean to check if the value has been set.
func (o *UpdateFilePermissionsVM) GetAdministrateOk() (*FileAdministrate, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Administrate, true
}

// SetAdministrate sets field value
func (o *UpdateFilePermissionsVM) SetAdministrate(v FileAdministrate) {
	o.Administrate = v
}

// GetT returns the T field value
func (o *UpdateFilePermissionsVM) GetT() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.T
}

// GetTOk returns a tuple with the T field value
// and a boolean to check if the value has been set.
func (o *UpdateFilePermissionsVM) GetTOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.T, true
}

// SetT sets field value
func (o *UpdateFilePermissionsVM) SetT(v string) {
	o.T = v
}

func (o UpdateFilePermissionsVM) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateFilePermissionsVM) ToMap() (map[string]interface{}, error) {
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

func (o *UpdateFilePermissionsVM) UnmarshalJSON(data []byte) (err error) {
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

	varUpdateFilePermissionsVM := _UpdateFilePermissionsVM{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUpdateFilePermissionsVM)

	if err != nil {
		return err
	}

	*o = UpdateFilePermissionsVM(varUpdateFilePermissionsVM)

	return err
}

type NullableUpdateFilePermissionsVM struct {
	value *UpdateFilePermissionsVM
	isSet bool
}

func (v NullableUpdateFilePermissionsVM) Get() *UpdateFilePermissionsVM {
	return v.value
}

func (v *NullableUpdateFilePermissionsVM) Set(val *UpdateFilePermissionsVM) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateFilePermissionsVM) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateFilePermissionsVM) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateFilePermissionsVM(val *UpdateFilePermissionsVM) *NullableUpdateFilePermissionsVM {
	return &NullableUpdateFilePermissionsVM{value: val, isSet: true}
}

func (v NullableUpdateFilePermissionsVM) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateFilePermissionsVM) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


