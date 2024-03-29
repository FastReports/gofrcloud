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

// checks if the CreateThumbnailTemplateTaskVM type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateThumbnailTemplateTaskVM{}

// CreateThumbnailTemplateTaskVM struct for CreateThumbnailTemplateTaskVM
type CreateThumbnailTemplateTaskVM struct {
	TemplateId NullableString `json:"templateId,omitempty"`
}

// NewCreateThumbnailTemplateTaskVM instantiates a new CreateThumbnailTemplateTaskVM object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateThumbnailTemplateTaskVM(t string) *CreateThumbnailTemplateTaskVM {
	this := CreateThumbnailTemplateTaskVM{}
	this.T = t
	return &this
}

// NewCreateThumbnailTemplateTaskVMWithDefaults instantiates a new CreateThumbnailTemplateTaskVM object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateThumbnailTemplateTaskVMWithDefaults() *CreateThumbnailTemplateTaskVM {
	this := CreateThumbnailTemplateTaskVM{}
	return &this
}

// GetTemplateId returns the TemplateId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateThumbnailTemplateTaskVM) GetTemplateId() string {
	if o == nil || IsNil(o.TemplateId.Get()) {
		var ret string
		return ret
	}
	return *o.TemplateId.Get()
}

// GetTemplateIdOk returns a tuple with the TemplateId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateThumbnailTemplateTaskVM) GetTemplateIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TemplateId.Get(), o.TemplateId.IsSet()
}

// HasTemplateId returns a boolean if a field has been set.
func (o *CreateThumbnailTemplateTaskVM) HasTemplateId() bool {
	if o != nil && o.TemplateId.IsSet() {
		return true
	}

	return false
}

// SetTemplateId gets a reference to the given NullableString and assigns it to the TemplateId field.
func (o *CreateThumbnailTemplateTaskVM) SetTemplateId(v string) {
	o.TemplateId.Set(&v)
}
// SetTemplateIdNil sets the value for TemplateId to be an explicit nil
func (o *CreateThumbnailTemplateTaskVM) SetTemplateIdNil() {
	o.TemplateId.Set(nil)
}

// UnsetTemplateId ensures that no value is present for TemplateId, not even an explicit nil
func (o *CreateThumbnailTemplateTaskVM) UnsetTemplateId() {
	o.TemplateId.Unset()
}

func (o CreateThumbnailTemplateTaskVM) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateThumbnailTemplateTaskVM) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.TemplateId.IsSet() {
		toSerialize["templateId"] = o.TemplateId.Get()
	}
	return toSerialize, nil
}

type NullableCreateThumbnailTemplateTaskVM struct {
	value *CreateThumbnailTemplateTaskVM
	isSet bool
}

func (v NullableCreateThumbnailTemplateTaskVM) Get() *CreateThumbnailTemplateTaskVM {
	return v.value
}

func (v *NullableCreateThumbnailTemplateTaskVM) Set(val *CreateThumbnailTemplateTaskVM) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateThumbnailTemplateTaskVM) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateThumbnailTemplateTaskVM) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateThumbnailTemplateTaskVM(val *CreateThumbnailTemplateTaskVM) *NullableCreateThumbnailTemplateTaskVM {
	return &NullableCreateThumbnailTemplateTaskVM{value: val, isSet: true}
}

func (v NullableCreateThumbnailTemplateTaskVM) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateThumbnailTemplateTaskVM) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


