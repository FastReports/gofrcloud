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

// GroupCreateGroupGetGroupUpdateGroupDeleteGroupExecuteGroupAdministratePermission struct for GroupCreateGroupGetGroupUpdateGroupDeleteGroupExecuteGroupAdministratePermission
type GroupCreateGroupGetGroupUpdateGroupDeleteGroupExecuteGroupAdministratePermission struct {
	Create *GroupCreate `json:"create,omitempty"`
	Delete *GroupDelete `json:"delete,omitempty"`
	Execute *GroupExecute `json:"execute,omitempty"`
	Get *GroupGet `json:"get,omitempty"`
	Update *GroupUpdate `json:"update,omitempty"`
	Administrate *GroupAdministrate `json:"administrate,omitempty"`
}

// NewGroupCreateGroupGetGroupUpdateGroupDeleteGroupExecuteGroupAdministratePermission instantiates a new GroupCreateGroupGetGroupUpdateGroupDeleteGroupExecuteGroupAdministratePermission object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupCreateGroupGetGroupUpdateGroupDeleteGroupExecuteGroupAdministratePermission() *GroupCreateGroupGetGroupUpdateGroupDeleteGroupExecuteGroupAdministratePermission {
	this := GroupCreateGroupGetGroupUpdateGroupDeleteGroupExecuteGroupAdministratePermission{}
	return &this
}

// NewGroupCreateGroupGetGroupUpdateGroupDeleteGroupExecuteGroupAdministratePermissionWithDefaults instantiates a new GroupCreateGroupGetGroupUpdateGroupDeleteGroupExecuteGroupAdministratePermission object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupCreateGroupGetGroupUpdateGroupDeleteGroupExecuteGroupAdministratePermissionWithDefaults() *GroupCreateGroupGetGroupUpdateGroupDeleteGroupExecuteGroupAdministratePermission {
	this := GroupCreateGroupGetGroupUpdateGroupDeleteGroupExecuteGroupAdministratePermission{}
	return &this
}

// GetCreate returns the Create field value if set, zero value otherwise.
func (o *GroupCreateGroupGetGroupUpdateGroupDeleteGroupExecuteGroupAdministratePermission) GetCreate() GroupCreate {
	if o == nil || o.Create == nil {
		var ret GroupCreate
		return ret
	}
	return *o.Create
}

// GetCreateOk returns a tuple with the Create field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupCreateGroupGetGroupUpdateGroupDeleteGroupExecuteGroupAdministratePermission) GetCreateOk() (*GroupCreate, bool) {
	if o == nil || o.Create == nil {
		return nil, false
	}
	return o.Create, true
}

// HasCreate returns a boolean if a field has been set.
func (o *GroupCreateGroupGetGroupUpdateGroupDeleteGroupExecuteGroupAdministratePermission) HasCreate() bool {
	if o != nil && o.Create != nil {
		return true
	}

	return false
}

// SetCreate gets a reference to the given GroupCreate and assigns it to the Create field.
func (o *GroupCreateGroupGetGroupUpdateGroupDeleteGroupExecuteGroupAdministratePermission) SetCreate(v GroupCreate) {
	o.Create = &v
}

// GetDelete returns the Delete field value if set, zero value otherwise.
func (o *GroupCreateGroupGetGroupUpdateGroupDeleteGroupExecuteGroupAdministratePermission) GetDelete() GroupDelete {
	if o == nil || o.Delete == nil {
		var ret GroupDelete
		return ret
	}
	return *o.Delete
}

// GetDeleteOk returns a tuple with the Delete field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupCreateGroupGetGroupUpdateGroupDeleteGroupExecuteGroupAdministratePermission) GetDeleteOk() (*GroupDelete, bool) {
	if o == nil || o.Delete == nil {
		return nil, false
	}
	return o.Delete, true
}

// HasDelete returns a boolean if a field has been set.
func (o *GroupCreateGroupGetGroupUpdateGroupDeleteGroupExecuteGroupAdministratePermission) HasDelete() bool {
	if o != nil && o.Delete != nil {
		return true
	}

	return false
}

// SetDelete gets a reference to the given GroupDelete and assigns it to the Delete field.
func (o *GroupCreateGroupGetGroupUpdateGroupDeleteGroupExecuteGroupAdministratePermission) SetDelete(v GroupDelete) {
	o.Delete = &v
}

// GetExecute returns the Execute field value if set, zero value otherwise.
func (o *GroupCreateGroupGetGroupUpdateGroupDeleteGroupExecuteGroupAdministratePermission) GetExecute() GroupExecute {
	if o == nil || o.Execute == nil {
		var ret GroupExecute
		return ret
	}
	return *o.Execute
}

// GetExecuteOk returns a tuple with the Execute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupCreateGroupGetGroupUpdateGroupDeleteGroupExecuteGroupAdministratePermission) GetExecuteOk() (*GroupExecute, bool) {
	if o == nil || o.Execute == nil {
		return nil, false
	}
	return o.Execute, true
}

// HasExecute returns a boolean if a field has been set.
func (o *GroupCreateGroupGetGroupUpdateGroupDeleteGroupExecuteGroupAdministratePermission) HasExecute() bool {
	if o != nil && o.Execute != nil {
		return true
	}

	return false
}

// SetExecute gets a reference to the given GroupExecute and assigns it to the Execute field.
func (o *GroupCreateGroupGetGroupUpdateGroupDeleteGroupExecuteGroupAdministratePermission) SetExecute(v GroupExecute) {
	o.Execute = &v
}

// GetGet returns the Get field value if set, zero value otherwise.
func (o *GroupCreateGroupGetGroupUpdateGroupDeleteGroupExecuteGroupAdministratePermission) GetGet() GroupGet {
	if o == nil || o.Get == nil {
		var ret GroupGet
		return ret
	}
	return *o.Get
}

// GetGetOk returns a tuple with the Get field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupCreateGroupGetGroupUpdateGroupDeleteGroupExecuteGroupAdministratePermission) GetGetOk() (*GroupGet, bool) {
	if o == nil || o.Get == nil {
		return nil, false
	}
	return o.Get, true
}

// HasGet returns a boolean if a field has been set.
func (o *GroupCreateGroupGetGroupUpdateGroupDeleteGroupExecuteGroupAdministratePermission) HasGet() bool {
	if o != nil && o.Get != nil {
		return true
	}

	return false
}

// SetGet gets a reference to the given GroupGet and assigns it to the Get field.
func (o *GroupCreateGroupGetGroupUpdateGroupDeleteGroupExecuteGroupAdministratePermission) SetGet(v GroupGet) {
	o.Get = &v
}

// GetUpdate returns the Update field value if set, zero value otherwise.
func (o *GroupCreateGroupGetGroupUpdateGroupDeleteGroupExecuteGroupAdministratePermission) GetUpdate() GroupUpdate {
	if o == nil || o.Update == nil {
		var ret GroupUpdate
		return ret
	}
	return *o.Update
}

// GetUpdateOk returns a tuple with the Update field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupCreateGroupGetGroupUpdateGroupDeleteGroupExecuteGroupAdministratePermission) GetUpdateOk() (*GroupUpdate, bool) {
	if o == nil || o.Update == nil {
		return nil, false
	}
	return o.Update, true
}

// HasUpdate returns a boolean if a field has been set.
func (o *GroupCreateGroupGetGroupUpdateGroupDeleteGroupExecuteGroupAdministratePermission) HasUpdate() bool {
	if o != nil && o.Update != nil {
		return true
	}

	return false
}

// SetUpdate gets a reference to the given GroupUpdate and assigns it to the Update field.
func (o *GroupCreateGroupGetGroupUpdateGroupDeleteGroupExecuteGroupAdministratePermission) SetUpdate(v GroupUpdate) {
	o.Update = &v
}

// GetAdministrate returns the Administrate field value if set, zero value otherwise.
func (o *GroupCreateGroupGetGroupUpdateGroupDeleteGroupExecuteGroupAdministratePermission) GetAdministrate() GroupAdministrate {
	if o == nil || o.Administrate == nil {
		var ret GroupAdministrate
		return ret
	}
	return *o.Administrate
}

// GetAdministrateOk returns a tuple with the Administrate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupCreateGroupGetGroupUpdateGroupDeleteGroupExecuteGroupAdministratePermission) GetAdministrateOk() (*GroupAdministrate, bool) {
	if o == nil || o.Administrate == nil {
		return nil, false
	}
	return o.Administrate, true
}

// HasAdministrate returns a boolean if a field has been set.
func (o *GroupCreateGroupGetGroupUpdateGroupDeleteGroupExecuteGroupAdministratePermission) HasAdministrate() bool {
	if o != nil && o.Administrate != nil {
		return true
	}

	return false
}

// SetAdministrate gets a reference to the given GroupAdministrate and assigns it to the Administrate field.
func (o *GroupCreateGroupGetGroupUpdateGroupDeleteGroupExecuteGroupAdministratePermission) SetAdministrate(v GroupAdministrate) {
	o.Administrate = &v
}

func (o GroupCreateGroupGetGroupUpdateGroupDeleteGroupExecuteGroupAdministratePermission) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Create != nil {
		toSerialize["create"] = o.Create
	}
	if o.Delete != nil {
		toSerialize["delete"] = o.Delete
	}
	if o.Execute != nil {
		toSerialize["execute"] = o.Execute
	}
	if o.Get != nil {
		toSerialize["get"] = o.Get
	}
	if o.Update != nil {
		toSerialize["update"] = o.Update
	}
	if o.Administrate != nil {
		toSerialize["administrate"] = o.Administrate
	}
	return json.Marshal(toSerialize)
}

type NullableGroupCreateGroupGetGroupUpdateGroupDeleteGroupExecuteGroupAdministratePermission struct {
	value *GroupCreateGroupGetGroupUpdateGroupDeleteGroupExecuteGroupAdministratePermission
	isSet bool
}

func (v NullableGroupCreateGroupGetGroupUpdateGroupDeleteGroupExecuteGroupAdministratePermission) Get() *GroupCreateGroupGetGroupUpdateGroupDeleteGroupExecuteGroupAdministratePermission {
	return v.value
}

func (v *NullableGroupCreateGroupGetGroupUpdateGroupDeleteGroupExecuteGroupAdministratePermission) Set(val *GroupCreateGroupGetGroupUpdateGroupDeleteGroupExecuteGroupAdministratePermission) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupCreateGroupGetGroupUpdateGroupDeleteGroupExecuteGroupAdministratePermission) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupCreateGroupGetGroupUpdateGroupDeleteGroupExecuteGroupAdministratePermission) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupCreateGroupGetGroupUpdateGroupDeleteGroupExecuteGroupAdministratePermission(val *GroupCreateGroupGetGroupUpdateGroupDeleteGroupExecuteGroupAdministratePermission) *NullableGroupCreateGroupGetGroupUpdateGroupDeleteGroupExecuteGroupAdministratePermission {
	return &NullableGroupCreateGroupGetGroupUpdateGroupDeleteGroupExecuteGroupAdministratePermission{value: val, isSet: true}
}

func (v NullableGroupCreateGroupGetGroupUpdateGroupDeleteGroupExecuteGroupAdministratePermission) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupCreateGroupGetGroupUpdateGroupDeleteGroupExecuteGroupAdministratePermission) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


