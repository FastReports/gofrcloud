/*
FastReport Cloud

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gofrcloud

import (
	"encoding/json"
	"fmt"
)

// SaveMode the model 'SaveMode'
type SaveMode string

// List of SaveMode
const (
	ALL SaveMode = "All"
	ORIGINAL SaveMode = "Original"
	USER SaveMode = "User"
	ROLE SaveMode = "Role"
	SECURITY SaveMode = "Security"
	DENY SaveMode = "Deny"
	CUSTOM SaveMode = "Custom"
)

// All allowed values of SaveMode enum
var AllowedSaveModeEnumValues = []SaveMode{
	"All",
	"Original",
	"User",
	"Role",
	"Security",
	"Deny",
	"Custom",
}

func (v *SaveMode) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SaveMode(value)
	for _, existing := range AllowedSaveModeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SaveMode", value)
}

// NewSaveModeFromValue returns a pointer to a valid SaveMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSaveModeFromValue(v string) (*SaveMode, error) {
	ev := SaveMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SaveMode: valid values are %v", v, AllowedSaveModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SaveMode) IsValid() bool {
	for _, existing := range AllowedSaveModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SaveMode value
func (v SaveMode) Ptr() *SaveMode {
	return &v
}

type NullableSaveMode struct {
	value *SaveMode
	isSet bool
}

func (v NullableSaveMode) Get() *SaveMode {
	return v.value
}

func (v *NullableSaveMode) Set(val *SaveMode) {
	v.value = val
	v.isSet = true
}

func (v NullableSaveMode) IsSet() bool {
	return v.isSet
}

func (v *NullableSaveMode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSaveMode(val *SaveMode) *NullableSaveMode {
	return &NullableSaveMode{value: val, isSet: true}
}

func (v NullableSaveMode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSaveMode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

