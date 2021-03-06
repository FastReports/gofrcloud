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

// ExportFolderCreateVM struct for ExportFolderCreateVM
type ExportFolderCreateVM struct {
	Name NullableString `json:"name,omitempty"`
	Tags []string `json:"tags,omitempty"`
	Icon NullableString `json:"icon,omitempty"`
}

// NewExportFolderCreateVM instantiates a new ExportFolderCreateVM object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExportFolderCreateVM() *ExportFolderCreateVM {
	this := ExportFolderCreateVM{}
	return &this
}

// NewExportFolderCreateVMWithDefaults instantiates a new ExportFolderCreateVM object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExportFolderCreateVMWithDefaults() *ExportFolderCreateVM {
	this := ExportFolderCreateVM{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExportFolderCreateVM) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExportFolderCreateVM) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *ExportFolderCreateVM) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *ExportFolderCreateVM) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *ExportFolderCreateVM) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *ExportFolderCreateVM) UnsetName() {
	o.Name.Unset()
}

// GetTags returns the Tags field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExportFolderCreateVM) GetTags() []string {
	if o == nil  {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExportFolderCreateVM) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return &o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *ExportFolderCreateVM) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *ExportFolderCreateVM) SetTags(v []string) {
	o.Tags = v
}

// GetIcon returns the Icon field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExportFolderCreateVM) GetIcon() string {
	if o == nil || o.Icon.Get() == nil {
		var ret string
		return ret
	}
	return *o.Icon.Get()
}

// GetIconOk returns a tuple with the Icon field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExportFolderCreateVM) GetIconOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Icon.Get(), o.Icon.IsSet()
}

// HasIcon returns a boolean if a field has been set.
func (o *ExportFolderCreateVM) HasIcon() bool {
	if o != nil && o.Icon.IsSet() {
		return true
	}

	return false
}

// SetIcon gets a reference to the given NullableString and assigns it to the Icon field.
func (o *ExportFolderCreateVM) SetIcon(v string) {
	o.Icon.Set(&v)
}
// SetIconNil sets the value for Icon to be an explicit nil
func (o *ExportFolderCreateVM) SetIconNil() {
	o.Icon.Set(nil)
}

// UnsetIcon ensures that no value is present for Icon, not even an explicit nil
func (o *ExportFolderCreateVM) UnsetIcon() {
	o.Icon.Unset()
}

func (o ExportFolderCreateVM) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if o.Icon.IsSet() {
		toSerialize["icon"] = o.Icon.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableExportFolderCreateVM struct {
	value *ExportFolderCreateVM
	isSet bool
}

func (v NullableExportFolderCreateVM) Get() *ExportFolderCreateVM {
	return v.value
}

func (v *NullableExportFolderCreateVM) Set(val *ExportFolderCreateVM) {
	v.value = val
	v.isSet = true
}

func (v NullableExportFolderCreateVM) IsSet() bool {
	return v.isSet
}

func (v *NullableExportFolderCreateVM) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExportFolderCreateVM(val *ExportFolderCreateVM) *NullableExportFolderCreateVM {
	return &NullableExportFolderCreateVM{value: val, isSet: true}
}

func (v NullableExportFolderCreateVM) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExportFolderCreateVM) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


