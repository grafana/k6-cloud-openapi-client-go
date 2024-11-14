# PaginatedScheduleList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int32** | Object count in the collection | [optional] 
**NextLink** | Pointer to **string** | A reference to the next page of results. The property is included until there are no more pages of results to retrieve. | [optional] 
**Value** | [**[]Schedule**](Schedule.md) | List of the resulting values | 

## Methods

### NewPaginatedScheduleList

`func NewPaginatedScheduleList(value []Schedule, ) *PaginatedScheduleList`

NewPaginatedScheduleList instantiates a new PaginatedScheduleList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginatedScheduleListWithDefaults

`func NewPaginatedScheduleListWithDefaults() *PaginatedScheduleList`

NewPaginatedScheduleListWithDefaults instantiates a new PaginatedScheduleList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *PaginatedScheduleList) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *PaginatedScheduleList) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *PaginatedScheduleList) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *PaginatedScheduleList) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetNextLink

`func (o *PaginatedScheduleList) GetNextLink() string`

GetNextLink returns the NextLink field if non-nil, zero value otherwise.

### GetNextLinkOk

`func (o *PaginatedScheduleList) GetNextLinkOk() (*string, bool)`

GetNextLinkOk returns a tuple with the NextLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextLink

`func (o *PaginatedScheduleList) SetNextLink(v string)`

SetNextLink sets NextLink field to given value.

### HasNextLink

`func (o *PaginatedScheduleList) HasNextLink() bool`

HasNextLink returns a boolean if a field has been set.

### GetValue

`func (o *PaginatedScheduleList) GetValue() []Schedule`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *PaginatedScheduleList) GetValueOk() (*[]Schedule, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *PaginatedScheduleList) SetValue(v []Schedule)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


