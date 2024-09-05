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

// checks if the UpdateUserVM type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateUserVM{}

// UpdateUserVM struct for UpdateUserVM
type UpdateUserVM struct {
	CloudBaseVM
	AdminPermission *AdminPermission `json:"adminPermission,omitempty"`
	Email NullableString `json:"email,omitempty"`
	Groups []string `json:"groups,omitempty"`
	IsAdmin NullableBool `json:"isAdmin,omitempty"`
	Name NullableString `json:"name,omitempty"`
	Password NullableString `json:"password,omitempty"`
	Provider NullableString `json:"provider,omitempty"`
	Subscriptions []string `json:"subscriptions,omitempty"`
	Username NullableString `json:"username,omitempty"`
	T string `json:"$t"`
}

type _UpdateUserVM UpdateUserVM

// NewUpdateUserVM instantiates a new UpdateUserVM object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateUserVM(t string) *UpdateUserVM {
	this := UpdateUserVM{}
	this.T = t
	return &this
}

// NewUpdateUserVMWithDefaults instantiates a new UpdateUserVM object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateUserVMWithDefaults() *UpdateUserVM {
	this := UpdateUserVM{}
	return &this
}

// GetAdminPermission returns the AdminPermission field value if set, zero value otherwise.
func (o *UpdateUserVM) GetAdminPermission() AdminPermission {
	if o == nil || IsNil(o.AdminPermission) {
		var ret AdminPermission
		return ret
	}
	return *o.AdminPermission
}

// GetAdminPermissionOk returns a tuple with the AdminPermission field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateUserVM) GetAdminPermissionOk() (*AdminPermission, bool) {
	if o == nil || IsNil(o.AdminPermission) {
		return nil, false
	}
	return o.AdminPermission, true
}

// HasAdminPermission returns a boolean if a field has been set.
func (o *UpdateUserVM) HasAdminPermission() bool {
	if o != nil && !IsNil(o.AdminPermission) {
		return true
	}

	return false
}

// SetAdminPermission gets a reference to the given AdminPermission and assigns it to the AdminPermission field.
func (o *UpdateUserVM) SetAdminPermission(v AdminPermission) {
	o.AdminPermission = &v
}

// GetEmail returns the Email field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateUserVM) GetEmail() string {
	if o == nil || IsNil(o.Email.Get()) {
		var ret string
		return ret
	}
	return *o.Email.Get()
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateUserVM) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Email.Get(), o.Email.IsSet()
}

// HasEmail returns a boolean if a field has been set.
func (o *UpdateUserVM) HasEmail() bool {
	if o != nil && o.Email.IsSet() {
		return true
	}

	return false
}

// SetEmail gets a reference to the given NullableString and assigns it to the Email field.
func (o *UpdateUserVM) SetEmail(v string) {
	o.Email.Set(&v)
}
// SetEmailNil sets the value for Email to be an explicit nil
func (o *UpdateUserVM) SetEmailNil() {
	o.Email.Set(nil)
}

// UnsetEmail ensures that no value is present for Email, not even an explicit nil
func (o *UpdateUserVM) UnsetEmail() {
	o.Email.Unset()
}

// GetGroups returns the Groups field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateUserVM) GetGroups() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Groups
}

// GetGroupsOk returns a tuple with the Groups field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateUserVM) GetGroupsOk() ([]string, bool) {
	if o == nil || IsNil(o.Groups) {
		return nil, false
	}
	return o.Groups, true
}

// HasGroups returns a boolean if a field has been set.
func (o *UpdateUserVM) HasGroups() bool {
	if o != nil && IsNil(o.Groups) {
		return true
	}

	return false
}

// SetGroups gets a reference to the given []string and assigns it to the Groups field.
func (o *UpdateUserVM) SetGroups(v []string) {
	o.Groups = v
}

// GetIsAdmin returns the IsAdmin field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateUserVM) GetIsAdmin() bool {
	if o == nil || IsNil(o.IsAdmin.Get()) {
		var ret bool
		return ret
	}
	return *o.IsAdmin.Get()
}

// GetIsAdminOk returns a tuple with the IsAdmin field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateUserVM) GetIsAdminOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsAdmin.Get(), o.IsAdmin.IsSet()
}

// HasIsAdmin returns a boolean if a field has been set.
func (o *UpdateUserVM) HasIsAdmin() bool {
	if o != nil && o.IsAdmin.IsSet() {
		return true
	}

	return false
}

// SetIsAdmin gets a reference to the given NullableBool and assigns it to the IsAdmin field.
func (o *UpdateUserVM) SetIsAdmin(v bool) {
	o.IsAdmin.Set(&v)
}
// SetIsAdminNil sets the value for IsAdmin to be an explicit nil
func (o *UpdateUserVM) SetIsAdminNil() {
	o.IsAdmin.Set(nil)
}

// UnsetIsAdmin ensures that no value is present for IsAdmin, not even an explicit nil
func (o *UpdateUserVM) UnsetIsAdmin() {
	o.IsAdmin.Unset()
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateUserVM) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateUserVM) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *UpdateUserVM) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *UpdateUserVM) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *UpdateUserVM) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *UpdateUserVM) UnsetName() {
	o.Name.Unset()
}

// GetPassword returns the Password field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateUserVM) GetPassword() string {
	if o == nil || IsNil(o.Password.Get()) {
		var ret string
		return ret
	}
	return *o.Password.Get()
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateUserVM) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Password.Get(), o.Password.IsSet()
}

// HasPassword returns a boolean if a field has been set.
func (o *UpdateUserVM) HasPassword() bool {
	if o != nil && o.Password.IsSet() {
		return true
	}

	return false
}

// SetPassword gets a reference to the given NullableString and assigns it to the Password field.
func (o *UpdateUserVM) SetPassword(v string) {
	o.Password.Set(&v)
}
// SetPasswordNil sets the value for Password to be an explicit nil
func (o *UpdateUserVM) SetPasswordNil() {
	o.Password.Set(nil)
}

// UnsetPassword ensures that no value is present for Password, not even an explicit nil
func (o *UpdateUserVM) UnsetPassword() {
	o.Password.Unset()
}

// GetProvider returns the Provider field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateUserVM) GetProvider() string {
	if o == nil || IsNil(o.Provider.Get()) {
		var ret string
		return ret
	}
	return *o.Provider.Get()
}

// GetProviderOk returns a tuple with the Provider field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateUserVM) GetProviderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Provider.Get(), o.Provider.IsSet()
}

// HasProvider returns a boolean if a field has been set.
func (o *UpdateUserVM) HasProvider() bool {
	if o != nil && o.Provider.IsSet() {
		return true
	}

	return false
}

// SetProvider gets a reference to the given NullableString and assigns it to the Provider field.
func (o *UpdateUserVM) SetProvider(v string) {
	o.Provider.Set(&v)
}
// SetProviderNil sets the value for Provider to be an explicit nil
func (o *UpdateUserVM) SetProviderNil() {
	o.Provider.Set(nil)
}

// UnsetProvider ensures that no value is present for Provider, not even an explicit nil
func (o *UpdateUserVM) UnsetProvider() {
	o.Provider.Unset()
}

// GetSubscriptions returns the Subscriptions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateUserVM) GetSubscriptions() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Subscriptions
}

// GetSubscriptionsOk returns a tuple with the Subscriptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateUserVM) GetSubscriptionsOk() ([]string, bool) {
	if o == nil || IsNil(o.Subscriptions) {
		return nil, false
	}
	return o.Subscriptions, true
}

// HasSubscriptions returns a boolean if a field has been set.
func (o *UpdateUserVM) HasSubscriptions() bool {
	if o != nil && IsNil(o.Subscriptions) {
		return true
	}

	return false
}

// SetSubscriptions gets a reference to the given []string and assigns it to the Subscriptions field.
func (o *UpdateUserVM) SetSubscriptions(v []string) {
	o.Subscriptions = v
}

// GetUsername returns the Username field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateUserVM) GetUsername() string {
	if o == nil || IsNil(o.Username.Get()) {
		var ret string
		return ret
	}
	return *o.Username.Get()
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateUserVM) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Username.Get(), o.Username.IsSet()
}

// HasUsername returns a boolean if a field has been set.
func (o *UpdateUserVM) HasUsername() bool {
	if o != nil && o.Username.IsSet() {
		return true
	}

	return false
}

// SetUsername gets a reference to the given NullableString and assigns it to the Username field.
func (o *UpdateUserVM) SetUsername(v string) {
	o.Username.Set(&v)
}
// SetUsernameNil sets the value for Username to be an explicit nil
func (o *UpdateUserVM) SetUsernameNil() {
	o.Username.Set(nil)
}

// UnsetUsername ensures that no value is present for Username, not even an explicit nil
func (o *UpdateUserVM) UnsetUsername() {
	o.Username.Unset()
}

// GetT returns the T field value
func (o *UpdateUserVM) GetT() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.T
}

// GetTOk returns a tuple with the T field value
// and a boolean to check if the value has been set.
func (o *UpdateUserVM) GetTOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.T, true
}

// SetT sets field value
func (o *UpdateUserVM) SetT(v string) {
	o.T = v
}

func (o UpdateUserVM) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateUserVM) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedCloudBaseVM, errCloudBaseVM := json.Marshal(o.CloudBaseVM)
	if errCloudBaseVM != nil {
		return map[string]interface{}{}, errCloudBaseVM
	}
	errCloudBaseVM = json.Unmarshal([]byte(serializedCloudBaseVM), &toSerialize)
	if errCloudBaseVM != nil {
		return map[string]interface{}{}, errCloudBaseVM
	}
	if !IsNil(o.AdminPermission) {
		toSerialize["adminPermission"] = o.AdminPermission
	}
	if o.Email.IsSet() {
		toSerialize["email"] = o.Email.Get()
	}
	if o.Groups != nil {
		toSerialize["groups"] = o.Groups
	}
	if o.IsAdmin.IsSet() {
		toSerialize["isAdmin"] = o.IsAdmin.Get()
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.Password.IsSet() {
		toSerialize["password"] = o.Password.Get()
	}
	if o.Provider.IsSet() {
		toSerialize["provider"] = o.Provider.Get()
	}
	if o.Subscriptions != nil {
		toSerialize["subscriptions"] = o.Subscriptions
	}
	if o.Username.IsSet() {
		toSerialize["username"] = o.Username.Get()
	}
	toSerialize["$t"] = o.T
	return toSerialize, nil
}

func (o *UpdateUserVM) UnmarshalJSON(data []byte) (err error) {
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

	varUpdateUserVM := _UpdateUserVM{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUpdateUserVM)

	if err != nil {
		return err
	}

	*o = UpdateUserVM(varUpdateUserVM)

	return err
}

type NullableUpdateUserVM struct {
	value *UpdateUserVM
	isSet bool
}

func (v NullableUpdateUserVM) Get() *UpdateUserVM {
	return v.value
}

func (v *NullableUpdateUserVM) Set(val *UpdateUserVM) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateUserVM) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateUserVM) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateUserVM(val *UpdateUserVM) *NullableUpdateUserVM {
	return &NullableUpdateUserVM{value: val, isSet: true}
}

func (v NullableUpdateUserVM) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateUserVM) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

