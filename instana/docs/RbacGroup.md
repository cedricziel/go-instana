# RbacGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | **string** |  | 
**Members** | [**[]Member**](Member.md) |  | 
**PermissionSet** | Pointer to [**PermissionSet**](PermissionSet.md) |  | [optional] 

## Methods

### NewRbacGroup

`func NewRbacGroup(name string, members []Member, ) *RbacGroup`

NewRbacGroup instantiates a new RbacGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRbacGroupWithDefaults

`func NewRbacGroupWithDefaults() *RbacGroup`

NewRbacGroupWithDefaults instantiates a new RbacGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RbacGroup) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RbacGroup) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RbacGroup) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RbacGroup) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *RbacGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RbacGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RbacGroup) SetName(v string)`

SetName sets Name field to given value.


### GetMembers

`func (o *RbacGroup) GetMembers() []Member`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *RbacGroup) GetMembersOk() (*[]Member, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *RbacGroup) SetMembers(v []Member)`

SetMembers sets Members field to given value.


### GetPermissionSet

`func (o *RbacGroup) GetPermissionSet() PermissionSet`

GetPermissionSet returns the PermissionSet field if non-nil, zero value otherwise.

### GetPermissionSetOk

`func (o *RbacGroup) GetPermissionSetOk() (*PermissionSet, bool)`

GetPermissionSetOk returns a tuple with the PermissionSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionSet

`func (o *RbacGroup) SetPermissionSet(v PermissionSet)`

SetPermissionSet sets PermissionSet field to given value.

### HasPermissionSet

`func (o *RbacGroup) HasPermissionSet() bool`

HasPermissionSet returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


