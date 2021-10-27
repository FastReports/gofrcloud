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

// GroupPermissions struct for GroupPermissions
type GroupPermissions struct {
	OwnerId NullableString `json:"ownerId,omitempty"`
	Owner *GroupPermission `json:"owner,omitempty"`
	Groups map[string]GroupPermission `json:"groups,omitempty"`
	Other *GroupPermission `json:"other,omitempty"`
	Anon *GroupPermission `json:"anon,omitempty"`
}

// NewGroupPermissions instantiates a new GroupPermissions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupPermissions() *GroupPermissions {
	this := GroupPermissions{}
	return &this
}

// NewGroupPermissionsWithDefaults instantiates a new GroupPermissions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupPermissionsWithDefaults() *GroupPermissions {
	this := GroupPermissions{}
	return &this
}

// GetOwnerId returns the OwnerId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GroupPermissions) GetOwnerId() string {
	if o == nil || o.OwnerId.Get() == nil {
		var ret string
		return ret
	}
	return *o.OwnerId.Get()
}

// GetOwnerIdOk returns a tuple with the OwnerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GroupPermissions) GetOwnerIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.OwnerId.Get(), o.OwnerId.IsSet()
}

// HasOwnerId returns a boolean if a field has been set.
func (o *GroupPermissions) HasOwnerId() bool {
	if o != nil && o.OwnerId.IsSet() {
		return true
	}

	return false
}

// SetOwnerId gets a reference to the given NullableString and assigns it to the OwnerId field.
func (o *GroupPermissions) SetOwnerId(v string) {
	o.OwnerId.Set(&v)
}
// SetOwnerIdNil sets the value for OwnerId to be an explicit nil
func (o *GroupPermissions) SetOwnerIdNil() {
	o.OwnerId.Set(nil)
}

// UnsetOwnerId ensures that no value is present for OwnerId, not even an explicit nil
func (o *GroupPermissions) UnsetOwnerId() {
	o.OwnerId.Unset()
}

// GetOwner returns the Owner field value if set, zero value otherwise.
func (o *GroupPermissions) GetOwner() GroupPermission {
	if o == nil || o.Owner == nil {
		var ret GroupPermission
		return ret
	}
	return *o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupPermissions) GetOwnerOk() (*GroupPermission, bool) {
	if o == nil || o.Owner == nil {
		return nil, false
	}
	return o.Owner, true
}

// HasOwner returns a boolean if a field has been set.
func (o *GroupPermissions) HasOwner() bool {
	if o != nil && o.Owner != nil {
		return true
	}

	return false
}

// SetOwner gets a reference to the given GroupPermission and assigns it to the Owner field.
func (o *GroupPermissions) SetOwner(v GroupPermission) {
	o.Owner = &v
}

// GetGroups returns the Groups field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GroupPermissions) GetGroups() map[string]GroupPermission {
	if o == nil  {
		var ret map[string]GroupPermission
		return ret
	}
	return o.Groups
}

// GetGroupsOk returns a tuple with the Groups field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GroupPermissions) GetGroupsOk() (*map[string]GroupPermission, bool) {
	if o == nil || o.Groups == nil {
		return nil, false
	}
	return &o.Groups, true
}

// HasGroups returns a boolean if a field has been set.
func (o *GroupPermissions) HasGroups() bool {
	if o != nil && o.Groups != nil {
		return true
	}

	return false
}

// SetGroups gets a reference to the given map[string]GroupPermission and assigns it to the Groups field.
func (o *GroupPermissions) SetGroups(v map[string]GroupPermission) {
	o.Groups = v
}

// GetOther returns the Other field value if set, zero value otherwise.
func (o *GroupPermissions) GetOther() GroupPermission {
	if o == nil || o.Other == nil {
		var ret GroupPermission
		return ret
	}
	return *o.Other
}

// GetOtherOk returns a tuple with the Other field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupPermissions) GetOtherOk() (*GroupPermission, bool) {
	if o == nil || o.Other == nil {
		return nil, false
	}
	return o.Other, true
}

// HasOther returns a boolean if a field has been set.
func (o *GroupPermissions) HasOther() bool {
	if o != nil && o.Other != nil {
		return true
	}

	return false
}

// SetOther gets a reference to the given GroupPermission and assigns it to the Other field.
func (o *GroupPermissions) SetOther(v GroupPermission) {
	o.Other = &v
}

// GetAnon returns the Anon field value if set, zero value otherwise.
func (o *GroupPermissions) GetAnon() GroupPermission {
	if o == nil || o.Anon == nil {
		var ret GroupPermission
		return ret
	}
	return *o.Anon
}

// GetAnonOk returns a tuple with the Anon field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupPermissions) GetAnonOk() (*GroupPermission, bool) {
	if o == nil || o.Anon == nil {
		return nil, false
	}
	return o.Anon, true
}

// HasAnon returns a boolean if a field has been set.
func (o *GroupPermissions) HasAnon() bool {
	if o != nil && o.Anon != nil {
		return true
	}

	return false
}

// SetAnon gets a reference to the given GroupPermission and assigns it to the Anon field.
func (o *GroupPermissions) SetAnon(v GroupPermission) {
	o.Anon = &v
}

func (o GroupPermissions) MarshalJSON() ([]byte, error) {
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

type NullableGroupPermissions struct {
	value *GroupPermissions
	isSet bool
}

func (v NullableGroupPermissions) Get() *GroupPermissions {
	return v.value
}

func (v *NullableGroupPermissions) Set(val *GroupPermissions) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupPermissions) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupPermissions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupPermissions(val *GroupPermissions) *NullableGroupPermissions {
	return &NullableGroupPermissions{value: val, isSet: true}
}

func (v NullableGroupPermissions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupPermissions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


