/*
FastReport Cloud

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gofrcloud

import (
	"encoding/json"
)

// checks if the CreateDataSourceVM type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateDataSourceVM{}

// CreateDataSourceVM struct for CreateDataSourceVM
type CreateDataSourceVM struct {
	Name NullableString `json:"name,omitempty"`
	ConnectionString string `json:"connectionString"`
	SubscriptionId string `json:"subscriptionId"`
	ConnectionType *DataSourceConnectionType `json:"connectionType,omitempty"`
}

// NewCreateDataSourceVM instantiates a new CreateDataSourceVM object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateDataSourceVM(connectionString string, subscriptionId string) *CreateDataSourceVM {
	this := CreateDataSourceVM{}
	this.ConnectionString = connectionString
	this.SubscriptionId = subscriptionId
	return &this
}

// NewCreateDataSourceVMWithDefaults instantiates a new CreateDataSourceVM object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateDataSourceVMWithDefaults() *CreateDataSourceVM {
	this := CreateDataSourceVM{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateDataSourceVM) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateDataSourceVM) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *CreateDataSourceVM) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *CreateDataSourceVM) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *CreateDataSourceVM) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *CreateDataSourceVM) UnsetName() {
	o.Name.Unset()
}

// GetConnectionString returns the ConnectionString field value
func (o *CreateDataSourceVM) GetConnectionString() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ConnectionString
}

// GetConnectionStringOk returns a tuple with the ConnectionString field value
// and a boolean to check if the value has been set.
func (o *CreateDataSourceVM) GetConnectionStringOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConnectionString, true
}

// SetConnectionString sets field value
func (o *CreateDataSourceVM) SetConnectionString(v string) {
	o.ConnectionString = v
}

// GetSubscriptionId returns the SubscriptionId field value
func (o *CreateDataSourceVM) GetSubscriptionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SubscriptionId
}

// GetSubscriptionIdOk returns a tuple with the SubscriptionId field value
// and a boolean to check if the value has been set.
func (o *CreateDataSourceVM) GetSubscriptionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SubscriptionId, true
}

// SetSubscriptionId sets field value
func (o *CreateDataSourceVM) SetSubscriptionId(v string) {
	o.SubscriptionId = v
}

// GetConnectionType returns the ConnectionType field value if set, zero value otherwise.
func (o *CreateDataSourceVM) GetConnectionType() DataSourceConnectionType {
	if o == nil || IsNil(o.ConnectionType) {
		var ret DataSourceConnectionType
		return ret
	}
	return *o.ConnectionType
}

// GetConnectionTypeOk returns a tuple with the ConnectionType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDataSourceVM) GetConnectionTypeOk() (*DataSourceConnectionType, bool) {
	if o == nil || IsNil(o.ConnectionType) {
		return nil, false
	}
	return o.ConnectionType, true
}

// HasConnectionType returns a boolean if a field has been set.
func (o *CreateDataSourceVM) HasConnectionType() bool {
	if o != nil && !IsNil(o.ConnectionType) {
		return true
	}

	return false
}

// SetConnectionType gets a reference to the given DataSourceConnectionType and assigns it to the ConnectionType field.
func (o *CreateDataSourceVM) SetConnectionType(v DataSourceConnectionType) {
	o.ConnectionType = &v
}

func (o CreateDataSourceVM) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateDataSourceVM) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	toSerialize["connectionString"] = o.ConnectionString
	toSerialize["subscriptionId"] = o.SubscriptionId
	if !IsNil(o.ConnectionType) {
		toSerialize["connectionType"] = o.ConnectionType
	}
	return toSerialize, nil
}

type NullableCreateDataSourceVM struct {
	value *CreateDataSourceVM
	isSet bool
}

func (v NullableCreateDataSourceVM) Get() *CreateDataSourceVM {
	return v.value
}

func (v *NullableCreateDataSourceVM) Set(val *CreateDataSourceVM) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateDataSourceVM) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateDataSourceVM) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateDataSourceVM(val *CreateDataSourceVM) *NullableCreateDataSourceVM {
	return &NullableCreateDataSourceVM{value: val, isSet: true}
}

func (v NullableCreateDataSourceVM) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateDataSourceVM) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


