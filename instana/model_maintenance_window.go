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

// MaintenanceWindow struct for MaintenanceWindow
type MaintenanceWindow struct {
	Id    string `json:"id"`
	Start *int64 `json:"start,omitempty"`
	End   *int64 `json:"end,omitempty"`
}

// NewMaintenanceWindow instantiates a new MaintenanceWindow object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMaintenanceWindow(id string) *MaintenanceWindow {
	this := MaintenanceWindow{}
	this.Id = id
	return &this
}

// NewMaintenanceWindowWithDefaults instantiates a new MaintenanceWindow object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMaintenanceWindowWithDefaults() *MaintenanceWindow {
	this := MaintenanceWindow{}
	return &this
}

// GetId returns the Id field value
func (o *MaintenanceWindow) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *MaintenanceWindow) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *MaintenanceWindow) SetId(v string) {
	o.Id = v
}

// GetStart returns the Start field value if set, zero value otherwise.
func (o *MaintenanceWindow) GetStart() int64 {
	if o == nil || o.Start == nil {
		var ret int64
		return ret
	}
	return *o.Start
}

// GetStartOk returns a tuple with the Start field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MaintenanceWindow) GetStartOk() (*int64, bool) {
	if o == nil || o.Start == nil {
		return nil, false
	}
	return o.Start, true
}

// HasStart returns a boolean if a field has been set.
func (o *MaintenanceWindow) HasStart() bool {
	if o != nil && o.Start != nil {
		return true
	}

	return false
}

// SetStart gets a reference to the given int64 and assigns it to the Start field.
func (o *MaintenanceWindow) SetStart(v int64) {
	o.Start = &v
}

// GetEnd returns the End field value if set, zero value otherwise.
func (o *MaintenanceWindow) GetEnd() int64 {
	if o == nil || o.End == nil {
		var ret int64
		return ret
	}
	return *o.End
}

// GetEndOk returns a tuple with the End field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MaintenanceWindow) GetEndOk() (*int64, bool) {
	if o == nil || o.End == nil {
		return nil, false
	}
	return o.End, true
}

// HasEnd returns a boolean if a field has been set.
func (o *MaintenanceWindow) HasEnd() bool {
	if o != nil && o.End != nil {
		return true
	}

	return false
}

// SetEnd gets a reference to the given int64 and assigns it to the End field.
func (o *MaintenanceWindow) SetEnd(v int64) {
	o.End = &v
}

func (o MaintenanceWindow) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if o.Start != nil {
		toSerialize["start"] = o.Start
	}
	if o.End != nil {
		toSerialize["end"] = o.End
	}
	return json.Marshal(toSerialize)
}

type NullableMaintenanceWindow struct {
	value *MaintenanceWindow
	isSet bool
}

func (v NullableMaintenanceWindow) Get() *MaintenanceWindow {
	return v.value
}

func (v *NullableMaintenanceWindow) Set(val *MaintenanceWindow) {
	v.value = val
	v.isSet = true
}

func (v NullableMaintenanceWindow) IsSet() bool {
	return v.isSet
}

func (v *NullableMaintenanceWindow) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMaintenanceWindow(val *MaintenanceWindow) *NullableMaintenanceWindow {
	return &NullableMaintenanceWindow{value: val, isSet: true}
}

func (v NullableMaintenanceWindow) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMaintenanceWindow) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
