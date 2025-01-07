/*


HTTP API for interacting with k6 Cloud.

API version: 0.0.0
Contact: info@grafana.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package k6

import (
	"encoding/json"
)

// checks if the LoadZone type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LoadZone{}

// LoadZone struct for LoadZone
type LoadZone struct {
	// Id of the load zone.
	Id int32 `json:"id"`
	// Unique load zone name used in k6 script.
	Name NullableString `json:"name"`
	// Human readable load zone name.
	Title string `json:"title"`
	// If the load zone is a private load zone owned by the organization.
	Private bool `json:"private"`
	// Name of the load zone provider. Usually, `amazon` for public load zones and `k8s` for the private ones.
	ProviderName string `json:"provider_name"`
	// Provider specific details.
	ProviderDetails      map[string]interface{} `json:"provider_details"`
	AdditionalProperties map[string]interface{}
}

type _LoadZone LoadZone

// NewLoadZone instantiates a new LoadZone object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLoadZone(id int32, name NullableString, title string, private bool, providerName string, providerDetails map[string]interface{}) *LoadZone {
	this := LoadZone{}
	this.Id = id
	this.Name = name
	this.Title = title
	this.Private = private
	this.ProviderName = providerName
	this.ProviderDetails = providerDetails
	return &this
}

// NewLoadZoneWithDefaults instantiates a new LoadZone object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLoadZoneWithDefaults() *LoadZone {
	this := LoadZone{}
	return &this
}

// GetId returns the Id field value
func (o *LoadZone) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *LoadZone) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *LoadZone) SetId(v int32) {
	o.Id = v
}

// GetName returns the Name field value
// If the value is explicit nil, the zero value for string will be returned
func (o *LoadZone) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}

	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LoadZone) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// SetName sets field value
func (o *LoadZone) SetName(v string) {
	o.Name.Set(&v)
}

// GetTitle returns the Title field value
func (o *LoadZone) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *LoadZone) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *LoadZone) SetTitle(v string) {
	o.Title = v
}

// GetPrivate returns the Private field value
func (o *LoadZone) GetPrivate() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Private
}

// GetPrivateOk returns a tuple with the Private field value
// and a boolean to check if the value has been set.
func (o *LoadZone) GetPrivateOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Private, true
}

// SetPrivate sets field value
func (o *LoadZone) SetPrivate(v bool) {
	o.Private = v
}

// GetProviderName returns the ProviderName field value
func (o *LoadZone) GetProviderName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProviderName
}

// GetProviderNameOk returns a tuple with the ProviderName field value
// and a boolean to check if the value has been set.
func (o *LoadZone) GetProviderNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProviderName, true
}

// SetProviderName sets field value
func (o *LoadZone) SetProviderName(v string) {
	o.ProviderName = v
}

// GetProviderDetails returns the ProviderDetails field value
func (o *LoadZone) GetProviderDetails() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.ProviderDetails
}

// GetProviderDetailsOk returns a tuple with the ProviderDetails field value
// and a boolean to check if the value has been set.
func (o *LoadZone) GetProviderDetailsOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.ProviderDetails, true
}

// SetProviderDetails sets field value
func (o *LoadZone) SetProviderDetails(v map[string]interface{}) {
	o.ProviderDetails = v
}

func (o LoadZone) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LoadZone) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name.Get()
	toSerialize["title"] = o.Title
	toSerialize["private"] = o.Private
	toSerialize["provider_name"] = o.ProviderName
	toSerialize["provider_details"] = o.ProviderDetails

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *LoadZone) UnmarshalJSON(data []byte) (err error) {
	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	varLoadZone := _LoadZone{}

	err = json.Unmarshal(data, &varLoadZone)

	if err != nil {
		return err
	}

	*o = LoadZone(varLoadZone)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "title")
		delete(additionalProperties, "private")
		delete(additionalProperties, "provider_name")
		delete(additionalProperties, "provider_details")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableLoadZone struct {
	value *LoadZone
	isSet bool
}

func (v NullableLoadZone) Get() *LoadZone {
	return v.value
}

func (v *NullableLoadZone) Set(val *LoadZone) {
	v.value = val
	v.isSet = true
}

func (v NullableLoadZone) IsSet() bool {
	return v.isSet
}

func (v *NullableLoadZone) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLoadZone(val *LoadZone) *NullableLoadZone {
	return &NullableLoadZone{value: val, isSet: true}
}

func (v NullableLoadZone) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLoadZone) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
