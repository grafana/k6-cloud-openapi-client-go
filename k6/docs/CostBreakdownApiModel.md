# CostBreakdownApiModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProtocolVuh** | [**ProtocolVuh**](ProtocolVuh.md) |  | 
**BrowserVuh** | [**BrowserVuh**](BrowserVuh.md) |  | 
**BaseTotalVuh** | [**BaseTotalVuh**](BaseTotalVuh.md) |  | 
**ReductionRate** | [**ReductionRate**](ReductionRate.md) |  | 
**ReductionRateBreakdown** | [**map[string]ReductionRateBreakdownValue**](ReductionRateBreakdownValue.md) | The individual reduction rates applied to the base VUH usage. | 

## Methods

### NewCostBreakdownApiModel

`func NewCostBreakdownApiModel(protocolVuh ProtocolVuh, browserVuh BrowserVuh, baseTotalVuh BaseTotalVuh, reductionRate ReductionRate, reductionRateBreakdown map[string]ReductionRateBreakdownValue, ) *CostBreakdownApiModel`

NewCostBreakdownApiModel instantiates a new CostBreakdownApiModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCostBreakdownApiModelWithDefaults

`func NewCostBreakdownApiModelWithDefaults() *CostBreakdownApiModel`

NewCostBreakdownApiModelWithDefaults instantiates a new CostBreakdownApiModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProtocolVuh

`func (o *CostBreakdownApiModel) GetProtocolVuh() ProtocolVuh`

GetProtocolVuh returns the ProtocolVuh field if non-nil, zero value otherwise.

### GetProtocolVuhOk

`func (o *CostBreakdownApiModel) GetProtocolVuhOk() (*ProtocolVuh, bool)`

GetProtocolVuhOk returns a tuple with the ProtocolVuh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocolVuh

`func (o *CostBreakdownApiModel) SetProtocolVuh(v ProtocolVuh)`

SetProtocolVuh sets ProtocolVuh field to given value.


### GetBrowserVuh

`func (o *CostBreakdownApiModel) GetBrowserVuh() BrowserVuh`

GetBrowserVuh returns the BrowserVuh field if non-nil, zero value otherwise.

### GetBrowserVuhOk

`func (o *CostBreakdownApiModel) GetBrowserVuhOk() (*BrowserVuh, bool)`

GetBrowserVuhOk returns a tuple with the BrowserVuh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrowserVuh

`func (o *CostBreakdownApiModel) SetBrowserVuh(v BrowserVuh)`

SetBrowserVuh sets BrowserVuh field to given value.


### GetBaseTotalVuh

`func (o *CostBreakdownApiModel) GetBaseTotalVuh() BaseTotalVuh`

GetBaseTotalVuh returns the BaseTotalVuh field if non-nil, zero value otherwise.

### GetBaseTotalVuhOk

`func (o *CostBreakdownApiModel) GetBaseTotalVuhOk() (*BaseTotalVuh, bool)`

GetBaseTotalVuhOk returns a tuple with the BaseTotalVuh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseTotalVuh

`func (o *CostBreakdownApiModel) SetBaseTotalVuh(v BaseTotalVuh)`

SetBaseTotalVuh sets BaseTotalVuh field to given value.


### GetReductionRate

`func (o *CostBreakdownApiModel) GetReductionRate() ReductionRate`

GetReductionRate returns the ReductionRate field if non-nil, zero value otherwise.

### GetReductionRateOk

`func (o *CostBreakdownApiModel) GetReductionRateOk() (*ReductionRate, bool)`

GetReductionRateOk returns a tuple with the ReductionRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReductionRate

`func (o *CostBreakdownApiModel) SetReductionRate(v ReductionRate)`

SetReductionRate sets ReductionRate field to given value.


### GetReductionRateBreakdown

`func (o *CostBreakdownApiModel) GetReductionRateBreakdown() map[string]ReductionRateBreakdownValue`

GetReductionRateBreakdown returns the ReductionRateBreakdown field if non-nil, zero value otherwise.

### GetReductionRateBreakdownOk

`func (o *CostBreakdownApiModel) GetReductionRateBreakdownOk() (*map[string]ReductionRateBreakdownValue, bool)`

GetReductionRateBreakdownOk returns a tuple with the ReductionRateBreakdown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReductionRateBreakdown

`func (o *CostBreakdownApiModel) SetReductionRateBreakdown(v map[string]ReductionRateBreakdownValue)`

SetReductionRateBreakdown sets ReductionRateBreakdown field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


