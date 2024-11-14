# TestRunStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Type of the test run status. | 
**Entered** | **time.Time** | Date and time when the test run entered the status. | 
**Extra** | Pointer to [**TestRunStatusExtra**](TestRunStatusExtra.md) |  | [optional] 

## Methods

### NewTestRunStatus

`func NewTestRunStatus(type_ string, entered time.Time, ) *TestRunStatus`

NewTestRunStatus instantiates a new TestRunStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestRunStatusWithDefaults

`func NewTestRunStatusWithDefaults() *TestRunStatus`

NewTestRunStatusWithDefaults instantiates a new TestRunStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *TestRunStatus) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TestRunStatus) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TestRunStatus) SetType(v string)`

SetType sets Type field to given value.


### GetEntered

`func (o *TestRunStatus) GetEntered() time.Time`

GetEntered returns the Entered field if non-nil, zero value otherwise.

### GetEnteredOk

`func (o *TestRunStatus) GetEnteredOk() (*time.Time, bool)`

GetEnteredOk returns a tuple with the Entered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntered

`func (o *TestRunStatus) SetEntered(v time.Time)`

SetEntered sets Entered field to given value.


### GetExtra

`func (o *TestRunStatus) GetExtra() TestRunStatusExtra`

GetExtra returns the Extra field if non-nil, zero value otherwise.

### GetExtraOk

`func (o *TestRunStatus) GetExtraOk() (*TestRunStatusExtra, bool)`

GetExtraOk returns a tuple with the Extra field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtra

`func (o *TestRunStatus) SetExtra(v TestRunStatusExtra)`

SetExtra sets Extra field to given value.

### HasExtra

`func (o *TestRunStatus) HasExtra() bool`

HasExtra returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


