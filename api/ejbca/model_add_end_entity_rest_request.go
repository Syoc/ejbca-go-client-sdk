/*
Copyright 2024 Keyfactor

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

EJBCA REST Interface

API reference documentation.

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ejbca

import (
	"encoding/json"
)

// checks if the AddEndEntityRestRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddEndEntityRestRequest{}

// AddEndEntityRestRequest struct for AddEndEntityRestRequest
type AddEndEntityRestRequest struct {
	// Username
	Username *string `json:"username,omitempty"`
	// Password
	Password *string `json:"password,omitempty"`
	// Subject Distinguished Name
	SubjectDn *string `json:"subject_dn,omitempty"`
	// Subject Alternative Name (SAN)
	SubjectAltName *string `json:"subject_alt_name,omitempty"`
	// Email
	Email         *string                                   `json:"email,omitempty"`
	ExtensionData []ExtendedInformationRestRequestComponent `json:"extension_data,omitempty"`
	// Certificate Authority (CA) name
	CaName *string `json:"ca_name,omitempty"`
	// Certificate profile name
	CertificateProfileName *string `json:"certificate_profile_name,omitempty"`
	// End Entity profile name
	EndEntityProfileName *string `json:"end_entity_profile_name,omitempty"`
	// Token type property
	Token *string `json:"token,omitempty"`
	// Account Binding ID
	AccountBindingId     *string `json:"account_binding_id,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AddEndEntityRestRequest AddEndEntityRestRequest

// NewAddEndEntityRestRequest instantiates a new AddEndEntityRestRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddEndEntityRestRequest() *AddEndEntityRestRequest {
	this := AddEndEntityRestRequest{}
	return &this
}

// NewAddEndEntityRestRequestWithDefaults instantiates a new AddEndEntityRestRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddEndEntityRestRequestWithDefaults() *AddEndEntityRestRequest {
	this := AddEndEntityRestRequest{}
	return &this
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *AddEndEntityRestRequest) GetUsername() string {
	if o == nil || isNil(o.Username) {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddEndEntityRestRequest) GetUsernameOk() (*string, bool) {
	if o == nil || isNil(o.Username) {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *AddEndEntityRestRequest) HasUsername() bool {
	if o != nil && !isNil(o.Username) {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *AddEndEntityRestRequest) SetUsername(v string) {
	o.Username = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *AddEndEntityRestRequest) GetPassword() string {
	if o == nil || isNil(o.Password) {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddEndEntityRestRequest) GetPasswordOk() (*string, bool) {
	if o == nil || isNil(o.Password) {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *AddEndEntityRestRequest) HasPassword() bool {
	if o != nil && !isNil(o.Password) {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *AddEndEntityRestRequest) SetPassword(v string) {
	o.Password = &v
}

// GetSubjectDn returns the SubjectDn field value if set, zero value otherwise.
func (o *AddEndEntityRestRequest) GetSubjectDn() string {
	if o == nil || isNil(o.SubjectDn) {
		var ret string
		return ret
	}
	return *o.SubjectDn
}

// GetSubjectDnOk returns a tuple with the SubjectDn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddEndEntityRestRequest) GetSubjectDnOk() (*string, bool) {
	if o == nil || isNil(o.SubjectDn) {
		return nil, false
	}
	return o.SubjectDn, true
}

// HasSubjectDn returns a boolean if a field has been set.
func (o *AddEndEntityRestRequest) HasSubjectDn() bool {
	if o != nil && !isNil(o.SubjectDn) {
		return true
	}

	return false
}

// SetSubjectDn gets a reference to the given string and assigns it to the SubjectDn field.
func (o *AddEndEntityRestRequest) SetSubjectDn(v string) {
	o.SubjectDn = &v
}

// GetSubjectAltName returns the SubjectAltName field value if set, zero value otherwise.
func (o *AddEndEntityRestRequest) GetSubjectAltName() string {
	if o == nil || isNil(o.SubjectAltName) {
		var ret string
		return ret
	}
	return *o.SubjectAltName
}

// GetSubjectAltNameOk returns a tuple with the SubjectAltName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddEndEntityRestRequest) GetSubjectAltNameOk() (*string, bool) {
	if o == nil || isNil(o.SubjectAltName) {
		return nil, false
	}
	return o.SubjectAltName, true
}

// HasSubjectAltName returns a boolean if a field has been set.
func (o *AddEndEntityRestRequest) HasSubjectAltName() bool {
	if o != nil && !isNil(o.SubjectAltName) {
		return true
	}

	return false
}

// SetSubjectAltName gets a reference to the given string and assigns it to the SubjectAltName field.
func (o *AddEndEntityRestRequest) SetSubjectAltName(v string) {
	o.SubjectAltName = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *AddEndEntityRestRequest) GetEmail() string {
	if o == nil || isNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddEndEntityRestRequest) GetEmailOk() (*string, bool) {
	if o == nil || isNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *AddEndEntityRestRequest) HasEmail() bool {
	if o != nil && !isNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *AddEndEntityRestRequest) SetEmail(v string) {
	o.Email = &v
}

// GetExtensionData returns the ExtensionData field value if set, zero value otherwise.
func (o *AddEndEntityRestRequest) GetExtensionData() []ExtendedInformationRestRequestComponent {
	if o == nil || isNil(o.ExtensionData) {
		var ret []ExtendedInformationRestRequestComponent
		return ret
	}
	return o.ExtensionData
}

// GetExtensionDataOk returns a tuple with the ExtensionData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddEndEntityRestRequest) GetExtensionDataOk() ([]ExtendedInformationRestRequestComponent, bool) {
	if o == nil || isNil(o.ExtensionData) {
		return nil, false
	}
	return o.ExtensionData, true
}

// HasExtensionData returns a boolean if a field has been set.
func (o *AddEndEntityRestRequest) HasExtensionData() bool {
	if o != nil && !isNil(o.ExtensionData) {
		return true
	}

	return false
}

// SetExtensionData gets a reference to the given []ExtendedInformationRestRequestComponent and assigns it to the ExtensionData field.
func (o *AddEndEntityRestRequest) SetExtensionData(v []ExtendedInformationRestRequestComponent) {
	o.ExtensionData = v
}

// GetCaName returns the CaName field value if set, zero value otherwise.
func (o *AddEndEntityRestRequest) GetCaName() string {
	if o == nil || isNil(o.CaName) {
		var ret string
		return ret
	}
	return *o.CaName
}

// GetCaNameOk returns a tuple with the CaName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddEndEntityRestRequest) GetCaNameOk() (*string, bool) {
	if o == nil || isNil(o.CaName) {
		return nil, false
	}
	return o.CaName, true
}

// HasCaName returns a boolean if a field has been set.
func (o *AddEndEntityRestRequest) HasCaName() bool {
	if o != nil && !isNil(o.CaName) {
		return true
	}

	return false
}

// SetCaName gets a reference to the given string and assigns it to the CaName field.
func (o *AddEndEntityRestRequest) SetCaName(v string) {
	o.CaName = &v
}

// GetCertificateProfileName returns the CertificateProfileName field value if set, zero value otherwise.
func (o *AddEndEntityRestRequest) GetCertificateProfileName() string {
	if o == nil || isNil(o.CertificateProfileName) {
		var ret string
		return ret
	}
	return *o.CertificateProfileName
}

// GetCertificateProfileNameOk returns a tuple with the CertificateProfileName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddEndEntityRestRequest) GetCertificateProfileNameOk() (*string, bool) {
	if o == nil || isNil(o.CertificateProfileName) {
		return nil, false
	}
	return o.CertificateProfileName, true
}

// HasCertificateProfileName returns a boolean if a field has been set.
func (o *AddEndEntityRestRequest) HasCertificateProfileName() bool {
	if o != nil && !isNil(o.CertificateProfileName) {
		return true
	}

	return false
}

// SetCertificateProfileName gets a reference to the given string and assigns it to the CertificateProfileName field.
func (o *AddEndEntityRestRequest) SetCertificateProfileName(v string) {
	o.CertificateProfileName = &v
}

// GetEndEntityProfileName returns the EndEntityProfileName field value if set, zero value otherwise.
func (o *AddEndEntityRestRequest) GetEndEntityProfileName() string {
	if o == nil || isNil(o.EndEntityProfileName) {
		var ret string
		return ret
	}
	return *o.EndEntityProfileName
}

// GetEndEntityProfileNameOk returns a tuple with the EndEntityProfileName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddEndEntityRestRequest) GetEndEntityProfileNameOk() (*string, bool) {
	if o == nil || isNil(o.EndEntityProfileName) {
		return nil, false
	}
	return o.EndEntityProfileName, true
}

// HasEndEntityProfileName returns a boolean if a field has been set.
func (o *AddEndEntityRestRequest) HasEndEntityProfileName() bool {
	if o != nil && !isNil(o.EndEntityProfileName) {
		return true
	}

	return false
}

// SetEndEntityProfileName gets a reference to the given string and assigns it to the EndEntityProfileName field.
func (o *AddEndEntityRestRequest) SetEndEntityProfileName(v string) {
	o.EndEntityProfileName = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *AddEndEntityRestRequest) GetToken() string {
	if o == nil || isNil(o.Token) {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddEndEntityRestRequest) GetTokenOk() (*string, bool) {
	if o == nil || isNil(o.Token) {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *AddEndEntityRestRequest) HasToken() bool {
	if o != nil && !isNil(o.Token) {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *AddEndEntityRestRequest) SetToken(v string) {
	o.Token = &v
}

// GetAccountBindingId returns the AccountBindingId field value if set, zero value otherwise.
func (o *AddEndEntityRestRequest) GetAccountBindingId() string {
	if o == nil || isNil(o.AccountBindingId) {
		var ret string
		return ret
	}
	return *o.AccountBindingId
}

// GetAccountBindingIdOk returns a tuple with the AccountBindingId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddEndEntityRestRequest) GetAccountBindingIdOk() (*string, bool) {
	if o == nil || isNil(o.AccountBindingId) {
		return nil, false
	}
	return o.AccountBindingId, true
}

// HasAccountBindingId returns a boolean if a field has been set.
func (o *AddEndEntityRestRequest) HasAccountBindingId() bool {
	if o != nil && !isNil(o.AccountBindingId) {
		return true
	}

	return false
}

// SetAccountBindingId gets a reference to the given string and assigns it to the AccountBindingId field.
func (o *AddEndEntityRestRequest) SetAccountBindingId(v string) {
	o.AccountBindingId = &v
}

func (o AddEndEntityRestRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddEndEntityRestRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Username) {
		toSerialize["username"] = o.Username
	}
	if !isNil(o.Password) {
		toSerialize["password"] = o.Password
	}
	if !isNil(o.SubjectDn) {
		toSerialize["subject_dn"] = o.SubjectDn
	}
	if !isNil(o.SubjectAltName) {
		toSerialize["subject_alt_name"] = o.SubjectAltName
	}
	if !isNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !isNil(o.ExtensionData) {
		toSerialize["extension_data"] = o.ExtensionData
	}
	if !isNil(o.CaName) {
		toSerialize["ca_name"] = o.CaName
	}
	if !isNil(o.CertificateProfileName) {
		toSerialize["certificate_profile_name"] = o.CertificateProfileName
	}
	if !isNil(o.EndEntityProfileName) {
		toSerialize["end_entity_profile_name"] = o.EndEntityProfileName
	}
	if !isNil(o.Token) {
		toSerialize["token"] = o.Token
	}
	if !isNil(o.AccountBindingId) {
		toSerialize["account_binding_id"] = o.AccountBindingId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AddEndEntityRestRequest) UnmarshalJSON(bytes []byte) (err error) {
	varAddEndEntityRestRequest := _AddEndEntityRestRequest{}

	if err = json.Unmarshal(bytes, &varAddEndEntityRestRequest); err == nil {
		*o = AddEndEntityRestRequest(varAddEndEntityRestRequest)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "username")
		delete(additionalProperties, "password")
		delete(additionalProperties, "subject_dn")
		delete(additionalProperties, "subject_alt_name")
		delete(additionalProperties, "email")
		delete(additionalProperties, "extension_data")
		delete(additionalProperties, "ca_name")
		delete(additionalProperties, "certificate_profile_name")
		delete(additionalProperties, "end_entity_profile_name")
		delete(additionalProperties, "token")
		delete(additionalProperties, "account_binding_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAddEndEntityRestRequest struct {
	value *AddEndEntityRestRequest
	isSet bool
}

func (v NullableAddEndEntityRestRequest) Get() *AddEndEntityRestRequest {
	return v.value
}

func (v *NullableAddEndEntityRestRequest) Set(val *AddEndEntityRestRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddEndEntityRestRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddEndEntityRestRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddEndEntityRestRequest(val *AddEndEntityRestRequest) *NullableAddEndEntityRestRequest {
	return &NullableAddEndEntityRestRequest{value: val, isSet: true}
}

func (v NullableAddEndEntityRestRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddEndEntityRestRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
