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

// checks if the Route type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Route{}

// Route struct for Route
type Route struct {
	Athlete *SummaryAthlete `json:"athlete,omitempty"`
	// The description of the route
	Description *string `json:"description,omitempty"`
	// The route's distance, in meters
	Distance *float32 `json:"distance,omitempty"`
	// The route's elevation gain.
	ElevationGain *float32 `json:"elevation_gain,omitempty"`
	// The unique identifier of this route
	Id *int64 `json:"id,omitempty"`
	// The unique identifier of the route in string format
	IdStr *string `json:"id_str,omitempty"`
	Map *PolylineMap `json:"map,omitempty"`
	// The name of this route
	Name *string `json:"name,omitempty"`
	// Whether this route is private
	Private *bool `json:"private,omitempty"`
	// Whether this route is starred by the logged-in athlete
	Starred *bool `json:"starred,omitempty"`
	// An epoch timestamp of when the route was created
	Timestamp *int32 `json:"timestamp,omitempty"`
	// This route's type (1 for ride, 2 for runs)
	Type *int32 `json:"type,omitempty"`
	// This route's sub-type (1 for road, 2 for mountain bike, 3 for cross, 4 for trail, 5 for mixed)
	SubType *int32 `json:"sub_type,omitempty"`
	// The time at which the route was created
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// The time at which the route was last updated
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	// Estimated time in seconds for the authenticated athlete to complete route
	EstimatedMovingTime *int32 `json:"estimated_moving_time,omitempty"`
	// The segments traversed by this route
	Segments []SummarySegment `json:"segments,omitempty"`
}

// NewRoute instantiates a new Route object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRoute() *Route {
	this := Route{}
	return &this
}

// NewRouteWithDefaults instantiates a new Route object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRouteWithDefaults() *Route {
	this := Route{}
	return &this
}

// GetAthlete returns the Athlete field value if set, zero value otherwise.
func (o *Route) GetAthlete() SummaryAthlete {
	if o == nil || IsNil(o.Athlete) {
		var ret SummaryAthlete
		return ret
	}
	return *o.Athlete
}

// GetAthleteOk returns a tuple with the Athlete field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Route) GetAthleteOk() (*SummaryAthlete, bool) {
	if o == nil || IsNil(o.Athlete) {
		return nil, false
	}
	return o.Athlete, true
}

// HasAthlete returns a boolean if a field has been set.
func (o *Route) HasAthlete() bool {
	if o != nil && !IsNil(o.Athlete) {
		return true
	}

	return false
}

// SetAthlete gets a reference to the given SummaryAthlete and assigns it to the Athlete field.
func (o *Route) SetAthlete(v SummaryAthlete) {
	o.Athlete = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Route) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Route) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Route) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Route) SetDescription(v string) {
	o.Description = &v
}

// GetDistance returns the Distance field value if set, zero value otherwise.
func (o *Route) GetDistance() float32 {
	if o == nil || IsNil(o.Distance) {
		var ret float32
		return ret
	}
	return *o.Distance
}

// GetDistanceOk returns a tuple with the Distance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Route) GetDistanceOk() (*float32, bool) {
	if o == nil || IsNil(o.Distance) {
		return nil, false
	}
	return o.Distance, true
}

// HasDistance returns a boolean if a field has been set.
func (o *Route) HasDistance() bool {
	if o != nil && !IsNil(o.Distance) {
		return true
	}

	return false
}

// SetDistance gets a reference to the given float32 and assigns it to the Distance field.
func (o *Route) SetDistance(v float32) {
	o.Distance = &v
}

// GetElevationGain returns the ElevationGain field value if set, zero value otherwise.
func (o *Route) GetElevationGain() float32 {
	if o == nil || IsNil(o.ElevationGain) {
		var ret float32
		return ret
	}
	return *o.ElevationGain
}

// GetElevationGainOk returns a tuple with the ElevationGain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Route) GetElevationGainOk() (*float32, bool) {
	if o == nil || IsNil(o.ElevationGain) {
		return nil, false
	}
	return o.ElevationGain, true
}

// HasElevationGain returns a boolean if a field has been set.
func (o *Route) HasElevationGain() bool {
	if o != nil && !IsNil(o.ElevationGain) {
		return true
	}

	return false
}

// SetElevationGain gets a reference to the given float32 and assigns it to the ElevationGain field.
func (o *Route) SetElevationGain(v float32) {
	o.ElevationGain = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Route) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Route) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Route) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *Route) SetId(v int64) {
	o.Id = &v
}

// GetIdStr returns the IdStr field value if set, zero value otherwise.
func (o *Route) GetIdStr() string {
	if o == nil || IsNil(o.IdStr) {
		var ret string
		return ret
	}
	return *o.IdStr
}

// GetIdStrOk returns a tuple with the IdStr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Route) GetIdStrOk() (*string, bool) {
	if o == nil || IsNil(o.IdStr) {
		return nil, false
	}
	return o.IdStr, true
}

// HasIdStr returns a boolean if a field has been set.
func (o *Route) HasIdStr() bool {
	if o != nil && !IsNil(o.IdStr) {
		return true
	}

	return false
}

// SetIdStr gets a reference to the given string and assigns it to the IdStr field.
func (o *Route) SetIdStr(v string) {
	o.IdStr = &v
}

// GetMap returns the Map field value if set, zero value otherwise.
func (o *Route) GetMap() PolylineMap {
	if o == nil || IsNil(o.Map) {
		var ret PolylineMap
		return ret
	}
	return *o.Map
}

// GetMapOk returns a tuple with the Map field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Route) GetMapOk() (*PolylineMap, bool) {
	if o == nil || IsNil(o.Map) {
		return nil, false
	}
	return o.Map, true
}

// HasMap returns a boolean if a field has been set.
func (o *Route) HasMap() bool {
	if o != nil && !IsNil(o.Map) {
		return true
	}

	return false
}

// SetMap gets a reference to the given PolylineMap and assigns it to the Map field.
func (o *Route) SetMap(v PolylineMap) {
	o.Map = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Route) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Route) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Route) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Route) SetName(v string) {
	o.Name = &v
}

// GetPrivate returns the Private field value if set, zero value otherwise.
func (o *Route) GetPrivate() bool {
	if o == nil || IsNil(o.Private) {
		var ret bool
		return ret
	}
	return *o.Private
}

// GetPrivateOk returns a tuple with the Private field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Route) GetPrivateOk() (*bool, bool) {
	if o == nil || IsNil(o.Private) {
		return nil, false
	}
	return o.Private, true
}

// HasPrivate returns a boolean if a field has been set.
func (o *Route) HasPrivate() bool {
	if o != nil && !IsNil(o.Private) {
		return true
	}

	return false
}

// SetPrivate gets a reference to the given bool and assigns it to the Private field.
func (o *Route) SetPrivate(v bool) {
	o.Private = &v
}

// GetStarred returns the Starred field value if set, zero value otherwise.
func (o *Route) GetStarred() bool {
	if o == nil || IsNil(o.Starred) {
		var ret bool
		return ret
	}
	return *o.Starred
}

// GetStarredOk returns a tuple with the Starred field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Route) GetStarredOk() (*bool, bool) {
	if o == nil || IsNil(o.Starred) {
		return nil, false
	}
	return o.Starred, true
}

// HasStarred returns a boolean if a field has been set.
func (o *Route) HasStarred() bool {
	if o != nil && !IsNil(o.Starred) {
		return true
	}

	return false
}

// SetStarred gets a reference to the given bool and assigns it to the Starred field.
func (o *Route) SetStarred(v bool) {
	o.Starred = &v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *Route) GetTimestamp() int32 {
	if o == nil || IsNil(o.Timestamp) {
		var ret int32
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Route) GetTimestampOk() (*int32, bool) {
	if o == nil || IsNil(o.Timestamp) {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *Route) HasTimestamp() bool {
	if o != nil && !IsNil(o.Timestamp) {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given int32 and assigns it to the Timestamp field.
func (o *Route) SetTimestamp(v int32) {
	o.Timestamp = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *Route) GetType() int32 {
	if o == nil || IsNil(o.Type) {
		var ret int32
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Route) GetTypeOk() (*int32, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *Route) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given int32 and assigns it to the Type field.
func (o *Route) SetType(v int32) {
	o.Type = &v
}

// GetSubType returns the SubType field value if set, zero value otherwise.
func (o *Route) GetSubType() int32 {
	if o == nil || IsNil(o.SubType) {
		var ret int32
		return ret
	}
	return *o.SubType
}

// GetSubTypeOk returns a tuple with the SubType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Route) GetSubTypeOk() (*int32, bool) {
	if o == nil || IsNil(o.SubType) {
		return nil, false
	}
	return o.SubType, true
}

// HasSubType returns a boolean if a field has been set.
func (o *Route) HasSubType() bool {
	if o != nil && !IsNil(o.SubType) {
		return true
	}

	return false
}

// SetSubType gets a reference to the given int32 and assigns it to the SubType field.
func (o *Route) SetSubType(v int32) {
	o.SubType = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *Route) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Route) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *Route) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *Route) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *Route) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Route) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *Route) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *Route) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetEstimatedMovingTime returns the EstimatedMovingTime field value if set, zero value otherwise.
func (o *Route) GetEstimatedMovingTime() int32 {
	if o == nil || IsNil(o.EstimatedMovingTime) {
		var ret int32
		return ret
	}
	return *o.EstimatedMovingTime
}

// GetEstimatedMovingTimeOk returns a tuple with the EstimatedMovingTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Route) GetEstimatedMovingTimeOk() (*int32, bool) {
	if o == nil || IsNil(o.EstimatedMovingTime) {
		return nil, false
	}
	return o.EstimatedMovingTime, true
}

// HasEstimatedMovingTime returns a boolean if a field has been set.
func (o *Route) HasEstimatedMovingTime() bool {
	if o != nil && !IsNil(o.EstimatedMovingTime) {
		return true
	}

	return false
}

// SetEstimatedMovingTime gets a reference to the given int32 and assigns it to the EstimatedMovingTime field.
func (o *Route) SetEstimatedMovingTime(v int32) {
	o.EstimatedMovingTime = &v
}

// GetSegments returns the Segments field value if set, zero value otherwise.
func (o *Route) GetSegments() []SummarySegment {
	if o == nil || IsNil(o.Segments) {
		var ret []SummarySegment
		return ret
	}
	return o.Segments
}

// GetSegmentsOk returns a tuple with the Segments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Route) GetSegmentsOk() ([]SummarySegment, bool) {
	if o == nil || IsNil(o.Segments) {
		return nil, false
	}
	return o.Segments, true
}

// HasSegments returns a boolean if a field has been set.
func (o *Route) HasSegments() bool {
	if o != nil && !IsNil(o.Segments) {
		return true
	}

	return false
}

// SetSegments gets a reference to the given []SummarySegment and assigns it to the Segments field.
func (o *Route) SetSegments(v []SummarySegment) {
	o.Segments = v
}

func (o Route) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Route) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Athlete) {
		toSerialize["athlete"] = o.Athlete
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Distance) {
		toSerialize["distance"] = o.Distance
	}
	if !IsNil(o.ElevationGain) {
		toSerialize["elevation_gain"] = o.ElevationGain
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.IdStr) {
		toSerialize["id_str"] = o.IdStr
	}
	if !IsNil(o.Map) {
		toSerialize["map"] = o.Map
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Private) {
		toSerialize["private"] = o.Private
	}
	if !IsNil(o.Starred) {
		toSerialize["starred"] = o.Starred
	}
	if !IsNil(o.Timestamp) {
		toSerialize["timestamp"] = o.Timestamp
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.SubType) {
		toSerialize["sub_type"] = o.SubType
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if !IsNil(o.EstimatedMovingTime) {
		toSerialize["estimated_moving_time"] = o.EstimatedMovingTime
	}
	if !IsNil(o.Segments) {
		toSerialize["segments"] = o.Segments
	}
	return toSerialize, nil
}

type NullableRoute struct {
	value *Route
	isSet bool
}

func (v NullableRoute) Get() *Route {
	return v.value
}

func (v *NullableRoute) Set(val *Route) {
	v.value = val
	v.isSet = true
}

func (v NullableRoute) IsSet() bool {
	return v.isSet
}

func (v *NullableRoute) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoute(val *Route) *NullableRoute {
	return &NullableRoute{value: val, isSet: true}
}

func (v NullableRoute) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoute) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


