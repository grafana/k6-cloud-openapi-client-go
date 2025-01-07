# StatusApiModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Type of simple test run status: &#x60;created&#x60;, &#x60;queued&#x60;, &#x60;initializing&#x60;, &#x60;running&#x60;, &#x60;processing_metrics&#x60;, &#x60;completed&#x60; or &#x60;aborted&#x60;. | 
**Entered** | **time.Time** | Date and time when the test run entered the status. | 
**Extra** | Pointer to [**NullableStatusExtraApiModel**](StatusExtraApiModel.md) |  | [optional] 

## Methods

### NewStatusApiModel

`func NewStatusApiModel(type_ string, entered time.Time, ) *StatusApiModel`

NewStatusApiModel instantiates a new StatusApiModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatusApiModelWithDefaults

`func NewStatusApiModelWithDefaults() *StatusApiModel`

NewStatusApiModelWithDefaults instantiates a new StatusApiModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *StatusApiModel) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *StatusApiModel) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *StatusApiModel) SetType(v string)`

SetType sets Type field to given value.


### GetEntered

`func (o *StatusApiModel) GetEntered() time.Time`

GetEntered returns the Entered field if non-nil, zero value otherwise.

### GetEnteredOk

`func (o *StatusApiModel) GetEnteredOk() (*time.Time, bool)`

GetEnteredOk returns a tuple with the Entered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntered

`func (o *StatusApiModel) SetEntered(v time.Time)`

SetEntered sets Entered field to given value.


### GetExtra

`func (o *StatusApiModel) GetExtra() StatusExtraApiModel`

GetExtra returns the Extra field if non-nil, zero value otherwise.

### GetExtraOk

`func (o *StatusApiModel) GetExtraOk() (*StatusExtraApiModel, bool)`

GetExtraOk returns a tuple with the Extra field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtra

`func (o *StatusApiModel) SetExtra(v StatusExtraApiModel)`

SetExtra sets Extra field to given value.

### HasExtra

`func (o *StatusApiModel) HasExtra() bool`

HasExtra returns a boolean if a field has been set.

### SetExtraNil

`func (o *StatusApiModel) SetExtraNil(b bool)`

 SetExtraNil sets the value for Extra to be an explicit nil

### UnsetExtra
`func (o *StatusApiModel) UnsetExtra()`

UnsetExtra ensures that no value is present for Extra, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


