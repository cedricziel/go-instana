# TreeNode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SnapshotId** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Children** | Pointer to [**[]TreeNode**](TreeNode.md) |  | [optional] 

## Methods

### NewTreeNode

`func NewTreeNode() *TreeNode`

NewTreeNode instantiates a new TreeNode object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTreeNodeWithDefaults

`func NewTreeNodeWithDefaults() *TreeNode`

NewTreeNodeWithDefaults instantiates a new TreeNode object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSnapshotId

`func (o *TreeNode) GetSnapshotId() string`

GetSnapshotId returns the SnapshotId field if non-nil, zero value otherwise.

### GetSnapshotIdOk

`func (o *TreeNode) GetSnapshotIdOk() (*string, bool)`

GetSnapshotIdOk returns a tuple with the SnapshotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotId

`func (o *TreeNode) SetSnapshotId(v string)`

SetSnapshotId sets SnapshotId field to given value.

### HasSnapshotId

`func (o *TreeNode) HasSnapshotId() bool`

HasSnapshotId returns a boolean if a field has been set.

### GetType

`func (o *TreeNode) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TreeNode) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TreeNode) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *TreeNode) HasType() bool`

HasType returns a boolean if a field has been set.

### GetChildren

`func (o *TreeNode) GetChildren() []TreeNode`

GetChildren returns the Children field if non-nil, zero value otherwise.

### GetChildrenOk

`func (o *TreeNode) GetChildrenOk() (*[]TreeNode, bool)`

GetChildrenOk returns a tuple with the Children field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildren

`func (o *TreeNode) SetChildren(v []TreeNode)`

SetChildren sets Children field to given value.

### HasChildren

`func (o *TreeNode) HasChildren() bool`

HasChildren returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


