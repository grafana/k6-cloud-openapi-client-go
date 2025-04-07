/*
Grafana Cloud k6

HTTP API for interacting with Grafana Cloud k6.

API version: 1.1.0
Contact: info@grafana.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package k6

import (
	"encoding/json"
	"fmt"
)

// checks if the CostBreakdownApiModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CostBreakdownApiModel{}

// CostBreakdownApiModel Breakdown details of the test cost.
type CostBreakdownApiModel struct {
	ProtocolVuh   ProtocolVuh   `json:"protocol_vuh"`
	BrowserVuh    BrowserVuh    `json:"browser_vuh"`
	BaseTotalVuh  BaseTotalVuh  `json:"base_total_vuh"`
	ReductionRate ReductionRate `json:"reduction_rate"`
	// The individual reduction rates applied to the base VUH usage.
	ReductionRateBreakdown map[string]ReductionRateBreakdownValue `json:"reduction_rate_breakdown"`
	AdditionalProperties   map[string]interface{}
}

type _CostBreakdownApiModel CostBreakdownApiModel

// NewCostBreakdownApiModel instantiates a new CostBreakdownApiModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCostBreakdownApiModel(protocolVuh ProtocolVuh, browserVuh BrowserVuh, baseTotalVuh BaseTotalVuh, reductionRate ReductionRate, reductionRateBreakdown map[string]ReductionRateBreakdownValue) *CostBreakdownApiModel {
	this := CostBreakdownApiModel{}
	this.ProtocolVuh = protocolVuh
	this.BrowserVuh = browserVuh
	this.BaseTotalVuh = baseTotalVuh
	this.ReductionRate = reductionRate
	this.ReductionRateBreakdown = reductionRateBreakdown
	return &this
}

// NewCostBreakdownApiModelWithDefaults instantiates a new CostBreakdownApiModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCostBreakdownApiModelWithDefaults() *CostBreakdownApiModel {
	this := CostBreakdownApiModel{}
	return &this
}

// GetProtocolVuh returns the ProtocolVuh field value
func (o *CostBreakdownApiModel) GetProtocolVuh() ProtocolVuh {
	if o == nil {
		var ret ProtocolVuh
		return ret
	}

	return o.ProtocolVuh
}

// GetProtocolVuhOk returns a tuple with the ProtocolVuh field value
// and a boolean to check if the value has been set.
func (o *CostBreakdownApiModel) GetProtocolVuhOk() (*ProtocolVuh, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProtocolVuh, true
}

// SetProtocolVuh sets field value
func (o *CostBreakdownApiModel) SetProtocolVuh(v ProtocolVuh) {
	o.ProtocolVuh = v
}

// GetBrowserVuh returns the BrowserVuh field value
func (o *CostBreakdownApiModel) GetBrowserVuh() BrowserVuh {
	if o == nil {
		var ret BrowserVuh
		return ret
	}

	return o.BrowserVuh
}

// GetBrowserVuhOk returns a tuple with the BrowserVuh field value
// and a boolean to check if the value has been set.
func (o *CostBreakdownApiModel) GetBrowserVuhOk() (*BrowserVuh, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BrowserVuh, true
}

// SetBrowserVuh sets field value
func (o *CostBreakdownApiModel) SetBrowserVuh(v BrowserVuh) {
	o.BrowserVuh = v
}

// GetBaseTotalVuh returns the BaseTotalVuh field value
func (o *CostBreakdownApiModel) GetBaseTotalVuh() BaseTotalVuh {
	if o == nil {
		var ret BaseTotalVuh
		return ret
	}

	return o.BaseTotalVuh
}

// GetBaseTotalVuhOk returns a tuple with the BaseTotalVuh field value
// and a boolean to check if the value has been set.
func (o *CostBreakdownApiModel) GetBaseTotalVuhOk() (*BaseTotalVuh, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BaseTotalVuh, true
}

// SetBaseTotalVuh sets field value
func (o *CostBreakdownApiModel) SetBaseTotalVuh(v BaseTotalVuh) {
	o.BaseTotalVuh = v
}

// GetReductionRate returns the ReductionRate field value
func (o *CostBreakdownApiModel) GetReductionRate() ReductionRate {
	if o == nil {
		var ret ReductionRate
		return ret
	}

	return o.ReductionRate
}

// GetReductionRateOk returns a tuple with the ReductionRate field value
// and a boolean to check if the value has been set.
func (o *CostBreakdownApiModel) GetReductionRateOk() (*ReductionRate, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReductionRate, true
}

// SetReductionRate sets field value
func (o *CostBreakdownApiModel) SetReductionRate(v ReductionRate) {
	o.ReductionRate = v
}

// GetReductionRateBreakdown returns the ReductionRateBreakdown field value
func (o *CostBreakdownApiModel) GetReductionRateBreakdown() map[string]ReductionRateBreakdownValue {
	if o == nil {
		var ret map[string]ReductionRateBreakdownValue
		return ret
	}

	return o.ReductionRateBreakdown
}

// GetReductionRateBreakdownOk returns a tuple with the ReductionRateBreakdown field value
// and a boolean to check if the value has been set.
func (o *CostBreakdownApiModel) GetReductionRateBreakdownOk() (*map[string]ReductionRateBreakdownValue, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReductionRateBreakdown, true
}

// SetReductionRateBreakdown sets field value
func (o *CostBreakdownApiModel) SetReductionRateBreakdown(v map[string]ReductionRateBreakdownValue) {
	o.ReductionRateBreakdown = v
}

func (o CostBreakdownApiModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CostBreakdownApiModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["protocol_vuh"] = o.ProtocolVuh
	toSerialize["browser_vuh"] = o.BrowserVuh
	toSerialize["base_total_vuh"] = o.BaseTotalVuh
	toSerialize["reduction_rate"] = o.ReductionRate
	toSerialize["reduction_rate_breakdown"] = o.ReductionRateBreakdown

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CostBreakdownApiModel) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"protocol_vuh",
		"browser_vuh",
		"base_total_vuh",
		"reduction_rate",
		"reduction_rate_breakdown",
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

	varCostBreakdownApiModel := _CostBreakdownApiModel{}

	err = json.Unmarshal(data, &varCostBreakdownApiModel)

	if err != nil {
		return err
	}

	*o = CostBreakdownApiModel(varCostBreakdownApiModel)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "protocol_vuh")
		delete(additionalProperties, "browser_vuh")
		delete(additionalProperties, "base_total_vuh")
		delete(additionalProperties, "reduction_rate")
		delete(additionalProperties, "reduction_rate_breakdown")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCostBreakdownApiModel struct {
	value *CostBreakdownApiModel
	isSet bool
}

func (v NullableCostBreakdownApiModel) Get() *CostBreakdownApiModel {
	return v.value
}

func (v *NullableCostBreakdownApiModel) Set(val *CostBreakdownApiModel) {
	v.value = val
	v.isSet = true
}

func (v NullableCostBreakdownApiModel) IsSet() bool {
	return v.isSet
}

func (v *NullableCostBreakdownApiModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCostBreakdownApiModel(val *CostBreakdownApiModel) *NullableCostBreakdownApiModel {
	return &NullableCostBreakdownApiModel{value: val, isSet: true}
}

func (v NullableCostBreakdownApiModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCostBreakdownApiModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
