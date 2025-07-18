/*
Grafana Cloud k6

HTTP API for interacting with Grafana Cloud k6.

API version: 1.5.0
Contact: info@grafana.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package k6

import (
	"encoding/json"
	"fmt"
)

// checks if the ProjectLimitsListResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProjectLimitsListResponse{}

// ProjectLimitsListResponse struct for ProjectLimitsListResponse
type ProjectLimitsListResponse struct {
	// List of the resulting values.
	Value []ProjectLimitsApiModel `json:"value"`
	// Object count in the collection.
	Count *int32 `json:"@count,omitempty"`
	// A reference to the next page of results. The property is included until there are no more pages of results to retrieve.
	NextLink *string `json:"@nextLink,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ProjectLimitsListResponse ProjectLimitsListResponse

// NewProjectLimitsListResponse instantiates a new ProjectLimitsListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectLimitsListResponse(value []ProjectLimitsApiModel) *ProjectLimitsListResponse {
	this := ProjectLimitsListResponse{}
	this.Value = value
	return &this
}

// NewProjectLimitsListResponseWithDefaults instantiates a new ProjectLimitsListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectLimitsListResponseWithDefaults() *ProjectLimitsListResponse {
	this := ProjectLimitsListResponse{}
	return &this
}

// GetValue returns the Value field value
func (o *ProjectLimitsListResponse) GetValue() []ProjectLimitsApiModel {
	if o == nil {
		var ret []ProjectLimitsApiModel
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *ProjectLimitsListResponse) GetValueOk() ([]ProjectLimitsApiModel, bool) {
	if o == nil {
		return nil, false
	}
	return o.Value, true
}

// SetValue sets field value
func (o *ProjectLimitsListResponse) SetValue(v []ProjectLimitsApiModel) {
	o.Value = v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *ProjectLimitsListResponse) GetCount() int32 {
	if o == nil || IsNil(o.Count) {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectLimitsListResponse) GetCountOk() (*int32, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *ProjectLimitsListResponse) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *ProjectLimitsListResponse) SetCount(v int32) {
	o.Count = &v
}

// GetNextLink returns the NextLink field value if set, zero value otherwise.
func (o *ProjectLimitsListResponse) GetNextLink() string {
	if o == nil || IsNil(o.NextLink) {
		var ret string
		return ret
	}
	return *o.NextLink
}

// GetNextLinkOk returns a tuple with the NextLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectLimitsListResponse) GetNextLinkOk() (*string, bool) {
	if o == nil || IsNil(o.NextLink) {
		return nil, false
	}
	return o.NextLink, true
}

// HasNextLink returns a boolean if a field has been set.
func (o *ProjectLimitsListResponse) HasNextLink() bool {
	if o != nil && !IsNil(o.NextLink) {
		return true
	}

	return false
}

// SetNextLink gets a reference to the given string and assigns it to the NextLink field.
func (o *ProjectLimitsListResponse) SetNextLink(v string) {
	o.NextLink = &v
}

func (o ProjectLimitsListResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProjectLimitsListResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["value"] = o.Value
	if !IsNil(o.Count) {
		toSerialize["@count"] = o.Count
	}
	if !IsNil(o.NextLink) {
		toSerialize["@nextLink"] = o.NextLink
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ProjectLimitsListResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"value",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varProjectLimitsListResponse := _ProjectLimitsListResponse{}

	err = json.Unmarshal(data, &varProjectLimitsListResponse)

	if err != nil {
		return err
	}

	*o = ProjectLimitsListResponse(varProjectLimitsListResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "value")
		delete(additionalProperties, "@count")
		delete(additionalProperties, "@nextLink")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableProjectLimitsListResponse struct {
	value *ProjectLimitsListResponse
	isSet bool
}

func (v NullableProjectLimitsListResponse) Get() *ProjectLimitsListResponse {
	return v.value
}

func (v *NullableProjectLimitsListResponse) Set(val *ProjectLimitsListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectLimitsListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectLimitsListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectLimitsListResponse(val *ProjectLimitsListResponse) *NullableProjectLimitsListResponse {
	return &NullableProjectLimitsListResponse{value: val, isSet: true}
}

func (v NullableProjectLimitsListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectLimitsListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


