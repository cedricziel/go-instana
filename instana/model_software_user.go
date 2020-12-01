/*
 * Introduction to Instana public APIs
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.192.86
 * Contact: support@instana.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package instana

import (
	"encoding/json"
)

// SoftwareUser struct for SoftwareUser
type SoftwareUser struct {
	Host      *string `json:"host,omitempty"`
	Container *string `json:"container,omitempty"`
	Process   *string `json:"process,omitempty"`
}

// NewSoftwareUser instantiates a new SoftwareUser object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSoftwareUser() *SoftwareUser {
	this := SoftwareUser{}
	return &this
}

// NewSoftwareUserWithDefaults instantiates a new SoftwareUser object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSoftwareUserWithDefaults() *SoftwareUser {
	this := SoftwareUser{}
	return &this
}

// GetHost returns the Host field value if set, zero value otherwise.
func (o *SoftwareUser) GetHost() string {
	if o == nil || o.Host == nil {
		var ret string
		return ret
	}
	return *o.Host
}

// GetHostOk returns a tuple with the Host field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwareUser) GetHostOk() (*string, bool) {
	if o == nil || o.Host == nil {
		return nil, false
	}
	return o.Host, true
}

// HasHost returns a boolean if a field has been set.
func (o *SoftwareUser) HasHost() bool {
	if o != nil && o.Host != nil {
		return true
	}

	return false
}

// SetHost gets a reference to the given string and assigns it to the Host field.
func (o *SoftwareUser) SetHost(v string) {
	o.Host = &v
}

// GetContainer returns the Container field value if set, zero value otherwise.
func (o *SoftwareUser) GetContainer() string {
	if o == nil || o.Container == nil {
		var ret string
		return ret
	}
	return *o.Container
}

// GetContainerOk returns a tuple with the Container field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwareUser) GetContainerOk() (*string, bool) {
	if o == nil || o.Container == nil {
		return nil, false
	}
	return o.Container, true
}

// HasContainer returns a boolean if a field has been set.
func (o *SoftwareUser) HasContainer() bool {
	if o != nil && o.Container != nil {
		return true
	}

	return false
}

// SetContainer gets a reference to the given string and assigns it to the Container field.
func (o *SoftwareUser) SetContainer(v string) {
	o.Container = &v
}

// GetProcess returns the Process field value if set, zero value otherwise.
func (o *SoftwareUser) GetProcess() string {
	if o == nil || o.Process == nil {
		var ret string
		return ret
	}
	return *o.Process
}

// GetProcessOk returns a tuple with the Process field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwareUser) GetProcessOk() (*string, bool) {
	if o == nil || o.Process == nil {
		return nil, false
	}
	return o.Process, true
}

// HasProcess returns a boolean if a field has been set.
func (o *SoftwareUser) HasProcess() bool {
	if o != nil && o.Process != nil {
		return true
	}

	return false
}

// SetProcess gets a reference to the given string and assigns it to the Process field.
func (o *SoftwareUser) SetProcess(v string) {
	o.Process = &v
}

func (o SoftwareUser) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Host != nil {
		toSerialize["host"] = o.Host
	}
	if o.Container != nil {
		toSerialize["container"] = o.Container
	}
	if o.Process != nil {
		toSerialize["process"] = o.Process
	}
	return json.Marshal(toSerialize)
}

type NullableSoftwareUser struct {
	value *SoftwareUser
	isSet bool
}

func (v NullableSoftwareUser) Get() *SoftwareUser {
	return v.value
}

func (v *NullableSoftwareUser) Set(val *SoftwareUser) {
	v.value = val
	v.isSet = true
}

func (v NullableSoftwareUser) IsSet() bool {
	return v.isSet
}

func (v *NullableSoftwareUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSoftwareUser(val *SoftwareUser) *NullableSoftwareUser {
	return &NullableSoftwareUser{value: val, isSet: true}
}

func (v NullableSoftwareUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSoftwareUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
