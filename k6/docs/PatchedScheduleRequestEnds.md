# PatchedScheduleRequestEnds

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OnDate** | Pointer to **time.Time** | A datetime instance specifying the upper-bound limit of the recurrence. | [optional] 
**AfterRuns** | Pointer to **int32** | Determines how many times the schedule will start the test. | [optional] 

## Methods

### NewPatchedScheduleRequestEnds

`func NewPatchedScheduleRequestEnds() *PatchedScheduleRequestEnds`

NewPatchedScheduleRequestEnds instantiates a new PatchedScheduleRequestEnds object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedScheduleRequestEndsWithDefaults

`func NewPatchedScheduleRequestEndsWithDefaults() *PatchedScheduleRequestEnds`

NewPatchedScheduleRequestEndsWithDefaults instantiates a new PatchedScheduleRequestEnds object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOnDate

`func (o *PatchedScheduleRequestEnds) GetOnDate() time.Time`

GetOnDate returns the OnDate field if non-nil, zero value otherwise.

### GetOnDateOk

`func (o *PatchedScheduleRequestEnds) GetOnDateOk() (*time.Time, bool)`

GetOnDateOk returns a tuple with the OnDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnDate

`func (o *PatchedScheduleRequestEnds) SetOnDate(v time.Time)`

SetOnDate sets OnDate field to given value.

### HasOnDate

`func (o *PatchedScheduleRequestEnds) HasOnDate() bool`

HasOnDate returns a boolean if a field has been set.

### GetAfterRuns

`func (o *PatchedScheduleRequestEnds) GetAfterRuns() int32`

GetAfterRuns returns the AfterRuns field if non-nil, zero value otherwise.

### GetAfterRunsOk

`func (o *PatchedScheduleRequestEnds) GetAfterRunsOk() (*int32, bool)`

GetAfterRunsOk returns a tuple with the AfterRuns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfterRuns

`func (o *PatchedScheduleRequestEnds) SetAfterRuns(v int32)`

SetAfterRuns sets AfterRuns field to given value.

### HasAfterRuns

`func (o *PatchedScheduleRequestEnds) HasAfterRuns() bool`

HasAfterRuns returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


