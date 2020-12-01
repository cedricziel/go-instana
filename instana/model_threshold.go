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

// Threshold struct for Threshold
type Threshold struct {
	Type        string `json:"type"`
	Operator    string `json:"operator"`
	LastUpdated *int64 `json:"lastUpdated,omitempty"`
}

// NewThreshold instantiates a new Threshold object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewThreshold(type_ string, operator string) *Threshold {
	this := Threshold{}
	this.Type = type_
	this.Operator = operator
	return &this
}

// NewThresholdWithDefaults instantiates a new Threshold object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewThresholdWithDefaults() *Threshold {
	this := Threshold{}
	return &this
}

// GetType returns the Type field value
func (o *Threshold) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *Threshold) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *Threshold) SetType(v string) {
	o.Type = v
}

// GetOperator returns the Operator field value
func (o *Threshold) GetOperator() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Operator
}

// GetOperatorOk returns a tuple with the Operator field value
// and a boolean to check if the value has been set.
func (o *Threshold) GetOperatorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Operator, true
}

// SetOperator sets field value
func (o *Threshold) SetOperator(v string) {
	o.Operator = v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *Threshold) GetLastUpdated() int64 {
	if o == nil || o.LastUpdated == nil {
		var ret int64
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Threshold) GetLastUpdatedOk() (*int64, bool) {
	if o == nil || o.LastUpdated == nil {
		return nil, false
	}
	return o.LastUpdated, true
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *Threshold) HasLastUpdated() bool {
	if o != nil && o.LastUpdated != nil {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given int64 and assigns it to the LastUpdated field.
func (o *Threshold) SetLastUpdated(v int64) {
	o.LastUpdated = &v
}

func (o Threshold) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["operator"] = o.Operator
	}
	if o.LastUpdated != nil {
		toSerialize["lastUpdated"] = o.LastUpdated
	}
	return json.Marshal(toSerialize)
}

type NullableThreshold struct {
	value *Threshold
	isSet bool
}

func (v NullableThreshold) Get() *Threshold {
	return v.value
}

func (v *NullableThreshold) Set(val *Threshold) {
	v.value = val
	v.isSet = true
}

func (v NullableThreshold) IsSet() bool {
	return v.isSet
}

func (v *NullableThreshold) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableThreshold(val *Threshold) *NullableThreshold {
	return &NullableThreshold{value: val, isSet: true}
}

func (v NullableThreshold) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableThreshold) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
