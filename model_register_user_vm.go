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

// checks if the RegisterUserVM type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RegisterUserVM{}

// RegisterUserVM struct for RegisterUserVM
type RegisterUserVM struct {
	CloudBaseVM
	Id NullableString `json:"id,omitempty"`
	Subscriptions []string `json:"subscriptions,omitempty"`
	Groups []string `json:"groups,omitempty"`
	AdminPermission *AdminPermission `json:"adminPermission,omitempty"`
	Name NullableString `json:"name,omitempty"`
	Username NullableString `json:"username,omitempty"`
	Email NullableString `json:"email,omitempty"`
	Password NullableString `json:"password,omitempty"`
	IsAdmin *bool `json:"isAdmin,omitempty"`
	Provider NullableString `json:"provider,omitempty"`
	T string `json:"$t"`
}

type _RegisterUserVM RegisterUserVM

// NewRegisterUserVM instantiates a new RegisterUserVM object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRegisterUserVM(t string) *RegisterUserVM {
	this := RegisterUserVM{}
	this.T = t
	return &this
}

// NewRegisterUserVMWithDefaults instantiates a new RegisterUserVM object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRegisterUserVMWithDefaults() *RegisterUserVM {
	this := RegisterUserVM{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RegisterUserVM) GetId() string {
	if o == nil || IsNil(o.Id.Get()) {
		var ret string
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RegisterUserVM) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *RegisterUserVM) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableString and assigns it to the Id field.
func (o *RegisterUserVM) SetId(v string) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *RegisterUserVM) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *RegisterUserVM) UnsetId() {
	o.Id.Unset()
}

// GetSubscriptions returns the Subscriptions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RegisterUserVM) GetSubscriptions() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Subscriptions
}

// GetSubscriptionsOk returns a tuple with the Subscriptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RegisterUserVM) GetSubscriptionsOk() ([]string, bool) {
	if o == nil || IsNil(o.Subscriptions) {
		return nil, false
	}
	return o.Subscriptions, true
}

// HasSubscriptions returns a boolean if a field has been set.
func (o *RegisterUserVM) HasSubscriptions() bool {
	if o != nil && IsNil(o.Subscriptions) {
		return true
	}

	return false
}

// SetSubscriptions gets a reference to the given []string and assigns it to the Subscriptions field.
func (o *RegisterUserVM) SetSubscriptions(v []string) {
	o.Subscriptions = v
}

// GetGroups returns the Groups field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RegisterUserVM) GetGroups() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Groups
}

// GetGroupsOk returns a tuple with the Groups field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RegisterUserVM) GetGroupsOk() ([]string, bool) {
	if o == nil || IsNil(o.Groups) {
		return nil, false
	}
	return o.Groups, true
}

// HasGroups returns a boolean if a field has been set.
func (o *RegisterUserVM) HasGroups() bool {
	if o != nil && IsNil(o.Groups) {
		return true
	}

	return false
}

// SetGroups gets a reference to the given []string and assigns it to the Groups field.
func (o *RegisterUserVM) SetGroups(v []string) {
	o.Groups = v
}

// GetAdminPermission returns the AdminPermission field value if set, zero value otherwise.
func (o *RegisterUserVM) GetAdminPermission() AdminPermission {
	if o == nil || IsNil(o.AdminPermission) {
		var ret AdminPermission
		return ret
	}
	return *o.AdminPermission
}

// GetAdminPermissionOk returns a tuple with the AdminPermission field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegisterUserVM) GetAdminPermissionOk() (*AdminPermission, bool) {
	if o == nil || IsNil(o.AdminPermission) {
		return nil, false
	}
	return o.AdminPermission, true
}

// HasAdminPermission returns a boolean if a field has been set.
func (o *RegisterUserVM) HasAdminPermission() bool {
	if o != nil && !IsNil(o.AdminPermission) {
		return true
	}

	return false
}

// SetAdminPermission gets a reference to the given AdminPermission and assigns it to the AdminPermission field.
func (o *RegisterUserVM) SetAdminPermission(v AdminPermission) {
	o.AdminPermission = &v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RegisterUserVM) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RegisterUserVM) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *RegisterUserVM) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *RegisterUserVM) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *RegisterUserVM) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *RegisterUserVM) UnsetName() {
	o.Name.Unset()
}

// GetUsername returns the Username field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RegisterUserVM) GetUsername() string {
	if o == nil || IsNil(o.Username.Get()) {
		var ret string
		return ret
	}
	return *o.Username.Get()
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RegisterUserVM) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Username.Get(), o.Username.IsSet()
}

// HasUsername returns a boolean if a field has been set.
func (o *RegisterUserVM) HasUsername() bool {
	if o != nil && o.Username.IsSet() {
		return true
	}

	return false
}

// SetUsername gets a reference to the given NullableString and assigns it to the Username field.
func (o *RegisterUserVM) SetUsername(v string) {
	o.Username.Set(&v)
}
// SetUsernameNil sets the value for Username to be an explicit nil
func (o *RegisterUserVM) SetUsernameNil() {
	o.Username.Set(nil)
}

// UnsetUsername ensures that no value is present for Username, not even an explicit nil
func (o *RegisterUserVM) UnsetUsername() {
	o.Username.Unset()
}

// GetEmail returns the Email field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RegisterUserVM) GetEmail() string {
	if o == nil || IsNil(o.Email.Get()) {
		var ret string
		return ret
	}
	return *o.Email.Get()
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RegisterUserVM) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Email.Get(), o.Email.IsSet()
}

// HasEmail returns a boolean if a field has been set.
func (o *RegisterUserVM) HasEmail() bool {
	if o != nil && o.Email.IsSet() {
		return true
	}

	return false
}

// SetEmail gets a reference to the given NullableString and assigns it to the Email field.
func (o *RegisterUserVM) SetEmail(v string) {
	o.Email.Set(&v)
}
// SetEmailNil sets the value for Email to be an explicit nil
func (o *RegisterUserVM) SetEmailNil() {
	o.Email.Set(nil)
}

// UnsetEmail ensures that no value is present for Email, not even an explicit nil
func (o *RegisterUserVM) UnsetEmail() {
	o.Email.Unset()
}

// GetPassword returns the Password field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RegisterUserVM) GetPassword() string {
	if o == nil || IsNil(o.Password.Get()) {
		var ret string
		return ret
	}
	return *o.Password.Get()
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RegisterUserVM) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Password.Get(), o.Password.IsSet()
}

// HasPassword returns a boolean if a field has been set.
func (o *RegisterUserVM) HasPassword() bool {
	if o != nil && o.Password.IsSet() {
		return true
	}

	return false
}

// SetPassword gets a reference to the given NullableString and assigns it to the Password field.
func (o *RegisterUserVM) SetPassword(v string) {
	o.Password.Set(&v)
}
// SetPasswordNil sets the value for Password to be an explicit nil
func (o *RegisterUserVM) SetPasswordNil() {
	o.Password.Set(nil)
}

// UnsetPassword ensures that no value is present for Password, not even an explicit nil
func (o *RegisterUserVM) UnsetPassword() {
	o.Password.Unset()
}

// GetIsAdmin returns the IsAdmin field value if set, zero value otherwise.
func (o *RegisterUserVM) GetIsAdmin() bool {
	if o == nil || IsNil(o.IsAdmin) {
		var ret bool
		return ret
	}
	return *o.IsAdmin
}

// GetIsAdminOk returns a tuple with the IsAdmin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegisterUserVM) GetIsAdminOk() (*bool, bool) {
	if o == nil || IsNil(o.IsAdmin) {
		return nil, false
	}
	return o.IsAdmin, true
}

// HasIsAdmin returns a boolean if a field has been set.
func (o *RegisterUserVM) HasIsAdmin() bool {
	if o != nil && !IsNil(o.IsAdmin) {
		return true
	}

	return false
}

// SetIsAdmin gets a reference to the given bool and assigns it to the IsAdmin field.
func (o *RegisterUserVM) SetIsAdmin(v bool) {
	o.IsAdmin = &v
}

// GetProvider returns the Provider field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RegisterUserVM) GetProvider() string {
	if o == nil || IsNil(o.Provider.Get()) {
		var ret string
		return ret
	}
	return *o.Provider.Get()
}

// GetProviderOk returns a tuple with the Provider field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RegisterUserVM) GetProviderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Provider.Get(), o.Provider.IsSet()
}

// HasProvider returns a boolean if a field has been set.
func (o *RegisterUserVM) HasProvider() bool {
	if o != nil && o.Provider.IsSet() {
		return true
	}

	return false
}

// SetProvider gets a reference to the given NullableString and assigns it to the Provider field.
func (o *RegisterUserVM) SetProvider(v string) {
	o.Provider.Set(&v)
}
// SetProviderNil sets the value for Provider to be an explicit nil
func (o *RegisterUserVM) SetProviderNil() {
	o.Provider.Set(nil)
}

// UnsetProvider ensures that no value is present for Provider, not even an explicit nil
func (o *RegisterUserVM) UnsetProvider() {
	o.Provider.Unset()
}

// GetT returns the T field value
func (o *RegisterUserVM) GetT() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.T
}

// GetTOk returns a tuple with the T field value
// and a boolean to check if the value has been set.
func (o *RegisterUserVM) GetTOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.T, true
}

// SetT sets field value
func (o *RegisterUserVM) SetT(v string) {
	o.T = v
}

func (o RegisterUserVM) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RegisterUserVM) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedCloudBaseVM, errCloudBaseVM := json.Marshal(o.CloudBaseVM)
	if errCloudBaseVM != nil {
		return map[string]interface{}{}, errCloudBaseVM
	}
	errCloudBaseVM = json.Unmarshal([]byte(serializedCloudBaseVM), &toSerialize)
	if errCloudBaseVM != nil {
		return map[string]interface{}{}, errCloudBaseVM
	}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if o.Subscriptions != nil {
		toSerialize["subscriptions"] = o.Subscriptions
	}
	if o.Groups != nil {
		toSerialize["groups"] = o.Groups
	}
	if !IsNil(o.AdminPermission) {
		toSerialize["adminPermission"] = o.AdminPermission
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.Username.IsSet() {
		toSerialize["username"] = o.Username.Get()
	}
	if o.Email.IsSet() {
		toSerialize["email"] = o.Email.Get()
	}
	if o.Password.IsSet() {
		toSerialize["password"] = o.Password.Get()
	}
	if !IsNil(o.IsAdmin) {
		toSerialize["isAdmin"] = o.IsAdmin
	}
	if o.Provider.IsSet() {
		toSerialize["provider"] = o.Provider.Get()
	}
	toSerialize["$t"] = o.T
	return toSerialize, nil
}

func (o *RegisterUserVM) UnmarshalJSON(data []byte) (err error) {
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

	varRegisterUserVM := _RegisterUserVM{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRegisterUserVM)

	if err != nil {
		return err
	}

	*o = RegisterUserVM(varRegisterUserVM)

	return err
}

type NullableRegisterUserVM struct {
	value *RegisterUserVM
	isSet bool
}

func (v NullableRegisterUserVM) Get() *RegisterUserVM {
	return v.value
}

func (v *NullableRegisterUserVM) Set(val *RegisterUserVM) {
	v.value = val
	v.isSet = true
}

func (v NullableRegisterUserVM) IsSet() bool {
	return v.isSet
}

func (v *NullableRegisterUserVM) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRegisterUserVM(val *RegisterUserVM) *NullableRegisterUserVM {
	return &NullableRegisterUserVM{value: val, isSet: true}
}

func (v NullableRegisterUserVM) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRegisterUserVM) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

