# TestRunDistributionApiModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Distribution** | [**map[string]LoadZoneDistributionApiModel**](LoadZoneDistributionApiModel.md) | Distribution of nodes across the load generators for the test run. | 

## Methods

### NewTestRunDistributionApiModel

`func NewTestRunDistributionApiModel(distribution map[string]LoadZoneDistributionApiModel, ) *TestRunDistributionApiModel`

NewTestRunDistributionApiModel instantiates a new TestRunDistributionApiModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestRunDistributionApiModelWithDefaults

`func NewTestRunDistributionApiModelWithDefaults() *TestRunDistributionApiModel`

NewTestRunDistributionApiModelWithDefaults instantiates a new TestRunDistributionApiModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDistribution

`func (o *TestRunDistributionApiModel) GetDistribution() map[string]LoadZoneDistributionApiModel`

GetDistribution returns the Distribution field if non-nil, zero value otherwise.

### GetDistributionOk

`func (o *TestRunDistributionApiModel) GetDistributionOk() (*map[string]LoadZoneDistributionApiModel, bool)`

GetDistributionOk returns a tuple with the Distribution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistribution

`func (o *TestRunDistributionApiModel) SetDistribution(v map[string]LoadZoneDistributionApiModel)`

SetDistribution sets Distribution field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


