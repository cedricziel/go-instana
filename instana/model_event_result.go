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

// EventResult struct for EventResult
type EventResult struct {
	EventId        *string                   `json:"eventId,omitempty"`
	Start          *int64                    `json:"start,omitempty"`
	End            *int64                    `json:"end,omitempty"`
	TriggeringTime *int64                    `json:"triggeringTime,omitempty"`
	Type           *string                   `json:"type,omitempty"`
	State          *string                   `json:"state,omitempty"`
	Problem        *string                   `json:"problem,omitempty"`
	FixSuggestion  *string                   `json:"fixSuggestion,omitempty"`
	Severity       *int32                    `json:"severity,omitempty"`
	SnapshotId     *string                   `json:"snapshotId,omitempty"`
	Metrics        *[]map[string]interface{} `json:"metrics,omitempty"`
	RecentEvents   *[]map[string]interface{} `json:"recentEvents,omitempty"`
}

// NewEventResult instantiates a new EventResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEventResult() *EventResult {
	this := EventResult{}
	return &this
}

// NewEventResultWithDefaults instantiates a new EventResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEventResultWithDefaults() *EventResult {
	this := EventResult{}
	return &this
}

// GetEventId returns the EventId field value if set, zero value otherwise.
func (o *EventResult) GetEventId() string {
	if o == nil || o.EventId == nil {
		var ret string
		return ret
	}
	return *o.EventId
}

// GetEventIdOk returns a tuple with the EventId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventResult) GetEventIdOk() (*string, bool) {
	if o == nil || o.EventId == nil {
		return nil, false
	}
	return o.EventId, true
}

// HasEventId returns a boolean if a field has been set.
func (o *EventResult) HasEventId() bool {
	if o != nil && o.EventId != nil {
		return true
	}

	return false
}

// SetEventId gets a reference to the given string and assigns it to the EventId field.
func (o *EventResult) SetEventId(v string) {
	o.EventId = &v
}

// GetStart returns the Start field value if set, zero value otherwise.
func (o *EventResult) GetStart() int64 {
	if o == nil || o.Start == nil {
		var ret int64
		return ret
	}
	return *o.Start
}

// GetStartOk returns a tuple with the Start field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventResult) GetStartOk() (*int64, bool) {
	if o == nil || o.Start == nil {
		return nil, false
	}
	return o.Start, true
}

// HasStart returns a boolean if a field has been set.
func (o *EventResult) HasStart() bool {
	if o != nil && o.Start != nil {
		return true
	}

	return false
}

// SetStart gets a reference to the given int64 and assigns it to the Start field.
func (o *EventResult) SetStart(v int64) {
	o.Start = &v
}

// GetEnd returns the End field value if set, zero value otherwise.
func (o *EventResult) GetEnd() int64 {
	if o == nil || o.End == nil {
		var ret int64
		return ret
	}
	return *o.End
}

// GetEndOk returns a tuple with the End field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventResult) GetEndOk() (*int64, bool) {
	if o == nil || o.End == nil {
		return nil, false
	}
	return o.End, true
}

// HasEnd returns a boolean if a field has been set.
func (o *EventResult) HasEnd() bool {
	if o != nil && o.End != nil {
		return true
	}

	return false
}

// SetEnd gets a reference to the given int64 and assigns it to the End field.
func (o *EventResult) SetEnd(v int64) {
	o.End = &v
}

// GetTriggeringTime returns the TriggeringTime field value if set, zero value otherwise.
func (o *EventResult) GetTriggeringTime() int64 {
	if o == nil || o.TriggeringTime == nil {
		var ret int64
		return ret
	}
	return *o.TriggeringTime
}

// GetTriggeringTimeOk returns a tuple with the TriggeringTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventResult) GetTriggeringTimeOk() (*int64, bool) {
	if o == nil || o.TriggeringTime == nil {
		return nil, false
	}
	return o.TriggeringTime, true
}

// HasTriggeringTime returns a boolean if a field has been set.
func (o *EventResult) HasTriggeringTime() bool {
	if o != nil && o.TriggeringTime != nil {
		return true
	}

	return false
}

// SetTriggeringTime gets a reference to the given int64 and assigns it to the TriggeringTime field.
func (o *EventResult) SetTriggeringTime(v int64) {
	o.TriggeringTime = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *EventResult) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventResult) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *EventResult) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *EventResult) SetType(v string) {
	o.Type = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *EventResult) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventResult) GetStateOk() (*string, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *EventResult) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *EventResult) SetState(v string) {
	o.State = &v
}

// GetProblem returns the Problem field value if set, zero value otherwise.
func (o *EventResult) GetProblem() string {
	if o == nil || o.Problem == nil {
		var ret string
		return ret
	}
	return *o.Problem
}

// GetProblemOk returns a tuple with the Problem field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventResult) GetProblemOk() (*string, bool) {
	if o == nil || o.Problem == nil {
		return nil, false
	}
	return o.Problem, true
}

// HasProblem returns a boolean if a field has been set.
func (o *EventResult) HasProblem() bool {
	if o != nil && o.Problem != nil {
		return true
	}

	return false
}

// SetProblem gets a reference to the given string and assigns it to the Problem field.
func (o *EventResult) SetProblem(v string) {
	o.Problem = &v
}

// GetFixSuggestion returns the FixSuggestion field value if set, zero value otherwise.
func (o *EventResult) GetFixSuggestion() string {
	if o == nil || o.FixSuggestion == nil {
		var ret string
		return ret
	}
	return *o.FixSuggestion
}

// GetFixSuggestionOk returns a tuple with the FixSuggestion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventResult) GetFixSuggestionOk() (*string, bool) {
	if o == nil || o.FixSuggestion == nil {
		return nil, false
	}
	return o.FixSuggestion, true
}

// HasFixSuggestion returns a boolean if a field has been set.
func (o *EventResult) HasFixSuggestion() bool {
	if o != nil && o.FixSuggestion != nil {
		return true
	}

	return false
}

// SetFixSuggestion gets a reference to the given string and assigns it to the FixSuggestion field.
func (o *EventResult) SetFixSuggestion(v string) {
	o.FixSuggestion = &v
}

// GetSeverity returns the Severity field value if set, zero value otherwise.
func (o *EventResult) GetSeverity() int32 {
	if o == nil || o.Severity == nil {
		var ret int32
		return ret
	}
	return *o.Severity
}

// GetSeverityOk returns a tuple with the Severity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventResult) GetSeverityOk() (*int32, bool) {
	if o == nil || o.Severity == nil {
		return nil, false
	}
	return o.Severity, true
}

// HasSeverity returns a boolean if a field has been set.
func (o *EventResult) HasSeverity() bool {
	if o != nil && o.Severity != nil {
		return true
	}

	return false
}

// SetSeverity gets a reference to the given int32 and assigns it to the Severity field.
func (o *EventResult) SetSeverity(v int32) {
	o.Severity = &v
}

// GetSnapshotId returns the SnapshotId field value if set, zero value otherwise.
func (o *EventResult) GetSnapshotId() string {
	if o == nil || o.SnapshotId == nil {
		var ret string
		return ret
	}
	return *o.SnapshotId
}

// GetSnapshotIdOk returns a tuple with the SnapshotId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventResult) GetSnapshotIdOk() (*string, bool) {
	if o == nil || o.SnapshotId == nil {
		return nil, false
	}
	return o.SnapshotId, true
}

// HasSnapshotId returns a boolean if a field has been set.
func (o *EventResult) HasSnapshotId() bool {
	if o != nil && o.SnapshotId != nil {
		return true
	}

	return false
}

// SetSnapshotId gets a reference to the given string and assigns it to the SnapshotId field.
func (o *EventResult) SetSnapshotId(v string) {
	o.SnapshotId = &v
}

// GetMetrics returns the Metrics field value if set, zero value otherwise.
func (o *EventResult) GetMetrics() []map[string]interface{} {
	if o == nil || o.Metrics == nil {
		var ret []map[string]interface{}
		return ret
	}
	return *o.Metrics
}

// GetMetricsOk returns a tuple with the Metrics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventResult) GetMetricsOk() (*[]map[string]interface{}, bool) {
	if o == nil || o.Metrics == nil {
		return nil, false
	}
	return o.Metrics, true
}

// HasMetrics returns a boolean if a field has been set.
func (o *EventResult) HasMetrics() bool {
	if o != nil && o.Metrics != nil {
		return true
	}

	return false
}

// SetMetrics gets a reference to the given []map[string]interface{} and assigns it to the Metrics field.
func (o *EventResult) SetMetrics(v []map[string]interface{}) {
	o.Metrics = &v
}

// GetRecentEvents returns the RecentEvents field value if set, zero value otherwise.
func (o *EventResult) GetRecentEvents() []map[string]interface{} {
	if o == nil || o.RecentEvents == nil {
		var ret []map[string]interface{}
		return ret
	}
	return *o.RecentEvents
}

// GetRecentEventsOk returns a tuple with the RecentEvents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventResult) GetRecentEventsOk() (*[]map[string]interface{}, bool) {
	if o == nil || o.RecentEvents == nil {
		return nil, false
	}
	return o.RecentEvents, true
}

// HasRecentEvents returns a boolean if a field has been set.
func (o *EventResult) HasRecentEvents() bool {
	if o != nil && o.RecentEvents != nil {
		return true
	}

	return false
}

// SetRecentEvents gets a reference to the given []map[string]interface{} and assigns it to the RecentEvents field.
func (o *EventResult) SetRecentEvents(v []map[string]interface{}) {
	o.RecentEvents = &v
}

func (o EventResult) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.EventId != nil {
		toSerialize["eventId"] = o.EventId
	}
	if o.Start != nil {
		toSerialize["start"] = o.Start
	}
	if o.End != nil {
		toSerialize["end"] = o.End
	}
	if o.TriggeringTime != nil {
		toSerialize["triggeringTime"] = o.TriggeringTime
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.State != nil {
		toSerialize["state"] = o.State
	}
	if o.Problem != nil {
		toSerialize["problem"] = o.Problem
	}
	if o.FixSuggestion != nil {
		toSerialize["fixSuggestion"] = o.FixSuggestion
	}
	if o.Severity != nil {
		toSerialize["severity"] = o.Severity
	}
	if o.SnapshotId != nil {
		toSerialize["snapshotId"] = o.SnapshotId
	}
	if o.Metrics != nil {
		toSerialize["metrics"] = o.Metrics
	}
	if o.RecentEvents != nil {
		toSerialize["recentEvents"] = o.RecentEvents
	}
	return json.Marshal(toSerialize)
}

type NullableEventResult struct {
	value *EventResult
	isSet bool
}

func (v NullableEventResult) Get() *EventResult {
	return v.value
}

func (v *NullableEventResult) Set(val *EventResult) {
	v.value = val
	v.isSet = true
}

func (v NullableEventResult) IsSet() bool {
	return v.isSet
}

func (v *NullableEventResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventResult(val *EventResult) *NullableEventResult {
	return &NullableEventResult{value: val, isSet: true}
}

func (v NullableEventResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
