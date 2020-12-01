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

// UnsupportedHttpPathSegmentMatchingRule struct for UnsupportedHttpPathSegmentMatchingRule
type UnsupportedHttpPathSegmentMatchingRule struct {
	HttpPathSegmentMatchingRule
	UnsupportedType *string `json:"unsupportedType,omitempty"`
}

// NewUnsupportedHttpPathSegmentMatchingRule instantiates a new UnsupportedHttpPathSegmentMatchingRule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnsupportedHttpPathSegmentMatchingRule() *UnsupportedHttpPathSegmentMatchingRule {
	this := UnsupportedHttpPathSegmentMatchingRule{}
	return &this
}

// NewUnsupportedHttpPathSegmentMatchingRuleWithDefaults instantiates a new UnsupportedHttpPathSegmentMatchingRule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnsupportedHttpPathSegmentMatchingRuleWithDefaults() *UnsupportedHttpPathSegmentMatchingRule {
	this := UnsupportedHttpPathSegmentMatchingRule{}
	return &this
}

// GetUnsupportedType returns the UnsupportedType field value if set, zero value otherwise.
func (o *UnsupportedHttpPathSegmentMatchingRule) GetUnsupportedType() string {
	if o == nil || o.UnsupportedType == nil {
		var ret string
		return ret
	}
	return *o.UnsupportedType
}

// GetUnsupportedTypeOk returns a tuple with the UnsupportedType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnsupportedHttpPathSegmentMatchingRule) GetUnsupportedTypeOk() (*string, bool) {
	if o == nil || o.UnsupportedType == nil {
		return nil, false
	}
	return o.UnsupportedType, true
}

// HasUnsupportedType returns a boolean if a field has been set.
func (o *UnsupportedHttpPathSegmentMatchingRule) HasUnsupportedType() bool {
	if o != nil && o.UnsupportedType != nil {
		return true
	}

	return false
}

// SetUnsupportedType gets a reference to the given string and assigns it to the UnsupportedType field.
func (o *UnsupportedHttpPathSegmentMatchingRule) SetUnsupportedType(v string) {
	o.UnsupportedType = &v
}

func (o UnsupportedHttpPathSegmentMatchingRule) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedHttpPathSegmentMatchingRule, errHttpPathSegmentMatchingRule := json.Marshal(o.HttpPathSegmentMatchingRule)
	if errHttpPathSegmentMatchingRule != nil {
		return []byte{}, errHttpPathSegmentMatchingRule
	}
	errHttpPathSegmentMatchingRule = json.Unmarshal([]byte(serializedHttpPathSegmentMatchingRule), &toSerialize)
	if errHttpPathSegmentMatchingRule != nil {
		return []byte{}, errHttpPathSegmentMatchingRule
	}
	if o.UnsupportedType != nil {
		toSerialize["unsupportedType"] = o.UnsupportedType
	}
	return json.Marshal(toSerialize)
}

type NullableUnsupportedHttpPathSegmentMatchingRule struct {
	value *UnsupportedHttpPathSegmentMatchingRule
	isSet bool
}

func (v NullableUnsupportedHttpPathSegmentMatchingRule) Get() *UnsupportedHttpPathSegmentMatchingRule {
	return v.value
}

func (v *NullableUnsupportedHttpPathSegmentMatchingRule) Set(val *UnsupportedHttpPathSegmentMatchingRule) {
	v.value = val
	v.isSet = true
}

func (v NullableUnsupportedHttpPathSegmentMatchingRule) IsSet() bool {
	return v.isSet
}

func (v *NullableUnsupportedHttpPathSegmentMatchingRule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnsupportedHttpPathSegmentMatchingRule(val *UnsupportedHttpPathSegmentMatchingRule) *NullableUnsupportedHttpPathSegmentMatchingRule {
	return &NullableUnsupportedHttpPathSegmentMatchingRule{value: val, isSet: true}
}

func (v NullableUnsupportedHttpPathSegmentMatchingRule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnsupportedHttpPathSegmentMatchingRule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
