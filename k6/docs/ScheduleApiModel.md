# ScheduleApiModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | ID of the schedule. | 
**LoadTestId** | **int32** | ID of the test to run. | 
**Starts** | **time.Time** | The date on which the schedule will start running the test. | 
**RecurrenceRule** | [**NullableScheduleRecurrenceRule**](ScheduleRecurrenceRule.md) |  | 
**Cron** | [**NullableScheduleCron**](ScheduleCron.md) |  | 
**Deactivated** | **bool** | Whether the schedule is deactivated. A deactivated schedule will not trigger new test runs, but the schedule recurrence rule and expiration is not affected. | 
**NextRun** | **NullableTime** | The date of the next scheduled test run. The value is &#x60;null&#x60; if the schedule is expired and no more occurrences are expected in the future according to the recurrence rule. | 
**CreatedBy** | **NullableString** | The email of the user who created the schedule if applicable. | 

## Methods

### NewScheduleApiModel

`func NewScheduleApiModel(id int32, loadTestId int32, starts time.Time, recurrenceRule NullableScheduleRecurrenceRule, cron NullableScheduleCron, deactivated bool, nextRun NullableTime, createdBy NullableString, ) *ScheduleApiModel`

NewScheduleApiModel instantiates a new ScheduleApiModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScheduleApiModelWithDefaults

`func NewScheduleApiModelWithDefaults() *ScheduleApiModel`

NewScheduleApiModelWithDefaults instantiates a new ScheduleApiModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ScheduleApiModel) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ScheduleApiModel) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ScheduleApiModel) SetId(v int32)`

SetId sets Id field to given value.


### GetLoadTestId

`func (o *ScheduleApiModel) GetLoadTestId() int32`

GetLoadTestId returns the LoadTestId field if non-nil, zero value otherwise.

### GetLoadTestIdOk

`func (o *ScheduleApiModel) GetLoadTestIdOk() (*int32, bool)`

GetLoadTestIdOk returns a tuple with the LoadTestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadTestId

`func (o *ScheduleApiModel) SetLoadTestId(v int32)`

SetLoadTestId sets LoadTestId field to given value.


### GetStarts

`func (o *ScheduleApiModel) GetStarts() time.Time`

GetStarts returns the Starts field if non-nil, zero value otherwise.

### GetStartsOk

`func (o *ScheduleApiModel) GetStartsOk() (*time.Time, bool)`

GetStartsOk returns a tuple with the Starts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStarts

`func (o *ScheduleApiModel) SetStarts(v time.Time)`

SetStarts sets Starts field to given value.


### GetRecurrenceRule

`func (o *ScheduleApiModel) GetRecurrenceRule() ScheduleRecurrenceRule`

GetRecurrenceRule returns the RecurrenceRule field if non-nil, zero value otherwise.

### GetRecurrenceRuleOk

`func (o *ScheduleApiModel) GetRecurrenceRuleOk() (*ScheduleRecurrenceRule, bool)`

GetRecurrenceRuleOk returns a tuple with the RecurrenceRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurrenceRule

`func (o *ScheduleApiModel) SetRecurrenceRule(v ScheduleRecurrenceRule)`

SetRecurrenceRule sets RecurrenceRule field to given value.


### SetRecurrenceRuleNil

`func (o *ScheduleApiModel) SetRecurrenceRuleNil(b bool)`

 SetRecurrenceRuleNil sets the value for RecurrenceRule to be an explicit nil

### UnsetRecurrenceRule
`func (o *ScheduleApiModel) UnsetRecurrenceRule()`

UnsetRecurrenceRule ensures that no value is present for RecurrenceRule, not even an explicit nil
### GetCron

`func (o *ScheduleApiModel) GetCron() ScheduleCron`

GetCron returns the Cron field if non-nil, zero value otherwise.

### GetCronOk

`func (o *ScheduleApiModel) GetCronOk() (*ScheduleCron, bool)`

GetCronOk returns a tuple with the Cron field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCron

`func (o *ScheduleApiModel) SetCron(v ScheduleCron)`

SetCron sets Cron field to given value.


### SetCronNil

`func (o *ScheduleApiModel) SetCronNil(b bool)`

 SetCronNil sets the value for Cron to be an explicit nil

### UnsetCron
`func (o *ScheduleApiModel) UnsetCron()`

UnsetCron ensures that no value is present for Cron, not even an explicit nil
### GetDeactivated

`func (o *ScheduleApiModel) GetDeactivated() bool`

GetDeactivated returns the Deactivated field if non-nil, zero value otherwise.

### GetDeactivatedOk

`func (o *ScheduleApiModel) GetDeactivatedOk() (*bool, bool)`

GetDeactivatedOk returns a tuple with the Deactivated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeactivated

`func (o *ScheduleApiModel) SetDeactivated(v bool)`

SetDeactivated sets Deactivated field to given value.


### GetNextRun

`func (o *ScheduleApiModel) GetNextRun() time.Time`

GetNextRun returns the NextRun field if non-nil, zero value otherwise.

### GetNextRunOk

`func (o *ScheduleApiModel) GetNextRunOk() (*time.Time, bool)`

GetNextRunOk returns a tuple with the NextRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextRun

`func (o *ScheduleApiModel) SetNextRun(v time.Time)`

SetNextRun sets NextRun field to given value.


### SetNextRunNil

`func (o *ScheduleApiModel) SetNextRunNil(b bool)`

 SetNextRunNil sets the value for NextRun to be an explicit nil

### UnsetNextRun
`func (o *ScheduleApiModel) UnsetNextRun()`

UnsetNextRun ensures that no value is present for NextRun, not even an explicit nil
### GetCreatedBy

`func (o *ScheduleApiModel) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *ScheduleApiModel) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *ScheduleApiModel) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### SetCreatedByNil

`func (o *ScheduleApiModel) SetCreatedByNil(b bool)`

 SetCreatedByNil sets the value for CreatedBy to be an explicit nil

### UnsetCreatedBy
`func (o *ScheduleApiModel) UnsetCreatedBy()`

UnsetCreatedBy ensures that no value is present for CreatedBy, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


