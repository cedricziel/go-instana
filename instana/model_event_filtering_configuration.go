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

// EventFilteringConfiguration struct for EventFilteringConfiguration
type EventFilteringConfiguration struct {
	Query      *string   `json:"query,omitempty"`
	RuleIds    *[]string `json:"ruleIds,omitempty"`
	EventTypes *[]string `json:"eventTypes,omitempty"`
}

// NewEventFilteringConfiguration instantiates a new EventFilteringConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEventFilteringConfiguration() *EventFilteringConfiguration {
	this := EventFilteringConfiguration{}
	return &this
}

// NewEventFilteringConfigurationWithDefaults instantiates a new EventFilteringConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEventFilteringConfigurationWithDefaults() *EventFilteringConfiguration {
	this := EventFilteringConfiguration{}
	return &this
}

// GetQuery returns the Query field value if set, zero value otherwise.
func (o *EventFilteringConfiguration) GetQuery() string {
	if o == nil || o.Query == nil {
		var ret string
		return ret
	}
	return *o.Query
}

// GetQueryOk returns a tuple with the Query field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventFilteringConfiguration) GetQueryOk() (*string, bool) {
	if o == nil || o.Query == nil {
		return nil, false
	}
	return o.Query, true
}

// HasQuery returns a boolean if a field has been set.
func (o *EventFilteringConfiguration) HasQuery() bool {
	if o != nil && o.Query != nil {
		return true
	}

	return false
}

// SetQuery gets a reference to the given string and assigns it to the Query field.
func (o *EventFilteringConfiguration) SetQuery(v string) {
	o.Query = &v
}

// GetRuleIds returns the RuleIds field value if set, zero value otherwise.
func (o *EventFilteringConfiguration) GetRuleIds() []string {
	if o == nil || o.RuleIds == nil {
		var ret []string
		return ret
	}
	return *o.RuleIds
}

// GetRuleIdsOk returns a tuple with the RuleIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventFilteringConfiguration) GetRuleIdsOk() (*[]string, bool) {
	if o == nil || o.RuleIds == nil {
		return nil, false
	}
	return o.RuleIds, true
}

// HasRuleIds returns a boolean if a field has been set.
func (o *EventFilteringConfiguration) HasRuleIds() bool {
	if o != nil && o.RuleIds != nil {
		return true
	}

	return false
}

// SetRuleIds gets a reference to the given []string and assigns it to the RuleIds field.
func (o *EventFilteringConfiguration) SetRuleIds(v []string) {
	o.RuleIds = &v
}

// GetEventTypes returns the EventTypes field value if set, zero value otherwise.
func (o *EventFilteringConfiguration) GetEventTypes() []string {
	if o == nil || o.EventTypes == nil {
		var ret []string
		return ret
	}
	return *o.EventTypes
}

// GetEventTypesOk returns a tuple with the EventTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventFilteringConfiguration) GetEventTypesOk() (*[]string, bool) {
	if o == nil || o.EventTypes == nil {
		return nil, false
	}
	return o.EventTypes, true
}

// HasEventTypes returns a boolean if a field has been set.
func (o *EventFilteringConfiguration) HasEventTypes() bool {
	if o != nil && o.EventTypes != nil {
		return true
	}

	return false
}

// SetEventTypes gets a reference to the given []string and assigns it to the EventTypes field.
func (o *EventFilteringConfiguration) SetEventTypes(v []string) {
	o.EventTypes = &v
}

func (o EventFilteringConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Query != nil {
		toSerialize["query"] = o.Query
	}
	if o.RuleIds != nil {
		toSerialize["ruleIds"] = o.RuleIds
	}
	if o.EventTypes != nil {
		toSerialize["eventTypes"] = o.EventTypes
	}
	return json.Marshal(toSerialize)
}

type NullableEventFilteringConfiguration struct {
	value *EventFilteringConfiguration
	isSet bool
}

func (v NullableEventFilteringConfiguration) Get() *EventFilteringConfiguration {
	return v.value
}

func (v *NullableEventFilteringConfiguration) Set(val *EventFilteringConfiguration) {
	v.value = val
	v.isSet = true
}

func (v NullableEventFilteringConfiguration) IsSet() bool {
	return v.isSet
}

func (v *NullableEventFilteringConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventFilteringConfiguration(val *EventFilteringConfiguration) *NullableEventFilteringConfiguration {
	return &NullableEventFilteringConfiguration{value: val, isSet: true}
}

func (v NullableEventFilteringConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventFilteringConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
