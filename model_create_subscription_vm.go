/*
FastReport Cloud

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gofrcloud

import (
	"encoding/json"
	"time"
	"bytes"
	"fmt"
)

// checks if the CreateSubscriptionVM type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateSubscriptionVM{}

// CreateSubscriptionVM struct for CreateSubscriptionVM
type CreateSubscriptionVM struct {
	CloudBaseVM
	PlanId string `json:"planId"`
	Start NullableTime `json:"start,omitempty"`
	End NullableTime `json:"end,omitempty"`
	Name string `json:"name"`
	UserId string `json:"userId"`
	Tags []string `json:"tags,omitempty"`
	T string `json:"$t"`
}

type _CreateSubscriptionVM CreateSubscriptionVM

// NewCreateSubscriptionVM instantiates a new CreateSubscriptionVM object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateSubscriptionVM(planId string, name string, userId string, t string) *CreateSubscriptionVM {
	this := CreateSubscriptionVM{}
	this.T = t
	return &this
}

// NewCreateSubscriptionVMWithDefaults instantiates a new CreateSubscriptionVM object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateSubscriptionVMWithDefaults() *CreateSubscriptionVM {
	this := CreateSubscriptionVM{}
	return &this
}

// GetPlanId returns the PlanId field value
func (o *CreateSubscriptionVM) GetPlanId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PlanId
}

// GetPlanIdOk returns a tuple with the PlanId field value
// and a boolean to check if the value has been set.
func (o *CreateSubscriptionVM) GetPlanIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PlanId, true
}

// SetPlanId sets field value
func (o *CreateSubscriptionVM) SetPlanId(v string) {
	o.PlanId = v
}

// GetStart returns the Start field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateSubscriptionVM) GetStart() time.Time {
	if o == nil || IsNil(o.Start.Get()) {
		var ret time.Time
		return ret
	}
	return *o.Start.Get()
}

// GetStartOk returns a tuple with the Start field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateSubscriptionVM) GetStartOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Start.Get(), o.Start.IsSet()
}

// HasStart returns a boolean if a field has been set.
func (o *CreateSubscriptionVM) HasStart() bool {
	if o != nil && o.Start.IsSet() {
		return true
	}

	return false
}

// SetStart gets a reference to the given NullableTime and assigns it to the Start field.
func (o *CreateSubscriptionVM) SetStart(v time.Time) {
	o.Start.Set(&v)
}
// SetStartNil sets the value for Start to be an explicit nil
func (o *CreateSubscriptionVM) SetStartNil() {
	o.Start.Set(nil)
}

// UnsetStart ensures that no value is present for Start, not even an explicit nil
func (o *CreateSubscriptionVM) UnsetStart() {
	o.Start.Unset()
}

// GetEnd returns the End field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateSubscriptionVM) GetEnd() time.Time {
	if o == nil || IsNil(o.End.Get()) {
		var ret time.Time
		return ret
	}
	return *o.End.Get()
}

// GetEndOk returns a tuple with the End field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateSubscriptionVM) GetEndOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.End.Get(), o.End.IsSet()
}

// HasEnd returns a boolean if a field has been set.
func (o *CreateSubscriptionVM) HasEnd() bool {
	if o != nil && o.End.IsSet() {
		return true
	}

	return false
}

// SetEnd gets a reference to the given NullableTime and assigns it to the End field.
func (o *CreateSubscriptionVM) SetEnd(v time.Time) {
	o.End.Set(&v)
}
// SetEndNil sets the value for End to be an explicit nil
func (o *CreateSubscriptionVM) SetEndNil() {
	o.End.Set(nil)
}

// UnsetEnd ensures that no value is present for End, not even an explicit nil
func (o *CreateSubscriptionVM) UnsetEnd() {
	o.End.Unset()
}

// GetName returns the Name field value
func (o *CreateSubscriptionVM) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateSubscriptionVM) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateSubscriptionVM) SetName(v string) {
	o.Name = v
}

// GetUserId returns the UserId field value
func (o *CreateSubscriptionVM) GetUserId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value
// and a boolean to check if the value has been set.
func (o *CreateSubscriptionVM) GetUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserId, true
}

// SetUserId sets field value
func (o *CreateSubscriptionVM) SetUserId(v string) {
	o.UserId = v
}

// GetTags returns the Tags field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateSubscriptionVM) GetTags() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateSubscriptionVM) GetTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *CreateSubscriptionVM) HasTags() bool {
	if o != nil && IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *CreateSubscriptionVM) SetTags(v []string) {
	o.Tags = v
}

// GetT returns the T field value
func (o *CreateSubscriptionVM) GetT() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.T
}

// GetTOk returns a tuple with the T field value
// and a boolean to check if the value has been set.
func (o *CreateSubscriptionVM) GetTOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.T, true
}

// SetT sets field value
func (o *CreateSubscriptionVM) SetT(v string) {
	o.T = v
}

func (o CreateSubscriptionVM) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateSubscriptionVM) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedCloudBaseVM, errCloudBaseVM := json.Marshal(o.CloudBaseVM)
	if errCloudBaseVM != nil {
		return map[string]interface{}{}, errCloudBaseVM
	}
	errCloudBaseVM = json.Unmarshal([]byte(serializedCloudBaseVM), &toSerialize)
	if errCloudBaseVM != nil {
		return map[string]interface{}{}, errCloudBaseVM
	}
	toSerialize["planId"] = o.PlanId
	if o.Start.IsSet() {
		toSerialize["start"] = o.Start.Get()
	}
	if o.End.IsSet() {
		toSerialize["end"] = o.End.Get()
	}
	toSerialize["name"] = o.Name
	toSerialize["userId"] = o.UserId
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	toSerialize["$t"] = o.T
	return toSerialize, nil
}

func (o *CreateSubscriptionVM) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"planId",
		"name",
		"userId",
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

	varCreateSubscriptionVM := _CreateSubscriptionVM{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateSubscriptionVM)

	if err != nil {
		return err
	}

	*o = CreateSubscriptionVM(varCreateSubscriptionVM)

	return err
}

type NullableCreateSubscriptionVM struct {
	value *CreateSubscriptionVM
	isSet bool
}

func (v NullableCreateSubscriptionVM) Get() *CreateSubscriptionVM {
	return v.value
}

func (v *NullableCreateSubscriptionVM) Set(val *CreateSubscriptionVM) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateSubscriptionVM) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateSubscriptionVM) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateSubscriptionVM(val *CreateSubscriptionVM) *NullableCreateSubscriptionVM {
	return &NullableCreateSubscriptionVM{value: val, isSet: true}
}

func (v NullableCreateSubscriptionVM) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateSubscriptionVM) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


