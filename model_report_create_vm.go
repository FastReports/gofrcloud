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

// ReportCreateVM struct for ReportCreateVM
type ReportCreateVM struct {
	TemplateId NullableString `json:"templateId,omitempty"`
	Name NullableString `json:"name,omitempty"`
	Tags []string `json:"tags,omitempty"`
	Icon NullableString `json:"icon,omitempty"`
	Content NullableString `json:"content,omitempty"`
}

// NewReportCreateVM instantiates a new ReportCreateVM object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReportCreateVM() *ReportCreateVM {
	this := ReportCreateVM{}
	return &this
}

// NewReportCreateVMWithDefaults instantiates a new ReportCreateVM object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReportCreateVMWithDefaults() *ReportCreateVM {
	this := ReportCreateVM{}
	return &this
}

// GetTemplateId returns the TemplateId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReportCreateVM) GetTemplateId() string {
	if o == nil || o.TemplateId.Get() == nil {
		var ret string
		return ret
	}
	return *o.TemplateId.Get()
}

// GetTemplateIdOk returns a tuple with the TemplateId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReportCreateVM) GetTemplateIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.TemplateId.Get(), o.TemplateId.IsSet()
}

// HasTemplateId returns a boolean if a field has been set.
func (o *ReportCreateVM) HasTemplateId() bool {
	if o != nil && o.TemplateId.IsSet() {
		return true
	}

	return false
}

// SetTemplateId gets a reference to the given NullableString and assigns it to the TemplateId field.
func (o *ReportCreateVM) SetTemplateId(v string) {
	o.TemplateId.Set(&v)
}
// SetTemplateIdNil sets the value for TemplateId to be an explicit nil
func (o *ReportCreateVM) SetTemplateIdNil() {
	o.TemplateId.Set(nil)
}

// UnsetTemplateId ensures that no value is present for TemplateId, not even an explicit nil
func (o *ReportCreateVM) UnsetTemplateId() {
	o.TemplateId.Unset()
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReportCreateVM) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReportCreateVM) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *ReportCreateVM) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *ReportCreateVM) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *ReportCreateVM) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *ReportCreateVM) UnsetName() {
	o.Name.Unset()
}

// GetTags returns the Tags field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReportCreateVM) GetTags() []string {
	if o == nil  {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReportCreateVM) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return &o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *ReportCreateVM) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *ReportCreateVM) SetTags(v []string) {
	o.Tags = v
}

// GetIcon returns the Icon field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReportCreateVM) GetIcon() string {
	if o == nil || o.Icon.Get() == nil {
		var ret string
		return ret
	}
	return *o.Icon.Get()
}

// GetIconOk returns a tuple with the Icon field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReportCreateVM) GetIconOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Icon.Get(), o.Icon.IsSet()
}

// HasIcon returns a boolean if a field has been set.
func (o *ReportCreateVM) HasIcon() bool {
	if o != nil && o.Icon.IsSet() {
		return true
	}

	return false
}

// SetIcon gets a reference to the given NullableString and assigns it to the Icon field.
func (o *ReportCreateVM) SetIcon(v string) {
	o.Icon.Set(&v)
}
// SetIconNil sets the value for Icon to be an explicit nil
func (o *ReportCreateVM) SetIconNil() {
	o.Icon.Set(nil)
}

// UnsetIcon ensures that no value is present for Icon, not even an explicit nil
func (o *ReportCreateVM) UnsetIcon() {
	o.Icon.Unset()
}

// GetContent returns the Content field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReportCreateVM) GetContent() string {
	if o == nil || o.Content.Get() == nil {
		var ret string
		return ret
	}
	return *o.Content.Get()
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReportCreateVM) GetContentOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Content.Get(), o.Content.IsSet()
}

// HasContent returns a boolean if a field has been set.
func (o *ReportCreateVM) HasContent() bool {
	if o != nil && o.Content.IsSet() {
		return true
	}

	return false
}

// SetContent gets a reference to the given NullableString and assigns it to the Content field.
func (o *ReportCreateVM) SetContent(v string) {
	o.Content.Set(&v)
}
// SetContentNil sets the value for Content to be an explicit nil
func (o *ReportCreateVM) SetContentNil() {
	o.Content.Set(nil)
}

// UnsetContent ensures that no value is present for Content, not even an explicit nil
func (o *ReportCreateVM) UnsetContent() {
	o.Content.Unset()
}

func (o ReportCreateVM) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.TemplateId.IsSet() {
		toSerialize["templateId"] = o.TemplateId.Get()
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

type NullableReportCreateVM struct {
	value *ReportCreateVM
	isSet bool
}

func (v NullableReportCreateVM) Get() *ReportCreateVM {
	return v.value
}

func (v *NullableReportCreateVM) Set(val *ReportCreateVM) {
	v.value = val
	v.isSet = true
}

func (v NullableReportCreateVM) IsSet() bool {
	return v.isSet
}

func (v *NullableReportCreateVM) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReportCreateVM(val *ReportCreateVM) *NullableReportCreateVM {
	return &NullableReportCreateVM{value: val, isSet: true}
}

func (v NullableReportCreateVM) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReportCreateVM) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


