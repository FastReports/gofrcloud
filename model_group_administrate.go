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

// GroupAdministrate the model 'GroupAdministrate'
type GroupAdministrate int32

// List of GroupAdministrate
const (
	_0 GroupAdministrate = 0
	_1 GroupAdministrate = 1
	_2 GroupAdministrate = 2
	_4 GroupAdministrate = 4
	_8 GroupAdministrate = 8
	_MINUS_1 GroupAdministrate = -1
)

// All allowed values of GroupAdministrate enum
var AllowedGroupAdministrateEnumValues = []GroupAdministrate{
	0,
	1,
	2,
	4,
	8,
	-1,
}

func (v *GroupAdministrate) UnmarshalJSON(src []byte) error {
	var value int32
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := GroupAdministrate(value)
	for _, existing := range AllowedGroupAdministrateEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid GroupAdministrate", value)
}

// NewGroupAdministrateFromValue returns a pointer to a valid GroupAdministrate
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewGroupAdministrateFromValue(v int32) (*GroupAdministrate, error) {
	ev := GroupAdministrate(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for GroupAdministrate: valid values are %v", v, AllowedGroupAdministrateEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v GroupAdministrate) IsValid() bool {
	for _, existing := range AllowedGroupAdministrateEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to GroupAdministrate value
func (v GroupAdministrate) Ptr() *GroupAdministrate {
	return &v
}

type NullableGroupAdministrate struct {
	value *GroupAdministrate
	isSet bool
}

func (v NullableGroupAdministrate) Get() *GroupAdministrate {
	return v.value
}

func (v *NullableGroupAdministrate) Set(val *GroupAdministrate) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupAdministrate) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupAdministrate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupAdministrate(val *GroupAdministrate) *NullableGroupAdministrate {
	return &NullableGroupAdministrate{value: val, isSet: true}
}

func (v NullableGroupAdministrate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupAdministrate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

