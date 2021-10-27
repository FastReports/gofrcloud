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

// MyPermissionsVM struct for MyPermissionsVM
type MyPermissionsVM struct {
	Subscription *SubscriptionPermission `json:"subscription,omitempty"`
	Files *FilePermission `json:"files,omitempty"`
	Datasources *DataSourcePermission `json:"datasources,omitempty"`
	Groups *GroupPermission `json:"groups,omitempty"`
}

// NewMyPermissionsVM instantiates a new MyPermissionsVM object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMyPermissionsVM() *MyPermissionsVM {
	this := MyPermissionsVM{}
	return &this
}

// NewMyPermissionsVMWithDefaults instantiates a new MyPermissionsVM object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMyPermissionsVMWithDefaults() *MyPermissionsVM {
	this := MyPermissionsVM{}
	return &this
}

// GetSubscription returns the Subscription field value if set, zero value otherwise.
func (o *MyPermissionsVM) GetSubscription() SubscriptionPermission {
	if o == nil || o.Subscription == nil {
		var ret SubscriptionPermission
		return ret
	}
	return *o.Subscription
}

// GetSubscriptionOk returns a tuple with the Subscription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MyPermissionsVM) GetSubscriptionOk() (*SubscriptionPermission, bool) {
	if o == nil || o.Subscription == nil {
		return nil, false
	}
	return o.Subscription, true
}

// HasSubscription returns a boolean if a field has been set.
func (o *MyPermissionsVM) HasSubscription() bool {
	if o != nil && o.Subscription != nil {
		return true
	}

	return false
}

// SetSubscription gets a reference to the given SubscriptionPermission and assigns it to the Subscription field.
func (o *MyPermissionsVM) SetSubscription(v SubscriptionPermission) {
	o.Subscription = &v
}

// GetFiles returns the Files field value if set, zero value otherwise.
func (o *MyPermissionsVM) GetFiles() FilePermission {
	if o == nil || o.Files == nil {
		var ret FilePermission
		return ret
	}
	return *o.Files
}

// GetFilesOk returns a tuple with the Files field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MyPermissionsVM) GetFilesOk() (*FilePermission, bool) {
	if o == nil || o.Files == nil {
		return nil, false
	}
	return o.Files, true
}

// HasFiles returns a boolean if a field has been set.
func (o *MyPermissionsVM) HasFiles() bool {
	if o != nil && o.Files != nil {
		return true
	}

	return false
}

// SetFiles gets a reference to the given FilePermission and assigns it to the Files field.
func (o *MyPermissionsVM) SetFiles(v FilePermission) {
	o.Files = &v
}

// GetDatasources returns the Datasources field value if set, zero value otherwise.
func (o *MyPermissionsVM) GetDatasources() DataSourcePermission {
	if o == nil || o.Datasources == nil {
		var ret DataSourcePermission
		return ret
	}
	return *o.Datasources
}

// GetDatasourcesOk returns a tuple with the Datasources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MyPermissionsVM) GetDatasourcesOk() (*DataSourcePermission, bool) {
	if o == nil || o.Datasources == nil {
		return nil, false
	}
	return o.Datasources, true
}

// HasDatasources returns a boolean if a field has been set.
func (o *MyPermissionsVM) HasDatasources() bool {
	if o != nil && o.Datasources != nil {
		return true
	}

	return false
}

// SetDatasources gets a reference to the given DataSourcePermission and assigns it to the Datasources field.
func (o *MyPermissionsVM) SetDatasources(v DataSourcePermission) {
	o.Datasources = &v
}

// GetGroups returns the Groups field value if set, zero value otherwise.
func (o *MyPermissionsVM) GetGroups() GroupPermission {
	if o == nil || o.Groups == nil {
		var ret GroupPermission
		return ret
	}
	return *o.Groups
}

// GetGroupsOk returns a tuple with the Groups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MyPermissionsVM) GetGroupsOk() (*GroupPermission, bool) {
	if o == nil || o.Groups == nil {
		return nil, false
	}
	return o.Groups, true
}

// HasGroups returns a boolean if a field has been set.
func (o *MyPermissionsVM) HasGroups() bool {
	if o != nil && o.Groups != nil {
		return true
	}

	return false
}

// SetGroups gets a reference to the given GroupPermission and assigns it to the Groups field.
func (o *MyPermissionsVM) SetGroups(v GroupPermission) {
	o.Groups = &v
}

func (o MyPermissionsVM) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Subscription != nil {
		toSerialize["subscription"] = o.Subscription
	}
	if o.Files != nil {
		toSerialize["files"] = o.Files
	}
	if o.Datasources != nil {
		toSerialize["datasources"] = o.Datasources
	}
	if o.Groups != nil {
		toSerialize["groups"] = o.Groups
	}
	return json.Marshal(toSerialize)
}

type NullableMyPermissionsVM struct {
	value *MyPermissionsVM
	isSet bool
}

func (v NullableMyPermissionsVM) Get() *MyPermissionsVM {
	return v.value
}

func (v *NullableMyPermissionsVM) Set(val *MyPermissionsVM) {
	v.value = val
	v.isSet = true
}

func (v NullableMyPermissionsVM) IsSet() bool {
	return v.isSet
}

func (v *NullableMyPermissionsVM) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMyPermissionsVM(val *MyPermissionsVM) *NullableMyPermissionsVM {
	return &NullableMyPermissionsVM{value: val, isSet: true}
}

func (v NullableMyPermissionsVM) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMyPermissionsVM) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


