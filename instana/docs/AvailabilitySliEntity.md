# AvailabilitySliEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicationId** | Pointer to **string** |  | [optional] 
**ServiceId** | Pointer to **string** |  | [optional] 
**EndpointId** | Pointer to **string** |  | [optional] 
**BoundaryScope** | **string** |  | 
**GoodEventFilters** | [**[]TagFilter**](TagFilter.md) |  | 
**BadEventFilters** | [**[]TagFilter**](TagFilter.md) |  | 

## Methods

### NewAvailabilitySliEntity

`func NewAvailabilitySliEntity(boundaryScope string, goodEventFilters []TagFilter, badEventFilters []TagFilter, ) *AvailabilitySliEntity`

NewAvailabilitySliEntity instantiates a new AvailabilitySliEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAvailabilitySliEntityWithDefaults

`func NewAvailabilitySliEntityWithDefaults() *AvailabilitySliEntity`

NewAvailabilitySliEntityWithDefaults instantiates a new AvailabilitySliEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplicationId

`func (o *AvailabilitySliEntity) GetApplicationId() string`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *AvailabilitySliEntity) GetApplicationIdOk() (*string, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *AvailabilitySliEntity) SetApplicationId(v string)`

SetApplicationId sets ApplicationId field to given value.

### HasApplicationId

`func (o *AvailabilitySliEntity) HasApplicationId() bool`

HasApplicationId returns a boolean if a field has been set.

### GetServiceId

`func (o *AvailabilitySliEntity) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *AvailabilitySliEntity) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *AvailabilitySliEntity) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.

### HasServiceId

`func (o *AvailabilitySliEntity) HasServiceId() bool`

HasServiceId returns a boolean if a field has been set.

### GetEndpointId

`func (o *AvailabilitySliEntity) GetEndpointId() string`

GetEndpointId returns the EndpointId field if non-nil, zero value otherwise.

### GetEndpointIdOk

`func (o *AvailabilitySliEntity) GetEndpointIdOk() (*string, bool)`

GetEndpointIdOk returns a tuple with the EndpointId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointId

`func (o *AvailabilitySliEntity) SetEndpointId(v string)`

SetEndpointId sets EndpointId field to given value.

### HasEndpointId

`func (o *AvailabilitySliEntity) HasEndpointId() bool`

HasEndpointId returns a boolean if a field has been set.

### GetBoundaryScope

`func (o *AvailabilitySliEntity) GetBoundaryScope() string`

GetBoundaryScope returns the BoundaryScope field if non-nil, zero value otherwise.

### GetBoundaryScopeOk

`func (o *AvailabilitySliEntity) GetBoundaryScopeOk() (*string, bool)`

GetBoundaryScopeOk returns a tuple with the BoundaryScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundaryScope

`func (o *AvailabilitySliEntity) SetBoundaryScope(v string)`

SetBoundaryScope sets BoundaryScope field to given value.


### GetGoodEventFilters

`func (o *AvailabilitySliEntity) GetGoodEventFilters() []TagFilter`

GetGoodEventFilters returns the GoodEventFilters field if non-nil, zero value otherwise.

### GetGoodEventFiltersOk

`func (o *AvailabilitySliEntity) GetGoodEventFiltersOk() (*[]TagFilter, bool)`

GetGoodEventFiltersOk returns a tuple with the GoodEventFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoodEventFilters

`func (o *AvailabilitySliEntity) SetGoodEventFilters(v []TagFilter)`

SetGoodEventFilters sets GoodEventFilters field to given value.


### GetBadEventFilters

`func (o *AvailabilitySliEntity) GetBadEventFilters() []TagFilter`

GetBadEventFilters returns the BadEventFilters field if non-nil, zero value otherwise.

### GetBadEventFiltersOk

`func (o *AvailabilitySliEntity) GetBadEventFiltersOk() (*[]TagFilter, bool)`

GetBadEventFiltersOk returns a tuple with the BadEventFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBadEventFilters

`func (o *AvailabilitySliEntity) SetBadEventFilters(v []TagFilter)`

SetBadEventFilters sets BadEventFilters field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


