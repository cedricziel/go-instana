# CustomEventSpecification

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

## Methods

### NewCustomEventSpecification

`func NewCustomEventSpecification(id string, name string, entityType string, rules []AbstractRule, ) *CustomEventSpecification`

NewCustomEventSpecification instantiates a new CustomEventSpecification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomEventSpecificationWithDefaults

`func NewCustomEventSpecificationWithDefaults() *CustomEventSpecification`

NewCustomEventSpecificationWithDefaults instantiates a new CustomEventSpecification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CustomEventSpecification) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CustomEventSpecification) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CustomEventSpecification) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *CustomEventSpecification) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CustomEventSpecification) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CustomEventSpecification) SetName(v string)`

SetName sets Name field to given value.


### GetEntityType

`func (o *CustomEventSpecification) GetEntityType() string`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *CustomEventSpecification) GetEntityTypeOk() (*string, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *CustomEventSpecification) SetEntityType(v string)`

SetEntityType sets EntityType field to given value.


### GetQuery

`func (o *CustomEventSpecification) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *CustomEventSpecification) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *CustomEventSpecification) SetQuery(v string)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *CustomEventSpecification) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### GetTriggering

`func (o *CustomEventSpecification) GetTriggering() bool`

GetTriggering returns the Triggering field if non-nil, zero value otherwise.

### GetTriggeringOk

`func (o *CustomEventSpecification) GetTriggeringOk() (*bool, bool)`

GetTriggeringOk returns a tuple with the Triggering field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggering

`func (o *CustomEventSpecification) SetTriggering(v bool)`

SetTriggering sets Triggering field to given value.

### HasTriggering

`func (o *CustomEventSpecification) HasTriggering() bool`

HasTriggering returns a boolean if a field has been set.

### GetDescription

`func (o *CustomEventSpecification) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CustomEventSpecification) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CustomEventSpecification) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CustomEventSpecification) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExpirationTime

`func (o *CustomEventSpecification) GetExpirationTime() int64`

GetExpirationTime returns the ExpirationTime field if non-nil, zero value otherwise.

### GetExpirationTimeOk

`func (o *CustomEventSpecification) GetExpirationTimeOk() (*int64, bool)`

GetExpirationTimeOk returns a tuple with the ExpirationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationTime

`func (o *CustomEventSpecification) SetExpirationTime(v int64)`

SetExpirationTime sets ExpirationTime field to given value.

### HasExpirationTime

`func (o *CustomEventSpecification) HasExpirationTime() bool`

HasExpirationTime returns a boolean if a field has been set.

### GetEnabled

`func (o *CustomEventSpecification) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CustomEventSpecification) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CustomEventSpecification) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *CustomEventSpecification) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetRules

`func (o *CustomEventSpecification) GetRules() []AbstractRule`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *CustomEventSpecification) GetRulesOk() (*[]AbstractRule, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *CustomEventSpecification) SetRules(v []AbstractRule)`

SetRules sets Rules field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


