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

// checks if the FetchTaskVM type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FetchTaskVM{}

// FetchTaskVM struct for FetchTaskVM
type FetchTaskVM struct {
	TaskBaseVM
	DataSourceId NullableString `json:"dataSourceId,omitempty"`
	T string `json:"$t"`
}

type _FetchTaskVM FetchTaskVM

// NewFetchTaskVM instantiates a new FetchTaskVM object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFetchTaskVM(t string) *FetchTaskVM {
	this := FetchTaskVM{}
	this.T = t
	return &this
}

// NewFetchTaskVMWithDefaults instantiates a new FetchTaskVM object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFetchTaskVMWithDefaults() *FetchTaskVM {
	this := FetchTaskVM{}
	return &this
}

// GetDataSourceId returns the DataSourceId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FetchTaskVM) GetDataSourceId() string {
	if o == nil || IsNil(o.DataSourceId.Get()) {
		var ret string
		return ret
	}
	return *o.DataSourceId.Get()
}

// GetDataSourceIdOk returns a tuple with the DataSourceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FetchTaskVM) GetDataSourceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DataSourceId.Get(), o.DataSourceId.IsSet()
}

// HasDataSourceId returns a boolean if a field has been set.
func (o *FetchTaskVM) HasDataSourceId() bool {
	if o != nil && o.DataSourceId.IsSet() {
		return true
	}

	return false
}

// SetDataSourceId gets a reference to the given NullableString and assigns it to the DataSourceId field.
func (o *FetchTaskVM) SetDataSourceId(v string) {
	o.DataSourceId.Set(&v)
}
// SetDataSourceIdNil sets the value for DataSourceId to be an explicit nil
func (o *FetchTaskVM) SetDataSourceIdNil() {
	o.DataSourceId.Set(nil)
}

// UnsetDataSourceId ensures that no value is present for DataSourceId, not even an explicit nil
func (o *FetchTaskVM) UnsetDataSourceId() {
	o.DataSourceId.Unset()
}

// GetT returns the T field value
func (o *FetchTaskVM) GetT() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.T
}

// GetTOk returns a tuple with the T field value
// and a boolean to check if the value has been set.
func (o *FetchTaskVM) GetTOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.T, true
}

// SetT sets field value
func (o *FetchTaskVM) SetT(v string) {
	o.T = v
}

func (o FetchTaskVM) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FetchTaskVM) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedTaskBaseVM, errTaskBaseVM := json.Marshal(o.TaskBaseVM)
	if errTaskBaseVM != nil {
		return map[string]interface{}{}, errTaskBaseVM
	}
	errTaskBaseVM = json.Unmarshal([]byte(serializedTaskBaseVM), &toSerialize)
	if errTaskBaseVM != nil {
		return map[string]interface{}{}, errTaskBaseVM
	}
	if o.DataSourceId.IsSet() {
		toSerialize["dataSourceId"] = o.DataSourceId.Get()
	}
	toSerialize["$t"] = o.T
	return toSerialize, nil
}

func (o *FetchTaskVM) UnmarshalJSON(data []byte) (err error) {
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

	varFetchTaskVM := _FetchTaskVM{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varFetchTaskVM)

	if err != nil {
		return err
	}

	*o = FetchTaskVM(varFetchTaskVM)

	return err
}

type NullableFetchTaskVM struct {
	value *FetchTaskVM
	isSet bool
}

func (v NullableFetchTaskVM) Get() *FetchTaskVM {
	return v.value
}

func (v *NullableFetchTaskVM) Set(val *FetchTaskVM) {
	v.value = val
	v.isSet = true
}

func (v NullableFetchTaskVM) IsSet() bool {
	return v.isSet
}

func (v *NullableFetchTaskVM) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFetchTaskVM(val *FetchTaskVM) *NullableFetchTaskVM {
	return &NullableFetchTaskVM{value: val, isSet: true}
}

func (v NullableFetchTaskVM) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFetchTaskVM) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


