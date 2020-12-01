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

// CloudfoundryPhysicalContext struct for CloudfoundryPhysicalContext
type CloudfoundryPhysicalContext struct {
	Application     *SnapshotPreview `json:"application,omitempty"`
	Space           *SnapshotPreview `json:"space,omitempty"`
	Organization    *SnapshotPreview `json:"organization,omitempty"`
	CfInstanceIndex *string          `json:"cfInstanceIndex,omitempty"`
}

// NewCloudfoundryPhysicalContext instantiates a new CloudfoundryPhysicalContext object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloudfoundryPhysicalContext() *CloudfoundryPhysicalContext {
	this := CloudfoundryPhysicalContext{}
	return &this
}

// NewCloudfoundryPhysicalContextWithDefaults instantiates a new CloudfoundryPhysicalContext object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudfoundryPhysicalContextWithDefaults() *CloudfoundryPhysicalContext {
	this := CloudfoundryPhysicalContext{}
	return &this
}

// GetApplication returns the Application field value if set, zero value otherwise.
func (o *CloudfoundryPhysicalContext) GetApplication() SnapshotPreview {
	if o == nil || o.Application == nil {
		var ret SnapshotPreview
		return ret
	}
	return *o.Application
}

// GetApplicationOk returns a tuple with the Application field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudfoundryPhysicalContext) GetApplicationOk() (*SnapshotPreview, bool) {
	if o == nil || o.Application == nil {
		return nil, false
	}
	return o.Application, true
}

// HasApplication returns a boolean if a field has been set.
func (o *CloudfoundryPhysicalContext) HasApplication() bool {
	if o != nil && o.Application != nil {
		return true
	}

	return false
}

// SetApplication gets a reference to the given SnapshotPreview and assigns it to the Application field.
func (o *CloudfoundryPhysicalContext) SetApplication(v SnapshotPreview) {
	o.Application = &v
}

// GetSpace returns the Space field value if set, zero value otherwise.
func (o *CloudfoundryPhysicalContext) GetSpace() SnapshotPreview {
	if o == nil || o.Space == nil {
		var ret SnapshotPreview
		return ret
	}
	return *o.Space
}

// GetSpaceOk returns a tuple with the Space field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudfoundryPhysicalContext) GetSpaceOk() (*SnapshotPreview, bool) {
	if o == nil || o.Space == nil {
		return nil, false
	}
	return o.Space, true
}

// HasSpace returns a boolean if a field has been set.
func (o *CloudfoundryPhysicalContext) HasSpace() bool {
	if o != nil && o.Space != nil {
		return true
	}

	return false
}

// SetSpace gets a reference to the given SnapshotPreview and assigns it to the Space field.
func (o *CloudfoundryPhysicalContext) SetSpace(v SnapshotPreview) {
	o.Space = &v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *CloudfoundryPhysicalContext) GetOrganization() SnapshotPreview {
	if o == nil || o.Organization == nil {
		var ret SnapshotPreview
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudfoundryPhysicalContext) GetOrganizationOk() (*SnapshotPreview, bool) {
	if o == nil || o.Organization == nil {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *CloudfoundryPhysicalContext) HasOrganization() bool {
	if o != nil && o.Organization != nil {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given SnapshotPreview and assigns it to the Organization field.
func (o *CloudfoundryPhysicalContext) SetOrganization(v SnapshotPreview) {
	o.Organization = &v
}

// GetCfInstanceIndex returns the CfInstanceIndex field value if set, zero value otherwise.
func (o *CloudfoundryPhysicalContext) GetCfInstanceIndex() string {
	if o == nil || o.CfInstanceIndex == nil {
		var ret string
		return ret
	}
	return *o.CfInstanceIndex
}

// GetCfInstanceIndexOk returns a tuple with the CfInstanceIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudfoundryPhysicalContext) GetCfInstanceIndexOk() (*string, bool) {
	if o == nil || o.CfInstanceIndex == nil {
		return nil, false
	}
	return o.CfInstanceIndex, true
}

// HasCfInstanceIndex returns a boolean if a field has been set.
func (o *CloudfoundryPhysicalContext) HasCfInstanceIndex() bool {
	if o != nil && o.CfInstanceIndex != nil {
		return true
	}

	return false
}

// SetCfInstanceIndex gets a reference to the given string and assigns it to the CfInstanceIndex field.
func (o *CloudfoundryPhysicalContext) SetCfInstanceIndex(v string) {
	o.CfInstanceIndex = &v
}

func (o CloudfoundryPhysicalContext) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Application != nil {
		toSerialize["application"] = o.Application
	}
	if o.Space != nil {
		toSerialize["space"] = o.Space
	}
	if o.Organization != nil {
		toSerialize["organization"] = o.Organization
	}
	if o.CfInstanceIndex != nil {
		toSerialize["cfInstanceIndex"] = o.CfInstanceIndex
	}
	return json.Marshal(toSerialize)
}

type NullableCloudfoundryPhysicalContext struct {
	value *CloudfoundryPhysicalContext
	isSet bool
}

func (v NullableCloudfoundryPhysicalContext) Get() *CloudfoundryPhysicalContext {
	return v.value
}

func (v *NullableCloudfoundryPhysicalContext) Set(val *CloudfoundryPhysicalContext) {
	v.value = val
	v.isSet = true
}

func (v NullableCloudfoundryPhysicalContext) IsSet() bool {
	return v.isSet
}

func (v *NullableCloudfoundryPhysicalContext) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloudfoundryPhysicalContext(val *CloudfoundryPhysicalContext) *NullableCloudfoundryPhysicalContext {
	return &NullableCloudfoundryPhysicalContext{value: val, isSet: true}
}

func (v NullableCloudfoundryPhysicalContext) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloudfoundryPhysicalContext) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
