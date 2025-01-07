# StaticIP

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Id of the Static IP. | [readonly] 
**Ip** | Pointer to **NullableString** | Acquired IP address or null if the IP is still in the &#x60;provisioning&#x60; status. | [optional] 
**Created** | **time.Time** | The date and time the Static IP acquisition was requested. | [readonly] 
**ProvisioningStatus** | **string** | Provisioning status of the Static IP. Only the IPs with status &#x60;provisioned&#x60; areavailable in tests.  * &#x60;provisioning&#x60; - provisioning * &#x60;provisioned&#x60; - provisioned * &#x60;releasing&#x60; - releasing | 
**LoadZone** | **string** | Load zone identifier in form of (&lt;vendor:countrycode:city&gt;) | [readonly] 

## Methods

### NewStaticIP

`func NewStaticIP(id int32, created time.Time, provisioningStatus string, loadZone string, ) *StaticIP`

NewStaticIP instantiates a new StaticIP object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStaticIPWithDefaults

`func NewStaticIPWithDefaults() *StaticIP`

NewStaticIPWithDefaults instantiates a new StaticIP object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *StaticIP) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StaticIP) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StaticIP) SetId(v int32)`

SetId sets Id field to given value.


### GetIp

`func (o *StaticIP) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *StaticIP) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *StaticIP) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *StaticIP) HasIp() bool`

HasIp returns a boolean if a field has been set.

### SetIpNil

`func (o *StaticIP) SetIpNil(b bool)`

 SetIpNil sets the value for Ip to be an explicit nil

### UnsetIp
`func (o *StaticIP) UnsetIp()`

UnsetIp ensures that no value is present for Ip, not even an explicit nil
### GetCreated

`func (o *StaticIP) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *StaticIP) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *StaticIP) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetProvisioningStatus

`func (o *StaticIP) GetProvisioningStatus() string`

GetProvisioningStatus returns the ProvisioningStatus field if non-nil, zero value otherwise.

### GetProvisioningStatusOk

`func (o *StaticIP) GetProvisioningStatusOk() (*string, bool)`

GetProvisioningStatusOk returns a tuple with the ProvisioningStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningStatus

`func (o *StaticIP) SetProvisioningStatus(v string)`

SetProvisioningStatus sets ProvisioningStatus field to given value.


### GetLoadZone

`func (o *StaticIP) GetLoadZone() string`

GetLoadZone returns the LoadZone field if non-nil, zero value otherwise.

### GetLoadZoneOk

`func (o *StaticIP) GetLoadZoneOk() (*string, bool)`

GetLoadZoneOk returns a tuple with the LoadZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadZone

`func (o *StaticIP) SetLoadZone(v string)`

SetLoadZone sets LoadZone field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


