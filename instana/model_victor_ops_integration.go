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

// VictorOpsIntegration struct for VictorOpsIntegration
type VictorOpsIntegration struct {
	AbstractIntegration
	ApiKey     string `json:"apiKey"`
	RoutingKey string `json:"routingKey"`
}

// NewVictorOpsIntegration instantiates a new VictorOpsIntegration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVictorOpsIntegration(apiKey string, routingKey string) *VictorOpsIntegration {
	this := VictorOpsIntegration{}
	this.ApiKey = apiKey
	this.RoutingKey = routingKey
	return &this
}

// NewVictorOpsIntegrationWithDefaults instantiates a new VictorOpsIntegration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVictorOpsIntegrationWithDefaults() *VictorOpsIntegration {
	this := VictorOpsIntegration{}
	return &this
}

// GetApiKey returns the ApiKey field value
func (o *VictorOpsIntegration) GetApiKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ApiKey
}

// GetApiKeyOk returns a tuple with the ApiKey field value
// and a boolean to check if the value has been set.
func (o *VictorOpsIntegration) GetApiKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApiKey, true
}

// SetApiKey sets field value
func (o *VictorOpsIntegration) SetApiKey(v string) {
	o.ApiKey = v
}

// GetRoutingKey returns the RoutingKey field value
func (o *VictorOpsIntegration) GetRoutingKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RoutingKey
}

// GetRoutingKeyOk returns a tuple with the RoutingKey field value
// and a boolean to check if the value has been set.
func (o *VictorOpsIntegration) GetRoutingKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RoutingKey, true
}

// SetRoutingKey sets field value
func (o *VictorOpsIntegration) SetRoutingKey(v string) {
	o.RoutingKey = v
}

func (o VictorOpsIntegration) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedAbstractIntegration, errAbstractIntegration := json.Marshal(o.AbstractIntegration)
	if errAbstractIntegration != nil {
		return []byte{}, errAbstractIntegration
	}
	errAbstractIntegration = json.Unmarshal([]byte(serializedAbstractIntegration), &toSerialize)
	if errAbstractIntegration != nil {
		return []byte{}, errAbstractIntegration
	}
	if true {
		toSerialize["apiKey"] = o.ApiKey
	}
	if true {
		toSerialize["routingKey"] = o.RoutingKey
	}
	return json.Marshal(toSerialize)
}

type NullableVictorOpsIntegration struct {
	value *VictorOpsIntegration
	isSet bool
}

func (v NullableVictorOpsIntegration) Get() *VictorOpsIntegration {
	return v.value
}

func (v *NullableVictorOpsIntegration) Set(val *VictorOpsIntegration) {
	v.value = val
	v.isSet = true
}

func (v NullableVictorOpsIntegration) IsSet() bool {
	return v.isSet
}

func (v *NullableVictorOpsIntegration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVictorOpsIntegration(val *VictorOpsIntegration) *NullableVictorOpsIntegration {
	return &NullableVictorOpsIntegration{value: val, isSet: true}
}

func (v NullableVictorOpsIntegration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVictorOpsIntegration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
