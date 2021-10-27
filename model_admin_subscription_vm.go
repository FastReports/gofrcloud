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

// AdminSubscriptionVM struct for AdminSubscriptionVM
type AdminSubscriptionVM struct {
	DefaultPermissions *DefaultPermissions `json:"defaultPermissions,omitempty"`
	Id NullableString `json:"id,omitempty"`
	Name NullableString `json:"name,omitempty"`
	Locale NullableString `json:"locale,omitempty"`
	Current *SubscriptionPeriodVM `json:"current,omitempty"`
	Old []SubscriptionPeriodVM `json:"old,omitempty"`
	TemplatesFolder *SubscriptionFolder `json:"templatesFolder,omitempty"`
	ReportsFolder *SubscriptionFolder `json:"reportsFolder,omitempty"`
	ExportsFolder *SubscriptionFolder `json:"exportsFolder,omitempty"`
}

// NewAdminSubscriptionVM instantiates a new AdminSubscriptionVM object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdminSubscriptionVM() *AdminSubscriptionVM {
	this := AdminSubscriptionVM{}
	return &this
}

// NewAdminSubscriptionVMWithDefaults instantiates a new AdminSubscriptionVM object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdminSubscriptionVMWithDefaults() *AdminSubscriptionVM {
	this := AdminSubscriptionVM{}
	return &this
}

// GetDefaultPermissions returns the DefaultPermissions field value if set, zero value otherwise.
func (o *AdminSubscriptionVM) GetDefaultPermissions() DefaultPermissions {
	if o == nil || o.DefaultPermissions == nil {
		var ret DefaultPermissions
		return ret
	}
	return *o.DefaultPermissions
}

// GetDefaultPermissionsOk returns a tuple with the DefaultPermissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdminSubscriptionVM) GetDefaultPermissionsOk() (*DefaultPermissions, bool) {
	if o == nil || o.DefaultPermissions == nil {
		return nil, false
	}
	return o.DefaultPermissions, true
}

// HasDefaultPermissions returns a boolean if a field has been set.
func (o *AdminSubscriptionVM) HasDefaultPermissions() bool {
	if o != nil && o.DefaultPermissions != nil {
		return true
	}

	return false
}

// SetDefaultPermissions gets a reference to the given DefaultPermissions and assigns it to the DefaultPermissions field.
func (o *AdminSubscriptionVM) SetDefaultPermissions(v DefaultPermissions) {
	o.DefaultPermissions = &v
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AdminSubscriptionVM) GetId() string {
	if o == nil || o.Id.Get() == nil {
		var ret string
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AdminSubscriptionVM) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *AdminSubscriptionVM) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableString and assigns it to the Id field.
func (o *AdminSubscriptionVM) SetId(v string) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *AdminSubscriptionVM) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *AdminSubscriptionVM) UnsetId() {
	o.Id.Unset()
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AdminSubscriptionVM) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AdminSubscriptionVM) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *AdminSubscriptionVM) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *AdminSubscriptionVM) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *AdminSubscriptionVM) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *AdminSubscriptionVM) UnsetName() {
	o.Name.Unset()
}

// GetLocale returns the Locale field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AdminSubscriptionVM) GetLocale() string {
	if o == nil || o.Locale.Get() == nil {
		var ret string
		return ret
	}
	return *o.Locale.Get()
}

// GetLocaleOk returns a tuple with the Locale field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AdminSubscriptionVM) GetLocaleOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Locale.Get(), o.Locale.IsSet()
}

// HasLocale returns a boolean if a field has been set.
func (o *AdminSubscriptionVM) HasLocale() bool {
	if o != nil && o.Locale.IsSet() {
		return true
	}

	return false
}

// SetLocale gets a reference to the given NullableString and assigns it to the Locale field.
func (o *AdminSubscriptionVM) SetLocale(v string) {
	o.Locale.Set(&v)
}
// SetLocaleNil sets the value for Locale to be an explicit nil
func (o *AdminSubscriptionVM) SetLocaleNil() {
	o.Locale.Set(nil)
}

// UnsetLocale ensures that no value is present for Locale, not even an explicit nil
func (o *AdminSubscriptionVM) UnsetLocale() {
	o.Locale.Unset()
}

// GetCurrent returns the Current field value if set, zero value otherwise.
func (o *AdminSubscriptionVM) GetCurrent() SubscriptionPeriodVM {
	if o == nil || o.Current == nil {
		var ret SubscriptionPeriodVM
		return ret
	}
	return *o.Current
}

// GetCurrentOk returns a tuple with the Current field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdminSubscriptionVM) GetCurrentOk() (*SubscriptionPeriodVM, bool) {
	if o == nil || o.Current == nil {
		return nil, false
	}
	return o.Current, true
}

// HasCurrent returns a boolean if a field has been set.
func (o *AdminSubscriptionVM) HasCurrent() bool {
	if o != nil && o.Current != nil {
		return true
	}

	return false
}

// SetCurrent gets a reference to the given SubscriptionPeriodVM and assigns it to the Current field.
func (o *AdminSubscriptionVM) SetCurrent(v SubscriptionPeriodVM) {
	o.Current = &v
}

// GetOld returns the Old field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AdminSubscriptionVM) GetOld() []SubscriptionPeriodVM {
	if o == nil  {
		var ret []SubscriptionPeriodVM
		return ret
	}
	return o.Old
}

// GetOldOk returns a tuple with the Old field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AdminSubscriptionVM) GetOldOk() (*[]SubscriptionPeriodVM, bool) {
	if o == nil || o.Old == nil {
		return nil, false
	}
	return &o.Old, true
}

// HasOld returns a boolean if a field has been set.
func (o *AdminSubscriptionVM) HasOld() bool {
	if o != nil && o.Old != nil {
		return true
	}

	return false
}

// SetOld gets a reference to the given []SubscriptionPeriodVM and assigns it to the Old field.
func (o *AdminSubscriptionVM) SetOld(v []SubscriptionPeriodVM) {
	o.Old = v
}

// GetTemplatesFolder returns the TemplatesFolder field value if set, zero value otherwise.
func (o *AdminSubscriptionVM) GetTemplatesFolder() SubscriptionFolder {
	if o == nil || o.TemplatesFolder == nil {
		var ret SubscriptionFolder
		return ret
	}
	return *o.TemplatesFolder
}

// GetTemplatesFolderOk returns a tuple with the TemplatesFolder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdminSubscriptionVM) GetTemplatesFolderOk() (*SubscriptionFolder, bool) {
	if o == nil || o.TemplatesFolder == nil {
		return nil, false
	}
	return o.TemplatesFolder, true
}

// HasTemplatesFolder returns a boolean if a field has been set.
func (o *AdminSubscriptionVM) HasTemplatesFolder() bool {
	if o != nil && o.TemplatesFolder != nil {
		return true
	}

	return false
}

// SetTemplatesFolder gets a reference to the given SubscriptionFolder and assigns it to the TemplatesFolder field.
func (o *AdminSubscriptionVM) SetTemplatesFolder(v SubscriptionFolder) {
	o.TemplatesFolder = &v
}

// GetReportsFolder returns the ReportsFolder field value if set, zero value otherwise.
func (o *AdminSubscriptionVM) GetReportsFolder() SubscriptionFolder {
	if o == nil || o.ReportsFolder == nil {
		var ret SubscriptionFolder
		return ret
	}
	return *o.ReportsFolder
}

// GetReportsFolderOk returns a tuple with the ReportsFolder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdminSubscriptionVM) GetReportsFolderOk() (*SubscriptionFolder, bool) {
	if o == nil || o.ReportsFolder == nil {
		return nil, false
	}
	return o.ReportsFolder, true
}

// HasReportsFolder returns a boolean if a field has been set.
func (o *AdminSubscriptionVM) HasReportsFolder() bool {
	if o != nil && o.ReportsFolder != nil {
		return true
	}

	return false
}

// SetReportsFolder gets a reference to the given SubscriptionFolder and assigns it to the ReportsFolder field.
func (o *AdminSubscriptionVM) SetReportsFolder(v SubscriptionFolder) {
	o.ReportsFolder = &v
}

// GetExportsFolder returns the ExportsFolder field value if set, zero value otherwise.
func (o *AdminSubscriptionVM) GetExportsFolder() SubscriptionFolder {
	if o == nil || o.ExportsFolder == nil {
		var ret SubscriptionFolder
		return ret
	}
	return *o.ExportsFolder
}

// GetExportsFolderOk returns a tuple with the ExportsFolder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdminSubscriptionVM) GetExportsFolderOk() (*SubscriptionFolder, bool) {
	if o == nil || o.ExportsFolder == nil {
		return nil, false
	}
	return o.ExportsFolder, true
}

// HasExportsFolder returns a boolean if a field has been set.
func (o *AdminSubscriptionVM) HasExportsFolder() bool {
	if o != nil && o.ExportsFolder != nil {
		return true
	}

	return false
}

// SetExportsFolder gets a reference to the given SubscriptionFolder and assigns it to the ExportsFolder field.
func (o *AdminSubscriptionVM) SetExportsFolder(v SubscriptionFolder) {
	o.ExportsFolder = &v
}

func (o AdminSubscriptionVM) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DefaultPermissions != nil {
		toSerialize["defaultPermissions"] = o.DefaultPermissions
	}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.Locale.IsSet() {
		toSerialize["locale"] = o.Locale.Get()
	}
	if o.Current != nil {
		toSerialize["current"] = o.Current
	}
	if o.Old != nil {
		toSerialize["old"] = o.Old
	}
	if o.TemplatesFolder != nil {
		toSerialize["templatesFolder"] = o.TemplatesFolder
	}
	if o.ReportsFolder != nil {
		toSerialize["reportsFolder"] = o.ReportsFolder
	}
	if o.ExportsFolder != nil {
		toSerialize["exportsFolder"] = o.ExportsFolder
	}
	return json.Marshal(toSerialize)
}

type NullableAdminSubscriptionVM struct {
	value *AdminSubscriptionVM
	isSet bool
}

func (v NullableAdminSubscriptionVM) Get() *AdminSubscriptionVM {
	return v.value
}

func (v *NullableAdminSubscriptionVM) Set(val *AdminSubscriptionVM) {
	v.value = val
	v.isSet = true
}

func (v NullableAdminSubscriptionVM) IsSet() bool {
	return v.isSet
}

func (v *NullableAdminSubscriptionVM) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdminSubscriptionVM(val *AdminSubscriptionVM) *NullableAdminSubscriptionVM {
	return &NullableAdminSubscriptionVM{value: val, isSet: true}
}

func (v NullableAdminSubscriptionVM) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdminSubscriptionVM) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


