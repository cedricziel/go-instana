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

// SyntheticCallWithDefaultsConfig struct for SyntheticCallWithDefaultsConfig
type SyntheticCallWithDefaultsConfig struct {
	DefaultRulesEnabled *bool               `json:"defaultRulesEnabled,omitempty"`
	CustomRules         []SyntheticCallRule `json:"customRules"`
	DefaultRules        []SyntheticCallRule `json:"defaultRules"`
}

// NewSyntheticCallWithDefaultsConfig instantiates a new SyntheticCallWithDefaultsConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSyntheticCallWithDefaultsConfig(customRules []SyntheticCallRule, defaultRules []SyntheticCallRule) *SyntheticCallWithDefaultsConfig {
	this := SyntheticCallWithDefaultsConfig{}
	this.CustomRules = customRules
	this.DefaultRules = defaultRules
	return &this
}

// NewSyntheticCallWithDefaultsConfigWithDefaults instantiates a new SyntheticCallWithDefaultsConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSyntheticCallWithDefaultsConfigWithDefaults() *SyntheticCallWithDefaultsConfig {
	this := SyntheticCallWithDefaultsConfig{}
	return &this
}

// GetDefaultRulesEnabled returns the DefaultRulesEnabled field value if set, zero value otherwise.
func (o *SyntheticCallWithDefaultsConfig) GetDefaultRulesEnabled() bool {
	if o == nil || o.DefaultRulesEnabled == nil {
		var ret bool
		return ret
	}
	return *o.DefaultRulesEnabled
}

// GetDefaultRulesEnabledOk returns a tuple with the DefaultRulesEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticCallWithDefaultsConfig) GetDefaultRulesEnabledOk() (*bool, bool) {
	if o == nil || o.DefaultRulesEnabled == nil {
		return nil, false
	}
	return o.DefaultRulesEnabled, true
}

// HasDefaultRulesEnabled returns a boolean if a field has been set.
func (o *SyntheticCallWithDefaultsConfig) HasDefaultRulesEnabled() bool {
	if o != nil && o.DefaultRulesEnabled != nil {
		return true
	}

	return false
}

// SetDefaultRulesEnabled gets a reference to the given bool and assigns it to the DefaultRulesEnabled field.
func (o *SyntheticCallWithDefaultsConfig) SetDefaultRulesEnabled(v bool) {
	o.DefaultRulesEnabled = &v
}

// GetCustomRules returns the CustomRules field value
func (o *SyntheticCallWithDefaultsConfig) GetCustomRules() []SyntheticCallRule {
	if o == nil {
		var ret []SyntheticCallRule
		return ret
	}

	return o.CustomRules
}

// GetCustomRulesOk returns a tuple with the CustomRules field value
// and a boolean to check if the value has been set.
func (o *SyntheticCallWithDefaultsConfig) GetCustomRulesOk() (*[]SyntheticCallRule, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CustomRules, true
}

// SetCustomRules sets field value
func (o *SyntheticCallWithDefaultsConfig) SetCustomRules(v []SyntheticCallRule) {
	o.CustomRules = v
}

// GetDefaultRules returns the DefaultRules field value
func (o *SyntheticCallWithDefaultsConfig) GetDefaultRules() []SyntheticCallRule {
	if o == nil {
		var ret []SyntheticCallRule
		return ret
	}

	return o.DefaultRules
}

// GetDefaultRulesOk returns a tuple with the DefaultRules field value
// and a boolean to check if the value has been set.
func (o *SyntheticCallWithDefaultsConfig) GetDefaultRulesOk() (*[]SyntheticCallRule, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DefaultRules, true
}

// SetDefaultRules sets field value
func (o *SyntheticCallWithDefaultsConfig) SetDefaultRules(v []SyntheticCallRule) {
	o.DefaultRules = v
}

func (o SyntheticCallWithDefaultsConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DefaultRulesEnabled != nil {
		toSerialize["defaultRulesEnabled"] = o.DefaultRulesEnabled
	}
	if true {
		toSerialize["customRules"] = o.CustomRules
	}
	if true {
		toSerialize["defaultRules"] = o.DefaultRules
	}
	return json.Marshal(toSerialize)
}

type NullableSyntheticCallWithDefaultsConfig struct {
	value *SyntheticCallWithDefaultsConfig
	isSet bool
}

func (v NullableSyntheticCallWithDefaultsConfig) Get() *SyntheticCallWithDefaultsConfig {
	return v.value
}

func (v *NullableSyntheticCallWithDefaultsConfig) Set(val *SyntheticCallWithDefaultsConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableSyntheticCallWithDefaultsConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableSyntheticCallWithDefaultsConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSyntheticCallWithDefaultsConfig(val *SyntheticCallWithDefaultsConfig) *NullableSyntheticCallWithDefaultsConfig {
	return &NullableSyntheticCallWithDefaultsConfig{value: val, isSet: true}
}

func (v NullableSyntheticCallWithDefaultsConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSyntheticCallWithDefaultsConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
