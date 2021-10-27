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

// SubscriptionPlanVM struct for SubscriptionPlanVM
type SubscriptionPlanVM struct {
	Id NullableString `json:"id,omitempty"`
	IsActive NullableBool `json:"isActive,omitempty"`
	DisplayName NullableString `json:"displayName,omitempty"`
	TimePeriodType *TimePeriodType `json:"timePeriodType,omitempty"`
	TimePeriod NullableInt32 `json:"timePeriod,omitempty"`
	ReadonlyTimeLimitType *TimePeriodType `json:"readonlyTimeLimitType,omitempty"`
	ReadonlyTimeLimit *int32 `json:"readonlyTimeLimit,omitempty"`
	TemplatesSpaceLimit NullableInt64 `json:"templatesSpaceLimit,omitempty"`
	ReportsSpaceLimit NullableInt64 `json:"reportsSpaceLimit,omitempty"`
	ExportsSpaceLimit NullableInt64 `json:"exportsSpaceLimit,omitempty"`
	FileUploadSizeLimit NullableInt64 `json:"fileUploadSizeLimit,omitempty"`
	DataSourceLimit NullableInt32 `json:"dataSourceLimit,omitempty"`
	MaxUsersCount NullableInt32 `json:"maxUsersCount,omitempty"`
	HasSpaceOverdraft NullableBool `json:"hasSpaceOverdraft,omitempty"`
	GroupLimit NullableInt32 `json:"groupLimit,omitempty"`
	OnlineDesigner NullableBool `json:"onlineDesigner,omitempty"`
	IsDemo NullableBool `json:"isDemo,omitempty"`
	UrlToBuy NullableString `json:"urlToBuy,omitempty"`
	UnlimitedPage *bool `json:"unlimitedPage,omitempty"`
	PageLimit *int32 `json:"pageLimit,omitempty"`
	Tasks *TaskSettingsVM `json:"tasks,omitempty"`
}

// NewSubscriptionPlanVM instantiates a new SubscriptionPlanVM object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscriptionPlanVM() *SubscriptionPlanVM {
	this := SubscriptionPlanVM{}
	return &this
}

// NewSubscriptionPlanVMWithDefaults instantiates a new SubscriptionPlanVM object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionPlanVMWithDefaults() *SubscriptionPlanVM {
	this := SubscriptionPlanVM{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SubscriptionPlanVM) GetId() string {
	if o == nil || o.Id.Get() == nil {
		var ret string
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SubscriptionPlanVM) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *SubscriptionPlanVM) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableString and assigns it to the Id field.
func (o *SubscriptionPlanVM) SetId(v string) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *SubscriptionPlanVM) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *SubscriptionPlanVM) UnsetId() {
	o.Id.Unset()
}

// GetIsActive returns the IsActive field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SubscriptionPlanVM) GetIsActive() bool {
	if o == nil || o.IsActive.Get() == nil {
		var ret bool
		return ret
	}
	return *o.IsActive.Get()
}

// GetIsActiveOk returns a tuple with the IsActive field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SubscriptionPlanVM) GetIsActiveOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.IsActive.Get(), o.IsActive.IsSet()
}

// HasIsActive returns a boolean if a field has been set.
func (o *SubscriptionPlanVM) HasIsActive() bool {
	if o != nil && o.IsActive.IsSet() {
		return true
	}

	return false
}

// SetIsActive gets a reference to the given NullableBool and assigns it to the IsActive field.
func (o *SubscriptionPlanVM) SetIsActive(v bool) {
	o.IsActive.Set(&v)
}
// SetIsActiveNil sets the value for IsActive to be an explicit nil
func (o *SubscriptionPlanVM) SetIsActiveNil() {
	o.IsActive.Set(nil)
}

// UnsetIsActive ensures that no value is present for IsActive, not even an explicit nil
func (o *SubscriptionPlanVM) UnsetIsActive() {
	o.IsActive.Unset()
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SubscriptionPlanVM) GetDisplayName() string {
	if o == nil || o.DisplayName.Get() == nil {
		var ret string
		return ret
	}
	return *o.DisplayName.Get()
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SubscriptionPlanVM) GetDisplayNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DisplayName.Get(), o.DisplayName.IsSet()
}

// HasDisplayName returns a boolean if a field has been set.
func (o *SubscriptionPlanVM) HasDisplayName() bool {
	if o != nil && o.DisplayName.IsSet() {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given NullableString and assigns it to the DisplayName field.
func (o *SubscriptionPlanVM) SetDisplayName(v string) {
	o.DisplayName.Set(&v)
}
// SetDisplayNameNil sets the value for DisplayName to be an explicit nil
func (o *SubscriptionPlanVM) SetDisplayNameNil() {
	o.DisplayName.Set(nil)
}

// UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
func (o *SubscriptionPlanVM) UnsetDisplayName() {
	o.DisplayName.Unset()
}

// GetTimePeriodType returns the TimePeriodType field value if set, zero value otherwise.
func (o *SubscriptionPlanVM) GetTimePeriodType() TimePeriodType {
	if o == nil || o.TimePeriodType == nil {
		var ret TimePeriodType
		return ret
	}
	return *o.TimePeriodType
}

// GetTimePeriodTypeOk returns a tuple with the TimePeriodType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionPlanVM) GetTimePeriodTypeOk() (*TimePeriodType, bool) {
	if o == nil || o.TimePeriodType == nil {
		return nil, false
	}
	return o.TimePeriodType, true
}

// HasTimePeriodType returns a boolean if a field has been set.
func (o *SubscriptionPlanVM) HasTimePeriodType() bool {
	if o != nil && o.TimePeriodType != nil {
		return true
	}

	return false
}

// SetTimePeriodType gets a reference to the given TimePeriodType and assigns it to the TimePeriodType field.
func (o *SubscriptionPlanVM) SetTimePeriodType(v TimePeriodType) {
	o.TimePeriodType = &v
}

// GetTimePeriod returns the TimePeriod field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SubscriptionPlanVM) GetTimePeriod() int32 {
	if o == nil || o.TimePeriod.Get() == nil {
		var ret int32
		return ret
	}
	return *o.TimePeriod.Get()
}

// GetTimePeriodOk returns a tuple with the TimePeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SubscriptionPlanVM) GetTimePeriodOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.TimePeriod.Get(), o.TimePeriod.IsSet()
}

// HasTimePeriod returns a boolean if a field has been set.
func (o *SubscriptionPlanVM) HasTimePeriod() bool {
	if o != nil && o.TimePeriod.IsSet() {
		return true
	}

	return false
}

// SetTimePeriod gets a reference to the given NullableInt32 and assigns it to the TimePeriod field.
func (o *SubscriptionPlanVM) SetTimePeriod(v int32) {
	o.TimePeriod.Set(&v)
}
// SetTimePeriodNil sets the value for TimePeriod to be an explicit nil
func (o *SubscriptionPlanVM) SetTimePeriodNil() {
	o.TimePeriod.Set(nil)
}

// UnsetTimePeriod ensures that no value is present for TimePeriod, not even an explicit nil
func (o *SubscriptionPlanVM) UnsetTimePeriod() {
	o.TimePeriod.Unset()
}

// GetReadonlyTimeLimitType returns the ReadonlyTimeLimitType field value if set, zero value otherwise.
func (o *SubscriptionPlanVM) GetReadonlyTimeLimitType() TimePeriodType {
	if o == nil || o.ReadonlyTimeLimitType == nil {
		var ret TimePeriodType
		return ret
	}
	return *o.ReadonlyTimeLimitType
}

// GetReadonlyTimeLimitTypeOk returns a tuple with the ReadonlyTimeLimitType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionPlanVM) GetReadonlyTimeLimitTypeOk() (*TimePeriodType, bool) {
	if o == nil || o.ReadonlyTimeLimitType == nil {
		return nil, false
	}
	return o.ReadonlyTimeLimitType, true
}

// HasReadonlyTimeLimitType returns a boolean if a field has been set.
func (o *SubscriptionPlanVM) HasReadonlyTimeLimitType() bool {
	if o != nil && o.ReadonlyTimeLimitType != nil {
		return true
	}

	return false
}

// SetReadonlyTimeLimitType gets a reference to the given TimePeriodType and assigns it to the ReadonlyTimeLimitType field.
func (o *SubscriptionPlanVM) SetReadonlyTimeLimitType(v TimePeriodType) {
	o.ReadonlyTimeLimitType = &v
}

// GetReadonlyTimeLimit returns the ReadonlyTimeLimit field value if set, zero value otherwise.
func (o *SubscriptionPlanVM) GetReadonlyTimeLimit() int32 {
	if o == nil || o.ReadonlyTimeLimit == nil {
		var ret int32
		return ret
	}
	return *o.ReadonlyTimeLimit
}

// GetReadonlyTimeLimitOk returns a tuple with the ReadonlyTimeLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionPlanVM) GetReadonlyTimeLimitOk() (*int32, bool) {
	if o == nil || o.ReadonlyTimeLimit == nil {
		return nil, false
	}
	return o.ReadonlyTimeLimit, true
}

// HasReadonlyTimeLimit returns a boolean if a field has been set.
func (o *SubscriptionPlanVM) HasReadonlyTimeLimit() bool {
	if o != nil && o.ReadonlyTimeLimit != nil {
		return true
	}

	return false
}

// SetReadonlyTimeLimit gets a reference to the given int32 and assigns it to the ReadonlyTimeLimit field.
func (o *SubscriptionPlanVM) SetReadonlyTimeLimit(v int32) {
	o.ReadonlyTimeLimit = &v
}

// GetTemplatesSpaceLimit returns the TemplatesSpaceLimit field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SubscriptionPlanVM) GetTemplatesSpaceLimit() int64 {
	if o == nil || o.TemplatesSpaceLimit.Get() == nil {
		var ret int64
		return ret
	}
	return *o.TemplatesSpaceLimit.Get()
}

// GetTemplatesSpaceLimitOk returns a tuple with the TemplatesSpaceLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SubscriptionPlanVM) GetTemplatesSpaceLimitOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.TemplatesSpaceLimit.Get(), o.TemplatesSpaceLimit.IsSet()
}

// HasTemplatesSpaceLimit returns a boolean if a field has been set.
func (o *SubscriptionPlanVM) HasTemplatesSpaceLimit() bool {
	if o != nil && o.TemplatesSpaceLimit.IsSet() {
		return true
	}

	return false
}

// SetTemplatesSpaceLimit gets a reference to the given NullableInt64 and assigns it to the TemplatesSpaceLimit field.
func (o *SubscriptionPlanVM) SetTemplatesSpaceLimit(v int64) {
	o.TemplatesSpaceLimit.Set(&v)
}
// SetTemplatesSpaceLimitNil sets the value for TemplatesSpaceLimit to be an explicit nil
func (o *SubscriptionPlanVM) SetTemplatesSpaceLimitNil() {
	o.TemplatesSpaceLimit.Set(nil)
}

// UnsetTemplatesSpaceLimit ensures that no value is present for TemplatesSpaceLimit, not even an explicit nil
func (o *SubscriptionPlanVM) UnsetTemplatesSpaceLimit() {
	o.TemplatesSpaceLimit.Unset()
}

// GetReportsSpaceLimit returns the ReportsSpaceLimit field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SubscriptionPlanVM) GetReportsSpaceLimit() int64 {
	if o == nil || o.ReportsSpaceLimit.Get() == nil {
		var ret int64
		return ret
	}
	return *o.ReportsSpaceLimit.Get()
}

// GetReportsSpaceLimitOk returns a tuple with the ReportsSpaceLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SubscriptionPlanVM) GetReportsSpaceLimitOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ReportsSpaceLimit.Get(), o.ReportsSpaceLimit.IsSet()
}

// HasReportsSpaceLimit returns a boolean if a field has been set.
func (o *SubscriptionPlanVM) HasReportsSpaceLimit() bool {
	if o != nil && o.ReportsSpaceLimit.IsSet() {
		return true
	}

	return false
}

// SetReportsSpaceLimit gets a reference to the given NullableInt64 and assigns it to the ReportsSpaceLimit field.
func (o *SubscriptionPlanVM) SetReportsSpaceLimit(v int64) {
	o.ReportsSpaceLimit.Set(&v)
}
// SetReportsSpaceLimitNil sets the value for ReportsSpaceLimit to be an explicit nil
func (o *SubscriptionPlanVM) SetReportsSpaceLimitNil() {
	o.ReportsSpaceLimit.Set(nil)
}

// UnsetReportsSpaceLimit ensures that no value is present for ReportsSpaceLimit, not even an explicit nil
func (o *SubscriptionPlanVM) UnsetReportsSpaceLimit() {
	o.ReportsSpaceLimit.Unset()
}

// GetExportsSpaceLimit returns the ExportsSpaceLimit field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SubscriptionPlanVM) GetExportsSpaceLimit() int64 {
	if o == nil || o.ExportsSpaceLimit.Get() == nil {
		var ret int64
		return ret
	}
	return *o.ExportsSpaceLimit.Get()
}

// GetExportsSpaceLimitOk returns a tuple with the ExportsSpaceLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SubscriptionPlanVM) GetExportsSpaceLimitOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ExportsSpaceLimit.Get(), o.ExportsSpaceLimit.IsSet()
}

// HasExportsSpaceLimit returns a boolean if a field has been set.
func (o *SubscriptionPlanVM) HasExportsSpaceLimit() bool {
	if o != nil && o.ExportsSpaceLimit.IsSet() {
		return true
	}

	return false
}

// SetExportsSpaceLimit gets a reference to the given NullableInt64 and assigns it to the ExportsSpaceLimit field.
func (o *SubscriptionPlanVM) SetExportsSpaceLimit(v int64) {
	o.ExportsSpaceLimit.Set(&v)
}
// SetExportsSpaceLimitNil sets the value for ExportsSpaceLimit to be an explicit nil
func (o *SubscriptionPlanVM) SetExportsSpaceLimitNil() {
	o.ExportsSpaceLimit.Set(nil)
}

// UnsetExportsSpaceLimit ensures that no value is present for ExportsSpaceLimit, not even an explicit nil
func (o *SubscriptionPlanVM) UnsetExportsSpaceLimit() {
	o.ExportsSpaceLimit.Unset()
}

// GetFileUploadSizeLimit returns the FileUploadSizeLimit field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SubscriptionPlanVM) GetFileUploadSizeLimit() int64 {
	if o == nil || o.FileUploadSizeLimit.Get() == nil {
		var ret int64
		return ret
	}
	return *o.FileUploadSizeLimit.Get()
}

// GetFileUploadSizeLimitOk returns a tuple with the FileUploadSizeLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SubscriptionPlanVM) GetFileUploadSizeLimitOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.FileUploadSizeLimit.Get(), o.FileUploadSizeLimit.IsSet()
}

// HasFileUploadSizeLimit returns a boolean if a field has been set.
func (o *SubscriptionPlanVM) HasFileUploadSizeLimit() bool {
	if o != nil && o.FileUploadSizeLimit.IsSet() {
		return true
	}

	return false
}

// SetFileUploadSizeLimit gets a reference to the given NullableInt64 and assigns it to the FileUploadSizeLimit field.
func (o *SubscriptionPlanVM) SetFileUploadSizeLimit(v int64) {
	o.FileUploadSizeLimit.Set(&v)
}
// SetFileUploadSizeLimitNil sets the value for FileUploadSizeLimit to be an explicit nil
func (o *SubscriptionPlanVM) SetFileUploadSizeLimitNil() {
	o.FileUploadSizeLimit.Set(nil)
}

// UnsetFileUploadSizeLimit ensures that no value is present for FileUploadSizeLimit, not even an explicit nil
func (o *SubscriptionPlanVM) UnsetFileUploadSizeLimit() {
	o.FileUploadSizeLimit.Unset()
}

// GetDataSourceLimit returns the DataSourceLimit field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SubscriptionPlanVM) GetDataSourceLimit() int32 {
	if o == nil || o.DataSourceLimit.Get() == nil {
		var ret int32
		return ret
	}
	return *o.DataSourceLimit.Get()
}

// GetDataSourceLimitOk returns a tuple with the DataSourceLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SubscriptionPlanVM) GetDataSourceLimitOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DataSourceLimit.Get(), o.DataSourceLimit.IsSet()
}

// HasDataSourceLimit returns a boolean if a field has been set.
func (o *SubscriptionPlanVM) HasDataSourceLimit() bool {
	if o != nil && o.DataSourceLimit.IsSet() {
		return true
	}

	return false
}

// SetDataSourceLimit gets a reference to the given NullableInt32 and assigns it to the DataSourceLimit field.
func (o *SubscriptionPlanVM) SetDataSourceLimit(v int32) {
	o.DataSourceLimit.Set(&v)
}
// SetDataSourceLimitNil sets the value for DataSourceLimit to be an explicit nil
func (o *SubscriptionPlanVM) SetDataSourceLimitNil() {
	o.DataSourceLimit.Set(nil)
}

// UnsetDataSourceLimit ensures that no value is present for DataSourceLimit, not even an explicit nil
func (o *SubscriptionPlanVM) UnsetDataSourceLimit() {
	o.DataSourceLimit.Unset()
}

// GetMaxUsersCount returns the MaxUsersCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SubscriptionPlanVM) GetMaxUsersCount() int32 {
	if o == nil || o.MaxUsersCount.Get() == nil {
		var ret int32
		return ret
	}
	return *o.MaxUsersCount.Get()
}

// GetMaxUsersCountOk returns a tuple with the MaxUsersCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SubscriptionPlanVM) GetMaxUsersCountOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.MaxUsersCount.Get(), o.MaxUsersCount.IsSet()
}

// HasMaxUsersCount returns a boolean if a field has been set.
func (o *SubscriptionPlanVM) HasMaxUsersCount() bool {
	if o != nil && o.MaxUsersCount.IsSet() {
		return true
	}

	return false
}

// SetMaxUsersCount gets a reference to the given NullableInt32 and assigns it to the MaxUsersCount field.
func (o *SubscriptionPlanVM) SetMaxUsersCount(v int32) {
	o.MaxUsersCount.Set(&v)
}
// SetMaxUsersCountNil sets the value for MaxUsersCount to be an explicit nil
func (o *SubscriptionPlanVM) SetMaxUsersCountNil() {
	o.MaxUsersCount.Set(nil)
}

// UnsetMaxUsersCount ensures that no value is present for MaxUsersCount, not even an explicit nil
func (o *SubscriptionPlanVM) UnsetMaxUsersCount() {
	o.MaxUsersCount.Unset()
}

// GetHasSpaceOverdraft returns the HasSpaceOverdraft field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SubscriptionPlanVM) GetHasSpaceOverdraft() bool {
	if o == nil || o.HasSpaceOverdraft.Get() == nil {
		var ret bool
		return ret
	}
	return *o.HasSpaceOverdraft.Get()
}

// GetHasSpaceOverdraftOk returns a tuple with the HasSpaceOverdraft field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SubscriptionPlanVM) GetHasSpaceOverdraftOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.HasSpaceOverdraft.Get(), o.HasSpaceOverdraft.IsSet()
}

// HasHasSpaceOverdraft returns a boolean if a field has been set.
func (o *SubscriptionPlanVM) HasHasSpaceOverdraft() bool {
	if o != nil && o.HasSpaceOverdraft.IsSet() {
		return true
	}

	return false
}

// SetHasSpaceOverdraft gets a reference to the given NullableBool and assigns it to the HasSpaceOverdraft field.
func (o *SubscriptionPlanVM) SetHasSpaceOverdraft(v bool) {
	o.HasSpaceOverdraft.Set(&v)
}
// SetHasSpaceOverdraftNil sets the value for HasSpaceOverdraft to be an explicit nil
func (o *SubscriptionPlanVM) SetHasSpaceOverdraftNil() {
	o.HasSpaceOverdraft.Set(nil)
}

// UnsetHasSpaceOverdraft ensures that no value is present for HasSpaceOverdraft, not even an explicit nil
func (o *SubscriptionPlanVM) UnsetHasSpaceOverdraft() {
	o.HasSpaceOverdraft.Unset()
}

// GetGroupLimit returns the GroupLimit field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SubscriptionPlanVM) GetGroupLimit() int32 {
	if o == nil || o.GroupLimit.Get() == nil {
		var ret int32
		return ret
	}
	return *o.GroupLimit.Get()
}

// GetGroupLimitOk returns a tuple with the GroupLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SubscriptionPlanVM) GetGroupLimitOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.GroupLimit.Get(), o.GroupLimit.IsSet()
}

// HasGroupLimit returns a boolean if a field has been set.
func (o *SubscriptionPlanVM) HasGroupLimit() bool {
	if o != nil && o.GroupLimit.IsSet() {
		return true
	}

	return false
}

// SetGroupLimit gets a reference to the given NullableInt32 and assigns it to the GroupLimit field.
func (o *SubscriptionPlanVM) SetGroupLimit(v int32) {
	o.GroupLimit.Set(&v)
}
// SetGroupLimitNil sets the value for GroupLimit to be an explicit nil
func (o *SubscriptionPlanVM) SetGroupLimitNil() {
	o.GroupLimit.Set(nil)
}

// UnsetGroupLimit ensures that no value is present for GroupLimit, not even an explicit nil
func (o *SubscriptionPlanVM) UnsetGroupLimit() {
	o.GroupLimit.Unset()
}

// GetOnlineDesigner returns the OnlineDesigner field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SubscriptionPlanVM) GetOnlineDesigner() bool {
	if o == nil || o.OnlineDesigner.Get() == nil {
		var ret bool
		return ret
	}
	return *o.OnlineDesigner.Get()
}

// GetOnlineDesignerOk returns a tuple with the OnlineDesigner field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SubscriptionPlanVM) GetOnlineDesignerOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.OnlineDesigner.Get(), o.OnlineDesigner.IsSet()
}

// HasOnlineDesigner returns a boolean if a field has been set.
func (o *SubscriptionPlanVM) HasOnlineDesigner() bool {
	if o != nil && o.OnlineDesigner.IsSet() {
		return true
	}

	return false
}

// SetOnlineDesigner gets a reference to the given NullableBool and assigns it to the OnlineDesigner field.
func (o *SubscriptionPlanVM) SetOnlineDesigner(v bool) {
	o.OnlineDesigner.Set(&v)
}
// SetOnlineDesignerNil sets the value for OnlineDesigner to be an explicit nil
func (o *SubscriptionPlanVM) SetOnlineDesignerNil() {
	o.OnlineDesigner.Set(nil)
}

// UnsetOnlineDesigner ensures that no value is present for OnlineDesigner, not even an explicit nil
func (o *SubscriptionPlanVM) UnsetOnlineDesigner() {
	o.OnlineDesigner.Unset()
}

// GetIsDemo returns the IsDemo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SubscriptionPlanVM) GetIsDemo() bool {
	if o == nil || o.IsDemo.Get() == nil {
		var ret bool
		return ret
	}
	return *o.IsDemo.Get()
}

// GetIsDemoOk returns a tuple with the IsDemo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SubscriptionPlanVM) GetIsDemoOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.IsDemo.Get(), o.IsDemo.IsSet()
}

// HasIsDemo returns a boolean if a field has been set.
func (o *SubscriptionPlanVM) HasIsDemo() bool {
	if o != nil && o.IsDemo.IsSet() {
		return true
	}

	return false
}

// SetIsDemo gets a reference to the given NullableBool and assigns it to the IsDemo field.
func (o *SubscriptionPlanVM) SetIsDemo(v bool) {
	o.IsDemo.Set(&v)
}
// SetIsDemoNil sets the value for IsDemo to be an explicit nil
func (o *SubscriptionPlanVM) SetIsDemoNil() {
	o.IsDemo.Set(nil)
}

// UnsetIsDemo ensures that no value is present for IsDemo, not even an explicit nil
func (o *SubscriptionPlanVM) UnsetIsDemo() {
	o.IsDemo.Unset()
}

// GetUrlToBuy returns the UrlToBuy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SubscriptionPlanVM) GetUrlToBuy() string {
	if o == nil || o.UrlToBuy.Get() == nil {
		var ret string
		return ret
	}
	return *o.UrlToBuy.Get()
}

// GetUrlToBuyOk returns a tuple with the UrlToBuy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SubscriptionPlanVM) GetUrlToBuyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.UrlToBuy.Get(), o.UrlToBuy.IsSet()
}

// HasUrlToBuy returns a boolean if a field has been set.
func (o *SubscriptionPlanVM) HasUrlToBuy() bool {
	if o != nil && o.UrlToBuy.IsSet() {
		return true
	}

	return false
}

// SetUrlToBuy gets a reference to the given NullableString and assigns it to the UrlToBuy field.
func (o *SubscriptionPlanVM) SetUrlToBuy(v string) {
	o.UrlToBuy.Set(&v)
}
// SetUrlToBuyNil sets the value for UrlToBuy to be an explicit nil
func (o *SubscriptionPlanVM) SetUrlToBuyNil() {
	o.UrlToBuy.Set(nil)
}

// UnsetUrlToBuy ensures that no value is present for UrlToBuy, not even an explicit nil
func (o *SubscriptionPlanVM) UnsetUrlToBuy() {
	o.UrlToBuy.Unset()
}

// GetUnlimitedPage returns the UnlimitedPage field value if set, zero value otherwise.
func (o *SubscriptionPlanVM) GetUnlimitedPage() bool {
	if o == nil || o.UnlimitedPage == nil {
		var ret bool
		return ret
	}
	return *o.UnlimitedPage
}

// GetUnlimitedPageOk returns a tuple with the UnlimitedPage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionPlanVM) GetUnlimitedPageOk() (*bool, bool) {
	if o == nil || o.UnlimitedPage == nil {
		return nil, false
	}
	return o.UnlimitedPage, true
}

// HasUnlimitedPage returns a boolean if a field has been set.
func (o *SubscriptionPlanVM) HasUnlimitedPage() bool {
	if o != nil && o.UnlimitedPage != nil {
		return true
	}

	return false
}

// SetUnlimitedPage gets a reference to the given bool and assigns it to the UnlimitedPage field.
func (o *SubscriptionPlanVM) SetUnlimitedPage(v bool) {
	o.UnlimitedPage = &v
}

// GetPageLimit returns the PageLimit field value if set, zero value otherwise.
func (o *SubscriptionPlanVM) GetPageLimit() int32 {
	if o == nil || o.PageLimit == nil {
		var ret int32
		return ret
	}
	return *o.PageLimit
}

// GetPageLimitOk returns a tuple with the PageLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionPlanVM) GetPageLimitOk() (*int32, bool) {
	if o == nil || o.PageLimit == nil {
		return nil, false
	}
	return o.PageLimit, true
}

// HasPageLimit returns a boolean if a field has been set.
func (o *SubscriptionPlanVM) HasPageLimit() bool {
	if o != nil && o.PageLimit != nil {
		return true
	}

	return false
}

// SetPageLimit gets a reference to the given int32 and assigns it to the PageLimit field.
func (o *SubscriptionPlanVM) SetPageLimit(v int32) {
	o.PageLimit = &v
}

// GetTasks returns the Tasks field value if set, zero value otherwise.
func (o *SubscriptionPlanVM) GetTasks() TaskSettingsVM {
	if o == nil || o.Tasks == nil {
		var ret TaskSettingsVM
		return ret
	}
	return *o.Tasks
}

// GetTasksOk returns a tuple with the Tasks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionPlanVM) GetTasksOk() (*TaskSettingsVM, bool) {
	if o == nil || o.Tasks == nil {
		return nil, false
	}
	return o.Tasks, true
}

// HasTasks returns a boolean if a field has been set.
func (o *SubscriptionPlanVM) HasTasks() bool {
	if o != nil && o.Tasks != nil {
		return true
	}

	return false
}

// SetTasks gets a reference to the given TaskSettingsVM and assigns it to the Tasks field.
func (o *SubscriptionPlanVM) SetTasks(v TaskSettingsVM) {
	o.Tasks = &v
}

func (o SubscriptionPlanVM) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if o.IsActive.IsSet() {
		toSerialize["isActive"] = o.IsActive.Get()
	}
	if o.DisplayName.IsSet() {
		toSerialize["displayName"] = o.DisplayName.Get()
	}
	if o.TimePeriodType != nil {
		toSerialize["timePeriodType"] = o.TimePeriodType
	}
	if o.TimePeriod.IsSet() {
		toSerialize["timePeriod"] = o.TimePeriod.Get()
	}
	if o.ReadonlyTimeLimitType != nil {
		toSerialize["readonlyTimeLimitType"] = o.ReadonlyTimeLimitType
	}
	if o.ReadonlyTimeLimit != nil {
		toSerialize["readonlyTimeLimit"] = o.ReadonlyTimeLimit
	}
	if o.TemplatesSpaceLimit.IsSet() {
		toSerialize["templatesSpaceLimit"] = o.TemplatesSpaceLimit.Get()
	}
	if o.ReportsSpaceLimit.IsSet() {
		toSerialize["reportsSpaceLimit"] = o.ReportsSpaceLimit.Get()
	}
	if o.ExportsSpaceLimit.IsSet() {
		toSerialize["exportsSpaceLimit"] = o.ExportsSpaceLimit.Get()
	}
	if o.FileUploadSizeLimit.IsSet() {
		toSerialize["fileUploadSizeLimit"] = o.FileUploadSizeLimit.Get()
	}
	if o.DataSourceLimit.IsSet() {
		toSerialize["dataSourceLimit"] = o.DataSourceLimit.Get()
	}
	if o.MaxUsersCount.IsSet() {
		toSerialize["maxUsersCount"] = o.MaxUsersCount.Get()
	}
	if o.HasSpaceOverdraft.IsSet() {
		toSerialize["hasSpaceOverdraft"] = o.HasSpaceOverdraft.Get()
	}
	if o.GroupLimit.IsSet() {
		toSerialize["groupLimit"] = o.GroupLimit.Get()
	}
	if o.OnlineDesigner.IsSet() {
		toSerialize["onlineDesigner"] = o.OnlineDesigner.Get()
	}
	if o.IsDemo.IsSet() {
		toSerialize["isDemo"] = o.IsDemo.Get()
	}
	if o.UrlToBuy.IsSet() {
		toSerialize["urlToBuy"] = o.UrlToBuy.Get()
	}
	if o.UnlimitedPage != nil {
		toSerialize["unlimitedPage"] = o.UnlimitedPage
	}
	if o.PageLimit != nil {
		toSerialize["pageLimit"] = o.PageLimit
	}
	if o.Tasks != nil {
		toSerialize["tasks"] = o.Tasks
	}
	return json.Marshal(toSerialize)
}

type NullableSubscriptionPlanVM struct {
	value *SubscriptionPlanVM
	isSet bool
}

func (v NullableSubscriptionPlanVM) Get() *SubscriptionPlanVM {
	return v.value
}

func (v *NullableSubscriptionPlanVM) Set(val *SubscriptionPlanVM) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionPlanVM) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionPlanVM) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionPlanVM(val *SubscriptionPlanVM) *NullableSubscriptionPlanVM {
	return &NullableSubscriptionPlanVM{value: val, isSet: true}
}

func (v NullableSubscriptionPlanVM) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionPlanVM) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


