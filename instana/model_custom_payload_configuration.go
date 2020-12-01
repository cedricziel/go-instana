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

// CustomPayloadConfiguration struct for CustomPayloadConfiguration
type CustomPayloadConfiguration struct {
	Fields []CustomPayloadField `json:"fields"`
}

// NewCustomPayloadConfiguration instantiates a new CustomPayloadConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomPayloadConfiguration(fields []CustomPayloadField) *CustomPayloadConfiguration {
	this := CustomPayloadConfiguration{}
	this.Fields = fields
	return &this
}

// NewCustomPayloadConfigurationWithDefaults instantiates a new CustomPayloadConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomPayloadConfigurationWithDefaults() *CustomPayloadConfiguration {
	this := CustomPayloadConfiguration{}
	return &this
}

// GetFields returns the Fields field value
func (o *CustomPayloadConfiguration) GetFields() []CustomPayloadField {
	if o == nil {
		var ret []CustomPayloadField
		return ret
	}

	return o.Fields
}

// GetFieldsOk returns a tuple with the Fields field value
// and a boolean to check if the value has been set.
func (o *CustomPayloadConfiguration) GetFieldsOk() (*[]CustomPayloadField, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Fields, true
}

// SetFields sets field value
func (o *CustomPayloadConfiguration) SetFields(v []CustomPayloadField) {
	o.Fields = v
}

func (o CustomPayloadConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["fields"] = o.Fields
	}
	return json.Marshal(toSerialize)
}

type NullableCustomPayloadConfiguration struct {
	value *CustomPayloadConfiguration
	isSet bool
}

func (v NullableCustomPayloadConfiguration) Get() *CustomPayloadConfiguration {
	return v.value
}

func (v *NullableCustomPayloadConfiguration) Set(val *CustomPayloadConfiguration) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomPayloadConfiguration) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomPayloadConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomPayloadConfiguration(val *CustomPayloadConfiguration) *NullableCustomPayloadConfiguration {
	return &NullableCustomPayloadConfiguration{value: val, isSet: true}
}

func (v NullableCustomPayloadConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomPayloadConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
