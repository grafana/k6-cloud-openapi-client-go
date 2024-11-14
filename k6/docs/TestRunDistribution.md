# TestRunDistribution

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LoadZone** | **string** | Name of the load zone. | 
**Percent** | **int32** | Percentage of the load that runs from the load zone. | 

## Methods

### NewTestRunDistribution

`func NewTestRunDistribution(loadZone string, percent int32, ) *TestRunDistribution`

NewTestRunDistribution instantiates a new TestRunDistribution object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestRunDistributionWithDefaults

`func NewTestRunDistributionWithDefaults() *TestRunDistribution`

NewTestRunDistributionWithDefaults instantiates a new TestRunDistribution object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLoadZone

`func (o *TestRunDistribution) GetLoadZone() string`

GetLoadZone returns the LoadZone field if non-nil, zero value otherwise.

### GetLoadZoneOk

`func (o *TestRunDistribution) GetLoadZoneOk() (*string, bool)`

GetLoadZoneOk returns a tuple with the LoadZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadZone

`func (o *TestRunDistribution) SetLoadZone(v string)`

SetLoadZone sets LoadZone field to given value.


### GetPercent

`func (o *TestRunDistribution) GetPercent() int32`

GetPercent returns the Percent field if non-nil, zero value otherwise.

### GetPercentOk

`func (o *TestRunDistribution) GetPercentOk() (*int32, bool)`

GetPercentOk returns a tuple with the Percent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercent

`func (o *TestRunDistribution) SetPercent(v int32)`

SetPercent sets Percent field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


