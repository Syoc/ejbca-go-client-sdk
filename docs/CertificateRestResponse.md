# CertificateRestResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Certificate** | Pointer to **string** | Certificate | [optional] 
**SerialNumber** | Pointer to **string** | Hex Serial Number | [optional] 
**ResponseFormat** | Pointer to **string** | Response format | [optional] 
**CertificateChain** | Pointer to **[]string** | Certificate chain | [optional] 
**CertificateProfile** | Pointer to **string** | Certificate profile name | [optional] 
**EndEntityProfile** | Pointer to **string** | End Entity profile name | [optional] 

## Methods

### NewCertificateRestResponse

`func NewCertificateRestResponse() *CertificateRestResponse`

NewCertificateRestResponse instantiates a new CertificateRestResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertificateRestResponseWithDefaults

`func NewCertificateRestResponseWithDefaults() *CertificateRestResponse`

NewCertificateRestResponseWithDefaults instantiates a new CertificateRestResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCertificate

`func (o *CertificateRestResponse) GetCertificate() string`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *CertificateRestResponse) GetCertificateOk() (*string, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *CertificateRestResponse) SetCertificate(v string)`

SetCertificate sets Certificate field to given value.

### HasCertificate

`func (o *CertificateRestResponse) HasCertificate() bool`

HasCertificate returns a boolean if a field has been set.

### GetSerialNumber

`func (o *CertificateRestResponse) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *CertificateRestResponse) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *CertificateRestResponse) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *CertificateRestResponse) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

### GetResponseFormat

`func (o *CertificateRestResponse) GetResponseFormat() string`

GetResponseFormat returns the ResponseFormat field if non-nil, zero value otherwise.

### GetResponseFormatOk

`func (o *CertificateRestResponse) GetResponseFormatOk() (*string, bool)`

GetResponseFormatOk returns a tuple with the ResponseFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseFormat

`func (o *CertificateRestResponse) SetResponseFormat(v string)`

SetResponseFormat sets ResponseFormat field to given value.

### HasResponseFormat

`func (o *CertificateRestResponse) HasResponseFormat() bool`

HasResponseFormat returns a boolean if a field has been set.

### GetCertificateChain

`func (o *CertificateRestResponse) GetCertificateChain() []string`

GetCertificateChain returns the CertificateChain field if non-nil, zero value otherwise.

### GetCertificateChainOk

`func (o *CertificateRestResponse) GetCertificateChainOk() (*[]string, bool)`

GetCertificateChainOk returns a tuple with the CertificateChain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateChain

`func (o *CertificateRestResponse) SetCertificateChain(v []string)`

SetCertificateChain sets CertificateChain field to given value.

### HasCertificateChain

`func (o *CertificateRestResponse) HasCertificateChain() bool`

HasCertificateChain returns a boolean if a field has been set.

### GetCertificateProfile

`func (o *CertificateRestResponse) GetCertificateProfile() string`

GetCertificateProfile returns the CertificateProfile field if non-nil, zero value otherwise.

### GetCertificateProfileOk

`func (o *CertificateRestResponse) GetCertificateProfileOk() (*string, bool)`

GetCertificateProfileOk returns a tuple with the CertificateProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateProfile

`func (o *CertificateRestResponse) SetCertificateProfile(v string)`

SetCertificateProfile sets CertificateProfile field to given value.

### HasCertificateProfile

`func (o *CertificateRestResponse) HasCertificateProfile() bool`

HasCertificateProfile returns a boolean if a field has been set.

### GetEndEntityProfile

`func (o *CertificateRestResponse) GetEndEntityProfile() string`

GetEndEntityProfile returns the EndEntityProfile field if non-nil, zero value otherwise.

### GetEndEntityProfileOk

`func (o *CertificateRestResponse) GetEndEntityProfileOk() (*string, bool)`

GetEndEntityProfileOk returns a tuple with the EndEntityProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndEntityProfile

`func (o *CertificateRestResponse) SetEndEntityProfile(v string)`

SetEndEntityProfile sets EndEntityProfile field to given value.

### HasEndEntityProfile

`func (o *CertificateRestResponse) HasEndEntityProfile() bool`

HasEndEntityProfile returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


