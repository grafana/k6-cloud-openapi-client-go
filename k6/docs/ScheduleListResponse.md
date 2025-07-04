# ScheduleListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | [**[]ScheduleApiModel**](ScheduleApiModel.md) | List of the resulting values. | 
**Count** | Pointer to **int32** | Object count in the collection. | [optional] 
**NextLink** | Pointer to **string** | A reference to the next page of results. The property is included until there are no more pages of results to retrieve. | [optional] 

## Methods

### NewScheduleListResponse

`func NewScheduleListResponse(value []ScheduleApiModel, ) *ScheduleListResponse`

NewScheduleListResponse instantiates a new ScheduleListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScheduleListResponseWithDefaults

`func NewScheduleListResponseWithDefaults() *ScheduleListResponse`

NewScheduleListResponseWithDefaults instantiates a new ScheduleListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *ScheduleListResponse) GetValue() []ScheduleApiModel`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ScheduleListResponse) GetValueOk() (*[]ScheduleApiModel, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ScheduleListResponse) SetValue(v []ScheduleApiModel)`

SetValue sets Value field to given value.


### GetCount

`func (o *ScheduleListResponse) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *ScheduleListResponse) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *ScheduleListResponse) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *ScheduleListResponse) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetNextLink

`func (o *ScheduleListResponse) GetNextLink() string`

GetNextLink returns the NextLink field if non-nil, zero value otherwise.

### GetNextLinkOk

`func (o *ScheduleListResponse) GetNextLinkOk() (*string, bool)`

GetNextLinkOk returns a tuple with the NextLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextLink

`func (o *ScheduleListResponse) SetNextLink(v string)`

SetNextLink sets NextLink field to given value.

### HasNextLink

`func (o *ScheduleListResponse) HasNextLink() bool`

HasNextLink returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


