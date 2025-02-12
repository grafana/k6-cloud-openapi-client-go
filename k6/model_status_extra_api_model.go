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

// checks if the StatusExtraApiModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StatusExtraApiModel{}

// StatusExtraApiModel struct for StatusExtraApiModel
type StatusExtraApiModel struct {
	// Email of the user that set the status if applicable.
	ByUser NullableString `json:"by_user"`
	// Human-readable string describing the error if applicable.
	Message NullableString `json:"message"`
	// Service-defined error code if applicable.
	Code                 NullableInt32 `json:"code"`
	AdditionalProperties map[string]interface{}
}

type _StatusExtraApiModel StatusExtraApiModel

// NewStatusExtraApiModel instantiates a new StatusExtraApiModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStatusExtraApiModel(byUser NullableString, message NullableString, code NullableInt32) *StatusExtraApiModel {
	this := StatusExtraApiModel{}
	this.ByUser = byUser
	this.Message = message
	this.Code = code
	return &this
}

// NewStatusExtraApiModelWithDefaults instantiates a new StatusExtraApiModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStatusExtraApiModelWithDefaults() *StatusExtraApiModel {
	this := StatusExtraApiModel{}
	return &this
}

// GetByUser returns the ByUser field value
// If the value is explicit nil, the zero value for string will be returned
func (o *StatusExtraApiModel) GetByUser() string {
	if o == nil || o.ByUser.Get() == nil {
		var ret string
		return ret
	}

	return *o.ByUser.Get()
}

// GetByUserOk returns a tuple with the ByUser field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StatusExtraApiModel) GetByUserOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ByUser.Get(), o.ByUser.IsSet()
}

// SetByUser sets field value
func (o *StatusExtraApiModel) SetByUser(v string) {
	o.ByUser.Set(&v)
}

// GetMessage returns the Message field value
// If the value is explicit nil, the zero value for string will be returned
func (o *StatusExtraApiModel) GetMessage() string {
	if o == nil || o.Message.Get() == nil {
		var ret string
		return ret
	}

	return *o.Message.Get()
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StatusExtraApiModel) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Message.Get(), o.Message.IsSet()
}

// SetMessage sets field value
func (o *StatusExtraApiModel) SetMessage(v string) {
	o.Message.Set(&v)
}

// GetCode returns the Code field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *StatusExtraApiModel) GetCode() int32 {
	if o == nil || o.Code.Get() == nil {
		var ret int32
		return ret
	}

	return *o.Code.Get()
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StatusExtraApiModel) GetCodeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Code.Get(), o.Code.IsSet()
}

// SetCode sets field value
func (o *StatusExtraApiModel) SetCode(v int32) {
	o.Code.Set(&v)
}

func (o StatusExtraApiModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StatusExtraApiModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["by_user"] = o.ByUser.Get()
	toSerialize["message"] = o.Message.Get()
	toSerialize["code"] = o.Code.Get()

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *StatusExtraApiModel) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"by_user",
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

	varStatusExtraApiModel := _StatusExtraApiModel{}

	err = json.Unmarshal(data, &varStatusExtraApiModel)

	if err != nil {
		return err
	}

	*o = StatusExtraApiModel(varStatusExtraApiModel)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "by_user")
		delete(additionalProperties, "message")
		delete(additionalProperties, "code")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableStatusExtraApiModel struct {
	value *StatusExtraApiModel
	isSet bool
}

func (v NullableStatusExtraApiModel) Get() *StatusExtraApiModel {
	return v.value
}

func (v *NullableStatusExtraApiModel) Set(val *StatusExtraApiModel) {
	v.value = val
	v.isSet = true
}

func (v NullableStatusExtraApiModel) IsSet() bool {
	return v.isSet
}

func (v *NullableStatusExtraApiModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStatusExtraApiModel(val *StatusExtraApiModel) *NullableStatusExtraApiModel {
	return &NullableStatusExtraApiModel{value: val, isSet: true}
}

func (v NullableStatusExtraApiModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStatusExtraApiModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
