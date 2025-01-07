# ProjectLimitsApiModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProjectId** | **int32** | ID of the related project. | 
**VuhMaxPerMonth** | **NullableInt32** | Max amount of virtual user hours (VUH) used per one calendar month. | 
**VuMaxPerTest** | **NullableInt32** | Max number of concurrent virtual users (VUs) used in one test. | 
**VuBrowserMaxPerTest** | **NullableInt32** | Max number of concurrent browser virtual users (VUs) used in one test. | 
**DurationMaxPerTest** | **NullableInt32** | Max duration of a test in seconds. | 

## Methods

### NewProjectLimitsApiModel

`func NewProjectLimitsApiModel(projectId int32, vuhMaxPerMonth NullableInt32, vuMaxPerTest NullableInt32, vuBrowserMaxPerTest NullableInt32, durationMaxPerTest NullableInt32, ) *ProjectLimitsApiModel`

NewProjectLimitsApiModel instantiates a new ProjectLimitsApiModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectLimitsApiModelWithDefaults

`func NewProjectLimitsApiModelWithDefaults() *ProjectLimitsApiModel`

NewProjectLimitsApiModelWithDefaults instantiates a new ProjectLimitsApiModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjectId

`func (o *ProjectLimitsApiModel) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *ProjectLimitsApiModel) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *ProjectLimitsApiModel) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.


### GetVuhMaxPerMonth

`func (o *ProjectLimitsApiModel) GetVuhMaxPerMonth() int32`

GetVuhMaxPerMonth returns the VuhMaxPerMonth field if non-nil, zero value otherwise.

### GetVuhMaxPerMonthOk

`func (o *ProjectLimitsApiModel) GetVuhMaxPerMonthOk() (*int32, bool)`

GetVuhMaxPerMonthOk returns a tuple with the VuhMaxPerMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVuhMaxPerMonth

`func (o *ProjectLimitsApiModel) SetVuhMaxPerMonth(v int32)`

SetVuhMaxPerMonth sets VuhMaxPerMonth field to given value.


### SetVuhMaxPerMonthNil

`func (o *ProjectLimitsApiModel) SetVuhMaxPerMonthNil(b bool)`

 SetVuhMaxPerMonthNil sets the value for VuhMaxPerMonth to be an explicit nil

### UnsetVuhMaxPerMonth
`func (o *ProjectLimitsApiModel) UnsetVuhMaxPerMonth()`

UnsetVuhMaxPerMonth ensures that no value is present for VuhMaxPerMonth, not even an explicit nil
### GetVuMaxPerTest

`func (o *ProjectLimitsApiModel) GetVuMaxPerTest() int32`

GetVuMaxPerTest returns the VuMaxPerTest field if non-nil, zero value otherwise.

### GetVuMaxPerTestOk

`func (o *ProjectLimitsApiModel) GetVuMaxPerTestOk() (*int32, bool)`

GetVuMaxPerTestOk returns a tuple with the VuMaxPerTest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVuMaxPerTest

`func (o *ProjectLimitsApiModel) SetVuMaxPerTest(v int32)`

SetVuMaxPerTest sets VuMaxPerTest field to given value.


### SetVuMaxPerTestNil

`func (o *ProjectLimitsApiModel) SetVuMaxPerTestNil(b bool)`

 SetVuMaxPerTestNil sets the value for VuMaxPerTest to be an explicit nil

### UnsetVuMaxPerTest
`func (o *ProjectLimitsApiModel) UnsetVuMaxPerTest()`

UnsetVuMaxPerTest ensures that no value is present for VuMaxPerTest, not even an explicit nil
### GetVuBrowserMaxPerTest

`func (o *ProjectLimitsApiModel) GetVuBrowserMaxPerTest() int32`

GetVuBrowserMaxPerTest returns the VuBrowserMaxPerTest field if non-nil, zero value otherwise.

### GetVuBrowserMaxPerTestOk

`func (o *ProjectLimitsApiModel) GetVuBrowserMaxPerTestOk() (*int32, bool)`

GetVuBrowserMaxPerTestOk returns a tuple with the VuBrowserMaxPerTest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVuBrowserMaxPerTest

`func (o *ProjectLimitsApiModel) SetVuBrowserMaxPerTest(v int32)`

SetVuBrowserMaxPerTest sets VuBrowserMaxPerTest field to given value.


### SetVuBrowserMaxPerTestNil

`func (o *ProjectLimitsApiModel) SetVuBrowserMaxPerTestNil(b bool)`

 SetVuBrowserMaxPerTestNil sets the value for VuBrowserMaxPerTest to be an explicit nil

### UnsetVuBrowserMaxPerTest
`func (o *ProjectLimitsApiModel) UnsetVuBrowserMaxPerTest()`

UnsetVuBrowserMaxPerTest ensures that no value is present for VuBrowserMaxPerTest, not even an explicit nil
### GetDurationMaxPerTest

`func (o *ProjectLimitsApiModel) GetDurationMaxPerTest() int32`

GetDurationMaxPerTest returns the DurationMaxPerTest field if non-nil, zero value otherwise.

### GetDurationMaxPerTestOk

`func (o *ProjectLimitsApiModel) GetDurationMaxPerTestOk() (*int32, bool)`

GetDurationMaxPerTestOk returns a tuple with the DurationMaxPerTest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationMaxPerTest

`func (o *ProjectLimitsApiModel) SetDurationMaxPerTest(v int32)`

SetDurationMaxPerTest sets DurationMaxPerTest field to given value.


### SetDurationMaxPerTestNil

`func (o *ProjectLimitsApiModel) SetDurationMaxPerTestNil(b bool)`

 SetDurationMaxPerTestNil sets the value for DurationMaxPerTest to be an explicit nil

### UnsetDurationMaxPerTest
`func (o *ProjectLimitsApiModel) UnsetDurationMaxPerTest()`

UnsetDurationMaxPerTest ensures that no value is present for DurationMaxPerTest, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


