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

// SliConfiguration struct for SliConfiguration
type SliConfiguration struct {
	Id                         string               `json:"id"`
	SliName                    string               `json:"sliName"`
	InitialEvaluationTimestamp *int64               `json:"initialEvaluationTimestamp,omitempty"`
	MetricConfiguration        *MetricConfiguration `json:"metricConfiguration,omitempty"`
	SliEntity                  SliEntity            `json:"sliEntity"`
}

// NewSliConfiguration instantiates a new SliConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSliConfiguration(id string, sliName string, sliEntity SliEntity) *SliConfiguration {
	this := SliConfiguration{}
	this.Id = id
	this.SliName = sliName
	this.SliEntity = sliEntity
	return &this
}

// NewSliConfigurationWithDefaults instantiates a new SliConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSliConfigurationWithDefaults() *SliConfiguration {
	this := SliConfiguration{}
	return &this
}

// GetId returns the Id field value
func (o *SliConfiguration) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *SliConfiguration) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *SliConfiguration) SetId(v string) {
	o.Id = v
}

// GetSliName returns the SliName field value
func (o *SliConfiguration) GetSliName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SliName
}

// GetSliNameOk returns a tuple with the SliName field value
// and a boolean to check if the value has been set.
func (o *SliConfiguration) GetSliNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SliName, true
}

// SetSliName sets field value
func (o *SliConfiguration) SetSliName(v string) {
	o.SliName = v
}

// GetInitialEvaluationTimestamp returns the InitialEvaluationTimestamp field value if set, zero value otherwise.
func (o *SliConfiguration) GetInitialEvaluationTimestamp() int64 {
	if o == nil || o.InitialEvaluationTimestamp == nil {
		var ret int64
		return ret
	}
	return *o.InitialEvaluationTimestamp
}

// GetInitialEvaluationTimestampOk returns a tuple with the InitialEvaluationTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SliConfiguration) GetInitialEvaluationTimestampOk() (*int64, bool) {
	if o == nil || o.InitialEvaluationTimestamp == nil {
		return nil, false
	}
	return o.InitialEvaluationTimestamp, true
}

// HasInitialEvaluationTimestamp returns a boolean if a field has been set.
func (o *SliConfiguration) HasInitialEvaluationTimestamp() bool {
	if o != nil && o.InitialEvaluationTimestamp != nil {
		return true
	}

	return false
}

// SetInitialEvaluationTimestamp gets a reference to the given int64 and assigns it to the InitialEvaluationTimestamp field.
func (o *SliConfiguration) SetInitialEvaluationTimestamp(v int64) {
	o.InitialEvaluationTimestamp = &v
}

// GetMetricConfiguration returns the MetricConfiguration field value if set, zero value otherwise.
func (o *SliConfiguration) GetMetricConfiguration() MetricConfiguration {
	if o == nil || o.MetricConfiguration == nil {
		var ret MetricConfiguration
		return ret
	}
	return *o.MetricConfiguration
}

// GetMetricConfigurationOk returns a tuple with the MetricConfiguration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SliConfiguration) GetMetricConfigurationOk() (*MetricConfiguration, bool) {
	if o == nil || o.MetricConfiguration == nil {
		return nil, false
	}
	return o.MetricConfiguration, true
}

// HasMetricConfiguration returns a boolean if a field has been set.
func (o *SliConfiguration) HasMetricConfiguration() bool {
	if o != nil && o.MetricConfiguration != nil {
		return true
	}

	return false
}

// SetMetricConfiguration gets a reference to the given MetricConfiguration and assigns it to the MetricConfiguration field.
func (o *SliConfiguration) SetMetricConfiguration(v MetricConfiguration) {
	o.MetricConfiguration = &v
}

// GetSliEntity returns the SliEntity field value
func (o *SliConfiguration) GetSliEntity() SliEntity {
	if o == nil {
		var ret SliEntity
		return ret
	}

	return o.SliEntity
}

// GetSliEntityOk returns a tuple with the SliEntity field value
// and a boolean to check if the value has been set.
func (o *SliConfiguration) GetSliEntityOk() (*SliEntity, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SliEntity, true
}

// SetSliEntity sets field value
func (o *SliConfiguration) SetSliEntity(v SliEntity) {
	o.SliEntity = v
}

func (o SliConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["sliName"] = o.SliName
	}
	if o.InitialEvaluationTimestamp != nil {
		toSerialize["initialEvaluationTimestamp"] = o.InitialEvaluationTimestamp
	}
	if o.MetricConfiguration != nil {
		toSerialize["metricConfiguration"] = o.MetricConfiguration
	}
	if true {
		toSerialize["sliEntity"] = o.SliEntity
	}
	return json.Marshal(toSerialize)
}

type NullableSliConfiguration struct {
	value *SliConfiguration
	isSet bool
}

func (v NullableSliConfiguration) Get() *SliConfiguration {
	return v.value
}

func (v *NullableSliConfiguration) Set(val *SliConfiguration) {
	v.value = val
	v.isSet = true
}

func (v NullableSliConfiguration) IsSet() bool {
	return v.isSet
}

func (v *NullableSliConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSliConfiguration(val *SliConfiguration) *NullableSliConfiguration {
	return &NullableSliConfiguration{value: val, isSet: true}
}

func (v NullableSliConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSliConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
