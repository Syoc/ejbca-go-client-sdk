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

// checks if the ExtendedInformationRestResponseComponent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExtendedInformationRestResponseComponent{}

// ExtendedInformationRestResponseComponent struct for ExtendedInformationRestResponseComponent
type ExtendedInformationRestResponseComponent struct {
	// Extended Information property name
	Name *string `json:"name,omitempty"`
	// Property value
	Value *string `json:"value,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ExtendedInformationRestResponseComponent ExtendedInformationRestResponseComponent

// NewExtendedInformationRestResponseComponent instantiates a new ExtendedInformationRestResponseComponent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExtendedInformationRestResponseComponent() *ExtendedInformationRestResponseComponent {
	this := ExtendedInformationRestResponseComponent{}
	return &this
}

// NewExtendedInformationRestResponseComponentWithDefaults instantiates a new ExtendedInformationRestResponseComponent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExtendedInformationRestResponseComponentWithDefaults() *ExtendedInformationRestResponseComponent {
	this := ExtendedInformationRestResponseComponent{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ExtendedInformationRestResponseComponent) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtendedInformationRestResponseComponent) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ExtendedInformationRestResponseComponent) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ExtendedInformationRestResponseComponent) SetName(v string) {
	o.Name = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *ExtendedInformationRestResponseComponent) GetValue() string {
	if o == nil || isNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtendedInformationRestResponseComponent) GetValueOk() (*string, bool) {
	if o == nil || isNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *ExtendedInformationRestResponseComponent) HasValue() bool {
	if o != nil && !isNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *ExtendedInformationRestResponseComponent) SetValue(v string) {
	o.Value = &v
}

func (o ExtendedInformationRestResponseComponent) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExtendedInformationRestResponseComponent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Value) {
		toSerialize["value"] = o.Value
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ExtendedInformationRestResponseComponent) UnmarshalJSON(bytes []byte) (err error) {
	varExtendedInformationRestResponseComponent := _ExtendedInformationRestResponseComponent{}

	if err = json.Unmarshal(bytes, &varExtendedInformationRestResponseComponent); err == nil {
		*o = ExtendedInformationRestResponseComponent(varExtendedInformationRestResponseComponent)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "value")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableExtendedInformationRestResponseComponent struct {
	value *ExtendedInformationRestResponseComponent
	isSet bool
}

func (v NullableExtendedInformationRestResponseComponent) Get() *ExtendedInformationRestResponseComponent {
	return v.value
}

func (v *NullableExtendedInformationRestResponseComponent) Set(val *ExtendedInformationRestResponseComponent) {
	v.value = val
	v.isSet = true
}

func (v NullableExtendedInformationRestResponseComponent) IsSet() bool {
	return v.isSet
}

func (v *NullableExtendedInformationRestResponseComponent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExtendedInformationRestResponseComponent(val *ExtendedInformationRestResponseComponent) *NullableExtendedInformationRestResponseComponent {
	return &NullableExtendedInformationRestResponseComponent{value: val, isSet: true}
}

func (v NullableExtendedInformationRestResponseComponent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExtendedInformationRestResponseComponent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

