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

// RunEmailTaskVM struct for RunEmailTaskVM
type RunEmailTaskVM struct {
	Password NullableString `json:"password,omitempty"`
	Body NullableString `json:"body,omitempty"`
	IsBodyHtml *bool `json:"isBodyHtml,omitempty"`
	Subject NullableString `json:"subject,omitempty"`
	To []string `json:"to,omitempty"`
	From NullableString `json:"from,omitempty"`
	Username NullableString `json:"username,omitempty"`
	Server NullableString `json:"server,omitempty"`
	Port *int32 `json:"port,omitempty"`
	EnableSsl *bool `json:"enableSsl,omitempty"`
	SubscriptionId NullableString `json:"subscriptionId,omitempty"`
	Type *TaskType `json:"type,omitempty"`
}

// NewRunEmailTaskVM instantiates a new RunEmailTaskVM object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRunEmailTaskVM() *RunEmailTaskVM {
	this := RunEmailTaskVM{}
	return &this
}

// NewRunEmailTaskVMWithDefaults instantiates a new RunEmailTaskVM object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRunEmailTaskVMWithDefaults() *RunEmailTaskVM {
	this := RunEmailTaskVM{}
	return &this
}

// GetPassword returns the Password field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RunEmailTaskVM) GetPassword() string {
	if o == nil || o.Password.Get() == nil {
		var ret string
		return ret
	}
	return *o.Password.Get()
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RunEmailTaskVM) GetPasswordOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Password.Get(), o.Password.IsSet()
}

// HasPassword returns a boolean if a field has been set.
func (o *RunEmailTaskVM) HasPassword() bool {
	if o != nil && o.Password.IsSet() {
		return true
	}

	return false
}

// SetPassword gets a reference to the given NullableString and assigns it to the Password field.
func (o *RunEmailTaskVM) SetPassword(v string) {
	o.Password.Set(&v)
}
// SetPasswordNil sets the value for Password to be an explicit nil
func (o *RunEmailTaskVM) SetPasswordNil() {
	o.Password.Set(nil)
}

// UnsetPassword ensures that no value is present for Password, not even an explicit nil
func (o *RunEmailTaskVM) UnsetPassword() {
	o.Password.Unset()
}

// GetBody returns the Body field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RunEmailTaskVM) GetBody() string {
	if o == nil || o.Body.Get() == nil {
		var ret string
		return ret
	}
	return *o.Body.Get()
}

// GetBodyOk returns a tuple with the Body field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RunEmailTaskVM) GetBodyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Body.Get(), o.Body.IsSet()
}

// HasBody returns a boolean if a field has been set.
func (o *RunEmailTaskVM) HasBody() bool {
	if o != nil && o.Body.IsSet() {
		return true
	}

	return false
}

// SetBody gets a reference to the given NullableString and assigns it to the Body field.
func (o *RunEmailTaskVM) SetBody(v string) {
	o.Body.Set(&v)
}
// SetBodyNil sets the value for Body to be an explicit nil
func (o *RunEmailTaskVM) SetBodyNil() {
	o.Body.Set(nil)
}

// UnsetBody ensures that no value is present for Body, not even an explicit nil
func (o *RunEmailTaskVM) UnsetBody() {
	o.Body.Unset()
}

// GetIsBodyHtml returns the IsBodyHtml field value if set, zero value otherwise.
func (o *RunEmailTaskVM) GetIsBodyHtml() bool {
	if o == nil || o.IsBodyHtml == nil {
		var ret bool
		return ret
	}
	return *o.IsBodyHtml
}

// GetIsBodyHtmlOk returns a tuple with the IsBodyHtml field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RunEmailTaskVM) GetIsBodyHtmlOk() (*bool, bool) {
	if o == nil || o.IsBodyHtml == nil {
		return nil, false
	}
	return o.IsBodyHtml, true
}

// HasIsBodyHtml returns a boolean if a field has been set.
func (o *RunEmailTaskVM) HasIsBodyHtml() bool {
	if o != nil && o.IsBodyHtml != nil {
		return true
	}

	return false
}

// SetIsBodyHtml gets a reference to the given bool and assigns it to the IsBodyHtml field.
func (o *RunEmailTaskVM) SetIsBodyHtml(v bool) {
	o.IsBodyHtml = &v
}

// GetSubject returns the Subject field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RunEmailTaskVM) GetSubject() string {
	if o == nil || o.Subject.Get() == nil {
		var ret string
		return ret
	}
	return *o.Subject.Get()
}

// GetSubjectOk returns a tuple with the Subject field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RunEmailTaskVM) GetSubjectOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Subject.Get(), o.Subject.IsSet()
}

// HasSubject returns a boolean if a field has been set.
func (o *RunEmailTaskVM) HasSubject() bool {
	if o != nil && o.Subject.IsSet() {
		return true
	}

	return false
}

// SetSubject gets a reference to the given NullableString and assigns it to the Subject field.
func (o *RunEmailTaskVM) SetSubject(v string) {
	o.Subject.Set(&v)
}
// SetSubjectNil sets the value for Subject to be an explicit nil
func (o *RunEmailTaskVM) SetSubjectNil() {
	o.Subject.Set(nil)
}

// UnsetSubject ensures that no value is present for Subject, not even an explicit nil
func (o *RunEmailTaskVM) UnsetSubject() {
	o.Subject.Unset()
}

// GetTo returns the To field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RunEmailTaskVM) GetTo() []string {
	if o == nil  {
		var ret []string
		return ret
	}
	return o.To
}

// GetToOk returns a tuple with the To field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RunEmailTaskVM) GetToOk() (*[]string, bool) {
	if o == nil || o.To == nil {
		return nil, false
	}
	return &o.To, true
}

// HasTo returns a boolean if a field has been set.
func (o *RunEmailTaskVM) HasTo() bool {
	if o != nil && o.To != nil {
		return true
	}

	return false
}

// SetTo gets a reference to the given []string and assigns it to the To field.
func (o *RunEmailTaskVM) SetTo(v []string) {
	o.To = v
}

// GetFrom returns the From field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RunEmailTaskVM) GetFrom() string {
	if o == nil || o.From.Get() == nil {
		var ret string
		return ret
	}
	return *o.From.Get()
}

// GetFromOk returns a tuple with the From field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RunEmailTaskVM) GetFromOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.From.Get(), o.From.IsSet()
}

// HasFrom returns a boolean if a field has been set.
func (o *RunEmailTaskVM) HasFrom() bool {
	if o != nil && o.From.IsSet() {
		return true
	}

	return false
}

// SetFrom gets a reference to the given NullableString and assigns it to the From field.
func (o *RunEmailTaskVM) SetFrom(v string) {
	o.From.Set(&v)
}
// SetFromNil sets the value for From to be an explicit nil
func (o *RunEmailTaskVM) SetFromNil() {
	o.From.Set(nil)
}

// UnsetFrom ensures that no value is present for From, not even an explicit nil
func (o *RunEmailTaskVM) UnsetFrom() {
	o.From.Unset()
}

// GetUsername returns the Username field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RunEmailTaskVM) GetUsername() string {
	if o == nil || o.Username.Get() == nil {
		var ret string
		return ret
	}
	return *o.Username.Get()
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RunEmailTaskVM) GetUsernameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Username.Get(), o.Username.IsSet()
}

// HasUsername returns a boolean if a field has been set.
func (o *RunEmailTaskVM) HasUsername() bool {
	if o != nil && o.Username.IsSet() {
		return true
	}

	return false
}

// SetUsername gets a reference to the given NullableString and assigns it to the Username field.
func (o *RunEmailTaskVM) SetUsername(v string) {
	o.Username.Set(&v)
}
// SetUsernameNil sets the value for Username to be an explicit nil
func (o *RunEmailTaskVM) SetUsernameNil() {
	o.Username.Set(nil)
}

// UnsetUsername ensures that no value is present for Username, not even an explicit nil
func (o *RunEmailTaskVM) UnsetUsername() {
	o.Username.Unset()
}

// GetServer returns the Server field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RunEmailTaskVM) GetServer() string {
	if o == nil || o.Server.Get() == nil {
		var ret string
		return ret
	}
	return *o.Server.Get()
}

// GetServerOk returns a tuple with the Server field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RunEmailTaskVM) GetServerOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Server.Get(), o.Server.IsSet()
}

// HasServer returns a boolean if a field has been set.
func (o *RunEmailTaskVM) HasServer() bool {
	if o != nil && o.Server.IsSet() {
		return true
	}

	return false
}

// SetServer gets a reference to the given NullableString and assigns it to the Server field.
func (o *RunEmailTaskVM) SetServer(v string) {
	o.Server.Set(&v)
}
// SetServerNil sets the value for Server to be an explicit nil
func (o *RunEmailTaskVM) SetServerNil() {
	o.Server.Set(nil)
}

// UnsetServer ensures that no value is present for Server, not even an explicit nil
func (o *RunEmailTaskVM) UnsetServer() {
	o.Server.Unset()
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *RunEmailTaskVM) GetPort() int32 {
	if o == nil || o.Port == nil {
		var ret int32
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RunEmailTaskVM) GetPortOk() (*int32, bool) {
	if o == nil || o.Port == nil {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *RunEmailTaskVM) HasPort() bool {
	if o != nil && o.Port != nil {
		return true
	}

	return false
}

// SetPort gets a reference to the given int32 and assigns it to the Port field.
func (o *RunEmailTaskVM) SetPort(v int32) {
	o.Port = &v
}

// GetEnableSsl returns the EnableSsl field value if set, zero value otherwise.
func (o *RunEmailTaskVM) GetEnableSsl() bool {
	if o == nil || o.EnableSsl == nil {
		var ret bool
		return ret
	}
	return *o.EnableSsl
}

// GetEnableSslOk returns a tuple with the EnableSsl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RunEmailTaskVM) GetEnableSslOk() (*bool, bool) {
	if o == nil || o.EnableSsl == nil {
		return nil, false
	}
	return o.EnableSsl, true
}

// HasEnableSsl returns a boolean if a field has been set.
func (o *RunEmailTaskVM) HasEnableSsl() bool {
	if o != nil && o.EnableSsl != nil {
		return true
	}

	return false
}

// SetEnableSsl gets a reference to the given bool and assigns it to the EnableSsl field.
func (o *RunEmailTaskVM) SetEnableSsl(v bool) {
	o.EnableSsl = &v
}

// GetSubscriptionId returns the SubscriptionId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RunEmailTaskVM) GetSubscriptionId() string {
	if o == nil || o.SubscriptionId.Get() == nil {
		var ret string
		return ret
	}
	return *o.SubscriptionId.Get()
}

// GetSubscriptionIdOk returns a tuple with the SubscriptionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RunEmailTaskVM) GetSubscriptionIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SubscriptionId.Get(), o.SubscriptionId.IsSet()
}

// HasSubscriptionId returns a boolean if a field has been set.
func (o *RunEmailTaskVM) HasSubscriptionId() bool {
	if o != nil && o.SubscriptionId.IsSet() {
		return true
	}

	return false
}

// SetSubscriptionId gets a reference to the given NullableString and assigns it to the SubscriptionId field.
func (o *RunEmailTaskVM) SetSubscriptionId(v string) {
	o.SubscriptionId.Set(&v)
}
// SetSubscriptionIdNil sets the value for SubscriptionId to be an explicit nil
func (o *RunEmailTaskVM) SetSubscriptionIdNil() {
	o.SubscriptionId.Set(nil)
}

// UnsetSubscriptionId ensures that no value is present for SubscriptionId, not even an explicit nil
func (o *RunEmailTaskVM) UnsetSubscriptionId() {
	o.SubscriptionId.Unset()
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *RunEmailTaskVM) GetType() TaskType {
	if o == nil || o.Type == nil {
		var ret TaskType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RunEmailTaskVM) GetTypeOk() (*TaskType, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *RunEmailTaskVM) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given TaskType and assigns it to the Type field.
func (o *RunEmailTaskVM) SetType(v TaskType) {
	o.Type = &v
}

func (o RunEmailTaskVM) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Password.IsSet() {
		toSerialize["password"] = o.Password.Get()
	}
	if o.Body.IsSet() {
		toSerialize["body"] = o.Body.Get()
	}
	if o.IsBodyHtml != nil {
		toSerialize["isBodyHtml"] = o.IsBodyHtml
	}
	if o.Subject.IsSet() {
		toSerialize["subject"] = o.Subject.Get()
	}
	if o.To != nil {
		toSerialize["to"] = o.To
	}
	if o.From.IsSet() {
		toSerialize["from"] = o.From.Get()
	}
	if o.Username.IsSet() {
		toSerialize["username"] = o.Username.Get()
	}
	if o.Server.IsSet() {
		toSerialize["server"] = o.Server.Get()
	}
	if o.Port != nil {
		toSerialize["port"] = o.Port
	}
	if o.EnableSsl != nil {
		toSerialize["enableSsl"] = o.EnableSsl
	}
	if o.SubscriptionId.IsSet() {
		toSerialize["subscriptionId"] = o.SubscriptionId.Get()
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableRunEmailTaskVM struct {
	value *RunEmailTaskVM
	isSet bool
}

func (v NullableRunEmailTaskVM) Get() *RunEmailTaskVM {
	return v.value
}

func (v *NullableRunEmailTaskVM) Set(val *RunEmailTaskVM) {
	v.value = val
	v.isSet = true
}

func (v NullableRunEmailTaskVM) IsSet() bool {
	return v.isSet
}

func (v *NullableRunEmailTaskVM) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRunEmailTaskVM(val *RunEmailTaskVM) *NullableRunEmailTaskVM {
	return &NullableRunEmailTaskVM{value: val, isSet: true}
}

func (v NullableRunEmailTaskVM) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRunEmailTaskVM) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


