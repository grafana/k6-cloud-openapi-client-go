/*
Grafana Cloud k6

HTTP API for interacting with Grafana Cloud k6.

API version: 1.1.0
Contact: info@grafana.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package k6

import (
	"encoding/json"
	"fmt"
)

// checks if the ValidateOptionsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ValidateOptionsRequest{}

// ValidateOptionsRequest struct for ValidateOptionsRequest
type ValidateOptionsRequest struct {
	// ID of a project where the test belongs.
	ProjectId NullableInt32 `json:"project_id,omitempty"`
	// k6 script options object to validate.
	Options              map[string]interface{} `json:"options"`
	AdditionalProperties map[string]interface{}
}

type _ValidateOptionsRequest ValidateOptionsRequest

// NewValidateOptionsRequest instantiates a new ValidateOptionsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewValidateOptionsRequest(options map[string]interface{}) *ValidateOptionsRequest {
	this := ValidateOptionsRequest{}
	this.Options = options
	return &this
}

// NewValidateOptionsRequestWithDefaults instantiates a new ValidateOptionsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewValidateOptionsRequestWithDefaults() *ValidateOptionsRequest {
	this := ValidateOptionsRequest{}
	return &this
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ValidateOptionsRequest) GetProjectId() int32 {
	if o == nil || IsNil(o.ProjectId.Get()) {
		var ret int32
		return ret
	}
	return *o.ProjectId.Get()
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ValidateOptionsRequest) GetProjectIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProjectId.Get(), o.ProjectId.IsSet()
}

// HasProjectId returns a boolean if a field has been set.
func (o *ValidateOptionsRequest) HasProjectId() bool {
	if o != nil && o.ProjectId.IsSet() {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given NullableInt32 and assigns it to the ProjectId field.
func (o *ValidateOptionsRequest) SetProjectId(v int32) {
	o.ProjectId.Set(&v)
}

// SetProjectIdNil sets the value for ProjectId to be an explicit nil
func (o *ValidateOptionsRequest) SetProjectIdNil() {
	o.ProjectId.Set(nil)
}

// UnsetProjectId ensures that no value is present for ProjectId, not even an explicit nil
func (o *ValidateOptionsRequest) UnsetProjectId() {
	o.ProjectId.Unset()
}

// GetOptions returns the Options field value
func (o *ValidateOptionsRequest) GetOptions() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value
// and a boolean to check if the value has been set.
func (o *ValidateOptionsRequest) GetOptionsOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Options, true
}

// SetOptions sets field value
func (o *ValidateOptionsRequest) SetOptions(v map[string]interface{}) {
	o.Options = v
}

func (o ValidateOptionsRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ValidateOptionsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.ProjectId.IsSet() {
		toSerialize["project_id"] = o.ProjectId.Get()
	}
	toSerialize["options"] = o.Options

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ValidateOptionsRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"options",
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

	varValidateOptionsRequest := _ValidateOptionsRequest{}

	err = json.Unmarshal(data, &varValidateOptionsRequest)

	if err != nil {
		return err
	}

	*o = ValidateOptionsRequest(varValidateOptionsRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "project_id")
		delete(additionalProperties, "options")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableValidateOptionsRequest struct {
	value *ValidateOptionsRequest
	isSet bool
}

func (v NullableValidateOptionsRequest) Get() *ValidateOptionsRequest {
	return v.value
}

func (v *NullableValidateOptionsRequest) Set(val *ValidateOptionsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableValidateOptionsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableValidateOptionsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableValidateOptionsRequest(val *ValidateOptionsRequest) *NullableValidateOptionsRequest {
	return &NullableValidateOptionsRequest{value: val, isSet: true}
}

func (v NullableValidateOptionsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableValidateOptionsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
