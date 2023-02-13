# EnrollCertificateRestRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CertificateRequest** | Pointer to **string** | Certificate request | [optional] 
**CertificateProfileName** | Pointer to **string** | Certificate profile name | [optional] 
**EndEntityProfileName** | Pointer to **string** | End Entity profile name | [optional] 
**CertificateAuthorityName** | Pointer to **string** | Certificate Authority (CA) name | [optional] 
**Username** | Pointer to **string** | Username | [optional] 
**Password** | Pointer to **string** | Password | [optional] 
**AccountBindingId** | Pointer to **string** | Account Binding ID | [optional] 
**IncludeChain** | Pointer to **bool** |  | [optional] 
**Email** | Pointer to **string** | Email | [optional] 

## Methods

### NewEnrollCertificateRestRequest

`func NewEnrollCertificateRestRequest() *EnrollCertificateRestRequest`

NewEnrollCertificateRestRequest instantiates a new EnrollCertificateRestRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnrollCertificateRestRequestWithDefaults

`func NewEnrollCertificateRestRequestWithDefaults() *EnrollCertificateRestRequest`

NewEnrollCertificateRestRequestWithDefaults instantiates a new EnrollCertificateRestRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCertificateRequest

`func (o *EnrollCertificateRestRequest) GetCertificateRequest() string`

GetCertificateRequest returns the CertificateRequest field if non-nil, zero value otherwise.

### GetCertificateRequestOk

`func (o *EnrollCertificateRestRequest) GetCertificateRequestOk() (*string, bool)`

GetCertificateRequestOk returns a tuple with the CertificateRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateRequest

`func (o *EnrollCertificateRestRequest) SetCertificateRequest(v string)`

SetCertificateRequest sets CertificateRequest field to given value.

### HasCertificateRequest

`func (o *EnrollCertificateRestRequest) HasCertificateRequest() bool`

HasCertificateRequest returns a boolean if a field has been set.

### GetCertificateProfileName

`func (o *EnrollCertificateRestRequest) GetCertificateProfileName() string`

GetCertificateProfileName returns the CertificateProfileName field if non-nil, zero value otherwise.

### GetCertificateProfileNameOk

`func (o *EnrollCertificateRestRequest) GetCertificateProfileNameOk() (*string, bool)`

GetCertificateProfileNameOk returns a tuple with the CertificateProfileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateProfileName

`func (o *EnrollCertificateRestRequest) SetCertificateProfileName(v string)`

SetCertificateProfileName sets CertificateProfileName field to given value.

### HasCertificateProfileName

`func (o *EnrollCertificateRestRequest) HasCertificateProfileName() bool`

HasCertificateProfileName returns a boolean if a field has been set.

### GetEndEntityProfileName

`func (o *EnrollCertificateRestRequest) GetEndEntityProfileName() string`

GetEndEntityProfileName returns the EndEntityProfileName field if non-nil, zero value otherwise.

### GetEndEntityProfileNameOk

`func (o *EnrollCertificateRestRequest) GetEndEntityProfileNameOk() (*string, bool)`

GetEndEntityProfileNameOk returns a tuple with the EndEntityProfileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndEntityProfileName

`func (o *EnrollCertificateRestRequest) SetEndEntityProfileName(v string)`

SetEndEntityProfileName sets EndEntityProfileName field to given value.

### HasEndEntityProfileName

`func (o *EnrollCertificateRestRequest) HasEndEntityProfileName() bool`

HasEndEntityProfileName returns a boolean if a field has been set.

### GetCertificateAuthorityName

`func (o *EnrollCertificateRestRequest) GetCertificateAuthorityName() string`

GetCertificateAuthorityName returns the CertificateAuthorityName field if non-nil, zero value otherwise.

### GetCertificateAuthorityNameOk

`func (o *EnrollCertificateRestRequest) GetCertificateAuthorityNameOk() (*string, bool)`

GetCertificateAuthorityNameOk returns a tuple with the CertificateAuthorityName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateAuthorityName

`func (o *EnrollCertificateRestRequest) SetCertificateAuthorityName(v string)`

SetCertificateAuthorityName sets CertificateAuthorityName field to given value.

### HasCertificateAuthorityName

`func (o *EnrollCertificateRestRequest) HasCertificateAuthorityName() bool`

HasCertificateAuthorityName returns a boolean if a field has been set.

### GetUsername

`func (o *EnrollCertificateRestRequest) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *EnrollCertificateRestRequest) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *EnrollCertificateRestRequest) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *EnrollCertificateRestRequest) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetPassword

`func (o *EnrollCertificateRestRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *EnrollCertificateRestRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *EnrollCertificateRestRequest) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *EnrollCertificateRestRequest) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetAccountBindingId

`func (o *EnrollCertificateRestRequest) GetAccountBindingId() string`

GetAccountBindingId returns the AccountBindingId field if non-nil, zero value otherwise.

### GetAccountBindingIdOk

`func (o *EnrollCertificateRestRequest) GetAccountBindingIdOk() (*string, bool)`

GetAccountBindingIdOk returns a tuple with the AccountBindingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountBindingId

`func (o *EnrollCertificateRestRequest) SetAccountBindingId(v string)`

SetAccountBindingId sets AccountBindingId field to given value.

### HasAccountBindingId

`func (o *EnrollCertificateRestRequest) HasAccountBindingId() bool`

HasAccountBindingId returns a boolean if a field has been set.

### GetIncludeChain

`func (o *EnrollCertificateRestRequest) GetIncludeChain() bool`

GetIncludeChain returns the IncludeChain field if non-nil, zero value otherwise.

### GetIncludeChainOk

`func (o *EnrollCertificateRestRequest) GetIncludeChainOk() (*bool, bool)`

GetIncludeChainOk returns a tuple with the IncludeChain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeChain

`func (o *EnrollCertificateRestRequest) SetIncludeChain(v bool)`

SetIncludeChain sets IncludeChain field to given value.

### HasIncludeChain

`func (o *EnrollCertificateRestRequest) HasIncludeChain() bool`

HasIncludeChain returns a boolean if a field has been set.

### GetEmail

`func (o *EnrollCertificateRestRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *EnrollCertificateRestRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *EnrollCertificateRestRequest) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *EnrollCertificateRestRequest) HasEmail() bool`

HasEmail returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


