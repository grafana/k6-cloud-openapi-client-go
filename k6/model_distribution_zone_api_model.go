/*
Grafana Cloud k6

HTTP API for interacting with Grafana Cloud k6.

API version: 1.2.0
Contact: info@grafana.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package k6

import (
	"encoding/json"
	"fmt"
)

// checks if the DistributionZoneApiModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DistributionZoneApiModel{}

// DistributionZoneApiModel struct for DistributionZoneApiModel
type DistributionZoneApiModel struct {
	// Name of the load zone.
	LoadZone string `json:"load_zone"`
	// Percentage of the load that runs from the load zone.
	Percent              int32 `json:"percent"`
	AdditionalProperties map[string]interface{}
}

type _DistributionZoneApiModel DistributionZoneApiModel

// NewDistributionZoneApiModel instantiates a new DistributionZoneApiModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDistributionZoneApiModel(loadZone string, percent int32) *DistributionZoneApiModel {
	this := DistributionZoneApiModel{}
	this.LoadZone = loadZone
	this.Percent = percent
	return &this
}

// NewDistributionZoneApiModelWithDefaults instantiates a new DistributionZoneApiModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDistributionZoneApiModelWithDefaults() *DistributionZoneApiModel {
	this := DistributionZoneApiModel{}
	return &this
}

// GetLoadZone returns the LoadZone field value
func (o *DistributionZoneApiModel) GetLoadZone() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LoadZone
}

// GetLoadZoneOk returns a tuple with the LoadZone field value
// and a boolean to check if the value has been set.
func (o *DistributionZoneApiModel) GetLoadZoneOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LoadZone, true
}

// SetLoadZone sets field value
func (o *DistributionZoneApiModel) SetLoadZone(v string) {
	o.LoadZone = v
}

// GetPercent returns the Percent field value
func (o *DistributionZoneApiModel) GetPercent() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Percent
}

// GetPercentOk returns a tuple with the Percent field value
// and a boolean to check if the value has been set.
func (o *DistributionZoneApiModel) GetPercentOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Percent, true
}

// SetPercent sets field value
func (o *DistributionZoneApiModel) SetPercent(v int32) {
	o.Percent = v
}

func (o DistributionZoneApiModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DistributionZoneApiModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["load_zone"] = o.LoadZone
	toSerialize["percent"] = o.Percent

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DistributionZoneApiModel) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"load_zone",
		"percent",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varDistributionZoneApiModel := _DistributionZoneApiModel{}

	err = json.Unmarshal(data, &varDistributionZoneApiModel)

	if err != nil {
		return err
	}

	*o = DistributionZoneApiModel(varDistributionZoneApiModel)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "load_zone")
		delete(additionalProperties, "percent")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDistributionZoneApiModel struct {
	value *DistributionZoneApiModel
	isSet bool
}

func (v NullableDistributionZoneApiModel) Get() *DistributionZoneApiModel {
	return v.value
}

func (v *NullableDistributionZoneApiModel) Set(val *DistributionZoneApiModel) {
	v.value = val
	v.isSet = true
}

func (v NullableDistributionZoneApiModel) IsSet() bool {
	return v.isSet
}

func (v *NullableDistributionZoneApiModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDistributionZoneApiModel(val *DistributionZoneApiModel) *NullableDistributionZoneApiModel {
	return &NullableDistributionZoneApiModel{value: val, isSet: true}
}

func (v NullableDistributionZoneApiModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDistributionZoneApiModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
