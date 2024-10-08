/*
FastReport Cloud

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gofrcloud

import (
	"encoding/json"
	"time"
)

// checks if the TaskEnd type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TaskEnd{}

// TaskEnd struct for TaskEnd
type TaskEnd struct {
	After NullableInt32 `json:"after,omitempty"`
	On NullableTime `json:"on,omitempty"`
}

// NewTaskEnd instantiates a new TaskEnd object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTaskEnd() *TaskEnd {
	this := TaskEnd{}
	return &this
}

// NewTaskEndWithDefaults instantiates a new TaskEnd object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTaskEndWithDefaults() *TaskEnd {
	this := TaskEnd{}
	return &this
}

// GetAfter returns the After field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TaskEnd) GetAfter() int32 {
	if o == nil || IsNil(o.After.Get()) {
		var ret int32
		return ret
	}
	return *o.After.Get()
}

// GetAfterOk returns a tuple with the After field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TaskEnd) GetAfterOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.After.Get(), o.After.IsSet()
}

// HasAfter returns a boolean if a field has been set.
func (o *TaskEnd) HasAfter() bool {
	if o != nil && o.After.IsSet() {
		return true
	}

	return false
}

// SetAfter gets a reference to the given NullableInt32 and assigns it to the After field.
func (o *TaskEnd) SetAfter(v int32) {
	o.After.Set(&v)
}
// SetAfterNil sets the value for After to be an explicit nil
func (o *TaskEnd) SetAfterNil() {
	o.After.Set(nil)
}

// UnsetAfter ensures that no value is present for After, not even an explicit nil
func (o *TaskEnd) UnsetAfter() {
	o.After.Unset()
}

// GetOn returns the On field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TaskEnd) GetOn() time.Time {
	if o == nil || IsNil(o.On.Get()) {
		var ret time.Time
		return ret
	}
	return *o.On.Get()
}

// GetOnOk returns a tuple with the On field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TaskEnd) GetOnOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.On.Get(), o.On.IsSet()
}

// HasOn returns a boolean if a field has been set.
func (o *TaskEnd) HasOn() bool {
	if o != nil && o.On.IsSet() {
		return true
	}

	return false
}

// SetOn gets a reference to the given NullableTime and assigns it to the On field.
func (o *TaskEnd) SetOn(v time.Time) {
	o.On.Set(&v)
}
// SetOnNil sets the value for On to be an explicit nil
func (o *TaskEnd) SetOnNil() {
	o.On.Set(nil)
}

// UnsetOn ensures that no value is present for On, not even an explicit nil
func (o *TaskEnd) UnsetOn() {
	o.On.Unset()
}

func (o TaskEnd) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TaskEnd) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.After.IsSet() {
		toSerialize["after"] = o.After.Get()
	}
	if o.On.IsSet() {
		toSerialize["on"] = o.On.Get()
	}
	return toSerialize, nil
}

type NullableTaskEnd struct {
	value *TaskEnd
	isSet bool
}

func (v NullableTaskEnd) Get() *TaskEnd {
	return v.value
}

func (v *NullableTaskEnd) Set(val *TaskEnd) {
	v.value = val
	v.isSet = true
}

func (v NullableTaskEnd) IsSet() bool {
	return v.isSet
}

func (v *NullableTaskEnd) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTaskEnd(val *TaskEnd) *NullableTaskEnd {
	return &NullableTaskEnd{value: val, isSet: true}
}

func (v NullableTaskEnd) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTaskEnd) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


