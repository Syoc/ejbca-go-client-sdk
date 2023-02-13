# AddEndEntityRestRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Username** | Pointer to **string** | Username | [optional] 
**Password** | Pointer to **string** | Password | [optional] 
**SubjectDn** | Pointer to **string** | Subject Distinguished Name | [optional] 
**SubjectAltName** | Pointer to **string** | Subject Alternative Name (SAN) | [optional] 
**Email** | Pointer to **string** | Email | [optional] 
**ExtensionData** | Pointer to [**[]ExtendedInformationRestRequestComponent**](ExtendedInformationRestRequestComponent.md) |  | [optional] 
**CaName** | Pointer to **string** | Certificate Authority (CA) name | [optional] 
**CertificateProfileName** | Pointer to **string** | Certificate profile name | [optional] 
**EndEntityProfileName** | Pointer to **string** | End Entity profile name | [optional] 
**Token** | Pointer to **string** | Token type property | [optional] 
**AccountBindingId** | Pointer to **string** | Account Binding ID | [optional] 

## Methods

### NewAddEndEntityRestRequest

`func NewAddEndEntityRestRequest() *AddEndEntityRestRequest`

NewAddEndEntityRestRequest instantiates a new AddEndEntityRestRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddEndEntityRestRequestWithDefaults

`func NewAddEndEntityRestRequestWithDefaults() *AddEndEntityRestRequest`

NewAddEndEntityRestRequestWithDefaults instantiates a new AddEndEntityRestRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsername

`func (o *AddEndEntityRestRequest) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *AddEndEntityRestRequest) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *AddEndEntityRestRequest) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *AddEndEntityRestRequest) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetPassword

`func (o *AddEndEntityRestRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *AddEndEntityRestRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *AddEndEntityRestRequest) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *AddEndEntityRestRequest) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetSubjectDn

`func (o *AddEndEntityRestRequest) GetSubjectDn() string`

GetSubjectDn returns the SubjectDn field if non-nil, zero value otherwise.

### GetSubjectDnOk

`func (o *AddEndEntityRestRequest) GetSubjectDnOk() (*string, bool)`

GetSubjectDnOk returns a tuple with the SubjectDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectDn

`func (o *AddEndEntityRestRequest) SetSubjectDn(v string)`

SetSubjectDn sets SubjectDn field to given value.

### HasSubjectDn

`func (o *AddEndEntityRestRequest) HasSubjectDn() bool`

HasSubjectDn returns a boolean if a field has been set.

### GetSubjectAltName

`func (o *AddEndEntityRestRequest) GetSubjectAltName() string`

GetSubjectAltName returns the SubjectAltName field if non-nil, zero value otherwise.

### GetSubjectAltNameOk

`func (o *AddEndEntityRestRequest) GetSubjectAltNameOk() (*string, bool)`

GetSubjectAltNameOk returns a tuple with the SubjectAltName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectAltName

`func (o *AddEndEntityRestRequest) SetSubjectAltName(v string)`

SetSubjectAltName sets SubjectAltName field to given value.

### HasSubjectAltName

`func (o *AddEndEntityRestRequest) HasSubjectAltName() bool`

HasSubjectAltName returns a boolean if a field has been set.

### GetEmail

`func (o *AddEndEntityRestRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *AddEndEntityRestRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *AddEndEntityRestRequest) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *AddEndEntityRestRequest) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetExtensionData

`func (o *AddEndEntityRestRequest) GetExtensionData() []ExtendedInformationRestRequestComponent`

GetExtensionData returns the ExtensionData field if non-nil, zero value otherwise.

### GetExtensionDataOk

`func (o *AddEndEntityRestRequest) GetExtensionDataOk() (*[]ExtendedInformationRestRequestComponent, bool)`

GetExtensionDataOk returns a tuple with the ExtensionData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionData

`func (o *AddEndEntityRestRequest) SetExtensionData(v []ExtendedInformationRestRequestComponent)`

SetExtensionData sets ExtensionData field to given value.

### HasExtensionData

`func (o *AddEndEntityRestRequest) HasExtensionData() bool`

HasExtensionData returns a boolean if a field has been set.

### GetCaName

`func (o *AddEndEntityRestRequest) GetCaName() string`

GetCaName returns the CaName field if non-nil, zero value otherwise.

### GetCaNameOk

`func (o *AddEndEntityRestRequest) GetCaNameOk() (*string, bool)`

GetCaNameOk returns a tuple with the CaName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaName

`func (o *AddEndEntityRestRequest) SetCaName(v string)`

SetCaName sets CaName field to given value.

### HasCaName

`func (o *AddEndEntityRestRequest) HasCaName() bool`

HasCaName returns a boolean if a field has been set.

### GetCertificateProfileName

`func (o *AddEndEntityRestRequest) GetCertificateProfileName() string`

GetCertificateProfileName returns the CertificateProfileName field if non-nil, zero value otherwise.

### GetCertificateProfileNameOk

`func (o *AddEndEntityRestRequest) GetCertificateProfileNameOk() (*string, bool)`

GetCertificateProfileNameOk returns a tuple with the CertificateProfileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateProfileName

`func (o *AddEndEntityRestRequest) SetCertificateProfileName(v string)`

SetCertificateProfileName sets CertificateProfileName field to given value.

### HasCertificateProfileName

`func (o *AddEndEntityRestRequest) HasCertificateProfileName() bool`

HasCertificateProfileName returns a boolean if a field has been set.

### GetEndEntityProfileName

`func (o *AddEndEntityRestRequest) GetEndEntityProfileName() string`

GetEndEntityProfileName returns the EndEntityProfileName field if non-nil, zero value otherwise.

### GetEndEntityProfileNameOk

`func (o *AddEndEntityRestRequest) GetEndEntityProfileNameOk() (*string, bool)`

GetEndEntityProfileNameOk returns a tuple with the EndEntityProfileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndEntityProfileName

`func (o *AddEndEntityRestRequest) SetEndEntityProfileName(v string)`

SetEndEntityProfileName sets EndEntityProfileName field to given value.

### HasEndEntityProfileName

`func (o *AddEndEntityRestRequest) HasEndEntityProfileName() bool`

HasEndEntityProfileName returns a boolean if a field has been set.

### GetToken

`func (o *AddEndEntityRestRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *AddEndEntityRestRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *AddEndEntityRestRequest) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *AddEndEntityRestRequest) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetAccountBindingId

`func (o *AddEndEntityRestRequest) GetAccountBindingId() string`

GetAccountBindingId returns the AccountBindingId field if non-nil, zero value otherwise.

### GetAccountBindingIdOk

`func (o *AddEndEntityRestRequest) GetAccountBindingIdOk() (*string, bool)`

GetAccountBindingIdOk returns a tuple with the AccountBindingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountBindingId

`func (o *AddEndEntityRestRequest) SetAccountBindingId(v string)`

SetAccountBindingId sets AccountBindingId field to given value.

### HasAccountBindingId

`func (o *AddEndEntityRestRequest) HasAccountBindingId() bool`

HasAccountBindingId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


