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

// ServiceScopeWithMetadata struct for ServiceScopeWithMetadata
type ServiceScopeWithMetadata struct {
	Id       string                       `json:"id"`
	Name     *string                      `json:"name,omitempty"`
	ScopedTo *ServiceScopedToWithMetadata `json:"scopedTo,omitempty"`
}

// NewServiceScopeWithMetadata instantiates a new ServiceScopeWithMetadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceScopeWithMetadata(id string) *ServiceScopeWithMetadata {
	this := ServiceScopeWithMetadata{}
	this.Id = id
	return &this
}

// NewServiceScopeWithMetadataWithDefaults instantiates a new ServiceScopeWithMetadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceScopeWithMetadataWithDefaults() *ServiceScopeWithMetadata {
	this := ServiceScopeWithMetadata{}
	return &this
}

// GetId returns the Id field value
func (o *ServiceScopeWithMetadata) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ServiceScopeWithMetadata) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ServiceScopeWithMetadata) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ServiceScopeWithMetadata) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceScopeWithMetadata) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ServiceScopeWithMetadata) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ServiceScopeWithMetadata) SetName(v string) {
	o.Name = &v
}

// GetScopedTo returns the ScopedTo field value if set, zero value otherwise.
func (o *ServiceScopeWithMetadata) GetScopedTo() ServiceScopedToWithMetadata {
	if o == nil || o.ScopedTo == nil {
		var ret ServiceScopedToWithMetadata
		return ret
	}
	return *o.ScopedTo
}

// GetScopedToOk returns a tuple with the ScopedTo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceScopeWithMetadata) GetScopedToOk() (*ServiceScopedToWithMetadata, bool) {
	if o == nil || o.ScopedTo == nil {
		return nil, false
	}
	return o.ScopedTo, true
}

// HasScopedTo returns a boolean if a field has been set.
func (o *ServiceScopeWithMetadata) HasScopedTo() bool {
	if o != nil && o.ScopedTo != nil {
		return true
	}

	return false
}

// SetScopedTo gets a reference to the given ServiceScopedToWithMetadata and assigns it to the ScopedTo field.
func (o *ServiceScopeWithMetadata) SetScopedTo(v ServiceScopedToWithMetadata) {
	o.ScopedTo = &v
}

func (o ServiceScopeWithMetadata) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.ScopedTo != nil {
		toSerialize["scopedTo"] = o.ScopedTo
	}
	return json.Marshal(toSerialize)
}

type NullableServiceScopeWithMetadata struct {
	value *ServiceScopeWithMetadata
	isSet bool
}

func (v NullableServiceScopeWithMetadata) Get() *ServiceScopeWithMetadata {
	return v.value
}

func (v *NullableServiceScopeWithMetadata) Set(val *ServiceScopeWithMetadata) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceScopeWithMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceScopeWithMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceScopeWithMetadata(val *ServiceScopeWithMetadata) *NullableServiceScopeWithMetadata {
	return &NullableServiceScopeWithMetadata{value: val, isSet: true}
}

func (v NullableServiceScopeWithMetadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceScopeWithMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
