# EndEntityProfileResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EndEntityProfileName** | Pointer to **string** | End Entity profile name | [optional] 
**SubjectDistinguishedNameFields** | Pointer to **[]string** | List of Subject DN Attributes | [optional] 
**SubjectAlternativeNameFields** | Pointer to **[]string** | List of Subject Alternative Name fields | [optional] 
**AvailableCertificateProfiles** | Pointer to **[]string** | List of available Certificate Profiles | [optional] 
**AvailableCas** | Pointer to **[]string** | List of available Certificate Authorities (CAs) | [optional] 

## Methods

### NewEndEntityProfileResponse

`func NewEndEntityProfileResponse() *EndEntityProfileResponse`

NewEndEntityProfileResponse instantiates a new EndEntityProfileResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEndEntityProfileResponseWithDefaults

`func NewEndEntityProfileResponseWithDefaults() *EndEntityProfileResponse`

NewEndEntityProfileResponseWithDefaults instantiates a new EndEntityProfileResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndEntityProfileName

`func (o *EndEntityProfileResponse) GetEndEntityProfileName() string`

GetEndEntityProfileName returns the EndEntityProfileName field if non-nil, zero value otherwise.

### GetEndEntityProfileNameOk

`func (o *EndEntityProfileResponse) GetEndEntityProfileNameOk() (*string, bool)`

GetEndEntityProfileNameOk returns a tuple with the EndEntityProfileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndEntityProfileName

`func (o *EndEntityProfileResponse) SetEndEntityProfileName(v string)`

SetEndEntityProfileName sets EndEntityProfileName field to given value.

### HasEndEntityProfileName

`func (o *EndEntityProfileResponse) HasEndEntityProfileName() bool`

HasEndEntityProfileName returns a boolean if a field has been set.

### GetSubjectDistinguishedNameFields

`func (o *EndEntityProfileResponse) GetSubjectDistinguishedNameFields() []string`

GetSubjectDistinguishedNameFields returns the SubjectDistinguishedNameFields field if non-nil, zero value otherwise.

### GetSubjectDistinguishedNameFieldsOk

`func (o *EndEntityProfileResponse) GetSubjectDistinguishedNameFieldsOk() (*[]string, bool)`

GetSubjectDistinguishedNameFieldsOk returns a tuple with the SubjectDistinguishedNameFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectDistinguishedNameFields

`func (o *EndEntityProfileResponse) SetSubjectDistinguishedNameFields(v []string)`

SetSubjectDistinguishedNameFields sets SubjectDistinguishedNameFields field to given value.

### HasSubjectDistinguishedNameFields

`func (o *EndEntityProfileResponse) HasSubjectDistinguishedNameFields() bool`

HasSubjectDistinguishedNameFields returns a boolean if a field has been set.

### GetSubjectAlternativeNameFields

`func (o *EndEntityProfileResponse) GetSubjectAlternativeNameFields() []string`

GetSubjectAlternativeNameFields returns the SubjectAlternativeNameFields field if non-nil, zero value otherwise.

### GetSubjectAlternativeNameFieldsOk

`func (o *EndEntityProfileResponse) GetSubjectAlternativeNameFieldsOk() (*[]string, bool)`

GetSubjectAlternativeNameFieldsOk returns a tuple with the SubjectAlternativeNameFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectAlternativeNameFields

`func (o *EndEntityProfileResponse) SetSubjectAlternativeNameFields(v []string)`

SetSubjectAlternativeNameFields sets SubjectAlternativeNameFields field to given value.

### HasSubjectAlternativeNameFields

`func (o *EndEntityProfileResponse) HasSubjectAlternativeNameFields() bool`

HasSubjectAlternativeNameFields returns a boolean if a field has been set.

### GetAvailableCertificateProfiles

`func (o *EndEntityProfileResponse) GetAvailableCertificateProfiles() []string`

GetAvailableCertificateProfiles returns the AvailableCertificateProfiles field if non-nil, zero value otherwise.

### GetAvailableCertificateProfilesOk

`func (o *EndEntityProfileResponse) GetAvailableCertificateProfilesOk() (*[]string, bool)`

GetAvailableCertificateProfilesOk returns a tuple with the AvailableCertificateProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableCertificateProfiles

`func (o *EndEntityProfileResponse) SetAvailableCertificateProfiles(v []string)`

SetAvailableCertificateProfiles sets AvailableCertificateProfiles field to given value.

### HasAvailableCertificateProfiles

`func (o *EndEntityProfileResponse) HasAvailableCertificateProfiles() bool`

HasAvailableCertificateProfiles returns a boolean if a field has been set.

### GetAvailableCas

`func (o *EndEntityProfileResponse) GetAvailableCas() []string`

GetAvailableCas returns the AvailableCas field if non-nil, zero value otherwise.

### GetAvailableCasOk

`func (o *EndEntityProfileResponse) GetAvailableCasOk() (*[]string, bool)`

GetAvailableCasOk returns a tuple with the AvailableCas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableCas

`func (o *EndEntityProfileResponse) SetAvailableCas(v []string)`

SetAvailableCas sets AvailableCas field to given value.

### HasAvailableCas

`func (o *EndEntityProfileResponse) HasAvailableCas() bool`

HasAvailableCas returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


