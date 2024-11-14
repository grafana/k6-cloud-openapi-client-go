# ScheduleCreateRequestRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LoadTestId** | **int32** | Id of the test to run. | 
**Starts** | **time.Time** | The date after which the schedule will start running the test. | 
**Repeats** | [**PatchedScheduleRequestRepeats**](PatchedScheduleRequestRepeats.md) |  | 
**Ends** | [**PatchedScheduleRequestEnds**](PatchedScheduleRequestEnds.md) |  | 

## Methods

### NewScheduleCreateRequestRequest

`func NewScheduleCreateRequestRequest(loadTestId int32, starts time.Time, repeats PatchedScheduleRequestRepeats, ends PatchedScheduleRequestEnds, ) *ScheduleCreateRequestRequest`

NewScheduleCreateRequestRequest instantiates a new ScheduleCreateRequestRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScheduleCreateRequestRequestWithDefaults

`func NewScheduleCreateRequestRequestWithDefaults() *ScheduleCreateRequestRequest`

NewScheduleCreateRequestRequestWithDefaults instantiates a new ScheduleCreateRequestRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLoadTestId

`func (o *ScheduleCreateRequestRequest) GetLoadTestId() int32`

GetLoadTestId returns the LoadTestId field if non-nil, zero value otherwise.

### GetLoadTestIdOk

`func (o *ScheduleCreateRequestRequest) GetLoadTestIdOk() (*int32, bool)`

GetLoadTestIdOk returns a tuple with the LoadTestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadTestId

`func (o *ScheduleCreateRequestRequest) SetLoadTestId(v int32)`

SetLoadTestId sets LoadTestId field to given value.


### GetStarts

`func (o *ScheduleCreateRequestRequest) GetStarts() time.Time`

GetStarts returns the Starts field if non-nil, zero value otherwise.

### GetStartsOk

`func (o *ScheduleCreateRequestRequest) GetStartsOk() (*time.Time, bool)`

GetStartsOk returns a tuple with the Starts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStarts

`func (o *ScheduleCreateRequestRequest) SetStarts(v time.Time)`

SetStarts sets Starts field to given value.


### GetRepeats

`func (o *ScheduleCreateRequestRequest) GetRepeats() PatchedScheduleRequestRepeats`

GetRepeats returns the Repeats field if non-nil, zero value otherwise.

### GetRepeatsOk

`func (o *ScheduleCreateRequestRequest) GetRepeatsOk() (*PatchedScheduleRequestRepeats, bool)`

GetRepeatsOk returns a tuple with the Repeats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepeats

`func (o *ScheduleCreateRequestRequest) SetRepeats(v PatchedScheduleRequestRepeats)`

SetRepeats sets Repeats field to given value.


### GetEnds

`func (o *ScheduleCreateRequestRequest) GetEnds() PatchedScheduleRequestEnds`

GetEnds returns the Ends field if non-nil, zero value otherwise.

### GetEndsOk

`func (o *ScheduleCreateRequestRequest) GetEndsOk() (*PatchedScheduleRequestEnds, bool)`

GetEndsOk returns a tuple with the Ends field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnds

`func (o *ScheduleCreateRequestRequest) SetEnds(v PatchedScheduleRequestEnds)`

SetEnds sets Ends field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


