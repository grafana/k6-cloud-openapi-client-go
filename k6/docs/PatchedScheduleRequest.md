# PatchedScheduleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Starts** | Pointer to **time.Time** | The date after which the schedule will start running the test. | [optional] 
**Repeats** | Pointer to [**PatchedScheduleRequestRepeats**](PatchedScheduleRequestRepeats.md) |  | [optional] 
**Ends** | Pointer to [**PatchedScheduleRequestEnds**](PatchedScheduleRequestEnds.md) |  | [optional] 

## Methods

### NewPatchedScheduleRequest

`func NewPatchedScheduleRequest() *PatchedScheduleRequest`

NewPatchedScheduleRequest instantiates a new PatchedScheduleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedScheduleRequestWithDefaults

`func NewPatchedScheduleRequestWithDefaults() *PatchedScheduleRequest`

NewPatchedScheduleRequestWithDefaults instantiates a new PatchedScheduleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStarts

`func (o *PatchedScheduleRequest) GetStarts() time.Time`

GetStarts returns the Starts field if non-nil, zero value otherwise.

### GetStartsOk

`func (o *PatchedScheduleRequest) GetStartsOk() (*time.Time, bool)`

GetStartsOk returns a tuple with the Starts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStarts

`func (o *PatchedScheduleRequest) SetStarts(v time.Time)`

SetStarts sets Starts field to given value.

### HasStarts

`func (o *PatchedScheduleRequest) HasStarts() bool`

HasStarts returns a boolean if a field has been set.

### GetRepeats

`func (o *PatchedScheduleRequest) GetRepeats() PatchedScheduleRequestRepeats`

GetRepeats returns the Repeats field if non-nil, zero value otherwise.

### GetRepeatsOk

`func (o *PatchedScheduleRequest) GetRepeatsOk() (*PatchedScheduleRequestRepeats, bool)`

GetRepeatsOk returns a tuple with the Repeats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepeats

`func (o *PatchedScheduleRequest) SetRepeats(v PatchedScheduleRequestRepeats)`

SetRepeats sets Repeats field to given value.

### HasRepeats

`func (o *PatchedScheduleRequest) HasRepeats() bool`

HasRepeats returns a boolean if a field has been set.

### GetEnds

`func (o *PatchedScheduleRequest) GetEnds() PatchedScheduleRequestEnds`

GetEnds returns the Ends field if non-nil, zero value otherwise.

### GetEndsOk

`func (o *PatchedScheduleRequest) GetEndsOk() (*PatchedScheduleRequestEnds, bool)`

GetEndsOk returns a tuple with the Ends field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnds

`func (o *PatchedScheduleRequest) SetEnds(v PatchedScheduleRequestEnds)`

SetEnds sets Ends field to given value.

### HasEnds

`func (o *PatchedScheduleRequest) HasEnds() bool`

HasEnds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


