/*
Tweakwise Navigator Backend API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package backend

import (
	"encoding/json"
)

// ValidationErrorContainer fluentValidator error container
type ValidationErrorContainer struct {
	// Message
	Message *string `json:"Message,omitempty"`
	// List of validation Errors
	Errors *[]ValidationError `json:"Errors,omitempty"`
}

// NewValidationErrorContainer instantiates a new ValidationErrorContainer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewValidationErrorContainer() *ValidationErrorContainer {
	this := ValidationErrorContainer{}
	return &this
}

// NewValidationErrorContainerWithDefaults instantiates a new ValidationErrorContainer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewValidationErrorContainerWithDefaults() *ValidationErrorContainer {
	this := ValidationErrorContainer{}
	return &this
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *ValidationErrorContainer) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidationErrorContainer) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *ValidationErrorContainer) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *ValidationErrorContainer) SetMessage(v string) {
	o.Message = &v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *ValidationErrorContainer) GetErrors() []ValidationError {
	if o == nil || o.Errors == nil {
		var ret []ValidationError
		return ret
	}
	return *o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidationErrorContainer) GetErrorsOk() (*[]ValidationError, bool) {
	if o == nil || o.Errors == nil {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *ValidationErrorContainer) HasErrors() bool {
	if o != nil && o.Errors != nil {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []ValidationError and assigns it to the Errors field.
func (o *ValidationErrorContainer) SetErrors(v []ValidationError) {
	o.Errors = &v
}

func (o ValidationErrorContainer) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Message != nil {
		toSerialize["Message"] = o.Message
	}
	if o.Errors != nil {
		toSerialize["Errors"] = o.Errors
	}
	return json.Marshal(toSerialize)
}

type NullableValidationErrorContainer struct {
	value *ValidationErrorContainer
	isSet bool
}

func (v NullableValidationErrorContainer) Get() *ValidationErrorContainer {
	return v.value
}

func (v *NullableValidationErrorContainer) Set(val *ValidationErrorContainer) {
	v.value = val
	v.isSet = true
}

func (v NullableValidationErrorContainer) IsSet() bool {
	return v.isSet
}

func (v *NullableValidationErrorContainer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableValidationErrorContainer(val *ValidationErrorContainer) *NullableValidationErrorContainer {
	return &NullableValidationErrorContainer{value: val, isSet: true}
}

func (v NullableValidationErrorContainer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableValidationErrorContainer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

