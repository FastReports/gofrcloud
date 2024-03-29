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

// GroupGet the model 'GroupGet'
type GroupGet int32

// List of GroupGet
const (
	_0 GroupGet = 0
	_1 GroupGet = 1
	_2 GroupGet = 2
	_4 GroupGet = 4
	_MINUS_1 GroupGet = -1
)

// All allowed values of GroupGet enum
var AllowedGroupGetEnumValues = []GroupGet{
	0,
	1,
	2,
	4,
	-1,
}

func (v *GroupGet) UnmarshalJSON(src []byte) error {
	var value int32
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := GroupGet(value)
	for _, existing := range AllowedGroupGetEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid GroupGet", value)
}

// NewGroupGetFromValue returns a pointer to a valid GroupGet
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewGroupGetFromValue(v int32) (*GroupGet, error) {
	ev := GroupGet(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for GroupGet: valid values are %v", v, AllowedGroupGetEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v GroupGet) IsValid() bool {
	for _, existing := range AllowedGroupGetEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to GroupGet value
func (v GroupGet) Ptr() *GroupGet {
	return &v
}

type NullableGroupGet struct {
	value *GroupGet
	isSet bool
}

func (v NullableGroupGet) Get() *GroupGet {
	return v.value
}

func (v *NullableGroupGet) Set(val *GroupGet) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupGet) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupGet) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupGet(val *GroupGet) *NullableGroupGet {
	return &NullableGroupGet{value: val, isSet: true}
}

func (v NullableGroupGet) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupGet) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

