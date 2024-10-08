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

// checks if the CreateApiKeyVM type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateApiKeyVM{}

// CreateApiKeyVM struct for CreateApiKeyVM
type CreateApiKeyVM struct {
	CloudBaseVM
	Description NullableString `json:"description,omitempty"`
	Expired time.Time `json:"expired"`
	T string `json:"$t"`
}

type _CreateApiKeyVM CreateApiKeyVM

// NewCreateApiKeyVM instantiates a new CreateApiKeyVM object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateApiKeyVM(expired time.Time, t string) *CreateApiKeyVM {
	this := CreateApiKeyVM{}
	this.T = t
	return &this
}

// NewCreateApiKeyVMWithDefaults instantiates a new CreateApiKeyVM object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateApiKeyVMWithDefaults() *CreateApiKeyVM {
	this := CreateApiKeyVM{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateApiKeyVM) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateApiKeyVM) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *CreateApiKeyVM) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *CreateApiKeyVM) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *CreateApiKeyVM) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *CreateApiKeyVM) UnsetDescription() {
	o.Description.Unset()
}

// GetExpired returns the Expired field value
func (o *CreateApiKeyVM) GetExpired() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Expired
}

// GetExpiredOk returns a tuple with the Expired field value
// and a boolean to check if the value has been set.
func (o *CreateApiKeyVM) GetExpiredOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Expired, true
}

// SetExpired sets field value
func (o *CreateApiKeyVM) SetExpired(v time.Time) {
	o.Expired = v
}

// GetT returns the T field value
func (o *CreateApiKeyVM) GetT() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.T
}

// GetTOk returns a tuple with the T field value
// and a boolean to check if the value has been set.
func (o *CreateApiKeyVM) GetTOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.T, true
}

// SetT sets field value
func (o *CreateApiKeyVM) SetT(v string) {
	o.T = v
}

func (o CreateApiKeyVM) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateApiKeyVM) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedCloudBaseVM, errCloudBaseVM := json.Marshal(o.CloudBaseVM)
	if errCloudBaseVM != nil {
		return map[string]interface{}{}, errCloudBaseVM
	}
	errCloudBaseVM = json.Unmarshal([]byte(serializedCloudBaseVM), &toSerialize)
	if errCloudBaseVM != nil {
		return map[string]interface{}{}, errCloudBaseVM
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	toSerialize["expired"] = o.Expired
	toSerialize["$t"] = o.T
	return toSerialize, nil
}

func (o *CreateApiKeyVM) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"expired",
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

	varCreateApiKeyVM := _CreateApiKeyVM{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateApiKeyVM)

	if err != nil {
		return err
	}

	*o = CreateApiKeyVM(varCreateApiKeyVM)

	return err
}

type NullableCreateApiKeyVM struct {
	value *CreateApiKeyVM
	isSet bool
}

func (v NullableCreateApiKeyVM) Get() *CreateApiKeyVM {
	return v.value
}

func (v *NullableCreateApiKeyVM) Set(val *CreateApiKeyVM) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateApiKeyVM) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateApiKeyVM) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateApiKeyVM(val *CreateApiKeyVM) *NullableCreateApiKeyVM {
	return &NullableCreateApiKeyVM{value: val, isSet: true}
}

func (v NullableCreateApiKeyVM) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateApiKeyVM) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


