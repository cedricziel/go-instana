/*
 * Introduction to Instana public APIs
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.192.86
 * Contact: support@instana.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package instana

import (
	"encoding/json"
)

// HistoricBaseline struct for HistoricBaseline
type HistoricBaseline struct {
	Threshold
	Seasonality     string      `json:"seasonality"`
	Baseline        [][]float32 `json:"baseline"`
	DeviationFactor *float64    `json:"deviationFactor,omitempty"`
}

// NewHistoricBaseline instantiates a new HistoricBaseline object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHistoricBaseline(seasonality string, baseline [][]float32) *HistoricBaseline {
	this := HistoricBaseline{}
	this.Seasonality = seasonality
	this.Baseline = baseline
	return &this
}

// NewHistoricBaselineWithDefaults instantiates a new HistoricBaseline object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHistoricBaselineWithDefaults() *HistoricBaseline {
	this := HistoricBaseline{}
	return &this
}

// GetSeasonality returns the Seasonality field value
func (o *HistoricBaseline) GetSeasonality() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Seasonality
}

// GetSeasonalityOk returns a tuple with the Seasonality field value
// and a boolean to check if the value has been set.
func (o *HistoricBaseline) GetSeasonalityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Seasonality, true
}

// SetSeasonality sets field value
func (o *HistoricBaseline) SetSeasonality(v string) {
	o.Seasonality = v
}

// GetBaseline returns the Baseline field value
func (o *HistoricBaseline) GetBaseline() [][]float32 {
	if o == nil {
		var ret [][]float32
		return ret
	}

	return o.Baseline
}

// GetBaselineOk returns a tuple with the Baseline field value
// and a boolean to check if the value has been set.
func (o *HistoricBaseline) GetBaselineOk() (*[][]float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Baseline, true
}

// SetBaseline sets field value
func (o *HistoricBaseline) SetBaseline(v [][]float32) {
	o.Baseline = v
}

// GetDeviationFactor returns the DeviationFactor field value if set, zero value otherwise.
func (o *HistoricBaseline) GetDeviationFactor() float64 {
	if o == nil || o.DeviationFactor == nil {
		var ret float64
		return ret
	}
	return *o.DeviationFactor
}

// GetDeviationFactorOk returns a tuple with the DeviationFactor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HistoricBaseline) GetDeviationFactorOk() (*float64, bool) {
	if o == nil || o.DeviationFactor == nil {
		return nil, false
	}
	return o.DeviationFactor, true
}

// HasDeviationFactor returns a boolean if a field has been set.
func (o *HistoricBaseline) HasDeviationFactor() bool {
	if o != nil && o.DeviationFactor != nil {
		return true
	}

	return false
}

// SetDeviationFactor gets a reference to the given float64 and assigns it to the DeviationFactor field.
func (o *HistoricBaseline) SetDeviationFactor(v float64) {
	o.DeviationFactor = &v
}

func (o HistoricBaseline) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedThreshold, errThreshold := json.Marshal(o.Threshold)
	if errThreshold != nil {
		return []byte{}, errThreshold
	}
	errThreshold = json.Unmarshal([]byte(serializedThreshold), &toSerialize)
	if errThreshold != nil {
		return []byte{}, errThreshold
	}
	if true {
		toSerialize["seasonality"] = o.Seasonality
	}
	if true {
		toSerialize["baseline"] = o.Baseline
	}
	if o.DeviationFactor != nil {
		toSerialize["deviationFactor"] = o.DeviationFactor
	}
	return json.Marshal(toSerialize)
}

type NullableHistoricBaseline struct {
	value *HistoricBaseline
	isSet bool
}

func (v NullableHistoricBaseline) Get() *HistoricBaseline {
	return v.value
}

func (v *NullableHistoricBaseline) Set(val *HistoricBaseline) {
	v.value = val
	v.isSet = true
}

func (v NullableHistoricBaseline) IsSet() bool {
	return v.isSet
}

func (v *NullableHistoricBaseline) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHistoricBaseline(val *HistoricBaseline) *NullableHistoricBaseline {
	return &NullableHistoricBaseline{value: val, isSet: true}
}

func (v NullableHistoricBaseline) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHistoricBaseline) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
