# ConfigVersion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Created** | Pointer to **int64** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 

## Methods

### NewConfigVersion

`func NewConfigVersion(id string, ) *ConfigVersion`

NewConfigVersion instantiates a new ConfigVersion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigVersionWithDefaults

`func NewConfigVersionWithDefaults() *ConfigVersion`

NewConfigVersionWithDefaults instantiates a new ConfigVersion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ConfigVersion) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ConfigVersion) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ConfigVersion) SetId(v string)`

SetId sets Id field to given value.


### GetCreated

`func (o *ConfigVersion) GetCreated() int64`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ConfigVersion) GetCreatedOk() (*int64, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *ConfigVersion) SetCreated(v int64)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *ConfigVersion) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetEnabled

`func (o *ConfigVersion) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ConfigVersion) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ConfigVersion) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ConfigVersion) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


