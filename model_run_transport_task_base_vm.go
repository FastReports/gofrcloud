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

// RunTransportTaskBaseVM struct for RunTransportTaskBaseVM
type RunTransportTaskBaseVM struct {
	Files []RunInputFileVM `json:"files,omitempty"`
	SubscriptionId NullableString `json:"subscriptionId,omitempty"`
	Type *TaskType `json:"type,omitempty"`
}

// NewRunTransportTaskBaseVM instantiates a new RunTransportTaskBaseVM object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRunTransportTaskBaseVM() *RunTransportTaskBaseVM {
	this := RunTransportTaskBaseVM{}
	return &this
}

// NewRunTransportTaskBaseVMWithDefaults instantiates a new RunTransportTaskBaseVM object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRunTransportTaskBaseVMWithDefaults() *RunTransportTaskBaseVM {
	this := RunTransportTaskBaseVM{}
	return &this
}

// GetFiles returns the Files field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RunTransportTaskBaseVM) GetFiles() []RunInputFileVM {
	if o == nil  {
		var ret []RunInputFileVM
		return ret
	}
	return o.Files
}

// GetFilesOk returns a tuple with the Files field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RunTransportTaskBaseVM) GetFilesOk() (*[]RunInputFileVM, bool) {
	if o == nil || o.Files == nil {
		return nil, false
	}
	return &o.Files, true
}

// HasFiles returns a boolean if a field has been set.
func (o *RunTransportTaskBaseVM) HasFiles() bool {
	if o != nil && o.Files != nil {
		return true
	}

	return false
}

// SetFiles gets a reference to the given []RunInputFileVM and assigns it to the Files field.
func (o *RunTransportTaskBaseVM) SetFiles(v []RunInputFileVM) {
	o.Files = v
}

// GetSubscriptionId returns the SubscriptionId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RunTransportTaskBaseVM) GetSubscriptionId() string {
	if o == nil || o.SubscriptionId.Get() == nil {
		var ret string
		return ret
	}
	return *o.SubscriptionId.Get()
}

// GetSubscriptionIdOk returns a tuple with the SubscriptionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RunTransportTaskBaseVM) GetSubscriptionIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SubscriptionId.Get(), o.SubscriptionId.IsSet()
}

// HasSubscriptionId returns a boolean if a field has been set.
func (o *RunTransportTaskBaseVM) HasSubscriptionId() bool {
	if o != nil && o.SubscriptionId.IsSet() {
		return true
	}

	return false
}

// SetSubscriptionId gets a reference to the given NullableString and assigns it to the SubscriptionId field.
func (o *RunTransportTaskBaseVM) SetSubscriptionId(v string) {
	o.SubscriptionId.Set(&v)
}
// SetSubscriptionIdNil sets the value for SubscriptionId to be an explicit nil
func (o *RunTransportTaskBaseVM) SetSubscriptionIdNil() {
	o.SubscriptionId.Set(nil)
}

// UnsetSubscriptionId ensures that no value is present for SubscriptionId, not even an explicit nil
func (o *RunTransportTaskBaseVM) UnsetSubscriptionId() {
	o.SubscriptionId.Unset()
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *RunTransportTaskBaseVM) GetType() TaskType {
	if o == nil || o.Type == nil {
		var ret TaskType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RunTransportTaskBaseVM) GetTypeOk() (*TaskType, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *RunTransportTaskBaseVM) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given TaskType and assigns it to the Type field.
func (o *RunTransportTaskBaseVM) SetType(v TaskType) {
	o.Type = &v
}

func (o RunTransportTaskBaseVM) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Files != nil {
		toSerialize["files"] = o.Files
	}
	if o.SubscriptionId.IsSet() {
		toSerialize["subscriptionId"] = o.SubscriptionId.Get()
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableRunTransportTaskBaseVM struct {
	value *RunTransportTaskBaseVM
	isSet bool
}

func (v NullableRunTransportTaskBaseVM) Get() *RunTransportTaskBaseVM {
	return v.value
}

func (v *NullableRunTransportTaskBaseVM) Set(val *RunTransportTaskBaseVM) {
	v.value = val
	v.isSet = true
}

func (v NullableRunTransportTaskBaseVM) IsSet() bool {
	return v.isSet
}

func (v *NullableRunTransportTaskBaseVM) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRunTransportTaskBaseVM(val *RunTransportTaskBaseVM) *NullableRunTransportTaskBaseVM {
	return &NullableRunTransportTaskBaseVM{value: val, isSet: true}
}

func (v NullableRunTransportTaskBaseVM) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRunTransportTaskBaseVM) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

