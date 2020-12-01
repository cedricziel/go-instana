# GetCallGroups

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pagination** | Pointer to [**CursorPagination**](CursorPagination.md) |  | [optional] 
**Order** | Pointer to [**Order**](Order.md) |  | [optional] 
**TimeFrame** | Pointer to [**TimeFrame**](TimeFrame.md) |  | [optional] 
**Group** | [**Group**](Group.md) |  | 
**TagFilters** | Pointer to [**[]TagFilter**](TagFilter.md) |  | [optional] 
**Metrics** | [**[]MetricConfig**](MetricConfig.md) |  | 

## Methods

### NewGetCallGroups

`func NewGetCallGroups(group Group, metrics []MetricConfig, ) *GetCallGroups`

NewGetCallGroups instantiates a new GetCallGroups object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCallGroupsWithDefaults

`func NewGetCallGroupsWithDefaults() *GetCallGroups`

NewGetCallGroupsWithDefaults instantiates a new GetCallGroups object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPagination

`func (o *GetCallGroups) GetPagination() CursorPagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *GetCallGroups) GetPaginationOk() (*CursorPagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *GetCallGroups) SetPagination(v CursorPagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *GetCallGroups) HasPagination() bool`

HasPagination returns a boolean if a field has been set.

### GetOrder

`func (o *GetCallGroups) GetOrder() Order`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *GetCallGroups) GetOrderOk() (*Order, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *GetCallGroups) SetOrder(v Order)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *GetCallGroups) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetTimeFrame

`func (o *GetCallGroups) GetTimeFrame() TimeFrame`

GetTimeFrame returns the TimeFrame field if non-nil, zero value otherwise.

### GetTimeFrameOk

`func (o *GetCallGroups) GetTimeFrameOk() (*TimeFrame, bool)`

GetTimeFrameOk returns a tuple with the TimeFrame field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeFrame

`func (o *GetCallGroups) SetTimeFrame(v TimeFrame)`

SetTimeFrame sets TimeFrame field to given value.

### HasTimeFrame

`func (o *GetCallGroups) HasTimeFrame() bool`

HasTimeFrame returns a boolean if a field has been set.

### GetGroup

`func (o *GetCallGroups) GetGroup() Group`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *GetCallGroups) GetGroupOk() (*Group, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *GetCallGroups) SetGroup(v Group)`

SetGroup sets Group field to given value.


### GetTagFilters

`func (o *GetCallGroups) GetTagFilters() []TagFilter`

GetTagFilters returns the TagFilters field if non-nil, zero value otherwise.

### GetTagFiltersOk

`func (o *GetCallGroups) GetTagFiltersOk() (*[]TagFilter, bool)`

GetTagFiltersOk returns a tuple with the TagFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagFilters

`func (o *GetCallGroups) SetTagFilters(v []TagFilter)`

SetTagFilters sets TagFilters field to given value.

### HasTagFilters

`func (o *GetCallGroups) HasTagFilters() bool`

HasTagFilters returns a boolean if a field has been set.

### GetMetrics

`func (o *GetCallGroups) GetMetrics() []MetricConfig`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *GetCallGroups) GetMetricsOk() (*[]MetricConfig, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *GetCallGroups) SetMetrics(v []MetricConfig)`

SetMetrics sets Metrics field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


