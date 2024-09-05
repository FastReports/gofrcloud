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

// checks if the TemplateCreateFormVM type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TemplateCreateFormVM{}

// TemplateCreateFormVM struct for TemplateCreateFormVM
type TemplateCreateFormVM struct {
	FileCreateFormVM
	T string `json:"$t"`
}

type _TemplateCreateFormVM TemplateCreateFormVM

// NewTemplateCreateFormVM instantiates a new TemplateCreateFormVM object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTemplateCreateFormVM(t string, fileContent *os.File) *TemplateCreateFormVM {
	this := TemplateCreateFormVM{}
	this.FileContent = fileContent
	this.T = t
	return &this
}

// NewTemplateCreateFormVMWithDefaults instantiates a new TemplateCreateFormVM object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTemplateCreateFormVMWithDefaults() *TemplateCreateFormVM {
	this := TemplateCreateFormVM{}
	return &this
}

// GetT returns the T field value
func (o *TemplateCreateFormVM) GetT() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.T
}

// GetTOk returns a tuple with the T field value
// and a boolean to check if the value has been set.
func (o *TemplateCreateFormVM) GetTOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.T, true
}

// SetT sets field value
func (o *TemplateCreateFormVM) SetT(v string) {
	o.T = v
}

func (o TemplateCreateFormVM) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TemplateCreateFormVM) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedFileCreateFormVM, errFileCreateFormVM := json.Marshal(o.FileCreateFormVM)
	if errFileCreateFormVM != nil {
		return map[string]interface{}{}, errFileCreateFormVM
	}
	errFileCreateFormVM = json.Unmarshal([]byte(serializedFileCreateFormVM), &toSerialize)
	if errFileCreateFormVM != nil {
		return map[string]interface{}{}, errFileCreateFormVM
	}
	toSerialize["$t"] = o.T
	return toSerialize, nil
}

func (o *TemplateCreateFormVM) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"$t",
		"fileContent",
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

	varTemplateCreateFormVM := _TemplateCreateFormVM{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTemplateCreateFormVM)

	if err != nil {
		return err
	}

	*o = TemplateCreateFormVM(varTemplateCreateFormVM)

	return err
}

type NullableTemplateCreateFormVM struct {
	value *TemplateCreateFormVM
	isSet bool
}

func (v NullableTemplateCreateFormVM) Get() *TemplateCreateFormVM {
	return v.value
}

func (v *NullableTemplateCreateFormVM) Set(val *TemplateCreateFormVM) {
	v.value = val
	v.isSet = true
}

func (v NullableTemplateCreateFormVM) IsSet() bool {
	return v.isSet
}

func (v *NullableTemplateCreateFormVM) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTemplateCreateFormVM(val *TemplateCreateFormVM) *NullableTemplateCreateFormVM {
	return &NullableTemplateCreateFormVM{value: val, isSet: true}
}

func (v NullableTemplateCreateFormVM) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTemplateCreateFormVM) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

