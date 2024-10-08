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

// FileSorting the model 'FileSorting'
type FileSorting string

// List of FileSorting
const (
	CREATED_TIME FileSorting = "CreatedTime"
	EDITED_TIME FileSorting = "EditedTime"
	SIZE FileSorting = "Size"
	NAME FileSorting = "Name"
)

// All allowed values of FileSorting enum
var AllowedFileSortingEnumValues = []FileSorting{
	"CreatedTime",
	"EditedTime",
	"Size",
	"Name",
}

func (v *FileSorting) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := FileSorting(value)
	for _, existing := range AllowedFileSortingEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid FileSorting", value)
}

// NewFileSortingFromValue returns a pointer to a valid FileSorting
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFileSortingFromValue(v string) (*FileSorting, error) {
	ev := FileSorting(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for FileSorting: valid values are %v", v, AllowedFileSortingEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FileSorting) IsValid() bool {
	for _, existing := range AllowedFileSortingEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FileSorting value
func (v FileSorting) Ptr() *FileSorting {
	return &v
}

type NullableFileSorting struct {
	value *FileSorting
	isSet bool
}

func (v NullableFileSorting) Get() *FileSorting {
	return v.value
}

func (v *NullableFileSorting) Set(val *FileSorting) {
	v.value = val
	v.isSet = true
}

func (v NullableFileSorting) IsSet() bool {
	return v.isSet
}

func (v *NullableFileSorting) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFileSorting(val *FileSorting) *NullableFileSorting {
	return &NullableFileSorting{value: val, isSet: true}
}

func (v NullableFileSorting) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFileSorting) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

