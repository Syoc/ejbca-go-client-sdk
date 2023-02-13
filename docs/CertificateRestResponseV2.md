# CertificateRestResponseV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Fingerprint** | Pointer to **string** | Certificate fingerprint | [optional] [readonly] 
**CAFingerprint** | Pointer to **string** | Certificate Authority fingerprint | [optional] [readonly] 
**CertificateProfileId** | Pointer to **int32** | Certificate Profile Identifier | [optional] [readonly] 
**EndEntityProfileId** | Pointer to **int32** | End Entity Profile Identifier | [optional] [readonly] 
**ExpireDate** | Pointer to **int64** | Date after which certificate should be considered expired | [optional] [readonly] 
**IssuerDN** | Pointer to **string** | Issuer Distinguished Name | [optional] [readonly] 
**NotBefore** | Pointer to **int64** | Date at which certificate became valid | [optional] [readonly] 
**RevocationDate** | Pointer to **int64** | Revocation date | [optional] [readonly] 
**RevocationReason** | Pointer to **int32** | Revocation reson | [optional] [readonly] 
**SerialNumber** | Pointer to **string** | Hex Serial Number | [optional] [readonly] 
**Status** | Pointer to **int32** | Certificate status | [optional] [readonly] 
**SubjectAltName** | Pointer to **string** | Subject Alternative Name (SAN) | [optional] [readonly] 
**SubjectDN** | Pointer to **string** | Subject Distinguished Name | [optional] [readonly] 
**SubjectKeyId** | Pointer to **string** | Subject Key Identifier | [optional] [readonly] 
**Tag** | Pointer to **string** |  | [optional] [readonly] 
**Type** | Pointer to **int32** |  | [optional] [readonly] 
**UdpateTime** | Pointer to **int64** | Update time | [optional] [readonly] 
**Username** | Pointer to **string** | Username | [optional] [readonly] 
**Base64Cert** | Pointer to **string** | Base64 encoded certificate | [optional] [readonly] 
**CertificateRequest** | Pointer to **string** | Certificate request | [optional] [readonly] 
**CrlPartitionIndex** | Pointer to **int32** | CRL partition index | [optional] [readonly] 

## Methods

### NewCertificateRestResponseV2

`func NewCertificateRestResponseV2() *CertificateRestResponseV2`

NewCertificateRestResponseV2 instantiates a new CertificateRestResponseV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertificateRestResponseV2WithDefaults

`func NewCertificateRestResponseV2WithDefaults() *CertificateRestResponseV2`

NewCertificateRestResponseV2WithDefaults instantiates a new CertificateRestResponseV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFingerprint

`func (o *CertificateRestResponseV2) GetFingerprint() string`

GetFingerprint returns the Fingerprint field if non-nil, zero value otherwise.

### GetFingerprintOk

`func (o *CertificateRestResponseV2) GetFingerprintOk() (*string, bool)`

GetFingerprintOk returns a tuple with the Fingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFingerprint

`func (o *CertificateRestResponseV2) SetFingerprint(v string)`

SetFingerprint sets Fingerprint field to given value.

### HasFingerprint

`func (o *CertificateRestResponseV2) HasFingerprint() bool`

HasFingerprint returns a boolean if a field has been set.

### GetCAFingerprint

`func (o *CertificateRestResponseV2) GetCAFingerprint() string`

GetCAFingerprint returns the CAFingerprint field if non-nil, zero value otherwise.

### GetCAFingerprintOk

`func (o *CertificateRestResponseV2) GetCAFingerprintOk() (*string, bool)`

GetCAFingerprintOk returns a tuple with the CAFingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCAFingerprint

`func (o *CertificateRestResponseV2) SetCAFingerprint(v string)`

SetCAFingerprint sets CAFingerprint field to given value.

### HasCAFingerprint

`func (o *CertificateRestResponseV2) HasCAFingerprint() bool`

HasCAFingerprint returns a boolean if a field has been set.

### GetCertificateProfileId

`func (o *CertificateRestResponseV2) GetCertificateProfileId() int32`

GetCertificateProfileId returns the CertificateProfileId field if non-nil, zero value otherwise.

### GetCertificateProfileIdOk

`func (o *CertificateRestResponseV2) GetCertificateProfileIdOk() (*int32, bool)`

GetCertificateProfileIdOk returns a tuple with the CertificateProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateProfileId

`func (o *CertificateRestResponseV2) SetCertificateProfileId(v int32)`

SetCertificateProfileId sets CertificateProfileId field to given value.

### HasCertificateProfileId

`func (o *CertificateRestResponseV2) HasCertificateProfileId() bool`

HasCertificateProfileId returns a boolean if a field has been set.

### GetEndEntityProfileId

`func (o *CertificateRestResponseV2) GetEndEntityProfileId() int32`

GetEndEntityProfileId returns the EndEntityProfileId field if non-nil, zero value otherwise.

### GetEndEntityProfileIdOk

`func (o *CertificateRestResponseV2) GetEndEntityProfileIdOk() (*int32, bool)`

GetEndEntityProfileIdOk returns a tuple with the EndEntityProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndEntityProfileId

`func (o *CertificateRestResponseV2) SetEndEntityProfileId(v int32)`

SetEndEntityProfileId sets EndEntityProfileId field to given value.

### HasEndEntityProfileId

`func (o *CertificateRestResponseV2) HasEndEntityProfileId() bool`

HasEndEntityProfileId returns a boolean if a field has been set.

### GetExpireDate

`func (o *CertificateRestResponseV2) GetExpireDate() int64`

GetExpireDate returns the ExpireDate field if non-nil, zero value otherwise.

### GetExpireDateOk

`func (o *CertificateRestResponseV2) GetExpireDateOk() (*int64, bool)`

GetExpireDateOk returns a tuple with the ExpireDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpireDate

`func (o *CertificateRestResponseV2) SetExpireDate(v int64)`

SetExpireDate sets ExpireDate field to given value.

### HasExpireDate

`func (o *CertificateRestResponseV2) HasExpireDate() bool`

HasExpireDate returns a boolean if a field has been set.

### GetIssuerDN

`func (o *CertificateRestResponseV2) GetIssuerDN() string`

GetIssuerDN returns the IssuerDN field if non-nil, zero value otherwise.

### GetIssuerDNOk

`func (o *CertificateRestResponseV2) GetIssuerDNOk() (*string, bool)`

GetIssuerDNOk returns a tuple with the IssuerDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerDN

`func (o *CertificateRestResponseV2) SetIssuerDN(v string)`

SetIssuerDN sets IssuerDN field to given value.

### HasIssuerDN

`func (o *CertificateRestResponseV2) HasIssuerDN() bool`

HasIssuerDN returns a boolean if a field has been set.

### GetNotBefore

`func (o *CertificateRestResponseV2) GetNotBefore() int64`

GetNotBefore returns the NotBefore field if non-nil, zero value otherwise.

### GetNotBeforeOk

`func (o *CertificateRestResponseV2) GetNotBeforeOk() (*int64, bool)`

GetNotBeforeOk returns a tuple with the NotBefore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotBefore

`func (o *CertificateRestResponseV2) SetNotBefore(v int64)`

SetNotBefore sets NotBefore field to given value.

### HasNotBefore

`func (o *CertificateRestResponseV2) HasNotBefore() bool`

HasNotBefore returns a boolean if a field has been set.

### GetRevocationDate

`func (o *CertificateRestResponseV2) GetRevocationDate() int64`

GetRevocationDate returns the RevocationDate field if non-nil, zero value otherwise.

### GetRevocationDateOk

`func (o *CertificateRestResponseV2) GetRevocationDateOk() (*int64, bool)`

GetRevocationDateOk returns a tuple with the RevocationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevocationDate

`func (o *CertificateRestResponseV2) SetRevocationDate(v int64)`

SetRevocationDate sets RevocationDate field to given value.

### HasRevocationDate

`func (o *CertificateRestResponseV2) HasRevocationDate() bool`

HasRevocationDate returns a boolean if a field has been set.

### GetRevocationReason

`func (o *CertificateRestResponseV2) GetRevocationReason() int32`

GetRevocationReason returns the RevocationReason field if non-nil, zero value otherwise.

### GetRevocationReasonOk

`func (o *CertificateRestResponseV2) GetRevocationReasonOk() (*int32, bool)`

GetRevocationReasonOk returns a tuple with the RevocationReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevocationReason

`func (o *CertificateRestResponseV2) SetRevocationReason(v int32)`

SetRevocationReason sets RevocationReason field to given value.

### HasRevocationReason

`func (o *CertificateRestResponseV2) HasRevocationReason() bool`

HasRevocationReason returns a boolean if a field has been set.

### GetSerialNumber

`func (o *CertificateRestResponseV2) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *CertificateRestResponseV2) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *CertificateRestResponseV2) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *CertificateRestResponseV2) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

### GetStatus

`func (o *CertificateRestResponseV2) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CertificateRestResponseV2) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CertificateRestResponseV2) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CertificateRestResponseV2) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSubjectAltName

`func (o *CertificateRestResponseV2) GetSubjectAltName() string`

GetSubjectAltName returns the SubjectAltName field if non-nil, zero value otherwise.

### GetSubjectAltNameOk

`func (o *CertificateRestResponseV2) GetSubjectAltNameOk() (*string, bool)`

GetSubjectAltNameOk returns a tuple with the SubjectAltName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectAltName

`func (o *CertificateRestResponseV2) SetSubjectAltName(v string)`

SetSubjectAltName sets SubjectAltName field to given value.

### HasSubjectAltName

`func (o *CertificateRestResponseV2) HasSubjectAltName() bool`

HasSubjectAltName returns a boolean if a field has been set.

### GetSubjectDN

`func (o *CertificateRestResponseV2) GetSubjectDN() string`

GetSubjectDN returns the SubjectDN field if non-nil, zero value otherwise.

### GetSubjectDNOk

`func (o *CertificateRestResponseV2) GetSubjectDNOk() (*string, bool)`

GetSubjectDNOk returns a tuple with the SubjectDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectDN

`func (o *CertificateRestResponseV2) SetSubjectDN(v string)`

SetSubjectDN sets SubjectDN field to given value.

### HasSubjectDN

`func (o *CertificateRestResponseV2) HasSubjectDN() bool`

HasSubjectDN returns a boolean if a field has been set.

### GetSubjectKeyId

`func (o *CertificateRestResponseV2) GetSubjectKeyId() string`

GetSubjectKeyId returns the SubjectKeyId field if non-nil, zero value otherwise.

### GetSubjectKeyIdOk

`func (o *CertificateRestResponseV2) GetSubjectKeyIdOk() (*string, bool)`

GetSubjectKeyIdOk returns a tuple with the SubjectKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectKeyId

`func (o *CertificateRestResponseV2) SetSubjectKeyId(v string)`

SetSubjectKeyId sets SubjectKeyId field to given value.

### HasSubjectKeyId

`func (o *CertificateRestResponseV2) HasSubjectKeyId() bool`

HasSubjectKeyId returns a boolean if a field has been set.

### GetTag

`func (o *CertificateRestResponseV2) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *CertificateRestResponseV2) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *CertificateRestResponseV2) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *CertificateRestResponseV2) HasTag() bool`

HasTag returns a boolean if a field has been set.

### GetType

`func (o *CertificateRestResponseV2) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CertificateRestResponseV2) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CertificateRestResponseV2) SetType(v int32)`

SetType sets Type field to given value.

### HasType

`func (o *CertificateRestResponseV2) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUdpateTime

`func (o *CertificateRestResponseV2) GetUdpateTime() int64`

GetUdpateTime returns the UdpateTime field if non-nil, zero value otherwise.

### GetUdpateTimeOk

`func (o *CertificateRestResponseV2) GetUdpateTimeOk() (*int64, bool)`

GetUdpateTimeOk returns a tuple with the UdpateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUdpateTime

`func (o *CertificateRestResponseV2) SetUdpateTime(v int64)`

SetUdpateTime sets UdpateTime field to given value.

### HasUdpateTime

`func (o *CertificateRestResponseV2) HasUdpateTime() bool`

HasUdpateTime returns a boolean if a field has been set.

### GetUsername

`func (o *CertificateRestResponseV2) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *CertificateRestResponseV2) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *CertificateRestResponseV2) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *CertificateRestResponseV2) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetBase64Cert

`func (o *CertificateRestResponseV2) GetBase64Cert() string`

GetBase64Cert returns the Base64Cert field if non-nil, zero value otherwise.

### GetBase64CertOk

`func (o *CertificateRestResponseV2) GetBase64CertOk() (*string, bool)`

GetBase64CertOk returns a tuple with the Base64Cert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBase64Cert

`func (o *CertificateRestResponseV2) SetBase64Cert(v string)`

SetBase64Cert sets Base64Cert field to given value.

### HasBase64Cert

`func (o *CertificateRestResponseV2) HasBase64Cert() bool`

HasBase64Cert returns a boolean if a field has been set.

### GetCertificateRequest

`func (o *CertificateRestResponseV2) GetCertificateRequest() string`

GetCertificateRequest returns the CertificateRequest field if non-nil, zero value otherwise.

### GetCertificateRequestOk

`func (o *CertificateRestResponseV2) GetCertificateRequestOk() (*string, bool)`

GetCertificateRequestOk returns a tuple with the CertificateRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateRequest

`func (o *CertificateRestResponseV2) SetCertificateRequest(v string)`

SetCertificateRequest sets CertificateRequest field to given value.

### HasCertificateRequest

`func (o *CertificateRestResponseV2) HasCertificateRequest() bool`

HasCertificateRequest returns a boolean if a field has been set.

### GetCrlPartitionIndex

`func (o *CertificateRestResponseV2) GetCrlPartitionIndex() int32`

GetCrlPartitionIndex returns the CrlPartitionIndex field if non-nil, zero value otherwise.

### GetCrlPartitionIndexOk

`func (o *CertificateRestResponseV2) GetCrlPartitionIndexOk() (*int32, bool)`

GetCrlPartitionIndexOk returns a tuple with the CrlPartitionIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrlPartitionIndex

`func (o *CertificateRestResponseV2) SetCrlPartitionIndex(v int32)`

SetCrlPartitionIndex sets CrlPartitionIndex field to given value.

### HasCrlPartitionIndex

`func (o *CertificateRestResponseV2) HasCrlPartitionIndex() bool`

HasCrlPartitionIndex returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


