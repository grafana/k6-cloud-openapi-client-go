# LoadTestApiModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | ID of the load test. | 
**ProjectId** | **int64** | ID of the parent project. | 
**Name** | **string** | Unique name of the test within the project. | 
**BaselineTestRunId** | **NullableInt64** | ID of a baseline test run used for results comparison. Deprecated: baselines are being replaced by the star/unstar test run APIs and this field is scheduled for removal on 2026-09-01. | 
**K6Version** | Pointer to **NullableInt32** | Identifier of the k6 version used to run the test. | [optional] 
**Created** | **time.Time** | The date when the test was created. | 
**Updated** | **time.Time** | The date when the test was last updated. | 

## Methods

### NewLoadTestApiModel

`func NewLoadTestApiModel(id int64, projectId int64, name string, baselineTestRunId NullableInt64, created time.Time, updated time.Time, ) *LoadTestApiModel`

NewLoadTestApiModel instantiates a new LoadTestApiModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoadTestApiModelWithDefaults

`func NewLoadTestApiModelWithDefaults() *LoadTestApiModel`

NewLoadTestApiModelWithDefaults instantiates a new LoadTestApiModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *LoadTestApiModel) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LoadTestApiModel) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LoadTestApiModel) SetId(v int64)`

SetId sets Id field to given value.


### GetProjectId

`func (o *LoadTestApiModel) GetProjectId() int64`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *LoadTestApiModel) GetProjectIdOk() (*int64, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *LoadTestApiModel) SetProjectId(v int64)`

SetProjectId sets ProjectId field to given value.


### GetName

`func (o *LoadTestApiModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LoadTestApiModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LoadTestApiModel) SetName(v string)`

SetName sets Name field to given value.


### GetBaselineTestRunId

`func (o *LoadTestApiModel) GetBaselineTestRunId() int64`

GetBaselineTestRunId returns the BaselineTestRunId field if non-nil, zero value otherwise.

### GetBaselineTestRunIdOk

`func (o *LoadTestApiModel) GetBaselineTestRunIdOk() (*int64, bool)`

GetBaselineTestRunIdOk returns a tuple with the BaselineTestRunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaselineTestRunId

`func (o *LoadTestApiModel) SetBaselineTestRunId(v int64)`

SetBaselineTestRunId sets BaselineTestRunId field to given value.


### SetBaselineTestRunIdNil

`func (o *LoadTestApiModel) SetBaselineTestRunIdNil(b bool)`

 SetBaselineTestRunIdNil sets the value for BaselineTestRunId to be an explicit nil

### UnsetBaselineTestRunId
`func (o *LoadTestApiModel) UnsetBaselineTestRunId()`

UnsetBaselineTestRunId ensures that no value is present for BaselineTestRunId, not even an explicit nil
### GetK6Version

`func (o *LoadTestApiModel) GetK6Version() int32`

GetK6Version returns the K6Version field if non-nil, zero value otherwise.

### GetK6VersionOk

`func (o *LoadTestApiModel) GetK6VersionOk() (*int32, bool)`

GetK6VersionOk returns a tuple with the K6Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetK6Version

`func (o *LoadTestApiModel) SetK6Version(v int32)`

SetK6Version sets K6Version field to given value.

### HasK6Version

`func (o *LoadTestApiModel) HasK6Version() bool`

HasK6Version returns a boolean if a field has been set.

### SetK6VersionNil

`func (o *LoadTestApiModel) SetK6VersionNil(b bool)`

 SetK6VersionNil sets the value for K6Version to be an explicit nil

### UnsetK6Version
`func (o *LoadTestApiModel) UnsetK6Version()`

UnsetK6Version ensures that no value is present for K6Version, not even an explicit nil
### GetCreated

`func (o *LoadTestApiModel) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *LoadTestApiModel) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *LoadTestApiModel) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetUpdated

`func (o *LoadTestApiModel) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *LoadTestApiModel) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *LoadTestApiModel) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


