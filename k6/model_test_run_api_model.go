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
	"time"
)

// checks if the TestRunApiModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TestRunApiModel{}

// TestRunApiModel The API model for a test run.
type TestRunApiModel struct {
	// ID of the test run.
	Id int32 `json:"id"`
	// ID of the parent test.
	TestId int32 `json:"test_id"`
	// ID of the parent project.
	ProjectId int32 `json:"project_id"`
	// Email of the user who started the test if started with a user token.
	StartedBy NullableString `json:"started_by"`
	// Date and time when the test run was started.
	Created time.Time `json:"created"`
	// Date and time when the test run ended. Unset if the test is still running.
	Ended NullableTime `json:"ended"`
	// User-defined note for the test run.
	Note string `json:"note"`
	// Timestamp after which the test run results are deleted or null if the test run is saved.
	RetentionExpiry NullableTime             `json:"retention_expiry"`
	Cost            NullableTestCostApiModel `json:"cost"`
	// Current test run status.
	Status string `json:"status"`
	// Details of the current test run status.
	StatusDetails StatusApiModel `json:"status_details"`
	// List of test run status objects sorted by the enter time representing the status history.
	StatusHistory []StatusApiModel `json:"status_history"`
	// List of the load zones configured for the test and the corresponding distribution percentages.
	Distribution []DistributionZoneApiModel `json:"distribution"`
	// Test run result. `passed` if there were no issues, `failed` if thresholds were breached, `error` if the execution was not completed.
	Result NullableString `json:"result"`
	// Additional information about the test run result.
	ResultDetails map[string]interface{} `json:"result_details"`
	// The original options object if available.
	Options              map[string]interface{} `json:"options"`
	AdditionalProperties map[string]interface{}
}

type _TestRunApiModel TestRunApiModel

// NewTestRunApiModel instantiates a new TestRunApiModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTestRunApiModel(id int32, testId int32, projectId int32, startedBy NullableString, created time.Time, ended NullableTime, note string, retentionExpiry NullableTime, cost NullableTestCostApiModel, status string, statusDetails StatusApiModel, statusHistory []StatusApiModel, distribution []DistributionZoneApiModel, result NullableString, resultDetails map[string]interface{}, options map[string]interface{}) *TestRunApiModel {
	this := TestRunApiModel{}
	this.Id = id
	this.TestId = testId
	this.ProjectId = projectId
	this.StartedBy = startedBy
	this.Created = created
	this.Ended = ended
	this.Note = note
	this.RetentionExpiry = retentionExpiry
	this.Cost = cost
	this.Status = status
	this.StatusDetails = statusDetails
	this.StatusHistory = statusHistory
	this.Distribution = distribution
	this.Result = result
	this.ResultDetails = resultDetails
	this.Options = options
	return &this
}

// NewTestRunApiModelWithDefaults instantiates a new TestRunApiModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTestRunApiModelWithDefaults() *TestRunApiModel {
	this := TestRunApiModel{}
	return &this
}

// GetId returns the Id field value
func (o *TestRunApiModel) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *TestRunApiModel) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *TestRunApiModel) SetId(v int32) {
	o.Id = v
}

// GetTestId returns the TestId field value
func (o *TestRunApiModel) GetTestId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TestId
}

// GetTestIdOk returns a tuple with the TestId field value
// and a boolean to check if the value has been set.
func (o *TestRunApiModel) GetTestIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TestId, true
}

// SetTestId sets field value
func (o *TestRunApiModel) SetTestId(v int32) {
	o.TestId = v
}

// GetProjectId returns the ProjectId field value
func (o *TestRunApiModel) GetProjectId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
func (o *TestRunApiModel) GetProjectIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectId, true
}

// SetProjectId sets field value
func (o *TestRunApiModel) SetProjectId(v int32) {
	o.ProjectId = v
}

// GetStartedBy returns the StartedBy field value
// If the value is explicit nil, the zero value for string will be returned
func (o *TestRunApiModel) GetStartedBy() string {
	if o == nil || o.StartedBy.Get() == nil {
		var ret string
		return ret
	}

	return *o.StartedBy.Get()
}

// GetStartedByOk returns a tuple with the StartedBy field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestRunApiModel) GetStartedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.StartedBy.Get(), o.StartedBy.IsSet()
}

// SetStartedBy sets field value
func (o *TestRunApiModel) SetStartedBy(v string) {
	o.StartedBy.Set(&v)
}

// GetCreated returns the Created field value
func (o *TestRunApiModel) GetCreated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Created
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
func (o *TestRunApiModel) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Created, true
}

// SetCreated sets field value
func (o *TestRunApiModel) SetCreated(v time.Time) {
	o.Created = v
}

// GetEnded returns the Ended field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *TestRunApiModel) GetEnded() time.Time {
	if o == nil || o.Ended.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.Ended.Get()
}

// GetEndedOk returns a tuple with the Ended field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestRunApiModel) GetEndedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Ended.Get(), o.Ended.IsSet()
}

// SetEnded sets field value
func (o *TestRunApiModel) SetEnded(v time.Time) {
	o.Ended.Set(&v)
}

// GetNote returns the Note field value
func (o *TestRunApiModel) GetNote() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Note
}

// GetNoteOk returns a tuple with the Note field value
// and a boolean to check if the value has been set.
func (o *TestRunApiModel) GetNoteOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Note, true
}

// SetNote sets field value
func (o *TestRunApiModel) SetNote(v string) {
	o.Note = v
}

// GetRetentionExpiry returns the RetentionExpiry field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *TestRunApiModel) GetRetentionExpiry() time.Time {
	if o == nil || o.RetentionExpiry.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.RetentionExpiry.Get()
}

// GetRetentionExpiryOk returns a tuple with the RetentionExpiry field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestRunApiModel) GetRetentionExpiryOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.RetentionExpiry.Get(), o.RetentionExpiry.IsSet()
}

// SetRetentionExpiry sets field value
func (o *TestRunApiModel) SetRetentionExpiry(v time.Time) {
	o.RetentionExpiry.Set(&v)
}

// GetCost returns the Cost field value
// If the value is explicit nil, the zero value for TestCostApiModel will be returned
func (o *TestRunApiModel) GetCost() TestCostApiModel {
	if o == nil || o.Cost.Get() == nil {
		var ret TestCostApiModel
		return ret
	}

	return *o.Cost.Get()
}

// GetCostOk returns a tuple with the Cost field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestRunApiModel) GetCostOk() (*TestCostApiModel, bool) {
	if o == nil {
		return nil, false
	}
	return o.Cost.Get(), o.Cost.IsSet()
}

// SetCost sets field value
func (o *TestRunApiModel) SetCost(v TestCostApiModel) {
	o.Cost.Set(&v)
}

// GetStatus returns the Status field value
func (o *TestRunApiModel) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *TestRunApiModel) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *TestRunApiModel) SetStatus(v string) {
	o.Status = v
}

// GetStatusDetails returns the StatusDetails field value
func (o *TestRunApiModel) GetStatusDetails() StatusApiModel {
	if o == nil {
		var ret StatusApiModel
		return ret
	}

	return o.StatusDetails
}

// GetStatusDetailsOk returns a tuple with the StatusDetails field value
// and a boolean to check if the value has been set.
func (o *TestRunApiModel) GetStatusDetailsOk() (*StatusApiModel, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StatusDetails, true
}

// SetStatusDetails sets field value
func (o *TestRunApiModel) SetStatusDetails(v StatusApiModel) {
	o.StatusDetails = v
}

// GetStatusHistory returns the StatusHistory field value
func (o *TestRunApiModel) GetStatusHistory() []StatusApiModel {
	if o == nil {
		var ret []StatusApiModel
		return ret
	}

	return o.StatusHistory
}

// GetStatusHistoryOk returns a tuple with the StatusHistory field value
// and a boolean to check if the value has been set.
func (o *TestRunApiModel) GetStatusHistoryOk() ([]StatusApiModel, bool) {
	if o == nil {
		return nil, false
	}
	return o.StatusHistory, true
}

// SetStatusHistory sets field value
func (o *TestRunApiModel) SetStatusHistory(v []StatusApiModel) {
	o.StatusHistory = v
}

// GetDistribution returns the Distribution field value
// If the value is explicit nil, the zero value for []DistributionZoneApiModel will be returned
func (o *TestRunApiModel) GetDistribution() []DistributionZoneApiModel {
	if o == nil {
		var ret []DistributionZoneApiModel
		return ret
	}

	return o.Distribution
}

// GetDistributionOk returns a tuple with the Distribution field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestRunApiModel) GetDistributionOk() ([]DistributionZoneApiModel, bool) {
	if o == nil || IsNil(o.Distribution) {
		return nil, false
	}
	return o.Distribution, true
}

// SetDistribution sets field value
func (o *TestRunApiModel) SetDistribution(v []DistributionZoneApiModel) {
	o.Distribution = v
}

// GetResult returns the Result field value
// If the value is explicit nil, the zero value for string will be returned
func (o *TestRunApiModel) GetResult() string {
	if o == nil || o.Result.Get() == nil {
		var ret string
		return ret
	}

	return *o.Result.Get()
}

// GetResultOk returns a tuple with the Result field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestRunApiModel) GetResultOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Result.Get(), o.Result.IsSet()
}

// SetResult sets field value
func (o *TestRunApiModel) SetResult(v string) {
	o.Result.Set(&v)
}

// GetResultDetails returns the ResultDetails field value
// If the value is explicit nil, the zero value for map[string]interface{} will be returned
func (o *TestRunApiModel) GetResultDetails() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.ResultDetails
}

// GetResultDetailsOk returns a tuple with the ResultDetails field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestRunApiModel) GetResultDetailsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.ResultDetails) {
		return map[string]interface{}{}, false
	}
	return o.ResultDetails, true
}

// SetResultDetails sets field value
func (o *TestRunApiModel) SetResultDetails(v map[string]interface{}) {
	o.ResultDetails = v
}

// GetOptions returns the Options field value
// If the value is explicit nil, the zero value for map[string]interface{} will be returned
func (o *TestRunApiModel) GetOptions() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestRunApiModel) GetOptionsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Options) {
		return map[string]interface{}{}, false
	}
	return o.Options, true
}

// SetOptions sets field value
func (o *TestRunApiModel) SetOptions(v map[string]interface{}) {
	o.Options = v
}

func (o TestRunApiModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TestRunApiModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["test_id"] = o.TestId
	toSerialize["project_id"] = o.ProjectId
	toSerialize["started_by"] = o.StartedBy.Get()
	toSerialize["created"] = o.Created
	toSerialize["ended"] = o.Ended.Get()
	toSerialize["note"] = o.Note
	toSerialize["retention_expiry"] = o.RetentionExpiry.Get()
	toSerialize["cost"] = o.Cost.Get()
	toSerialize["status"] = o.Status
	toSerialize["status_details"] = o.StatusDetails
	toSerialize["status_history"] = o.StatusHistory
	if o.Distribution != nil {
		toSerialize["distribution"] = o.Distribution
	}
	toSerialize["result"] = o.Result.Get()
	if o.ResultDetails != nil {
		toSerialize["result_details"] = o.ResultDetails
	}
	if o.Options != nil {
		toSerialize["options"] = o.Options
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TestRunApiModel) UnmarshalJSON(data []byte) (err error) {
	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	varTestRunApiModel := _TestRunApiModel{}

	err = json.Unmarshal(data, &varTestRunApiModel)

	if err != nil {
		return err
	}

	*o = TestRunApiModel(varTestRunApiModel)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "test_id")
		delete(additionalProperties, "project_id")
		delete(additionalProperties, "started_by")
		delete(additionalProperties, "created")
		delete(additionalProperties, "ended")
		delete(additionalProperties, "note")
		delete(additionalProperties, "retention_expiry")
		delete(additionalProperties, "cost")
		delete(additionalProperties, "status")
		delete(additionalProperties, "status_details")
		delete(additionalProperties, "status_history")
		delete(additionalProperties, "distribution")
		delete(additionalProperties, "result")
		delete(additionalProperties, "result_details")
		delete(additionalProperties, "options")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTestRunApiModel struct {
	value *TestRunApiModel
	isSet bool
}

func (v NullableTestRunApiModel) Get() *TestRunApiModel {
	return v.value
}

func (v *NullableTestRunApiModel) Set(val *TestRunApiModel) {
	v.value = val
	v.isSet = true
}

func (v NullableTestRunApiModel) IsSet() bool {
	return v.isSet
}

func (v *NullableTestRunApiModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTestRunApiModel(val *TestRunApiModel) *NullableTestRunApiModel {
	return &NullableTestRunApiModel{value: val, isSet: true}
}

func (v NullableTestRunApiModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTestRunApiModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
