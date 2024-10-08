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

// checks if the RestOfSpaceVM type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RestOfSpaceVM{}

// RestOfSpaceVM struct for RestOfSpaceVM
type RestOfSpaceVM struct {
	CloudBaseVM
	RestOfSpace *int64 `json:"restOfSpace,omitempty"`
	SubscriptionPlan *SubscriptionPlanVM `json:"subscriptionPlan,omitempty"`
	T string `json:"$t"`
}

type _RestOfSpaceVM RestOfSpaceVM

// NewRestOfSpaceVM instantiates a new RestOfSpaceVM object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRestOfSpaceVM(t string) *RestOfSpaceVM {
	this := RestOfSpaceVM{}
	this.T = t
	return &this
}

// NewRestOfSpaceVMWithDefaults instantiates a new RestOfSpaceVM object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRestOfSpaceVMWithDefaults() *RestOfSpaceVM {
	this := RestOfSpaceVM{}
	return &this
}

// GetRestOfSpace returns the RestOfSpace field value if set, zero value otherwise.
func (o *RestOfSpaceVM) GetRestOfSpace() int64 {
	if o == nil || IsNil(o.RestOfSpace) {
		var ret int64
		return ret
	}
	return *o.RestOfSpace
}

// GetRestOfSpaceOk returns a tuple with the RestOfSpace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestOfSpaceVM) GetRestOfSpaceOk() (*int64, bool) {
	if o == nil || IsNil(o.RestOfSpace) {
		return nil, false
	}
	return o.RestOfSpace, true
}

// HasRestOfSpace returns a boolean if a field has been set.
func (o *RestOfSpaceVM) HasRestOfSpace() bool {
	if o != nil && !IsNil(o.RestOfSpace) {
		return true
	}

	return false
}

// SetRestOfSpace gets a reference to the given int64 and assigns it to the RestOfSpace field.
func (o *RestOfSpaceVM) SetRestOfSpace(v int64) {
	o.RestOfSpace = &v
}

// GetSubscriptionPlan returns the SubscriptionPlan field value if set, zero value otherwise.
func (o *RestOfSpaceVM) GetSubscriptionPlan() SubscriptionPlanVM {
	if o == nil || IsNil(o.SubscriptionPlan) {
		var ret SubscriptionPlanVM
		return ret
	}
	return *o.SubscriptionPlan
}

// GetSubscriptionPlanOk returns a tuple with the SubscriptionPlan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestOfSpaceVM) GetSubscriptionPlanOk() (*SubscriptionPlanVM, bool) {
	if o == nil || IsNil(o.SubscriptionPlan) {
		return nil, false
	}
	return o.SubscriptionPlan, true
}

// HasSubscriptionPlan returns a boolean if a field has been set.
func (o *RestOfSpaceVM) HasSubscriptionPlan() bool {
	if o != nil && !IsNil(o.SubscriptionPlan) {
		return true
	}

	return false
}

// SetSubscriptionPlan gets a reference to the given SubscriptionPlanVM and assigns it to the SubscriptionPlan field.
func (o *RestOfSpaceVM) SetSubscriptionPlan(v SubscriptionPlanVM) {
	o.SubscriptionPlan = &v
}

// GetT returns the T field value
func (o *RestOfSpaceVM) GetT() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.T
}

// GetTOk returns a tuple with the T field value
// and a boolean to check if the value has been set.
func (o *RestOfSpaceVM) GetTOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.T, true
}

// SetT sets field value
func (o *RestOfSpaceVM) SetT(v string) {
	o.T = v
}

func (o RestOfSpaceVM) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RestOfSpaceVM) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedCloudBaseVM, errCloudBaseVM := json.Marshal(o.CloudBaseVM)
	if errCloudBaseVM != nil {
		return map[string]interface{}{}, errCloudBaseVM
	}
	errCloudBaseVM = json.Unmarshal([]byte(serializedCloudBaseVM), &toSerialize)
	if errCloudBaseVM != nil {
		return map[string]interface{}{}, errCloudBaseVM
	}
	if !IsNil(o.RestOfSpace) {
		toSerialize["restOfSpace"] = o.RestOfSpace
	}
	if !IsNil(o.SubscriptionPlan) {
		toSerialize["subscriptionPlan"] = o.SubscriptionPlan
	}
	toSerialize["$t"] = o.T
	return toSerialize, nil
}

func (o *RestOfSpaceVM) UnmarshalJSON(data []byte) (err error) {
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

	varRestOfSpaceVM := _RestOfSpaceVM{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRestOfSpaceVM)

	if err != nil {
		return err
	}

	*o = RestOfSpaceVM(varRestOfSpaceVM)

	return err
}

type NullableRestOfSpaceVM struct {
	value *RestOfSpaceVM
	isSet bool
}

func (v NullableRestOfSpaceVM) Get() *RestOfSpaceVM {
	return v.value
}

func (v *NullableRestOfSpaceVM) Set(val *RestOfSpaceVM) {
	v.value = val
	v.isSet = true
}

func (v NullableRestOfSpaceVM) IsSet() bool {
	return v.isSet
}

func (v *NullableRestOfSpaceVM) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRestOfSpaceVM(val *RestOfSpaceVM) *NullableRestOfSpaceVM {
	return &NullableRestOfSpaceVM{value: val, isSet: true}
}

func (v NullableRestOfSpaceVM) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRestOfSpaceVM) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


