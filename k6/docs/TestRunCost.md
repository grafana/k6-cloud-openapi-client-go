# TestRunCost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalVuh** | **float64** | Total number of VUh charged for the test run. | 
**Breakdown** | [**TestCostBreakdown**](TestCostBreakdown.md) | Breakdown details of the test cost. | 

## Methods

### NewTestRunCost

`func NewTestRunCost(totalVuh float64, breakdown TestCostBreakdown, ) *TestRunCost`

NewTestRunCost instantiates a new TestRunCost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestRunCostWithDefaults

`func NewTestRunCostWithDefaults() *TestRunCost`

NewTestRunCostWithDefaults instantiates a new TestRunCost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalVuh

`func (o *TestRunCost) GetTotalVuh() float64`

GetTotalVuh returns the TotalVuh field if non-nil, zero value otherwise.

### GetTotalVuhOk

`func (o *TestRunCost) GetTotalVuhOk() (*float64, bool)`

GetTotalVuhOk returns a tuple with the TotalVuh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalVuh

`func (o *TestRunCost) SetTotalVuh(v float64)`

SetTotalVuh sets TotalVuh field to given value.


### GetBreakdown

`func (o *TestRunCost) GetBreakdown() TestCostBreakdown`

GetBreakdown returns the Breakdown field if non-nil, zero value otherwise.

### GetBreakdownOk

`func (o *TestRunCost) GetBreakdownOk() (*TestCostBreakdown, bool)`

GetBreakdownOk returns a tuple with the Breakdown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBreakdown

`func (o *TestRunCost) SetBreakdown(v TestCostBreakdown)`

SetBreakdown sets Breakdown field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


