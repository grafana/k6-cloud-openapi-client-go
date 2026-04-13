# StartLoadTestResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | ID of the test run. | 
**TestId** | **int32** | ID of the parent test. | 
**ProjectId** | **int32** | ID of the parent project. | 
**StartedBy** | **NullableString** | Email of the user who started the test if started with a user token. | 
**Created** | **time.Time** | Date and time when the test run was started. | 
**Ended** | **NullableTime** | Date and time when the test run ended. Unset if the test is still running. | 
**Note** | **string** | User-defined note for the test run. | 
**RetentionExpiry** | **NullableTime** | Timestamp after which the test run results are deleted or null if the test run is saved. | 
**Cost** | [**NullableTestCostApiModel**](TestCostApiModel.md) |  | 
**Status** | **string** | Current test run status. | 
**StatusDetails** | [**StatusApiModel**](StatusApiModel.md) | Details of the current test run status. | 
**StatusHistory** | [**[]StatusApiModel**](StatusApiModel.md) | List of test run status objects sorted by the status start time. The list represents the test run status history. | 
**Distribution** | [**[]DistributionZoneApiModel**](DistributionZoneApiModel.md) | List of the load zones configured for the test and the corresponding distribution percentages. | 
**Result** | **NullableString** | Test run result. &#x60;passed&#x60; if there were no issues, &#x60;failed&#x60; if thresholds were breached, &#x60;error&#x60; if the execution was not completed. | 
**ResultDetails** | **map[string]interface{}** | Additional information about the test run result. | 
**Options** | **map[string]interface{}** | The original options object if available. | 
**K6Dependencies** | **map[string]string** | The requested version of k6 and extensions that was part of the script/archive. | 
**K6Versions** | **map[string]string** | The computed version for k6 and extensions used to run the test. | 
**MaxVus** | **NullableInt32** | The maximum number of total VUs (browser and protocol) at any stage of the execution plan. | 
**MaxBrowserVus** | **NullableInt32** | The maximum number of browser VUs at any stage of the execution plan. | 
**EstimatedDuration** | **NullableInt32** | The estimated duration of the test run in seconds. | 
**ExecutionDuration** | **int32** | The real billable duration of the test run in seconds. | 
**TestRunDetailsPageUrl** | **string** | URL to the Grafana web app test run page. | 

## Methods

### NewStartLoadTestResponse

`func NewStartLoadTestResponse(id int32, testId int32, projectId int32, startedBy NullableString, created time.Time, ended NullableTime, note string, retentionExpiry NullableTime, cost NullableTestCostApiModel, status string, statusDetails StatusApiModel, statusHistory []StatusApiModel, distribution []DistributionZoneApiModel, result NullableString, resultDetails map[string]interface{}, options map[string]interface{}, k6Dependencies map[string]string, k6Versions map[string]string, maxVus NullableInt32, maxBrowserVus NullableInt32, estimatedDuration NullableInt32, executionDuration int32, testRunDetailsPageUrl string, ) *StartLoadTestResponse`

NewStartLoadTestResponse instantiates a new StartLoadTestResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStartLoadTestResponseWithDefaults

`func NewStartLoadTestResponseWithDefaults() *StartLoadTestResponse`

NewStartLoadTestResponseWithDefaults instantiates a new StartLoadTestResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *StartLoadTestResponse) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StartLoadTestResponse) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StartLoadTestResponse) SetId(v int32)`

SetId sets Id field to given value.


### GetTestId

`func (o *StartLoadTestResponse) GetTestId() int32`

GetTestId returns the TestId field if non-nil, zero value otherwise.

### GetTestIdOk

`func (o *StartLoadTestResponse) GetTestIdOk() (*int32, bool)`

GetTestIdOk returns a tuple with the TestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestId

`func (o *StartLoadTestResponse) SetTestId(v int32)`

SetTestId sets TestId field to given value.


### GetProjectId

`func (o *StartLoadTestResponse) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *StartLoadTestResponse) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *StartLoadTestResponse) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.


### GetStartedBy

`func (o *StartLoadTestResponse) GetStartedBy() string`

GetStartedBy returns the StartedBy field if non-nil, zero value otherwise.

### GetStartedByOk

`func (o *StartLoadTestResponse) GetStartedByOk() (*string, bool)`

GetStartedByOk returns a tuple with the StartedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedBy

`func (o *StartLoadTestResponse) SetStartedBy(v string)`

SetStartedBy sets StartedBy field to given value.


### SetStartedByNil

`func (o *StartLoadTestResponse) SetStartedByNil(b bool)`

 SetStartedByNil sets the value for StartedBy to be an explicit nil

### UnsetStartedBy
`func (o *StartLoadTestResponse) UnsetStartedBy()`

UnsetStartedBy ensures that no value is present for StartedBy, not even an explicit nil
### GetCreated

`func (o *StartLoadTestResponse) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *StartLoadTestResponse) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *StartLoadTestResponse) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetEnded

`func (o *StartLoadTestResponse) GetEnded() time.Time`

GetEnded returns the Ended field if non-nil, zero value otherwise.

### GetEndedOk

`func (o *StartLoadTestResponse) GetEndedOk() (*time.Time, bool)`

GetEndedOk returns a tuple with the Ended field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnded

`func (o *StartLoadTestResponse) SetEnded(v time.Time)`

SetEnded sets Ended field to given value.


### SetEndedNil

`func (o *StartLoadTestResponse) SetEndedNil(b bool)`

 SetEndedNil sets the value for Ended to be an explicit nil

### UnsetEnded
`func (o *StartLoadTestResponse) UnsetEnded()`

UnsetEnded ensures that no value is present for Ended, not even an explicit nil
### GetNote

`func (o *StartLoadTestResponse) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *StartLoadTestResponse) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *StartLoadTestResponse) SetNote(v string)`

SetNote sets Note field to given value.


### GetRetentionExpiry

`func (o *StartLoadTestResponse) GetRetentionExpiry() time.Time`

GetRetentionExpiry returns the RetentionExpiry field if non-nil, zero value otherwise.

### GetRetentionExpiryOk

`func (o *StartLoadTestResponse) GetRetentionExpiryOk() (*time.Time, bool)`

GetRetentionExpiryOk returns a tuple with the RetentionExpiry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionExpiry

`func (o *StartLoadTestResponse) SetRetentionExpiry(v time.Time)`

SetRetentionExpiry sets RetentionExpiry field to given value.


### SetRetentionExpiryNil

`func (o *StartLoadTestResponse) SetRetentionExpiryNil(b bool)`

 SetRetentionExpiryNil sets the value for RetentionExpiry to be an explicit nil

### UnsetRetentionExpiry
`func (o *StartLoadTestResponse) UnsetRetentionExpiry()`

UnsetRetentionExpiry ensures that no value is present for RetentionExpiry, not even an explicit nil
### GetCost

`func (o *StartLoadTestResponse) GetCost() TestCostApiModel`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *StartLoadTestResponse) GetCostOk() (*TestCostApiModel, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *StartLoadTestResponse) SetCost(v TestCostApiModel)`

SetCost sets Cost field to given value.


### SetCostNil

`func (o *StartLoadTestResponse) SetCostNil(b bool)`

 SetCostNil sets the value for Cost to be an explicit nil

### UnsetCost
`func (o *StartLoadTestResponse) UnsetCost()`

UnsetCost ensures that no value is present for Cost, not even an explicit nil
### GetStatus

`func (o *StartLoadTestResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *StartLoadTestResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *StartLoadTestResponse) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetStatusDetails

`func (o *StartLoadTestResponse) GetStatusDetails() StatusApiModel`

GetStatusDetails returns the StatusDetails field if non-nil, zero value otherwise.

### GetStatusDetailsOk

`func (o *StartLoadTestResponse) GetStatusDetailsOk() (*StatusApiModel, bool)`

GetStatusDetailsOk returns a tuple with the StatusDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusDetails

`func (o *StartLoadTestResponse) SetStatusDetails(v StatusApiModel)`

SetStatusDetails sets StatusDetails field to given value.


### GetStatusHistory

`func (o *StartLoadTestResponse) GetStatusHistory() []StatusApiModel`

GetStatusHistory returns the StatusHistory field if non-nil, zero value otherwise.

### GetStatusHistoryOk

`func (o *StartLoadTestResponse) GetStatusHistoryOk() (*[]StatusApiModel, bool)`

GetStatusHistoryOk returns a tuple with the StatusHistory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusHistory

`func (o *StartLoadTestResponse) SetStatusHistory(v []StatusApiModel)`

SetStatusHistory sets StatusHistory field to given value.


### GetDistribution

`func (o *StartLoadTestResponse) GetDistribution() []DistributionZoneApiModel`

GetDistribution returns the Distribution field if non-nil, zero value otherwise.

### GetDistributionOk

`func (o *StartLoadTestResponse) GetDistributionOk() (*[]DistributionZoneApiModel, bool)`

GetDistributionOk returns a tuple with the Distribution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistribution

`func (o *StartLoadTestResponse) SetDistribution(v []DistributionZoneApiModel)`

SetDistribution sets Distribution field to given value.


### SetDistributionNil

`func (o *StartLoadTestResponse) SetDistributionNil(b bool)`

 SetDistributionNil sets the value for Distribution to be an explicit nil

### UnsetDistribution
`func (o *StartLoadTestResponse) UnsetDistribution()`

UnsetDistribution ensures that no value is present for Distribution, not even an explicit nil
### GetResult

`func (o *StartLoadTestResponse) GetResult() string`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *StartLoadTestResponse) GetResultOk() (*string, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *StartLoadTestResponse) SetResult(v string)`

SetResult sets Result field to given value.


### SetResultNil

`func (o *StartLoadTestResponse) SetResultNil(b bool)`

 SetResultNil sets the value for Result to be an explicit nil

### UnsetResult
`func (o *StartLoadTestResponse) UnsetResult()`

UnsetResult ensures that no value is present for Result, not even an explicit nil
### GetResultDetails

`func (o *StartLoadTestResponse) GetResultDetails() map[string]interface{}`

GetResultDetails returns the ResultDetails field if non-nil, zero value otherwise.

### GetResultDetailsOk

`func (o *StartLoadTestResponse) GetResultDetailsOk() (*map[string]interface{}, bool)`

GetResultDetailsOk returns a tuple with the ResultDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultDetails

`func (o *StartLoadTestResponse) SetResultDetails(v map[string]interface{})`

SetResultDetails sets ResultDetails field to given value.


### SetResultDetailsNil

`func (o *StartLoadTestResponse) SetResultDetailsNil(b bool)`

 SetResultDetailsNil sets the value for ResultDetails to be an explicit nil

### UnsetResultDetails
`func (o *StartLoadTestResponse) UnsetResultDetails()`

UnsetResultDetails ensures that no value is present for ResultDetails, not even an explicit nil
### GetOptions

`func (o *StartLoadTestResponse) GetOptions() map[string]interface{}`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *StartLoadTestResponse) GetOptionsOk() (*map[string]interface{}, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *StartLoadTestResponse) SetOptions(v map[string]interface{})`

SetOptions sets Options field to given value.


### SetOptionsNil

`func (o *StartLoadTestResponse) SetOptionsNil(b bool)`

 SetOptionsNil sets the value for Options to be an explicit nil

### UnsetOptions
`func (o *StartLoadTestResponse) UnsetOptions()`

UnsetOptions ensures that no value is present for Options, not even an explicit nil
### GetK6Dependencies

`func (o *StartLoadTestResponse) GetK6Dependencies() map[string]string`

GetK6Dependencies returns the K6Dependencies field if non-nil, zero value otherwise.

### GetK6DependenciesOk

`func (o *StartLoadTestResponse) GetK6DependenciesOk() (*map[string]string, bool)`

GetK6DependenciesOk returns a tuple with the K6Dependencies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetK6Dependencies

`func (o *StartLoadTestResponse) SetK6Dependencies(v map[string]string)`

SetK6Dependencies sets K6Dependencies field to given value.


### GetK6Versions

`func (o *StartLoadTestResponse) GetK6Versions() map[string]string`

GetK6Versions returns the K6Versions field if non-nil, zero value otherwise.

### GetK6VersionsOk

`func (o *StartLoadTestResponse) GetK6VersionsOk() (*map[string]string, bool)`

GetK6VersionsOk returns a tuple with the K6Versions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetK6Versions

`func (o *StartLoadTestResponse) SetK6Versions(v map[string]string)`

SetK6Versions sets K6Versions field to given value.


### GetMaxVus

`func (o *StartLoadTestResponse) GetMaxVus() int32`

GetMaxVus returns the MaxVus field if non-nil, zero value otherwise.

### GetMaxVusOk

`func (o *StartLoadTestResponse) GetMaxVusOk() (*int32, bool)`

GetMaxVusOk returns a tuple with the MaxVus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxVus

`func (o *StartLoadTestResponse) SetMaxVus(v int32)`

SetMaxVus sets MaxVus field to given value.


### SetMaxVusNil

`func (o *StartLoadTestResponse) SetMaxVusNil(b bool)`

 SetMaxVusNil sets the value for MaxVus to be an explicit nil

### UnsetMaxVus
`func (o *StartLoadTestResponse) UnsetMaxVus()`

UnsetMaxVus ensures that no value is present for MaxVus, not even an explicit nil
### GetMaxBrowserVus

`func (o *StartLoadTestResponse) GetMaxBrowserVus() int32`

GetMaxBrowserVus returns the MaxBrowserVus field if non-nil, zero value otherwise.

### GetMaxBrowserVusOk

`func (o *StartLoadTestResponse) GetMaxBrowserVusOk() (*int32, bool)`

GetMaxBrowserVusOk returns a tuple with the MaxBrowserVus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxBrowserVus

`func (o *StartLoadTestResponse) SetMaxBrowserVus(v int32)`

SetMaxBrowserVus sets MaxBrowserVus field to given value.


### SetMaxBrowserVusNil

`func (o *StartLoadTestResponse) SetMaxBrowserVusNil(b bool)`

 SetMaxBrowserVusNil sets the value for MaxBrowserVus to be an explicit nil

### UnsetMaxBrowserVus
`func (o *StartLoadTestResponse) UnsetMaxBrowserVus()`

UnsetMaxBrowserVus ensures that no value is present for MaxBrowserVus, not even an explicit nil
### GetEstimatedDuration

`func (o *StartLoadTestResponse) GetEstimatedDuration() int32`

GetEstimatedDuration returns the EstimatedDuration field if non-nil, zero value otherwise.

### GetEstimatedDurationOk

`func (o *StartLoadTestResponse) GetEstimatedDurationOk() (*int32, bool)`

GetEstimatedDurationOk returns a tuple with the EstimatedDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedDuration

`func (o *StartLoadTestResponse) SetEstimatedDuration(v int32)`

SetEstimatedDuration sets EstimatedDuration field to given value.


### SetEstimatedDurationNil

`func (o *StartLoadTestResponse) SetEstimatedDurationNil(b bool)`

 SetEstimatedDurationNil sets the value for EstimatedDuration to be an explicit nil

### UnsetEstimatedDuration
`func (o *StartLoadTestResponse) UnsetEstimatedDuration()`

UnsetEstimatedDuration ensures that no value is present for EstimatedDuration, not even an explicit nil
### GetExecutionDuration

`func (o *StartLoadTestResponse) GetExecutionDuration() int32`

GetExecutionDuration returns the ExecutionDuration field if non-nil, zero value otherwise.

### GetExecutionDurationOk

`func (o *StartLoadTestResponse) GetExecutionDurationOk() (*int32, bool)`

GetExecutionDurationOk returns a tuple with the ExecutionDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionDuration

`func (o *StartLoadTestResponse) SetExecutionDuration(v int32)`

SetExecutionDuration sets ExecutionDuration field to given value.


### GetTestRunDetailsPageUrl

`func (o *StartLoadTestResponse) GetTestRunDetailsPageUrl() string`

GetTestRunDetailsPageUrl returns the TestRunDetailsPageUrl field if non-nil, zero value otherwise.

### GetTestRunDetailsPageUrlOk

`func (o *StartLoadTestResponse) GetTestRunDetailsPageUrlOk() (*string, bool)`

GetTestRunDetailsPageUrlOk returns a tuple with the TestRunDetailsPageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestRunDetailsPageUrl

`func (o *StartLoadTestResponse) SetTestRunDetailsPageUrl(v string)`

SetTestRunDetailsPageUrl sets TestRunDetailsPageUrl field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


