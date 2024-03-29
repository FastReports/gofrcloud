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

// checks if the FileVMFilesVMBase type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FileVMFilesVMBase{}

// FileVMFilesVMBase struct for FileVMFilesVMBase
type FileVMFilesVMBase struct {
	Files []FileVM `json:"files,omitempty"`
	Count *int64 `json:"count,omitempty"`
	Skip *int32 `json:"skip,omitempty"`
	Take *int32 `json:"take,omitempty"`
}

// NewFileVMFilesVMBase instantiates a new FileVMFilesVMBase object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFileVMFilesVMBase() *FileVMFilesVMBase {
	this := FileVMFilesVMBase{}
	return &this
}

// NewFileVMFilesVMBaseWithDefaults instantiates a new FileVMFilesVMBase object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFileVMFilesVMBaseWithDefaults() *FileVMFilesVMBase {
	this := FileVMFilesVMBase{}
	return &this
}

// GetFiles returns the Files field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FileVMFilesVMBase) GetFiles() []FileVM {
	if o == nil {
		var ret []FileVM
		return ret
	}
	return o.Files
}

// GetFilesOk returns a tuple with the Files field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FileVMFilesVMBase) GetFilesOk() ([]FileVM, bool) {
	if o == nil || IsNil(o.Files) {
		return nil, false
	}
	return o.Files, true
}

// HasFiles returns a boolean if a field has been set.
func (o *FileVMFilesVMBase) HasFiles() bool {
	if o != nil && IsNil(o.Files) {
		return true
	}

	return false
}

// SetFiles gets a reference to the given []FileVM and assigns it to the Files field.
func (o *FileVMFilesVMBase) SetFiles(v []FileVM) {
	o.Files = v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *FileVMFilesVMBase) GetCount() int64 {
	if o == nil || IsNil(o.Count) {
		var ret int64
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileVMFilesVMBase) GetCountOk() (*int64, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *FileVMFilesVMBase) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int64 and assigns it to the Count field.
func (o *FileVMFilesVMBase) SetCount(v int64) {
	o.Count = &v
}

// GetSkip returns the Skip field value if set, zero value otherwise.
func (o *FileVMFilesVMBase) GetSkip() int32 {
	if o == nil || IsNil(o.Skip) {
		var ret int32
		return ret
	}
	return *o.Skip
}

// GetSkipOk returns a tuple with the Skip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileVMFilesVMBase) GetSkipOk() (*int32, bool) {
	if o == nil || IsNil(o.Skip) {
		return nil, false
	}
	return o.Skip, true
}

// HasSkip returns a boolean if a field has been set.
func (o *FileVMFilesVMBase) HasSkip() bool {
	if o != nil && !IsNil(o.Skip) {
		return true
	}

	return false
}

// SetSkip gets a reference to the given int32 and assigns it to the Skip field.
func (o *FileVMFilesVMBase) SetSkip(v int32) {
	o.Skip = &v
}

// GetTake returns the Take field value if set, zero value otherwise.
func (o *FileVMFilesVMBase) GetTake() int32 {
	if o == nil || IsNil(o.Take) {
		var ret int32
		return ret
	}
	return *o.Take
}

// GetTakeOk returns a tuple with the Take field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileVMFilesVMBase) GetTakeOk() (*int32, bool) {
	if o == nil || IsNil(o.Take) {
		return nil, false
	}
	return o.Take, true
}

// HasTake returns a boolean if a field has been set.
func (o *FileVMFilesVMBase) HasTake() bool {
	if o != nil && !IsNil(o.Take) {
		return true
	}

	return false
}

// SetTake gets a reference to the given int32 and assigns it to the Take field.
func (o *FileVMFilesVMBase) SetTake(v int32) {
	o.Take = &v
}

func (o FileVMFilesVMBase) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FileVMFilesVMBase) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Files != nil {
		toSerialize["files"] = o.Files
	}
	if !IsNil(o.Count) {
		toSerialize["count"] = o.Count
	}
	if !IsNil(o.Skip) {
		toSerialize["skip"] = o.Skip
	}
	if !IsNil(o.Take) {
		toSerialize["take"] = o.Take
	}
	return toSerialize, nil
}

type NullableFileVMFilesVMBase struct {
	value *FileVMFilesVMBase
	isSet bool
}

func (v NullableFileVMFilesVMBase) Get() *FileVMFilesVMBase {
	return v.value
}

func (v *NullableFileVMFilesVMBase) Set(val *FileVMFilesVMBase) {
	v.value = val
	v.isSet = true
}

func (v NullableFileVMFilesVMBase) IsSet() bool {
	return v.isSet
}

func (v *NullableFileVMFilesVMBase) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFileVMFilesVMBase(val *FileVMFilesVMBase) *NullableFileVMFilesVMBase {
	return &NullableFileVMFilesVMBase{value: val, isSet: true}
}

func (v NullableFileVMFilesVMBase) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFileVMFilesVMBase) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


