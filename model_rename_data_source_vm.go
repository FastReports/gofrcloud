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

// RenameDataSourceVM struct for RenameDataSourceVM
type RenameDataSourceVM struct {
	Name string `json:"name"`
}

// NewRenameDataSourceVM instantiates a new RenameDataSourceVM object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRenameDataSourceVM(name string) *RenameDataSourceVM {
	this := RenameDataSourceVM{}
	this.Name = name
	return &this
}

// NewRenameDataSourceVMWithDefaults instantiates a new RenameDataSourceVM object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRenameDataSourceVMWithDefaults() *RenameDataSourceVM {
	this := RenameDataSourceVM{}
	return &this
}

// GetName returns the Name field value
func (o *RenameDataSourceVM) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *RenameDataSourceVM) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *RenameDataSourceVM) SetName(v string) {
	o.Name = v
}

func (o RenameDataSourceVM) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableRenameDataSourceVM struct {
	value *RenameDataSourceVM
	isSet bool
}

func (v NullableRenameDataSourceVM) Get() *RenameDataSourceVM {
	return v.value
}

func (v *NullableRenameDataSourceVM) Set(val *RenameDataSourceVM) {
	v.value = val
	v.isSet = true
}

func (v NullableRenameDataSourceVM) IsSet() bool {
	return v.isSet
}

func (v *NullableRenameDataSourceVM) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRenameDataSourceVM(val *RenameDataSourceVM) *NullableRenameDataSourceVM {
	return &NullableRenameDataSourceVM{value: val, isSet: true}
}

func (v NullableRenameDataSourceVM) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRenameDataSourceVM) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


