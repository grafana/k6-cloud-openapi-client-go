# LoadZoneDistributionApiModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Percentage** | **int32** | Percentage of the total load handled by this load zone. | 
**Nodes** | [**[]NodeApiModel**](NodeApiModel.md) | List of nodes allocated to this load zone. | 

## Methods

### NewLoadZoneDistributionApiModel

`func NewLoadZoneDistributionApiModel(percentage int32, nodes []NodeApiModel, ) *LoadZoneDistributionApiModel`

NewLoadZoneDistributionApiModel instantiates a new LoadZoneDistributionApiModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoadZoneDistributionApiModelWithDefaults

`func NewLoadZoneDistributionApiModelWithDefaults() *LoadZoneDistributionApiModel`

NewLoadZoneDistributionApiModelWithDefaults instantiates a new LoadZoneDistributionApiModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPercentage

`func (o *LoadZoneDistributionApiModel) GetPercentage() int32`

GetPercentage returns the Percentage field if non-nil, zero value otherwise.

### GetPercentageOk

`func (o *LoadZoneDistributionApiModel) GetPercentageOk() (*int32, bool)`

GetPercentageOk returns a tuple with the Percentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentage

`func (o *LoadZoneDistributionApiModel) SetPercentage(v int32)`

SetPercentage sets Percentage field to given value.


### GetNodes

`func (o *LoadZoneDistributionApiModel) GetNodes() []NodeApiModel`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *LoadZoneDistributionApiModel) GetNodesOk() (*[]NodeApiModel, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *LoadZoneDistributionApiModel) SetNodes(v []NodeApiModel)`

SetNodes sets Nodes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


