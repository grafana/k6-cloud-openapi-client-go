# TestRunApiModel

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

## Methods

### NewTestRunApiModel

`func NewTestRunApiModel(id int32, testId int32, projectId int32, startedBy NullableString, created time.Time, ended NullableTime, note string, retentionExpiry NullableTime, cost NullableTestCostApiModel, status string, statusDetails StatusApiModel, statusHistory []StatusApiModel, distribution []DistributionZoneApiModel, result NullableString, resultDetails map[string]interface{}, options map[string]interface{}, k6Dependencies map[string]string, k6Versions map[string]string, ) *TestRunApiModel`

NewTestRunApiModel instantiates a new TestRunApiModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestRunApiModelWithDefaults

`func NewTestRunApiModelWithDefaults() *TestRunApiModel`

NewTestRunApiModelWithDefaults instantiates a new TestRunApiModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TestRunApiModel) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TestRunApiModel) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TestRunApiModel) SetId(v int32)`

SetId sets Id field to given value.


### GetTestId

`func (o *TestRunApiModel) GetTestId() int32`

GetTestId returns the TestId field if non-nil, zero value otherwise.

### GetTestIdOk

`func (o *TestRunApiModel) GetTestIdOk() (*int32, bool)`

GetTestIdOk returns a tuple with the TestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestId

`func (o *TestRunApiModel) SetTestId(v int32)`

SetTestId sets TestId field to given value.


### GetProjectId

`func (o *TestRunApiModel) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *TestRunApiModel) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *TestRunApiModel) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.


### GetStartedBy

`func (o *TestRunApiModel) GetStartedBy() string`

GetStartedBy returns the StartedBy field if non-nil, zero value otherwise.

### GetStartedByOk

`func (o *TestRunApiModel) GetStartedByOk() (*string, bool)`

GetStartedByOk returns a tuple with the StartedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedBy

`func (o *TestRunApiModel) SetStartedBy(v string)`

SetStartedBy sets StartedBy field to given value.


### SetStartedByNil

`func (o *TestRunApiModel) SetStartedByNil(b bool)`

 SetStartedByNil sets the value for StartedBy to be an explicit nil

### UnsetStartedBy
`func (o *TestRunApiModel) UnsetStartedBy()`

UnsetStartedBy ensures that no value is present for StartedBy, not even an explicit nil
### GetCreated

`func (o *TestRunApiModel) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *TestRunApiModel) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *TestRunApiModel) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetEnded

`func (o *TestRunApiModel) GetEnded() time.Time`

GetEnded returns the Ended field if non-nil, zero value otherwise.

### GetEndedOk

`func (o *TestRunApiModel) GetEndedOk() (*time.Time, bool)`

GetEndedOk returns a tuple with the Ended field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnded

`func (o *TestRunApiModel) SetEnded(v time.Time)`

SetEnded sets Ended field to given value.


### SetEndedNil

`func (o *TestRunApiModel) SetEndedNil(b bool)`

 SetEndedNil sets the value for Ended to be an explicit nil

### UnsetEnded
`func (o *TestRunApiModel) UnsetEnded()`

UnsetEnded ensures that no value is present for Ended, not even an explicit nil
### GetNote

`func (o *TestRunApiModel) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *TestRunApiModel) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *TestRunApiModel) SetNote(v string)`

SetNote sets Note field to given value.


### GetRetentionExpiry

`func (o *TestRunApiModel) GetRetentionExpiry() time.Time`

GetRetentionExpiry returns the RetentionExpiry field if non-nil, zero value otherwise.

### GetRetentionExpiryOk

`func (o *TestRunApiModel) GetRetentionExpiryOk() (*time.Time, bool)`

GetRetentionExpiryOk returns a tuple with the RetentionExpiry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionExpiry

`func (o *TestRunApiModel) SetRetentionExpiry(v time.Time)`

SetRetentionExpiry sets RetentionExpiry field to given value.


### SetRetentionExpiryNil

`func (o *TestRunApiModel) SetRetentionExpiryNil(b bool)`

 SetRetentionExpiryNil sets the value for RetentionExpiry to be an explicit nil

### UnsetRetentionExpiry
`func (o *TestRunApiModel) UnsetRetentionExpiry()`

UnsetRetentionExpiry ensures that no value is present for RetentionExpiry, not even an explicit nil
### GetCost

`func (o *TestRunApiModel) GetCost() TestCostApiModel`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *TestRunApiModel) GetCostOk() (*TestCostApiModel, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *TestRunApiModel) SetCost(v TestCostApiModel)`

SetCost sets Cost field to given value.


### SetCostNil

`func (o *TestRunApiModel) SetCostNil(b bool)`

 SetCostNil sets the value for Cost to be an explicit nil

### UnsetCost
`func (o *TestRunApiModel) UnsetCost()`

UnsetCost ensures that no value is present for Cost, not even an explicit nil
### GetStatus

`func (o *TestRunApiModel) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TestRunApiModel) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TestRunApiModel) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetStatusDetails

`func (o *TestRunApiModel) GetStatusDetails() StatusApiModel`

GetStatusDetails returns the StatusDetails field if non-nil, zero value otherwise.

### GetStatusDetailsOk

`func (o *TestRunApiModel) GetStatusDetailsOk() (*StatusApiModel, bool)`

GetStatusDetailsOk returns a tuple with the StatusDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusDetails

`func (o *TestRunApiModel) SetStatusDetails(v StatusApiModel)`

SetStatusDetails sets StatusDetails field to given value.


### GetStatusHistory

`func (o *TestRunApiModel) GetStatusHistory() []StatusApiModel`

GetStatusHistory returns the StatusHistory field if non-nil, zero value otherwise.

### GetStatusHistoryOk

`func (o *TestRunApiModel) GetStatusHistoryOk() (*[]StatusApiModel, bool)`

GetStatusHistoryOk returns a tuple with the StatusHistory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusHistory

`func (o *TestRunApiModel) SetStatusHistory(v []StatusApiModel)`

SetStatusHistory sets StatusHistory field to given value.


### GetDistribution

`func (o *TestRunApiModel) GetDistribution() []DistributionZoneApiModel`

GetDistribution returns the Distribution field if non-nil, zero value otherwise.

### GetDistributionOk

`func (o *TestRunApiModel) GetDistributionOk() (*[]DistributionZoneApiModel, bool)`

GetDistributionOk returns a tuple with the Distribution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistribution

`func (o *TestRunApiModel) SetDistribution(v []DistributionZoneApiModel)`

SetDistribution sets Distribution field to given value.


### SetDistributionNil

`func (o *TestRunApiModel) SetDistributionNil(b bool)`

 SetDistributionNil sets the value for Distribution to be an explicit nil

### UnsetDistribution
`func (o *TestRunApiModel) UnsetDistribution()`

UnsetDistribution ensures that no value is present for Distribution, not even an explicit nil
### GetResult

`func (o *TestRunApiModel) GetResult() string`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *TestRunApiModel) GetResultOk() (*string, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *TestRunApiModel) SetResult(v string)`

SetResult sets Result field to given value.


### SetResultNil

`func (o *TestRunApiModel) SetResultNil(b bool)`

 SetResultNil sets the value for Result to be an explicit nil

### UnsetResult
`func (o *TestRunApiModel) UnsetResult()`

UnsetResult ensures that no value is present for Result, not even an explicit nil
### GetResultDetails

`func (o *TestRunApiModel) GetResultDetails() map[string]interface{}`

GetResultDetails returns the ResultDetails field if non-nil, zero value otherwise.

### GetResultDetailsOk

`func (o *TestRunApiModel) GetResultDetailsOk() (*map[string]interface{}, bool)`

GetResultDetailsOk returns a tuple with the ResultDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultDetails

`func (o *TestRunApiModel) SetResultDetails(v map[string]interface{})`

SetResultDetails sets ResultDetails field to given value.


### SetResultDetailsNil

`func (o *TestRunApiModel) SetResultDetailsNil(b bool)`

 SetResultDetailsNil sets the value for ResultDetails to be an explicit nil

### UnsetResultDetails
`func (o *TestRunApiModel) UnsetResultDetails()`

UnsetResultDetails ensures that no value is present for ResultDetails, not even an explicit nil
### GetOptions

`func (o *TestRunApiModel) GetOptions() map[string]interface{}`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *TestRunApiModel) GetOptionsOk() (*map[string]interface{}, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *TestRunApiModel) SetOptions(v map[string]interface{})`

SetOptions sets Options field to given value.


### SetOptionsNil

`func (o *TestRunApiModel) SetOptionsNil(b bool)`

 SetOptionsNil sets the value for Options to be an explicit nil

### UnsetOptions
`func (o *TestRunApiModel) UnsetOptions()`

UnsetOptions ensures that no value is present for Options, not even an explicit nil
### GetK6Dependencies

`func (o *TestRunApiModel) GetK6Dependencies() map[string]string`

GetK6Dependencies returns the K6Dependencies field if non-nil, zero value otherwise.

### GetK6DependenciesOk

`func (o *TestRunApiModel) GetK6DependenciesOk() (*map[string]string, bool)`

GetK6DependenciesOk returns a tuple with the K6Dependencies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetK6Dependencies

`func (o *TestRunApiModel) SetK6Dependencies(v map[string]string)`

SetK6Dependencies sets K6Dependencies field to given value.


### GetK6Versions

`func (o *TestRunApiModel) GetK6Versions() map[string]string`

GetK6Versions returns the K6Versions field if non-nil, zero value otherwise.

### GetK6VersionsOk

`func (o *TestRunApiModel) GetK6VersionsOk() (*map[string]string, bool)`

GetK6VersionsOk returns a tuple with the K6Versions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetK6Versions

`func (o *TestRunApiModel) SetK6Versions(v map[string]string)`

SetK6Versions sets K6Versions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


