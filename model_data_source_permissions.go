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

// DataSourcePermissions struct for DataSourcePermissions
type DataSourcePermissions struct {
	OwnerId NullableString `json:"ownerId,omitempty"`
	Owner *DataSourcePermission `json:"owner,omitempty"`
	Groups map[string]DataSourcePermission `json:"groups,omitempty"`
	Other *DataSourcePermission `json:"other,omitempty"`
	Anon *DataSourcePermission `json:"anon,omitempty"`
}

// NewDataSourcePermissions instantiates a new DataSourcePermissions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDataSourcePermissions() *DataSourcePermissions {
	this := DataSourcePermissions{}
	return &this
}

// NewDataSourcePermissionsWithDefaults instantiates a new DataSourcePermissions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDataSourcePermissionsWithDefaults() *DataSourcePermissions {
	this := DataSourcePermissions{}
	return &this
}

// GetOwnerId returns the OwnerId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DataSourcePermissions) GetOwnerId() string {
	if o == nil || o.OwnerId.Get() == nil {
		var ret string
		return ret
	}
	return *o.OwnerId.Get()
}

// GetOwnerIdOk returns a tuple with the OwnerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DataSourcePermissions) GetOwnerIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.OwnerId.Get(), o.OwnerId.IsSet()
}

// HasOwnerId returns a boolean if a field has been set.
func (o *DataSourcePermissions) HasOwnerId() bool {
	if o != nil && o.OwnerId.IsSet() {
		return true
	}

	return false
}

// SetOwnerId gets a reference to the given NullableString and assigns it to the OwnerId field.
func (o *DataSourcePermissions) SetOwnerId(v string) {
	o.OwnerId.Set(&v)
}
// SetOwnerIdNil sets the value for OwnerId to be an explicit nil
func (o *DataSourcePermissions) SetOwnerIdNil() {
	o.OwnerId.Set(nil)
}

// UnsetOwnerId ensures that no value is present for OwnerId, not even an explicit nil
func (o *DataSourcePermissions) UnsetOwnerId() {
	o.OwnerId.Unset()
}

// GetOwner returns the Owner field value if set, zero value otherwise.
func (o *DataSourcePermissions) GetOwner() DataSourcePermission {
	if o == nil || o.Owner == nil {
		var ret DataSourcePermission
		return ret
	}
	return *o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataSourcePermissions) GetOwnerOk() (*DataSourcePermission, bool) {
	if o == nil || o.Owner == nil {
		return nil, false
	}
	return o.Owner, true
}

// HasOwner returns a boolean if a field has been set.
func (o *DataSourcePermissions) HasOwner() bool {
	if o != nil && o.Owner != nil {
		return true
	}

	return false
}

// SetOwner gets a reference to the given DataSourcePermission and assigns it to the Owner field.
func (o *DataSourcePermissions) SetOwner(v DataSourcePermission) {
	o.Owner = &v
}

// GetGroups returns the Groups field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DataSourcePermissions) GetGroups() map[string]DataSourcePermission {
	if o == nil  {
		var ret map[string]DataSourcePermission
		return ret
	}
	return o.Groups
}

// GetGroupsOk returns a tuple with the Groups field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DataSourcePermissions) GetGroupsOk() (*map[string]DataSourcePermission, bool) {
	if o == nil || o.Groups == nil {
		return nil, false
	}
	return &o.Groups, true
}

// HasGroups returns a boolean if a field has been set.
func (o *DataSourcePermissions) HasGroups() bool {
	if o != nil && o.Groups != nil {
		return true
	}

	return false
}

// SetGroups gets a reference to the given map[string]DataSourcePermission and assigns it to the Groups field.
func (o *DataSourcePermissions) SetGroups(v map[string]DataSourcePermission) {
	o.Groups = v
}

// GetOther returns the Other field value if set, zero value otherwise.
func (o *DataSourcePermissions) GetOther() DataSourcePermission {
	if o == nil || o.Other == nil {
		var ret DataSourcePermission
		return ret
	}
	return *o.Other
}

// GetOtherOk returns a tuple with the Other field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataSourcePermissions) GetOtherOk() (*DataSourcePermission, bool) {
	if o == nil || o.Other == nil {
		return nil, false
	}
	return o.Other, true
}

// HasOther returns a boolean if a field has been set.
func (o *DataSourcePermissions) HasOther() bool {
	if o != nil && o.Other != nil {
		return true
	}

	return false
}

// SetOther gets a reference to the given DataSourcePermission and assigns it to the Other field.
func (o *DataSourcePermissions) SetOther(v DataSourcePermission) {
	o.Other = &v
}

// GetAnon returns the Anon field value if set, zero value otherwise.
func (o *DataSourcePermissions) GetAnon() DataSourcePermission {
	if o == nil || o.Anon == nil {
		var ret DataSourcePermission
		return ret
	}
	return *o.Anon
}

// GetAnonOk returns a tuple with the Anon field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataSourcePermissions) GetAnonOk() (*DataSourcePermission, bool) {
	if o == nil || o.Anon == nil {
		return nil, false
	}
	return o.Anon, true
}

// HasAnon returns a boolean if a field has been set.
func (o *DataSourcePermissions) HasAnon() bool {
	if o != nil && o.Anon != nil {
		return true
	}

	return false
}

// SetAnon gets a reference to the given DataSourcePermission and assigns it to the Anon field.
func (o *DataSourcePermissions) SetAnon(v DataSourcePermission) {
	o.Anon = &v
}

func (o DataSourcePermissions) MarshalJSON() ([]byte, error) {
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

type NullableDataSourcePermissions struct {
	value *DataSourcePermissions
	isSet bool
}

func (v NullableDataSourcePermissions) Get() *DataSourcePermissions {
	return v.value
}

func (v *NullableDataSourcePermissions) Set(val *DataSourcePermissions) {
	v.value = val
	v.isSet = true
}

func (v NullableDataSourcePermissions) IsSet() bool {
	return v.isSet
}

func (v *NullableDataSourcePermissions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDataSourcePermissions(val *DataSourcePermissions) *NullableDataSourcePermissions {
	return &NullableDataSourcePermissions{value: val, isSet: true}
}

func (v NullableDataSourcePermissions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDataSourcePermissions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


