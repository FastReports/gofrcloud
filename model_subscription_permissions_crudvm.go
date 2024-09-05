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

// checks if the SubscriptionPermissionsCRUDVM type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubscriptionPermissionsCRUDVM{}

// SubscriptionPermissionsCRUDVM struct for SubscriptionPermissionsCRUDVM
type SubscriptionPermissionsCRUDVM struct {
	CloudBaseVM
	OwnerId NullableString `json:"ownerId,omitempty"`
	Owner *SubscriptionPermissionCRUDVM `json:"owner,omitempty"`
	Groups map[string]SubscriptionPermissionCRUDVM `json:"groups,omitempty"`
	Other *SubscriptionPermissionCRUDVM `json:"other,omitempty"`
	Anon *SubscriptionPermissionCRUDVM `json:"anon,omitempty"`
	T string `json:"$t"`
}

type _SubscriptionPermissionsCRUDVM SubscriptionPermissionsCRUDVM

// NewSubscriptionPermissionsCRUDVM instantiates a new SubscriptionPermissionsCRUDVM object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscriptionPermissionsCRUDVM(t string) *SubscriptionPermissionsCRUDVM {
	this := SubscriptionPermissionsCRUDVM{}
	this.T = t
	return &this
}

// NewSubscriptionPermissionsCRUDVMWithDefaults instantiates a new SubscriptionPermissionsCRUDVM object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionPermissionsCRUDVMWithDefaults() *SubscriptionPermissionsCRUDVM {
	this := SubscriptionPermissionsCRUDVM{}
	return &this
}

// GetOwnerId returns the OwnerId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SubscriptionPermissionsCRUDVM) GetOwnerId() string {
	if o == nil || IsNil(o.OwnerId.Get()) {
		var ret string
		return ret
	}
	return *o.OwnerId.Get()
}

// GetOwnerIdOk returns a tuple with the OwnerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SubscriptionPermissionsCRUDVM) GetOwnerIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OwnerId.Get(), o.OwnerId.IsSet()
}

// HasOwnerId returns a boolean if a field has been set.
func (o *SubscriptionPermissionsCRUDVM) HasOwnerId() bool {
	if o != nil && o.OwnerId.IsSet() {
		return true
	}

	return false
}

// SetOwnerId gets a reference to the given NullableString and assigns it to the OwnerId field.
func (o *SubscriptionPermissionsCRUDVM) SetOwnerId(v string) {
	o.OwnerId.Set(&v)
}
// SetOwnerIdNil sets the value for OwnerId to be an explicit nil
func (o *SubscriptionPermissionsCRUDVM) SetOwnerIdNil() {
	o.OwnerId.Set(nil)
}

// UnsetOwnerId ensures that no value is present for OwnerId, not even an explicit nil
func (o *SubscriptionPermissionsCRUDVM) UnsetOwnerId() {
	o.OwnerId.Unset()
}

// GetOwner returns the Owner field value if set, zero value otherwise.
func (o *SubscriptionPermissionsCRUDVM) GetOwner() SubscriptionPermissionCRUDVM {
	if o == nil || IsNil(o.Owner) {
		var ret SubscriptionPermissionCRUDVM
		return ret
	}
	return *o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionPermissionsCRUDVM) GetOwnerOk() (*SubscriptionPermissionCRUDVM, bool) {
	if o == nil || IsNil(o.Owner) {
		return nil, false
	}
	return o.Owner, true
}

// HasOwner returns a boolean if a field has been set.
func (o *SubscriptionPermissionsCRUDVM) HasOwner() bool {
	if o != nil && !IsNil(o.Owner) {
		return true
	}

	return false
}

// SetOwner gets a reference to the given SubscriptionPermissionCRUDVM and assigns it to the Owner field.
func (o *SubscriptionPermissionsCRUDVM) SetOwner(v SubscriptionPermissionCRUDVM) {
	o.Owner = &v
}

// GetGroups returns the Groups field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SubscriptionPermissionsCRUDVM) GetGroups() map[string]SubscriptionPermissionCRUDVM {
	if o == nil {
		var ret map[string]SubscriptionPermissionCRUDVM
		return ret
	}
	return o.Groups
}

// GetGroupsOk returns a tuple with the Groups field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SubscriptionPermissionsCRUDVM) GetGroupsOk() (*map[string]SubscriptionPermissionCRUDVM, bool) {
	if o == nil || IsNil(o.Groups) {
		return nil, false
	}
	return &o.Groups, true
}

// HasGroups returns a boolean if a field has been set.
func (o *SubscriptionPermissionsCRUDVM) HasGroups() bool {
	if o != nil && IsNil(o.Groups) {
		return true
	}

	return false
}

// SetGroups gets a reference to the given map[string]SubscriptionPermissionCRUDVM and assigns it to the Groups field.
func (o *SubscriptionPermissionsCRUDVM) SetGroups(v map[string]SubscriptionPermissionCRUDVM) {
	o.Groups = v
}

// GetOther returns the Other field value if set, zero value otherwise.
func (o *SubscriptionPermissionsCRUDVM) GetOther() SubscriptionPermissionCRUDVM {
	if o == nil || IsNil(o.Other) {
		var ret SubscriptionPermissionCRUDVM
		return ret
	}
	return *o.Other
}

// GetOtherOk returns a tuple with the Other field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionPermissionsCRUDVM) GetOtherOk() (*SubscriptionPermissionCRUDVM, bool) {
	if o == nil || IsNil(o.Other) {
		return nil, false
	}
	return o.Other, true
}

// HasOther returns a boolean if a field has been set.
func (o *SubscriptionPermissionsCRUDVM) HasOther() bool {
	if o != nil && !IsNil(o.Other) {
		return true
	}

	return false
}

// SetOther gets a reference to the given SubscriptionPermissionCRUDVM and assigns it to the Other field.
func (o *SubscriptionPermissionsCRUDVM) SetOther(v SubscriptionPermissionCRUDVM) {
	o.Other = &v
}

// GetAnon returns the Anon field value if set, zero value otherwise.
func (o *SubscriptionPermissionsCRUDVM) GetAnon() SubscriptionPermissionCRUDVM {
	if o == nil || IsNil(o.Anon) {
		var ret SubscriptionPermissionCRUDVM
		return ret
	}
	return *o.Anon
}

// GetAnonOk returns a tuple with the Anon field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionPermissionsCRUDVM) GetAnonOk() (*SubscriptionPermissionCRUDVM, bool) {
	if o == nil || IsNil(o.Anon) {
		return nil, false
	}
	return o.Anon, true
}

// HasAnon returns a boolean if a field has been set.
func (o *SubscriptionPermissionsCRUDVM) HasAnon() bool {
	if o != nil && !IsNil(o.Anon) {
		return true
	}

	return false
}

// SetAnon gets a reference to the given SubscriptionPermissionCRUDVM and assigns it to the Anon field.
func (o *SubscriptionPermissionsCRUDVM) SetAnon(v SubscriptionPermissionCRUDVM) {
	o.Anon = &v
}

// GetT returns the T field value
func (o *SubscriptionPermissionsCRUDVM) GetT() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.T
}

// GetTOk returns a tuple with the T field value
// and a boolean to check if the value has been set.
func (o *SubscriptionPermissionsCRUDVM) GetTOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.T, true
}

// SetT sets field value
func (o *SubscriptionPermissionsCRUDVM) SetT(v string) {
	o.T = v
}

func (o SubscriptionPermissionsCRUDVM) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubscriptionPermissionsCRUDVM) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedCloudBaseVM, errCloudBaseVM := json.Marshal(o.CloudBaseVM)
	if errCloudBaseVM != nil {
		return map[string]interface{}{}, errCloudBaseVM
	}
	errCloudBaseVM = json.Unmarshal([]byte(serializedCloudBaseVM), &toSerialize)
	if errCloudBaseVM != nil {
		return map[string]interface{}{}, errCloudBaseVM
	}
	if o.OwnerId.IsSet() {
		toSerialize["ownerId"] = o.OwnerId.Get()
	}
	if !IsNil(o.Owner) {
		toSerialize["owner"] = o.Owner
	}
	if o.Groups != nil {
		toSerialize["groups"] = o.Groups
	}
	if !IsNil(o.Other) {
		toSerialize["other"] = o.Other
	}
	if !IsNil(o.Anon) {
		toSerialize["anon"] = o.Anon
	}
	toSerialize["$t"] = o.T
	return toSerialize, nil
}

func (o *SubscriptionPermissionsCRUDVM) UnmarshalJSON(data []byte) (err error) {
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

	varSubscriptionPermissionsCRUDVM := _SubscriptionPermissionsCRUDVM{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSubscriptionPermissionsCRUDVM)

	if err != nil {
		return err
	}

	*o = SubscriptionPermissionsCRUDVM(varSubscriptionPermissionsCRUDVM)

	return err
}

type NullableSubscriptionPermissionsCRUDVM struct {
	value *SubscriptionPermissionsCRUDVM
	isSet bool
}

func (v NullableSubscriptionPermissionsCRUDVM) Get() *SubscriptionPermissionsCRUDVM {
	return v.value
}

func (v *NullableSubscriptionPermissionsCRUDVM) Set(val *SubscriptionPermissionsCRUDVM) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionPermissionsCRUDVM) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionPermissionsCRUDVM) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionPermissionsCRUDVM(val *SubscriptionPermissionsCRUDVM) *NullableSubscriptionPermissionsCRUDVM {
	return &NullableSubscriptionPermissionsCRUDVM{value: val, isSet: true}
}

func (v NullableSubscriptionPermissionsCRUDVM) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionPermissionsCRUDVM) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

