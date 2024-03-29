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

// ProfileVisibility the model 'ProfileVisibility'
type ProfileVisibility int32

// List of ProfileVisibility
const (
	_0 ProfileVisibility = 0
	_1 ProfileVisibility = 1
	_2 ProfileVisibility = 2
	_4 ProfileVisibility = 4
)

// All allowed values of ProfileVisibility enum
var AllowedProfileVisibilityEnumValues = []ProfileVisibility{
	0,
	1,
	2,
	4,
}

func (v *ProfileVisibility) UnmarshalJSON(src []byte) error {
	var value int32
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ProfileVisibility(value)
	for _, existing := range AllowedProfileVisibilityEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ProfileVisibility", value)
}

// NewProfileVisibilityFromValue returns a pointer to a valid ProfileVisibility
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProfileVisibilityFromValue(v int32) (*ProfileVisibility, error) {
	ev := ProfileVisibility(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProfileVisibility: valid values are %v", v, AllowedProfileVisibilityEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProfileVisibility) IsValid() bool {
	for _, existing := range AllowedProfileVisibilityEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ProfileVisibility value
func (v ProfileVisibility) Ptr() *ProfileVisibility {
	return &v
}

type NullableProfileVisibility struct {
	value *ProfileVisibility
	isSet bool
}

func (v NullableProfileVisibility) Get() *ProfileVisibility {
	return v.value
}

func (v *NullableProfileVisibility) Set(val *ProfileVisibility) {
	v.value = val
	v.isSet = true
}

func (v NullableProfileVisibility) IsSet() bool {
	return v.isSet
}

func (v *NullableProfileVisibility) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProfileVisibility(val *ProfileVisibility) *NullableProfileVisibility {
	return &NullableProfileVisibility{value: val, isSet: true}
}

func (v NullableProfileVisibility) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProfileVisibility) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

