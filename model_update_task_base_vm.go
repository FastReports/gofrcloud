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

// checks if the UpdateTaskBaseVM type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateTaskBaseVM{}

// UpdateTaskBaseVM struct for UpdateTaskBaseVM
type UpdateTaskBaseVM struct {
	CronExpression NullableString `json:"cronExpression,omitempty"`
	DelayedRunTime NullableTime `json:"delayedRunTime,omitempty"`
	Name NullableString `json:"name,omitempty"`
	T string `json:"$t"`
}

// NewUpdateTaskBaseVM instantiates a new UpdateTaskBaseVM object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateTaskBaseVM(t string) *UpdateTaskBaseVM {
	this := UpdateTaskBaseVM{}
	this.T = t
	return &this
}

// NewUpdateTaskBaseVMWithDefaults instantiates a new UpdateTaskBaseVM object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateTaskBaseVMWithDefaults() *UpdateTaskBaseVM {
	this := UpdateTaskBaseVM{}
	return &this
}

// GetCronExpression returns the CronExpression field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateTaskBaseVM) GetCronExpression() string {
	if o == nil || IsNil(o.CronExpression.Get()) {
		var ret string
		return ret
	}
	return *o.CronExpression.Get()
}

// GetCronExpressionOk returns a tuple with the CronExpression field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateTaskBaseVM) GetCronExpressionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CronExpression.Get(), o.CronExpression.IsSet()
}

// HasCronExpression returns a boolean if a field has been set.
func (o *UpdateTaskBaseVM) HasCronExpression() bool {
	if o != nil && o.CronExpression.IsSet() {
		return true
	}

	return false
}

// SetCronExpression gets a reference to the given NullableString and assigns it to the CronExpression field.
func (o *UpdateTaskBaseVM) SetCronExpression(v string) {
	o.CronExpression.Set(&v)
}
// SetCronExpressionNil sets the value for CronExpression to be an explicit nil
func (o *UpdateTaskBaseVM) SetCronExpressionNil() {
	o.CronExpression.Set(nil)
}

// UnsetCronExpression ensures that no value is present for CronExpression, not even an explicit nil
func (o *UpdateTaskBaseVM) UnsetCronExpression() {
	o.CronExpression.Unset()
}

// GetDelayedRunTime returns the DelayedRunTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateTaskBaseVM) GetDelayedRunTime() time.Time {
	if o == nil || IsNil(o.DelayedRunTime.Get()) {
		var ret time.Time
		return ret
	}
	return *o.DelayedRunTime.Get()
}

// GetDelayedRunTimeOk returns a tuple with the DelayedRunTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateTaskBaseVM) GetDelayedRunTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.DelayedRunTime.Get(), o.DelayedRunTime.IsSet()
}

// HasDelayedRunTime returns a boolean if a field has been set.
func (o *UpdateTaskBaseVM) HasDelayedRunTime() bool {
	if o != nil && o.DelayedRunTime.IsSet() {
		return true
	}

	return false
}

// SetDelayedRunTime gets a reference to the given NullableTime and assigns it to the DelayedRunTime field.
func (o *UpdateTaskBaseVM) SetDelayedRunTime(v time.Time) {
	o.DelayedRunTime.Set(&v)
}
// SetDelayedRunTimeNil sets the value for DelayedRunTime to be an explicit nil
func (o *UpdateTaskBaseVM) SetDelayedRunTimeNil() {
	o.DelayedRunTime.Set(nil)
}

// UnsetDelayedRunTime ensures that no value is present for DelayedRunTime, not even an explicit nil
func (o *UpdateTaskBaseVM) UnsetDelayedRunTime() {
	o.DelayedRunTime.Unset()
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateTaskBaseVM) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateTaskBaseVM) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *UpdateTaskBaseVM) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *UpdateTaskBaseVM) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *UpdateTaskBaseVM) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *UpdateTaskBaseVM) UnsetName() {
	o.Name.Unset()
}

// GetT returns the T field value
func (o *UpdateTaskBaseVM) GetT() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.T
}

// GetTOk returns a tuple with the T field value
// and a boolean to check if the value has been set.
func (o *UpdateTaskBaseVM) GetTOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.T, true
}

// SetT sets field value
func (o *UpdateTaskBaseVM) SetT(v string) {
	o.T = v
}

func (o UpdateTaskBaseVM) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateTaskBaseVM) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.CronExpression.IsSet() {
		toSerialize["cronExpression"] = o.CronExpression.Get()
	}
	if o.DelayedRunTime.IsSet() {
		toSerialize["delayedRunTime"] = o.DelayedRunTime.Get()
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	toSerialize["$t"] = o.T
	return toSerialize, nil
}

type NullableUpdateTaskBaseVM struct {
	value *UpdateTaskBaseVM
	isSet bool
}

func (v NullableUpdateTaskBaseVM) Get() *UpdateTaskBaseVM {
	return v.value
}

func (v *NullableUpdateTaskBaseVM) Set(val *UpdateTaskBaseVM) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateTaskBaseVM) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateTaskBaseVM) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateTaskBaseVM(val *UpdateTaskBaseVM) *NullableUpdateTaskBaseVM {
	return &NullableUpdateTaskBaseVM{value: val, isSet: true}
}

func (v NullableUpdateTaskBaseVM) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateTaskBaseVM) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


