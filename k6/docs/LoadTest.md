# LoadTest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Id of the load test | [readonly] 
**ProjectId** | **int32** | Id of the parent project | [readonly] 
**Name** | **string** | Unique name of the test within the project | 
**BaselineTestRunId** | **int32** | Id of a baseline test run used for results comparison | 
**Created** | **time.Time** | The date when the test was created | [readonly] 
**Updated** | **time.Time** | The date when the test was last time updated | [readonly] 

## Methods

### NewLoadTest

`func NewLoadTest(id int32, projectId int32, name string, baselineTestRunId int32, created time.Time, updated time.Time, ) *LoadTest`

NewLoadTest instantiates a new LoadTest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoadTestWithDefaults

`func NewLoadTestWithDefaults() *LoadTest`

NewLoadTestWithDefaults instantiates a new LoadTest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *LoadTest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LoadTest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LoadTest) SetId(v int32)`

SetId sets Id field to given value.


### GetProjectId

`func (o *LoadTest) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *LoadTest) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *LoadTest) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.


### GetName

`func (o *LoadTest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LoadTest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LoadTest) SetName(v string)`

SetName sets Name field to given value.


### GetBaselineTestRunId

`func (o *LoadTest) GetBaselineTestRunId() int32`

GetBaselineTestRunId returns the BaselineTestRunId field if non-nil, zero value otherwise.

### GetBaselineTestRunIdOk

`func (o *LoadTest) GetBaselineTestRunIdOk() (*int32, bool)`

GetBaselineTestRunIdOk returns a tuple with the BaselineTestRunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaselineTestRunId

`func (o *LoadTest) SetBaselineTestRunId(v int32)`

SetBaselineTestRunId sets BaselineTestRunId field to given value.


### GetCreated

`func (o *LoadTest) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *LoadTest) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *LoadTest) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetUpdated

`func (o *LoadTest) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *LoadTest) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *LoadTest) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


