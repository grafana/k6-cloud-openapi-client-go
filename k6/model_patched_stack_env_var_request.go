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

// checks if the PatchedStackEnvVarRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchedStackEnvVarRequest{}

// PatchedStackEnvVarRequest struct for PatchedStackEnvVarRequest
type PatchedStackEnvVarRequest struct {
	// Name of the variable.
	Name *string `json:"name,omitempty"`
	// Variable value.
	Value *string `json:"value,omitempty"`
	// Description of the variable.
	Description          NullableString `json:"description,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PatchedStackEnvVarRequest PatchedStackEnvVarRequest

// NewPatchedStackEnvVarRequest instantiates a new PatchedStackEnvVarRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedStackEnvVarRequest() *PatchedStackEnvVarRequest {
	this := PatchedStackEnvVarRequest{}
	return &this
}

// NewPatchedStackEnvVarRequestWithDefaults instantiates a new PatchedStackEnvVarRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedStackEnvVarRequestWithDefaults() *PatchedStackEnvVarRequest {
	this := PatchedStackEnvVarRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PatchedStackEnvVarRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedStackEnvVarRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchedStackEnvVarRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PatchedStackEnvVarRequest) SetName(v string) {
	o.Name = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *PatchedStackEnvVarRequest) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedStackEnvVarRequest) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *PatchedStackEnvVarRequest) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *PatchedStackEnvVarRequest) SetValue(v string) {
	o.Value = &v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedStackEnvVarRequest) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedStackEnvVarRequest) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *PatchedStackEnvVarRequest) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *PatchedStackEnvVarRequest) SetDescription(v string) {
	o.Description.Set(&v)
}

// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *PatchedStackEnvVarRequest) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *PatchedStackEnvVarRequest) UnsetDescription() {
	o.Description.Unset()
}

func (o PatchedStackEnvVarRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchedStackEnvVarRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PatchedStackEnvVarRequest) UnmarshalJSON(data []byte) (err error) {
	varPatchedStackEnvVarRequest := _PatchedStackEnvVarRequest{}

	err = json.Unmarshal(data, &varPatchedStackEnvVarRequest)

	if err != nil {
		return err
	}

	*o = PatchedStackEnvVarRequest(varPatchedStackEnvVarRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "value")
		delete(additionalProperties, "description")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePatchedStackEnvVarRequest struct {
	value *PatchedStackEnvVarRequest
	isSet bool
}

func (v NullablePatchedStackEnvVarRequest) Get() *PatchedStackEnvVarRequest {
	return v.value
}

func (v *NullablePatchedStackEnvVarRequest) Set(val *PatchedStackEnvVarRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedStackEnvVarRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedStackEnvVarRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedStackEnvVarRequest(val *PatchedStackEnvVarRequest) *NullablePatchedStackEnvVarRequest {
	return &NullablePatchedStackEnvVarRequest{value: val, isSet: true}
}

func (v NullablePatchedStackEnvVarRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedStackEnvVarRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
