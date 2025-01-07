# ScheduleCreateRequestRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LoadTestId** | **int32** | Id of the test to run. | 
**Starts** | **NullableTime** | The date after which the schedule will start running the test. | 
**Repeats** | [**NullableScheduleRepeatsRequest**](ScheduleRepeatsRequest.md) |  | 
**Ends** | [**NullableScheduleEndsRequest**](ScheduleEndsRequest.md) |  | 

## Methods

### NewScheduleCreateRequestRequest

`func NewScheduleCreateRequestRequest(loadTestId int32, starts NullableTime, repeats NullableScheduleRepeatsRequest, ends NullableScheduleEndsRequest, ) *ScheduleCreateRequestRequest`

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


### SetStartsNil

`func (o *ScheduleCreateRequestRequest) SetStartsNil(b bool)`

 SetStartsNil sets the value for Starts to be an explicit nil

### UnsetStarts
`func (o *ScheduleCreateRequestRequest) UnsetStarts()`

UnsetStarts ensures that no value is present for Starts, not even an explicit nil
### GetRepeats

`func (o *ScheduleCreateRequestRequest) GetRepeats() ScheduleRepeatsRequest`

GetRepeats returns the Repeats field if non-nil, zero value otherwise.

### GetRepeatsOk

`func (o *ScheduleCreateRequestRequest) GetRepeatsOk() (*ScheduleRepeatsRequest, bool)`

GetRepeatsOk returns a tuple with the Repeats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepeats

`func (o *ScheduleCreateRequestRequest) SetRepeats(v ScheduleRepeatsRequest)`

SetRepeats sets Repeats field to given value.


### SetRepeatsNil

`func (o *ScheduleCreateRequestRequest) SetRepeatsNil(b bool)`

 SetRepeatsNil sets the value for Repeats to be an explicit nil

### UnsetRepeats
`func (o *ScheduleCreateRequestRequest) UnsetRepeats()`

UnsetRepeats ensures that no value is present for Repeats, not even an explicit nil
### GetEnds

`func (o *ScheduleCreateRequestRequest) GetEnds() ScheduleEndsRequest`

GetEnds returns the Ends field if non-nil, zero value otherwise.

### GetEndsOk

`func (o *ScheduleCreateRequestRequest) GetEndsOk() (*ScheduleEndsRequest, bool)`

GetEndsOk returns a tuple with the Ends field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnds

`func (o *ScheduleCreateRequestRequest) SetEnds(v ScheduleEndsRequest)`

SetEnds sets Ends field to given value.


### SetEndsNil

`func (o *ScheduleCreateRequestRequest) SetEndsNil(b bool)`

 SetEndsNil sets the value for Ends to be an explicit nil

### UnsetEnds
`func (o *ScheduleCreateRequestRequest) UnsetEnds()`

UnsetEnds ensures that no value is present for Ends, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


