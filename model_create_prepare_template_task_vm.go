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

// CreatePrepareTemplateTaskVM struct for CreatePrepareTemplateTaskVM
type CreatePrepareTemplateTaskVM struct {
	Exports []CreateExportReportTaskVM `json:"exports,omitempty"`
	PagesCount NullableInt32 `json:"pagesCount,omitempty"`
	ReportParameters map[string]string `json:"reportParameters,omitempty"`
	Name NullableString `json:"name,omitempty"`
	SubscriptionId NullableString `json:"subscriptionId,omitempty"`
	Type *TaskType `json:"type,omitempty"`
}

// NewCreatePrepareTemplateTaskVM instantiates a new CreatePrepareTemplateTaskVM object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreatePrepareTemplateTaskVM() *CreatePrepareTemplateTaskVM {
	this := CreatePrepareTemplateTaskVM{}
	return &this
}

// NewCreatePrepareTemplateTaskVMWithDefaults instantiates a new CreatePrepareTemplateTaskVM object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreatePrepareTemplateTaskVMWithDefaults() *CreatePrepareTemplateTaskVM {
	this := CreatePrepareTemplateTaskVM{}
	return &this
}

// GetExports returns the Exports field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreatePrepareTemplateTaskVM) GetExports() []CreateExportReportTaskVM {
	if o == nil  {
		var ret []CreateExportReportTaskVM
		return ret
	}
	return o.Exports
}

// GetExportsOk returns a tuple with the Exports field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreatePrepareTemplateTaskVM) GetExportsOk() (*[]CreateExportReportTaskVM, bool) {
	if o == nil || o.Exports == nil {
		return nil, false
	}
	return &o.Exports, true
}

// HasExports returns a boolean if a field has been set.
func (o *CreatePrepareTemplateTaskVM) HasExports() bool {
	if o != nil && o.Exports != nil {
		return true
	}

	return false
}

// SetExports gets a reference to the given []CreateExportReportTaskVM and assigns it to the Exports field.
func (o *CreatePrepareTemplateTaskVM) SetExports(v []CreateExportReportTaskVM) {
	o.Exports = v
}

// GetPagesCount returns the PagesCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreatePrepareTemplateTaskVM) GetPagesCount() int32 {
	if o == nil || o.PagesCount.Get() == nil {
		var ret int32
		return ret
	}
	return *o.PagesCount.Get()
}

// GetPagesCountOk returns a tuple with the PagesCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreatePrepareTemplateTaskVM) GetPagesCountOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PagesCount.Get(), o.PagesCount.IsSet()
}

// HasPagesCount returns a boolean if a field has been set.
func (o *CreatePrepareTemplateTaskVM) HasPagesCount() bool {
	if o != nil && o.PagesCount.IsSet() {
		return true
	}

	return false
}

// SetPagesCount gets a reference to the given NullableInt32 and assigns it to the PagesCount field.
func (o *CreatePrepareTemplateTaskVM) SetPagesCount(v int32) {
	o.PagesCount.Set(&v)
}
// SetPagesCountNil sets the value for PagesCount to be an explicit nil
func (o *CreatePrepareTemplateTaskVM) SetPagesCountNil() {
	o.PagesCount.Set(nil)
}

// UnsetPagesCount ensures that no value is present for PagesCount, not even an explicit nil
func (o *CreatePrepareTemplateTaskVM) UnsetPagesCount() {
	o.PagesCount.Unset()
}

// GetReportParameters returns the ReportParameters field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreatePrepareTemplateTaskVM) GetReportParameters() map[string]string {
	if o == nil  {
		var ret map[string]string
		return ret
	}
	return o.ReportParameters
}

// GetReportParametersOk returns a tuple with the ReportParameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreatePrepareTemplateTaskVM) GetReportParametersOk() (*map[string]string, bool) {
	if o == nil || o.ReportParameters == nil {
		return nil, false
	}
	return &o.ReportParameters, true
}

// HasReportParameters returns a boolean if a field has been set.
func (o *CreatePrepareTemplateTaskVM) HasReportParameters() bool {
	if o != nil && o.ReportParameters != nil {
		return true
	}

	return false
}

// SetReportParameters gets a reference to the given map[string]string and assigns it to the ReportParameters field.
func (o *CreatePrepareTemplateTaskVM) SetReportParameters(v map[string]string) {
	o.ReportParameters = v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreatePrepareTemplateTaskVM) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreatePrepareTemplateTaskVM) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *CreatePrepareTemplateTaskVM) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *CreatePrepareTemplateTaskVM) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *CreatePrepareTemplateTaskVM) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *CreatePrepareTemplateTaskVM) UnsetName() {
	o.Name.Unset()
}

// GetSubscriptionId returns the SubscriptionId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreatePrepareTemplateTaskVM) GetSubscriptionId() string {
	if o == nil || o.SubscriptionId.Get() == nil {
		var ret string
		return ret
	}
	return *o.SubscriptionId.Get()
}

// GetSubscriptionIdOk returns a tuple with the SubscriptionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreatePrepareTemplateTaskVM) GetSubscriptionIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SubscriptionId.Get(), o.SubscriptionId.IsSet()
}

// HasSubscriptionId returns a boolean if a field has been set.
func (o *CreatePrepareTemplateTaskVM) HasSubscriptionId() bool {
	if o != nil && o.SubscriptionId.IsSet() {
		return true
	}

	return false
}

// SetSubscriptionId gets a reference to the given NullableString and assigns it to the SubscriptionId field.
func (o *CreatePrepareTemplateTaskVM) SetSubscriptionId(v string) {
	o.SubscriptionId.Set(&v)
}
// SetSubscriptionIdNil sets the value for SubscriptionId to be an explicit nil
func (o *CreatePrepareTemplateTaskVM) SetSubscriptionIdNil() {
	o.SubscriptionId.Set(nil)
}

// UnsetSubscriptionId ensures that no value is present for SubscriptionId, not even an explicit nil
func (o *CreatePrepareTemplateTaskVM) UnsetSubscriptionId() {
	o.SubscriptionId.Unset()
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *CreatePrepareTemplateTaskVM) GetType() TaskType {
	if o == nil || o.Type == nil {
		var ret TaskType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatePrepareTemplateTaskVM) GetTypeOk() (*TaskType, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *CreatePrepareTemplateTaskVM) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given TaskType and assigns it to the Type field.
func (o *CreatePrepareTemplateTaskVM) SetType(v TaskType) {
	o.Type = &v
}

func (o CreatePrepareTemplateTaskVM) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Exports != nil {
		toSerialize["exports"] = o.Exports
	}
	if o.PagesCount.IsSet() {
		toSerialize["pagesCount"] = o.PagesCount.Get()
	}
	if o.ReportParameters != nil {
		toSerialize["reportParameters"] = o.ReportParameters
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

type NullableCreatePrepareTemplateTaskVM struct {
	value *CreatePrepareTemplateTaskVM
	isSet bool
}

func (v NullableCreatePrepareTemplateTaskVM) Get() *CreatePrepareTemplateTaskVM {
	return v.value
}

func (v *NullableCreatePrepareTemplateTaskVM) Set(val *CreatePrepareTemplateTaskVM) {
	v.value = val
	v.isSet = true
}

func (v NullableCreatePrepareTemplateTaskVM) IsSet() bool {
	return v.isSet
}

func (v *NullableCreatePrepareTemplateTaskVM) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreatePrepareTemplateTaskVM(val *CreatePrepareTemplateTaskVM) *NullableCreatePrepareTemplateTaskVM {
	return &NullableCreatePrepareTemplateTaskVM{value: val, isSet: true}
}

func (v NullableCreatePrepareTemplateTaskVM) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreatePrepareTemplateTaskVM) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


