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

// ServerConfigurationVM struct for ServerConfigurationVM
type ServerConfigurationVM struct {
	Title NullableString `json:"title,omitempty"`
	CorporateServerMode *bool `json:"corporateServerMode,omitempty"`
	AppMixins *AppMixins `json:"appMixins,omitempty"`
}

// NewServerConfigurationVM instantiates a new ServerConfigurationVM object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServerConfigurationVM() *ServerConfigurationVM {
	this := ServerConfigurationVM{}
	return &this
}

// NewServerConfigurationVMWithDefaults instantiates a new ServerConfigurationVM object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServerConfigurationVMWithDefaults() *ServerConfigurationVM {
	this := ServerConfigurationVM{}
	return &this
}

// GetTitle returns the Title field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ServerConfigurationVM) GetTitle() string {
	if o == nil || o.Title.Get() == nil {
		var ret string
		return ret
	}
	return *o.Title.Get()
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ServerConfigurationVM) GetTitleOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Title.Get(), o.Title.IsSet()
}

// HasTitle returns a boolean if a field has been set.
func (o *ServerConfigurationVM) HasTitle() bool {
	if o != nil && o.Title.IsSet() {
		return true
	}

	return false
}

// SetTitle gets a reference to the given NullableString and assigns it to the Title field.
func (o *ServerConfigurationVM) SetTitle(v string) {
	o.Title.Set(&v)
}
// SetTitleNil sets the value for Title to be an explicit nil
func (o *ServerConfigurationVM) SetTitleNil() {
	o.Title.Set(nil)
}

// UnsetTitle ensures that no value is present for Title, not even an explicit nil
func (o *ServerConfigurationVM) UnsetTitle() {
	o.Title.Unset()
}

// GetCorporateServerMode returns the CorporateServerMode field value if set, zero value otherwise.
func (o *ServerConfigurationVM) GetCorporateServerMode() bool {
	if o == nil || o.CorporateServerMode == nil {
		var ret bool
		return ret
	}
	return *o.CorporateServerMode
}

// GetCorporateServerModeOk returns a tuple with the CorporateServerMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerConfigurationVM) GetCorporateServerModeOk() (*bool, bool) {
	if o == nil || o.CorporateServerMode == nil {
		return nil, false
	}
	return o.CorporateServerMode, true
}

// HasCorporateServerMode returns a boolean if a field has been set.
func (o *ServerConfigurationVM) HasCorporateServerMode() bool {
	if o != nil && o.CorporateServerMode != nil {
		return true
	}

	return false
}

// SetCorporateServerMode gets a reference to the given bool and assigns it to the CorporateServerMode field.
func (o *ServerConfigurationVM) SetCorporateServerMode(v bool) {
	o.CorporateServerMode = &v
}

// GetAppMixins returns the AppMixins field value if set, zero value otherwise.
func (o *ServerConfigurationVM) GetAppMixins() AppMixins {
	if o == nil || o.AppMixins == nil {
		var ret AppMixins
		return ret
	}
	return *o.AppMixins
}

// GetAppMixinsOk returns a tuple with the AppMixins field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerConfigurationVM) GetAppMixinsOk() (*AppMixins, bool) {
	if o == nil || o.AppMixins == nil {
		return nil, false
	}
	return o.AppMixins, true
}

// HasAppMixins returns a boolean if a field has been set.
func (o *ServerConfigurationVM) HasAppMixins() bool {
	if o != nil && o.AppMixins != nil {
		return true
	}

	return false
}

// SetAppMixins gets a reference to the given AppMixins and assigns it to the AppMixins field.
func (o *ServerConfigurationVM) SetAppMixins(v AppMixins) {
	o.AppMixins = &v
}

func (o ServerConfigurationVM) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Title.IsSet() {
		toSerialize["title"] = o.Title.Get()
	}
	if o.CorporateServerMode != nil {
		toSerialize["corporateServerMode"] = o.CorporateServerMode
	}
	if o.AppMixins != nil {
		toSerialize["appMixins"] = o.AppMixins
	}
	return json.Marshal(toSerialize)
}

type NullableServerConfigurationVM struct {
	value *ServerConfigurationVM
	isSet bool
}

func (v NullableServerConfigurationVM) Get() *ServerConfigurationVM {
	return v.value
}

func (v *NullableServerConfigurationVM) Set(val *ServerConfigurationVM) {
	v.value = val
	v.isSet = true
}

func (v NullableServerConfigurationVM) IsSet() bool {
	return v.isSet
}

func (v *NullableServerConfigurationVM) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServerConfigurationVM(val *ServerConfigurationVM) *NullableServerConfigurationVM {
	return &NullableServerConfigurationVM{value: val, isSet: true}
}

func (v NullableServerConfigurationVM) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServerConfigurationVM) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


