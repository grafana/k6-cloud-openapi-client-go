# NodeApiModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Size** | **string** | Size/type of the node (e.g., &#39;m5.large&#39;). | 
**PublicIp** | **string** | Public IP address of the node. | 

## Methods

### NewNodeApiModel

`func NewNodeApiModel(size string, publicIp string, ) *NodeApiModel`

NewNodeApiModel instantiates a new NodeApiModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNodeApiModelWithDefaults

`func NewNodeApiModelWithDefaults() *NodeApiModel`

NewNodeApiModelWithDefaults instantiates a new NodeApiModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSize

`func (o *NodeApiModel) GetSize() string`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *NodeApiModel) GetSizeOk() (*string, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *NodeApiModel) SetSize(v string)`

SetSize sets Size field to given value.


### GetPublicIp

`func (o *NodeApiModel) GetPublicIp() string`

GetPublicIp returns the PublicIp field if non-nil, zero value otherwise.

### GetPublicIpOk

`func (o *NodeApiModel) GetPublicIpOk() (*string, bool)`

GetPublicIpOk returns a tuple with the PublicIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicIp

`func (o *NodeApiModel) SetPublicIp(v string)`

SetPublicIp sets PublicIp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


