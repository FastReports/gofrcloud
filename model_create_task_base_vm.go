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

// checks if the CreateTaskBaseVM type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateTaskBaseVM{}

// CreateTaskBaseVM struct for CreateTaskBaseVM
type CreateTaskBaseVM struct {
	CloudBaseVM
	CronExpression NullableString `json:"cronExpression,omitempty"`
	StartsOn NullableTime `json:"startsOn,omitempty"`
	Ends *CreateTaskEndVM `json:"ends,omitempty"`
	Name NullableString `json:"name,omitempty"`
	SubscriptionId NullableString `json:"subscriptionId,omitempty"`
	T string `json:"$t"`
}

type _CreateTaskBaseVM CreateTaskBaseVM

// NewCreateTaskBaseVM instantiates a new CreateTaskBaseVM object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateTaskBaseVM(t string) *CreateTaskBaseVM {
	this := CreateTaskBaseVM{}
	this.T = t
	return &this
}

// NewCreateTaskBaseVMWithDefaults instantiates a new CreateTaskBaseVM object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateTaskBaseVMWithDefaults() *CreateTaskBaseVM {
	this := CreateTaskBaseVM{}
	return &this
}

// GetCronExpression returns the CronExpression field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateTaskBaseVM) GetCronExpression() string {
	if o == nil || IsNil(o.CronExpression.Get()) {
		var ret string
		return ret
	}
	return *o.CronExpression.Get()
}

// GetCronExpressionOk returns a tuple with the CronExpression field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateTaskBaseVM) GetCronExpressionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CronExpression.Get(), o.CronExpression.IsSet()
}

// HasCronExpression returns a boolean if a field has been set.
func (o *CreateTaskBaseVM) HasCronExpression() bool {
	if o != nil && o.CronExpression.IsSet() {
		return true
	}

	return false
}

// SetCronExpression gets a reference to the given NullableString and assigns it to the CronExpression field.
func (o *CreateTaskBaseVM) SetCronExpression(v string) {
	o.CronExpression.Set(&v)
}
// SetCronExpressionNil sets the value for CronExpression to be an explicit nil
func (o *CreateTaskBaseVM) SetCronExpressionNil() {
	o.CronExpression.Set(nil)
}

// UnsetCronExpression ensures that no value is present for CronExpression, not even an explicit nil
func (o *CreateTaskBaseVM) UnsetCronExpression() {
	o.CronExpression.Unset()
}

// GetStartsOn returns the StartsOn field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateTaskBaseVM) GetStartsOn() time.Time {
	if o == nil || IsNil(o.StartsOn.Get()) {
		var ret time.Time
		return ret
	}
	return *o.StartsOn.Get()
}

// GetStartsOnOk returns a tuple with the StartsOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateTaskBaseVM) GetStartsOnOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.StartsOn.Get(), o.StartsOn.IsSet()
}

// HasStartsOn returns a boolean if a field has been set.
func (o *CreateTaskBaseVM) HasStartsOn() bool {
	if o != nil && o.StartsOn.IsSet() {
		return true
	}

	return false
}

// SetStartsOn gets a reference to the given NullableTime and assigns it to the StartsOn field.
func (o *CreateTaskBaseVM) SetStartsOn(v time.Time) {
	o.StartsOn.Set(&v)
}
// SetStartsOnNil sets the value for StartsOn to be an explicit nil
func (o *CreateTaskBaseVM) SetStartsOnNil() {
	o.StartsOn.Set(nil)
}

// UnsetStartsOn ensures that no value is present for StartsOn, not even an explicit nil
func (o *CreateTaskBaseVM) UnsetStartsOn() {
	o.StartsOn.Unset()
}

// GetEnds returns the Ends field value if set, zero value otherwise.
func (o *CreateTaskBaseVM) GetEnds() CreateTaskEndVM {
	if o == nil || IsNil(o.Ends) {
		var ret CreateTaskEndVM
		return ret
	}
	return *o.Ends
}

// GetEndsOk returns a tuple with the Ends field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateTaskBaseVM) GetEndsOk() (*CreateTaskEndVM, bool) {
	if o == nil || IsNil(o.Ends) {
		return nil, false
	}
	return o.Ends, true
}

// HasEnds returns a boolean if a field has been set.
func (o *CreateTaskBaseVM) HasEnds() bool {
	if o != nil && !IsNil(o.Ends) {
		return true
	}

	return false
}

// SetEnds gets a reference to the given CreateTaskEndVM and assigns it to the Ends field.
func (o *CreateTaskBaseVM) SetEnds(v CreateTaskEndVM) {
	o.Ends = &v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateTaskBaseVM) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateTaskBaseVM) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *CreateTaskBaseVM) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *CreateTaskBaseVM) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *CreateTaskBaseVM) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *CreateTaskBaseVM) UnsetName() {
	o.Name.Unset()
}

// GetSubscriptionId returns the SubscriptionId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateTaskBaseVM) GetSubscriptionId() string {
	if o == nil || IsNil(o.SubscriptionId.Get()) {
		var ret string
		return ret
	}
	return *o.SubscriptionId.Get()
}

// GetSubscriptionIdOk returns a tuple with the SubscriptionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateTaskBaseVM) GetSubscriptionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SubscriptionId.Get(), o.SubscriptionId.IsSet()
}

// HasSubscriptionId returns a boolean if a field has been set.
func (o *CreateTaskBaseVM) HasSubscriptionId() bool {
	if o != nil && o.SubscriptionId.IsSet() {
		return true
	}

	return false
}

// SetSubscriptionId gets a reference to the given NullableString and assigns it to the SubscriptionId field.
func (o *CreateTaskBaseVM) SetSubscriptionId(v string) {
	o.SubscriptionId.Set(&v)
}
// SetSubscriptionIdNil sets the value for SubscriptionId to be an explicit nil
func (o *CreateTaskBaseVM) SetSubscriptionIdNil() {
	o.SubscriptionId.Set(nil)
}

// UnsetSubscriptionId ensures that no value is present for SubscriptionId, not even an explicit nil
func (o *CreateTaskBaseVM) UnsetSubscriptionId() {
	o.SubscriptionId.Unset()
}

// GetT returns the T field value
func (o *CreateTaskBaseVM) GetT() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.T
}

// GetTOk returns a tuple with the T field value
// and a boolean to check if the value has been set.
func (o *CreateTaskBaseVM) GetTOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.T, true
}

// SetT sets field value
func (o *CreateTaskBaseVM) SetT(v string) {
	o.T = v
}

func (o CreateTaskBaseVM) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateTaskBaseVM) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedCloudBaseVM, errCloudBaseVM := json.Marshal(o.CloudBaseVM)
	if errCloudBaseVM != nil {
		return map[string]interface{}{}, errCloudBaseVM
	}
	errCloudBaseVM = json.Unmarshal([]byte(serializedCloudBaseVM), &toSerialize)
	if errCloudBaseVM != nil {
		return map[string]interface{}{}, errCloudBaseVM
	}
	if o.CronExpression.IsSet() {
		toSerialize["cronExpression"] = o.CronExpression.Get()
	}
	if o.StartsOn.IsSet() {
		toSerialize["startsOn"] = o.StartsOn.Get()
	}
	if !IsNil(o.Ends) {
		toSerialize["ends"] = o.Ends
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.SubscriptionId.IsSet() {
		toSerialize["subscriptionId"] = o.SubscriptionId.Get()
	}
	toSerialize["$t"] = o.T
	return toSerialize, nil
}

func (o *CreateTaskBaseVM) UnmarshalJSON(data []byte) (err error) {
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

	varCreateTaskBaseVM := _CreateTaskBaseVM{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateTaskBaseVM)

	if err != nil {
		return err
	}

	*o = CreateTaskBaseVM(varCreateTaskBaseVM)

	return err
}

type NullableCreateTaskBaseVM struct {
	value *CreateTaskBaseVM
	isSet bool
}

func (v NullableCreateTaskBaseVM) Get() *CreateTaskBaseVM {
	return v.value
}

func (v *NullableCreateTaskBaseVM) Set(val *CreateTaskBaseVM) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateTaskBaseVM) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateTaskBaseVM) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateTaskBaseVM(val *CreateTaskBaseVM) *NullableCreateTaskBaseVM {
	return &NullableCreateTaskBaseVM{value: val, isSet: true}
}

func (v NullableCreateTaskBaseVM) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateTaskBaseVM) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


