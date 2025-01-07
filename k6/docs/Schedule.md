# Schedule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Id of the schedule. | [readonly] 
**LoadTestId** | **int32** | Id of the test to run. | [readonly] 
**Active** | **bool** | Id of the schedule. | [readonly] 
**Starts** | **NullableTime** | The date after which the schedule will start running the test. | 
**Repeats** | [**NullableScheduleRepeats**](ScheduleRepeats.md) |  | 
**Ends** | [**NullableScheduleEnds**](ScheduleEnds.md) |  | 
**NextRun** | **NullableTime** | The date of the next scheduled test run. If the schedule is inactive or past the end date or all runs have been used up, the value is null. | [readonly] 

## Methods

### NewSchedule

`func NewSchedule(id int32, loadTestId int32, active bool, starts NullableTime, repeats NullableScheduleRepeats, ends NullableScheduleEnds, nextRun NullableTime, ) *Schedule`

NewSchedule instantiates a new Schedule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScheduleWithDefaults

`func NewScheduleWithDefaults() *Schedule`

NewScheduleWithDefaults instantiates a new Schedule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Schedule) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Schedule) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Schedule) SetId(v int32)`

SetId sets Id field to given value.


### GetLoadTestId

`func (o *Schedule) GetLoadTestId() int32`

GetLoadTestId returns the LoadTestId field if non-nil, zero value otherwise.

### GetLoadTestIdOk

`func (o *Schedule) GetLoadTestIdOk() (*int32, bool)`

GetLoadTestIdOk returns a tuple with the LoadTestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadTestId

`func (o *Schedule) SetLoadTestId(v int32)`

SetLoadTestId sets LoadTestId field to given value.


### GetActive

`func (o *Schedule) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *Schedule) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *Schedule) SetActive(v bool)`

SetActive sets Active field to given value.


### GetStarts

`func (o *Schedule) GetStarts() time.Time`

GetStarts returns the Starts field if non-nil, zero value otherwise.

### GetStartsOk

`func (o *Schedule) GetStartsOk() (*time.Time, bool)`

GetStartsOk returns a tuple with the Starts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStarts

`func (o *Schedule) SetStarts(v time.Time)`

SetStarts sets Starts field to given value.


### SetStartsNil

`func (o *Schedule) SetStartsNil(b bool)`

 SetStartsNil sets the value for Starts to be an explicit nil

### UnsetStarts
`func (o *Schedule) UnsetStarts()`

UnsetStarts ensures that no value is present for Starts, not even an explicit nil
### GetRepeats

`func (o *Schedule) GetRepeats() ScheduleRepeats`

GetRepeats returns the Repeats field if non-nil, zero value otherwise.

### GetRepeatsOk

`func (o *Schedule) GetRepeatsOk() (*ScheduleRepeats, bool)`

GetRepeatsOk returns a tuple with the Repeats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepeats

`func (o *Schedule) SetRepeats(v ScheduleRepeats)`

SetRepeats sets Repeats field to given value.


### SetRepeatsNil

`func (o *Schedule) SetRepeatsNil(b bool)`

 SetRepeatsNil sets the value for Repeats to be an explicit nil

### UnsetRepeats
`func (o *Schedule) UnsetRepeats()`

UnsetRepeats ensures that no value is present for Repeats, not even an explicit nil
### GetEnds

`func (o *Schedule) GetEnds() ScheduleEnds`

GetEnds returns the Ends field if non-nil, zero value otherwise.

### GetEndsOk

`func (o *Schedule) GetEndsOk() (*ScheduleEnds, bool)`

GetEndsOk returns a tuple with the Ends field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnds

`func (o *Schedule) SetEnds(v ScheduleEnds)`

SetEnds sets Ends field to given value.


### SetEndsNil

`func (o *Schedule) SetEndsNil(b bool)`

 SetEndsNil sets the value for Ends to be an explicit nil

### UnsetEnds
`func (o *Schedule) UnsetEnds()`

UnsetEnds ensures that no value is present for Ends, not even an explicit nil
### GetNextRun

`func (o *Schedule) GetNextRun() time.Time`

GetNextRun returns the NextRun field if non-nil, zero value otherwise.

### GetNextRunOk

`func (o *Schedule) GetNextRunOk() (*time.Time, bool)`

GetNextRunOk returns a tuple with the NextRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextRun

`func (o *Schedule) SetNextRun(v time.Time)`

SetNextRun sets NextRun field to given value.


### SetNextRunNil

`func (o *Schedule) SetNextRunNil(b bool)`

 SetNextRunNil sets the value for NextRun to be an explicit nil

### UnsetNextRun
`func (o *Schedule) UnsetNextRun()`

UnsetNextRun ensures that no value is present for NextRun, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


