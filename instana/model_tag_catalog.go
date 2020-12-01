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

// TagCatalog struct for TagCatalog
type TagCatalog struct {
	TagTree []TagTreeLevel `json:"tagTree"`
	Tags    []Tag          `json:"tags"`
}

// NewTagCatalog instantiates a new TagCatalog object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTagCatalog(tagTree []TagTreeLevel, tags []Tag) *TagCatalog {
	this := TagCatalog{}
	this.TagTree = tagTree
	this.Tags = tags
	return &this
}

// NewTagCatalogWithDefaults instantiates a new TagCatalog object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTagCatalogWithDefaults() *TagCatalog {
	this := TagCatalog{}
	return &this
}

// GetTagTree returns the TagTree field value
func (o *TagCatalog) GetTagTree() []TagTreeLevel {
	if o == nil {
		var ret []TagTreeLevel
		return ret
	}

	return o.TagTree
}

// GetTagTreeOk returns a tuple with the TagTree field value
// and a boolean to check if the value has been set.
func (o *TagCatalog) GetTagTreeOk() (*[]TagTreeLevel, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TagTree, true
}

// SetTagTree sets field value
func (o *TagCatalog) SetTagTree(v []TagTreeLevel) {
	o.TagTree = v
}

// GetTags returns the Tags field value
func (o *TagCatalog) GetTags() []Tag {
	if o == nil {
		var ret []Tag
		return ret
	}

	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value
// and a boolean to check if the value has been set.
func (o *TagCatalog) GetTagsOk() (*[]Tag, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Tags, true
}

// SetTags sets field value
func (o *TagCatalog) SetTags(v []Tag) {
	o.Tags = v
}

func (o TagCatalog) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["tagTree"] = o.TagTree
	}
	if true {
		toSerialize["tags"] = o.Tags
	}
	return json.Marshal(toSerialize)
}

type NullableTagCatalog struct {
	value *TagCatalog
	isSet bool
}

func (v NullableTagCatalog) Get() *TagCatalog {
	return v.value
}

func (v *NullableTagCatalog) Set(val *TagCatalog) {
	v.value = val
	v.isSet = true
}

func (v NullableTagCatalog) IsSet() bool {
	return v.isSet
}

func (v *NullableTagCatalog) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTagCatalog(val *TagCatalog) *NullableTagCatalog {
	return &NullableTagCatalog{value: val, isSet: true}
}

func (v NullableTagCatalog) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTagCatalog) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
