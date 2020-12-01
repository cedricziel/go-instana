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

// MetricItem struct for MetricItem
type MetricItem struct {
	SnapshotId *string                 `json:"snapshotId,omitempty"`
	Plugin     *string                 `json:"plugin,omitempty"`
	From       *int64                  `json:"from,omitempty"`
	To         *int64                  `json:"to,omitempty"`
	Tags       *[]string               `json:"tags,omitempty"`
	Label      *string                 `json:"label,omitempty"`
	Host       *string                 `json:"host,omitempty"`
	Metrics    *map[string][][]float32 `json:"metrics,omitempty"`
}

// NewMetricItem instantiates a new MetricItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetricItem() *MetricItem {
	this := MetricItem{}
	return &this
}

// NewMetricItemWithDefaults instantiates a new MetricItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetricItemWithDefaults() *MetricItem {
	this := MetricItem{}
	return &this
}

// GetSnapshotId returns the SnapshotId field value if set, zero value otherwise.
func (o *MetricItem) GetSnapshotId() string {
	if o == nil || o.SnapshotId == nil {
		var ret string
		return ret
	}
	return *o.SnapshotId
}

// GetSnapshotIdOk returns a tuple with the SnapshotId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricItem) GetSnapshotIdOk() (*string, bool) {
	if o == nil || o.SnapshotId == nil {
		return nil, false
	}
	return o.SnapshotId, true
}

// HasSnapshotId returns a boolean if a field has been set.
func (o *MetricItem) HasSnapshotId() bool {
	if o != nil && o.SnapshotId != nil {
		return true
	}

	return false
}

// SetSnapshotId gets a reference to the given string and assigns it to the SnapshotId field.
func (o *MetricItem) SetSnapshotId(v string) {
	o.SnapshotId = &v
}

// GetPlugin returns the Plugin field value if set, zero value otherwise.
func (o *MetricItem) GetPlugin() string {
	if o == nil || o.Plugin == nil {
		var ret string
		return ret
	}
	return *o.Plugin
}

// GetPluginOk returns a tuple with the Plugin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricItem) GetPluginOk() (*string, bool) {
	if o == nil || o.Plugin == nil {
		return nil, false
	}
	return o.Plugin, true
}

// HasPlugin returns a boolean if a field has been set.
func (o *MetricItem) HasPlugin() bool {
	if o != nil && o.Plugin != nil {
		return true
	}

	return false
}

// SetPlugin gets a reference to the given string and assigns it to the Plugin field.
func (o *MetricItem) SetPlugin(v string) {
	o.Plugin = &v
}

// GetFrom returns the From field value if set, zero value otherwise.
func (o *MetricItem) GetFrom() int64 {
	if o == nil || o.From == nil {
		var ret int64
		return ret
	}
	return *o.From
}

// GetFromOk returns a tuple with the From field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricItem) GetFromOk() (*int64, bool) {
	if o == nil || o.From == nil {
		return nil, false
	}
	return o.From, true
}

// HasFrom returns a boolean if a field has been set.
func (o *MetricItem) HasFrom() bool {
	if o != nil && o.From != nil {
		return true
	}

	return false
}

// SetFrom gets a reference to the given int64 and assigns it to the From field.
func (o *MetricItem) SetFrom(v int64) {
	o.From = &v
}

// GetTo returns the To field value if set, zero value otherwise.
func (o *MetricItem) GetTo() int64 {
	if o == nil || o.To == nil {
		var ret int64
		return ret
	}
	return *o.To
}

// GetToOk returns a tuple with the To field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricItem) GetToOk() (*int64, bool) {
	if o == nil || o.To == nil {
		return nil, false
	}
	return o.To, true
}

// HasTo returns a boolean if a field has been set.
func (o *MetricItem) HasTo() bool {
	if o != nil && o.To != nil {
		return true
	}

	return false
}

// SetTo gets a reference to the given int64 and assigns it to the To field.
func (o *MetricItem) SetTo(v int64) {
	o.To = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *MetricItem) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricItem) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *MetricItem) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *MetricItem) SetTags(v []string) {
	o.Tags = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *MetricItem) GetLabel() string {
	if o == nil || o.Label == nil {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricItem) GetLabelOk() (*string, bool) {
	if o == nil || o.Label == nil {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *MetricItem) HasLabel() bool {
	if o != nil && o.Label != nil {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *MetricItem) SetLabel(v string) {
	o.Label = &v
}

// GetHost returns the Host field value if set, zero value otherwise.
func (o *MetricItem) GetHost() string {
	if o == nil || o.Host == nil {
		var ret string
		return ret
	}
	return *o.Host
}

// GetHostOk returns a tuple with the Host field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricItem) GetHostOk() (*string, bool) {
	if o == nil || o.Host == nil {
		return nil, false
	}
	return o.Host, true
}

// HasHost returns a boolean if a field has been set.
func (o *MetricItem) HasHost() bool {
	if o != nil && o.Host != nil {
		return true
	}

	return false
}

// SetHost gets a reference to the given string and assigns it to the Host field.
func (o *MetricItem) SetHost(v string) {
	o.Host = &v
}

// GetMetrics returns the Metrics field value if set, zero value otherwise.
func (o *MetricItem) GetMetrics() map[string][][]float32 {
	if o == nil || o.Metrics == nil {
		var ret map[string][][]float32
		return ret
	}
	return *o.Metrics
}

// GetMetricsOk returns a tuple with the Metrics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricItem) GetMetricsOk() (*map[string][][]float32, bool) {
	if o == nil || o.Metrics == nil {
		return nil, false
	}
	return o.Metrics, true
}

// HasMetrics returns a boolean if a field has been set.
func (o *MetricItem) HasMetrics() bool {
	if o != nil && o.Metrics != nil {
		return true
	}

	return false
}

// SetMetrics gets a reference to the given map[string][][]float32 and assigns it to the Metrics field.
func (o *MetricItem) SetMetrics(v map[string][][]float32) {
	o.Metrics = &v
}

func (o MetricItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.SnapshotId != nil {
		toSerialize["snapshotId"] = o.SnapshotId
	}
	if o.Plugin != nil {
		toSerialize["plugin"] = o.Plugin
	}
	if o.From != nil {
		toSerialize["from"] = o.From
	}
	if o.To != nil {
		toSerialize["to"] = o.To
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if o.Label != nil {
		toSerialize["label"] = o.Label
	}
	if o.Host != nil {
		toSerialize["host"] = o.Host
	}
	if o.Metrics != nil {
		toSerialize["metrics"] = o.Metrics
	}
	return json.Marshal(toSerialize)
}

type NullableMetricItem struct {
	value *MetricItem
	isSet bool
}

func (v NullableMetricItem) Get() *MetricItem {
	return v.value
}

func (v *NullableMetricItem) Set(val *MetricItem) {
	v.value = val
	v.isSet = true
}

func (v NullableMetricItem) IsSet() bool {
	return v.isSet
}

func (v *NullableMetricItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetricItem(val *MetricItem) *NullableMetricItem {
	return &NullableMetricItem{value: val, isSet: true}
}

func (v NullableMetricItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetricItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
