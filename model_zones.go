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

// checks if the Zones type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Zones{}

// Zones struct for Zones
type Zones struct {
	HeartRate *HeartRateZoneRanges `json:"heart_rate,omitempty"`
	Power *PowerZoneRanges `json:"power,omitempty"`
}

// NewZones instantiates a new Zones object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewZones() *Zones {
	this := Zones{}
	return &this
}

// NewZonesWithDefaults instantiates a new Zones object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewZonesWithDefaults() *Zones {
	this := Zones{}
	return &this
}

// GetHeartRate returns the HeartRate field value if set, zero value otherwise.
func (o *Zones) GetHeartRate() HeartRateZoneRanges {
	if o == nil || IsNil(o.HeartRate) {
		var ret HeartRateZoneRanges
		return ret
	}
	return *o.HeartRate
}

// GetHeartRateOk returns a tuple with the HeartRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Zones) GetHeartRateOk() (*HeartRateZoneRanges, bool) {
	if o == nil || IsNil(o.HeartRate) {
		return nil, false
	}
	return o.HeartRate, true
}

// HasHeartRate returns a boolean if a field has been set.
func (o *Zones) HasHeartRate() bool {
	if o != nil && !IsNil(o.HeartRate) {
		return true
	}

	return false
}

// SetHeartRate gets a reference to the given HeartRateZoneRanges and assigns it to the HeartRate field.
func (o *Zones) SetHeartRate(v HeartRateZoneRanges) {
	o.HeartRate = &v
}

// GetPower returns the Power field value if set, zero value otherwise.
func (o *Zones) GetPower() PowerZoneRanges {
	if o == nil || IsNil(o.Power) {
		var ret PowerZoneRanges
		return ret
	}
	return *o.Power
}

// GetPowerOk returns a tuple with the Power field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Zones) GetPowerOk() (*PowerZoneRanges, bool) {
	if o == nil || IsNil(o.Power) {
		return nil, false
	}
	return o.Power, true
}

// HasPower returns a boolean if a field has been set.
func (o *Zones) HasPower() bool {
	if o != nil && !IsNil(o.Power) {
		return true
	}

	return false
}

// SetPower gets a reference to the given PowerZoneRanges and assigns it to the Power field.
func (o *Zones) SetPower(v PowerZoneRanges) {
	o.Power = &v
}

func (o Zones) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Zones) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.HeartRate) {
		toSerialize["heart_rate"] = o.HeartRate
	}
	if !IsNil(o.Power) {
		toSerialize["power"] = o.Power
	}
	return toSerialize, nil
}

type NullableZones struct {
	value *Zones
	isSet bool
}

func (v NullableZones) Get() *Zones {
	return v.value
}

func (v *NullableZones) Set(val *Zones) {
	v.value = val
	v.isSet = true
}

func (v NullableZones) IsSet() bool {
	return v.isSet
}

func (v *NullableZones) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableZones(val *Zones) *NullableZones {
	return &NullableZones{value: val, isSet: true}
}

func (v NullableZones) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableZones) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


