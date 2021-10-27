/*
 * FastReport Cloud
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gofrcloud

import (
	"encoding/json"
)

// CreateWebhookTaskVM struct for CreateWebhookTaskVM
type CreateWebhookTaskVM struct {
	Endpoints []CreateEndpointVM `json:"endpoints,omitempty"`
	Name NullableString `json:"name,omitempty"`
	SubscriptionId NullableString `json:"subscriptionId,omitempty"`
	Type *TaskType `json:"type,omitempty"`
}

// NewCreateWebhookTaskVM instantiates a new CreateWebhookTaskVM object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateWebhookTaskVM() *CreateWebhookTaskVM {
	this := CreateWebhookTaskVM{}
	return &this
}

// NewCreateWebhookTaskVMWithDefaults instantiates a new CreateWebhookTaskVM object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateWebhookTaskVMWithDefaults() *CreateWebhookTaskVM {
	this := CreateWebhookTaskVM{}
	return &this
}

// GetEndpoints returns the Endpoints field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateWebhookTaskVM) GetEndpoints() []CreateEndpointVM {
	if o == nil  {
		var ret []CreateEndpointVM
		return ret
	}
	return o.Endpoints
}

// GetEndpointsOk returns a tuple with the Endpoints field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateWebhookTaskVM) GetEndpointsOk() (*[]CreateEndpointVM, bool) {
	if o == nil || o.Endpoints == nil {
		return nil, false
	}
	return &o.Endpoints, true
}

// HasEndpoints returns a boolean if a field has been set.
func (o *CreateWebhookTaskVM) HasEndpoints() bool {
	if o != nil && o.Endpoints != nil {
		return true
	}

	return false
}

// SetEndpoints gets a reference to the given []CreateEndpointVM and assigns it to the Endpoints field.
func (o *CreateWebhookTaskVM) SetEndpoints(v []CreateEndpointVM) {
	o.Endpoints = v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateWebhookTaskVM) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateWebhookTaskVM) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *CreateWebhookTaskVM) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *CreateWebhookTaskVM) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *CreateWebhookTaskVM) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *CreateWebhookTaskVM) UnsetName() {
	o.Name.Unset()
}

// GetSubscriptionId returns the SubscriptionId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateWebhookTaskVM) GetSubscriptionId() string {
	if o == nil || o.SubscriptionId.Get() == nil {
		var ret string
		return ret
	}
	return *o.SubscriptionId.Get()
}

// GetSubscriptionIdOk returns a tuple with the SubscriptionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateWebhookTaskVM) GetSubscriptionIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SubscriptionId.Get(), o.SubscriptionId.IsSet()
}

// HasSubscriptionId returns a boolean if a field has been set.
func (o *CreateWebhookTaskVM) HasSubscriptionId() bool {
	if o != nil && o.SubscriptionId.IsSet() {
		return true
	}

	return false
}

// SetSubscriptionId gets a reference to the given NullableString and assigns it to the SubscriptionId field.
func (o *CreateWebhookTaskVM) SetSubscriptionId(v string) {
	o.SubscriptionId.Set(&v)
}
// SetSubscriptionIdNil sets the value for SubscriptionId to be an explicit nil
func (o *CreateWebhookTaskVM) SetSubscriptionIdNil() {
	o.SubscriptionId.Set(nil)
}

// UnsetSubscriptionId ensures that no value is present for SubscriptionId, not even an explicit nil
func (o *CreateWebhookTaskVM) UnsetSubscriptionId() {
	o.SubscriptionId.Unset()
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *CreateWebhookTaskVM) GetType() TaskType {
	if o == nil || o.Type == nil {
		var ret TaskType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateWebhookTaskVM) GetTypeOk() (*TaskType, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *CreateWebhookTaskVM) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given TaskType and assigns it to the Type field.
func (o *CreateWebhookTaskVM) SetType(v TaskType) {
	o.Type = &v
}

func (o CreateWebhookTaskVM) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Endpoints != nil {
		toSerialize["endpoints"] = o.Endpoints
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.SubscriptionId.IsSet() {
		toSerialize["subscriptionId"] = o.SubscriptionId.Get()
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableCreateWebhookTaskVM struct {
	value *CreateWebhookTaskVM
	isSet bool
}

func (v NullableCreateWebhookTaskVM) Get() *CreateWebhookTaskVM {
	return v.value
}

func (v *NullableCreateWebhookTaskVM) Set(val *CreateWebhookTaskVM) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateWebhookTaskVM) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateWebhookTaskVM) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateWebhookTaskVM(val *CreateWebhookTaskVM) *NullableCreateWebhookTaskVM {
	return &NullableCreateWebhookTaskVM{value: val, isSet: true}
}

func (v NullableCreateWebhookTaskVM) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateWebhookTaskVM) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


