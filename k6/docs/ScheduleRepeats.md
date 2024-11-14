# ScheduleRepeats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Frequency** | **string** | The unit type of the recurrence interval.  * &#x60;hourly&#x60; - hourly * &#x60;daily&#x60; - daily * &#x60;weekly&#x60; - weekly * &#x60;monthly&#x60; - monthly * &#x60;yearly&#x60; - yearly | 
**Interval** | Pointer to **int32** | The interval between each freq iteration. An interval of 2 with &#39;hourly&#39; frequency means once every 2 hours. | [optional] [default to 1]
**Days** | Pointer to **[]string** | When given, define the days of the week where the recurrence will be applied. The default is every day of the week. | [optional] 

## Methods

### NewScheduleRepeats

`func NewScheduleRepeats(frequency string, ) *ScheduleRepeats`

NewScheduleRepeats instantiates a new ScheduleRepeats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScheduleRepeatsWithDefaults

`func NewScheduleRepeatsWithDefaults() *ScheduleRepeats`

NewScheduleRepeatsWithDefaults instantiates a new ScheduleRepeats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFrequency

`func (o *ScheduleRepeats) GetFrequency() string`

GetFrequency returns the Frequency field if non-nil, zero value otherwise.

### GetFrequencyOk

`func (o *ScheduleRepeats) GetFrequencyOk() (*string, bool)`

GetFrequencyOk returns a tuple with the Frequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequency

`func (o *ScheduleRepeats) SetFrequency(v string)`

SetFrequency sets Frequency field to given value.


### GetInterval

`func (o *ScheduleRepeats) GetInterval() int32`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *ScheduleRepeats) GetIntervalOk() (*int32, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *ScheduleRepeats) SetInterval(v int32)`

SetInterval sets Interval field to given value.

### HasInterval

`func (o *ScheduleRepeats) HasInterval() bool`

HasInterval returns a boolean if a field has been set.

### GetDays

`func (o *ScheduleRepeats) GetDays() []string`

GetDays returns the Days field if non-nil, zero value otherwise.

### GetDaysOk

`func (o *ScheduleRepeats) GetDaysOk() (*[]string, bool)`

GetDaysOk returns a tuple with the Days field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDays

`func (o *ScheduleRepeats) SetDays(v []string)`

SetDays sets Days field to given value.

### HasDays

`func (o *ScheduleRepeats) HasDays() bool`

HasDays returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


