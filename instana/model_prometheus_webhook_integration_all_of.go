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

// PrometheusWebhookIntegrationAllOf struct for PrometheusWebhookIntegrationAllOf
type PrometheusWebhookIntegrationAllOf struct {
	WebhookUrl *string `json:"webhookUrl,omitempty"`
	Receiver   *string `json:"receiver,omitempty"`
}

// NewPrometheusWebhookIntegrationAllOf instantiates a new PrometheusWebhookIntegrationAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPrometheusWebhookIntegrationAllOf() *PrometheusWebhookIntegrationAllOf {
	this := PrometheusWebhookIntegrationAllOf{}
	return &this
}

// NewPrometheusWebhookIntegrationAllOfWithDefaults instantiates a new PrometheusWebhookIntegrationAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPrometheusWebhookIntegrationAllOfWithDefaults() *PrometheusWebhookIntegrationAllOf {
	this := PrometheusWebhookIntegrationAllOf{}
	return &this
}

// GetWebhookUrl returns the WebhookUrl field value if set, zero value otherwise.
func (o *PrometheusWebhookIntegrationAllOf) GetWebhookUrl() string {
	if o == nil || o.WebhookUrl == nil {
		var ret string
		return ret
	}
	return *o.WebhookUrl
}

// GetWebhookUrlOk returns a tuple with the WebhookUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrometheusWebhookIntegrationAllOf) GetWebhookUrlOk() (*string, bool) {
	if o == nil || o.WebhookUrl == nil {
		return nil, false
	}
	return o.WebhookUrl, true
}

// HasWebhookUrl returns a boolean if a field has been set.
func (o *PrometheusWebhookIntegrationAllOf) HasWebhookUrl() bool {
	if o != nil && o.WebhookUrl != nil {
		return true
	}

	return false
}

// SetWebhookUrl gets a reference to the given string and assigns it to the WebhookUrl field.
func (o *PrometheusWebhookIntegrationAllOf) SetWebhookUrl(v string) {
	o.WebhookUrl = &v
}

// GetReceiver returns the Receiver field value if set, zero value otherwise.
func (o *PrometheusWebhookIntegrationAllOf) GetReceiver() string {
	if o == nil || o.Receiver == nil {
		var ret string
		return ret
	}
	return *o.Receiver
}

// GetReceiverOk returns a tuple with the Receiver field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrometheusWebhookIntegrationAllOf) GetReceiverOk() (*string, bool) {
	if o == nil || o.Receiver == nil {
		return nil, false
	}
	return o.Receiver, true
}

// HasReceiver returns a boolean if a field has been set.
func (o *PrometheusWebhookIntegrationAllOf) HasReceiver() bool {
	if o != nil && o.Receiver != nil {
		return true
	}

	return false
}

// SetReceiver gets a reference to the given string and assigns it to the Receiver field.
func (o *PrometheusWebhookIntegrationAllOf) SetReceiver(v string) {
	o.Receiver = &v
}

func (o PrometheusWebhookIntegrationAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.WebhookUrl != nil {
		toSerialize["webhookUrl"] = o.WebhookUrl
	}
	if o.Receiver != nil {
		toSerialize["receiver"] = o.Receiver
	}
	return json.Marshal(toSerialize)
}

type NullablePrometheusWebhookIntegrationAllOf struct {
	value *PrometheusWebhookIntegrationAllOf
	isSet bool
}

func (v NullablePrometheusWebhookIntegrationAllOf) Get() *PrometheusWebhookIntegrationAllOf {
	return v.value
}

func (v *NullablePrometheusWebhookIntegrationAllOf) Set(val *PrometheusWebhookIntegrationAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullablePrometheusWebhookIntegrationAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullablePrometheusWebhookIntegrationAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePrometheusWebhookIntegrationAllOf(val *PrometheusWebhookIntegrationAllOf) *NullablePrometheusWebhookIntegrationAllOf {
	return &NullablePrometheusWebhookIntegrationAllOf{value: val, isSet: true}
}

func (v NullablePrometheusWebhookIntegrationAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePrometheusWebhookIntegrationAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
