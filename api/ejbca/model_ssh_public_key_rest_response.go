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

// checks if the SshPublicKeyRestResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SshPublicKeyRestResponse{}

// SshPublicKeyRestResponse struct for SshPublicKeyRestResponse
type SshPublicKeyRestResponse struct {
	// Certificate Authority (CA) name
	CaName *string `json:"ca_name,omitempty"`
	// CA’s public key
	Response *string `json:"response,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SshPublicKeyRestResponse SshPublicKeyRestResponse

// NewSshPublicKeyRestResponse instantiates a new SshPublicKeyRestResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSshPublicKeyRestResponse() *SshPublicKeyRestResponse {
	this := SshPublicKeyRestResponse{}
	return &this
}

// NewSshPublicKeyRestResponseWithDefaults instantiates a new SshPublicKeyRestResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSshPublicKeyRestResponseWithDefaults() *SshPublicKeyRestResponse {
	this := SshPublicKeyRestResponse{}
	return &this
}

// GetCaName returns the CaName field value if set, zero value otherwise.
func (o *SshPublicKeyRestResponse) GetCaName() string {
	if o == nil || isNil(o.CaName) {
		var ret string
		return ret
	}
	return *o.CaName
}

// GetCaNameOk returns a tuple with the CaName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SshPublicKeyRestResponse) GetCaNameOk() (*string, bool) {
	if o == nil || isNil(o.CaName) {
		return nil, false
	}
	return o.CaName, true
}

// HasCaName returns a boolean if a field has been set.
func (o *SshPublicKeyRestResponse) HasCaName() bool {
	if o != nil && !isNil(o.CaName) {
		return true
	}

	return false
}

// SetCaName gets a reference to the given string and assigns it to the CaName field.
func (o *SshPublicKeyRestResponse) SetCaName(v string) {
	o.CaName = &v
}

// GetResponse returns the Response field value if set, zero value otherwise.
func (o *SshPublicKeyRestResponse) GetResponse() string {
	if o == nil || isNil(o.Response) {
		var ret string
		return ret
	}
	return *o.Response
}

// GetResponseOk returns a tuple with the Response field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SshPublicKeyRestResponse) GetResponseOk() (*string, bool) {
	if o == nil || isNil(o.Response) {
		return nil, false
	}
	return o.Response, true
}

// HasResponse returns a boolean if a field has been set.
func (o *SshPublicKeyRestResponse) HasResponse() bool {
	if o != nil && !isNil(o.Response) {
		return true
	}

	return false
}

// SetResponse gets a reference to the given string and assigns it to the Response field.
func (o *SshPublicKeyRestResponse) SetResponse(v string) {
	o.Response = &v
}

func (o SshPublicKeyRestResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SshPublicKeyRestResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.CaName) {
		toSerialize["ca_name"] = o.CaName
	}
	if !isNil(o.Response) {
		toSerialize["response"] = o.Response
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SshPublicKeyRestResponse) UnmarshalJSON(bytes []byte) (err error) {
	varSshPublicKeyRestResponse := _SshPublicKeyRestResponse{}

	if err = json.Unmarshal(bytes, &varSshPublicKeyRestResponse); err == nil {
		*o = SshPublicKeyRestResponse(varSshPublicKeyRestResponse)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ca_name")
		delete(additionalProperties, "response")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSshPublicKeyRestResponse struct {
	value *SshPublicKeyRestResponse
	isSet bool
}

func (v NullableSshPublicKeyRestResponse) Get() *SshPublicKeyRestResponse {
	return v.value
}

func (v *NullableSshPublicKeyRestResponse) Set(val *SshPublicKeyRestResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSshPublicKeyRestResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSshPublicKeyRestResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSshPublicKeyRestResponse(val *SshPublicKeyRestResponse) *NullableSshPublicKeyRestResponse {
	return &NullableSshPublicKeyRestResponse{value: val, isSet: true}
}

func (v NullableSshPublicKeyRestResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSshPublicKeyRestResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

