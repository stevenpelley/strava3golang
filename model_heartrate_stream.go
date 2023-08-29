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

// checks if the HeartrateStream type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HeartrateStream{}

// HeartrateStream struct for HeartrateStream
type HeartrateStream struct {
	// The number of data points in this stream
	OriginalSize *int32 `json:"original_size,omitempty"`
	// The level of detail (sampling) in which this stream was returned
	Resolution *string `json:"resolution,omitempty"`
	// The base series used in the case the stream was downsampled
	SeriesType *string `json:"series_type,omitempty"`
	// The sequence of heart rate values for this stream, in beats per minute
	Data []int32 `json:"data,omitempty"`
}

// NewHeartrateStream instantiates a new HeartrateStream object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHeartrateStream() *HeartrateStream {
	this := HeartrateStream{}
	return &this
}

// NewHeartrateStreamWithDefaults instantiates a new HeartrateStream object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHeartrateStreamWithDefaults() *HeartrateStream {
	this := HeartrateStream{}
	return &this
}

// GetOriginalSize returns the OriginalSize field value if set, zero value otherwise.
func (o *HeartrateStream) GetOriginalSize() int32 {
	if o == nil || IsNil(o.OriginalSize) {
		var ret int32
		return ret
	}
	return *o.OriginalSize
}

// GetOriginalSizeOk returns a tuple with the OriginalSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HeartrateStream) GetOriginalSizeOk() (*int32, bool) {
	if o == nil || IsNil(o.OriginalSize) {
		return nil, false
	}
	return o.OriginalSize, true
}

// HasOriginalSize returns a boolean if a field has been set.
func (o *HeartrateStream) HasOriginalSize() bool {
	if o != nil && !IsNil(o.OriginalSize) {
		return true
	}

	return false
}

// SetOriginalSize gets a reference to the given int32 and assigns it to the OriginalSize field.
func (o *HeartrateStream) SetOriginalSize(v int32) {
	o.OriginalSize = &v
}

// GetResolution returns the Resolution field value if set, zero value otherwise.
func (o *HeartrateStream) GetResolution() string {
	if o == nil || IsNil(o.Resolution) {
		var ret string
		return ret
	}
	return *o.Resolution
}

// GetResolutionOk returns a tuple with the Resolution field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HeartrateStream) GetResolutionOk() (*string, bool) {
	if o == nil || IsNil(o.Resolution) {
		return nil, false
	}
	return o.Resolution, true
}

// HasResolution returns a boolean if a field has been set.
func (o *HeartrateStream) HasResolution() bool {
	if o != nil && !IsNil(o.Resolution) {
		return true
	}

	return false
}

// SetResolution gets a reference to the given string and assigns it to the Resolution field.
func (o *HeartrateStream) SetResolution(v string) {
	o.Resolution = &v
}

// GetSeriesType returns the SeriesType field value if set, zero value otherwise.
func (o *HeartrateStream) GetSeriesType() string {
	if o == nil || IsNil(o.SeriesType) {
		var ret string
		return ret
	}
	return *o.SeriesType
}

// GetSeriesTypeOk returns a tuple with the SeriesType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HeartrateStream) GetSeriesTypeOk() (*string, bool) {
	if o == nil || IsNil(o.SeriesType) {
		return nil, false
	}
	return o.SeriesType, true
}

// HasSeriesType returns a boolean if a field has been set.
func (o *HeartrateStream) HasSeriesType() bool {
	if o != nil && !IsNil(o.SeriesType) {
		return true
	}

	return false
}

// SetSeriesType gets a reference to the given string and assigns it to the SeriesType field.
func (o *HeartrateStream) SetSeriesType(v string) {
	o.SeriesType = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *HeartrateStream) GetData() []int32 {
	if o == nil || IsNil(o.Data) {
		var ret []int32
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HeartrateStream) GetDataOk() ([]int32, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *HeartrateStream) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []int32 and assigns it to the Data field.
func (o *HeartrateStream) SetData(v []int32) {
	o.Data = v
}

func (o HeartrateStream) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HeartrateStream) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.OriginalSize) {
		toSerialize["original_size"] = o.OriginalSize
	}
	if !IsNil(o.Resolution) {
		toSerialize["resolution"] = o.Resolution
	}
	if !IsNil(o.SeriesType) {
		toSerialize["series_type"] = o.SeriesType
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableHeartrateStream struct {
	value *HeartrateStream
	isSet bool
}

func (v NullableHeartrateStream) Get() *HeartrateStream {
	return v.value
}

func (v *NullableHeartrateStream) Set(val *HeartrateStream) {
	v.value = val
	v.isSet = true
}

func (v NullableHeartrateStream) IsSet() bool {
	return v.isSet
}

func (v *NullableHeartrateStream) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHeartrateStream(val *HeartrateStream) *NullableHeartrateStream {
	return &NullableHeartrateStream{value: val, isSet: true}
}

func (v NullableHeartrateStream) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHeartrateStream) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

