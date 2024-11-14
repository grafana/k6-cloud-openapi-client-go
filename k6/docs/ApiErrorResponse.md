# ApiErrorResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | [**ApiError**](ApiError.md) | Details of the error. | 

## Methods

### NewApiErrorResponse

`func NewApiErrorResponse(error_ ApiError, ) *ApiErrorResponse`

NewApiErrorResponse instantiates a new ApiErrorResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiErrorResponseWithDefaults

`func NewApiErrorResponseWithDefaults() *ApiErrorResponse`

NewApiErrorResponseWithDefaults instantiates a new ApiErrorResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *ApiErrorResponse) GetError() ApiError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ApiErrorResponse) GetErrorOk() (*ApiError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ApiErrorResponse) SetError(v ApiError)`

SetError sets Error field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


