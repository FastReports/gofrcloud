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

// CreateTransformTaskBaseVM struct for CreateTransformTaskBaseVM
type CreateTransformTaskBaseVM struct {
	Locale NullableString `json:"locale,omitempty"`
	InputFile *InputFileVM `json:"inputFile,omitempty"`
	OutputFile *OutputFileVM `json:"outputFile,omitempty"`
	Transports []CreateTransportTaskBaseVM `json:"transports,omitempty"`
	Name NullableString `json:"name,omitempty"`
	SubscriptionId NullableString `json:"subscriptionId,omitempty"`
	Type *TaskType `json:"type,omitempty"`
}

// NewCreateTransformTaskBaseVM instantiates a new CreateTransformTaskBaseVM object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateTransformTaskBaseVM() *CreateTransformTaskBaseVM {
	this := CreateTransformTaskBaseVM{}
	return &this
}

// NewCreateTransformTaskBaseVMWithDefaults instantiates a new CreateTransformTaskBaseVM object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateTransformTaskBaseVMWithDefaults() *CreateTransformTaskBaseVM {
	this := CreateTransformTaskBaseVM{}
	return &this
}

// GetLocale returns the Locale field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateTransformTaskBaseVM) GetLocale() string {
	if o == nil || o.Locale.Get() == nil {
		var ret string
		return ret
	}
	return *o.Locale.Get()
}

// GetLocaleOk returns a tuple with the Locale field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateTransformTaskBaseVM) GetLocaleOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Locale.Get(), o.Locale.IsSet()
}

// HasLocale returns a boolean if a field has been set.
func (o *CreateTransformTaskBaseVM) HasLocale() bool {
	if o != nil && o.Locale.IsSet() {
		return true
	}

	return false
}

// SetLocale gets a reference to the given NullableString and assigns it to the Locale field.
func (o *CreateTransformTaskBaseVM) SetLocale(v string) {
	o.Locale.Set(&v)
}
// SetLocaleNil sets the value for Locale to be an explicit nil
func (o *CreateTransformTaskBaseVM) SetLocaleNil() {
	o.Locale.Set(nil)
}

// UnsetLocale ensures that no value is present for Locale, not even an explicit nil
func (o *CreateTransformTaskBaseVM) UnsetLocale() {
	o.Locale.Unset()
}

// GetInputFile returns the InputFile field value if set, zero value otherwise.
func (o *CreateTransformTaskBaseVM) GetInputFile() InputFileVM {
	if o == nil || o.InputFile == nil {
		var ret InputFileVM
		return ret
	}
	return *o.InputFile
}

// GetInputFileOk returns a tuple with the InputFile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateTransformTaskBaseVM) GetInputFileOk() (*InputFileVM, bool) {
	if o == nil || o.InputFile == nil {
		return nil, false
	}
	return o.InputFile, true
}

// HasInputFile returns a boolean if a field has been set.
func (o *CreateTransformTaskBaseVM) HasInputFile() bool {
	if o != nil && o.InputFile != nil {
		return true
	}

	return false
}

// SetInputFile gets a reference to the given InputFileVM and assigns it to the InputFile field.
func (o *CreateTransformTaskBaseVM) SetInputFile(v InputFileVM) {
	o.InputFile = &v
}

// GetOutputFile returns the OutputFile field value if set, zero value otherwise.
func (o *CreateTransformTaskBaseVM) GetOutputFile() OutputFileVM {
	if o == nil || o.OutputFile == nil {
		var ret OutputFileVM
		return ret
	}
	return *o.OutputFile
}

// GetOutputFileOk returns a tuple with the OutputFile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateTransformTaskBaseVM) GetOutputFileOk() (*OutputFileVM, bool) {
	if o == nil || o.OutputFile == nil {
		return nil, false
	}
	return o.OutputFile, true
}

// HasOutputFile returns a boolean if a field has been set.
func (o *CreateTransformTaskBaseVM) HasOutputFile() bool {
	if o != nil && o.OutputFile != nil {
		return true
	}

	return false
}

// SetOutputFile gets a reference to the given OutputFileVM and assigns it to the OutputFile field.
func (o *CreateTransformTaskBaseVM) SetOutputFile(v OutputFileVM) {
	o.OutputFile = &v
}

// GetTransports returns the Transports field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateTransformTaskBaseVM) GetTransports() []CreateTransportTaskBaseVM {
	if o == nil  {
		var ret []CreateTransportTaskBaseVM
		return ret
	}
	return o.Transports
}

// GetTransportsOk returns a tuple with the Transports field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateTransformTaskBaseVM) GetTransportsOk() (*[]CreateTransportTaskBaseVM, bool) {
	if o == nil || o.Transports == nil {
		return nil, false
	}
	return &o.Transports, true
}

// HasTransports returns a boolean if a field has been set.
func (o *CreateTransformTaskBaseVM) HasTransports() bool {
	if o != nil && o.Transports != nil {
		return true
	}

	return false
}

// SetTransports gets a reference to the given []CreateTransportTaskBaseVM and assigns it to the Transports field.
func (o *CreateTransformTaskBaseVM) SetTransports(v []CreateTransportTaskBaseVM) {
	o.Transports = v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateTransformTaskBaseVM) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateTransformTaskBaseVM) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *CreateTransformTaskBaseVM) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *CreateTransformTaskBaseVM) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *CreateTransformTaskBaseVM) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *CreateTransformTaskBaseVM) UnsetName() {
	o.Name.Unset()
}

// GetSubscriptionId returns the SubscriptionId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateTransformTaskBaseVM) GetSubscriptionId() string {
	if o == nil || o.SubscriptionId.Get() == nil {
		var ret string
		return ret
	}
	return *o.SubscriptionId.Get()
}

// GetSubscriptionIdOk returns a tuple with the SubscriptionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateTransformTaskBaseVM) GetSubscriptionIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SubscriptionId.Get(), o.SubscriptionId.IsSet()
}

// HasSubscriptionId returns a boolean if a field has been set.
func (o *CreateTransformTaskBaseVM) HasSubscriptionId() bool {
	if o != nil && o.SubscriptionId.IsSet() {
		return true
	}

	return false
}

// SetSubscriptionId gets a reference to the given NullableString and assigns it to the SubscriptionId field.
func (o *CreateTransformTaskBaseVM) SetSubscriptionId(v string) {
	o.SubscriptionId.Set(&v)
}
// SetSubscriptionIdNil sets the value for SubscriptionId to be an explicit nil
func (o *CreateTransformTaskBaseVM) SetSubscriptionIdNil() {
	o.SubscriptionId.Set(nil)
}

// UnsetSubscriptionId ensures that no value is present for SubscriptionId, not even an explicit nil
func (o *CreateTransformTaskBaseVM) UnsetSubscriptionId() {
	o.SubscriptionId.Unset()
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *CreateTransformTaskBaseVM) GetType() TaskType {
	if o == nil || o.Type == nil {
		var ret TaskType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateTransformTaskBaseVM) GetTypeOk() (*TaskType, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *CreateTransformTaskBaseVM) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given TaskType and assigns it to the Type field.
func (o *CreateTransformTaskBaseVM) SetType(v TaskType) {
	o.Type = &v
}

func (o CreateTransformTaskBaseVM) MarshalJSON() ([]byte, error) {
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

type NullableCreateTransformTaskBaseVM struct {
	value *CreateTransformTaskBaseVM
	isSet bool
}

func (v NullableCreateTransformTaskBaseVM) Get() *CreateTransformTaskBaseVM {
	return v.value
}

func (v *NullableCreateTransformTaskBaseVM) Set(val *CreateTransformTaskBaseVM) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateTransformTaskBaseVM) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateTransformTaskBaseVM) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateTransformTaskBaseVM(val *CreateTransformTaskBaseVM) *NullableCreateTransformTaskBaseVM {
	return &NullableCreateTransformTaskBaseVM{value: val, isSet: true}
}

func (v NullableCreateTransformTaskBaseVM) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateTransformTaskBaseVM) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

