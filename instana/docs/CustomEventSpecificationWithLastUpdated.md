# CustomEventSpecificationWithLastUpdated

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
**EntityType** | **string** |  | 
**Query** | Pointer to **string** |  | [optional] 
**Triggering** | Pointer to **bool** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**ExpirationTime** | Pointer to **int64** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Rules** | [**[]AbstractRule**](AbstractRule.md) |  | 
**LastUpdated** | Pointer to **int64** |  | [optional] 

## Methods

### NewCustomEventSpecificationWithLastUpdated

`func NewCustomEventSpecificationWithLastUpdated(id string, name string, entityType string, rules []AbstractRule, ) *CustomEventSpecificationWithLastUpdated`

NewCustomEventSpecificationWithLastUpdated instantiates a new CustomEventSpecificationWithLastUpdated object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomEventSpecificationWithLastUpdatedWithDefaults

`func NewCustomEventSpecificationWithLastUpdatedWithDefaults() *CustomEventSpecificationWithLastUpdated`

NewCustomEventSpecificationWithLastUpdatedWithDefaults instantiates a new CustomEventSpecificationWithLastUpdated object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CustomEventSpecificationWithLastUpdated) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CustomEventSpecificationWithLastUpdated) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CustomEventSpecificationWithLastUpdated) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *CustomEventSpecificationWithLastUpdated) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CustomEventSpecificationWithLastUpdated) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CustomEventSpecificationWithLastUpdated) SetName(v string)`

SetName sets Name field to given value.


### GetEntityType

`func (o *CustomEventSpecificationWithLastUpdated) GetEntityType() string`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *CustomEventSpecificationWithLastUpdated) GetEntityTypeOk() (*string, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *CustomEventSpecificationWithLastUpdated) SetEntityType(v string)`

SetEntityType sets EntityType field to given value.


### GetQuery

`func (o *CustomEventSpecificationWithLastUpdated) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *CustomEventSpecificationWithLastUpdated) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *CustomEventSpecificationWithLastUpdated) SetQuery(v string)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *CustomEventSpecificationWithLastUpdated) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### GetTriggering

`func (o *CustomEventSpecificationWithLastUpdated) GetTriggering() bool`

GetTriggering returns the Triggering field if non-nil, zero value otherwise.

### GetTriggeringOk

`func (o *CustomEventSpecificationWithLastUpdated) GetTriggeringOk() (*bool, bool)`

GetTriggeringOk returns a tuple with the Triggering field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggering

`func (o *CustomEventSpecificationWithLastUpdated) SetTriggering(v bool)`

SetTriggering sets Triggering field to given value.

### HasTriggering

`func (o *CustomEventSpecificationWithLastUpdated) HasTriggering() bool`

HasTriggering returns a boolean if a field has been set.

### GetDescription

`func (o *CustomEventSpecificationWithLastUpdated) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CustomEventSpecificationWithLastUpdated) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CustomEventSpecificationWithLastUpdated) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CustomEventSpecificationWithLastUpdated) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExpirationTime

`func (o *CustomEventSpecificationWithLastUpdated) GetExpirationTime() int64`

GetExpirationTime returns the ExpirationTime field if non-nil, zero value otherwise.

### GetExpirationTimeOk

`func (o *CustomEventSpecificationWithLastUpdated) GetExpirationTimeOk() (*int64, bool)`

GetExpirationTimeOk returns a tuple with the ExpirationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationTime

`func (o *CustomEventSpecificationWithLastUpdated) SetExpirationTime(v int64)`

SetExpirationTime sets ExpirationTime field to given value.

### HasExpirationTime

`func (o *CustomEventSpecificationWithLastUpdated) HasExpirationTime() bool`

HasExpirationTime returns a boolean if a field has been set.

### GetEnabled

`func (o *CustomEventSpecificationWithLastUpdated) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CustomEventSpecificationWithLastUpdated) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CustomEventSpecificationWithLastUpdated) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *CustomEventSpecificationWithLastUpdated) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetRules

`func (o *CustomEventSpecificationWithLastUpdated) GetRules() []AbstractRule`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *CustomEventSpecificationWithLastUpdated) GetRulesOk() (*[]AbstractRule, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *CustomEventSpecificationWithLastUpdated) SetRules(v []AbstractRule)`

SetRules sets Rules field to given value.


### GetLastUpdated

`func (o *CustomEventSpecificationWithLastUpdated) GetLastUpdated() int64`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *CustomEventSpecificationWithLastUpdated) GetLastUpdatedOk() (*int64, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *CustomEventSpecificationWithLastUpdated) SetLastUpdated(v int64)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *CustomEventSpecificationWithLastUpdated) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


