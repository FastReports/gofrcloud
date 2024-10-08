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

// checks if the RenameSubscriptionVM type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RenameSubscriptionVM{}

// RenameSubscriptionVM struct for RenameSubscriptionVM
type RenameSubscriptionVM struct {
	CloudBaseVM
	Name string `json:"name"`
	T string `json:"$t"`
}

type _RenameSubscriptionVM RenameSubscriptionVM

// NewRenameSubscriptionVM instantiates a new RenameSubscriptionVM object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRenameSubscriptionVM(name string, t string) *RenameSubscriptionVM {
	this := RenameSubscriptionVM{}
	this.T = t
	return &this
}

// NewRenameSubscriptionVMWithDefaults instantiates a new RenameSubscriptionVM object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRenameSubscriptionVMWithDefaults() *RenameSubscriptionVM {
	this := RenameSubscriptionVM{}
	return &this
}

// GetName returns the Name field value
func (o *RenameSubscriptionVM) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *RenameSubscriptionVM) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *RenameSubscriptionVM) SetName(v string) {
	o.Name = v
}

// GetT returns the T field value
func (o *RenameSubscriptionVM) GetT() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.T
}

// GetTOk returns a tuple with the T field value
// and a boolean to check if the value has been set.
func (o *RenameSubscriptionVM) GetTOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.T, true
}

// SetT sets field value
func (o *RenameSubscriptionVM) SetT(v string) {
	o.T = v
}

func (o RenameSubscriptionVM) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RenameSubscriptionVM) ToMap() (map[string]interface{}, error) {
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
	toSerialize["$t"] = o.T
	return toSerialize, nil
}

func (o *RenameSubscriptionVM) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
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

	varRenameSubscriptionVM := _RenameSubscriptionVM{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRenameSubscriptionVM)

	if err != nil {
		return err
	}

	*o = RenameSubscriptionVM(varRenameSubscriptionVM)

	return err
}

type NullableRenameSubscriptionVM struct {
	value *RenameSubscriptionVM
	isSet bool
}

func (v NullableRenameSubscriptionVM) Get() *RenameSubscriptionVM {
	return v.value
}

func (v *NullableRenameSubscriptionVM) Set(val *RenameSubscriptionVM) {
	v.value = val
	v.isSet = true
}

func (v NullableRenameSubscriptionVM) IsSet() bool {
	return v.isSet
}

func (v *NullableRenameSubscriptionVM) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRenameSubscriptionVM(val *RenameSubscriptionVM) *NullableRenameSubscriptionVM {
	return &NullableRenameSubscriptionVM{value: val, isSet: true}
}

func (v NullableRenameSubscriptionVM) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRenameSubscriptionVM) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


