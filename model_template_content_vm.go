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

// checks if the TemplateContentVM type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TemplateContentVM{}

// TemplateContentVM struct for TemplateContentVM
type TemplateContentVM struct {
	CloudBaseVM
	Content NullableString `json:"content,omitempty"`
	T string `json:"$t"`
}

type _TemplateContentVM TemplateContentVM

// NewTemplateContentVM instantiates a new TemplateContentVM object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTemplateContentVM(t string) *TemplateContentVM {
	this := TemplateContentVM{}
	this.T = t
	return &this
}

// NewTemplateContentVMWithDefaults instantiates a new TemplateContentVM object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTemplateContentVMWithDefaults() *TemplateContentVM {
	this := TemplateContentVM{}
	return &this
}

// GetContent returns the Content field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TemplateContentVM) GetContent() string {
	if o == nil || IsNil(o.Content.Get()) {
		var ret string
		return ret
	}
	return *o.Content.Get()
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TemplateContentVM) GetContentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Content.Get(), o.Content.IsSet()
}

// HasContent returns a boolean if a field has been set.
func (o *TemplateContentVM) HasContent() bool {
	if o != nil && o.Content.IsSet() {
		return true
	}

	return false
}

// SetContent gets a reference to the given NullableString and assigns it to the Content field.
func (o *TemplateContentVM) SetContent(v string) {
	o.Content.Set(&v)
}
// SetContentNil sets the value for Content to be an explicit nil
func (o *TemplateContentVM) SetContentNil() {
	o.Content.Set(nil)
}

// UnsetContent ensures that no value is present for Content, not even an explicit nil
func (o *TemplateContentVM) UnsetContent() {
	o.Content.Unset()
}

// GetT returns the T field value
func (o *TemplateContentVM) GetT() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.T
}

// GetTOk returns a tuple with the T field value
// and a boolean to check if the value has been set.
func (o *TemplateContentVM) GetTOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.T, true
}

// SetT sets field value
func (o *TemplateContentVM) SetT(v string) {
	o.T = v
}

func (o TemplateContentVM) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TemplateContentVM) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedCloudBaseVM, errCloudBaseVM := json.Marshal(o.CloudBaseVM)
	if errCloudBaseVM != nil {
		return map[string]interface{}{}, errCloudBaseVM
	}
	errCloudBaseVM = json.Unmarshal([]byte(serializedCloudBaseVM), &toSerialize)
	if errCloudBaseVM != nil {
		return map[string]interface{}{}, errCloudBaseVM
	}
	if o.Content.IsSet() {
		toSerialize["content"] = o.Content.Get()
	}
	toSerialize["$t"] = o.T
	return toSerialize, nil
}

func (o *TemplateContentVM) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"$t",
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

	varTemplateContentVM := _TemplateContentVM{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTemplateContentVM)

	if err != nil {
		return err
	}

	*o = TemplateContentVM(varTemplateContentVM)

	return err
}

type NullableTemplateContentVM struct {
	value *TemplateContentVM
	isSet bool
}

func (v NullableTemplateContentVM) Get() *TemplateContentVM {
	return v.value
}

func (v *NullableTemplateContentVM) Set(val *TemplateContentVM) {
	v.value = val
	v.isSet = true
}

func (v NullableTemplateContentVM) IsSet() bool {
	return v.isSet
}

func (v *NullableTemplateContentVM) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTemplateContentVM(val *TemplateContentVM) *NullableTemplateContentVM {
	return &NullableTemplateContentVM{value: val, isSet: true}
}

func (v NullableTemplateContentVM) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTemplateContentVM) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


