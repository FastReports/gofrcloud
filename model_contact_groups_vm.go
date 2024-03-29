/*
FastReport Cloud

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gofrcloud

import (
	"encoding/json"
)

// checks if the ContactGroupsVM type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContactGroupsVM{}

// ContactGroupsVM struct for ContactGroupsVM
type ContactGroupsVM struct {
	Groups []ContactGroupVM `json:"groups,omitempty"`
	Skip *int32 `json:"skip,omitempty"`
	Take *int32 `json:"take,omitempty"`
	Count *int64 `json:"count,omitempty"`
}

// NewContactGroupsVM instantiates a new ContactGroupsVM object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContactGroupsVM() *ContactGroupsVM {
	this := ContactGroupsVM{}
	return &this
}

// NewContactGroupsVMWithDefaults instantiates a new ContactGroupsVM object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContactGroupsVMWithDefaults() *ContactGroupsVM {
	this := ContactGroupsVM{}
	return &this
}

// GetGroups returns the Groups field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ContactGroupsVM) GetGroups() []ContactGroupVM {
	if o == nil {
		var ret []ContactGroupVM
		return ret
	}
	return o.Groups
}

// GetGroupsOk returns a tuple with the Groups field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ContactGroupsVM) GetGroupsOk() ([]ContactGroupVM, bool) {
	if o == nil || IsNil(o.Groups) {
		return nil, false
	}
	return o.Groups, true
}

// HasGroups returns a boolean if a field has been set.
func (o *ContactGroupsVM) HasGroups() bool {
	if o != nil && IsNil(o.Groups) {
		return true
	}

	return false
}

// SetGroups gets a reference to the given []ContactGroupVM and assigns it to the Groups field.
func (o *ContactGroupsVM) SetGroups(v []ContactGroupVM) {
	o.Groups = v
}

// GetSkip returns the Skip field value if set, zero value otherwise.
func (o *ContactGroupsVM) GetSkip() int32 {
	if o == nil || IsNil(o.Skip) {
		var ret int32
		return ret
	}
	return *o.Skip
}

// GetSkipOk returns a tuple with the Skip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContactGroupsVM) GetSkipOk() (*int32, bool) {
	if o == nil || IsNil(o.Skip) {
		return nil, false
	}
	return o.Skip, true
}

// HasSkip returns a boolean if a field has been set.
func (o *ContactGroupsVM) HasSkip() bool {
	if o != nil && !IsNil(o.Skip) {
		return true
	}

	return false
}

// SetSkip gets a reference to the given int32 and assigns it to the Skip field.
func (o *ContactGroupsVM) SetSkip(v int32) {
	o.Skip = &v
}

// GetTake returns the Take field value if set, zero value otherwise.
func (o *ContactGroupsVM) GetTake() int32 {
	if o == nil || IsNil(o.Take) {
		var ret int32
		return ret
	}
	return *o.Take
}

// GetTakeOk returns a tuple with the Take field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContactGroupsVM) GetTakeOk() (*int32, bool) {
	if o == nil || IsNil(o.Take) {
		return nil, false
	}
	return o.Take, true
}

// HasTake returns a boolean if a field has been set.
func (o *ContactGroupsVM) HasTake() bool {
	if o != nil && !IsNil(o.Take) {
		return true
	}

	return false
}

// SetTake gets a reference to the given int32 and assigns it to the Take field.
func (o *ContactGroupsVM) SetTake(v int32) {
	o.Take = &v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *ContactGroupsVM) GetCount() int64 {
	if o == nil || IsNil(o.Count) {
		var ret int64
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContactGroupsVM) GetCountOk() (*int64, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *ContactGroupsVM) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int64 and assigns it to the Count field.
func (o *ContactGroupsVM) SetCount(v int64) {
	o.Count = &v
}

func (o ContactGroupsVM) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContactGroupsVM) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Groups != nil {
		toSerialize["groups"] = o.Groups
	}
	if !IsNil(o.Skip) {
		toSerialize["skip"] = o.Skip
	}
	if !IsNil(o.Take) {
		toSerialize["take"] = o.Take
	}
	if !IsNil(o.Count) {
		toSerialize["count"] = o.Count
	}
	return toSerialize, nil
}

type NullableContactGroupsVM struct {
	value *ContactGroupsVM
	isSet bool
}

func (v NullableContactGroupsVM) Get() *ContactGroupsVM {
	return v.value
}

func (v *NullableContactGroupsVM) Set(val *ContactGroupsVM) {
	v.value = val
	v.isSet = true
}

func (v NullableContactGroupsVM) IsSet() bool {
	return v.isSet
}

func (v *NullableContactGroupsVM) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContactGroupsVM(val *ContactGroupsVM) *NullableContactGroupsVM {
	return &NullableContactGroupsVM{value: val, isSet: true}
}

func (v NullableContactGroupsVM) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContactGroupsVM) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


