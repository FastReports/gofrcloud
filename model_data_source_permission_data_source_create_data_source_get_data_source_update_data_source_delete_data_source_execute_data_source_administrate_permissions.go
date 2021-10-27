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

// DataSourcePermissionDataSourceCreateDataSourceGetDataSourceUpdateDataSourceDeleteDataSourceExecuteDataSourceAdministratePermissions struct for DataSourcePermissionDataSourceCreateDataSourceGetDataSourceUpdateDataSourceDeleteDataSourceExecuteDataSourceAdministratePermissions
type DataSourcePermissionDataSourceCreateDataSourceGetDataSourceUpdateDataSourceDeleteDataSourceExecuteDataSourceAdministratePermissions struct {
	OwnerId NullableString `json:"ownerId,omitempty"`
	Owner *DataSourcePermission `json:"owner,omitempty"`
	Groups map[string]DataSourcePermission `json:"groups,omitempty"`
	Other *DataSourcePermission `json:"other,omitempty"`
	Anon *DataSourcePermission `json:"anon,omitempty"`
}

// NewDataSourcePermissionDataSourceCreateDataSourceGetDataSourceUpdateDataSourceDeleteDataSourceExecuteDataSourceAdministratePermissions instantiates a new DataSourcePermissionDataSourceCreateDataSourceGetDataSourceUpdateDataSourceDeleteDataSourceExecuteDataSourceAdministratePermissions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDataSourcePermissionDataSourceCreateDataSourceGetDataSourceUpdateDataSourceDeleteDataSourceExecuteDataSourceAdministratePermissions() *DataSourcePermissionDataSourceCreateDataSourceGetDataSourceUpdateDataSourceDeleteDataSourceExecuteDataSourceAdministratePermissions {
	this := DataSourcePermissionDataSourceCreateDataSourceGetDataSourceUpdateDataSourceDeleteDataSourceExecuteDataSourceAdministratePermissions{}
	return &this
}

// NewDataSourcePermissionDataSourceCreateDataSourceGetDataSourceUpdateDataSourceDeleteDataSourceExecuteDataSourceAdministratePermissionsWithDefaults instantiates a new DataSourcePermissionDataSourceCreateDataSourceGetDataSourceUpdateDataSourceDeleteDataSourceExecuteDataSourceAdministratePermissions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDataSourcePermissionDataSourceCreateDataSourceGetDataSourceUpdateDataSourceDeleteDataSourceExecuteDataSourceAdministratePermissionsWithDefaults() *DataSourcePermissionDataSourceCreateDataSourceGetDataSourceUpdateDataSourceDeleteDataSourceExecuteDataSourceAdministratePermissions {
	this := DataSourcePermissionDataSourceCreateDataSourceGetDataSourceUpdateDataSourceDeleteDataSourceExecuteDataSourceAdministratePermissions{}
	return &this
}

// GetOwnerId returns the OwnerId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DataSourcePermissionDataSourceCreateDataSourceGetDataSourceUpdateDataSourceDeleteDataSourceExecuteDataSourceAdministratePermissions) GetOwnerId() string {
	if o == nil || o.OwnerId.Get() == nil {
		var ret string
		return ret
	}
	return *o.OwnerId.Get()
}

// GetOwnerIdOk returns a tuple with the OwnerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DataSourcePermissionDataSourceCreateDataSourceGetDataSourceUpdateDataSourceDeleteDataSourceExecuteDataSourceAdministratePermissions) GetOwnerIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.OwnerId.Get(), o.OwnerId.IsSet()
}

// HasOwnerId returns a boolean if a field has been set.
func (o *DataSourcePermissionDataSourceCreateDataSourceGetDataSourceUpdateDataSourceDeleteDataSourceExecuteDataSourceAdministratePermissions) HasOwnerId() bool {
	if o != nil && o.OwnerId.IsSet() {
		return true
	}

	return false
}

// SetOwnerId gets a reference to the given NullableString and assigns it to the OwnerId field.
func (o *DataSourcePermissionDataSourceCreateDataSourceGetDataSourceUpdateDataSourceDeleteDataSourceExecuteDataSourceAdministratePermissions) SetOwnerId(v string) {
	o.OwnerId.Set(&v)
}
// SetOwnerIdNil sets the value for OwnerId to be an explicit nil
func (o *DataSourcePermissionDataSourceCreateDataSourceGetDataSourceUpdateDataSourceDeleteDataSourceExecuteDataSourceAdministratePermissions) SetOwnerIdNil() {
	o.OwnerId.Set(nil)
}

// UnsetOwnerId ensures that no value is present for OwnerId, not even an explicit nil
func (o *DataSourcePermissionDataSourceCreateDataSourceGetDataSourceUpdateDataSourceDeleteDataSourceExecuteDataSourceAdministratePermissions) UnsetOwnerId() {
	o.OwnerId.Unset()
}

// GetOwner returns the Owner field value if set, zero value otherwise.
func (o *DataSourcePermissionDataSourceCreateDataSourceGetDataSourceUpdateDataSourceDeleteDataSourceExecuteDataSourceAdministratePermissions) GetOwner() DataSourcePermission {
	if o == nil || o.Owner == nil {
		var ret DataSourcePermission
		return ret
	}
	return *o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataSourcePermissionDataSourceCreateDataSourceGetDataSourceUpdateDataSourceDeleteDataSourceExecuteDataSourceAdministratePermissions) GetOwnerOk() (*DataSourcePermission, bool) {
	if o == nil || o.Owner == nil {
		return nil, false
	}
	return o.Owner, true
}

// HasOwner returns a boolean if a field has been set.
func (o *DataSourcePermissionDataSourceCreateDataSourceGetDataSourceUpdateDataSourceDeleteDataSourceExecuteDataSourceAdministratePermissions) HasOwner() bool {
	if o != nil && o.Owner != nil {
		return true
	}

	return false
}

// SetOwner gets a reference to the given DataSourcePermission and assigns it to the Owner field.
func (o *DataSourcePermissionDataSourceCreateDataSourceGetDataSourceUpdateDataSourceDeleteDataSourceExecuteDataSourceAdministratePermissions) SetOwner(v DataSourcePermission) {
	o.Owner = &v
}

// GetGroups returns the Groups field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DataSourcePermissionDataSourceCreateDataSourceGetDataSourceUpdateDataSourceDeleteDataSourceExecuteDataSourceAdministratePermissions) GetGroups() map[string]DataSourcePermission {
	if o == nil  {
		var ret map[string]DataSourcePermission
		return ret
	}
	return o.Groups
}

// GetGroupsOk returns a tuple with the Groups field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DataSourcePermissionDataSourceCreateDataSourceGetDataSourceUpdateDataSourceDeleteDataSourceExecuteDataSourceAdministratePermissions) GetGroupsOk() (*map[string]DataSourcePermission, bool) {
	if o == nil || o.Groups == nil {
		return nil, false
	}
	return &o.Groups, true
}

// HasGroups returns a boolean if a field has been set.
func (o *DataSourcePermissionDataSourceCreateDataSourceGetDataSourceUpdateDataSourceDeleteDataSourceExecuteDataSourceAdministratePermissions) HasGroups() bool {
	if o != nil && o.Groups != nil {
		return true
	}

	return false
}

// SetGroups gets a reference to the given map[string]DataSourcePermission and assigns it to the Groups field.
func (o *DataSourcePermissionDataSourceCreateDataSourceGetDataSourceUpdateDataSourceDeleteDataSourceExecuteDataSourceAdministratePermissions) SetGroups(v map[string]DataSourcePermission) {
	o.Groups = v
}

// GetOther returns the Other field value if set, zero value otherwise.
func (o *DataSourcePermissionDataSourceCreateDataSourceGetDataSourceUpdateDataSourceDeleteDataSourceExecuteDataSourceAdministratePermissions) GetOther() DataSourcePermission {
	if o == nil || o.Other == nil {
		var ret DataSourcePermission
		return ret
	}
	return *o.Other
}

// GetOtherOk returns a tuple with the Other field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataSourcePermissionDataSourceCreateDataSourceGetDataSourceUpdateDataSourceDeleteDataSourceExecuteDataSourceAdministratePermissions) GetOtherOk() (*DataSourcePermission, bool) {
	if o == nil || o.Other == nil {
		return nil, false
	}
	return o.Other, true
}

// HasOther returns a boolean if a field has been set.
func (o *DataSourcePermissionDataSourceCreateDataSourceGetDataSourceUpdateDataSourceDeleteDataSourceExecuteDataSourceAdministratePermissions) HasOther() bool {
	if o != nil && o.Other != nil {
		return true
	}

	return false
}

// SetOther gets a reference to the given DataSourcePermission and assigns it to the Other field.
func (o *DataSourcePermissionDataSourceCreateDataSourceGetDataSourceUpdateDataSourceDeleteDataSourceExecuteDataSourceAdministratePermissions) SetOther(v DataSourcePermission) {
	o.Other = &v
}

// GetAnon returns the Anon field value if set, zero value otherwise.
func (o *DataSourcePermissionDataSourceCreateDataSourceGetDataSourceUpdateDataSourceDeleteDataSourceExecuteDataSourceAdministratePermissions) GetAnon() DataSourcePermission {
	if o == nil || o.Anon == nil {
		var ret DataSourcePermission
		return ret
	}
	return *o.Anon
}

// GetAnonOk returns a tuple with the Anon field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataSourcePermissionDataSourceCreateDataSourceGetDataSourceUpdateDataSourceDeleteDataSourceExecuteDataSourceAdministratePermissions) GetAnonOk() (*DataSourcePermission, bool) {
	if o == nil || o.Anon == nil {
		return nil, false
	}
	return o.Anon, true
}

// HasAnon returns a boolean if a field has been set.
func (o *DataSourcePermissionDataSourceCreateDataSourceGetDataSourceUpdateDataSourceDeleteDataSourceExecuteDataSourceAdministratePermissions) HasAnon() bool {
	if o != nil && o.Anon != nil {
		return true
	}

	return false
}

// SetAnon gets a reference to the given DataSourcePermission and assigns it to the Anon field.
func (o *DataSourcePermissionDataSourceCreateDataSourceGetDataSourceUpdateDataSourceDeleteDataSourceExecuteDataSourceAdministratePermissions) SetAnon(v DataSourcePermission) {
	o.Anon = &v
}

func (o DataSourcePermissionDataSourceCreateDataSourceGetDataSourceUpdateDataSourceDeleteDataSourceExecuteDataSourceAdministratePermissions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.OwnerId.IsSet() {
		toSerialize["ownerId"] = o.OwnerId.Get()
	}
	if o.Owner != nil {
		toSerialize["owner"] = o.Owner
	}
	if o.Groups != nil {
		toSerialize["groups"] = o.Groups
	}
	if o.Other != nil {
		toSerialize["other"] = o.Other
	}
	if o.Anon != nil {
		toSerialize["anon"] = o.Anon
	}
	return json.Marshal(toSerialize)
}

type NullableDataSourcePermissionDataSourceCreateDataSourceGetDataSourceUpdateDataSourceDeleteDataSourceExecuteDataSourceAdministratePermissions struct {
	value *DataSourcePermissionDataSourceCreateDataSourceGetDataSourceUpdateDataSourceDeleteDataSourceExecuteDataSourceAdministratePermissions
	isSet bool
}

func (v NullableDataSourcePermissionDataSourceCreateDataSourceGetDataSourceUpdateDataSourceDeleteDataSourceExecuteDataSourceAdministratePermissions) Get() *DataSourcePermissionDataSourceCreateDataSourceGetDataSourceUpdateDataSourceDeleteDataSourceExecuteDataSourceAdministratePermissions {
	return v.value
}

func (v *NullableDataSourcePermissionDataSourceCreateDataSourceGetDataSourceUpdateDataSourceDeleteDataSourceExecuteDataSourceAdministratePermissions) Set(val *DataSourcePermissionDataSourceCreateDataSourceGetDataSourceUpdateDataSourceDeleteDataSourceExecuteDataSourceAdministratePermissions) {
	v.value = val
	v.isSet = true
}

func (v NullableDataSourcePermissionDataSourceCreateDataSourceGetDataSourceUpdateDataSourceDeleteDataSourceExecuteDataSourceAdministratePermissions) IsSet() bool {
	return v.isSet
}

func (v *NullableDataSourcePermissionDataSourceCreateDataSourceGetDataSourceUpdateDataSourceDeleteDataSourceExecuteDataSourceAdministratePermissions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDataSourcePermissionDataSourceCreateDataSourceGetDataSourceUpdateDataSourceDeleteDataSourceExecuteDataSourceAdministratePermissions(val *DataSourcePermissionDataSourceCreateDataSourceGetDataSourceUpdateDataSourceDeleteDataSourceExecuteDataSourceAdministratePermissions) *NullableDataSourcePermissionDataSourceCreateDataSourceGetDataSourceUpdateDataSourceDeleteDataSourceExecuteDataSourceAdministratePermissions {
	return &NullableDataSourcePermissionDataSourceCreateDataSourceGetDataSourceUpdateDataSourceDeleteDataSourceExecuteDataSourceAdministratePermissions{value: val, isSet: true}
}

func (v NullableDataSourcePermissionDataSourceCreateDataSourceGetDataSourceUpdateDataSourceDeleteDataSourceExecuteDataSourceAdministratePermissions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDataSourcePermissionDataSourceCreateDataSourceGetDataSourceUpdateDataSourceDeleteDataSourceExecuteDataSourceAdministratePermissions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


