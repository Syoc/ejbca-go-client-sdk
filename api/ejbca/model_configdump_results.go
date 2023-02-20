/*
EJBCA REST Interface

API reference documentation.

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ejbca

import (
	"encoding/json"
)

// checks if the ConfigdumpResults type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConfigdumpResults{}

// ConfigdumpResults struct for ConfigdumpResults
type ConfigdumpResults struct {
	Success *bool `json:"success,omitempty"`
	Errors []string `json:"errors,omitempty"`
	Warnings []string `json:"warnings,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ConfigdumpResults ConfigdumpResults

// NewConfigdumpResults instantiates a new ConfigdumpResults object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConfigdumpResults() *ConfigdumpResults {
	this := ConfigdumpResults{}
	return &this
}

// NewConfigdumpResultsWithDefaults instantiates a new ConfigdumpResults object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConfigdumpResultsWithDefaults() *ConfigdumpResults {
	this := ConfigdumpResults{}
	return &this
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *ConfigdumpResults) GetSuccess() bool {
	if o == nil || isNil(o.Success) {
		var ret bool
		return ret
	}
	return *o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigdumpResults) GetSuccessOk() (*bool, bool) {
	if o == nil || isNil(o.Success) {
		return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *ConfigdumpResults) HasSuccess() bool {
	if o != nil && !isNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given bool and assigns it to the Success field.
func (o *ConfigdumpResults) SetSuccess(v bool) {
	o.Success = &v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *ConfigdumpResults) GetErrors() []string {
	if o == nil || isNil(o.Errors) {
		var ret []string
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigdumpResults) GetErrorsOk() ([]string, bool) {
	if o == nil || isNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *ConfigdumpResults) HasErrors() bool {
	if o != nil && !isNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []string and assigns it to the Errors field.
func (o *ConfigdumpResults) SetErrors(v []string) {
	o.Errors = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *ConfigdumpResults) GetWarnings() []string {
	if o == nil || isNil(o.Warnings) {
		var ret []string
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigdumpResults) GetWarningsOk() ([]string, bool) {
	if o == nil || isNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *ConfigdumpResults) HasWarnings() bool {
	if o != nil && !isNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []string and assigns it to the Warnings field.
func (o *ConfigdumpResults) SetWarnings(v []string) {
	o.Warnings = v
}

func (o ConfigdumpResults) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConfigdumpResults) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Success) {
		toSerialize["success"] = o.Success
	}
	if !isNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	if !isNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ConfigdumpResults) UnmarshalJSON(bytes []byte) (err error) {
	varConfigdumpResults := _ConfigdumpResults{}

	if err = json.Unmarshal(bytes, &varConfigdumpResults); err == nil {
		*o = ConfigdumpResults(varConfigdumpResults)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "success")
		delete(additionalProperties, "errors")
		delete(additionalProperties, "warnings")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableConfigdumpResults struct {
	value *ConfigdumpResults
	isSet bool
}

func (v NullableConfigdumpResults) Get() *ConfigdumpResults {
	return v.value
}

func (v *NullableConfigdumpResults) Set(val *ConfigdumpResults) {
	v.value = val
	v.isSet = true
}

func (v NullableConfigdumpResults) IsSet() bool {
	return v.isSet
}

func (v *NullableConfigdumpResults) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConfigdumpResults(val *ConfigdumpResults) *NullableConfigdumpResults {
	return &NullableConfigdumpResults{value: val, isSet: true}
}

func (v NullableConfigdumpResults) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConfigdumpResults) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

