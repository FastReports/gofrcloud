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
	"fmt"
)

// FileCreate the model 'FileCreate'
type FileCreate int32

// List of FileCreate
const (
	_0 FileCreate = 0
	_1 FileCreate = 1
	_MINUS_1 FileCreate = -1
)

var allowedFileCreateEnumValues = []FileCreate{
	0,
	1,
	-1,
}

func (v *FileCreate) UnmarshalJSON(src []byte) error {
	var value int32
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := FileCreate(value)
	for _, existing := range allowedFileCreateEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid FileCreate", value)
}

// NewFileCreateFromValue returns a pointer to a valid FileCreate
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFileCreateFromValue(v int32) (*FileCreate, error) {
	ev := FileCreate(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for FileCreate: valid values are %v", v, allowedFileCreateEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FileCreate) IsValid() bool {
	for _, existing := range allowedFileCreateEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FileCreate value
func (v FileCreate) Ptr() *FileCreate {
	return &v
}

type NullableFileCreate struct {
	value *FileCreate
	isSet bool
}

func (v NullableFileCreate) Get() *FileCreate {
	return v.value
}

func (v *NullableFileCreate) Set(val *FileCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableFileCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableFileCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFileCreate(val *FileCreate) *NullableFileCreate {
	return &NullableFileCreate{value: val, isSet: true}
}

func (v NullableFileCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFileCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
