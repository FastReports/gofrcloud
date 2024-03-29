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

// GroupCreate the model 'GroupCreate'
type GroupCreate int32

// List of GroupCreate
const (
	_0 GroupCreate = 0
	_1 GroupCreate = 1
	_MINUS_1 GroupCreate = -1
)

// All allowed values of GroupCreate enum
var AllowedGroupCreateEnumValues = []GroupCreate{
	0,
	1,
	-1,
}

func (v *GroupCreate) UnmarshalJSON(src []byte) error {
	var value int32
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := GroupCreate(value)
	for _, existing := range AllowedGroupCreateEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid GroupCreate", value)
}

// NewGroupCreateFromValue returns a pointer to a valid GroupCreate
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewGroupCreateFromValue(v int32) (*GroupCreate, error) {
	ev := GroupCreate(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for GroupCreate: valid values are %v", v, AllowedGroupCreateEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v GroupCreate) IsValid() bool {
	for _, existing := range AllowedGroupCreateEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to GroupCreate value
func (v GroupCreate) Ptr() *GroupCreate {
	return &v
}

type NullableGroupCreate struct {
	value *GroupCreate
	isSet bool
}

func (v NullableGroupCreate) Get() *GroupCreate {
	return v.value
}

func (v *NullableGroupCreate) Set(val *GroupCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupCreate(val *GroupCreate) *NullableGroupCreate {
	return &NullableGroupCreate{value: val, isSet: true}
}

func (v NullableGroupCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

