# CreateScheduleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Starts** | **time.Time** | The timezone-aware date on which the schedule will start running the test. | 
**RecurrenceRule** | Pointer to [**NullableScheduleRecurrenceRule**](ScheduleRecurrenceRule.md) |  | [optional] 
**Cron** | Pointer to [**NullableScheduleCron**](ScheduleCron.md) |  | [optional] 

## Methods

### NewCreateScheduleRequest

`func NewCreateScheduleRequest(starts time.Time, ) *CreateScheduleRequest`

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

### HasRecurrenceRule

`func (o *CreateScheduleRequest) HasRecurrenceRule() bool`

HasRecurrenceRule returns a boolean if a field has been set.

### SetRecurrenceRuleNil

`func (o *CreateScheduleRequest) SetRecurrenceRuleNil(b bool)`

 SetRecurrenceRuleNil sets the value for RecurrenceRule to be an explicit nil

### UnsetRecurrenceRule
`func (o *CreateScheduleRequest) UnsetRecurrenceRule()`

UnsetRecurrenceRule ensures that no value is present for RecurrenceRule, not even an explicit nil
### GetCron

`func (o *CreateScheduleRequest) GetCron() ScheduleCron`

GetCron returns the Cron field if non-nil, zero value otherwise.

### GetCronOk

`func (o *CreateScheduleRequest) GetCronOk() (*ScheduleCron, bool)`

GetCronOk returns a tuple with the Cron field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCron

`func (o *CreateScheduleRequest) SetCron(v ScheduleCron)`

SetCron sets Cron field to given value.

### HasCron

`func (o *CreateScheduleRequest) HasCron() bool`

HasCron returns a boolean if a field has been set.

### SetCronNil

`func (o *CreateScheduleRequest) SetCronNil(b bool)`

 SetCronNil sets the value for Cron to be an explicit nil

### UnsetCron
`func (o *CreateScheduleRequest) UnsetCron()`

UnsetCron ensures that no value is present for Cron, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


