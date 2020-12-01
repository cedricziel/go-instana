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

// SlownessWebsiteAlertRuleAllOf struct for SlownessWebsiteAlertRuleAllOf
type SlownessWebsiteAlertRuleAllOf struct {
	Aggregation *string `json:"aggregation,omitempty"`
}

// NewSlownessWebsiteAlertRuleAllOf instantiates a new SlownessWebsiteAlertRuleAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSlownessWebsiteAlertRuleAllOf() *SlownessWebsiteAlertRuleAllOf {
	this := SlownessWebsiteAlertRuleAllOf{}
	return &this
}

// NewSlownessWebsiteAlertRuleAllOfWithDefaults instantiates a new SlownessWebsiteAlertRuleAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSlownessWebsiteAlertRuleAllOfWithDefaults() *SlownessWebsiteAlertRuleAllOf {
	this := SlownessWebsiteAlertRuleAllOf{}
	return &this
}

// GetAggregation returns the Aggregation field value if set, zero value otherwise.
func (o *SlownessWebsiteAlertRuleAllOf) GetAggregation() string {
	if o == nil || o.Aggregation == nil {
		var ret string
		return ret
	}
	return *o.Aggregation
}

// GetAggregationOk returns a tuple with the Aggregation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SlownessWebsiteAlertRuleAllOf) GetAggregationOk() (*string, bool) {
	if o == nil || o.Aggregation == nil {
		return nil, false
	}
	return o.Aggregation, true
}

// HasAggregation returns a boolean if a field has been set.
func (o *SlownessWebsiteAlertRuleAllOf) HasAggregation() bool {
	if o != nil && o.Aggregation != nil {
		return true
	}

	return false
}

// SetAggregation gets a reference to the given string and assigns it to the Aggregation field.
func (o *SlownessWebsiteAlertRuleAllOf) SetAggregation(v string) {
	o.Aggregation = &v
}

func (o SlownessWebsiteAlertRuleAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Aggregation != nil {
		toSerialize["aggregation"] = o.Aggregation
	}
	return json.Marshal(toSerialize)
}

type NullableSlownessWebsiteAlertRuleAllOf struct {
	value *SlownessWebsiteAlertRuleAllOf
	isSet bool
}

func (v NullableSlownessWebsiteAlertRuleAllOf) Get() *SlownessWebsiteAlertRuleAllOf {
	return v.value
}

func (v *NullableSlownessWebsiteAlertRuleAllOf) Set(val *SlownessWebsiteAlertRuleAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableSlownessWebsiteAlertRuleAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableSlownessWebsiteAlertRuleAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSlownessWebsiteAlertRuleAllOf(val *SlownessWebsiteAlertRuleAllOf) *NullableSlownessWebsiteAlertRuleAllOf {
	return &NullableSlownessWebsiteAlertRuleAllOf{value: val, isSet: true}
}

func (v NullableSlownessWebsiteAlertRuleAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSlownessWebsiteAlertRuleAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
