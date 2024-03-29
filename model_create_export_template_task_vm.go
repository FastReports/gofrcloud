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

// checks if the CreateExportTemplateTaskVM type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateExportTemplateTaskVM{}

// CreateExportTemplateTaskVM struct for CreateExportTemplateTaskVM
type CreateExportTemplateTaskVM struct {
	ReportParameters map[string]string `json:"reportParameters,omitempty"`
}

// NewCreateExportTemplateTaskVM instantiates a new CreateExportTemplateTaskVM object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateExportTemplateTaskVM(t string) *CreateExportTemplateTaskVM {
	this := CreateExportTemplateTaskVM{}
	this.T = t
	return &this
}

// NewCreateExportTemplateTaskVMWithDefaults instantiates a new CreateExportTemplateTaskVM object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateExportTemplateTaskVMWithDefaults() *CreateExportTemplateTaskVM {
	this := CreateExportTemplateTaskVM{}
	return &this
}

// GetReportParameters returns the ReportParameters field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateExportTemplateTaskVM) GetReportParameters() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}
	return o.ReportParameters
}

// GetReportParametersOk returns a tuple with the ReportParameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateExportTemplateTaskVM) GetReportParametersOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.ReportParameters) {
		return nil, false
	}
	return &o.ReportParameters, true
}

// HasReportParameters returns a boolean if a field has been set.
func (o *CreateExportTemplateTaskVM) HasReportParameters() bool {
	if o != nil && IsNil(o.ReportParameters) {
		return true
	}

	return false
}

// SetReportParameters gets a reference to the given map[string]string and assigns it to the ReportParameters field.
func (o *CreateExportTemplateTaskVM) SetReportParameters(v map[string]string) {
	o.ReportParameters = v
}

func (o CreateExportTemplateTaskVM) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateExportTemplateTaskVM) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.ReportParameters != nil {
		toSerialize["reportParameters"] = o.ReportParameters
	}
	return toSerialize, nil
}

type NullableCreateExportTemplateTaskVM struct {
	value *CreateExportTemplateTaskVM
	isSet bool
}

func (v NullableCreateExportTemplateTaskVM) Get() *CreateExportTemplateTaskVM {
	return v.value
}

func (v *NullableCreateExportTemplateTaskVM) Set(val *CreateExportTemplateTaskVM) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateExportTemplateTaskVM) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateExportTemplateTaskVM) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateExportTemplateTaskVM(val *CreateExportTemplateTaskVM) *NullableCreateExportTemplateTaskVM {
	return &NullableCreateExportTemplateTaskVM{value: val, isSet: true}
}

func (v NullableCreateExportTemplateTaskVM) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateExportTemplateTaskVM) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


