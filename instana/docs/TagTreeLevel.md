# TagTreeLevel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Label** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**Icon** | Pointer to **string** |  | [optional] 
**Searchable** | **bool** |  | 
**Children** | [**[]TagTreeNode**](TagTreeNode.md) |  | 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewTagTreeLevel

`func NewTagTreeLevel(label string, searchable bool, children []TagTreeNode, ) *TagTreeLevel`

NewTagTreeLevel instantiates a new TagTreeLevel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTagTreeLevelWithDefaults

`func NewTagTreeLevelWithDefaults() *TagTreeLevel`

NewTagTreeLevelWithDefaults instantiates a new TagTreeLevel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabel

`func (o *TagTreeLevel) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *TagTreeLevel) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *TagTreeLevel) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetDescription

`func (o *TagTreeLevel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TagTreeLevel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TagTreeLevel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TagTreeLevel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIcon

`func (o *TagTreeLevel) GetIcon() string`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *TagTreeLevel) GetIconOk() (*string, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *TagTreeLevel) SetIcon(v string)`

SetIcon sets Icon field to given value.

### HasIcon

`func (o *TagTreeLevel) HasIcon() bool`

HasIcon returns a boolean if a field has been set.

### GetSearchable

`func (o *TagTreeLevel) GetSearchable() bool`

GetSearchable returns the Searchable field if non-nil, zero value otherwise.

### GetSearchableOk

`func (o *TagTreeLevel) GetSearchableOk() (*bool, bool)`

GetSearchableOk returns a tuple with the Searchable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchable

`func (o *TagTreeLevel) SetSearchable(v bool)`

SetSearchable sets Searchable field to given value.


### GetChildren

`func (o *TagTreeLevel) GetChildren() []TagTreeNode`

GetChildren returns the Children field if non-nil, zero value otherwise.

### GetChildrenOk

`func (o *TagTreeLevel) GetChildrenOk() (*[]TagTreeNode, bool)`

GetChildrenOk returns a tuple with the Children field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildren

`func (o *TagTreeLevel) SetChildren(v []TagTreeNode)`

SetChildren sets Children field to given value.


### GetType

`func (o *TagTreeLevel) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TagTreeLevel) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TagTreeLevel) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *TagTreeLevel) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


