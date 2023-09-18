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

// checks if the UpdateFileContentVM type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateFileContentVM{}

// UpdateFileContentVM struct for UpdateFileContentVM
type UpdateFileContentVM struct {
	Content string `json:"content"`
}

// NewUpdateFileContentVM instantiates a new UpdateFileContentVM object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateFileContentVM(content string) *UpdateFileContentVM {
	this := UpdateFileContentVM{}
	this.Content = content
	return &this
}

// NewUpdateFileContentVMWithDefaults instantiates a new UpdateFileContentVM object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateFileContentVMWithDefaults() *UpdateFileContentVM {
	this := UpdateFileContentVM{}
	return &this
}

// GetContent returns the Content field value
func (o *UpdateFileContentVM) GetContent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Content
}

// GetContentOk returns a tuple with the Content field value
// and a boolean to check if the value has been set.
func (o *UpdateFileContentVM) GetContentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Content, true
}

// SetContent sets field value
func (o *UpdateFileContentVM) SetContent(v string) {
	o.Content = v
}

func (o UpdateFileContentVM) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateFileContentVM) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["content"] = o.Content
	return toSerialize, nil
}

type NullableUpdateFileContentVM struct {
	value *UpdateFileContentVM
	isSet bool
}

func (v NullableUpdateFileContentVM) Get() *UpdateFileContentVM {
	return v.value
}

func (v *NullableUpdateFileContentVM) Set(val *UpdateFileContentVM) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateFileContentVM) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateFileContentVM) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateFileContentVM(val *UpdateFileContentVM) *NullableUpdateFileContentVM {
	return &NullableUpdateFileContentVM{value: val, isSet: true}
}

func (v NullableUpdateFileContentVM) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateFileContentVM) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

