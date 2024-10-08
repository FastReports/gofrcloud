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

// DataSourceUpdate the model 'DataSourceUpdate'
type DataSourceUpdate int32

// List of DataSourceUpdate
const (
	_0 DataSourceUpdate = 0
	_1 DataSourceUpdate = 1
	_2 DataSourceUpdate = 2
	_4 DataSourceUpdate = 4
	_8 DataSourceUpdate = 8
	_16 DataSourceUpdate = 16
	_MINUS_1 DataSourceUpdate = -1
)

// All allowed values of DataSourceUpdate enum
var AllowedDataSourceUpdateEnumValues = []DataSourceUpdate{
	0,
	1,
	2,
	4,
	8,
	16,
	-1,
}

func (v *DataSourceUpdate) UnmarshalJSON(src []byte) error {
	var value int32
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := DataSourceUpdate(value)
	for _, existing := range AllowedDataSourceUpdateEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid DataSourceUpdate", value)
}

// NewDataSourceUpdateFromValue returns a pointer to a valid DataSourceUpdate
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDataSourceUpdateFromValue(v int32) (*DataSourceUpdate, error) {
	ev := DataSourceUpdate(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DataSourceUpdate: valid values are %v", v, AllowedDataSourceUpdateEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DataSourceUpdate) IsValid() bool {
	for _, existing := range AllowedDataSourceUpdateEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DataSourceUpdate value
func (v DataSourceUpdate) Ptr() *DataSourceUpdate {
	return &v
}

type NullableDataSourceUpdate struct {
	value *DataSourceUpdate
	isSet bool
}

func (v NullableDataSourceUpdate) Get() *DataSourceUpdate {
	return v.value
}

func (v *NullableDataSourceUpdate) Set(val *DataSourceUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableDataSourceUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableDataSourceUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDataSourceUpdate(val *DataSourceUpdate) *NullableDataSourceUpdate {
	return &NullableDataSourceUpdate{value: val, isSet: true}
}

func (v NullableDataSourceUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDataSourceUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

