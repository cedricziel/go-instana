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

// StaticBooleanFieldAllOf struct for StaticBooleanFieldAllOf
type StaticBooleanFieldAllOf struct {
	Value *bool `json:"value,omitempty"`
}

// NewStaticBooleanFieldAllOf instantiates a new StaticBooleanFieldAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStaticBooleanFieldAllOf() *StaticBooleanFieldAllOf {
	this := StaticBooleanFieldAllOf{}
	return &this
}

// NewStaticBooleanFieldAllOfWithDefaults instantiates a new StaticBooleanFieldAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStaticBooleanFieldAllOfWithDefaults() *StaticBooleanFieldAllOf {
	this := StaticBooleanFieldAllOf{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *StaticBooleanFieldAllOf) GetValue() bool {
	if o == nil || o.Value == nil {
		var ret bool
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StaticBooleanFieldAllOf) GetValueOk() (*bool, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *StaticBooleanFieldAllOf) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given bool and assigns it to the Value field.
func (o *StaticBooleanFieldAllOf) SetValue(v bool) {
	o.Value = &v
}

func (o StaticBooleanFieldAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableStaticBooleanFieldAllOf struct {
	value *StaticBooleanFieldAllOf
	isSet bool
}

func (v NullableStaticBooleanFieldAllOf) Get() *StaticBooleanFieldAllOf {
	return v.value
}

func (v *NullableStaticBooleanFieldAllOf) Set(val *StaticBooleanFieldAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableStaticBooleanFieldAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableStaticBooleanFieldAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStaticBooleanFieldAllOf(val *StaticBooleanFieldAllOf) *NullableStaticBooleanFieldAllOf {
	return &NullableStaticBooleanFieldAllOf{value: val, isSet: true}
}

func (v NullableStaticBooleanFieldAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStaticBooleanFieldAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}