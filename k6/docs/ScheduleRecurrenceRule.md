# ScheduleRecurrenceRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Frequency** | [**Frequency**](Frequency.md) |  | 
**Interval** | Pointer to **int32** | The interval between each frequency iteration. An interval of 2 with &#39;HOURLY&#39; frequency makes the test run once every 2 hours. | [optional] [default to 1]
**Byday** | Pointer to [**[]Weekday**](Weekday.md) | The weekdays when the &#39;WEEKLY&#39; recurrence will be applied. Cannot be set for other frequencies. | [optional] 
**Until** | Pointer to **NullableTime** | A datetime instance specifying the upper-bound time limit of the recurrence. | [optional] 
**Count** | Pointer to **NullableInt32** | Determines how many times the recurrence will repeat. | [optional] 

## Methods

### NewScheduleRecurrenceRule

`func NewScheduleRecurrenceRule(frequency Frequency, ) *ScheduleRecurrenceRule`

NewScheduleRecurrenceRule instantiates a new ScheduleRecurrenceRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScheduleRecurrenceRuleWithDefaults

`func NewScheduleRecurrenceRuleWithDefaults() *ScheduleRecurrenceRule`

NewScheduleRecurrenceRuleWithDefaults instantiates a new ScheduleRecurrenceRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFrequency

`func (o *ScheduleRecurrenceRule) GetFrequency() Frequency`

GetFrequency returns the Frequency field if non-nil, zero value otherwise.

### GetFrequencyOk

`func (o *ScheduleRecurrenceRule) GetFrequencyOk() (*Frequency, bool)`

GetFrequencyOk returns a tuple with the Frequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequency

`func (o *ScheduleRecurrenceRule) SetFrequency(v Frequency)`

SetFrequency sets Frequency field to given value.


### GetInterval

`func (o *ScheduleRecurrenceRule) GetInterval() int32`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *ScheduleRecurrenceRule) GetIntervalOk() (*int32, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *ScheduleRecurrenceRule) SetInterval(v int32)`

SetInterval sets Interval field to given value.

### HasInterval

`func (o *ScheduleRecurrenceRule) HasInterval() bool`

HasInterval returns a boolean if a field has been set.

### GetByday

`func (o *ScheduleRecurrenceRule) GetByday() []Weekday`

GetByday returns the Byday field if non-nil, zero value otherwise.

### GetBydayOk

`func (o *ScheduleRecurrenceRule) GetBydayOk() (*[]Weekday, bool)`

GetBydayOk returns a tuple with the Byday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByday

`func (o *ScheduleRecurrenceRule) SetByday(v []Weekday)`

SetByday sets Byday field to given value.

### HasByday

`func (o *ScheduleRecurrenceRule) HasByday() bool`

HasByday returns a boolean if a field has been set.

### SetBydayNil

`func (o *ScheduleRecurrenceRule) SetBydayNil(b bool)`

 SetBydayNil sets the value for Byday to be an explicit nil

### UnsetByday
`func (o *ScheduleRecurrenceRule) UnsetByday()`

UnsetByday ensures that no value is present for Byday, not even an explicit nil
### GetUntil

`func (o *ScheduleRecurrenceRule) GetUntil() time.Time`

GetUntil returns the Until field if non-nil, zero value otherwise.

### GetUntilOk

`func (o *ScheduleRecurrenceRule) GetUntilOk() (*time.Time, bool)`

GetUntilOk returns a tuple with the Until field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUntil

`func (o *ScheduleRecurrenceRule) SetUntil(v time.Time)`

SetUntil sets Until field to given value.

### HasUntil

`func (o *ScheduleRecurrenceRule) HasUntil() bool`

HasUntil returns a boolean if a field has been set.

### SetUntilNil

`func (o *ScheduleRecurrenceRule) SetUntilNil(b bool)`

 SetUntilNil sets the value for Until to be an explicit nil

### UnsetUntil
`func (o *ScheduleRecurrenceRule) UnsetUntil()`

UnsetUntil ensures that no value is present for Until, not even an explicit nil
### GetCount

`func (o *ScheduleRecurrenceRule) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *ScheduleRecurrenceRule) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *ScheduleRecurrenceRule) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *ScheduleRecurrenceRule) HasCount() bool`

HasCount returns a boolean if a field has been set.

### SetCountNil

`func (o *ScheduleRecurrenceRule) SetCountNil(b bool)`

 SetCountNil sets the value for Count to be an explicit nil

### UnsetCount
`func (o *ScheduleRecurrenceRule) UnsetCount()`

UnsetCount ensures that no value is present for Count, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


