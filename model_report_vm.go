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
	"time"
)

// ReportVM struct for ReportVM
type ReportVM struct {
	TemplateId NullableString `json:"templateId,omitempty"`
	ReportInfo *ReportInfo `json:"reportInfo,omitempty"`
	Id NullableString `json:"id,omitempty"`
	CreatedTime *time.Time `json:"createdTime,omitempty"`
	CreatorUserId NullableString `json:"creatorUserId,omitempty"`
	EditedTime *time.Time `json:"editedTime,omitempty"`
	EditorUserId NullableString `json:"editorUserId,omitempty"`
}

// NewReportVM instantiates a new ReportVM object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReportVM() *ReportVM {
	this := ReportVM{}
	return &this
}

// NewReportVMWithDefaults instantiates a new ReportVM object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReportVMWithDefaults() *ReportVM {
	this := ReportVM{}
	return &this
}

// GetTemplateId returns the TemplateId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReportVM) GetTemplateId() string {
	if o == nil || o.TemplateId.Get() == nil {
		var ret string
		return ret
	}
	return *o.TemplateId.Get()
}

// GetTemplateIdOk returns a tuple with the TemplateId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReportVM) GetTemplateIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.TemplateId.Get(), o.TemplateId.IsSet()
}

// HasTemplateId returns a boolean if a field has been set.
func (o *ReportVM) HasTemplateId() bool {
	if o != nil && o.TemplateId.IsSet() {
		return true
	}

	return false
}

// SetTemplateId gets a reference to the given NullableString and assigns it to the TemplateId field.
func (o *ReportVM) SetTemplateId(v string) {
	o.TemplateId.Set(&v)
}
// SetTemplateIdNil sets the value for TemplateId to be an explicit nil
func (o *ReportVM) SetTemplateIdNil() {
	o.TemplateId.Set(nil)
}

// UnsetTemplateId ensures that no value is present for TemplateId, not even an explicit nil
func (o *ReportVM) UnsetTemplateId() {
	o.TemplateId.Unset()
}

// GetReportInfo returns the ReportInfo field value if set, zero value otherwise.
func (o *ReportVM) GetReportInfo() ReportInfo {
	if o == nil || o.ReportInfo == nil {
		var ret ReportInfo
		return ret
	}
	return *o.ReportInfo
}

// GetReportInfoOk returns a tuple with the ReportInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportVM) GetReportInfoOk() (*ReportInfo, bool) {
	if o == nil || o.ReportInfo == nil {
		return nil, false
	}
	return o.ReportInfo, true
}

// HasReportInfo returns a boolean if a field has been set.
func (o *ReportVM) HasReportInfo() bool {
	if o != nil && o.ReportInfo != nil {
		return true
	}

	return false
}

// SetReportInfo gets a reference to the given ReportInfo and assigns it to the ReportInfo field.
func (o *ReportVM) SetReportInfo(v ReportInfo) {
	o.ReportInfo = &v
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReportVM) GetId() string {
	if o == nil || o.Id.Get() == nil {
		var ret string
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReportVM) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *ReportVM) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableString and assigns it to the Id field.
func (o *ReportVM) SetId(v string) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *ReportVM) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *ReportVM) UnsetId() {
	o.Id.Unset()
}

// GetCreatedTime returns the CreatedTime field value if set, zero value otherwise.
func (o *ReportVM) GetCreatedTime() time.Time {
	if o == nil || o.CreatedTime == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedTime
}

// GetCreatedTimeOk returns a tuple with the CreatedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportVM) GetCreatedTimeOk() (*time.Time, bool) {
	if o == nil || o.CreatedTime == nil {
		return nil, false
	}
	return o.CreatedTime, true
}

// HasCreatedTime returns a boolean if a field has been set.
func (o *ReportVM) HasCreatedTime() bool {
	if o != nil && o.CreatedTime != nil {
		return true
	}

	return false
}

// SetCreatedTime gets a reference to the given time.Time and assigns it to the CreatedTime field.
func (o *ReportVM) SetCreatedTime(v time.Time) {
	o.CreatedTime = &v
}

// GetCreatorUserId returns the CreatorUserId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReportVM) GetCreatorUserId() string {
	if o == nil || o.CreatorUserId.Get() == nil {
		var ret string
		return ret
	}
	return *o.CreatorUserId.Get()
}

// GetCreatorUserIdOk returns a tuple with the CreatorUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReportVM) GetCreatorUserIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CreatorUserId.Get(), o.CreatorUserId.IsSet()
}

// HasCreatorUserId returns a boolean if a field has been set.
func (o *ReportVM) HasCreatorUserId() bool {
	if o != nil && o.CreatorUserId.IsSet() {
		return true
	}

	return false
}

// SetCreatorUserId gets a reference to the given NullableString and assigns it to the CreatorUserId field.
func (o *ReportVM) SetCreatorUserId(v string) {
	o.CreatorUserId.Set(&v)
}
// SetCreatorUserIdNil sets the value for CreatorUserId to be an explicit nil
func (o *ReportVM) SetCreatorUserIdNil() {
	o.CreatorUserId.Set(nil)
}

// UnsetCreatorUserId ensures that no value is present for CreatorUserId, not even an explicit nil
func (o *ReportVM) UnsetCreatorUserId() {
	o.CreatorUserId.Unset()
}

// GetEditedTime returns the EditedTime field value if set, zero value otherwise.
func (o *ReportVM) GetEditedTime() time.Time {
	if o == nil || o.EditedTime == nil {
		var ret time.Time
		return ret
	}
	return *o.EditedTime
}

// GetEditedTimeOk returns a tuple with the EditedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportVM) GetEditedTimeOk() (*time.Time, bool) {
	if o == nil || o.EditedTime == nil {
		return nil, false
	}
	return o.EditedTime, true
}

// HasEditedTime returns a boolean if a field has been set.
func (o *ReportVM) HasEditedTime() bool {
	if o != nil && o.EditedTime != nil {
		return true
	}

	return false
}

// SetEditedTime gets a reference to the given time.Time and assigns it to the EditedTime field.
func (o *ReportVM) SetEditedTime(v time.Time) {
	o.EditedTime = &v
}

// GetEditorUserId returns the EditorUserId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReportVM) GetEditorUserId() string {
	if o == nil || o.EditorUserId.Get() == nil {
		var ret string
		return ret
	}
	return *o.EditorUserId.Get()
}

// GetEditorUserIdOk returns a tuple with the EditorUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReportVM) GetEditorUserIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.EditorUserId.Get(), o.EditorUserId.IsSet()
}

// HasEditorUserId returns a boolean if a field has been set.
func (o *ReportVM) HasEditorUserId() bool {
	if o != nil && o.EditorUserId.IsSet() {
		return true
	}

	return false
}

// SetEditorUserId gets a reference to the given NullableString and assigns it to the EditorUserId field.
func (o *ReportVM) SetEditorUserId(v string) {
	o.EditorUserId.Set(&v)
}
// SetEditorUserIdNil sets the value for EditorUserId to be an explicit nil
func (o *ReportVM) SetEditorUserIdNil() {
	o.EditorUserId.Set(nil)
}

// UnsetEditorUserId ensures that no value is present for EditorUserId, not even an explicit nil
func (o *ReportVM) UnsetEditorUserId() {
	o.EditorUserId.Unset()
}

func (o ReportVM) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.TemplateId.IsSet() {
		toSerialize["templateId"] = o.TemplateId.Get()
	}
	if o.ReportInfo != nil {
		toSerialize["reportInfo"] = o.ReportInfo
	}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if o.CreatedTime != nil {
		toSerialize["createdTime"] = o.CreatedTime
	}
	if o.CreatorUserId.IsSet() {
		toSerialize["creatorUserId"] = o.CreatorUserId.Get()
	}
	if o.EditedTime != nil {
		toSerialize["editedTime"] = o.EditedTime
	}
	if o.EditorUserId.IsSet() {
		toSerialize["editorUserId"] = o.EditorUserId.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableReportVM struct {
	value *ReportVM
	isSet bool
}

func (v NullableReportVM) Get() *ReportVM {
	return v.value
}

func (v *NullableReportVM) Set(val *ReportVM) {
	v.value = val
	v.isSet = true
}

func (v NullableReportVM) IsSet() bool {
	return v.isSet
}

func (v *NullableReportVM) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReportVM(val *ReportVM) *NullableReportVM {
	return &NullableReportVM{value: val, isSet: true}
}

func (v NullableReportVM) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReportVM) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


