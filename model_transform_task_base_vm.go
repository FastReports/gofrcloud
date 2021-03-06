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

// TransformTaskBaseVM struct for TransformTaskBaseVM
type TransformTaskBaseVM struct {
	Locale NullableString `json:"locale,omitempty"`
	InputFile *InputFileVM `json:"inputFile,omitempty"`
	OutputFile *OutputFileVM `json:"outputFile,omitempty"`
	Transports []TransportTaskBaseVM `json:"transports,omitempty"`
	Name NullableString `json:"name,omitempty"`
	SubscriptionId NullableString `json:"subscriptionId,omitempty"`
	Type *TaskType `json:"type,omitempty"`
}

// NewTransformTaskBaseVM instantiates a new TransformTaskBaseVM object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransformTaskBaseVM() *TransformTaskBaseVM {
	this := TransformTaskBaseVM{}
	return &this
}

// NewTransformTaskBaseVMWithDefaults instantiates a new TransformTaskBaseVM object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransformTaskBaseVMWithDefaults() *TransformTaskBaseVM {
	this := TransformTaskBaseVM{}
	return &this
}

// GetLocale returns the Locale field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TransformTaskBaseVM) GetLocale() string {
	if o == nil || o.Locale.Get() == nil {
		var ret string
		return ret
	}
	return *o.Locale.Get()
}

// GetLocaleOk returns a tuple with the Locale field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TransformTaskBaseVM) GetLocaleOk() (*string, bool) {
	if o == nil  {
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

// GetInputFile returns the InputFile field value if set, zero value otherwise.
func (o *TransformTaskBaseVM) GetInputFile() InputFileVM {
	if o == nil || o.InputFile == nil {
		var ret InputFileVM
		return ret
	}
	return *o.InputFile
}

// GetInputFileOk returns a tuple with the InputFile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransformTaskBaseVM) GetInputFileOk() (*InputFileVM, bool) {
	if o == nil || o.InputFile == nil {
		return nil, false
	}
	return o.InputFile, true
}

// HasInputFile returns a boolean if a field has been set.
func (o *TransformTaskBaseVM) HasInputFile() bool {
	if o != nil && o.InputFile != nil {
		return true
	}

	return false
}

// SetInputFile gets a reference to the given InputFileVM and assigns it to the InputFile field.
func (o *TransformTaskBaseVM) SetInputFile(v InputFileVM) {
	o.InputFile = &v
}

// GetOutputFile returns the OutputFile field value if set, zero value otherwise.
func (o *TransformTaskBaseVM) GetOutputFile() OutputFileVM {
	if o == nil || o.OutputFile == nil {
		var ret OutputFileVM
		return ret
	}
	return *o.OutputFile
}

// GetOutputFileOk returns a tuple with the OutputFile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransformTaskBaseVM) GetOutputFileOk() (*OutputFileVM, bool) {
	if o == nil || o.OutputFile == nil {
		return nil, false
	}
	return o.OutputFile, true
}

// HasOutputFile returns a boolean if a field has been set.
func (o *TransformTaskBaseVM) HasOutputFile() bool {
	if o != nil && o.OutputFile != nil {
		return true
	}

	return false
}

// SetOutputFile gets a reference to the given OutputFileVM and assigns it to the OutputFile field.
func (o *TransformTaskBaseVM) SetOutputFile(v OutputFileVM) {
	o.OutputFile = &v
}

// GetTransports returns the Transports field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TransformTaskBaseVM) GetTransports() []TransportTaskBaseVM {
	if o == nil  {
		var ret []TransportTaskBaseVM
		return ret
	}
	return o.Transports
}

// GetTransportsOk returns a tuple with the Transports field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TransformTaskBaseVM) GetTransportsOk() (*[]TransportTaskBaseVM, bool) {
	if o == nil || o.Transports == nil {
		return nil, false
	}
	return &o.Transports, true
}

// HasTransports returns a boolean if a field has been set.
func (o *TransformTaskBaseVM) HasTransports() bool {
	if o != nil && o.Transports != nil {
		return true
	}

	return false
}

// SetTransports gets a reference to the given []TransportTaskBaseVM and assigns it to the Transports field.
func (o *TransformTaskBaseVM) SetTransports(v []TransportTaskBaseVM) {
	o.Transports = v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TransformTaskBaseVM) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TransformTaskBaseVM) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *TransformTaskBaseVM) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *TransformTaskBaseVM) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *TransformTaskBaseVM) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *TransformTaskBaseVM) UnsetName() {
	o.Name.Unset()
}

// GetSubscriptionId returns the SubscriptionId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TransformTaskBaseVM) GetSubscriptionId() string {
	if o == nil || o.SubscriptionId.Get() == nil {
		var ret string
		return ret
	}
	return *o.SubscriptionId.Get()
}

// GetSubscriptionIdOk returns a tuple with the SubscriptionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TransformTaskBaseVM) GetSubscriptionIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SubscriptionId.Get(), o.SubscriptionId.IsSet()
}

// HasSubscriptionId returns a boolean if a field has been set.
func (o *TransformTaskBaseVM) HasSubscriptionId() bool {
	if o != nil && o.SubscriptionId.IsSet() {
		return true
	}

	return false
}

// SetSubscriptionId gets a reference to the given NullableString and assigns it to the SubscriptionId field.
func (o *TransformTaskBaseVM) SetSubscriptionId(v string) {
	o.SubscriptionId.Set(&v)
}
// SetSubscriptionIdNil sets the value for SubscriptionId to be an explicit nil
func (o *TransformTaskBaseVM) SetSubscriptionIdNil() {
	o.SubscriptionId.Set(nil)
}

// UnsetSubscriptionId ensures that no value is present for SubscriptionId, not even an explicit nil
func (o *TransformTaskBaseVM) UnsetSubscriptionId() {
	o.SubscriptionId.Unset()
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *TransformTaskBaseVM) GetType() TaskType {
	if o == nil || o.Type == nil {
		var ret TaskType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransformTaskBaseVM) GetTypeOk() (*TaskType, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *TransformTaskBaseVM) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given TaskType and assigns it to the Type field.
func (o *TransformTaskBaseVM) SetType(v TaskType) {
	o.Type = &v
}

func (o TransformTaskBaseVM) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Locale.IsSet() {
		toSerialize["locale"] = o.Locale.Get()
	}
	if o.InputFile != nil {
		toSerialize["inputFile"] = o.InputFile
	}
	if o.OutputFile != nil {
		toSerialize["outputFile"] = o.OutputFile
	}
	if o.Transports != nil {
		toSerialize["transports"] = o.Transports
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.SubscriptionId.IsSet() {
		toSerialize["subscriptionId"] = o.SubscriptionId.Get()
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
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


