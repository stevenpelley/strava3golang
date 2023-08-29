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

// checks if the SummarySegmentEffort type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SummarySegmentEffort{}

// SummarySegmentEffort struct for SummarySegmentEffort
type SummarySegmentEffort struct {
	// The unique identifier of this effort
	Id *int64 `json:"id,omitempty"`
	// The unique identifier of the activity related to this effort
	ActivityId *int64 `json:"activity_id,omitempty"`
	// The effort's elapsed time
	ElapsedTime *int32 `json:"elapsed_time,omitempty"`
	// The time at which the effort was started.
	StartDate *time.Time `json:"start_date,omitempty"`
	// The time at which the effort was started in the local timezone.
	StartDateLocal *time.Time `json:"start_date_local,omitempty"`
	// The effort's distance in meters
	Distance *float32 `json:"distance,omitempty"`
	// Whether this effort is the current best on the leaderboard
	IsKom *bool `json:"is_kom,omitempty"`
}

// NewSummarySegmentEffort instantiates a new SummarySegmentEffort object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSummarySegmentEffort() *SummarySegmentEffort {
	this := SummarySegmentEffort{}
	return &this
}

// NewSummarySegmentEffortWithDefaults instantiates a new SummarySegmentEffort object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSummarySegmentEffortWithDefaults() *SummarySegmentEffort {
	this := SummarySegmentEffort{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SummarySegmentEffort) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SummarySegmentEffort) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SummarySegmentEffort) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *SummarySegmentEffort) SetId(v int64) {
	o.Id = &v
}

// GetActivityId returns the ActivityId field value if set, zero value otherwise.
func (o *SummarySegmentEffort) GetActivityId() int64 {
	if o == nil || IsNil(o.ActivityId) {
		var ret int64
		return ret
	}
	return *o.ActivityId
}

// GetActivityIdOk returns a tuple with the ActivityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SummarySegmentEffort) GetActivityIdOk() (*int64, bool) {
	if o == nil || IsNil(o.ActivityId) {
		return nil, false
	}
	return o.ActivityId, true
}

// HasActivityId returns a boolean if a field has been set.
func (o *SummarySegmentEffort) HasActivityId() bool {
	if o != nil && !IsNil(o.ActivityId) {
		return true
	}

	return false
}

// SetActivityId gets a reference to the given int64 and assigns it to the ActivityId field.
func (o *SummarySegmentEffort) SetActivityId(v int64) {
	o.ActivityId = &v
}

// GetElapsedTime returns the ElapsedTime field value if set, zero value otherwise.
func (o *SummarySegmentEffort) GetElapsedTime() int32 {
	if o == nil || IsNil(o.ElapsedTime) {
		var ret int32
		return ret
	}
	return *o.ElapsedTime
}

// GetElapsedTimeOk returns a tuple with the ElapsedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SummarySegmentEffort) GetElapsedTimeOk() (*int32, bool) {
	if o == nil || IsNil(o.ElapsedTime) {
		return nil, false
	}
	return o.ElapsedTime, true
}

// HasElapsedTime returns a boolean if a field has been set.
func (o *SummarySegmentEffort) HasElapsedTime() bool {
	if o != nil && !IsNil(o.ElapsedTime) {
		return true
	}

	return false
}

// SetElapsedTime gets a reference to the given int32 and assigns it to the ElapsedTime field.
func (o *SummarySegmentEffort) SetElapsedTime(v int32) {
	o.ElapsedTime = &v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *SummarySegmentEffort) GetStartDate() time.Time {
	if o == nil || IsNil(o.StartDate) {
		var ret time.Time
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SummarySegmentEffort) GetStartDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.StartDate) {
		return nil, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *SummarySegmentEffort) HasStartDate() bool {
	if o != nil && !IsNil(o.StartDate) {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given time.Time and assigns it to the StartDate field.
func (o *SummarySegmentEffort) SetStartDate(v time.Time) {
	o.StartDate = &v
}

// GetStartDateLocal returns the StartDateLocal field value if set, zero value otherwise.
func (o *SummarySegmentEffort) GetStartDateLocal() time.Time {
	if o == nil || IsNil(o.StartDateLocal) {
		var ret time.Time
		return ret
	}
	return *o.StartDateLocal
}

// GetStartDateLocalOk returns a tuple with the StartDateLocal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SummarySegmentEffort) GetStartDateLocalOk() (*time.Time, bool) {
	if o == nil || IsNil(o.StartDateLocal) {
		return nil, false
	}
	return o.StartDateLocal, true
}

// HasStartDateLocal returns a boolean if a field has been set.
func (o *SummarySegmentEffort) HasStartDateLocal() bool {
	if o != nil && !IsNil(o.StartDateLocal) {
		return true
	}

	return false
}

// SetStartDateLocal gets a reference to the given time.Time and assigns it to the StartDateLocal field.
func (o *SummarySegmentEffort) SetStartDateLocal(v time.Time) {
	o.StartDateLocal = &v
}

// GetDistance returns the Distance field value if set, zero value otherwise.
func (o *SummarySegmentEffort) GetDistance() float32 {
	if o == nil || IsNil(o.Distance) {
		var ret float32
		return ret
	}
	return *o.Distance
}

// GetDistanceOk returns a tuple with the Distance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SummarySegmentEffort) GetDistanceOk() (*float32, bool) {
	if o == nil || IsNil(o.Distance) {
		return nil, false
	}
	return o.Distance, true
}

// HasDistance returns a boolean if a field has been set.
func (o *SummarySegmentEffort) HasDistance() bool {
	if o != nil && !IsNil(o.Distance) {
		return true
	}

	return false
}

// SetDistance gets a reference to the given float32 and assigns it to the Distance field.
func (o *SummarySegmentEffort) SetDistance(v float32) {
	o.Distance = &v
}

// GetIsKom returns the IsKom field value if set, zero value otherwise.
func (o *SummarySegmentEffort) GetIsKom() bool {
	if o == nil || IsNil(o.IsKom) {
		var ret bool
		return ret
	}
	return *o.IsKom
}

// GetIsKomOk returns a tuple with the IsKom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SummarySegmentEffort) GetIsKomOk() (*bool, bool) {
	if o == nil || IsNil(o.IsKom) {
		return nil, false
	}
	return o.IsKom, true
}

// HasIsKom returns a boolean if a field has been set.
func (o *SummarySegmentEffort) HasIsKom() bool {
	if o != nil && !IsNil(o.IsKom) {
		return true
	}

	return false
}

// SetIsKom gets a reference to the given bool and assigns it to the IsKom field.
func (o *SummarySegmentEffort) SetIsKom(v bool) {
	o.IsKom = &v
}

func (o SummarySegmentEffort) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SummarySegmentEffort) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.ActivityId) {
		toSerialize["activity_id"] = o.ActivityId
	}
	if !IsNil(o.ElapsedTime) {
		toSerialize["elapsed_time"] = o.ElapsedTime
	}
	if !IsNil(o.StartDate) {
		toSerialize["start_date"] = o.StartDate
	}
	if !IsNil(o.StartDateLocal) {
		toSerialize["start_date_local"] = o.StartDateLocal
	}
	if !IsNil(o.Distance) {
		toSerialize["distance"] = o.Distance
	}
	if !IsNil(o.IsKom) {
		toSerialize["is_kom"] = o.IsKom
	}
	return toSerialize, nil
}

type NullableSummarySegmentEffort struct {
	value *SummarySegmentEffort
	isSet bool
}

func (v NullableSummarySegmentEffort) Get() *SummarySegmentEffort {
	return v.value
}

func (v *NullableSummarySegmentEffort) Set(val *SummarySegmentEffort) {
	v.value = val
	v.isSet = true
}

func (v NullableSummarySegmentEffort) IsSet() bool {
	return v.isSet
}

func (v *NullableSummarySegmentEffort) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSummarySegmentEffort(val *SummarySegmentEffort) *NullableSummarySegmentEffort {
	return &NullableSummarySegmentEffort{value: val, isSet: true}
}

func (v NullableSummarySegmentEffort) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSummarySegmentEffort) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

