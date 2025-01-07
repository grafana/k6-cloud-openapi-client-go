# TestRun

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Id of the test run. | [readonly] 
**TestId** | **int32** | Id of the parent test. | [readonly] 
**ProjectId** | **NullableInt32** | Id of the parent project. | [readonly] 
**StartedBy** | **NullableInt32** | Id of the user who started the test if started with a user token. | [readonly] 
**Created** | **time.Time** | Date and time when the test run was started. | [readonly] 
**Ended** | **NullableTime** | Date and time when the test run ended. Unset if the test is still running. | [readonly] 
**Note** | **string** | User-defined note for the test run. | [readonly] 
**RetentionExpiry** | **NullableTime** | The expiry date of test run results retention beyond which the data is automatically deleted if the test tun is not saved, otherwise - null. | [readonly] 
**Cost** | [**NullableTestCost**](TestCost.md) |  | 
**Status** | **string** | Current test run status. | 
**StatusDetails** | [**TestRunStatus**](TestRunStatus.md) | Details of the current test run status. | 
**StatusHistory** | [**[]TestRunStatus**](TestRunStatus.md) | List of test run status objects sorted by enter time representing the status history. | 
**Result** | **NullableString** | Test run result. If thresholds are defined and have been tainted, the result is &#x60;&#39;passed&#39;&#x60;, otherwise - &#x60;&#39;failed&#39;&#x60;. If the execution had not completed successfully, the result is &#x60;&#39;error&#39;&#x60;. The result is available only after the test is no longer running, otherwise it is &#x60;null&#x60;. | 
**ResultDetails** | **map[string]interface{}** | Test run result details. | 
**Distribution** | [**[]TestRunDistribution**](TestRunDistribution.md) | List the load zones the test runs in and the corresponding load distribution percent. | 
**Options** | **map[string]interface{}** | The original options object if available. | 

## Methods

### NewTestRun

`func NewTestRun(id int32, testId int32, projectId NullableInt32, startedBy NullableInt32, created time.Time, ended NullableTime, note string, retentionExpiry NullableTime, cost NullableTestCost, status string, statusDetails TestRunStatus, statusHistory []TestRunStatus, result NullableString, resultDetails map[string]interface{}, distribution []TestRunDistribution, options map[string]interface{}, ) *TestRun`

NewTestRun instantiates a new TestRun object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestRunWithDefaults

`func NewTestRunWithDefaults() *TestRun`

NewTestRunWithDefaults instantiates a new TestRun object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TestRun) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TestRun) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TestRun) SetId(v int32)`

SetId sets Id field to given value.


### GetTestId

`func (o *TestRun) GetTestId() int32`

GetTestId returns the TestId field if non-nil, zero value otherwise.

### GetTestIdOk

`func (o *TestRun) GetTestIdOk() (*int32, bool)`

GetTestIdOk returns a tuple with the TestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestId

`func (o *TestRun) SetTestId(v int32)`

SetTestId sets TestId field to given value.


### GetProjectId

`func (o *TestRun) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *TestRun) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *TestRun) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.


### SetProjectIdNil

`func (o *TestRun) SetProjectIdNil(b bool)`

 SetProjectIdNil sets the value for ProjectId to be an explicit nil

### UnsetProjectId
`func (o *TestRun) UnsetProjectId()`

UnsetProjectId ensures that no value is present for ProjectId, not even an explicit nil
### GetStartedBy

`func (o *TestRun) GetStartedBy() int32`

GetStartedBy returns the StartedBy field if non-nil, zero value otherwise.

### GetStartedByOk

`func (o *TestRun) GetStartedByOk() (*int32, bool)`

GetStartedByOk returns a tuple with the StartedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedBy

`func (o *TestRun) SetStartedBy(v int32)`

SetStartedBy sets StartedBy field to given value.


### SetStartedByNil

`func (o *TestRun) SetStartedByNil(b bool)`

 SetStartedByNil sets the value for StartedBy to be an explicit nil

### UnsetStartedBy
`func (o *TestRun) UnsetStartedBy()`

UnsetStartedBy ensures that no value is present for StartedBy, not even an explicit nil
### GetCreated

`func (o *TestRun) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *TestRun) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *TestRun) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetEnded

`func (o *TestRun) GetEnded() time.Time`

GetEnded returns the Ended field if non-nil, zero value otherwise.

### GetEndedOk

`func (o *TestRun) GetEndedOk() (*time.Time, bool)`

GetEndedOk returns a tuple with the Ended field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnded

`func (o *TestRun) SetEnded(v time.Time)`

SetEnded sets Ended field to given value.


### SetEndedNil

`func (o *TestRun) SetEndedNil(b bool)`

 SetEndedNil sets the value for Ended to be an explicit nil

### UnsetEnded
`func (o *TestRun) UnsetEnded()`

UnsetEnded ensures that no value is present for Ended, not even an explicit nil
### GetNote

`func (o *TestRun) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *TestRun) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *TestRun) SetNote(v string)`

SetNote sets Note field to given value.


### GetRetentionExpiry

`func (o *TestRun) GetRetentionExpiry() time.Time`

GetRetentionExpiry returns the RetentionExpiry field if non-nil, zero value otherwise.

### GetRetentionExpiryOk

`func (o *TestRun) GetRetentionExpiryOk() (*time.Time, bool)`

GetRetentionExpiryOk returns a tuple with the RetentionExpiry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionExpiry

`func (o *TestRun) SetRetentionExpiry(v time.Time)`

SetRetentionExpiry sets RetentionExpiry field to given value.


### SetRetentionExpiryNil

`func (o *TestRun) SetRetentionExpiryNil(b bool)`

 SetRetentionExpiryNil sets the value for RetentionExpiry to be an explicit nil

### UnsetRetentionExpiry
`func (o *TestRun) UnsetRetentionExpiry()`

UnsetRetentionExpiry ensures that no value is present for RetentionExpiry, not even an explicit nil
### GetCost

`func (o *TestRun) GetCost() TestCost`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *TestRun) GetCostOk() (*TestCost, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *TestRun) SetCost(v TestCost)`

SetCost sets Cost field to given value.


### SetCostNil

`func (o *TestRun) SetCostNil(b bool)`

 SetCostNil sets the value for Cost to be an explicit nil

### UnsetCost
`func (o *TestRun) UnsetCost()`

UnsetCost ensures that no value is present for Cost, not even an explicit nil
### GetStatus

`func (o *TestRun) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TestRun) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TestRun) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetStatusDetails

`func (o *TestRun) GetStatusDetails() TestRunStatus`

GetStatusDetails returns the StatusDetails field if non-nil, zero value otherwise.

### GetStatusDetailsOk

`func (o *TestRun) GetStatusDetailsOk() (*TestRunStatus, bool)`

GetStatusDetailsOk returns a tuple with the StatusDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusDetails

`func (o *TestRun) SetStatusDetails(v TestRunStatus)`

SetStatusDetails sets StatusDetails field to given value.


### GetStatusHistory

`func (o *TestRun) GetStatusHistory() []TestRunStatus`

GetStatusHistory returns the StatusHistory field if non-nil, zero value otherwise.

### GetStatusHistoryOk

`func (o *TestRun) GetStatusHistoryOk() (*[]TestRunStatus, bool)`

GetStatusHistoryOk returns a tuple with the StatusHistory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusHistory

`func (o *TestRun) SetStatusHistory(v []TestRunStatus)`

SetStatusHistory sets StatusHistory field to given value.


### GetResult

`func (o *TestRun) GetResult() string`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *TestRun) GetResultOk() (*string, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *TestRun) SetResult(v string)`

SetResult sets Result field to given value.


### SetResultNil

`func (o *TestRun) SetResultNil(b bool)`

 SetResultNil sets the value for Result to be an explicit nil

### UnsetResult
`func (o *TestRun) UnsetResult()`

UnsetResult ensures that no value is present for Result, not even an explicit nil
### GetResultDetails

`func (o *TestRun) GetResultDetails() map[string]interface{}`

GetResultDetails returns the ResultDetails field if non-nil, zero value otherwise.

### GetResultDetailsOk

`func (o *TestRun) GetResultDetailsOk() (*map[string]interface{}, bool)`

GetResultDetailsOk returns a tuple with the ResultDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultDetails

`func (o *TestRun) SetResultDetails(v map[string]interface{})`

SetResultDetails sets ResultDetails field to given value.


### GetDistribution

`func (o *TestRun) GetDistribution() []TestRunDistribution`

GetDistribution returns the Distribution field if non-nil, zero value otherwise.

### GetDistributionOk

`func (o *TestRun) GetDistributionOk() (*[]TestRunDistribution, bool)`

GetDistributionOk returns a tuple with the Distribution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistribution

`func (o *TestRun) SetDistribution(v []TestRunDistribution)`

SetDistribution sets Distribution field to given value.


### SetDistributionNil

`func (o *TestRun) SetDistributionNil(b bool)`

 SetDistributionNil sets the value for Distribution to be an explicit nil

### UnsetDistribution
`func (o *TestRun) UnsetDistribution()`

UnsetDistribution ensures that no value is present for Distribution, not even an explicit nil
### GetOptions

`func (o *TestRun) GetOptions() map[string]interface{}`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *TestRun) GetOptionsOk() (*map[string]interface{}, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *TestRun) SetOptions(v map[string]interface{})`

SetOptions sets Options field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


