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

// SystemRuleAllOf struct for SystemRuleAllOf
type SystemRuleAllOf struct {
	SystemRuleId *string `json:"systemRuleId,omitempty"`
}

// NewSystemRuleAllOf instantiates a new SystemRuleAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSystemRuleAllOf() *SystemRuleAllOf {
	this := SystemRuleAllOf{}
	return &this
}

// NewSystemRuleAllOfWithDefaults instantiates a new SystemRuleAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSystemRuleAllOfWithDefaults() *SystemRuleAllOf {
	this := SystemRuleAllOf{}
	return &this
}

// GetSystemRuleId returns the SystemRuleId field value if set, zero value otherwise.
func (o *SystemRuleAllOf) GetSystemRuleId() string {
	if o == nil || o.SystemRuleId == nil {
		var ret string
		return ret
	}
	return *o.SystemRuleId
}

// GetSystemRuleIdOk returns a tuple with the SystemRuleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemRuleAllOf) GetSystemRuleIdOk() (*string, bool) {
	if o == nil || o.SystemRuleId == nil {
		return nil, false
	}
	return o.SystemRuleId, true
}

// HasSystemRuleId returns a boolean if a field has been set.
func (o *SystemRuleAllOf) HasSystemRuleId() bool {
	if o != nil && o.SystemRuleId != nil {
		return true
	}

	return false
}

// SetSystemRuleId gets a reference to the given string and assigns it to the SystemRuleId field.
func (o *SystemRuleAllOf) SetSystemRuleId(v string) {
	o.SystemRuleId = &v
}

func (o SystemRuleAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.SystemRuleId != nil {
		toSerialize["systemRuleId"] = o.SystemRuleId
	}
	return json.Marshal(toSerialize)
}

type NullableSystemRuleAllOf struct {
	value *SystemRuleAllOf
	isSet bool
}

func (v NullableSystemRuleAllOf) Get() *SystemRuleAllOf {
	return v.value
}

func (v *NullableSystemRuleAllOf) Set(val *SystemRuleAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableSystemRuleAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableSystemRuleAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSystemRuleAllOf(val *SystemRuleAllOf) *NullableSystemRuleAllOf {
	return &NullableSystemRuleAllOf{value: val, isSet: true}
}

func (v NullableSystemRuleAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSystemRuleAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
