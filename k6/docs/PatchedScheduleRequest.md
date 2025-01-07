# PatchedScheduleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Starts** | Pointer to **NullableTime** | The date after which the schedule will start running the test. | [optional] 
**Repeats** | Pointer to [**NullableScheduleRepeatsRequest**](ScheduleRepeatsRequest.md) |  | [optional] 
**Ends** | Pointer to [**NullableScheduleEndsRequest**](ScheduleEndsRequest.md) |  | [optional] 

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

### SetStartsNil

`func (o *PatchedScheduleRequest) SetStartsNil(b bool)`

 SetStartsNil sets the value for Starts to be an explicit nil

### UnsetStarts
`func (o *PatchedScheduleRequest) UnsetStarts()`

UnsetStarts ensures that no value is present for Starts, not even an explicit nil
### GetRepeats

`func (o *PatchedScheduleRequest) GetRepeats() ScheduleRepeatsRequest`

GetRepeats returns the Repeats field if non-nil, zero value otherwise.

### GetRepeatsOk

`func (o *PatchedScheduleRequest) GetRepeatsOk() (*ScheduleRepeatsRequest, bool)`

GetRepeatsOk returns a tuple with the Repeats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepeats

`func (o *PatchedScheduleRequest) SetRepeats(v ScheduleRepeatsRequest)`

SetRepeats sets Repeats field to given value.

### HasRepeats

`func (o *PatchedScheduleRequest) HasRepeats() bool`

HasRepeats returns a boolean if a field has been set.

### SetRepeatsNil

`func (o *PatchedScheduleRequest) SetRepeatsNil(b bool)`

 SetRepeatsNil sets the value for Repeats to be an explicit nil

### UnsetRepeats
`func (o *PatchedScheduleRequest) UnsetRepeats()`

UnsetRepeats ensures that no value is present for Repeats, not even an explicit nil
### GetEnds

`func (o *PatchedScheduleRequest) GetEnds() ScheduleEndsRequest`

GetEnds returns the Ends field if non-nil, zero value otherwise.

### GetEndsOk

`func (o *PatchedScheduleRequest) GetEndsOk() (*ScheduleEndsRequest, bool)`

GetEndsOk returns a tuple with the Ends field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnds

`func (o *PatchedScheduleRequest) SetEnds(v ScheduleEndsRequest)`

SetEnds sets Ends field to given value.

### HasEnds

`func (o *PatchedScheduleRequest) HasEnds() bool`

HasEnds returns a boolean if a field has been set.

### SetEndsNil

`func (o *PatchedScheduleRequest) SetEndsNil(b bool)`

 SetEndsNil sets the value for Ends to be an explicit nil

### UnsetEnds
`func (o *PatchedScheduleRequest) UnsetEnds()`

UnsetEnds ensures that no value is present for Ends, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


