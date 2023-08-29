/*
Strava API v3

The [Swagger Playground](https://developers.strava.com/playground) is the easiest way to familiarize yourself with the Strava API by submitting HTTP requests and observing the responses before you write any client code. It will show what a response will look like with different endpoints depending on the authorization scope you receive from your athletes. To use the Playground, go to https://www.strava.com/settings/api and change your “Authorization Callback Domain” to developers.strava.com. Please note, we only support Swagger 2.0. There is a known issue where you can only select one scope at a time. For more information, please check the section “client code” at https://developers.strava.com/docs.

API version: 3.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package strava3golang

import (
	"encoding/json"
)

// checks if the ExplorerSegment type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExplorerSegment{}

// ExplorerSegment struct for ExplorerSegment
type ExplorerSegment struct {
	// The unique identifier of this segment
	Id *int64 `json:"id,omitempty"`
	// The name of this segment
	Name *string `json:"name,omitempty"`
	// The category of the climb [0, 5]. Higher is harder ie. 5 is Hors catégorie, 0 is uncategorized in climb_category. If climb_category = 5, climb_category_desc = HC. If climb_category = 2, climb_category_desc = 3.
	ClimbCategory *int32 `json:"climb_category,omitempty"`
	// The description for the category of the climb
	ClimbCategoryDesc *string `json:"climb_category_desc,omitempty"`
	// The segment's average grade, in percents
	AvgGrade *float32 `json:"avg_grade,omitempty"`
	// A pair of latitude/longitude coordinates, represented as an array of 2 floating point numbers.
	StartLatlng []float32 `json:"start_latlng,omitempty"`
	// A pair of latitude/longitude coordinates, represented as an array of 2 floating point numbers.
	EndLatlng []float32 `json:"end_latlng,omitempty"`
	// The segments's evelation difference, in meters
	ElevDifference *float32 `json:"elev_difference,omitempty"`
	// The segment's distance, in meters
	Distance *float32 `json:"distance,omitempty"`
	// The polyline of the segment
	Points *string `json:"points,omitempty"`
}

// NewExplorerSegment instantiates a new ExplorerSegment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExplorerSegment() *ExplorerSegment {
	this := ExplorerSegment{}
	return &this
}

// NewExplorerSegmentWithDefaults instantiates a new ExplorerSegment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExplorerSegmentWithDefaults() *ExplorerSegment {
	this := ExplorerSegment{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ExplorerSegment) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExplorerSegment) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ExplorerSegment) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *ExplorerSegment) SetId(v int64) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ExplorerSegment) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExplorerSegment) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ExplorerSegment) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ExplorerSegment) SetName(v string) {
	o.Name = &v
}

// GetClimbCategory returns the ClimbCategory field value if set, zero value otherwise.
func (o *ExplorerSegment) GetClimbCategory() int32 {
	if o == nil || IsNil(o.ClimbCategory) {
		var ret int32
		return ret
	}
	return *o.ClimbCategory
}

// GetClimbCategoryOk returns a tuple with the ClimbCategory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExplorerSegment) GetClimbCategoryOk() (*int32, bool) {
	if o == nil || IsNil(o.ClimbCategory) {
		return nil, false
	}
	return o.ClimbCategory, true
}

// HasClimbCategory returns a boolean if a field has been set.
func (o *ExplorerSegment) HasClimbCategory() bool {
	if o != nil && !IsNil(o.ClimbCategory) {
		return true
	}

	return false
}

// SetClimbCategory gets a reference to the given int32 and assigns it to the ClimbCategory field.
func (o *ExplorerSegment) SetClimbCategory(v int32) {
	o.ClimbCategory = &v
}

// GetClimbCategoryDesc returns the ClimbCategoryDesc field value if set, zero value otherwise.
func (o *ExplorerSegment) GetClimbCategoryDesc() string {
	if o == nil || IsNil(o.ClimbCategoryDesc) {
		var ret string
		return ret
	}
	return *o.ClimbCategoryDesc
}

// GetClimbCategoryDescOk returns a tuple with the ClimbCategoryDesc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExplorerSegment) GetClimbCategoryDescOk() (*string, bool) {
	if o == nil || IsNil(o.ClimbCategoryDesc) {
		return nil, false
	}
	return o.ClimbCategoryDesc, true
}

// HasClimbCategoryDesc returns a boolean if a field has been set.
func (o *ExplorerSegment) HasClimbCategoryDesc() bool {
	if o != nil && !IsNil(o.ClimbCategoryDesc) {
		return true
	}

	return false
}

// SetClimbCategoryDesc gets a reference to the given string and assigns it to the ClimbCategoryDesc field.
func (o *ExplorerSegment) SetClimbCategoryDesc(v string) {
	o.ClimbCategoryDesc = &v
}

// GetAvgGrade returns the AvgGrade field value if set, zero value otherwise.
func (o *ExplorerSegment) GetAvgGrade() float32 {
	if o == nil || IsNil(o.AvgGrade) {
		var ret float32
		return ret
	}
	return *o.AvgGrade
}

// GetAvgGradeOk returns a tuple with the AvgGrade field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExplorerSegment) GetAvgGradeOk() (*float32, bool) {
	if o == nil || IsNil(o.AvgGrade) {
		return nil, false
	}
	return o.AvgGrade, true
}

// HasAvgGrade returns a boolean if a field has been set.
func (o *ExplorerSegment) HasAvgGrade() bool {
	if o != nil && !IsNil(o.AvgGrade) {
		return true
	}

	return false
}

// SetAvgGrade gets a reference to the given float32 and assigns it to the AvgGrade field.
func (o *ExplorerSegment) SetAvgGrade(v float32) {
	o.AvgGrade = &v
}

// GetStartLatlng returns the StartLatlng field value if set, zero value otherwise.
func (o *ExplorerSegment) GetStartLatlng() []float32 {
	if o == nil || IsNil(o.StartLatlng) {
		var ret []float32
		return ret
	}
	return o.StartLatlng
}

// GetStartLatlngOk returns a tuple with the StartLatlng field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExplorerSegment) GetStartLatlngOk() ([]float32, bool) {
	if o == nil || IsNil(o.StartLatlng) {
		return nil, false
	}
	return o.StartLatlng, true
}

// HasStartLatlng returns a boolean if a field has been set.
func (o *ExplorerSegment) HasStartLatlng() bool {
	if o != nil && !IsNil(o.StartLatlng) {
		return true
	}

	return false
}

// SetStartLatlng gets a reference to the given []float32 and assigns it to the StartLatlng field.
func (o *ExplorerSegment) SetStartLatlng(v []float32) {
	o.StartLatlng = v
}

// GetEndLatlng returns the EndLatlng field value if set, zero value otherwise.
func (o *ExplorerSegment) GetEndLatlng() []float32 {
	if o == nil || IsNil(o.EndLatlng) {
		var ret []float32
		return ret
	}
	return o.EndLatlng
}

// GetEndLatlngOk returns a tuple with the EndLatlng field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExplorerSegment) GetEndLatlngOk() ([]float32, bool) {
	if o == nil || IsNil(o.EndLatlng) {
		return nil, false
	}
	return o.EndLatlng, true
}

// HasEndLatlng returns a boolean if a field has been set.
func (o *ExplorerSegment) HasEndLatlng() bool {
	if o != nil && !IsNil(o.EndLatlng) {
		return true
	}

	return false
}

// SetEndLatlng gets a reference to the given []float32 and assigns it to the EndLatlng field.
func (o *ExplorerSegment) SetEndLatlng(v []float32) {
	o.EndLatlng = v
}

// GetElevDifference returns the ElevDifference field value if set, zero value otherwise.
func (o *ExplorerSegment) GetElevDifference() float32 {
	if o == nil || IsNil(o.ElevDifference) {
		var ret float32
		return ret
	}
	return *o.ElevDifference
}

// GetElevDifferenceOk returns a tuple with the ElevDifference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExplorerSegment) GetElevDifferenceOk() (*float32, bool) {
	if o == nil || IsNil(o.ElevDifference) {
		return nil, false
	}
	return o.ElevDifference, true
}

// HasElevDifference returns a boolean if a field has been set.
func (o *ExplorerSegment) HasElevDifference() bool {
	if o != nil && !IsNil(o.ElevDifference) {
		return true
	}

	return false
}

// SetElevDifference gets a reference to the given float32 and assigns it to the ElevDifference field.
func (o *ExplorerSegment) SetElevDifference(v float32) {
	o.ElevDifference = &v
}

// GetDistance returns the Distance field value if set, zero value otherwise.
func (o *ExplorerSegment) GetDistance() float32 {
	if o == nil || IsNil(o.Distance) {
		var ret float32
		return ret
	}
	return *o.Distance
}

// GetDistanceOk returns a tuple with the Distance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExplorerSegment) GetDistanceOk() (*float32, bool) {
	if o == nil || IsNil(o.Distance) {
		return nil, false
	}
	return o.Distance, true
}

// HasDistance returns a boolean if a field has been set.
func (o *ExplorerSegment) HasDistance() bool {
	if o != nil && !IsNil(o.Distance) {
		return true
	}

	return false
}

// SetDistance gets a reference to the given float32 and assigns it to the Distance field.
func (o *ExplorerSegment) SetDistance(v float32) {
	o.Distance = &v
}

// GetPoints returns the Points field value if set, zero value otherwise.
func (o *ExplorerSegment) GetPoints() string {
	if o == nil || IsNil(o.Points) {
		var ret string
		return ret
	}
	return *o.Points
}

// GetPointsOk returns a tuple with the Points field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExplorerSegment) GetPointsOk() (*string, bool) {
	if o == nil || IsNil(o.Points) {
		return nil, false
	}
	return o.Points, true
}

// HasPoints returns a boolean if a field has been set.
func (o *ExplorerSegment) HasPoints() bool {
	if o != nil && !IsNil(o.Points) {
		return true
	}

	return false
}

// SetPoints gets a reference to the given string and assigns it to the Points field.
func (o *ExplorerSegment) SetPoints(v string) {
	o.Points = &v
}

func (o ExplorerSegment) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExplorerSegment) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.ClimbCategory) {
		toSerialize["climb_category"] = o.ClimbCategory
	}
	if !IsNil(o.ClimbCategoryDesc) {
		toSerialize["climb_category_desc"] = o.ClimbCategoryDesc
	}
	if !IsNil(o.AvgGrade) {
		toSerialize["avg_grade"] = o.AvgGrade
	}
	if !IsNil(o.StartLatlng) {
		toSerialize["start_latlng"] = o.StartLatlng
	}
	if !IsNil(o.EndLatlng) {
		toSerialize["end_latlng"] = o.EndLatlng
	}
	if !IsNil(o.ElevDifference) {
		toSerialize["elev_difference"] = o.ElevDifference
	}
	if !IsNil(o.Distance) {
		toSerialize["distance"] = o.Distance
	}
	if !IsNil(o.Points) {
		toSerialize["points"] = o.Points
	}
	return toSerialize, nil
}

type NullableExplorerSegment struct {
	value *ExplorerSegment
	isSet bool
}

func (v NullableExplorerSegment) Get() *ExplorerSegment {
	return v.value
}

func (v *NullableExplorerSegment) Set(val *ExplorerSegment) {
	v.value = val
	v.isSet = true
}

func (v NullableExplorerSegment) IsSet() bool {
	return v.isSet
}

func (v *NullableExplorerSegment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExplorerSegment(val *ExplorerSegment) *NullableExplorerSegment {
	return &NullableExplorerSegment{value: val, isSet: true}
}

func (v NullableExplorerSegment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExplorerSegment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

