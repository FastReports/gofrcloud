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

// checks if the TransformTaskBaseVM type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TransformTaskBaseVM{}

// TransformTaskBaseVM struct for TransformTaskBaseVM
type TransformTaskBaseVM struct {
	TaskBaseVM
	InputFile *InputFileVM `json:"inputFile,omitempty"`
	Locale NullableString `json:"locale,omitempty"`
	OutputFile *OutputFileVM `json:"outputFile,omitempty"`
	TransportIds []string `json:"transportIds,omitempty"`
	T string `json:"$t"`
}

type _TransformTaskBaseVM TransformTaskBaseVM

// NewTransformTaskBaseVM instantiates a new TransformTaskBaseVM object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransformTaskBaseVM(t string) *TransformTaskBaseVM {
	this := TransformTaskBaseVM{}
	this.T = t
	return &this
}

// NewTransformTaskBaseVMWithDefaults instantiates a new TransformTaskBaseVM object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransformTaskBaseVMWithDefaults() *TransformTaskBaseVM {
	this := TransformTaskBaseVM{}
	return &this
}

// GetInputFile returns the InputFile field value if set, zero value otherwise.
func (o *TransformTaskBaseVM) GetInputFile() InputFileVM {
	if o == nil || IsNil(o.InputFile) {
		var ret InputFileVM
		return ret
	}
	return *o.InputFile
}

// GetInputFileOk returns a tuple with the InputFile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransformTaskBaseVM) GetInputFileOk() (*InputFileVM, bool) {
	if o == nil || IsNil(o.InputFile) {
		return nil, false
	}
	return o.InputFile, true
}

// HasInputFile returns a boolean if a field has been set.
func (o *TransformTaskBaseVM) HasInputFile() bool {
	if o != nil && !IsNil(o.InputFile) {
		return true
	}

	return false
}

// SetInputFile gets a reference to the given InputFileVM and assigns it to the InputFile field.
func (o *TransformTaskBaseVM) SetInputFile(v InputFileVM) {
	o.InputFile = &v
}

// GetLocale returns the Locale field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TransformTaskBaseVM) GetLocale() string {
	if o == nil || IsNil(o.Locale.Get()) {
		var ret string
		return ret
	}
	return *o.Locale.Get()
}

// GetLocaleOk returns a tuple with the Locale field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TransformTaskBaseVM) GetLocaleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Locale.Get(), o.Locale.IsSet()
}

// HasLocale returns a boolean if a field has been set.
func (o *TransformTaskBaseVM) HasLocale() bool {
	if o != nil && o.Locale.IsSet() {
		return true
	}

	return false
}

// SetLocale gets a reference to the given NullableString and assigns it to the Locale field.
func (o *TransformTaskBaseVM) SetLocale(v string) {
	o.Locale.Set(&v)
}
// SetLocaleNil sets the value for Locale to be an explicit nil
func (o *TransformTaskBaseVM) SetLocaleNil() {
	o.Locale.Set(nil)
}

// UnsetLocale ensures that no value is present for Locale, not even an explicit nil
func (o *TransformTaskBaseVM) UnsetLocale() {
	o.Locale.Unset()
}

// GetOutputFile returns the OutputFile field value if set, zero value otherwise.
func (o *TransformTaskBaseVM) GetOutputFile() OutputFileVM {
	if o == nil || IsNil(o.OutputFile) {
		var ret OutputFileVM
		return ret
	}
	return *o.OutputFile
}

// GetOutputFileOk returns a tuple with the OutputFile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransformTaskBaseVM) GetOutputFileOk() (*OutputFileVM, bool) {
	if o == nil || IsNil(o.OutputFile) {
		return nil, false
	}
	return o.OutputFile, true
}

// HasOutputFile returns a boolean if a field has been set.
func (o *TransformTaskBaseVM) HasOutputFile() bool {
	if o != nil && !IsNil(o.OutputFile) {
		return true
	}

	return false
}

// SetOutputFile gets a reference to the given OutputFileVM and assigns it to the OutputFile field.
func (o *TransformTaskBaseVM) SetOutputFile(v OutputFileVM) {
	o.OutputFile = &v
}

// GetTransportIds returns the TransportIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TransformTaskBaseVM) GetTransportIds() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.TransportIds
}

// GetTransportIdsOk returns a tuple with the TransportIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TransformTaskBaseVM) GetTransportIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.TransportIds) {
		return nil, false
	}
	return o.TransportIds, true
}

// HasTransportIds returns a boolean if a field has been set.
func (o *TransformTaskBaseVM) HasTransportIds() bool {
	if o != nil && IsNil(o.TransportIds) {
		return true
	}

	return false
}

// SetTransportIds gets a reference to the given []string and assigns it to the TransportIds field.
func (o *TransformTaskBaseVM) SetTransportIds(v []string) {
	o.TransportIds = v
}

// GetT returns the T field value
func (o *TransformTaskBaseVM) GetT() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.T
}

// GetTOk returns a tuple with the T field value
// and a boolean to check if the value has been set.
func (o *TransformTaskBaseVM) GetTOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.T, true
}

// SetT sets field value
func (o *TransformTaskBaseVM) SetT(v string) {
	o.T = v
}

func (o TransformTaskBaseVM) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TransformTaskBaseVM) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedTaskBaseVM, errTaskBaseVM := json.Marshal(o.TaskBaseVM)
	if errTaskBaseVM != nil {
		return map[string]interface{}{}, errTaskBaseVM
	}
	errTaskBaseVM = json.Unmarshal([]byte(serializedTaskBaseVM), &toSerialize)
	if errTaskBaseVM != nil {
		return map[string]interface{}{}, errTaskBaseVM
	}
	if !IsNil(o.InputFile) {
		toSerialize["inputFile"] = o.InputFile
	}
	if o.Locale.IsSet() {
		toSerialize["locale"] = o.Locale.Get()
	}
	if !IsNil(o.OutputFile) {
		toSerialize["outputFile"] = o.OutputFile
	}
	if o.TransportIds != nil {
		toSerialize["transportIds"] = o.TransportIds
	}
	toSerialize["$t"] = o.T
	return toSerialize, nil
}

func (o *TransformTaskBaseVM) UnmarshalJSON(data []byte) (err error) {
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

	varTransformTaskBaseVM := _TransformTaskBaseVM{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTransformTaskBaseVM)

	if err != nil {
		return err
	}

	*o = TransformTaskBaseVM(varTransformTaskBaseVM)

	return err
}

type NullableTransformTaskBaseVM struct {
	value *TransformTaskBaseVM
	isSet bool
}

func (v NullableTransformTaskBaseVM) Get() *TransformTaskBaseVM {
	return v.value
}

func (v *NullableTransformTaskBaseVM) Set(val *TransformTaskBaseVM) {
	v.value = val
	v.isSet = true
}

func (v NullableTransformTaskBaseVM) IsSet() bool {
	return v.isSet
}

func (v *NullableTransformTaskBaseVM) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransformTaskBaseVM(val *TransformTaskBaseVM) *NullableTransformTaskBaseVM {
	return &NullableTransformTaskBaseVM{value: val, isSet: true}
}

func (v NullableTransformTaskBaseVM) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransformTaskBaseVM) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


