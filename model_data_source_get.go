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

// DataSourceGet the model 'DataSourceGet'
type DataSourceGet int32

// List of DataSourceGet
const (
	_0 DataSourceGet = 0
	_1 DataSourceGet = 1
	_2 DataSourceGet = 2
	_4 DataSourceGet = 4
	_8 DataSourceGet = 8
	_16 DataSourceGet = 16
	_MINUS_1 DataSourceGet = -1
)

// All allowed values of DataSourceGet enum
var AllowedDataSourceGetEnumValues = []DataSourceGet{
	0,
	1,
	2,
	4,
	8,
	16,
	-1,
}

func (v *DataSourceGet) UnmarshalJSON(src []byte) error {
	var value int32
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := DataSourceGet(value)
	for _, existing := range AllowedDataSourceGetEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid DataSourceGet", value)
}

// NewDataSourceGetFromValue returns a pointer to a valid DataSourceGet
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDataSourceGetFromValue(v int32) (*DataSourceGet, error) {
	ev := DataSourceGet(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DataSourceGet: valid values are %v", v, AllowedDataSourceGetEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DataSourceGet) IsValid() bool {
	for _, existing := range AllowedDataSourceGetEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DataSourceGet value
func (v DataSourceGet) Ptr() *DataSourceGet {
	return &v
}

type NullableDataSourceGet struct {
	value *DataSourceGet
	isSet bool
}

func (v NullableDataSourceGet) Get() *DataSourceGet {
	return v.value
}

func (v *NullableDataSourceGet) Set(val *DataSourceGet) {
	v.value = val
	v.isSet = true
}

func (v NullableDataSourceGet) IsSet() bool {
	return v.isSet
}

func (v *NullableDataSourceGet) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDataSourceGet(val *DataSourceGet) *NullableDataSourceGet {
	return &NullableDataSourceGet{value: val, isSet: true}
}

func (v NullableDataSourceGet) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDataSourceGet) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

