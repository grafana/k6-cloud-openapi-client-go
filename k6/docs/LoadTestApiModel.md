# LoadTestApiModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | ID of the load test. | 
**ProjectId** | **int32** | ID of the parent project. | 
**Name** | **string** | Unique name of the test within the project. | 
**BaselineTestRunId** | **NullableInt32** | ID of a baseline test run used for results comparison. | 
**Created** | **time.Time** | The date when the test was created. | 
**Updated** | **time.Time** | The date when the test was last updated. | 

## Methods

### NewLoadTestApiModel

`func NewLoadTestApiModel(id int32, projectId int32, name string, baselineTestRunId NullableInt32, created time.Time, updated time.Time, ) *LoadTestApiModel`

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

`func (o *LoadTestApiModel) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LoadTestApiModel) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LoadTestApiModel) SetId(v int32)`

SetId sets Id field to given value.


### GetProjectId

`func (o *LoadTestApiModel) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *LoadTestApiModel) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *LoadTestApiModel) SetProjectId(v int32)`

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

`func (o *LoadTestApiModel) GetBaselineTestRunId() int32`

GetBaselineTestRunId returns the BaselineTestRunId field if non-nil, zero value otherwise.

### GetBaselineTestRunIdOk

`func (o *LoadTestApiModel) GetBaselineTestRunIdOk() (*int32, bool)`

GetBaselineTestRunIdOk returns a tuple with the BaselineTestRunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaselineTestRunId

`func (o *LoadTestApiModel) SetBaselineTestRunId(v int32)`

SetBaselineTestRunId sets BaselineTestRunId field to given value.


### SetBaselineTestRunIdNil

`func (o *LoadTestApiModel) SetBaselineTestRunIdNil(b bool)`

 SetBaselineTestRunIdNil sets the value for BaselineTestRunId to be an explicit nil

### UnsetBaselineTestRunId
`func (o *LoadTestApiModel) UnsetBaselineTestRunId()`

UnsetBaselineTestRunId ensures that no value is present for BaselineTestRunId, not even an explicit nil
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


