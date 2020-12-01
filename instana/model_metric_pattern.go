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

// MetricPattern struct for MetricPattern
type MetricPattern struct {
	Prefix      string  `json:"prefix"`
	Postfix     *string `json:"postfix,omitempty"`
	Placeholder *string `json:"placeholder,omitempty"`
	Operator    string  `json:"operator"`
}

// NewMetricPattern instantiates a new MetricPattern object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetricPattern(prefix string, operator string) *MetricPattern {
	this := MetricPattern{}
	this.Prefix = prefix
	this.Operator = operator
	return &this
}

// NewMetricPatternWithDefaults instantiates a new MetricPattern object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetricPatternWithDefaults() *MetricPattern {
	this := MetricPattern{}
	return &this
}

// GetPrefix returns the Prefix field value
func (o *MetricPattern) GetPrefix() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Prefix
}

// GetPrefixOk returns a tuple with the Prefix field value
// and a boolean to check if the value has been set.
func (o *MetricPattern) GetPrefixOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Prefix, true
}

// SetPrefix sets field value
func (o *MetricPattern) SetPrefix(v string) {
	o.Prefix = v
}

// GetPostfix returns the Postfix field value if set, zero value otherwise.
func (o *MetricPattern) GetPostfix() string {
	if o == nil || o.Postfix == nil {
		var ret string
		return ret
	}
	return *o.Postfix
}

// GetPostfixOk returns a tuple with the Postfix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricPattern) GetPostfixOk() (*string, bool) {
	if o == nil || o.Postfix == nil {
		return nil, false
	}
	return o.Postfix, true
}

// HasPostfix returns a boolean if a field has been set.
func (o *MetricPattern) HasPostfix() bool {
	if o != nil && o.Postfix != nil {
		return true
	}

	return false
}

// SetPostfix gets a reference to the given string and assigns it to the Postfix field.
func (o *MetricPattern) SetPostfix(v string) {
	o.Postfix = &v
}

// GetPlaceholder returns the Placeholder field value if set, zero value otherwise.
func (o *MetricPattern) GetPlaceholder() string {
	if o == nil || o.Placeholder == nil {
		var ret string
		return ret
	}
	return *o.Placeholder
}

// GetPlaceholderOk returns a tuple with the Placeholder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricPattern) GetPlaceholderOk() (*string, bool) {
	if o == nil || o.Placeholder == nil {
		return nil, false
	}
	return o.Placeholder, true
}

// HasPlaceholder returns a boolean if a field has been set.
func (o *MetricPattern) HasPlaceholder() bool {
	if o != nil && o.Placeholder != nil {
		return true
	}

	return false
}

// SetPlaceholder gets a reference to the given string and assigns it to the Placeholder field.
func (o *MetricPattern) SetPlaceholder(v string) {
	o.Placeholder = &v
}

// GetOperator returns the Operator field value
func (o *MetricPattern) GetOperator() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Operator
}

// GetOperatorOk returns a tuple with the Operator field value
// and a boolean to check if the value has been set.
func (o *MetricPattern) GetOperatorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Operator, true
}

// SetOperator sets field value
func (o *MetricPattern) SetOperator(v string) {
	o.Operator = v
}

func (o MetricPattern) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["prefix"] = o.Prefix
	}
	if o.Postfix != nil {
		toSerialize["postfix"] = o.Postfix
	}
	if o.Placeholder != nil {
		toSerialize["placeholder"] = o.Placeholder
	}
	if true {
		toSerialize["operator"] = o.Operator
	}
	return json.Marshal(toSerialize)
}

type NullableMetricPattern struct {
	value *MetricPattern
	isSet bool
}

func (v NullableMetricPattern) Get() *MetricPattern {
	return v.value
}

func (v *NullableMetricPattern) Set(val *MetricPattern) {
	v.value = val
	v.isSet = true
}

func (v NullableMetricPattern) IsSet() bool {
	return v.isSet
}

func (v *NullableMetricPattern) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetricPattern(val *MetricPattern) *NullableMetricPattern {
	return &NullableMetricPattern{value: val, isSet: true}
}

func (v NullableMetricPattern) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetricPattern) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
