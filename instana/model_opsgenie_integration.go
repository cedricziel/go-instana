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

// OpsgenieIntegration struct for OpsgenieIntegration
type OpsgenieIntegration struct {
	AbstractIntegration
	ApiKey string  `json:"apiKey"`
	Tags   *string `json:"tags,omitempty"`
	Region string  `json:"region"`
	Alias  *string `json:"alias,omitempty"`
}

// NewOpsgenieIntegration instantiates a new OpsgenieIntegration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOpsgenieIntegration(apiKey string, region string) *OpsgenieIntegration {
	this := OpsgenieIntegration{}
	this.ApiKey = apiKey
	this.Region = region
	return &this
}

// NewOpsgenieIntegrationWithDefaults instantiates a new OpsgenieIntegration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOpsgenieIntegrationWithDefaults() *OpsgenieIntegration {
	this := OpsgenieIntegration{}
	return &this
}

// GetApiKey returns the ApiKey field value
func (o *OpsgenieIntegration) GetApiKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ApiKey
}

// GetApiKeyOk returns a tuple with the ApiKey field value
// and a boolean to check if the value has been set.
func (o *OpsgenieIntegration) GetApiKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApiKey, true
}

// SetApiKey sets field value
func (o *OpsgenieIntegration) SetApiKey(v string) {
	o.ApiKey = v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *OpsgenieIntegration) GetTags() string {
	if o == nil || o.Tags == nil {
		var ret string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpsgenieIntegration) GetTagsOk() (*string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *OpsgenieIntegration) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given string and assigns it to the Tags field.
func (o *OpsgenieIntegration) SetTags(v string) {
	o.Tags = &v
}

// GetRegion returns the Region field value
func (o *OpsgenieIntegration) GetRegion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Region
}

// GetRegionOk returns a tuple with the Region field value
// and a boolean to check if the value has been set.
func (o *OpsgenieIntegration) GetRegionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Region, true
}

// SetRegion sets field value
func (o *OpsgenieIntegration) SetRegion(v string) {
	o.Region = v
}

// GetAlias returns the Alias field value if set, zero value otherwise.
func (o *OpsgenieIntegration) GetAlias() string {
	if o == nil || o.Alias == nil {
		var ret string
		return ret
	}
	return *o.Alias
}

// GetAliasOk returns a tuple with the Alias field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpsgenieIntegration) GetAliasOk() (*string, bool) {
	if o == nil || o.Alias == nil {
		return nil, false
	}
	return o.Alias, true
}

// HasAlias returns a boolean if a field has been set.
func (o *OpsgenieIntegration) HasAlias() bool {
	if o != nil && o.Alias != nil {
		return true
	}

	return false
}

// SetAlias gets a reference to the given string and assigns it to the Alias field.
func (o *OpsgenieIntegration) SetAlias(v string) {
	o.Alias = &v
}

func (o OpsgenieIntegration) MarshalJSON() ([]byte, error) {
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
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if true {
		toSerialize["region"] = o.Region
	}
	if o.Alias != nil {
		toSerialize["alias"] = o.Alias
	}
	return json.Marshal(toSerialize)
}

type NullableOpsgenieIntegration struct {
	value *OpsgenieIntegration
	isSet bool
}

func (v NullableOpsgenieIntegration) Get() *OpsgenieIntegration {
	return v.value
}

func (v *NullableOpsgenieIntegration) Set(val *OpsgenieIntegration) {
	v.value = val
	v.isSet = true
}

func (v NullableOpsgenieIntegration) IsSet() bool {
	return v.isSet
}

func (v *NullableOpsgenieIntegration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOpsgenieIntegration(val *OpsgenieIntegration) *NullableOpsgenieIntegration {
	return &NullableOpsgenieIntegration{value: val, isSet: true}
}

func (v NullableOpsgenieIntegration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOpsgenieIntegration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
