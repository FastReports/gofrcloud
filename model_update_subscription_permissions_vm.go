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

// checks if the UpdateSubscriptionPermissionsVM type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateSubscriptionPermissionsVM{}

// UpdateSubscriptionPermissionsVM struct for UpdateSubscriptionPermissionsVM
type UpdateSubscriptionPermissionsVM struct {
	CloudBaseVM
	NewPermissions SubscriptionPermissionsCRUDVM `json:"newPermissions"`
	Administrate SubscriptionAdministrate `json:"administrate"`
	T string `json:"$t"`
}

type _UpdateSubscriptionPermissionsVM UpdateSubscriptionPermissionsVM

// NewUpdateSubscriptionPermissionsVM instantiates a new UpdateSubscriptionPermissionsVM object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateSubscriptionPermissionsVM(newPermissions SubscriptionPermissionsCRUDVM, administrate SubscriptionAdministrate, t string) *UpdateSubscriptionPermissionsVM {
	this := UpdateSubscriptionPermissionsVM{}
	this.T = t
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
func (o *UpdateSubscriptionPermissionsVM) GetNewPermissions() SubscriptionPermissionsCRUDVM {
	if o == nil {
		var ret SubscriptionPermissionsCRUDVM
		return ret
	}

	return o.NewPermissions
}

// GetNewPermissionsOk returns a tuple with the NewPermissions field value
// and a boolean to check if the value has been set.
func (o *UpdateSubscriptionPermissionsVM) GetNewPermissionsOk() (*SubscriptionPermissionsCRUDVM, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NewPermissions, true
}

// SetNewPermissions sets field value
func (o *UpdateSubscriptionPermissionsVM) SetNewPermissions(v SubscriptionPermissionsCRUDVM) {
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
	if o == nil {
		return nil, false
	}
	return &o.Administrate, true
}

// SetAdministrate sets field value
func (o *UpdateSubscriptionPermissionsVM) SetAdministrate(v SubscriptionAdministrate) {
	o.Administrate = v
}

// GetT returns the T field value
func (o *UpdateSubscriptionPermissionsVM) GetT() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.T
}

// GetTOk returns a tuple with the T field value
// and a boolean to check if the value has been set.
func (o *UpdateSubscriptionPermissionsVM) GetTOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.T, true
}

// SetT sets field value
func (o *UpdateSubscriptionPermissionsVM) SetT(v string) {
	o.T = v
}

func (o UpdateSubscriptionPermissionsVM) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateSubscriptionPermissionsVM) ToMap() (map[string]interface{}, error) {
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

func (o *UpdateSubscriptionPermissionsVM) UnmarshalJSON(data []byte) (err error) {
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

	varUpdateSubscriptionPermissionsVM := _UpdateSubscriptionPermissionsVM{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUpdateSubscriptionPermissionsVM)

	if err != nil {
		return err
	}

	*o = UpdateSubscriptionPermissionsVM(varUpdateSubscriptionPermissionsVM)

	return err
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


