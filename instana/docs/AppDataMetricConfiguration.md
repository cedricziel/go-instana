# AppDataMetricConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metric** | **string** |  | 
**Granularity** | Pointer to **int32** |  | [optional] 
**Aggregation** | **string** |  | 

## Methods

### NewAppDataMetricConfiguration

`func NewAppDataMetricConfiguration(metric string, aggregation string, ) *AppDataMetricConfiguration`

NewAppDataMetricConfiguration instantiates a new AppDataMetricConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppDataMetricConfigurationWithDefaults

`func NewAppDataMetricConfigurationWithDefaults() *AppDataMetricConfiguration`

NewAppDataMetricConfigurationWithDefaults instantiates a new AppDataMetricConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetric

`func (o *AppDataMetricConfiguration) GetMetric() string`

GetMetric returns the Metric field if non-nil, zero value otherwise.

### GetMetricOk

`func (o *AppDataMetricConfiguration) GetMetricOk() (*string, bool)`

GetMetricOk returns a tuple with the Metric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetric

`func (o *AppDataMetricConfiguration) SetMetric(v string)`

SetMetric sets Metric field to given value.


### GetGranularity

`func (o *AppDataMetricConfiguration) GetGranularity() int32`

GetGranularity returns the Granularity field if non-nil, zero value otherwise.

### GetGranularityOk

`func (o *AppDataMetricConfiguration) GetGranularityOk() (*int32, bool)`

GetGranularityOk returns a tuple with the Granularity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGranularity

`func (o *AppDataMetricConfiguration) SetGranularity(v int32)`

SetGranularity sets Granularity field to given value.

### HasGranularity

`func (o *AppDataMetricConfiguration) HasGranularity() bool`

HasGranularity returns a boolean if a field has been set.

### GetAggregation

`func (o *AppDataMetricConfiguration) GetAggregation() string`

GetAggregation returns the Aggregation field if non-nil, zero value otherwise.

### GetAggregationOk

`func (o *AppDataMetricConfiguration) GetAggregationOk() (*string, bool)`

GetAggregationOk returns a tuple with the Aggregation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregation

`func (o *AppDataMetricConfiguration) SetAggregation(v string)`

SetAggregation sets Aggregation field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


