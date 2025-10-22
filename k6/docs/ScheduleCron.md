# ScheduleCron

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schedule** | **string** | A cron expression with exactly 5 entries, or an alias. The allowed aliases are: @yearly, @annually, @monthly, @weekly, @daily, @hourly. | 
**TimeZone** | **string** | The timezone of the cron expression. | 

## Methods

### NewScheduleCron

`func NewScheduleCron(schedule string, timeZone string, ) *ScheduleCron`

NewScheduleCron instantiates a new ScheduleCron object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScheduleCronWithDefaults

`func NewScheduleCronWithDefaults() *ScheduleCron`

NewScheduleCronWithDefaults instantiates a new ScheduleCron object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchedule

`func (o *ScheduleCron) GetSchedule() string`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *ScheduleCron) GetScheduleOk() (*string, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *ScheduleCron) SetSchedule(v string)`

SetSchedule sets Schedule field to given value.


### GetTimeZone

`func (o *ScheduleCron) GetTimeZone() string`

GetTimeZone returns the TimeZone field if non-nil, zero value otherwise.

### GetTimeZoneOk

`func (o *ScheduleCron) GetTimeZoneOk() (*string, bool)`

GetTimeZoneOk returns a tuple with the TimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZone

`func (o *ScheduleCron) SetTimeZone(v string)`

SetTimeZone sets TimeZone field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


