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

// checks if the DataSourceSelectCommandVM type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DataSourceSelectCommandVM{}

// DataSourceSelectCommandVM struct for DataSourceSelectCommandVM
type DataSourceSelectCommandVM struct {
	CloudBaseVM
	Name string `json:"name"`
	Command string `json:"command"`
	Parameters []DataSourceSelectCommandParameterVM `json:"parameters,omitempty"`
	T string `json:"$t"`
}

type _DataSourceSelectCommandVM DataSourceSelectCommandVM

// NewDataSourceSelectCommandVM instantiates a new DataSourceSelectCommandVM object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDataSourceSelectCommandVM(name string, command string, t string) *DataSourceSelectCommandVM {
	this := DataSourceSelectCommandVM{}
	this.T = t
	return &this
}

// NewDataSourceSelectCommandVMWithDefaults instantiates a new DataSourceSelectCommandVM object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDataSourceSelectCommandVMWithDefaults() *DataSourceSelectCommandVM {
	this := DataSourceSelectCommandVM{}
	return &this
}

// GetName returns the Name field value
func (o *DataSourceSelectCommandVM) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DataSourceSelectCommandVM) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *DataSourceSelectCommandVM) SetName(v string) {
	o.Name = v
}

// GetCommand returns the Command field value
func (o *DataSourceSelectCommandVM) GetCommand() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Command
}

// GetCommandOk returns a tuple with the Command field value
// and a boolean to check if the value has been set.
func (o *DataSourceSelectCommandVM) GetCommandOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Command, true
}

// SetCommand sets field value
func (o *DataSourceSelectCommandVM) SetCommand(v string) {
	o.Command = v
}

// GetParameters returns the Parameters field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DataSourceSelectCommandVM) GetParameters() []DataSourceSelectCommandParameterVM {
	if o == nil {
		var ret []DataSourceSelectCommandParameterVM
		return ret
	}
	return o.Parameters
}

// GetParametersOk returns a tuple with the Parameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DataSourceSelectCommandVM) GetParametersOk() ([]DataSourceSelectCommandParameterVM, bool) {
	if o == nil || IsNil(o.Parameters) {
		return nil, false
	}
	return o.Parameters, true
}

// HasParameters returns a boolean if a field has been set.
func (o *DataSourceSelectCommandVM) HasParameters() bool {
	if o != nil && IsNil(o.Parameters) {
		return true
	}

	return false
}

// SetParameters gets a reference to the given []DataSourceSelectCommandParameterVM and assigns it to the Parameters field.
func (o *DataSourceSelectCommandVM) SetParameters(v []DataSourceSelectCommandParameterVM) {
	o.Parameters = v
}

// GetT returns the T field value
func (o *DataSourceSelectCommandVM) GetT() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.T
}

// GetTOk returns a tuple with the T field value
// and a boolean to check if the value has been set.
func (o *DataSourceSelectCommandVM) GetTOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.T, true
}

// SetT sets field value
func (o *DataSourceSelectCommandVM) SetT(v string) {
	o.T = v
}

func (o DataSourceSelectCommandVM) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DataSourceSelectCommandVM) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedCloudBaseVM, errCloudBaseVM := json.Marshal(o.CloudBaseVM)
	if errCloudBaseVM != nil {
		return map[string]interface{}{}, errCloudBaseVM
	}
	errCloudBaseVM = json.Unmarshal([]byte(serializedCloudBaseVM), &toSerialize)
	if errCloudBaseVM != nil {
		return map[string]interface{}{}, errCloudBaseVM
	}
	toSerialize["name"] = o.Name
	toSerialize["command"] = o.Command
	if o.Parameters != nil {
		toSerialize["parameters"] = o.Parameters
	}
	toSerialize["$t"] = o.T
	return toSerialize, nil
}

func (o *DataSourceSelectCommandVM) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"command",
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

	varDataSourceSelectCommandVM := _DataSourceSelectCommandVM{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDataSourceSelectCommandVM)

	if err != nil {
		return err
	}

	*o = DataSourceSelectCommandVM(varDataSourceSelectCommandVM)

	return err
}

type NullableDataSourceSelectCommandVM struct {
	value *DataSourceSelectCommandVM
	isSet bool
}

func (v NullableDataSourceSelectCommandVM) Get() *DataSourceSelectCommandVM {
	return v.value
}

func (v *NullableDataSourceSelectCommandVM) Set(val *DataSourceSelectCommandVM) {
	v.value = val
	v.isSet = true
}

func (v NullableDataSourceSelectCommandVM) IsSet() bool {
	return v.isSet
}

func (v *NullableDataSourceSelectCommandVM) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDataSourceSelectCommandVM(val *DataSourceSelectCommandVM) *NullableDataSourceSelectCommandVM {
	return &NullableDataSourceSelectCommandVM{value: val, isSet: true}
}

func (v NullableDataSourceSelectCommandVM) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDataSourceSelectCommandVM) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


