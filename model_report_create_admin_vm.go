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

// checks if the ReportCreateAdminVM type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReportCreateAdminVM{}

// ReportCreateAdminVM struct for ReportCreateAdminVM
type ReportCreateAdminVM struct {
	OwnerId string `json:"ownerId"`
	ParentId string `json:"parentId"`
}

// NewReportCreateAdminVM instantiates a new ReportCreateAdminVM object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReportCreateAdminVM(ownerId string, parentId string) *ReportCreateAdminVM {
	this := ReportCreateAdminVM{}
	return &this
}

// NewReportCreateAdminVMWithDefaults instantiates a new ReportCreateAdminVM object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReportCreateAdminVMWithDefaults() *ReportCreateAdminVM {
	this := ReportCreateAdminVM{}
	return &this
}

// GetOwnerId returns the OwnerId field value
func (o *ReportCreateAdminVM) GetOwnerId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OwnerId
}

// GetOwnerIdOk returns a tuple with the OwnerId field value
// and a boolean to check if the value has been set.
func (o *ReportCreateAdminVM) GetOwnerIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OwnerId, true
}

// SetOwnerId sets field value
func (o *ReportCreateAdminVM) SetOwnerId(v string) {
	o.OwnerId = v
}

// GetParentId returns the ParentId field value
func (o *ReportCreateAdminVM) GetParentId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ParentId
}

// GetParentIdOk returns a tuple with the ParentId field value
// and a boolean to check if the value has been set.
func (o *ReportCreateAdminVM) GetParentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ParentId, true
}

// SetParentId sets field value
func (o *ReportCreateAdminVM) SetParentId(v string) {
	o.ParentId = v
}

func (o ReportCreateAdminVM) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReportCreateAdminVM) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["ownerId"] = o.OwnerId
	toSerialize["parentId"] = o.ParentId
	return toSerialize, nil
}

type NullableReportCreateAdminVM struct {
	value *ReportCreateAdminVM
	isSet bool
}

func (v NullableReportCreateAdminVM) Get() *ReportCreateAdminVM {
	return v.value
}

func (v *NullableReportCreateAdminVM) Set(val *ReportCreateAdminVM) {
	v.value = val
	v.isSet = true
}

func (v NullableReportCreateAdminVM) IsSet() bool {
	return v.isSet
}

func (v *NullableReportCreateAdminVM) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReportCreateAdminVM(val *ReportCreateAdminVM) *NullableReportCreateAdminVM {
	return &NullableReportCreateAdminVM{value: val, isSet: true}
}

func (v NullableReportCreateAdminVM) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReportCreateAdminVM) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


