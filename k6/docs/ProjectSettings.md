# ProjectSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VuhMaxPerMonth** | **int32** | Max amount of virtual user hours (VUh) used per one calendar month. | 
**VuMaxPerTest** | **int32** | Max number of concurrent virtual users (VUs) used in one test. | 
**VuBrowserMaxPerTest** | **int32** | Max number of concurrent browser virtual users (VUs) used in one test. | 
**DurationMaxPerTest** | **int32** | Max duration of a test in seconds. | 

## Methods

### NewProjectSettings

`func NewProjectSettings(vuhMaxPerMonth int32, vuMaxPerTest int32, vuBrowserMaxPerTest int32, durationMaxPerTest int32, ) *ProjectSettings`

NewProjectSettings instantiates a new ProjectSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectSettingsWithDefaults

`func NewProjectSettingsWithDefaults() *ProjectSettings`

NewProjectSettingsWithDefaults instantiates a new ProjectSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVuhMaxPerMonth

`func (o *ProjectSettings) GetVuhMaxPerMonth() int32`

GetVuhMaxPerMonth returns the VuhMaxPerMonth field if non-nil, zero value otherwise.

### GetVuhMaxPerMonthOk

`func (o *ProjectSettings) GetVuhMaxPerMonthOk() (*int32, bool)`

GetVuhMaxPerMonthOk returns a tuple with the VuhMaxPerMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVuhMaxPerMonth

`func (o *ProjectSettings) SetVuhMaxPerMonth(v int32)`

SetVuhMaxPerMonth sets VuhMaxPerMonth field to given value.


### GetVuMaxPerTest

`func (o *ProjectSettings) GetVuMaxPerTest() int32`

GetVuMaxPerTest returns the VuMaxPerTest field if non-nil, zero value otherwise.

### GetVuMaxPerTestOk

`func (o *ProjectSettings) GetVuMaxPerTestOk() (*int32, bool)`

GetVuMaxPerTestOk returns a tuple with the VuMaxPerTest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVuMaxPerTest

`func (o *ProjectSettings) SetVuMaxPerTest(v int32)`

SetVuMaxPerTest sets VuMaxPerTest field to given value.


### GetVuBrowserMaxPerTest

`func (o *ProjectSettings) GetVuBrowserMaxPerTest() int32`

GetVuBrowserMaxPerTest returns the VuBrowserMaxPerTest field if non-nil, zero value otherwise.

### GetVuBrowserMaxPerTestOk

`func (o *ProjectSettings) GetVuBrowserMaxPerTestOk() (*int32, bool)`

GetVuBrowserMaxPerTestOk returns a tuple with the VuBrowserMaxPerTest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVuBrowserMaxPerTest

`func (o *ProjectSettings) SetVuBrowserMaxPerTest(v int32)`

SetVuBrowserMaxPerTest sets VuBrowserMaxPerTest field to given value.


### GetDurationMaxPerTest

`func (o *ProjectSettings) GetDurationMaxPerTest() int32`

GetDurationMaxPerTest returns the DurationMaxPerTest field if non-nil, zero value otherwise.

### GetDurationMaxPerTestOk

`func (o *ProjectSettings) GetDurationMaxPerTestOk() (*int32, bool)`

GetDurationMaxPerTestOk returns a tuple with the DurationMaxPerTest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationMaxPerTest

`func (o *ProjectSettings) SetDurationMaxPerTest(v int32)`

SetDurationMaxPerTest sets DurationMaxPerTest field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


