/*
Strava API v3

The [Swagger Playground](https://developers.strava.com/playground) is the easiest way to familiarize yourself with the Strava API by submitting HTTP requests and observing the responses before you write any client code. It will show what a response will look like with different endpoints depending on the authorization scope you receive from your athletes. To use the Playground, go to https://www.strava.com/settings/api and change your “Authorization Callback Domain” to developers.strava.com. Please note, we only support Swagger 2.0. There is a known issue where you can only select one scope at a time. For more information, please check the section “client code” at https://developers.strava.com/docs.

API version: 3.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package strava3golang

import (
	"encoding/json"
	"time"
)

// checks if the SummaryPRSegmentEffort type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SummaryPRSegmentEffort{}

// SummaryPRSegmentEffort struct for SummaryPRSegmentEffort
type SummaryPRSegmentEffort struct {
	// The unique identifier of the activity related to the PR effort.
	PrActivityId *int64 `json:"pr_activity_id,omitempty"`
	// The elapsed time ot the PR effort.
	PrElapsedTime *int32 `json:"pr_elapsed_time,omitempty"`
	// The time at which the PR effort was started.
	PrDate *time.Time `json:"pr_date,omitempty"`
	// Number of efforts by the authenticated athlete on this segment.
	EffortCount *int32 `json:"effort_count,omitempty"`
}

// NewSummaryPRSegmentEffort instantiates a new SummaryPRSegmentEffort object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSummaryPRSegmentEffort() *SummaryPRSegmentEffort {
	this := SummaryPRSegmentEffort{}
	return &this
}

// NewSummaryPRSegmentEffortWithDefaults instantiates a new SummaryPRSegmentEffort object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSummaryPRSegmentEffortWithDefaults() *SummaryPRSegmentEffort {
	this := SummaryPRSegmentEffort{}
	return &this
}

// GetPrActivityId returns the PrActivityId field value if set, zero value otherwise.
func (o *SummaryPRSegmentEffort) GetPrActivityId() int64 {
	if o == nil || IsNil(o.PrActivityId) {
		var ret int64
		return ret
	}
	return *o.PrActivityId
}

// GetPrActivityIdOk returns a tuple with the PrActivityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SummaryPRSegmentEffort) GetPrActivityIdOk() (*int64, bool) {
	if o == nil || IsNil(o.PrActivityId) {
		return nil, false
	}
	return o.PrActivityId, true
}

// HasPrActivityId returns a boolean if a field has been set.
func (o *SummaryPRSegmentEffort) HasPrActivityId() bool {
	if o != nil && !IsNil(o.PrActivityId) {
		return true
	}

	return false
}

// SetPrActivityId gets a reference to the given int64 and assigns it to the PrActivityId field.
func (o *SummaryPRSegmentEffort) SetPrActivityId(v int64) {
	o.PrActivityId = &v
}

// GetPrElapsedTime returns the PrElapsedTime field value if set, zero value otherwise.
func (o *SummaryPRSegmentEffort) GetPrElapsedTime() int32 {
	if o == nil || IsNil(o.PrElapsedTime) {
		var ret int32
		return ret
	}
	return *o.PrElapsedTime
}

// GetPrElapsedTimeOk returns a tuple with the PrElapsedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SummaryPRSegmentEffort) GetPrElapsedTimeOk() (*int32, bool) {
	if o == nil || IsNil(o.PrElapsedTime) {
		return nil, false
	}
	return o.PrElapsedTime, true
}

// HasPrElapsedTime returns a boolean if a field has been set.
func (o *SummaryPRSegmentEffort) HasPrElapsedTime() bool {
	if o != nil && !IsNil(o.PrElapsedTime) {
		return true
	}

	return false
}

// SetPrElapsedTime gets a reference to the given int32 and assigns it to the PrElapsedTime field.
func (o *SummaryPRSegmentEffort) SetPrElapsedTime(v int32) {
	o.PrElapsedTime = &v
}

// GetPrDate returns the PrDate field value if set, zero value otherwise.
func (o *SummaryPRSegmentEffort) GetPrDate() time.Time {
	if o == nil || IsNil(o.PrDate) {
		var ret time.Time
		return ret
	}
	return *o.PrDate
}

// GetPrDateOk returns a tuple with the PrDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SummaryPRSegmentEffort) GetPrDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.PrDate) {
		return nil, false
	}
	return o.PrDate, true
}

// HasPrDate returns a boolean if a field has been set.
func (o *SummaryPRSegmentEffort) HasPrDate() bool {
	if o != nil && !IsNil(o.PrDate) {
		return true
	}

	return false
}

// SetPrDate gets a reference to the given time.Time and assigns it to the PrDate field.
func (o *SummaryPRSegmentEffort) SetPrDate(v time.Time) {
	o.PrDate = &v
}

// GetEffortCount returns the EffortCount field value if set, zero value otherwise.
func (o *SummaryPRSegmentEffort) GetEffortCount() int32 {
	if o == nil || IsNil(o.EffortCount) {
		var ret int32
		return ret
	}
	return *o.EffortCount
}

// GetEffortCountOk returns a tuple with the EffortCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SummaryPRSegmentEffort) GetEffortCountOk() (*int32, bool) {
	if o == nil || IsNil(o.EffortCount) {
		return nil, false
	}
	return o.EffortCount, true
}

// HasEffortCount returns a boolean if a field has been set.
func (o *SummaryPRSegmentEffort) HasEffortCount() bool {
	if o != nil && !IsNil(o.EffortCount) {
		return true
	}

	return false
}

// SetEffortCount gets a reference to the given int32 and assigns it to the EffortCount field.
func (o *SummaryPRSegmentEffort) SetEffortCount(v int32) {
	o.EffortCount = &v
}

func (o SummaryPRSegmentEffort) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SummaryPRSegmentEffort) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PrActivityId) {
		toSerialize["pr_activity_id"] = o.PrActivityId
	}
	if !IsNil(o.PrElapsedTime) {
		toSerialize["pr_elapsed_time"] = o.PrElapsedTime
	}
	if !IsNil(o.PrDate) {
		toSerialize["pr_date"] = o.PrDate
	}
	if !IsNil(o.EffortCount) {
		toSerialize["effort_count"] = o.EffortCount
	}
	return toSerialize, nil
}

type NullableSummaryPRSegmentEffort struct {
	value *SummaryPRSegmentEffort
	isSet bool
}

func (v NullableSummaryPRSegmentEffort) Get() *SummaryPRSegmentEffort {
	return v.value
}

func (v *NullableSummaryPRSegmentEffort) Set(val *SummaryPRSegmentEffort) {
	v.value = val
	v.isSet = true
}

func (v NullableSummaryPRSegmentEffort) IsSet() bool {
	return v.isSet
}

func (v *NullableSummaryPRSegmentEffort) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSummaryPRSegmentEffort(val *SummaryPRSegmentEffort) *NullableSummaryPRSegmentEffort {
	return &NullableSummaryPRSegmentEffort{value: val, isSet: true}
}

func (v NullableSummaryPRSegmentEffort) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSummaryPRSegmentEffort) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


