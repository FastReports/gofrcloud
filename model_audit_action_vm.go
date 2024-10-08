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

// checks if the AuditActionVM type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuditActionVM{}

// AuditActionVM struct for AuditActionVM
type AuditActionVM struct {
	CloudBaseVM
	UserId NullableString `json:"userId,omitempty"`
	EntityId NullableString `json:"entityId,omitempty"`
	SubscriptionId NullableString `json:"subscriptionId,omitempty"`
	Type *AuditType `json:"type,omitempty"`
	Id NullableString `json:"id,omitempty"`
	CreatedTime *time.Time `json:"createdTime,omitempty"`
	CreatorUserId NullableString `json:"creatorUserId,omitempty"`
	Name NullableString `json:"name,omitempty"`
	AdminAction *bool `json:"adminAction,omitempty"`
	T string `json:"$t"`
}

type _AuditActionVM AuditActionVM

// NewAuditActionVM instantiates a new AuditActionVM object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuditActionVM(t string) *AuditActionVM {
	this := AuditActionVM{}
	this.T = t
	return &this
}

// NewAuditActionVMWithDefaults instantiates a new AuditActionVM object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuditActionVMWithDefaults() *AuditActionVM {
	this := AuditActionVM{}
	return &this
}

// GetUserId returns the UserId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AuditActionVM) GetUserId() string {
	if o == nil || IsNil(o.UserId.Get()) {
		var ret string
		return ret
	}
	return *o.UserId.Get()
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AuditActionVM) GetUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.UserId.Get(), o.UserId.IsSet()
}

// HasUserId returns a boolean if a field has been set.
func (o *AuditActionVM) HasUserId() bool {
	if o != nil && o.UserId.IsSet() {
		return true
	}

	return false
}

// SetUserId gets a reference to the given NullableString and assigns it to the UserId field.
func (o *AuditActionVM) SetUserId(v string) {
	o.UserId.Set(&v)
}
// SetUserIdNil sets the value for UserId to be an explicit nil
func (o *AuditActionVM) SetUserIdNil() {
	o.UserId.Set(nil)
}

// UnsetUserId ensures that no value is present for UserId, not even an explicit nil
func (o *AuditActionVM) UnsetUserId() {
	o.UserId.Unset()
}

// GetEntityId returns the EntityId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AuditActionVM) GetEntityId() string {
	if o == nil || IsNil(o.EntityId.Get()) {
		var ret string
		return ret
	}
	return *o.EntityId.Get()
}

// GetEntityIdOk returns a tuple with the EntityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AuditActionVM) GetEntityIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.EntityId.Get(), o.EntityId.IsSet()
}

// HasEntityId returns a boolean if a field has been set.
func (o *AuditActionVM) HasEntityId() bool {
	if o != nil && o.EntityId.IsSet() {
		return true
	}

	return false
}

// SetEntityId gets a reference to the given NullableString and assigns it to the EntityId field.
func (o *AuditActionVM) SetEntityId(v string) {
	o.EntityId.Set(&v)
}
// SetEntityIdNil sets the value for EntityId to be an explicit nil
func (o *AuditActionVM) SetEntityIdNil() {
	o.EntityId.Set(nil)
}

// UnsetEntityId ensures that no value is present for EntityId, not even an explicit nil
func (o *AuditActionVM) UnsetEntityId() {
	o.EntityId.Unset()
}

// GetSubscriptionId returns the SubscriptionId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AuditActionVM) GetSubscriptionId() string {
	if o == nil || IsNil(o.SubscriptionId.Get()) {
		var ret string
		return ret
	}
	return *o.SubscriptionId.Get()
}

// GetSubscriptionIdOk returns a tuple with the SubscriptionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AuditActionVM) GetSubscriptionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SubscriptionId.Get(), o.SubscriptionId.IsSet()
}

// HasSubscriptionId returns a boolean if a field has been set.
func (o *AuditActionVM) HasSubscriptionId() bool {
	if o != nil && o.SubscriptionId.IsSet() {
		return true
	}

	return false
}

// SetSubscriptionId gets a reference to the given NullableString and assigns it to the SubscriptionId field.
func (o *AuditActionVM) SetSubscriptionId(v string) {
	o.SubscriptionId.Set(&v)
}
// SetSubscriptionIdNil sets the value for SubscriptionId to be an explicit nil
func (o *AuditActionVM) SetSubscriptionIdNil() {
	o.SubscriptionId.Set(nil)
}

// UnsetSubscriptionId ensures that no value is present for SubscriptionId, not even an explicit nil
func (o *AuditActionVM) UnsetSubscriptionId() {
	o.SubscriptionId.Unset()
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *AuditActionVM) GetType() AuditType {
	if o == nil || IsNil(o.Type) {
		var ret AuditType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditActionVM) GetTypeOk() (*AuditType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *AuditActionVM) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given AuditType and assigns it to the Type field.
func (o *AuditActionVM) SetType(v AuditType) {
	o.Type = &v
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AuditActionVM) GetId() string {
	if o == nil || IsNil(o.Id.Get()) {
		var ret string
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AuditActionVM) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *AuditActionVM) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableString and assigns it to the Id field.
func (o *AuditActionVM) SetId(v string) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *AuditActionVM) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *AuditActionVM) UnsetId() {
	o.Id.Unset()
}

// GetCreatedTime returns the CreatedTime field value if set, zero value otherwise.
func (o *AuditActionVM) GetCreatedTime() time.Time {
	if o == nil || IsNil(o.CreatedTime) {
		var ret time.Time
		return ret
	}
	return *o.CreatedTime
}

// GetCreatedTimeOk returns a tuple with the CreatedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditActionVM) GetCreatedTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedTime) {
		return nil, false
	}
	return o.CreatedTime, true
}

// HasCreatedTime returns a boolean if a field has been set.
func (o *AuditActionVM) HasCreatedTime() bool {
	if o != nil && !IsNil(o.CreatedTime) {
		return true
	}

	return false
}

// SetCreatedTime gets a reference to the given time.Time and assigns it to the CreatedTime field.
func (o *AuditActionVM) SetCreatedTime(v time.Time) {
	o.CreatedTime = &v
}

// GetCreatorUserId returns the CreatorUserId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AuditActionVM) GetCreatorUserId() string {
	if o == nil || IsNil(o.CreatorUserId.Get()) {
		var ret string
		return ret
	}
	return *o.CreatorUserId.Get()
}

// GetCreatorUserIdOk returns a tuple with the CreatorUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AuditActionVM) GetCreatorUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatorUserId.Get(), o.CreatorUserId.IsSet()
}

// HasCreatorUserId returns a boolean if a field has been set.
func (o *AuditActionVM) HasCreatorUserId() bool {
	if o != nil && o.CreatorUserId.IsSet() {
		return true
	}

	return false
}

// SetCreatorUserId gets a reference to the given NullableString and assigns it to the CreatorUserId field.
func (o *AuditActionVM) SetCreatorUserId(v string) {
	o.CreatorUserId.Set(&v)
}
// SetCreatorUserIdNil sets the value for CreatorUserId to be an explicit nil
func (o *AuditActionVM) SetCreatorUserIdNil() {
	o.CreatorUserId.Set(nil)
}

// UnsetCreatorUserId ensures that no value is present for CreatorUserId, not even an explicit nil
func (o *AuditActionVM) UnsetCreatorUserId() {
	o.CreatorUserId.Unset()
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AuditActionVM) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AuditActionVM) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *AuditActionVM) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *AuditActionVM) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *AuditActionVM) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *AuditActionVM) UnsetName() {
	o.Name.Unset()
}

// GetAdminAction returns the AdminAction field value if set, zero value otherwise.
func (o *AuditActionVM) GetAdminAction() bool {
	if o == nil || IsNil(o.AdminAction) {
		var ret bool
		return ret
	}
	return *o.AdminAction
}

// GetAdminActionOk returns a tuple with the AdminAction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditActionVM) GetAdminActionOk() (*bool, bool) {
	if o == nil || IsNil(o.AdminAction) {
		return nil, false
	}
	return o.AdminAction, true
}

// HasAdminAction returns a boolean if a field has been set.
func (o *AuditActionVM) HasAdminAction() bool {
	if o != nil && !IsNil(o.AdminAction) {
		return true
	}

	return false
}

// SetAdminAction gets a reference to the given bool and assigns it to the AdminAction field.
func (o *AuditActionVM) SetAdminAction(v bool) {
	o.AdminAction = &v
}

// GetT returns the T field value
func (o *AuditActionVM) GetT() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.T
}

// GetTOk returns a tuple with the T field value
// and a boolean to check if the value has been set.
func (o *AuditActionVM) GetTOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.T, true
}

// SetT sets field value
func (o *AuditActionVM) SetT(v string) {
	o.T = v
}

func (o AuditActionVM) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuditActionVM) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedCloudBaseVM, errCloudBaseVM := json.Marshal(o.CloudBaseVM)
	if errCloudBaseVM != nil {
		return map[string]interface{}{}, errCloudBaseVM
	}
	errCloudBaseVM = json.Unmarshal([]byte(serializedCloudBaseVM), &toSerialize)
	if errCloudBaseVM != nil {
		return map[string]interface{}{}, errCloudBaseVM
	}
	if o.UserId.IsSet() {
		toSerialize["userId"] = o.UserId.Get()
	}
	if o.EntityId.IsSet() {
		toSerialize["entityId"] = o.EntityId.Get()
	}
	if o.SubscriptionId.IsSet() {
		toSerialize["subscriptionId"] = o.SubscriptionId.Get()
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if !IsNil(o.CreatedTime) {
		toSerialize["createdTime"] = o.CreatedTime
	}
	if o.CreatorUserId.IsSet() {
		toSerialize["creatorUserId"] = o.CreatorUserId.Get()
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if !IsNil(o.AdminAction) {
		toSerialize["adminAction"] = o.AdminAction
	}
	toSerialize["$t"] = o.T
	return toSerialize, nil
}

func (o *AuditActionVM) UnmarshalJSON(data []byte) (err error) {
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

	varAuditActionVM := _AuditActionVM{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAuditActionVM)

	if err != nil {
		return err
	}

	*o = AuditActionVM(varAuditActionVM)

	return err
}

type NullableAuditActionVM struct {
	value *AuditActionVM
	isSet bool
}

func (v NullableAuditActionVM) Get() *AuditActionVM {
	return v.value
}

func (v *NullableAuditActionVM) Set(val *AuditActionVM) {
	v.value = val
	v.isSet = true
}

func (v NullableAuditActionVM) IsSet() bool {
	return v.isSet
}

func (v *NullableAuditActionVM) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuditActionVM(val *AuditActionVM) *NullableAuditActionVM {
	return &NullableAuditActionVM{value: val, isSet: true}
}

func (v NullableAuditActionVM) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuditActionVM) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


