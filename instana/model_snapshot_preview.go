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

// SnapshotPreview struct for SnapshotPreview
type SnapshotPreview struct {
	Id     string                  `json:"id"`
	Time   *int64                  `json:"time,omitempty"`
	Label  *string                 `json:"label,omitempty"`
	Plugin *string                 `json:"plugin,omitempty"`
	Data   *map[string]interface{} `json:"data,omitempty"`
}

// NewSnapshotPreview instantiates a new SnapshotPreview object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSnapshotPreview(id string) *SnapshotPreview {
	this := SnapshotPreview{}
	this.Id = id
	return &this
}

// NewSnapshotPreviewWithDefaults instantiates a new SnapshotPreview object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSnapshotPreviewWithDefaults() *SnapshotPreview {
	this := SnapshotPreview{}
	return &this
}

// GetId returns the Id field value
func (o *SnapshotPreview) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *SnapshotPreview) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *SnapshotPreview) SetId(v string) {
	o.Id = v
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *SnapshotPreview) GetTime() int64 {
	if o == nil || o.Time == nil {
		var ret int64
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnapshotPreview) GetTimeOk() (*int64, bool) {
	if o == nil || o.Time == nil {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *SnapshotPreview) HasTime() bool {
	if o != nil && o.Time != nil {
		return true
	}

	return false
}

// SetTime gets a reference to the given int64 and assigns it to the Time field.
func (o *SnapshotPreview) SetTime(v int64) {
	o.Time = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *SnapshotPreview) GetLabel() string {
	if o == nil || o.Label == nil {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnapshotPreview) GetLabelOk() (*string, bool) {
	if o == nil || o.Label == nil {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *SnapshotPreview) HasLabel() bool {
	if o != nil && o.Label != nil {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *SnapshotPreview) SetLabel(v string) {
	o.Label = &v
}

// GetPlugin returns the Plugin field value if set, zero value otherwise.
func (o *SnapshotPreview) GetPlugin() string {
	if o == nil || o.Plugin == nil {
		var ret string
		return ret
	}
	return *o.Plugin
}

// GetPluginOk returns a tuple with the Plugin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnapshotPreview) GetPluginOk() (*string, bool) {
	if o == nil || o.Plugin == nil {
		return nil, false
	}
	return o.Plugin, true
}

// HasPlugin returns a boolean if a field has been set.
func (o *SnapshotPreview) HasPlugin() bool {
	if o != nil && o.Plugin != nil {
		return true
	}

	return false
}

// SetPlugin gets a reference to the given string and assigns it to the Plugin field.
func (o *SnapshotPreview) SetPlugin(v string) {
	o.Plugin = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *SnapshotPreview) GetData() map[string]interface{} {
	if o == nil || o.Data == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnapshotPreview) GetDataOk() (*map[string]interface{}, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *SnapshotPreview) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given map[string]interface{} and assigns it to the Data field.
func (o *SnapshotPreview) SetData(v map[string]interface{}) {
	o.Data = &v
}

func (o SnapshotPreview) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if o.Time != nil {
		toSerialize["time"] = o.Time
	}
	if o.Label != nil {
		toSerialize["label"] = o.Label
	}
	if o.Plugin != nil {
		toSerialize["plugin"] = o.Plugin
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableSnapshotPreview struct {
	value *SnapshotPreview
	isSet bool
}

func (v NullableSnapshotPreview) Get() *SnapshotPreview {
	return v.value
}

func (v *NullableSnapshotPreview) Set(val *SnapshotPreview) {
	v.value = val
	v.isSet = true
}

func (v NullableSnapshotPreview) IsSet() bool {
	return v.isSet
}

func (v *NullableSnapshotPreview) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSnapshotPreview(val *SnapshotPreview) *NullableSnapshotPreview {
	return &NullableSnapshotPreview{value: val, isSet: true}
}

func (v NullableSnapshotPreview) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSnapshotPreview) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
