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

// checks if the ErrorApiModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ErrorApiModel{}

// ErrorApiModel Details of the error.
type ErrorApiModel struct {
	// Human-readable string describing the error.
	Message string `json:"message"`
	// Service-defined error code.
	Code string `json:"code"`
	// A string indicating the target of the error. For example, the name of the property in error.
	Target NullableString `json:"target,omitempty"`
	// Array of objects with more specific error information when applicable.
	Details              []ErrorDetailsApiModel `json:"details,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ErrorApiModel ErrorApiModel

// NewErrorApiModel instantiates a new ErrorApiModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewErrorApiModel(message string, code string) *ErrorApiModel {
	this := ErrorApiModel{}
	this.Message = message
	this.Code = code
	return &this
}

// NewErrorApiModelWithDefaults instantiates a new ErrorApiModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewErrorApiModelWithDefaults() *ErrorApiModel {
	this := ErrorApiModel{}
	return &this
}

// GetMessage returns the Message field value
func (o *ErrorApiModel) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *ErrorApiModel) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *ErrorApiModel) SetMessage(v string) {
	o.Message = v
}

// GetCode returns the Code field value
func (o *ErrorApiModel) GetCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *ErrorApiModel) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *ErrorApiModel) SetCode(v string) {
	o.Code = v
}

// GetTarget returns the Target field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ErrorApiModel) GetTarget() string {
	if o == nil || IsNil(o.Target.Get()) {
		var ret string
		return ret
	}
	return *o.Target.Get()
}

// GetTargetOk returns a tuple with the Target field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ErrorApiModel) GetTargetOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Target.Get(), o.Target.IsSet()
}

// HasTarget returns a boolean if a field has been set.
func (o *ErrorApiModel) HasTarget() bool {
	if o != nil && o.Target.IsSet() {
		return true
	}

	return false
}

// SetTarget gets a reference to the given NullableString and assigns it to the Target field.
func (o *ErrorApiModel) SetTarget(v string) {
	o.Target.Set(&v)
}

// SetTargetNil sets the value for Target to be an explicit nil
func (o *ErrorApiModel) SetTargetNil() {
	o.Target.Set(nil)
}

// UnsetTarget ensures that no value is present for Target, not even an explicit nil
func (o *ErrorApiModel) UnsetTarget() {
	o.Target.Unset()
}

// GetDetails returns the Details field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ErrorApiModel) GetDetails() []ErrorDetailsApiModel {
	if o == nil {
		var ret []ErrorDetailsApiModel
		return ret
	}
	return o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ErrorApiModel) GetDetailsOk() ([]ErrorDetailsApiModel, bool) {
	if o == nil || IsNil(o.Details) {
		return nil, false
	}
	return o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *ErrorApiModel) HasDetails() bool {
	if o != nil && !IsNil(o.Details) {
		return true
	}

	return false
}

// SetDetails gets a reference to the given []ErrorDetailsApiModel and assigns it to the Details field.
func (o *ErrorApiModel) SetDetails(v []ErrorDetailsApiModel) {
	o.Details = v
}

func (o ErrorApiModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ErrorApiModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["message"] = o.Message
	toSerialize["code"] = o.Code
	if o.Target.IsSet() {
		toSerialize["target"] = o.Target.Get()
	}
	if o.Details != nil {
		toSerialize["details"] = o.Details
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ErrorApiModel) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"message",
		"code",
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

	varErrorApiModel := _ErrorApiModel{}

	err = json.Unmarshal(data, &varErrorApiModel)

	if err != nil {
		return err
	}

	*o = ErrorApiModel(varErrorApiModel)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "message")
		delete(additionalProperties, "code")
		delete(additionalProperties, "target")
		delete(additionalProperties, "details")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableErrorApiModel struct {
	value *ErrorApiModel
	isSet bool
}

func (v NullableErrorApiModel) Get() *ErrorApiModel {
	return v.value
}

func (v *NullableErrorApiModel) Set(val *ErrorApiModel) {
	v.value = val
	v.isSet = true
}

func (v NullableErrorApiModel) IsSet() bool {
	return v.isSet
}

func (v *NullableErrorApiModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableErrorApiModel(val *ErrorApiModel) *NullableErrorApiModel {
	return &NullableErrorApiModel{value: val, isSet: true}
}

func (v NullableErrorApiModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableErrorApiModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
