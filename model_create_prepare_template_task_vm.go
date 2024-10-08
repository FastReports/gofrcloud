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

// checks if the CreatePrepareTemplateTaskVM type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreatePrepareTemplateTaskVM{}

// CreatePrepareTemplateTaskVM struct for CreatePrepareTemplateTaskVM
type CreatePrepareTemplateTaskVM struct {
	CreateTransformTaskBaseVM
	Exports []CreateExportReportTaskVM `json:"exports,omitempty"`
	PagesCount *int32 `json:"pagesCount,omitempty"`
	ReportParameters map[string]string `json:"reportParameters,omitempty"`
	T string `json:"$t"`
}

type _CreatePrepareTemplateTaskVM CreatePrepareTemplateTaskVM

// NewCreatePrepareTemplateTaskVM instantiates a new CreatePrepareTemplateTaskVM object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreatePrepareTemplateTaskVM(t string) *CreatePrepareTemplateTaskVM {
	this := CreatePrepareTemplateTaskVM{}
	this.T = t
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
	if o == nil {
		var ret []CreateExportReportTaskVM
		return ret
	}
	return o.Exports
}

// GetExportsOk returns a tuple with the Exports field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreatePrepareTemplateTaskVM) GetExportsOk() ([]CreateExportReportTaskVM, bool) {
	if o == nil || IsNil(o.Exports) {
		return nil, false
	}
	return o.Exports, true
}

// HasExports returns a boolean if a field has been set.
func (o *CreatePrepareTemplateTaskVM) HasExports() bool {
	if o != nil && IsNil(o.Exports) {
		return true
	}

	return false
}

// SetExports gets a reference to the given []CreateExportReportTaskVM and assigns it to the Exports field.
func (o *CreatePrepareTemplateTaskVM) SetExports(v []CreateExportReportTaskVM) {
	o.Exports = v
}

// GetPagesCount returns the PagesCount field value if set, zero value otherwise.
func (o *CreatePrepareTemplateTaskVM) GetPagesCount() int32 {
	if o == nil || IsNil(o.PagesCount) {
		var ret int32
		return ret
	}
	return *o.PagesCount
}

// GetPagesCountOk returns a tuple with the PagesCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatePrepareTemplateTaskVM) GetPagesCountOk() (*int32, bool) {
	if o == nil || IsNil(o.PagesCount) {
		return nil, false
	}
	return o.PagesCount, true
}

// HasPagesCount returns a boolean if a field has been set.
func (o *CreatePrepareTemplateTaskVM) HasPagesCount() bool {
	if o != nil && !IsNil(o.PagesCount) {
		return true
	}

	return false
}

// SetPagesCount gets a reference to the given int32 and assigns it to the PagesCount field.
func (o *CreatePrepareTemplateTaskVM) SetPagesCount(v int32) {
	o.PagesCount = &v
}

// GetReportParameters returns the ReportParameters field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreatePrepareTemplateTaskVM) GetReportParameters() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}
	return o.ReportParameters
}

// GetReportParametersOk returns a tuple with the ReportParameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreatePrepareTemplateTaskVM) GetReportParametersOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.ReportParameters) {
		return nil, false
	}
	return &o.ReportParameters, true
}

// HasReportParameters returns a boolean if a field has been set.
func (o *CreatePrepareTemplateTaskVM) HasReportParameters() bool {
	if o != nil && IsNil(o.ReportParameters) {
		return true
	}

	return false
}

// SetReportParameters gets a reference to the given map[string]string and assigns it to the ReportParameters field.
func (o *CreatePrepareTemplateTaskVM) SetReportParameters(v map[string]string) {
	o.ReportParameters = v
}

// GetT returns the T field value
func (o *CreatePrepareTemplateTaskVM) GetT() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.T
}

// GetTOk returns a tuple with the T field value
// and a boolean to check if the value has been set.
func (o *CreatePrepareTemplateTaskVM) GetTOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.T, true
}

// SetT sets field value
func (o *CreatePrepareTemplateTaskVM) SetT(v string) {
	o.T = v
}

func (o CreatePrepareTemplateTaskVM) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreatePrepareTemplateTaskVM) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedCreateTransformTaskBaseVM, errCreateTransformTaskBaseVM := json.Marshal(o.CreateTransformTaskBaseVM)
	if errCreateTransformTaskBaseVM != nil {
		return map[string]interface{}{}, errCreateTransformTaskBaseVM
	}
	errCreateTransformTaskBaseVM = json.Unmarshal([]byte(serializedCreateTransformTaskBaseVM), &toSerialize)
	if errCreateTransformTaskBaseVM != nil {
		return map[string]interface{}{}, errCreateTransformTaskBaseVM
	}
	if o.Exports != nil {
		toSerialize["exports"] = o.Exports
	}
	if !IsNil(o.PagesCount) {
		toSerialize["pagesCount"] = o.PagesCount
	}
	if o.ReportParameters != nil {
		toSerialize["reportParameters"] = o.ReportParameters
	}
	toSerialize["$t"] = o.T
	return toSerialize, nil
}

func (o *CreatePrepareTemplateTaskVM) UnmarshalJSON(data []byte) (err error) {
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

	varCreatePrepareTemplateTaskVM := _CreatePrepareTemplateTaskVM{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreatePrepareTemplateTaskVM)

	if err != nil {
		return err
	}

	*o = CreatePrepareTemplateTaskVM(varCreatePrepareTemplateTaskVM)

	return err
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


