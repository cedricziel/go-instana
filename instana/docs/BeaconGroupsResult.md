# BeaconGroupsResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | [**[]WebsiteBeaconGroupsItem**](WebsiteBeaconGroupsItem.md) |  | 
**CanLoadMore** | Pointer to **bool** |  | [optional] 
**TotalHits** | Pointer to **int64** |  | [optional] 
**TotalRepresentedItemCount** | Pointer to **int64** |  | [optional] 

## Methods

### NewBeaconGroupsResult

`func NewBeaconGroupsResult(items []WebsiteBeaconGroupsItem, ) *BeaconGroupsResult`

NewBeaconGroupsResult instantiates a new BeaconGroupsResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBeaconGroupsResultWithDefaults

`func NewBeaconGroupsResultWithDefaults() *BeaconGroupsResult`

NewBeaconGroupsResultWithDefaults instantiates a new BeaconGroupsResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *BeaconGroupsResult) GetItems() []WebsiteBeaconGroupsItem`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *BeaconGroupsResult) GetItemsOk() (*[]WebsiteBeaconGroupsItem, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *BeaconGroupsResult) SetItems(v []WebsiteBeaconGroupsItem)`

SetItems sets Items field to given value.


### GetCanLoadMore

`func (o *BeaconGroupsResult) GetCanLoadMore() bool`

GetCanLoadMore returns the CanLoadMore field if non-nil, zero value otherwise.

### GetCanLoadMoreOk

`func (o *BeaconGroupsResult) GetCanLoadMoreOk() (*bool, bool)`

GetCanLoadMoreOk returns a tuple with the CanLoadMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanLoadMore

`func (o *BeaconGroupsResult) SetCanLoadMore(v bool)`

SetCanLoadMore sets CanLoadMore field to given value.

### HasCanLoadMore

`func (o *BeaconGroupsResult) HasCanLoadMore() bool`

HasCanLoadMore returns a boolean if a field has been set.

### GetTotalHits

`func (o *BeaconGroupsResult) GetTotalHits() int64`

GetTotalHits returns the TotalHits field if non-nil, zero value otherwise.

### GetTotalHitsOk

`func (o *BeaconGroupsResult) GetTotalHitsOk() (*int64, bool)`

GetTotalHitsOk returns a tuple with the TotalHits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalHits

`func (o *BeaconGroupsResult) SetTotalHits(v int64)`

SetTotalHits sets TotalHits field to given value.

### HasTotalHits

`func (o *BeaconGroupsResult) HasTotalHits() bool`

HasTotalHits returns a boolean if a field has been set.

### GetTotalRepresentedItemCount

`func (o *BeaconGroupsResult) GetTotalRepresentedItemCount() int64`

GetTotalRepresentedItemCount returns the TotalRepresentedItemCount field if non-nil, zero value otherwise.

### GetTotalRepresentedItemCountOk

`func (o *BeaconGroupsResult) GetTotalRepresentedItemCountOk() (*int64, bool)`

GetTotalRepresentedItemCountOk returns a tuple with the TotalRepresentedItemCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalRepresentedItemCount

`func (o *BeaconGroupsResult) SetTotalRepresentedItemCount(v int64)`

SetTotalRepresentedItemCount sets TotalRepresentedItemCount field to given value.

### HasTotalRepresentedItemCount

`func (o *BeaconGroupsResult) HasTotalRepresentedItemCount() bool`

HasTotalRepresentedItemCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


