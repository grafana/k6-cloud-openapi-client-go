/*
Grafana Cloud k6

HTTP API for interacting with Grafana Cloud k6.

API version: 1.5.0
Contact: info@grafana.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package k6

import (
	"encoding/json"
	"fmt"
)

// checks if the TestCostApiModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TestCostApiModel{}

// TestCostApiModel struct for TestCostApiModel
type TestCostApiModel struct {
	TotalVuh TotalVuh `json:"total_vuh"`
	Breakdown CostBreakdownApiModel `json:"breakdown"`
	AdditionalProperties map[string]interface{}
}

type _TestCostApiModel TestCostApiModel

// NewTestCostApiModel instantiates a new TestCostApiModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTestCostApiModel(totalVuh TotalVuh, breakdown CostBreakdownApiModel) *TestCostApiModel {
	this := TestCostApiModel{}
	this.TotalVuh = totalVuh
	this.Breakdown = breakdown
	return &this
}

// NewTestCostApiModelWithDefaults instantiates a new TestCostApiModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTestCostApiModelWithDefaults() *TestCostApiModel {
	this := TestCostApiModel{}
	return &this
}

// GetTotalVuh returns the TotalVuh field value
func (o *TestCostApiModel) GetTotalVuh() TotalVuh {
	if o == nil {
		var ret TotalVuh
		return ret
	}

	return o.TotalVuh
}

// GetTotalVuhOk returns a tuple with the TotalVuh field value
// and a boolean to check if the value has been set.
func (o *TestCostApiModel) GetTotalVuhOk() (*TotalVuh, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalVuh, true
}

// SetTotalVuh sets field value
func (o *TestCostApiModel) SetTotalVuh(v TotalVuh) {
	o.TotalVuh = v
}

// GetBreakdown returns the Breakdown field value
func (o *TestCostApiModel) GetBreakdown() CostBreakdownApiModel {
	if o == nil {
		var ret CostBreakdownApiModel
		return ret
	}

	return o.Breakdown
}

// GetBreakdownOk returns a tuple with the Breakdown field value
// and a boolean to check if the value has been set.
func (o *TestCostApiModel) GetBreakdownOk() (*CostBreakdownApiModel, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Breakdown, true
}

// SetBreakdown sets field value
func (o *TestCostApiModel) SetBreakdown(v CostBreakdownApiModel) {
	o.Breakdown = v
}

func (o TestCostApiModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TestCostApiModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["total_vuh"] = o.TotalVuh
	toSerialize["breakdown"] = o.Breakdown

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TestCostApiModel) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"total_vuh",
		"breakdown",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varTestCostApiModel := _TestCostApiModel{}

	err = json.Unmarshal(data, &varTestCostApiModel)

	if err != nil {
		return err
	}

	*o = TestCostApiModel(varTestCostApiModel)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "total_vuh")
		delete(additionalProperties, "breakdown")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTestCostApiModel struct {
	value *TestCostApiModel
	isSet bool
}

func (v NullableTestCostApiModel) Get() *TestCostApiModel {
	return v.value
}

func (v *NullableTestCostApiModel) Set(val *TestCostApiModel) {
	v.value = val
	v.isSet = true
}

func (v NullableTestCostApiModel) IsSet() bool {
	return v.isSet
}

func (v *NullableTestCostApiModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTestCostApiModel(val *TestCostApiModel) *NullableTestCostApiModel {
	return &NullableTestCostApiModel{value: val, isSet: true}
}

func (v NullableTestCostApiModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTestCostApiModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


