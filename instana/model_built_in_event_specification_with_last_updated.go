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

// BuiltInEventSpecificationWithLastUpdated struct for BuiltInEventSpecificationWithLastUpdated
type BuiltInEventSpecificationWithLastUpdated struct {
	Id            string       `json:"id"`
	ShortPluginId string       `json:"shortPluginId"`
	Name          string       `json:"name"`
	Description   *string      `json:"description,omitempty"`
	HyperParams   []HyperParam `json:"hyperParams"`
	RuleInputs    []RuleInput  `json:"ruleInputs"`
	Severity      *int32       `json:"severity,omitempty"`
	Triggering    *bool        `json:"triggering,omitempty"`
	Enabled       *bool        `json:"enabled,omitempty"`
	LastUpdated   *int64       `json:"lastUpdated,omitempty"`
}

// NewBuiltInEventSpecificationWithLastUpdated instantiates a new BuiltInEventSpecificationWithLastUpdated object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBuiltInEventSpecificationWithLastUpdated(id string, shortPluginId string, name string, hyperParams []HyperParam, ruleInputs []RuleInput) *BuiltInEventSpecificationWithLastUpdated {
	this := BuiltInEventSpecificationWithLastUpdated{}
	this.Id = id
	this.ShortPluginId = shortPluginId
	this.Name = name
	this.HyperParams = hyperParams
	this.RuleInputs = ruleInputs
	return &this
}

// NewBuiltInEventSpecificationWithLastUpdatedWithDefaults instantiates a new BuiltInEventSpecificationWithLastUpdated object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBuiltInEventSpecificationWithLastUpdatedWithDefaults() *BuiltInEventSpecificationWithLastUpdated {
	this := BuiltInEventSpecificationWithLastUpdated{}
	return &this
}

// GetId returns the Id field value
func (o *BuiltInEventSpecificationWithLastUpdated) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *BuiltInEventSpecificationWithLastUpdated) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *BuiltInEventSpecificationWithLastUpdated) SetId(v string) {
	o.Id = v
}

// GetShortPluginId returns the ShortPluginId field value
func (o *BuiltInEventSpecificationWithLastUpdated) GetShortPluginId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ShortPluginId
}

// GetShortPluginIdOk returns a tuple with the ShortPluginId field value
// and a boolean to check if the value has been set.
func (o *BuiltInEventSpecificationWithLastUpdated) GetShortPluginIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ShortPluginId, true
}

// SetShortPluginId sets field value
func (o *BuiltInEventSpecificationWithLastUpdated) SetShortPluginId(v string) {
	o.ShortPluginId = v
}

// GetName returns the Name field value
func (o *BuiltInEventSpecificationWithLastUpdated) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *BuiltInEventSpecificationWithLastUpdated) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *BuiltInEventSpecificationWithLastUpdated) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *BuiltInEventSpecificationWithLastUpdated) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BuiltInEventSpecificationWithLastUpdated) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *BuiltInEventSpecificationWithLastUpdated) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *BuiltInEventSpecificationWithLastUpdated) SetDescription(v string) {
	o.Description = &v
}

// GetHyperParams returns the HyperParams field value
func (o *BuiltInEventSpecificationWithLastUpdated) GetHyperParams() []HyperParam {
	if o == nil {
		var ret []HyperParam
		return ret
	}

	return o.HyperParams
}

// GetHyperParamsOk returns a tuple with the HyperParams field value
// and a boolean to check if the value has been set.
func (o *BuiltInEventSpecificationWithLastUpdated) GetHyperParamsOk() (*[]HyperParam, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HyperParams, true
}

// SetHyperParams sets field value
func (o *BuiltInEventSpecificationWithLastUpdated) SetHyperParams(v []HyperParam) {
	o.HyperParams = v
}

// GetRuleInputs returns the RuleInputs field value
func (o *BuiltInEventSpecificationWithLastUpdated) GetRuleInputs() []RuleInput {
	if o == nil {
		var ret []RuleInput
		return ret
	}

	return o.RuleInputs
}

// GetRuleInputsOk returns a tuple with the RuleInputs field value
// and a boolean to check if the value has been set.
func (o *BuiltInEventSpecificationWithLastUpdated) GetRuleInputsOk() (*[]RuleInput, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RuleInputs, true
}

// SetRuleInputs sets field value
func (o *BuiltInEventSpecificationWithLastUpdated) SetRuleInputs(v []RuleInput) {
	o.RuleInputs = v
}

// GetSeverity returns the Severity field value if set, zero value otherwise.
func (o *BuiltInEventSpecificationWithLastUpdated) GetSeverity() int32 {
	if o == nil || o.Severity == nil {
		var ret int32
		return ret
	}
	return *o.Severity
}

// GetSeverityOk returns a tuple with the Severity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BuiltInEventSpecificationWithLastUpdated) GetSeverityOk() (*int32, bool) {
	if o == nil || o.Severity == nil {
		return nil, false
	}
	return o.Severity, true
}

// HasSeverity returns a boolean if a field has been set.
func (o *BuiltInEventSpecificationWithLastUpdated) HasSeverity() bool {
	if o != nil && o.Severity != nil {
		return true
	}

	return false
}

// SetSeverity gets a reference to the given int32 and assigns it to the Severity field.
func (o *BuiltInEventSpecificationWithLastUpdated) SetSeverity(v int32) {
	o.Severity = &v
}

// GetTriggering returns the Triggering field value if set, zero value otherwise.
func (o *BuiltInEventSpecificationWithLastUpdated) GetTriggering() bool {
	if o == nil || o.Triggering == nil {
		var ret bool
		return ret
	}
	return *o.Triggering
}

// GetTriggeringOk returns a tuple with the Triggering field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BuiltInEventSpecificationWithLastUpdated) GetTriggeringOk() (*bool, bool) {
	if o == nil || o.Triggering == nil {
		return nil, false
	}
	return o.Triggering, true
}

// HasTriggering returns a boolean if a field has been set.
func (o *BuiltInEventSpecificationWithLastUpdated) HasTriggering() bool {
	if o != nil && o.Triggering != nil {
		return true
	}

	return false
}

// SetTriggering gets a reference to the given bool and assigns it to the Triggering field.
func (o *BuiltInEventSpecificationWithLastUpdated) SetTriggering(v bool) {
	o.Triggering = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *BuiltInEventSpecificationWithLastUpdated) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BuiltInEventSpecificationWithLastUpdated) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *BuiltInEventSpecificationWithLastUpdated) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *BuiltInEventSpecificationWithLastUpdated) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *BuiltInEventSpecificationWithLastUpdated) GetLastUpdated() int64 {
	if o == nil || o.LastUpdated == nil {
		var ret int64
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BuiltInEventSpecificationWithLastUpdated) GetLastUpdatedOk() (*int64, bool) {
	if o == nil || o.LastUpdated == nil {
		return nil, false
	}
	return o.LastUpdated, true
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *BuiltInEventSpecificationWithLastUpdated) HasLastUpdated() bool {
	if o != nil && o.LastUpdated != nil {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given int64 and assigns it to the LastUpdated field.
func (o *BuiltInEventSpecificationWithLastUpdated) SetLastUpdated(v int64) {
	o.LastUpdated = &v
}

func (o BuiltInEventSpecificationWithLastUpdated) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["shortPluginId"] = o.ShortPluginId
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["hyperParams"] = o.HyperParams
	}
	if true {
		toSerialize["ruleInputs"] = o.RuleInputs
	}
	if o.Severity != nil {
		toSerialize["severity"] = o.Severity
	}
	if o.Triggering != nil {
		toSerialize["triggering"] = o.Triggering
	}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if o.LastUpdated != nil {
		toSerialize["lastUpdated"] = o.LastUpdated
	}
	return json.Marshal(toSerialize)
}

type NullableBuiltInEventSpecificationWithLastUpdated struct {
	value *BuiltInEventSpecificationWithLastUpdated
	isSet bool
}

func (v NullableBuiltInEventSpecificationWithLastUpdated) Get() *BuiltInEventSpecificationWithLastUpdated {
	return v.value
}

func (v *NullableBuiltInEventSpecificationWithLastUpdated) Set(val *BuiltInEventSpecificationWithLastUpdated) {
	v.value = val
	v.isSet = true
}

func (v NullableBuiltInEventSpecificationWithLastUpdated) IsSet() bool {
	return v.isSet
}

func (v *NullableBuiltInEventSpecificationWithLastUpdated) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBuiltInEventSpecificationWithLastUpdated(val *BuiltInEventSpecificationWithLastUpdated) *NullableBuiltInEventSpecificationWithLastUpdated {
	return &NullableBuiltInEventSpecificationWithLastUpdated{value: val, isSet: true}
}

func (v NullableBuiltInEventSpecificationWithLastUpdated) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBuiltInEventSpecificationWithLastUpdated) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
