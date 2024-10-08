/*
FastReport Cloud

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gofrcloud

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the AdminReportFolderCreateVM type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AdminReportFolderCreateVM{}

// AdminReportFolderCreateVM struct for AdminReportFolderCreateVM
type AdminReportFolderCreateVM struct {
	AdminFolderCreateVM
	T string `json:"$t"`
}

type _AdminReportFolderCreateVM AdminReportFolderCreateVM

// NewAdminReportFolderCreateVM instantiates a new AdminReportFolderCreateVM object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdminReportFolderCreateVM(t string, parentId string, ownerId string) *AdminReportFolderCreateVM {
	this := AdminReportFolderCreateVM{}
	this.ParentId = parentId
	this.OwnerId = ownerId
	this.T = t
	return &this
}

// NewAdminReportFolderCreateVMWithDefaults instantiates a new AdminReportFolderCreateVM object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdminReportFolderCreateVMWithDefaults() *AdminReportFolderCreateVM {
	this := AdminReportFolderCreateVM{}
	return &this
}

// GetT returns the T field value
func (o *AdminReportFolderCreateVM) GetT() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.T
}

// GetTOk returns a tuple with the T field value
// and a boolean to check if the value has been set.
func (o *AdminReportFolderCreateVM) GetTOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.T, true
}

// SetT sets field value
func (o *AdminReportFolderCreateVM) SetT(v string) {
	o.T = v
}

func (o AdminReportFolderCreateVM) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AdminReportFolderCreateVM) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedAdminFolderCreateVM, errAdminFolderCreateVM := json.Marshal(o.AdminFolderCreateVM)
	if errAdminFolderCreateVM != nil {
		return map[string]interface{}{}, errAdminFolderCreateVM
	}
	errAdminFolderCreateVM = json.Unmarshal([]byte(serializedAdminFolderCreateVM), &toSerialize)
	if errAdminFolderCreateVM != nil {
		return map[string]interface{}{}, errAdminFolderCreateVM
	}
	toSerialize["$t"] = o.T
	return toSerialize, nil
}

func (o *AdminReportFolderCreateVM) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"$t",
		"parentId",
		"ownerId",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varAdminReportFolderCreateVM := _AdminReportFolderCreateVM{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAdminReportFolderCreateVM)

	if err != nil {
		return err
	}

	*o = AdminReportFolderCreateVM(varAdminReportFolderCreateVM)

	return err
}

type NullableAdminReportFolderCreateVM struct {
	value *AdminReportFolderCreateVM
	isSet bool
}

func (v NullableAdminReportFolderCreateVM) Get() *AdminReportFolderCreateVM {
	return v.value
}

func (v *NullableAdminReportFolderCreateVM) Set(val *AdminReportFolderCreateVM) {
	v.value = val
	v.isSet = true
}

func (v NullableAdminReportFolderCreateVM) IsSet() bool {
	return v.isSet
}

func (v *NullableAdminReportFolderCreateVM) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdminReportFolderCreateVM(val *AdminReportFolderCreateVM) *NullableAdminReportFolderCreateVM {
	return &NullableAdminReportFolderCreateVM{value: val, isSet: true}
}

func (v NullableAdminReportFolderCreateVM) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdminReportFolderCreateVM) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


