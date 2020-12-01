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

// ValidatedMaintenanceConfigWithStatus struct for ValidatedMaintenanceConfigWithStatus
type ValidatedMaintenanceConfigWithStatus struct {
	Id          string               `json:"id"`
	Name        string               `json:"name"`
	Query       string               `json:"query"`
	Windows     *[]MaintenanceWindow `json:"windows,omitempty"`
	LastUpdated *int64               `json:"lastUpdated,omitempty"`
	Status      string               `json:"status"`
	Invalid     *bool                `json:"invalid,omitempty"`
}

// NewValidatedMaintenanceConfigWithStatus instantiates a new ValidatedMaintenanceConfigWithStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewValidatedMaintenanceConfigWithStatus(id string, name string, query string, status string) *ValidatedMaintenanceConfigWithStatus {
	this := ValidatedMaintenanceConfigWithStatus{}
	this.Id = id
	this.Name = name
	this.Query = query
	this.Status = status
	return &this
}

// NewValidatedMaintenanceConfigWithStatusWithDefaults instantiates a new ValidatedMaintenanceConfigWithStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewValidatedMaintenanceConfigWithStatusWithDefaults() *ValidatedMaintenanceConfigWithStatus {
	this := ValidatedMaintenanceConfigWithStatus{}
	return &this
}

// GetId returns the Id field value
func (o *ValidatedMaintenanceConfigWithStatus) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ValidatedMaintenanceConfigWithStatus) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ValidatedMaintenanceConfigWithStatus) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *ValidatedMaintenanceConfigWithStatus) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ValidatedMaintenanceConfigWithStatus) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ValidatedMaintenanceConfigWithStatus) SetName(v string) {
	o.Name = v
}

// GetQuery returns the Query field value
func (o *ValidatedMaintenanceConfigWithStatus) GetQuery() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Query
}

// GetQueryOk returns a tuple with the Query field value
// and a boolean to check if the value has been set.
func (o *ValidatedMaintenanceConfigWithStatus) GetQueryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Query, true
}

// SetQuery sets field value
func (o *ValidatedMaintenanceConfigWithStatus) SetQuery(v string) {
	o.Query = v
}

// GetWindows returns the Windows field value if set, zero value otherwise.
func (o *ValidatedMaintenanceConfigWithStatus) GetWindows() []MaintenanceWindow {
	if o == nil || o.Windows == nil {
		var ret []MaintenanceWindow
		return ret
	}
	return *o.Windows
}

// GetWindowsOk returns a tuple with the Windows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidatedMaintenanceConfigWithStatus) GetWindowsOk() (*[]MaintenanceWindow, bool) {
	if o == nil || o.Windows == nil {
		return nil, false
	}
	return o.Windows, true
}

// HasWindows returns a boolean if a field has been set.
func (o *ValidatedMaintenanceConfigWithStatus) HasWindows() bool {
	if o != nil && o.Windows != nil {
		return true
	}

	return false
}

// SetWindows gets a reference to the given []MaintenanceWindow and assigns it to the Windows field.
func (o *ValidatedMaintenanceConfigWithStatus) SetWindows(v []MaintenanceWindow) {
	o.Windows = &v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *ValidatedMaintenanceConfigWithStatus) GetLastUpdated() int64 {
	if o == nil || o.LastUpdated == nil {
		var ret int64
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidatedMaintenanceConfigWithStatus) GetLastUpdatedOk() (*int64, bool) {
	if o == nil || o.LastUpdated == nil {
		return nil, false
	}
	return o.LastUpdated, true
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *ValidatedMaintenanceConfigWithStatus) HasLastUpdated() bool {
	if o != nil && o.LastUpdated != nil {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given int64 and assigns it to the LastUpdated field.
func (o *ValidatedMaintenanceConfigWithStatus) SetLastUpdated(v int64) {
	o.LastUpdated = &v
}

// GetStatus returns the Status field value
func (o *ValidatedMaintenanceConfigWithStatus) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *ValidatedMaintenanceConfigWithStatus) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *ValidatedMaintenanceConfigWithStatus) SetStatus(v string) {
	o.Status = v
}

// GetInvalid returns the Invalid field value if set, zero value otherwise.
func (o *ValidatedMaintenanceConfigWithStatus) GetInvalid() bool {
	if o == nil || o.Invalid == nil {
		var ret bool
		return ret
	}
	return *o.Invalid
}

// GetInvalidOk returns a tuple with the Invalid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidatedMaintenanceConfigWithStatus) GetInvalidOk() (*bool, bool) {
	if o == nil || o.Invalid == nil {
		return nil, false
	}
	return o.Invalid, true
}

// HasInvalid returns a boolean if a field has been set.
func (o *ValidatedMaintenanceConfigWithStatus) HasInvalid() bool {
	if o != nil && o.Invalid != nil {
		return true
	}

	return false
}

// SetInvalid gets a reference to the given bool and assigns it to the Invalid field.
func (o *ValidatedMaintenanceConfigWithStatus) SetInvalid(v bool) {
	o.Invalid = &v
}

func (o ValidatedMaintenanceConfigWithStatus) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["query"] = o.Query
	}
	if o.Windows != nil {
		toSerialize["windows"] = o.Windows
	}
	if o.LastUpdated != nil {
		toSerialize["lastUpdated"] = o.LastUpdated
	}
	if true {
		toSerialize["status"] = o.Status
	}
	if o.Invalid != nil {
		toSerialize["invalid"] = o.Invalid
	}
	return json.Marshal(toSerialize)
}

type NullableValidatedMaintenanceConfigWithStatus struct {
	value *ValidatedMaintenanceConfigWithStatus
	isSet bool
}

func (v NullableValidatedMaintenanceConfigWithStatus) Get() *ValidatedMaintenanceConfigWithStatus {
	return v.value
}

func (v *NullableValidatedMaintenanceConfigWithStatus) Set(val *ValidatedMaintenanceConfigWithStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableValidatedMaintenanceConfigWithStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableValidatedMaintenanceConfigWithStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableValidatedMaintenanceConfigWithStatus(val *ValidatedMaintenanceConfigWithStatus) *NullableValidatedMaintenanceConfigWithStatus {
	return &NullableValidatedMaintenanceConfigWithStatus{value: val, isSet: true}
}

func (v NullableValidatedMaintenanceConfigWithStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableValidatedMaintenanceConfigWithStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
