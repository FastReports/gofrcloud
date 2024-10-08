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

// checks if the RunFetchTaskVM type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RunFetchTaskVM{}

// RunFetchTaskVM struct for RunFetchTaskVM
type RunFetchTaskVM struct {
	RunTaskBaseVM
	DataSourceId NullableString `json:"dataSourceId,omitempty"`
	T string `json:"$t"`
}

type _RunFetchTaskVM RunFetchTaskVM

// NewRunFetchTaskVM instantiates a new RunFetchTaskVM object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRunFetchTaskVM(t string) *RunFetchTaskVM {
	this := RunFetchTaskVM{}
	this.T = t
	return &this
}

// NewRunFetchTaskVMWithDefaults instantiates a new RunFetchTaskVM object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRunFetchTaskVMWithDefaults() *RunFetchTaskVM {
	this := RunFetchTaskVM{}
	return &this
}

// GetDataSourceId returns the DataSourceId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RunFetchTaskVM) GetDataSourceId() string {
	if o == nil || IsNil(o.DataSourceId.Get()) {
		var ret string
		return ret
	}
	return *o.DataSourceId.Get()
}

// GetDataSourceIdOk returns a tuple with the DataSourceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RunFetchTaskVM) GetDataSourceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DataSourceId.Get(), o.DataSourceId.IsSet()
}

// HasDataSourceId returns a boolean if a field has been set.
func (o *RunFetchTaskVM) HasDataSourceId() bool {
	if o != nil && o.DataSourceId.IsSet() {
		return true
	}

	return false
}

// SetDataSourceId gets a reference to the given NullableString and assigns it to the DataSourceId field.
func (o *RunFetchTaskVM) SetDataSourceId(v string) {
	o.DataSourceId.Set(&v)
}
// SetDataSourceIdNil sets the value for DataSourceId to be an explicit nil
func (o *RunFetchTaskVM) SetDataSourceIdNil() {
	o.DataSourceId.Set(nil)
}

// UnsetDataSourceId ensures that no value is present for DataSourceId, not even an explicit nil
func (o *RunFetchTaskVM) UnsetDataSourceId() {
	o.DataSourceId.Unset()
}

// GetT returns the T field value
func (o *RunFetchTaskVM) GetT() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.T
}

// GetTOk returns a tuple with the T field value
// and a boolean to check if the value has been set.
func (o *RunFetchTaskVM) GetTOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.T, true
}

// SetT sets field value
func (o *RunFetchTaskVM) SetT(v string) {
	o.T = v
}

func (o RunFetchTaskVM) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RunFetchTaskVM) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedRunTaskBaseVM, errRunTaskBaseVM := json.Marshal(o.RunTaskBaseVM)
	if errRunTaskBaseVM != nil {
		return map[string]interface{}{}, errRunTaskBaseVM
	}
	errRunTaskBaseVM = json.Unmarshal([]byte(serializedRunTaskBaseVM), &toSerialize)
	if errRunTaskBaseVM != nil {
		return map[string]interface{}{}, errRunTaskBaseVM
	}
	if o.DataSourceId.IsSet() {
		toSerialize["dataSourceId"] = o.DataSourceId.Get()
	}
	toSerialize["$t"] = o.T
	return toSerialize, nil
}

func (o *RunFetchTaskVM) UnmarshalJSON(data []byte) (err error) {
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

	varRunFetchTaskVM := _RunFetchTaskVM{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRunFetchTaskVM)

	if err != nil {
		return err
	}

	*o = RunFetchTaskVM(varRunFetchTaskVM)

	return err
}

type NullableRunFetchTaskVM struct {
	value *RunFetchTaskVM
	isSet bool
}

func (v NullableRunFetchTaskVM) Get() *RunFetchTaskVM {
	return v.value
}

func (v *NullableRunFetchTaskVM) Set(val *RunFetchTaskVM) {
	v.value = val
	v.isSet = true
}

func (v NullableRunFetchTaskVM) IsSet() bool {
	return v.isSet
}

func (v *NullableRunFetchTaskVM) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRunFetchTaskVM(val *RunFetchTaskVM) *NullableRunFetchTaskVM {
	return &NullableRunFetchTaskVM{value: val, isSet: true}
}

func (v NullableRunFetchTaskVM) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRunFetchTaskVM) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


