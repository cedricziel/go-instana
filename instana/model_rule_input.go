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

// RuleInput struct for RuleInput
type RuleInput struct {
	InputKind string `json:"inputKind"`
	InputName string `json:"inputName"`
}

// NewRuleInput instantiates a new RuleInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRuleInput(inputKind string, inputName string) *RuleInput {
	this := RuleInput{}
	this.InputKind = inputKind
	this.InputName = inputName
	return &this
}

// NewRuleInputWithDefaults instantiates a new RuleInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRuleInputWithDefaults() *RuleInput {
	this := RuleInput{}
	return &this
}

// GetInputKind returns the InputKind field value
func (o *RuleInput) GetInputKind() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.InputKind
}

// GetInputKindOk returns a tuple with the InputKind field value
// and a boolean to check if the value has been set.
func (o *RuleInput) GetInputKindOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InputKind, true
}

// SetInputKind sets field value
func (o *RuleInput) SetInputKind(v string) {
	o.InputKind = v
}

// GetInputName returns the InputName field value
func (o *RuleInput) GetInputName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.InputName
}

// GetInputNameOk returns a tuple with the InputName field value
// and a boolean to check if the value has been set.
func (o *RuleInput) GetInputNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InputName, true
}

// SetInputName sets field value
func (o *RuleInput) SetInputName(v string) {
	o.InputName = v
}

func (o RuleInput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["inputKind"] = o.InputKind
	}
	if true {
		toSerialize["inputName"] = o.InputName
	}
	return json.Marshal(toSerialize)
}

type NullableRuleInput struct {
	value *RuleInput
	isSet bool
}

func (v NullableRuleInput) Get() *RuleInput {
	return v.value
}

func (v *NullableRuleInput) Set(val *RuleInput) {
	v.value = val
	v.isSet = true
}

func (v NullableRuleInput) IsSet() bool {
	return v.isSet
}

func (v *NullableRuleInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRuleInput(val *RuleInput) *NullableRuleInput {
	return &NullableRuleInput{value: val, isSet: true}
}

func (v NullableRuleInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRuleInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
