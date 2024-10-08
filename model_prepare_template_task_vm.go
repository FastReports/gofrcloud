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

// checks if the PrepareTemplateTaskVM type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PrepareTemplateTaskVM{}

// PrepareTemplateTaskVM struct for PrepareTemplateTaskVM
type PrepareTemplateTaskVM struct {
	TransformTaskBaseVM
	ExportIds []string `json:"exportIds,omitempty"`
	PagesCount *int32 `json:"pagesCount,omitempty"`
	ReportParameters map[string]string `json:"reportParameters,omitempty"`
	T string `json:"$t"`
}

type _PrepareTemplateTaskVM PrepareTemplateTaskVM

// NewPrepareTemplateTaskVM instantiates a new PrepareTemplateTaskVM object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPrepareTemplateTaskVM(t string) *PrepareTemplateTaskVM {
	this := PrepareTemplateTaskVM{}
	this.T = t
	return &this
}

// NewPrepareTemplateTaskVMWithDefaults instantiates a new PrepareTemplateTaskVM object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPrepareTemplateTaskVMWithDefaults() *PrepareTemplateTaskVM {
	this := PrepareTemplateTaskVM{}
	return &this
}

// GetExportIds returns the ExportIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PrepareTemplateTaskVM) GetExportIds() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.ExportIds
}

// GetExportIdsOk returns a tuple with the ExportIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PrepareTemplateTaskVM) GetExportIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.ExportIds) {
		return nil, false
	}
	return o.ExportIds, true
}

// HasExportIds returns a boolean if a field has been set.
func (o *PrepareTemplateTaskVM) HasExportIds() bool {
	if o != nil && IsNil(o.ExportIds) {
		return true
	}

	return false
}

// SetExportIds gets a reference to the given []string and assigns it to the ExportIds field.
func (o *PrepareTemplateTaskVM) SetExportIds(v []string) {
	o.ExportIds = v
}

// GetPagesCount returns the PagesCount field value if set, zero value otherwise.
func (o *PrepareTemplateTaskVM) GetPagesCount() int32 {
	if o == nil || IsNil(o.PagesCount) {
		var ret int32
		return ret
	}
	return *o.PagesCount
}

// GetPagesCountOk returns a tuple with the PagesCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrepareTemplateTaskVM) GetPagesCountOk() (*int32, bool) {
	if o == nil || IsNil(o.PagesCount) {
		return nil, false
	}
	return o.PagesCount, true
}

// HasPagesCount returns a boolean if a field has been set.
func (o *PrepareTemplateTaskVM) HasPagesCount() bool {
	if o != nil && !IsNil(o.PagesCount) {
		return true
	}

	return false
}

// SetPagesCount gets a reference to the given int32 and assigns it to the PagesCount field.
func (o *PrepareTemplateTaskVM) SetPagesCount(v int32) {
	o.PagesCount = &v
}

// GetReportParameters returns the ReportParameters field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PrepareTemplateTaskVM) GetReportParameters() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}
	return o.ReportParameters
}

// GetReportParametersOk returns a tuple with the ReportParameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PrepareTemplateTaskVM) GetReportParametersOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.ReportParameters) {
		return nil, false
	}
	return &o.ReportParameters, true
}

// HasReportParameters returns a boolean if a field has been set.
func (o *PrepareTemplateTaskVM) HasReportParameters() bool {
	if o != nil && IsNil(o.ReportParameters) {
		return true
	}

	return false
}

// SetReportParameters gets a reference to the given map[string]string and assigns it to the ReportParameters field.
func (o *PrepareTemplateTaskVM) SetReportParameters(v map[string]string) {
	o.ReportParameters = v
}

// GetT returns the T field value
func (o *PrepareTemplateTaskVM) GetT() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.T
}

// GetTOk returns a tuple with the T field value
// and a boolean to check if the value has been set.
func (o *PrepareTemplateTaskVM) GetTOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.T, true
}

// SetT sets field value
func (o *PrepareTemplateTaskVM) SetT(v string) {
	o.T = v
}

func (o PrepareTemplateTaskVM) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PrepareTemplateTaskVM) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedTransformTaskBaseVM, errTransformTaskBaseVM := json.Marshal(o.TransformTaskBaseVM)
	if errTransformTaskBaseVM != nil {
		return map[string]interface{}{}, errTransformTaskBaseVM
	}
	errTransformTaskBaseVM = json.Unmarshal([]byte(serializedTransformTaskBaseVM), &toSerialize)
	if errTransformTaskBaseVM != nil {
		return map[string]interface{}{}, errTransformTaskBaseVM
	}
	if o.ExportIds != nil {
		toSerialize["exportIds"] = o.ExportIds
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

func (o *PrepareTemplateTaskVM) UnmarshalJSON(data []byte) (err error) {
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

	varPrepareTemplateTaskVM := _PrepareTemplateTaskVM{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPrepareTemplateTaskVM)

	if err != nil {
		return err
	}

	*o = PrepareTemplateTaskVM(varPrepareTemplateTaskVM)

	return err
}

type NullablePrepareTemplateTaskVM struct {
	value *PrepareTemplateTaskVM
	isSet bool
}

func (v NullablePrepareTemplateTaskVM) Get() *PrepareTemplateTaskVM {
	return v.value
}

func (v *NullablePrepareTemplateTaskVM) Set(val *PrepareTemplateTaskVM) {
	v.value = val
	v.isSet = true
}

func (v NullablePrepareTemplateTaskVM) IsSet() bool {
	return v.isSet
}

func (v *NullablePrepareTemplateTaskVM) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePrepareTemplateTaskVM(val *PrepareTemplateTaskVM) *NullablePrepareTemplateTaskVM {
	return &NullablePrepareTemplateTaskVM{value: val, isSet: true}
}

func (v NullablePrepareTemplateTaskVM) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePrepareTemplateTaskVM) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


