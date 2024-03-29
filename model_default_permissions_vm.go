/*
FastReport Cloud

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gofrcloud

import (
	"encoding/json"
)

// checks if the DefaultPermissionsVM type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DefaultPermissionsVM{}

// DefaultPermissionsVM struct for DefaultPermissionsVM
type DefaultPermissionsVM struct {
	FilePermissions *FilePermissions `json:"filePermissions,omitempty"`
	DataSourcePermissions *DataSourcePermissions `json:"dataSourcePermissions,omitempty"`
	GroupPermissions *GroupPermissions `json:"groupPermissions,omitempty"`
	TaskPermissions *TaskPermissions `json:"taskPermissions,omitempty"`
}

// NewDefaultPermissionsVM instantiates a new DefaultPermissionsVM object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDefaultPermissionsVM() *DefaultPermissionsVM {
	this := DefaultPermissionsVM{}
	return &this
}

// NewDefaultPermissionsVMWithDefaults instantiates a new DefaultPermissionsVM object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDefaultPermissionsVMWithDefaults() *DefaultPermissionsVM {
	this := DefaultPermissionsVM{}
	return &this
}

// GetFilePermissions returns the FilePermissions field value if set, zero value otherwise.
func (o *DefaultPermissionsVM) GetFilePermissions() FilePermissions {
	if o == nil || IsNil(o.FilePermissions) {
		var ret FilePermissions
		return ret
	}
	return *o.FilePermissions
}

// GetFilePermissionsOk returns a tuple with the FilePermissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DefaultPermissionsVM) GetFilePermissionsOk() (*FilePermissions, bool) {
	if o == nil || IsNil(o.FilePermissions) {
		return nil, false
	}
	return o.FilePermissions, true
}

// HasFilePermissions returns a boolean if a field has been set.
func (o *DefaultPermissionsVM) HasFilePermissions() bool {
	if o != nil && !IsNil(o.FilePermissions) {
		return true
	}

	return false
}

// SetFilePermissions gets a reference to the given FilePermissions and assigns it to the FilePermissions field.
func (o *DefaultPermissionsVM) SetFilePermissions(v FilePermissions) {
	o.FilePermissions = &v
}

// GetDataSourcePermissions returns the DataSourcePermissions field value if set, zero value otherwise.
func (o *DefaultPermissionsVM) GetDataSourcePermissions() DataSourcePermissions {
	if o == nil || IsNil(o.DataSourcePermissions) {
		var ret DataSourcePermissions
		return ret
	}
	return *o.DataSourcePermissions
}

// GetDataSourcePermissionsOk returns a tuple with the DataSourcePermissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DefaultPermissionsVM) GetDataSourcePermissionsOk() (*DataSourcePermissions, bool) {
	if o == nil || IsNil(o.DataSourcePermissions) {
		return nil, false
	}
	return o.DataSourcePermissions, true
}

// HasDataSourcePermissions returns a boolean if a field has been set.
func (o *DefaultPermissionsVM) HasDataSourcePermissions() bool {
	if o != nil && !IsNil(o.DataSourcePermissions) {
		return true
	}

	return false
}

// SetDataSourcePermissions gets a reference to the given DataSourcePermissions and assigns it to the DataSourcePermissions field.
func (o *DefaultPermissionsVM) SetDataSourcePermissions(v DataSourcePermissions) {
	o.DataSourcePermissions = &v
}

// GetGroupPermissions returns the GroupPermissions field value if set, zero value otherwise.
func (o *DefaultPermissionsVM) GetGroupPermissions() GroupPermissions {
	if o == nil || IsNil(o.GroupPermissions) {
		var ret GroupPermissions
		return ret
	}
	return *o.GroupPermissions
}

// GetGroupPermissionsOk returns a tuple with the GroupPermissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DefaultPermissionsVM) GetGroupPermissionsOk() (*GroupPermissions, bool) {
	if o == nil || IsNil(o.GroupPermissions) {
		return nil, false
	}
	return o.GroupPermissions, true
}

// HasGroupPermissions returns a boolean if a field has been set.
func (o *DefaultPermissionsVM) HasGroupPermissions() bool {
	if o != nil && !IsNil(o.GroupPermissions) {
		return true
	}

	return false
}

// SetGroupPermissions gets a reference to the given GroupPermissions and assigns it to the GroupPermissions field.
func (o *DefaultPermissionsVM) SetGroupPermissions(v GroupPermissions) {
	o.GroupPermissions = &v
}

// GetTaskPermissions returns the TaskPermissions field value if set, zero value otherwise.
func (o *DefaultPermissionsVM) GetTaskPermissions() TaskPermissions {
	if o == nil || IsNil(o.TaskPermissions) {
		var ret TaskPermissions
		return ret
	}
	return *o.TaskPermissions
}

// GetTaskPermissionsOk returns a tuple with the TaskPermissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DefaultPermissionsVM) GetTaskPermissionsOk() (*TaskPermissions, bool) {
	if o == nil || IsNil(o.TaskPermissions) {
		return nil, false
	}
	return o.TaskPermissions, true
}

// HasTaskPermissions returns a boolean if a field has been set.
func (o *DefaultPermissionsVM) HasTaskPermissions() bool {
	if o != nil && !IsNil(o.TaskPermissions) {
		return true
	}

	return false
}

// SetTaskPermissions gets a reference to the given TaskPermissions and assigns it to the TaskPermissions field.
func (o *DefaultPermissionsVM) SetTaskPermissions(v TaskPermissions) {
	o.TaskPermissions = &v
}

func (o DefaultPermissionsVM) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DefaultPermissionsVM) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FilePermissions) {
		toSerialize["filePermissions"] = o.FilePermissions
	}
	if !IsNil(o.DataSourcePermissions) {
		toSerialize["dataSourcePermissions"] = o.DataSourcePermissions
	}
	if !IsNil(o.GroupPermissions) {
		toSerialize["groupPermissions"] = o.GroupPermissions
	}
	if !IsNil(o.TaskPermissions) {
		toSerialize["taskPermissions"] = o.TaskPermissions
	}
	return toSerialize, nil
}

type NullableDefaultPermissionsVM struct {
	value *DefaultPermissionsVM
	isSet bool
}

func (v NullableDefaultPermissionsVM) Get() *DefaultPermissionsVM {
	return v.value
}

func (v *NullableDefaultPermissionsVM) Set(val *DefaultPermissionsVM) {
	v.value = val
	v.isSet = true
}

func (v NullableDefaultPermissionsVM) IsSet() bool {
	return v.isSet
}

func (v *NullableDefaultPermissionsVM) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDefaultPermissionsVM(val *DefaultPermissionsVM) *NullableDefaultPermissionsVM {
	return &NullableDefaultPermissionsVM{value: val, isSet: true}
}

func (v NullableDefaultPermissionsVM) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDefaultPermissionsVM) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


