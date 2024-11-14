# StaticIPReleaseRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ids** | **[]int32** | List of Static IPs IDs to release and delete. | 

## Methods

### NewStaticIPReleaseRequest

`func NewStaticIPReleaseRequest(ids []int32, ) *StaticIPReleaseRequest`

NewStaticIPReleaseRequest instantiates a new StaticIPReleaseRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStaticIPReleaseRequestWithDefaults

`func NewStaticIPReleaseRequestWithDefaults() *StaticIPReleaseRequest`

NewStaticIPReleaseRequestWithDefaults instantiates a new StaticIPReleaseRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIds

`func (o *StaticIPReleaseRequest) GetIds() []int32`

GetIds returns the Ids field if non-nil, zero value otherwise.

### GetIdsOk

`func (o *StaticIPReleaseRequest) GetIdsOk() (*[]int32, bool)`

GetIdsOk returns a tuple with the Ids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIds

`func (o *StaticIPReleaseRequest) SetIds(v []int32)`

SetIds sets Ids field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


