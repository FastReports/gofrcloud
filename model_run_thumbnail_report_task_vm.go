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
)

// RunThumbnailReportTaskVM struct for RunThumbnailReportTaskVM
type RunThumbnailReportTaskVM struct {
	SubscriptionId NullableString `json:"subscriptionId,omitempty"`
	Type *TaskType `json:"type,omitempty"`
}

// NewRunThumbnailReportTaskVM instantiates a new RunThumbnailReportTaskVM object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRunThumbnailReportTaskVM() *RunThumbnailReportTaskVM {
	this := RunThumbnailReportTaskVM{}
	return &this
}

// NewRunThumbnailReportTaskVMWithDefaults instantiates a new RunThumbnailReportTaskVM object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRunThumbnailReportTaskVMWithDefaults() *RunThumbnailReportTaskVM {
	this := RunThumbnailReportTaskVM{}
	return &this
}

// GetSubscriptionId returns the SubscriptionId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RunThumbnailReportTaskVM) GetSubscriptionId() string {
	if o == nil || o.SubscriptionId.Get() == nil {
		var ret string
		return ret
	}
	return *o.SubscriptionId.Get()
}

// GetSubscriptionIdOk returns a tuple with the SubscriptionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RunThumbnailReportTaskVM) GetSubscriptionIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SubscriptionId.Get(), o.SubscriptionId.IsSet()
}

// HasSubscriptionId returns a boolean if a field has been set.
func (o *RunThumbnailReportTaskVM) HasSubscriptionId() bool {
	if o != nil && o.SubscriptionId.IsSet() {
		return true
	}

	return false
}

// SetSubscriptionId gets a reference to the given NullableString and assigns it to the SubscriptionId field.
func (o *RunThumbnailReportTaskVM) SetSubscriptionId(v string) {
	o.SubscriptionId.Set(&v)
}
// SetSubscriptionIdNil sets the value for SubscriptionId to be an explicit nil
func (o *RunThumbnailReportTaskVM) SetSubscriptionIdNil() {
	o.SubscriptionId.Set(nil)
}

// UnsetSubscriptionId ensures that no value is present for SubscriptionId, not even an explicit nil
func (o *RunThumbnailReportTaskVM) UnsetSubscriptionId() {
	o.SubscriptionId.Unset()
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *RunThumbnailReportTaskVM) GetType() TaskType {
	if o == nil || o.Type == nil {
		var ret TaskType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RunThumbnailReportTaskVM) GetTypeOk() (*TaskType, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *RunThumbnailReportTaskVM) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given TaskType and assigns it to the Type field.
func (o *RunThumbnailReportTaskVM) SetType(v TaskType) {
	o.Type = &v
}

func (o RunThumbnailReportTaskVM) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.SubscriptionId.IsSet() {
		toSerialize["subscriptionId"] = o.SubscriptionId.Get()
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableRunThumbnailReportTaskVM struct {
	value *RunThumbnailReportTaskVM
	isSet bool
}

func (v NullableRunThumbnailReportTaskVM) Get() *RunThumbnailReportTaskVM {
	return v.value
}

func (v *NullableRunThumbnailReportTaskVM) Set(val *RunThumbnailReportTaskVM) {
	v.value = val
	v.isSet = true
}

func (v NullableRunThumbnailReportTaskVM) IsSet() bool {
	return v.isSet
}

func (v *NullableRunThumbnailReportTaskVM) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRunThumbnailReportTaskVM(val *RunThumbnailReportTaskVM) *NullableRunThumbnailReportTaskVM {
	return &NullableRunThumbnailReportTaskVM{value: val, isSet: true}
}

func (v NullableRunThumbnailReportTaskVM) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRunThumbnailReportTaskVM) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


