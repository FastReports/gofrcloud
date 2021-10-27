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

// SubscriptionUsersVM struct for SubscriptionUsersVM
type SubscriptionUsersVM struct {
	Users []SubscriptionUserVM `json:"users,omitempty"`
	Count *int64 `json:"count,omitempty"`
	Take *int32 `json:"take,omitempty"`
	Skip *int32 `json:"skip,omitempty"`
}

// NewSubscriptionUsersVM instantiates a new SubscriptionUsersVM object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscriptionUsersVM() *SubscriptionUsersVM {
	this := SubscriptionUsersVM{}
	return &this
}

// NewSubscriptionUsersVMWithDefaults instantiates a new SubscriptionUsersVM object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionUsersVMWithDefaults() *SubscriptionUsersVM {
	this := SubscriptionUsersVM{}
	return &this
}

// GetUsers returns the Users field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SubscriptionUsersVM) GetUsers() []SubscriptionUserVM {
	if o == nil  {
		var ret []SubscriptionUserVM
		return ret
	}
	return o.Users
}

// GetUsersOk returns a tuple with the Users field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SubscriptionUsersVM) GetUsersOk() (*[]SubscriptionUserVM, bool) {
	if o == nil || o.Users == nil {
		return nil, false
	}
	return &o.Users, true
}

// HasUsers returns a boolean if a field has been set.
func (o *SubscriptionUsersVM) HasUsers() bool {
	if o != nil && o.Users != nil {
		return true
	}

	return false
}

// SetUsers gets a reference to the given []SubscriptionUserVM and assigns it to the Users field.
func (o *SubscriptionUsersVM) SetUsers(v []SubscriptionUserVM) {
	o.Users = v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *SubscriptionUsersVM) GetCount() int64 {
	if o == nil || o.Count == nil {
		var ret int64
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionUsersVM) GetCountOk() (*int64, bool) {
	if o == nil || o.Count == nil {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *SubscriptionUsersVM) HasCount() bool {
	if o != nil && o.Count != nil {
		return true
	}

	return false
}

// SetCount gets a reference to the given int64 and assigns it to the Count field.
func (o *SubscriptionUsersVM) SetCount(v int64) {
	o.Count = &v
}

// GetTake returns the Take field value if set, zero value otherwise.
func (o *SubscriptionUsersVM) GetTake() int32 {
	if o == nil || o.Take == nil {
		var ret int32
		return ret
	}
	return *o.Take
}

// GetTakeOk returns a tuple with the Take field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionUsersVM) GetTakeOk() (*int32, bool) {
	if o == nil || o.Take == nil {
		return nil, false
	}
	return o.Take, true
}

// HasTake returns a boolean if a field has been set.
func (o *SubscriptionUsersVM) HasTake() bool {
	if o != nil && o.Take != nil {
		return true
	}

	return false
}

// SetTake gets a reference to the given int32 and assigns it to the Take field.
func (o *SubscriptionUsersVM) SetTake(v int32) {
	o.Take = &v
}

// GetSkip returns the Skip field value if set, zero value otherwise.
func (o *SubscriptionUsersVM) GetSkip() int32 {
	if o == nil || o.Skip == nil {
		var ret int32
		return ret
	}
	return *o.Skip
}

// GetSkipOk returns a tuple with the Skip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionUsersVM) GetSkipOk() (*int32, bool) {
	if o == nil || o.Skip == nil {
		return nil, false
	}
	return o.Skip, true
}

// HasSkip returns a boolean if a field has been set.
func (o *SubscriptionUsersVM) HasSkip() bool {
	if o != nil && o.Skip != nil {
		return true
	}

	return false
}

// SetSkip gets a reference to the given int32 and assigns it to the Skip field.
func (o *SubscriptionUsersVM) SetSkip(v int32) {
	o.Skip = &v
}

func (o SubscriptionUsersVM) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Users != nil {
		toSerialize["users"] = o.Users
	}
	if o.Count != nil {
		toSerialize["count"] = o.Count
	}
	if o.Take != nil {
		toSerialize["take"] = o.Take
	}
	if o.Skip != nil {
		toSerialize["skip"] = o.Skip
	}
	return json.Marshal(toSerialize)
}

type NullableSubscriptionUsersVM struct {
	value *SubscriptionUsersVM
	isSet bool
}

func (v NullableSubscriptionUsersVM) Get() *SubscriptionUsersVM {
	return v.value
}

func (v *NullableSubscriptionUsersVM) Set(val *SubscriptionUsersVM) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionUsersVM) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionUsersVM) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionUsersVM(val *SubscriptionUsersVM) *NullableSubscriptionUsersVM {
	return &NullableSubscriptionUsersVM{value: val, isSet: true}
}

func (v NullableSubscriptionUsersVM) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionUsersVM) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


