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

// TimeFrame struct for TimeFrame
type TimeFrame struct {
	WindowSize *int64 `json:"windowSize,omitempty"`
	To         *int64 `json:"to,omitempty"`
}

// NewTimeFrame instantiates a new TimeFrame object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTimeFrame() *TimeFrame {
	this := TimeFrame{}
	return &this
}

// NewTimeFrameWithDefaults instantiates a new TimeFrame object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTimeFrameWithDefaults() *TimeFrame {
	this := TimeFrame{}
	return &this
}

// GetWindowSize returns the WindowSize field value if set, zero value otherwise.
func (o *TimeFrame) GetWindowSize() int64 {
	if o == nil || o.WindowSize == nil {
		var ret int64
		return ret
	}
	return *o.WindowSize
}

// GetWindowSizeOk returns a tuple with the WindowSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeFrame) GetWindowSizeOk() (*int64, bool) {
	if o == nil || o.WindowSize == nil {
		return nil, false
	}
	return o.WindowSize, true
}

// HasWindowSize returns a boolean if a field has been set.
func (o *TimeFrame) HasWindowSize() bool {
	if o != nil && o.WindowSize != nil {
		return true
	}

	return false
}

// SetWindowSize gets a reference to the given int64 and assigns it to the WindowSize field.
func (o *TimeFrame) SetWindowSize(v int64) {
	o.WindowSize = &v
}

// GetTo returns the To field value if set, zero value otherwise.
func (o *TimeFrame) GetTo() int64 {
	if o == nil || o.To == nil {
		var ret int64
		return ret
	}
	return *o.To
}

// GetToOk returns a tuple with the To field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeFrame) GetToOk() (*int64, bool) {
	if o == nil || o.To == nil {
		return nil, false
	}
	return o.To, true
}

// HasTo returns a boolean if a field has been set.
func (o *TimeFrame) HasTo() bool {
	if o != nil && o.To != nil {
		return true
	}

	return false
}

// SetTo gets a reference to the given int64 and assigns it to the To field.
func (o *TimeFrame) SetTo(v int64) {
	o.To = &v
}

func (o TimeFrame) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.WindowSize != nil {
		toSerialize["windowSize"] = o.WindowSize
	}
	if o.To != nil {
		toSerialize["to"] = o.To
	}
	return json.Marshal(toSerialize)
}

type NullableTimeFrame struct {
	value *TimeFrame
	isSet bool
}

func (v NullableTimeFrame) Get() *TimeFrame {
	return v.value
}

func (v *NullableTimeFrame) Set(val *TimeFrame) {
	v.value = val
	v.isSet = true
}

func (v NullableTimeFrame) IsSet() bool {
	return v.isSet
}

func (v *NullableTimeFrame) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTimeFrame(val *TimeFrame) *NullableTimeFrame {
	return &NullableTimeFrame{value: val, isSet: true}
}

func (v NullableTimeFrame) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTimeFrame) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
