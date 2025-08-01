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
	"time"
	"fmt"
)

// checks if the LoadTestApiModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LoadTestApiModel{}

// LoadTestApiModel struct for LoadTestApiModel
type LoadTestApiModel struct {
	// ID of the load test.
	Id int32 `json:"id"`
	// ID of the parent project.
	ProjectId int32 `json:"project_id"`
	// Unique name of the test within the project.
	Name string `json:"name"`
	// ID of a baseline test run used for results comparison.
	BaselineTestRunId NullableInt32 `json:"baseline_test_run_id"`
	// The date when the test was created.
	Created time.Time `json:"created"`
	// The date when the test was last updated.
	Updated time.Time `json:"updated"`
	AdditionalProperties map[string]interface{}
}

type _LoadTestApiModel LoadTestApiModel

// NewLoadTestApiModel instantiates a new LoadTestApiModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLoadTestApiModel(id int32, projectId int32, name string, baselineTestRunId NullableInt32, created time.Time, updated time.Time) *LoadTestApiModel {
	this := LoadTestApiModel{}
	this.Id = id
	this.ProjectId = projectId
	this.Name = name
	this.BaselineTestRunId = baselineTestRunId
	this.Created = created
	this.Updated = updated
	return &this
}

// NewLoadTestApiModelWithDefaults instantiates a new LoadTestApiModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLoadTestApiModelWithDefaults() *LoadTestApiModel {
	this := LoadTestApiModel{}
	return &this
}

// GetId returns the Id field value
func (o *LoadTestApiModel) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *LoadTestApiModel) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *LoadTestApiModel) SetId(v int32) {
	o.Id = v
}

// GetProjectId returns the ProjectId field value
func (o *LoadTestApiModel) GetProjectId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
func (o *LoadTestApiModel) GetProjectIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectId, true
}

// SetProjectId sets field value
func (o *LoadTestApiModel) SetProjectId(v int32) {
	o.ProjectId = v
}

// GetName returns the Name field value
func (o *LoadTestApiModel) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *LoadTestApiModel) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *LoadTestApiModel) SetName(v string) {
	o.Name = v
}

// GetBaselineTestRunId returns the BaselineTestRunId field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *LoadTestApiModel) GetBaselineTestRunId() int32 {
	if o == nil || o.BaselineTestRunId.Get() == nil {
		var ret int32
		return ret
	}

	return *o.BaselineTestRunId.Get()
}

// GetBaselineTestRunIdOk returns a tuple with the BaselineTestRunId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LoadTestApiModel) GetBaselineTestRunIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.BaselineTestRunId.Get(), o.BaselineTestRunId.IsSet()
}

// SetBaselineTestRunId sets field value
func (o *LoadTestApiModel) SetBaselineTestRunId(v int32) {
	o.BaselineTestRunId.Set(&v)
}

// GetCreated returns the Created field value
func (o *LoadTestApiModel) GetCreated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Created
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
func (o *LoadTestApiModel) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Created, true
}

// SetCreated sets field value
func (o *LoadTestApiModel) SetCreated(v time.Time) {
	o.Created = v
}

// GetUpdated returns the Updated field value
func (o *LoadTestApiModel) GetUpdated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Updated
}

// GetUpdatedOk returns a tuple with the Updated field value
// and a boolean to check if the value has been set.
func (o *LoadTestApiModel) GetUpdatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Updated, true
}

// SetUpdated sets field value
func (o *LoadTestApiModel) SetUpdated(v time.Time) {
	o.Updated = v
}

func (o LoadTestApiModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LoadTestApiModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["project_id"] = o.ProjectId
	toSerialize["name"] = o.Name
	toSerialize["baseline_test_run_id"] = o.BaselineTestRunId.Get()
	toSerialize["created"] = o.Created
	toSerialize["updated"] = o.Updated

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *LoadTestApiModel) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"project_id",
		"name",
		"baseline_test_run_id",
		"created",
		"updated",
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

	varLoadTestApiModel := _LoadTestApiModel{}

	err = json.Unmarshal(data, &varLoadTestApiModel)

	if err != nil {
		return err
	}

	*o = LoadTestApiModel(varLoadTestApiModel)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "project_id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "baseline_test_run_id")
		delete(additionalProperties, "created")
		delete(additionalProperties, "updated")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableLoadTestApiModel struct {
	value *LoadTestApiModel
	isSet bool
}

func (v NullableLoadTestApiModel) Get() *LoadTestApiModel {
	return v.value
}

func (v *NullableLoadTestApiModel) Set(val *LoadTestApiModel) {
	v.value = val
	v.isSet = true
}

func (v NullableLoadTestApiModel) IsSet() bool {
	return v.isSet
}

func (v *NullableLoadTestApiModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLoadTestApiModel(val *LoadTestApiModel) *NullableLoadTestApiModel {
	return &NullableLoadTestApiModel{value: val, isSet: true}
}

func (v NullableLoadTestApiModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLoadTestApiModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


