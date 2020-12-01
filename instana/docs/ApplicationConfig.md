# ApplicationConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Label** | **string** |  | 
**MatchSpecification** | [**MatchExpressionDTO**](MatchExpressionDTO.md) |  | 
**Scope** | **string** |  | 
**BoundaryScope** | **string** |  | 

## Methods

### NewApplicationConfig

`func NewApplicationConfig(id string, label string, matchSpecification MatchExpressionDTO, scope string, boundaryScope string, ) *ApplicationConfig`

NewApplicationConfig instantiates a new ApplicationConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationConfigWithDefaults

`func NewApplicationConfigWithDefaults() *ApplicationConfig`

NewApplicationConfigWithDefaults instantiates a new ApplicationConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ApplicationConfig) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApplicationConfig) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApplicationConfig) SetId(v string)`

SetId sets Id field to given value.


### GetLabel

`func (o *ApplicationConfig) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *ApplicationConfig) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *ApplicationConfig) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetMatchSpecification

`func (o *ApplicationConfig) GetMatchSpecification() MatchExpressionDTO`

GetMatchSpecification returns the MatchSpecification field if non-nil, zero value otherwise.

### GetMatchSpecificationOk

`func (o *ApplicationConfig) GetMatchSpecificationOk() (*MatchExpressionDTO, bool)`

GetMatchSpecificationOk returns a tuple with the MatchSpecification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchSpecification

`func (o *ApplicationConfig) SetMatchSpecification(v MatchExpressionDTO)`

SetMatchSpecification sets MatchSpecification field to given value.


### GetScope

`func (o *ApplicationConfig) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *ApplicationConfig) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *ApplicationConfig) SetScope(v string)`

SetScope sets Scope field to given value.


### GetBoundaryScope

`func (o *ApplicationConfig) GetBoundaryScope() string`

GetBoundaryScope returns the BoundaryScope field if non-nil, zero value otherwise.

### GetBoundaryScopeOk

`func (o *ApplicationConfig) GetBoundaryScopeOk() (*string, bool)`

GetBoundaryScopeOk returns a tuple with the BoundaryScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundaryScope

`func (o *ApplicationConfig) SetBoundaryScope(v string)`

SetBoundaryScope sets BoundaryScope field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


