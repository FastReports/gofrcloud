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

// checks if the AuditStatsVM type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuditStatsVM{}

// AuditStatsVM struct for AuditStatsVM
type AuditStatsVM struct {
	CloudBaseVM
	Items map[string][]AuditStatVM `json:"items,omitempty"`
	T string `json:"$t"`
}

type _AuditStatsVM AuditStatsVM

// NewAuditStatsVM instantiates a new AuditStatsVM object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuditStatsVM(t string) *AuditStatsVM {
	this := AuditStatsVM{}
	this.T = t
	return &this
}

// NewAuditStatsVMWithDefaults instantiates a new AuditStatsVM object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuditStatsVMWithDefaults() *AuditStatsVM {
	this := AuditStatsVM{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AuditStatsVM) GetItems() map[string][]AuditStatVM {
	if o == nil {
		var ret map[string][]AuditStatVM
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AuditStatsVM) GetItemsOk() (*map[string][]AuditStatVM, bool) {
	if o == nil || IsNil(o.Items) {
		return nil, false
	}
	return &o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *AuditStatsVM) HasItems() bool {
	if o != nil && IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given map[string][]AuditStatVM and assigns it to the Items field.
func (o *AuditStatsVM) SetItems(v map[string][]AuditStatVM) {
	o.Items = v
}

// GetT returns the T field value
func (o *AuditStatsVM) GetT() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.T
}

// GetTOk returns a tuple with the T field value
// and a boolean to check if the value has been set.
func (o *AuditStatsVM) GetTOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.T, true
}

// SetT sets field value
func (o *AuditStatsVM) SetT(v string) {
	o.T = v
}

func (o AuditStatsVM) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuditStatsVM) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedCloudBaseVM, errCloudBaseVM := json.Marshal(o.CloudBaseVM)
	if errCloudBaseVM != nil {
		return map[string]interface{}{}, errCloudBaseVM
	}
	errCloudBaseVM = json.Unmarshal([]byte(serializedCloudBaseVM), &toSerialize)
	if errCloudBaseVM != nil {
		return map[string]interface{}{}, errCloudBaseVM
	}
	if o.Items != nil {
		toSerialize["items"] = o.Items
	}
	toSerialize["$t"] = o.T
	return toSerialize, nil
}

func (o *AuditStatsVM) UnmarshalJSON(data []byte) (err error) {
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

	varAuditStatsVM := _AuditStatsVM{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAuditStatsVM)

	if err != nil {
		return err
	}

	*o = AuditStatsVM(varAuditStatsVM)

	return err
}

type NullableAuditStatsVM struct {
	value *AuditStatsVM
	isSet bool
}

func (v NullableAuditStatsVM) Get() *AuditStatsVM {
	return v.value
}

func (v *NullableAuditStatsVM) Set(val *AuditStatsVM) {
	v.value = val
	v.isSet = true
}

func (v NullableAuditStatsVM) IsSet() bool {
	return v.isSet
}

func (v *NullableAuditStatsVM) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuditStatsVM(val *AuditStatsVM) *NullableAuditStatsVM {
	return &NullableAuditStatsVM{value: val, isSet: true}
}

func (v NullableAuditStatsVM) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuditStatsVM) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


