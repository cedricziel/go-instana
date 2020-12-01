# HistoricBaseline

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Seasonality** | **string** |  | 
**Baseline** | [**[][]float32**](array.md) |  | 
**DeviationFactor** | Pointer to **float64** |  | [optional] 

## Methods

### NewHistoricBaseline

`func NewHistoricBaseline(seasonality string, baseline [][]float32, ) *HistoricBaseline`

NewHistoricBaseline instantiates a new HistoricBaseline object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHistoricBaselineWithDefaults

`func NewHistoricBaselineWithDefaults() *HistoricBaseline`

NewHistoricBaselineWithDefaults instantiates a new HistoricBaseline object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSeasonality

`func (o *HistoricBaseline) GetSeasonality() string`

GetSeasonality returns the Seasonality field if non-nil, zero value otherwise.

### GetSeasonalityOk

`func (o *HistoricBaseline) GetSeasonalityOk() (*string, bool)`

GetSeasonalityOk returns a tuple with the Seasonality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeasonality

`func (o *HistoricBaseline) SetSeasonality(v string)`

SetSeasonality sets Seasonality field to given value.


### GetBaseline

`func (o *HistoricBaseline) GetBaseline() [][]float32`

GetBaseline returns the Baseline field if non-nil, zero value otherwise.

### GetBaselineOk

`func (o *HistoricBaseline) GetBaselineOk() (*[][]float32, bool)`

GetBaselineOk returns a tuple with the Baseline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseline

`func (o *HistoricBaseline) SetBaseline(v [][]float32)`

SetBaseline sets Baseline field to given value.


### GetDeviationFactor

`func (o *HistoricBaseline) GetDeviationFactor() float64`

GetDeviationFactor returns the DeviationFactor field if non-nil, zero value otherwise.

### GetDeviationFactorOk

`func (o *HistoricBaseline) GetDeviationFactorOk() (*float64, bool)`

GetDeviationFactorOk returns a tuple with the DeviationFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviationFactor

`func (o *HistoricBaseline) SetDeviationFactor(v float64)`

SetDeviationFactor sets DeviationFactor field to given value.

### HasDeviationFactor

`func (o *HistoricBaseline) HasDeviationFactor() bool`

HasDeviationFactor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


