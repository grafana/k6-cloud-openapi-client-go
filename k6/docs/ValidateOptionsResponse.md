# ValidateOptionsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VuhUsage** | **float32** | How many VUH will be charged for the test. | 
**Breakdown** | [**CostBreakdownApiModel**](CostBreakdownApiModel.md) | Breakdown of the VUH usage. | 

## Methods

### NewValidateOptionsResponse

`func NewValidateOptionsResponse(vuhUsage float32, breakdown CostBreakdownApiModel, ) *ValidateOptionsResponse`

NewValidateOptionsResponse instantiates a new ValidateOptionsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewValidateOptionsResponseWithDefaults

`func NewValidateOptionsResponseWithDefaults() *ValidateOptionsResponse`

NewValidateOptionsResponseWithDefaults instantiates a new ValidateOptionsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVuhUsage

`func (o *ValidateOptionsResponse) GetVuhUsage() float32`

GetVuhUsage returns the VuhUsage field if non-nil, zero value otherwise.

### GetVuhUsageOk

`func (o *ValidateOptionsResponse) GetVuhUsageOk() (*float32, bool)`

GetVuhUsageOk returns a tuple with the VuhUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVuhUsage

`func (o *ValidateOptionsResponse) SetVuhUsage(v float32)`

SetVuhUsage sets VuhUsage field to given value.


### GetBreakdown

`func (o *ValidateOptionsResponse) GetBreakdown() CostBreakdownApiModel`

GetBreakdown returns the Breakdown field if non-nil, zero value otherwise.

### GetBreakdownOk

`func (o *ValidateOptionsResponse) GetBreakdownOk() (*CostBreakdownApiModel, bool)`

GetBreakdownOk returns a tuple with the Breakdown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBreakdown

`func (o *ValidateOptionsResponse) SetBreakdown(v CostBreakdownApiModel)`

SetBreakdown sets Breakdown field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


