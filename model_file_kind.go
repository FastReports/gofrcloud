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

// FileKind the model 'FileKind'
type FileKind string

// List of FileKind
const (
	TEMPLATE FileKind = "Template"
	REPORT FileKind = "Report"
	EXPORT FileKind = "Export"
	DATA_SOURCE FileKind = "DataSource"
)

var allowedFileKindEnumValues = []FileKind{
	"Template",
	"Report",
	"Export",
	"DataSource",
}

func (v *FileKind) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := FileKind(value)
	for _, existing := range allowedFileKindEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid FileKind", value)
}

// NewFileKindFromValue returns a pointer to a valid FileKind
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFileKindFromValue(v string) (*FileKind, error) {
	ev := FileKind(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for FileKind: valid values are %v", v, allowedFileKindEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FileKind) IsValid() bool {
	for _, existing := range allowedFileKindEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FileKind value
func (v FileKind) Ptr() *FileKind {
	return &v
}

type NullableFileKind struct {
	value *FileKind
	isSet bool
}

func (v NullableFileKind) Get() *FileKind {
	return v.value
}

func (v *NullableFileKind) Set(val *FileKind) {
	v.value = val
	v.isSet = true
}

func (v NullableFileKind) IsSet() bool {
	return v.isSet
}

func (v *NullableFileKind) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFileKind(val *FileKind) *NullableFileKind {
	return &NullableFileKind{value: val, isSet: true}
}

func (v NullableFileKind) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFileKind) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

