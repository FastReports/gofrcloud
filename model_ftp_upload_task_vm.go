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

// checks if the FTPUploadTaskVM type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FTPUploadTaskVM{}

// FTPUploadTaskVM struct for FTPUploadTaskVM
type FTPUploadTaskVM struct {
	Archive *bool `json:"archive,omitempty"`
	ArchiveName NullableString `json:"archiveName,omitempty"`
	DestinationFolder NullableString `json:"destinationFolder,omitempty"`
	FtpHost NullableString `json:"ftpHost,omitempty"`
	FtpPort *int32 `json:"ftpPort,omitempty"`
	FtpUsername NullableString `json:"ftpUsername,omitempty"`
	UseSFTP *bool `json:"useSFTP,omitempty"`
}

// NewFTPUploadTaskVM instantiates a new FTPUploadTaskVM object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFTPUploadTaskVM(t string) *FTPUploadTaskVM {
	this := FTPUploadTaskVM{}
	this.T = t
	return &this
}

// NewFTPUploadTaskVMWithDefaults instantiates a new FTPUploadTaskVM object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFTPUploadTaskVMWithDefaults() *FTPUploadTaskVM {
	this := FTPUploadTaskVM{}
	return &this
}

// GetArchive returns the Archive field value if set, zero value otherwise.
func (o *FTPUploadTaskVM) GetArchive() bool {
	if o == nil || IsNil(o.Archive) {
		var ret bool
		return ret
	}
	return *o.Archive
}

// GetArchiveOk returns a tuple with the Archive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FTPUploadTaskVM) GetArchiveOk() (*bool, bool) {
	if o == nil || IsNil(o.Archive) {
		return nil, false
	}
	return o.Archive, true
}

// HasArchive returns a boolean if a field has been set.
func (o *FTPUploadTaskVM) HasArchive() bool {
	if o != nil && !IsNil(o.Archive) {
		return true
	}

	return false
}

// SetArchive gets a reference to the given bool and assigns it to the Archive field.
func (o *FTPUploadTaskVM) SetArchive(v bool) {
	o.Archive = &v
}

// GetArchiveName returns the ArchiveName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FTPUploadTaskVM) GetArchiveName() string {
	if o == nil || IsNil(o.ArchiveName.Get()) {
		var ret string
		return ret
	}
	return *o.ArchiveName.Get()
}

// GetArchiveNameOk returns a tuple with the ArchiveName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FTPUploadTaskVM) GetArchiveNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ArchiveName.Get(), o.ArchiveName.IsSet()
}

// HasArchiveName returns a boolean if a field has been set.
func (o *FTPUploadTaskVM) HasArchiveName() bool {
	if o != nil && o.ArchiveName.IsSet() {
		return true
	}

	return false
}

// SetArchiveName gets a reference to the given NullableString and assigns it to the ArchiveName field.
func (o *FTPUploadTaskVM) SetArchiveName(v string) {
	o.ArchiveName.Set(&v)
}
// SetArchiveNameNil sets the value for ArchiveName to be an explicit nil
func (o *FTPUploadTaskVM) SetArchiveNameNil() {
	o.ArchiveName.Set(nil)
}

// UnsetArchiveName ensures that no value is present for ArchiveName, not even an explicit nil
func (o *FTPUploadTaskVM) UnsetArchiveName() {
	o.ArchiveName.Unset()
}

// GetDestinationFolder returns the DestinationFolder field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FTPUploadTaskVM) GetDestinationFolder() string {
	if o == nil || IsNil(o.DestinationFolder.Get()) {
		var ret string
		return ret
	}
	return *o.DestinationFolder.Get()
}

// GetDestinationFolderOk returns a tuple with the DestinationFolder field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FTPUploadTaskVM) GetDestinationFolderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DestinationFolder.Get(), o.DestinationFolder.IsSet()
}

// HasDestinationFolder returns a boolean if a field has been set.
func (o *FTPUploadTaskVM) HasDestinationFolder() bool {
	if o != nil && o.DestinationFolder.IsSet() {
		return true
	}

	return false
}

// SetDestinationFolder gets a reference to the given NullableString and assigns it to the DestinationFolder field.
func (o *FTPUploadTaskVM) SetDestinationFolder(v string) {
	o.DestinationFolder.Set(&v)
}
// SetDestinationFolderNil sets the value for DestinationFolder to be an explicit nil
func (o *FTPUploadTaskVM) SetDestinationFolderNil() {
	o.DestinationFolder.Set(nil)
}

// UnsetDestinationFolder ensures that no value is present for DestinationFolder, not even an explicit nil
func (o *FTPUploadTaskVM) UnsetDestinationFolder() {
	o.DestinationFolder.Unset()
}

// GetFtpHost returns the FtpHost field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FTPUploadTaskVM) GetFtpHost() string {
	if o == nil || IsNil(o.FtpHost.Get()) {
		var ret string
		return ret
	}
	return *o.FtpHost.Get()
}

// GetFtpHostOk returns a tuple with the FtpHost field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FTPUploadTaskVM) GetFtpHostOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FtpHost.Get(), o.FtpHost.IsSet()
}

// HasFtpHost returns a boolean if a field has been set.
func (o *FTPUploadTaskVM) HasFtpHost() bool {
	if o != nil && o.FtpHost.IsSet() {
		return true
	}

	return false
}

// SetFtpHost gets a reference to the given NullableString and assigns it to the FtpHost field.
func (o *FTPUploadTaskVM) SetFtpHost(v string) {
	o.FtpHost.Set(&v)
}
// SetFtpHostNil sets the value for FtpHost to be an explicit nil
func (o *FTPUploadTaskVM) SetFtpHostNil() {
	o.FtpHost.Set(nil)
}

// UnsetFtpHost ensures that no value is present for FtpHost, not even an explicit nil
func (o *FTPUploadTaskVM) UnsetFtpHost() {
	o.FtpHost.Unset()
}

// GetFtpPort returns the FtpPort field value if set, zero value otherwise.
func (o *FTPUploadTaskVM) GetFtpPort() int32 {
	if o == nil || IsNil(o.FtpPort) {
		var ret int32
		return ret
	}
	return *o.FtpPort
}

// GetFtpPortOk returns a tuple with the FtpPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FTPUploadTaskVM) GetFtpPortOk() (*int32, bool) {
	if o == nil || IsNil(o.FtpPort) {
		return nil, false
	}
	return o.FtpPort, true
}

// HasFtpPort returns a boolean if a field has been set.
func (o *FTPUploadTaskVM) HasFtpPort() bool {
	if o != nil && !IsNil(o.FtpPort) {
		return true
	}

	return false
}

// SetFtpPort gets a reference to the given int32 and assigns it to the FtpPort field.
func (o *FTPUploadTaskVM) SetFtpPort(v int32) {
	o.FtpPort = &v
}

// GetFtpUsername returns the FtpUsername field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FTPUploadTaskVM) GetFtpUsername() string {
	if o == nil || IsNil(o.FtpUsername.Get()) {
		var ret string
		return ret
	}
	return *o.FtpUsername.Get()
}

// GetFtpUsernameOk returns a tuple with the FtpUsername field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FTPUploadTaskVM) GetFtpUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FtpUsername.Get(), o.FtpUsername.IsSet()
}

// HasFtpUsername returns a boolean if a field has been set.
func (o *FTPUploadTaskVM) HasFtpUsername() bool {
	if o != nil && o.FtpUsername.IsSet() {
		return true
	}

	return false
}

// SetFtpUsername gets a reference to the given NullableString and assigns it to the FtpUsername field.
func (o *FTPUploadTaskVM) SetFtpUsername(v string) {
	o.FtpUsername.Set(&v)
}
// SetFtpUsernameNil sets the value for FtpUsername to be an explicit nil
func (o *FTPUploadTaskVM) SetFtpUsernameNil() {
	o.FtpUsername.Set(nil)
}

// UnsetFtpUsername ensures that no value is present for FtpUsername, not even an explicit nil
func (o *FTPUploadTaskVM) UnsetFtpUsername() {
	o.FtpUsername.Unset()
}

// GetUseSFTP returns the UseSFTP field value if set, zero value otherwise.
func (o *FTPUploadTaskVM) GetUseSFTP() bool {
	if o == nil || IsNil(o.UseSFTP) {
		var ret bool
		return ret
	}
	return *o.UseSFTP
}

// GetUseSFTPOk returns a tuple with the UseSFTP field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FTPUploadTaskVM) GetUseSFTPOk() (*bool, bool) {
	if o == nil || IsNil(o.UseSFTP) {
		return nil, false
	}
	return o.UseSFTP, true
}

// HasUseSFTP returns a boolean if a field has been set.
func (o *FTPUploadTaskVM) HasUseSFTP() bool {
	if o != nil && !IsNil(o.UseSFTP) {
		return true
	}

	return false
}

// SetUseSFTP gets a reference to the given bool and assigns it to the UseSFTP field.
func (o *FTPUploadTaskVM) SetUseSFTP(v bool) {
	o.UseSFTP = &v
}

func (o FTPUploadTaskVM) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FTPUploadTaskVM) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Archive) {
		toSerialize["archive"] = o.Archive
	}
	if o.ArchiveName.IsSet() {
		toSerialize["archiveName"] = o.ArchiveName.Get()
	}
	if o.DestinationFolder.IsSet() {
		toSerialize["destinationFolder"] = o.DestinationFolder.Get()
	}
	if o.FtpHost.IsSet() {
		toSerialize["ftpHost"] = o.FtpHost.Get()
	}
	if !IsNil(o.FtpPort) {
		toSerialize["ftpPort"] = o.FtpPort
	}
	if o.FtpUsername.IsSet() {
		toSerialize["ftpUsername"] = o.FtpUsername.Get()
	}
	if !IsNil(o.UseSFTP) {
		toSerialize["useSFTP"] = o.UseSFTP
	}
	return toSerialize, nil
}

type NullableFTPUploadTaskVM struct {
	value *FTPUploadTaskVM
	isSet bool
}

func (v NullableFTPUploadTaskVM) Get() *FTPUploadTaskVM {
	return v.value
}

func (v *NullableFTPUploadTaskVM) Set(val *FTPUploadTaskVM) {
	v.value = val
	v.isSet = true
}

func (v NullableFTPUploadTaskVM) IsSet() bool {
	return v.isSet
}

func (v *NullableFTPUploadTaskVM) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFTPUploadTaskVM(val *FTPUploadTaskVM) *NullableFTPUploadTaskVM {
	return &NullableFTPUploadTaskVM{value: val, isSet: true}
}

func (v NullableFTPUploadTaskVM) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFTPUploadTaskVM) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

