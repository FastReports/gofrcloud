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

// checks if the CreateTransportTaskBaseVM type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateTransportTaskBaseVM{}

// CreateTransportTaskBaseVM struct for CreateTransportTaskBaseVM
type CreateTransportTaskBaseVM struct {
	CreateTaskBaseVM
	InputFile *InputFileVM `json:"inputFile,omitempty"`
	T string `json:"$t"`
}

type _CreateTransportTaskBaseVM CreateTransportTaskBaseVM

// NewCreateTransportTaskBaseVM instantiates a new CreateTransportTaskBaseVM object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateTransportTaskBaseVM(t string) *CreateTransportTaskBaseVM {
	this := CreateTransportTaskBaseVM{}
	this.T = t
	return &this
}

// NewCreateTransportTaskBaseVMWithDefaults instantiates a new CreateTransportTaskBaseVM object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateTransportTaskBaseVMWithDefaults() *CreateTransportTaskBaseVM {
	this := CreateTransportTaskBaseVM{}
	return &this
}

// GetInputFile returns the InputFile field value if set, zero value otherwise.
func (o *CreateTransportTaskBaseVM) GetInputFile() InputFileVM {
	if o == nil || IsNil(o.InputFile) {
		var ret InputFileVM
		return ret
	}
	return *o.InputFile
}

// GetInputFileOk returns a tuple with the InputFile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateTransportTaskBaseVM) GetInputFileOk() (*InputFileVM, bool) {
	if o == nil || IsNil(o.InputFile) {
		return nil, false
	}
	return o.InputFile, true
}

// HasInputFile returns a boolean if a field has been set.
func (o *CreateTransportTaskBaseVM) HasInputFile() bool {
	if o != nil && !IsNil(o.InputFile) {
		return true
	}

	return false
}

// SetInputFile gets a reference to the given InputFileVM and assigns it to the InputFile field.
func (o *CreateTransportTaskBaseVM) SetInputFile(v InputFileVM) {
	o.InputFile = &v
}

// GetT returns the T field value
func (o *CreateTransportTaskBaseVM) GetT() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.T
}

// GetTOk returns a tuple with the T field value
// and a boolean to check if the value has been set.
func (o *CreateTransportTaskBaseVM) GetTOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.T, true
}

// SetT sets field value
func (o *CreateTransportTaskBaseVM) SetT(v string) {
	o.T = v
}

func (o CreateTransportTaskBaseVM) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateTransportTaskBaseVM) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedCreateTaskBaseVM, errCreateTaskBaseVM := json.Marshal(o.CreateTaskBaseVM)
	if errCreateTaskBaseVM != nil {
		return map[string]interface{}{}, errCreateTaskBaseVM
	}
	errCreateTaskBaseVM = json.Unmarshal([]byte(serializedCreateTaskBaseVM), &toSerialize)
	if errCreateTaskBaseVM != nil {
		return map[string]interface{}{}, errCreateTaskBaseVM
	}
	if !IsNil(o.InputFile) {
		toSerialize["inputFile"] = o.InputFile
	}
	toSerialize["$t"] = o.T
	return toSerialize, nil
}

func (o *CreateTransportTaskBaseVM) UnmarshalJSON(data []byte) (err error) {
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

	varCreateTransportTaskBaseVM := _CreateTransportTaskBaseVM{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateTransportTaskBaseVM)

	if err != nil {
		return err
	}

	*o = CreateTransportTaskBaseVM(varCreateTransportTaskBaseVM)

	return err
}

type NullableCreateTransportTaskBaseVM struct {
	value *CreateTransportTaskBaseVM
	isSet bool
}

func (v NullableCreateTransportTaskBaseVM) Get() *CreateTransportTaskBaseVM {
	return v.value
}

func (v *NullableCreateTransportTaskBaseVM) Set(val *CreateTransportTaskBaseVM) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateTransportTaskBaseVM) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateTransportTaskBaseVM) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateTransportTaskBaseVM(val *CreateTransportTaskBaseVM) *NullableCreateTransportTaskBaseVM {
	return &NullableCreateTransportTaskBaseVM{value: val, isSet: true}
}

func (v NullableCreateTransportTaskBaseVM) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateTransportTaskBaseVM) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


