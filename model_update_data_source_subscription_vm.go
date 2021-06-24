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

// UpdateDataSourceSubscriptionVM struct for UpdateDataSourceSubscriptionVM
type UpdateDataSourceSubscriptionVM struct {
	SubscriptionId string `json:"subscriptionId"`
}

// NewUpdateDataSourceSubscriptionVM instantiates a new UpdateDataSourceSubscriptionVM object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateDataSourceSubscriptionVM(subscriptionId string) *UpdateDataSourceSubscriptionVM {
	this := UpdateDataSourceSubscriptionVM{}
	this.SubscriptionId = subscriptionId
	return &this
}

// NewUpdateDataSourceSubscriptionVMWithDefaults instantiates a new UpdateDataSourceSubscriptionVM object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateDataSourceSubscriptionVMWithDefaults() *UpdateDataSourceSubscriptionVM {
	this := UpdateDataSourceSubscriptionVM{}
	return &this
}

// GetSubscriptionId returns the SubscriptionId field value
func (o *UpdateDataSourceSubscriptionVM) GetSubscriptionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SubscriptionId
}

// GetSubscriptionIdOk returns a tuple with the SubscriptionId field value
// and a boolean to check if the value has been set.
func (o *UpdateDataSourceSubscriptionVM) GetSubscriptionIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SubscriptionId, true
}

// SetSubscriptionId sets field value
func (o *UpdateDataSourceSubscriptionVM) SetSubscriptionId(v string) {
	o.SubscriptionId = v
}

func (o UpdateDataSourceSubscriptionVM) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["subscriptionId"] = o.SubscriptionId
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateDataSourceSubscriptionVM struct {
	value *UpdateDataSourceSubscriptionVM
	isSet bool
}

func (v NullableUpdateDataSourceSubscriptionVM) Get() *UpdateDataSourceSubscriptionVM {
	return v.value
}

func (v *NullableUpdateDataSourceSubscriptionVM) Set(val *UpdateDataSourceSubscriptionVM) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateDataSourceSubscriptionVM) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateDataSourceSubscriptionVM) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateDataSourceSubscriptionVM(val *UpdateDataSourceSubscriptionVM) *NullableUpdateDataSourceSubscriptionVM {
	return &NullableUpdateDataSourceSubscriptionVM{value: val, isSet: true}
}

func (v NullableUpdateDataSourceSubscriptionVM) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateDataSourceSubscriptionVM) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


