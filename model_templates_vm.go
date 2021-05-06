/*
 * FastReport Cloud
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package FastReport.Cloud.SDK

import (
	"encoding/json"
)

// TemplatesVM struct for TemplatesVM
type TemplatesVM struct {
	Files *[]TemplateVM `json:"files,omitempty"`
	Count *int64 `json:"count,omitempty"`
	Skip *int32 `json:"skip,omitempty"`
	Take *int32 `json:"take,omitempty"`
}

// NewTemplatesVM instantiates a new TemplatesVM object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTemplatesVM() *TemplatesVM {
	this := TemplatesVM{}
	return &this
}

// NewTemplatesVMWithDefaults instantiates a new TemplatesVM object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTemplatesVMWithDefaults() *TemplatesVM {
	this := TemplatesVM{}
	return &this
}

// GetFiles returns the Files field value if set, zero value otherwise.
func (o *TemplatesVM) GetFiles() []TemplateVM {
	if o == nil || o.Files == nil {
		var ret []TemplateVM
		return ret
	}
	return *o.Files
}

// GetFilesOk returns a tuple with the Files field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplatesVM) GetFilesOk() (*[]TemplateVM, bool) {
	if o == nil || o.Files == nil {
		return nil, false
	}
	return o.Files, true
}

// HasFiles returns a boolean if a field has been set.
func (o *TemplatesVM) HasFiles() bool {
	if o != nil && o.Files != nil {
		return true
	}

	return false
}

// SetFiles gets a reference to the given []TemplateVM and assigns it to the Files field.
func (o *TemplatesVM) SetFiles(v []TemplateVM) {
	o.Files = &v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *TemplatesVM) GetCount() int64 {
	if o == nil || o.Count == nil {
		var ret int64
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplatesVM) GetCountOk() (*int64, bool) {
	if o == nil || o.Count == nil {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *TemplatesVM) HasCount() bool {
	if o != nil && o.Count != nil {
		return true
	}

	return false
}

// SetCount gets a reference to the given int64 and assigns it to the Count field.
func (o *TemplatesVM) SetCount(v int64) {
	o.Count = &v
}

// GetSkip returns the Skip field value if set, zero value otherwise.
func (o *TemplatesVM) GetSkip() int32 {
	if o == nil || o.Skip == nil {
		var ret int32
		return ret
	}
	return *o.Skip
}

// GetSkipOk returns a tuple with the Skip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplatesVM) GetSkipOk() (*int32, bool) {
	if o == nil || o.Skip == nil {
		return nil, false
	}
	return o.Skip, true
}

// HasSkip returns a boolean if a field has been set.
func (o *TemplatesVM) HasSkip() bool {
	if o != nil && o.Skip != nil {
		return true
	}

	return false
}

// SetSkip gets a reference to the given int32 and assigns it to the Skip field.
func (o *TemplatesVM) SetSkip(v int32) {
	o.Skip = &v
}

// GetTake returns the Take field value if set, zero value otherwise.
func (o *TemplatesVM) GetTake() int32 {
	if o == nil || o.Take == nil {
		var ret int32
		return ret
	}
	return *o.Take
}

// GetTakeOk returns a tuple with the Take field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplatesVM) GetTakeOk() (*int32, bool) {
	if o == nil || o.Take == nil {
		return nil, false
	}
	return o.Take, true
}

// HasTake returns a boolean if a field has been set.
func (o *TemplatesVM) HasTake() bool {
	if o != nil && o.Take != nil {
		return true
	}

	return false
}

// SetTake gets a reference to the given int32 and assigns it to the Take field.
func (o *TemplatesVM) SetTake(v int32) {
	o.Take = &v
}

func (o TemplatesVM) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Files != nil {
		toSerialize["files"] = o.Files
	}
	if o.Count != nil {
		toSerialize["count"] = o.Count
	}
	if o.Skip != nil {
		toSerialize["skip"] = o.Skip
	}
	if o.Take != nil {
		toSerialize["take"] = o.Take
	}
	return json.Marshal(toSerialize)
}

type NullableTemplatesVM struct {
	value *TemplatesVM
	isSet bool
}

func (v NullableTemplatesVM) Get() *TemplatesVM {
	return v.value
}

func (v *NullableTemplatesVM) Set(val *TemplatesVM) {
	v.value = val
	v.isSet = true
}

func (v NullableTemplatesVM) IsSet() bool {
	return v.isSet
}

func (v *NullableTemplatesVM) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTemplatesVM(val *TemplatesVM) *NullableTemplatesVM {
	return &NullableTemplatesVM{value: val, isSet: true}
}

func (v NullableTemplatesVM) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTemplatesVM) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

