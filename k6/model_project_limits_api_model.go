/*
Grafana Cloud k6

HTTP API for interacting with Grafana Cloud k6.

API version: 1.0.0
Contact: info@grafana.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package k6

import (
	"encoding/json"
	"fmt"
)

// checks if the ProjectLimitsApiModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProjectLimitsApiModel{}

// ProjectLimitsApiModel struct for ProjectLimitsApiModel
type ProjectLimitsApiModel struct {
	// ID of the related project.
	ProjectId int32 `json:"project_id"`
	// Max amount of virtual user hours (VUH) used per one calendar month.
	VuhMaxPerMonth NullableInt32 `json:"vuh_max_per_month"`
	// Max number of concurrent virtual users (VUs) used in one test.
	VuMaxPerTest NullableInt32 `json:"vu_max_per_test"`
	// Max number of concurrent browser virtual users (VUs) used in one test.
	VuBrowserMaxPerTest NullableInt32 `json:"vu_browser_max_per_test"`
	// Max duration of a test in seconds.
	DurationMaxPerTest   NullableInt32 `json:"duration_max_per_test"`
	AdditionalProperties map[string]interface{}
}

type _ProjectLimitsApiModel ProjectLimitsApiModel

// NewProjectLimitsApiModel instantiates a new ProjectLimitsApiModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectLimitsApiModel(projectId int32, vuhMaxPerMonth NullableInt32, vuMaxPerTest NullableInt32, vuBrowserMaxPerTest NullableInt32, durationMaxPerTest NullableInt32) *ProjectLimitsApiModel {
	this := ProjectLimitsApiModel{}
	this.ProjectId = projectId
	this.VuhMaxPerMonth = vuhMaxPerMonth
	this.VuMaxPerTest = vuMaxPerTest
	this.VuBrowserMaxPerTest = vuBrowserMaxPerTest
	this.DurationMaxPerTest = durationMaxPerTest
	return &this
}

// NewProjectLimitsApiModelWithDefaults instantiates a new ProjectLimitsApiModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectLimitsApiModelWithDefaults() *ProjectLimitsApiModel {
	this := ProjectLimitsApiModel{}
	return &this
}

// GetProjectId returns the ProjectId field value
func (o *ProjectLimitsApiModel) GetProjectId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
func (o *ProjectLimitsApiModel) GetProjectIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectId, true
}

// SetProjectId sets field value
func (o *ProjectLimitsApiModel) SetProjectId(v int32) {
	o.ProjectId = v
}

// GetVuhMaxPerMonth returns the VuhMaxPerMonth field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *ProjectLimitsApiModel) GetVuhMaxPerMonth() int32 {
	if o == nil || o.VuhMaxPerMonth.Get() == nil {
		var ret int32
		return ret
	}

	return *o.VuhMaxPerMonth.Get()
}

// GetVuhMaxPerMonthOk returns a tuple with the VuhMaxPerMonth field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectLimitsApiModel) GetVuhMaxPerMonthOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.VuhMaxPerMonth.Get(), o.VuhMaxPerMonth.IsSet()
}

// SetVuhMaxPerMonth sets field value
func (o *ProjectLimitsApiModel) SetVuhMaxPerMonth(v int32) {
	o.VuhMaxPerMonth.Set(&v)
}

// GetVuMaxPerTest returns the VuMaxPerTest field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *ProjectLimitsApiModel) GetVuMaxPerTest() int32 {
	if o == nil || o.VuMaxPerTest.Get() == nil {
		var ret int32
		return ret
	}

	return *o.VuMaxPerTest.Get()
}

// GetVuMaxPerTestOk returns a tuple with the VuMaxPerTest field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectLimitsApiModel) GetVuMaxPerTestOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.VuMaxPerTest.Get(), o.VuMaxPerTest.IsSet()
}

// SetVuMaxPerTest sets field value
func (o *ProjectLimitsApiModel) SetVuMaxPerTest(v int32) {
	o.VuMaxPerTest.Set(&v)
}

// GetVuBrowserMaxPerTest returns the VuBrowserMaxPerTest field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *ProjectLimitsApiModel) GetVuBrowserMaxPerTest() int32 {
	if o == nil || o.VuBrowserMaxPerTest.Get() == nil {
		var ret int32
		return ret
	}

	return *o.VuBrowserMaxPerTest.Get()
}

// GetVuBrowserMaxPerTestOk returns a tuple with the VuBrowserMaxPerTest field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectLimitsApiModel) GetVuBrowserMaxPerTestOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.VuBrowserMaxPerTest.Get(), o.VuBrowserMaxPerTest.IsSet()
}

// SetVuBrowserMaxPerTest sets field value
func (o *ProjectLimitsApiModel) SetVuBrowserMaxPerTest(v int32) {
	o.VuBrowserMaxPerTest.Set(&v)
}

// GetDurationMaxPerTest returns the DurationMaxPerTest field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *ProjectLimitsApiModel) GetDurationMaxPerTest() int32 {
	if o == nil || o.DurationMaxPerTest.Get() == nil {
		var ret int32
		return ret
	}

	return *o.DurationMaxPerTest.Get()
}

// GetDurationMaxPerTestOk returns a tuple with the DurationMaxPerTest field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectLimitsApiModel) GetDurationMaxPerTestOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.DurationMaxPerTest.Get(), o.DurationMaxPerTest.IsSet()
}

// SetDurationMaxPerTest sets field value
func (o *ProjectLimitsApiModel) SetDurationMaxPerTest(v int32) {
	o.DurationMaxPerTest.Set(&v)
}

func (o ProjectLimitsApiModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProjectLimitsApiModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["project_id"] = o.ProjectId
	toSerialize["vuh_max_per_month"] = o.VuhMaxPerMonth.Get()
	toSerialize["vu_max_per_test"] = o.VuMaxPerTest.Get()
	toSerialize["vu_browser_max_per_test"] = o.VuBrowserMaxPerTest.Get()
	toSerialize["duration_max_per_test"] = o.DurationMaxPerTest.Get()

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ProjectLimitsApiModel) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"project_id",
		"vuh_max_per_month",
		"vu_max_per_test",
		"vu_browser_max_per_test",
		"duration_max_per_test",
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

	varProjectLimitsApiModel := _ProjectLimitsApiModel{}

	err = json.Unmarshal(data, &varProjectLimitsApiModel)

	if err != nil {
		return err
	}

	*o = ProjectLimitsApiModel(varProjectLimitsApiModel)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "project_id")
		delete(additionalProperties, "vuh_max_per_month")
		delete(additionalProperties, "vu_max_per_test")
		delete(additionalProperties, "vu_browser_max_per_test")
		delete(additionalProperties, "duration_max_per_test")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableProjectLimitsApiModel struct {
	value *ProjectLimitsApiModel
	isSet bool
}

func (v NullableProjectLimitsApiModel) Get() *ProjectLimitsApiModel {
	return v.value
}

func (v *NullableProjectLimitsApiModel) Set(val *ProjectLimitsApiModel) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectLimitsApiModel) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectLimitsApiModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectLimitsApiModel(val *ProjectLimitsApiModel) *NullableProjectLimitsApiModel {
	return &NullableProjectLimitsApiModel{value: val, isSet: true}
}

func (v NullableProjectLimitsApiModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectLimitsApiModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
