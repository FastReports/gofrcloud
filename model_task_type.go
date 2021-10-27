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

// TaskType the model 'TaskType'
type TaskType string

// List of TaskType
const (
	EXPORT_TEMPLATE TaskType = "ExportTemplate"
	EXPORT_REPORT TaskType = "ExportReport"
	PREPARE TaskType = "Prepare"
	FETCH TaskType = "Fetch"
	EMAIL TaskType = "Email"
	WEBHOOK TaskType = "Webhook"
)

var allowedTaskTypeEnumValues = []TaskType{
	"ExportTemplate",
	"ExportReport",
	"Prepare",
	"Fetch",
	"Email",
	"Webhook",
}

func (v *TaskType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TaskType(value)
	for _, existing := range allowedTaskTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TaskType", value)
}

// NewTaskTypeFromValue returns a pointer to a valid TaskType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTaskTypeFromValue(v string) (*TaskType, error) {
	ev := TaskType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for TaskType: valid values are %v", v, allowedTaskTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TaskType) IsValid() bool {
	for _, existing := range allowedTaskTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TaskType value
func (v TaskType) Ptr() *TaskType {
	return &v
}

type NullableTaskType struct {
	value *TaskType
	isSet bool
}

func (v NullableTaskType) Get() *TaskType {
	return v.value
}

func (v *NullableTaskType) Set(val *TaskType) {
	v.value = val
	v.isSet = true
}

func (v NullableTaskType) IsSet() bool {
	return v.isSet
}

func (v *NullableTaskType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTaskType(val *TaskType) *NullableTaskType {
	return &NullableTaskType{value: val, isSet: true}
}

func (v NullableTaskType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTaskType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

