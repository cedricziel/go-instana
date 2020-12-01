# ScopedGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | **string** |  | 
**Members** | [**[]GroupMember**](GroupMember.md) |  | 
**Permissions** | [**[]GroupPermission**](GroupPermission.md) |  | 

## Methods

### NewScopedGroup

`func NewScopedGroup(name string, members []GroupMember, permissions []GroupPermission, ) *ScopedGroup`

NewScopedGroup instantiates a new ScopedGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScopedGroupWithDefaults

`func NewScopedGroupWithDefaults() *ScopedGroup`

NewScopedGroupWithDefaults instantiates a new ScopedGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ScopedGroup) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ScopedGroup) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ScopedGroup) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ScopedGroup) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ScopedGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ScopedGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ScopedGroup) SetName(v string)`

SetName sets Name field to given value.


### GetMembers

`func (o *ScopedGroup) GetMembers() []GroupMember`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *ScopedGroup) GetMembersOk() (*[]GroupMember, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *ScopedGroup) SetMembers(v []GroupMember)`

SetMembers sets Members field to given value.


### GetPermissions

`func (o *ScopedGroup) GetPermissions() []GroupPermission`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *ScopedGroup) GetPermissionsOk() (*[]GroupPermission, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *ScopedGroup) SetPermissions(v []GroupPermission)`

SetPermissions sets Permissions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


