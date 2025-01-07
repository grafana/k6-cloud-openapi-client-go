# PatchProjectLimitsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VuhMaxPerMonth** | Pointer to **NullableInt32** | Max amount of virtual user hours (VUH) used per one calendar month. | [optional] 
**VuMaxPerTest** | Pointer to **NullableInt32** | Max number of concurrent virtual users (VUs) used in one test. | [optional] 
**VuBrowserMaxPerTest** | Pointer to **NullableInt32** | Max number of concurrent browser virtual users (VUs) used in one test. | [optional] 
**DurationMaxPerTest** | Pointer to **NullableInt32** | Max duration of a test in seconds. | [optional] 

## Methods

### NewPatchProjectLimitsRequest

`func NewPatchProjectLimitsRequest() *PatchProjectLimitsRequest`

NewPatchProjectLimitsRequest instantiates a new PatchProjectLimitsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchProjectLimitsRequestWithDefaults

`func NewPatchProjectLimitsRequestWithDefaults() *PatchProjectLimitsRequest`

NewPatchProjectLimitsRequestWithDefaults instantiates a new PatchProjectLimitsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVuhMaxPerMonth

`func (o *PatchProjectLimitsRequest) GetVuhMaxPerMonth() int32`

GetVuhMaxPerMonth returns the VuhMaxPerMonth field if non-nil, zero value otherwise.

### GetVuhMaxPerMonthOk

`func (o *PatchProjectLimitsRequest) GetVuhMaxPerMonthOk() (*int32, bool)`

GetVuhMaxPerMonthOk returns a tuple with the VuhMaxPerMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVuhMaxPerMonth

`func (o *PatchProjectLimitsRequest) SetVuhMaxPerMonth(v int32)`

SetVuhMaxPerMonth sets VuhMaxPerMonth field to given value.

### HasVuhMaxPerMonth

`func (o *PatchProjectLimitsRequest) HasVuhMaxPerMonth() bool`

HasVuhMaxPerMonth returns a boolean if a field has been set.

### SetVuhMaxPerMonthNil

`func (o *PatchProjectLimitsRequest) SetVuhMaxPerMonthNil(b bool)`

 SetVuhMaxPerMonthNil sets the value for VuhMaxPerMonth to be an explicit nil

### UnsetVuhMaxPerMonth
`func (o *PatchProjectLimitsRequest) UnsetVuhMaxPerMonth()`

UnsetVuhMaxPerMonth ensures that no value is present for VuhMaxPerMonth, not even an explicit nil
### GetVuMaxPerTest

`func (o *PatchProjectLimitsRequest) GetVuMaxPerTest() int32`

GetVuMaxPerTest returns the VuMaxPerTest field if non-nil, zero value otherwise.

### GetVuMaxPerTestOk

`func (o *PatchProjectLimitsRequest) GetVuMaxPerTestOk() (*int32, bool)`

GetVuMaxPerTestOk returns a tuple with the VuMaxPerTest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVuMaxPerTest

`func (o *PatchProjectLimitsRequest) SetVuMaxPerTest(v int32)`

SetVuMaxPerTest sets VuMaxPerTest field to given value.

### HasVuMaxPerTest

`func (o *PatchProjectLimitsRequest) HasVuMaxPerTest() bool`

HasVuMaxPerTest returns a boolean if a field has been set.

### SetVuMaxPerTestNil

`func (o *PatchProjectLimitsRequest) SetVuMaxPerTestNil(b bool)`

 SetVuMaxPerTestNil sets the value for VuMaxPerTest to be an explicit nil

### UnsetVuMaxPerTest
`func (o *PatchProjectLimitsRequest) UnsetVuMaxPerTest()`

UnsetVuMaxPerTest ensures that no value is present for VuMaxPerTest, not even an explicit nil
### GetVuBrowserMaxPerTest

`func (o *PatchProjectLimitsRequest) GetVuBrowserMaxPerTest() int32`

GetVuBrowserMaxPerTest returns the VuBrowserMaxPerTest field if non-nil, zero value otherwise.

### GetVuBrowserMaxPerTestOk

`func (o *PatchProjectLimitsRequest) GetVuBrowserMaxPerTestOk() (*int32, bool)`

GetVuBrowserMaxPerTestOk returns a tuple with the VuBrowserMaxPerTest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVuBrowserMaxPerTest

`func (o *PatchProjectLimitsRequest) SetVuBrowserMaxPerTest(v int32)`

SetVuBrowserMaxPerTest sets VuBrowserMaxPerTest field to given value.

### HasVuBrowserMaxPerTest

`func (o *PatchProjectLimitsRequest) HasVuBrowserMaxPerTest() bool`

HasVuBrowserMaxPerTest returns a boolean if a field has been set.

### SetVuBrowserMaxPerTestNil

`func (o *PatchProjectLimitsRequest) SetVuBrowserMaxPerTestNil(b bool)`

 SetVuBrowserMaxPerTestNil sets the value for VuBrowserMaxPerTest to be an explicit nil

### UnsetVuBrowserMaxPerTest
`func (o *PatchProjectLimitsRequest) UnsetVuBrowserMaxPerTest()`

UnsetVuBrowserMaxPerTest ensures that no value is present for VuBrowserMaxPerTest, not even an explicit nil
### GetDurationMaxPerTest

`func (o *PatchProjectLimitsRequest) GetDurationMaxPerTest() int32`

GetDurationMaxPerTest returns the DurationMaxPerTest field if non-nil, zero value otherwise.

### GetDurationMaxPerTestOk

`func (o *PatchProjectLimitsRequest) GetDurationMaxPerTestOk() (*int32, bool)`

GetDurationMaxPerTestOk returns a tuple with the DurationMaxPerTest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationMaxPerTest

`func (o *PatchProjectLimitsRequest) SetDurationMaxPerTest(v int32)`

SetDurationMaxPerTest sets DurationMaxPerTest field to given value.

### HasDurationMaxPerTest

`func (o *PatchProjectLimitsRequest) HasDurationMaxPerTest() bool`

HasDurationMaxPerTest returns a boolean if a field has been set.

### SetDurationMaxPerTestNil

`func (o *PatchProjectLimitsRequest) SetDurationMaxPerTestNil(b bool)`

 SetDurationMaxPerTestNil sets the value for DurationMaxPerTest to be an explicit nil

### UnsetDurationMaxPerTest
`func (o *PatchProjectLimitsRequest) UnsetDurationMaxPerTest()`

UnsetDurationMaxPerTest ensures that no value is present for DurationMaxPerTest, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


