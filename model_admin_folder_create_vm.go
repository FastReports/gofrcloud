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

// checks if the AdminFolderCreateVM type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AdminFolderCreateVM{}

// AdminFolderCreateVM struct for AdminFolderCreateVM
type AdminFolderCreateVM struct {
	FolderCreateVM
	ParentId string `json:"parentId"`
	OwnerId string `json:"ownerId"`
	Force *bool `json:"force,omitempty"`
	T string `json:"$t"`
}

type _AdminFolderCreateVM AdminFolderCreateVM

// NewAdminFolderCreateVM instantiates a new AdminFolderCreateVM object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdminFolderCreateVM(parentId string, ownerId string, t string) *AdminFolderCreateVM {
	this := AdminFolderCreateVM{}
	this.T = t
	return &this
}

// NewAdminFolderCreateVMWithDefaults instantiates a new AdminFolderCreateVM object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdminFolderCreateVMWithDefaults() *AdminFolderCreateVM {
	this := AdminFolderCreateVM{}
	return &this
}

// GetParentId returns the ParentId field value
func (o *AdminFolderCreateVM) GetParentId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ParentId
}

// GetParentIdOk returns a tuple with the ParentId field value
// and a boolean to check if the value has been set.
func (o *AdminFolderCreateVM) GetParentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ParentId, true
}

// SetParentId sets field value
func (o *AdminFolderCreateVM) SetParentId(v string) {
	o.ParentId = v
}

// GetOwnerId returns the OwnerId field value
func (o *AdminFolderCreateVM) GetOwnerId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OwnerId
}

// GetOwnerIdOk returns a tuple with the OwnerId field value
// and a boolean to check if the value has been set.
func (o *AdminFolderCreateVM) GetOwnerIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OwnerId, true
}

// SetOwnerId sets field value
func (o *AdminFolderCreateVM) SetOwnerId(v string) {
	o.OwnerId = v
}

// GetForce returns the Force field value if set, zero value otherwise.
func (o *AdminFolderCreateVM) GetForce() bool {
	if o == nil || IsNil(o.Force) {
		var ret bool
		return ret
	}
	return *o.Force
}

// GetForceOk returns a tuple with the Force field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdminFolderCreateVM) GetForceOk() (*bool, bool) {
	if o == nil || IsNil(o.Force) {
		return nil, false
	}
	return o.Force, true
}

// HasForce returns a boolean if a field has been set.
func (o *AdminFolderCreateVM) HasForce() bool {
	if o != nil && !IsNil(o.Force) {
		return true
	}

	return false
}

// SetForce gets a reference to the given bool and assigns it to the Force field.
func (o *AdminFolderCreateVM) SetForce(v bool) {
	o.Force = &v
}

// GetT returns the T field value
func (o *AdminFolderCreateVM) GetT() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.T
}

// GetTOk returns a tuple with the T field value
// and a boolean to check if the value has been set.
func (o *AdminFolderCreateVM) GetTOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.T, true
}

// SetT sets field value
func (o *AdminFolderCreateVM) SetT(v string) {
	o.T = v
}

func (o AdminFolderCreateVM) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AdminFolderCreateVM) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedFolderCreateVM, errFolderCreateVM := json.Marshal(o.FolderCreateVM)
	if errFolderCreateVM != nil {
		return map[string]interface{}{}, errFolderCreateVM
	}
	errFolderCreateVM = json.Unmarshal([]byte(serializedFolderCreateVM), &toSerialize)
	if errFolderCreateVM != nil {
		return map[string]interface{}{}, errFolderCreateVM
	}
	toSerialize["parentId"] = o.ParentId
	toSerialize["ownerId"] = o.OwnerId
	if !IsNil(o.Force) {
		toSerialize["force"] = o.Force
	}
	toSerialize["$t"] = o.T
	return toSerialize, nil
}

func (o *AdminFolderCreateVM) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"parentId",
		"ownerId",
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

	varAdminFolderCreateVM := _AdminFolderCreateVM{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAdminFolderCreateVM)

	if err != nil {
		return err
	}

	*o = AdminFolderCreateVM(varAdminFolderCreateVM)

	return err
}

type NullableAdminFolderCreateVM struct {
	value *AdminFolderCreateVM
	isSet bool
}

func (v NullableAdminFolderCreateVM) Get() *AdminFolderCreateVM {
	return v.value
}

func (v *NullableAdminFolderCreateVM) Set(val *AdminFolderCreateVM) {
	v.value = val
	v.isSet = true
}

func (v NullableAdminFolderCreateVM) IsSet() bool {
	return v.isSet
}

func (v *NullableAdminFolderCreateVM) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdminFolderCreateVM(val *AdminFolderCreateVM) *NullableAdminFolderCreateVM {
	return &NullableAdminFolderCreateVM{value: val, isSet: true}
}

func (v NullableAdminFolderCreateVM) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdminFolderCreateVM) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


