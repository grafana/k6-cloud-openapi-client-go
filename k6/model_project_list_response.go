/*
Grafana Cloud k6

HTTP API for interacting with Grafana Cloud k6.

API version: 1.0.0
Contact: info@grafana.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package k6

import (
	"encoding/json"
	"fmt"
)

// checks if the ProjectListResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProjectListResponse{}

// ProjectListResponse struct for ProjectListResponse
type ProjectListResponse struct {
	// List of the resulting values.
	Value []ProjectApiModel `json:"value"`
	// Object count in the collection.
	Count *int32 `json:"@count,omitempty"`
	// A reference to the next page of results. The property is included until there are no more pages of results to retrieve.
	NextLink             *string `json:"@nextLink,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ProjectListResponse ProjectListResponse

// NewProjectListResponse instantiates a new ProjectListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectListResponse(value []ProjectApiModel) *ProjectListResponse {
	this := ProjectListResponse{}
	this.Value = value
	return &this
}

// NewProjectListResponseWithDefaults instantiates a new ProjectListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectListResponseWithDefaults() *ProjectListResponse {
	this := ProjectListResponse{}
	return &this
}

// GetValue returns the Value field value
func (o *ProjectListResponse) GetValue() []ProjectApiModel {
	if o == nil {
		var ret []ProjectApiModel
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *ProjectListResponse) GetValueOk() ([]ProjectApiModel, bool) {
	if o == nil {
		return nil, false
	}
	return o.Value, true
}

// SetValue sets field value
func (o *ProjectListResponse) SetValue(v []ProjectApiModel) {
	o.Value = v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *ProjectListResponse) GetCount() int32 {
	if o == nil || IsNil(o.Count) {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectListResponse) GetCountOk() (*int32, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *ProjectListResponse) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *ProjectListResponse) SetCount(v int32) {
	o.Count = &v
}

// GetNextLink returns the NextLink field value if set, zero value otherwise.
func (o *ProjectListResponse) GetNextLink() string {
	if o == nil || IsNil(o.NextLink) {
		var ret string
		return ret
	}
	return *o.NextLink
}

// GetNextLinkOk returns a tuple with the NextLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectListResponse) GetNextLinkOk() (*string, bool) {
	if o == nil || IsNil(o.NextLink) {
		return nil, false
	}
	return o.NextLink, true
}

// HasNextLink returns a boolean if a field has been set.
func (o *ProjectListResponse) HasNextLink() bool {
	if o != nil && !IsNil(o.NextLink) {
		return true
	}

	return false
}

// SetNextLink gets a reference to the given string and assigns it to the NextLink field.
func (o *ProjectListResponse) SetNextLink(v string) {
	o.NextLink = &v
}

func (o ProjectListResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProjectListResponse) ToMap() (map[string]interface{}, error) {
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

func (o *ProjectListResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"value",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varProjectListResponse := _ProjectListResponse{}

	err = json.Unmarshal(data, &varProjectListResponse)

	if err != nil {
		return err
	}

	*o = ProjectListResponse(varProjectListResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "value")
		delete(additionalProperties, "@count")
		delete(additionalProperties, "@nextLink")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableProjectListResponse struct {
	value *ProjectListResponse
	isSet bool
}

func (v NullableProjectListResponse) Get() *ProjectListResponse {
	return v.value
}

func (v *NullableProjectListResponse) Set(val *ProjectListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectListResponse(val *ProjectListResponse) *NullableProjectListResponse {
	return &NullableProjectListResponse{value: val, isSet: true}
}

func (v NullableProjectListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
