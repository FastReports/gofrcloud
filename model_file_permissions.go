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

// checks if the FilePermissions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FilePermissions{}

// FilePermissions struct for FilePermissions
type FilePermissions struct {
}

// NewFilePermissions instantiates a new FilePermissions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFilePermissions() *FilePermissions {
	this := FilePermissions{}
	return &this
}

// NewFilePermissionsWithDefaults instantiates a new FilePermissions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFilePermissionsWithDefaults() *FilePermissions {
	this := FilePermissions{}
	return &this
}

func (o FilePermissions) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FilePermissions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	return toSerialize, nil
}

type NullableFilePermissions struct {
	value *FilePermissions
	isSet bool
}

func (v NullableFilePermissions) Get() *FilePermissions {
	return v.value
}

func (v *NullableFilePermissions) Set(val *FilePermissions) {
	v.value = val
	v.isSet = true
}

func (v NullableFilePermissions) IsSet() bool {
	return v.isSet
}

func (v *NullableFilePermissions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFilePermissions(val *FilePermissions) *NullableFilePermissions {
	return &NullableFilePermissions{value: val, isSet: true}
}

func (v NullableFilePermissions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFilePermissions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


