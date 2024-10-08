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

// checks if the TaskPermissionCRUDVM type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TaskPermissionCRUDVM{}

// TaskPermissionCRUDVM struct for TaskPermissionCRUDVM
type TaskPermissionCRUDVM struct {
	CloudBaseVM
	Create *TaskCreate `json:"create,omitempty"`
	Delete *TaskDelete `json:"delete,omitempty"`
	Execute *TaskExecute `json:"execute,omitempty"`
	Get *TaskGet `json:"get,omitempty"`
	Update *TaskUpdate `json:"update,omitempty"`
	Administrate *TaskAdministrate `json:"administrate,omitempty"`
	T string `json:"$t"`
}

type _TaskPermissionCRUDVM TaskPermissionCRUDVM

// NewTaskPermissionCRUDVM instantiates a new TaskPermissionCRUDVM object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTaskPermissionCRUDVM(t string) *TaskPermissionCRUDVM {
	this := TaskPermissionCRUDVM{}
	this.T = t
	return &this
}

// NewTaskPermissionCRUDVMWithDefaults instantiates a new TaskPermissionCRUDVM object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTaskPermissionCRUDVMWithDefaults() *TaskPermissionCRUDVM {
	this := TaskPermissionCRUDVM{}
	return &this
}

// GetCreate returns the Create field value if set, zero value otherwise.
func (o *TaskPermissionCRUDVM) GetCreate() TaskCreate {
	if o == nil || IsNil(o.Create) {
		var ret TaskCreate
		return ret
	}
	return *o.Create
}

// GetCreateOk returns a tuple with the Create field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskPermissionCRUDVM) GetCreateOk() (*TaskCreate, bool) {
	if o == nil || IsNil(o.Create) {
		return nil, false
	}
	return o.Create, true
}

// HasCreate returns a boolean if a field has been set.
func (o *TaskPermissionCRUDVM) HasCreate() bool {
	if o != nil && !IsNil(o.Create) {
		return true
	}

	return false
}

// SetCreate gets a reference to the given TaskCreate and assigns it to the Create field.
func (o *TaskPermissionCRUDVM) SetCreate(v TaskCreate) {
	o.Create = &v
}

// GetDelete returns the Delete field value if set, zero value otherwise.
func (o *TaskPermissionCRUDVM) GetDelete() TaskDelete {
	if o == nil || IsNil(o.Delete) {
		var ret TaskDelete
		return ret
	}
	return *o.Delete
}

// GetDeleteOk returns a tuple with the Delete field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskPermissionCRUDVM) GetDeleteOk() (*TaskDelete, bool) {
	if o == nil || IsNil(o.Delete) {
		return nil, false
	}
	return o.Delete, true
}

// HasDelete returns a boolean if a field has been set.
func (o *TaskPermissionCRUDVM) HasDelete() bool {
	if o != nil && !IsNil(o.Delete) {
		return true
	}

	return false
}

// SetDelete gets a reference to the given TaskDelete and assigns it to the Delete field.
func (o *TaskPermissionCRUDVM) SetDelete(v TaskDelete) {
	o.Delete = &v
}

// GetExecute returns the Execute field value if set, zero value otherwise.
func (o *TaskPermissionCRUDVM) GetExecute() TaskExecute {
	if o == nil || IsNil(o.Execute) {
		var ret TaskExecute
		return ret
	}
	return *o.Execute
}

// GetExecuteOk returns a tuple with the Execute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskPermissionCRUDVM) GetExecuteOk() (*TaskExecute, bool) {
	if o == nil || IsNil(o.Execute) {
		return nil, false
	}
	return o.Execute, true
}

// HasExecute returns a boolean if a field has been set.
func (o *TaskPermissionCRUDVM) HasExecute() bool {
	if o != nil && !IsNil(o.Execute) {
		return true
	}

	return false
}

// SetExecute gets a reference to the given TaskExecute and assigns it to the Execute field.
func (o *TaskPermissionCRUDVM) SetExecute(v TaskExecute) {
	o.Execute = &v
}

// GetGet returns the Get field value if set, zero value otherwise.
func (o *TaskPermissionCRUDVM) GetGet() TaskGet {
	if o == nil || IsNil(o.Get) {
		var ret TaskGet
		return ret
	}
	return *o.Get
}

// GetGetOk returns a tuple with the Get field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskPermissionCRUDVM) GetGetOk() (*TaskGet, bool) {
	if o == nil || IsNil(o.Get) {
		return nil, false
	}
	return o.Get, true
}

// HasGet returns a boolean if a field has been set.
func (o *TaskPermissionCRUDVM) HasGet() bool {
	if o != nil && !IsNil(o.Get) {
		return true
	}

	return false
}

// SetGet gets a reference to the given TaskGet and assigns it to the Get field.
func (o *TaskPermissionCRUDVM) SetGet(v TaskGet) {
	o.Get = &v
}

// GetUpdate returns the Update field value if set, zero value otherwise.
func (o *TaskPermissionCRUDVM) GetUpdate() TaskUpdate {
	if o == nil || IsNil(o.Update) {
		var ret TaskUpdate
		return ret
	}
	return *o.Update
}

// GetUpdateOk returns a tuple with the Update field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskPermissionCRUDVM) GetUpdateOk() (*TaskUpdate, bool) {
	if o == nil || IsNil(o.Update) {
		return nil, false
	}
	return o.Update, true
}

// HasUpdate returns a boolean if a field has been set.
func (o *TaskPermissionCRUDVM) HasUpdate() bool {
	if o != nil && !IsNil(o.Update) {
		return true
	}

	return false
}

// SetUpdate gets a reference to the given TaskUpdate and assigns it to the Update field.
func (o *TaskPermissionCRUDVM) SetUpdate(v TaskUpdate) {
	o.Update = &v
}

// GetAdministrate returns the Administrate field value if set, zero value otherwise.
func (o *TaskPermissionCRUDVM) GetAdministrate() TaskAdministrate {
	if o == nil || IsNil(o.Administrate) {
		var ret TaskAdministrate
		return ret
	}
	return *o.Administrate
}

// GetAdministrateOk returns a tuple with the Administrate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskPermissionCRUDVM) GetAdministrateOk() (*TaskAdministrate, bool) {
	if o == nil || IsNil(o.Administrate) {
		return nil, false
	}
	return o.Administrate, true
}

// HasAdministrate returns a boolean if a field has been set.
func (o *TaskPermissionCRUDVM) HasAdministrate() bool {
	if o != nil && !IsNil(o.Administrate) {
		return true
	}

	return false
}

// SetAdministrate gets a reference to the given TaskAdministrate and assigns it to the Administrate field.
func (o *TaskPermissionCRUDVM) SetAdministrate(v TaskAdministrate) {
	o.Administrate = &v
}

// GetT returns the T field value
func (o *TaskPermissionCRUDVM) GetT() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.T
}

// GetTOk returns a tuple with the T field value
// and a boolean to check if the value has been set.
func (o *TaskPermissionCRUDVM) GetTOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.T, true
}

// SetT sets field value
func (o *TaskPermissionCRUDVM) SetT(v string) {
	o.T = v
}

func (o TaskPermissionCRUDVM) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TaskPermissionCRUDVM) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedCloudBaseVM, errCloudBaseVM := json.Marshal(o.CloudBaseVM)
	if errCloudBaseVM != nil {
		return map[string]interface{}{}, errCloudBaseVM
	}
	errCloudBaseVM = json.Unmarshal([]byte(serializedCloudBaseVM), &toSerialize)
	if errCloudBaseVM != nil {
		return map[string]interface{}{}, errCloudBaseVM
	}
	if !IsNil(o.Create) {
		toSerialize["create"] = o.Create
	}
	if !IsNil(o.Delete) {
		toSerialize["delete"] = o.Delete
	}
	if !IsNil(o.Execute) {
		toSerialize["execute"] = o.Execute
	}
	if !IsNil(o.Get) {
		toSerialize["get"] = o.Get
	}
	if !IsNil(o.Update) {
		toSerialize["update"] = o.Update
	}
	if !IsNil(o.Administrate) {
		toSerialize["administrate"] = o.Administrate
	}
	toSerialize["$t"] = o.T
	return toSerialize, nil
}

func (o *TaskPermissionCRUDVM) UnmarshalJSON(data []byte) (err error) {
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

	varTaskPermissionCRUDVM := _TaskPermissionCRUDVM{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTaskPermissionCRUDVM)

	if err != nil {
		return err
	}

	*o = TaskPermissionCRUDVM(varTaskPermissionCRUDVM)

	return err
}

type NullableTaskPermissionCRUDVM struct {
	value *TaskPermissionCRUDVM
	isSet bool
}

func (v NullableTaskPermissionCRUDVM) Get() *TaskPermissionCRUDVM {
	return v.value
}

func (v *NullableTaskPermissionCRUDVM) Set(val *TaskPermissionCRUDVM) {
	v.value = val
	v.isSet = true
}

func (v NullableTaskPermissionCRUDVM) IsSet() bool {
	return v.isSet
}

func (v *NullableTaskPermissionCRUDVM) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTaskPermissionCRUDVM(val *TaskPermissionCRUDVM) *NullableTaskPermissionCRUDVM {
	return &NullableTaskPermissionCRUDVM{value: val, isSet: true}
}

func (v NullableTaskPermissionCRUDVM) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTaskPermissionCRUDVM) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


