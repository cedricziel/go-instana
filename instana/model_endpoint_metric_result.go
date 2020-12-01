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

// EndpointMetricResult struct for EndpointMetricResult
type EndpointMetricResult struct {
	Items     []EndpointItem `json:"items"`
	Page      *int32         `json:"page,omitempty"`
	PageSize  *int32         `json:"pageSize,omitempty"`
	TotalHits *int64         `json:"totalHits,omitempty"`
}

// NewEndpointMetricResult instantiates a new EndpointMetricResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEndpointMetricResult(items []EndpointItem) *EndpointMetricResult {
	this := EndpointMetricResult{}
	this.Items = items
	return &this
}

// NewEndpointMetricResultWithDefaults instantiates a new EndpointMetricResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEndpointMetricResultWithDefaults() *EndpointMetricResult {
	this := EndpointMetricResult{}
	return &this
}

// GetItems returns the Items field value
func (o *EndpointMetricResult) GetItems() []EndpointItem {
	if o == nil {
		var ret []EndpointItem
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *EndpointMetricResult) GetItemsOk() (*[]EndpointItem, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Items, true
}

// SetItems sets field value
func (o *EndpointMetricResult) SetItems(v []EndpointItem) {
	o.Items = v
}

// GetPage returns the Page field value if set, zero value otherwise.
func (o *EndpointMetricResult) GetPage() int32 {
	if o == nil || o.Page == nil {
		var ret int32
		return ret
	}
	return *o.Page
}

// GetPageOk returns a tuple with the Page field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndpointMetricResult) GetPageOk() (*int32, bool) {
	if o == nil || o.Page == nil {
		return nil, false
	}
	return o.Page, true
}

// HasPage returns a boolean if a field has been set.
func (o *EndpointMetricResult) HasPage() bool {
	if o != nil && o.Page != nil {
		return true
	}

	return false
}

// SetPage gets a reference to the given int32 and assigns it to the Page field.
func (o *EndpointMetricResult) SetPage(v int32) {
	o.Page = &v
}

// GetPageSize returns the PageSize field value if set, zero value otherwise.
func (o *EndpointMetricResult) GetPageSize() int32 {
	if o == nil || o.PageSize == nil {
		var ret int32
		return ret
	}
	return *o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndpointMetricResult) GetPageSizeOk() (*int32, bool) {
	if o == nil || o.PageSize == nil {
		return nil, false
	}
	return o.PageSize, true
}

// HasPageSize returns a boolean if a field has been set.
func (o *EndpointMetricResult) HasPageSize() bool {
	if o != nil && o.PageSize != nil {
		return true
	}

	return false
}

// SetPageSize gets a reference to the given int32 and assigns it to the PageSize field.
func (o *EndpointMetricResult) SetPageSize(v int32) {
	o.PageSize = &v
}

// GetTotalHits returns the TotalHits field value if set, zero value otherwise.
func (o *EndpointMetricResult) GetTotalHits() int64 {
	if o == nil || o.TotalHits == nil {
		var ret int64
		return ret
	}
	return *o.TotalHits
}

// GetTotalHitsOk returns a tuple with the TotalHits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndpointMetricResult) GetTotalHitsOk() (*int64, bool) {
	if o == nil || o.TotalHits == nil {
		return nil, false
	}
	return o.TotalHits, true
}

// HasTotalHits returns a boolean if a field has been set.
func (o *EndpointMetricResult) HasTotalHits() bool {
	if o != nil && o.TotalHits != nil {
		return true
	}

	return false
}

// SetTotalHits gets a reference to the given int64 and assigns it to the TotalHits field.
func (o *EndpointMetricResult) SetTotalHits(v int64) {
	o.TotalHits = &v
}

func (o EndpointMetricResult) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["items"] = o.Items
	}
	if o.Page != nil {
		toSerialize["page"] = o.Page
	}
	if o.PageSize != nil {
		toSerialize["pageSize"] = o.PageSize
	}
	if o.TotalHits != nil {
		toSerialize["totalHits"] = o.TotalHits
	}
	return json.Marshal(toSerialize)
}

type NullableEndpointMetricResult struct {
	value *EndpointMetricResult
	isSet bool
}

func (v NullableEndpointMetricResult) Get() *EndpointMetricResult {
	return v.value
}

func (v *NullableEndpointMetricResult) Set(val *EndpointMetricResult) {
	v.value = val
	v.isSet = true
}

func (v NullableEndpointMetricResult) IsSet() bool {
	return v.isSet
}

func (v *NullableEndpointMetricResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEndpointMetricResult(val *EndpointMetricResult) *NullableEndpointMetricResult {
	return &NullableEndpointMetricResult{value: val, isSet: true}
}

func (v NullableEndpointMetricResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEndpointMetricResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
