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

// checks if the CreateFetchTaskVM type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateFetchTaskVM{}

// CreateFetchTaskVM struct for CreateFetchTaskVM
type CreateFetchTaskVM struct {
	CreateTaskBaseVM
	DataSourceId NullableString `json:"dataSourceId,omitempty"`
	T string `json:"$t"`
}

type _CreateFetchTaskVM CreateFetchTaskVM

// NewCreateFetchTaskVM instantiates a new CreateFetchTaskVM object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateFetchTaskVM(t string) *CreateFetchTaskVM {
	this := CreateFetchTaskVM{}
	this.T = t
	return &this
}

// NewCreateFetchTaskVMWithDefaults instantiates a new CreateFetchTaskVM object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateFetchTaskVMWithDefaults() *CreateFetchTaskVM {
	this := CreateFetchTaskVM{}
	return &this
}

// GetDataSourceId returns the DataSourceId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateFetchTaskVM) GetDataSourceId() string {
	if o == nil || IsNil(o.DataSourceId.Get()) {
		var ret string
		return ret
	}
	return *o.DataSourceId.Get()
}

// GetDataSourceIdOk returns a tuple with the DataSourceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateFetchTaskVM) GetDataSourceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DataSourceId.Get(), o.DataSourceId.IsSet()
}

// HasDataSourceId returns a boolean if a field has been set.
func (o *CreateFetchTaskVM) HasDataSourceId() bool {
	if o != nil && o.DataSourceId.IsSet() {
		return true
	}

	return false
}

// SetDataSourceId gets a reference to the given NullableString and assigns it to the DataSourceId field.
func (o *CreateFetchTaskVM) SetDataSourceId(v string) {
	o.DataSourceId.Set(&v)
}
// SetDataSourceIdNil sets the value for DataSourceId to be an explicit nil
func (o *CreateFetchTaskVM) SetDataSourceIdNil() {
	o.DataSourceId.Set(nil)
}

// UnsetDataSourceId ensures that no value is present for DataSourceId, not even an explicit nil
func (o *CreateFetchTaskVM) UnsetDataSourceId() {
	o.DataSourceId.Unset()
}

// GetT returns the T field value
func (o *CreateFetchTaskVM) GetT() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.T
}

// GetTOk returns a tuple with the T field value
// and a boolean to check if the value has been set.
func (o *CreateFetchTaskVM) GetTOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.T, true
}

// SetT sets field value
func (o *CreateFetchTaskVM) SetT(v string) {
	o.T = v
}

func (o CreateFetchTaskVM) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateFetchTaskVM) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedCreateTaskBaseVM, errCreateTaskBaseVM := json.Marshal(o.CreateTaskBaseVM)
	if errCreateTaskBaseVM != nil {
		return map[string]interface{}{}, errCreateTaskBaseVM
	}
	errCreateTaskBaseVM = json.Unmarshal([]byte(serializedCreateTaskBaseVM), &toSerialize)
	if errCreateTaskBaseVM != nil {
		return map[string]interface{}{}, errCreateTaskBaseVM
	}
	if o.DataSourceId.IsSet() {
		toSerialize["dataSourceId"] = o.DataSourceId.Get()
	}
	toSerialize["$t"] = o.T
	return toSerialize, nil
}

func (o *CreateFetchTaskVM) UnmarshalJSON(data []byte) (err error) {
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

	varCreateFetchTaskVM := _CreateFetchTaskVM{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateFetchTaskVM)

	if err != nil {
		return err
	}

	*o = CreateFetchTaskVM(varCreateFetchTaskVM)

	return err
}

type NullableCreateFetchTaskVM struct {
	value *CreateFetchTaskVM
	isSet bool
}

func (v NullableCreateFetchTaskVM) Get() *CreateFetchTaskVM {
	return v.value
}

func (v *NullableCreateFetchTaskVM) Set(val *CreateFetchTaskVM) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateFetchTaskVM) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateFetchTaskVM) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateFetchTaskVM(val *CreateFetchTaskVM) *NullableCreateFetchTaskVM {
	return &NullableCreateFetchTaskVM{value: val, isSet: true}
}

func (v NullableCreateFetchTaskVM) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateFetchTaskVM) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


