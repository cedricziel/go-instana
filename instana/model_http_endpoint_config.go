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

// HttpEndpointConfig struct for HttpEndpointConfig
type HttpEndpointConfig struct {
	ServiceId                                      string             `json:"serviceId"`
	EndpointNameByFirstPathSegmentRuleEnabled      *bool              `json:"endpointNameByFirstPathSegmentRuleEnabled,omitempty"`
	EndpointNameByCollectedPathTemplateRuleEnabled *bool              `json:"endpointNameByCollectedPathTemplateRuleEnabled,omitempty"`
	Rules                                          []HttpEndpointRule `json:"rules"`
}

// NewHttpEndpointConfig instantiates a new HttpEndpointConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHttpEndpointConfig(serviceId string, rules []HttpEndpointRule) *HttpEndpointConfig {
	this := HttpEndpointConfig{}
	this.ServiceId = serviceId
	this.Rules = rules
	return &this
}

// NewHttpEndpointConfigWithDefaults instantiates a new HttpEndpointConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHttpEndpointConfigWithDefaults() *HttpEndpointConfig {
	this := HttpEndpointConfig{}
	return &this
}

// GetServiceId returns the ServiceId field value
func (o *HttpEndpointConfig) GetServiceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceId
}

// GetServiceIdOk returns a tuple with the ServiceId field value
// and a boolean to check if the value has been set.
func (o *HttpEndpointConfig) GetServiceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceId, true
}

// SetServiceId sets field value
func (o *HttpEndpointConfig) SetServiceId(v string) {
	o.ServiceId = v
}

// GetEndpointNameByFirstPathSegmentRuleEnabled returns the EndpointNameByFirstPathSegmentRuleEnabled field value if set, zero value otherwise.
func (o *HttpEndpointConfig) GetEndpointNameByFirstPathSegmentRuleEnabled() bool {
	if o == nil || o.EndpointNameByFirstPathSegmentRuleEnabled == nil {
		var ret bool
		return ret
	}
	return *o.EndpointNameByFirstPathSegmentRuleEnabled
}

// GetEndpointNameByFirstPathSegmentRuleEnabledOk returns a tuple with the EndpointNameByFirstPathSegmentRuleEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HttpEndpointConfig) GetEndpointNameByFirstPathSegmentRuleEnabledOk() (*bool, bool) {
	if o == nil || o.EndpointNameByFirstPathSegmentRuleEnabled == nil {
		return nil, false
	}
	return o.EndpointNameByFirstPathSegmentRuleEnabled, true
}

// HasEndpointNameByFirstPathSegmentRuleEnabled returns a boolean if a field has been set.
func (o *HttpEndpointConfig) HasEndpointNameByFirstPathSegmentRuleEnabled() bool {
	if o != nil && o.EndpointNameByFirstPathSegmentRuleEnabled != nil {
		return true
	}

	return false
}

// SetEndpointNameByFirstPathSegmentRuleEnabled gets a reference to the given bool and assigns it to the EndpointNameByFirstPathSegmentRuleEnabled field.
func (o *HttpEndpointConfig) SetEndpointNameByFirstPathSegmentRuleEnabled(v bool) {
	o.EndpointNameByFirstPathSegmentRuleEnabled = &v
}

// GetEndpointNameByCollectedPathTemplateRuleEnabled returns the EndpointNameByCollectedPathTemplateRuleEnabled field value if set, zero value otherwise.
func (o *HttpEndpointConfig) GetEndpointNameByCollectedPathTemplateRuleEnabled() bool {
	if o == nil || o.EndpointNameByCollectedPathTemplateRuleEnabled == nil {
		var ret bool
		return ret
	}
	return *o.EndpointNameByCollectedPathTemplateRuleEnabled
}

// GetEndpointNameByCollectedPathTemplateRuleEnabledOk returns a tuple with the EndpointNameByCollectedPathTemplateRuleEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HttpEndpointConfig) GetEndpointNameByCollectedPathTemplateRuleEnabledOk() (*bool, bool) {
	if o == nil || o.EndpointNameByCollectedPathTemplateRuleEnabled == nil {
		return nil, false
	}
	return o.EndpointNameByCollectedPathTemplateRuleEnabled, true
}

// HasEndpointNameByCollectedPathTemplateRuleEnabled returns a boolean if a field has been set.
func (o *HttpEndpointConfig) HasEndpointNameByCollectedPathTemplateRuleEnabled() bool {
	if o != nil && o.EndpointNameByCollectedPathTemplateRuleEnabled != nil {
		return true
	}

	return false
}

// SetEndpointNameByCollectedPathTemplateRuleEnabled gets a reference to the given bool and assigns it to the EndpointNameByCollectedPathTemplateRuleEnabled field.
func (o *HttpEndpointConfig) SetEndpointNameByCollectedPathTemplateRuleEnabled(v bool) {
	o.EndpointNameByCollectedPathTemplateRuleEnabled = &v
}

// GetRules returns the Rules field value
func (o *HttpEndpointConfig) GetRules() []HttpEndpointRule {
	if o == nil {
		var ret []HttpEndpointRule
		return ret
	}

	return o.Rules
}

// GetRulesOk returns a tuple with the Rules field value
// and a boolean to check if the value has been set.
func (o *HttpEndpointConfig) GetRulesOk() (*[]HttpEndpointRule, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Rules, true
}

// SetRules sets field value
func (o *HttpEndpointConfig) SetRules(v []HttpEndpointRule) {
	o.Rules = v
}

func (o HttpEndpointConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["serviceId"] = o.ServiceId
	}
	if o.EndpointNameByFirstPathSegmentRuleEnabled != nil {
		toSerialize["endpointNameByFirstPathSegmentRuleEnabled"] = o.EndpointNameByFirstPathSegmentRuleEnabled
	}
	if o.EndpointNameByCollectedPathTemplateRuleEnabled != nil {
		toSerialize["endpointNameByCollectedPathTemplateRuleEnabled"] = o.EndpointNameByCollectedPathTemplateRuleEnabled
	}
	if true {
		toSerialize["rules"] = o.Rules
	}
	return json.Marshal(toSerialize)
}

type NullableHttpEndpointConfig struct {
	value *HttpEndpointConfig
	isSet bool
}

func (v NullableHttpEndpointConfig) Get() *HttpEndpointConfig {
	return v.value
}

func (v *NullableHttpEndpointConfig) Set(val *HttpEndpointConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableHttpEndpointConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableHttpEndpointConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHttpEndpointConfig(val *HttpEndpointConfig) *NullableHttpEndpointConfig {
	return &NullableHttpEndpointConfig{value: val, isSet: true}
}

func (v NullableHttpEndpointConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHttpEndpointConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
