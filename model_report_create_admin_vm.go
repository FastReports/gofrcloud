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

// ReportCreateAdminVM struct for ReportCreateAdminVM
type ReportCreateAdminVM struct {
	OwnerId string `json:"ownerId"`
	ParentId string `json:"parentId"`
	Name NullableString `json:"name,omitempty"`
	Tags []string `json:"tags,omitempty"`
	Icon NullableString `json:"icon,omitempty"`
	Content NullableString `json:"content,omitempty"`
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
	if o == nil  {
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
	if o == nil  {
		return nil, false
	}
	return &o.ParentId, true
}

// SetParentId sets field value
func (o *ReportCreateAdminVM) SetParentId(v string) {
	o.ParentId = v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReportCreateAdminVM) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReportCreateAdminVM) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *ReportCreateAdminVM) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *ReportCreateAdminVM) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *ReportCreateAdminVM) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *ReportCreateAdminVM) UnsetName() {
	o.Name.Unset()
}

// GetTags returns the Tags field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReportCreateAdminVM) GetTags() []string {
	if o == nil  {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReportCreateAdminVM) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return &o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *ReportCreateAdminVM) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *ReportCreateAdminVM) SetTags(v []string) {
	o.Tags = v
}

// GetIcon returns the Icon field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReportCreateAdminVM) GetIcon() string {
	if o == nil || o.Icon.Get() == nil {
		var ret string
		return ret
	}
	return *o.Icon.Get()
}

// GetIconOk returns a tuple with the Icon field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReportCreateAdminVM) GetIconOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Icon.Get(), o.Icon.IsSet()
}

// HasIcon returns a boolean if a field has been set.
func (o *ReportCreateAdminVM) HasIcon() bool {
	if o != nil && o.Icon.IsSet() {
		return true
	}

	return false
}

// SetIcon gets a reference to the given NullableString and assigns it to the Icon field.
func (o *ReportCreateAdminVM) SetIcon(v string) {
	o.Icon.Set(&v)
}
// SetIconNil sets the value for Icon to be an explicit nil
func (o *ReportCreateAdminVM) SetIconNil() {
	o.Icon.Set(nil)
}

// UnsetIcon ensures that no value is present for Icon, not even an explicit nil
func (o *ReportCreateAdminVM) UnsetIcon() {
	o.Icon.Unset()
}

// GetContent returns the Content field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReportCreateAdminVM) GetContent() string {
	if o == nil || o.Content.Get() == nil {
		var ret string
		return ret
	}
	return *o.Content.Get()
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReportCreateAdminVM) GetContentOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Content.Get(), o.Content.IsSet()
}

// HasContent returns a boolean if a field has been set.
func (o *ReportCreateAdminVM) HasContent() bool {
	if o != nil && o.Content.IsSet() {
		return true
	}

	return false
}

// SetContent gets a reference to the given NullableString and assigns it to the Content field.
func (o *ReportCreateAdminVM) SetContent(v string) {
	o.Content.Set(&v)
}
// SetContentNil sets the value for Content to be an explicit nil
func (o *ReportCreateAdminVM) SetContentNil() {
	o.Content.Set(nil)
}

// UnsetContent ensures that no value is present for Content, not even an explicit nil
func (o *ReportCreateAdminVM) UnsetContent() {
	o.Content.Unset()
}

func (o ReportCreateAdminVM) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ownerId"] = o.OwnerId
	}
	if true {
		toSerialize["parentId"] = o.ParentId
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if o.Icon.IsSet() {
		toSerialize["icon"] = o.Icon.Get()
	}
	if o.Content.IsSet() {
		toSerialize["content"] = o.Content.Get()
	}
	return json.Marshal(toSerialize)
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

