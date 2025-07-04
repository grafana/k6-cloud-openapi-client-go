# CreateScheduleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Starts** | **time.Time** | The timezone-aware date on which the schedule will start running the test. | 
**RecurrenceRule** | [**NullableScheduleRecurrenceRule**](ScheduleRecurrenceRule.md) |  | 

## Methods

### NewCreateScheduleRequest

`func NewCreateScheduleRequest(starts time.Time, recurrenceRule NullableScheduleRecurrenceRule, ) *CreateScheduleRequest`

NewCreateScheduleRequest instantiates a new CreateScheduleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateScheduleRequestWithDefaults

`func NewCreateScheduleRequestWithDefaults() *CreateScheduleRequest`

NewCreateScheduleRequestWithDefaults instantiates a new CreateScheduleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStarts

`func (o *CreateScheduleRequest) GetStarts() time.Time`

GetStarts returns the Starts field if non-nil, zero value otherwise.

### GetStartsOk

`func (o *CreateScheduleRequest) GetStartsOk() (*time.Time, bool)`

GetStartsOk returns a tuple with the Starts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStarts

`func (o *CreateScheduleRequest) SetStarts(v time.Time)`

SetStarts sets Starts field to given value.


### GetRecurrenceRule

`func (o *CreateScheduleRequest) GetRecurrenceRule() ScheduleRecurrenceRule`

GetRecurrenceRule returns the RecurrenceRule field if non-nil, zero value otherwise.

### GetRecurrenceRuleOk

`func (o *CreateScheduleRequest) GetRecurrenceRuleOk() (*ScheduleRecurrenceRule, bool)`

GetRecurrenceRuleOk returns a tuple with the RecurrenceRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurrenceRule

`func (o *CreateScheduleRequest) SetRecurrenceRule(v ScheduleRecurrenceRule)`

SetRecurrenceRule sets RecurrenceRule field to given value.


### SetRecurrenceRuleNil

`func (o *CreateScheduleRequest) SetRecurrenceRuleNil(b bool)`

 SetRecurrenceRuleNil sets the value for RecurrenceRule to be an explicit nil

### UnsetRecurrenceRule
`func (o *CreateScheduleRequest) UnsetRecurrenceRule()`

UnsetRecurrenceRule ensures that no value is present for RecurrenceRule, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


