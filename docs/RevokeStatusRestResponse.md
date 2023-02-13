# RevokeStatusRestResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IssuerDn** | Pointer to **string** | Issuer Distinguished Name | [optional] 
**SerialNumber** | Pointer to **string** | Hex Serial Number | [optional] 
**RevocationReason** | Pointer to **string** | RFC5280 revokation reason | [optional] 
**RevocationDate** | Pointer to **time.Time** | Revokation date | [optional] 
**Message** | Pointer to **string** | Message | [optional] 
**Revoked** | Pointer to **bool** |  | [optional] 

## Methods

### NewRevokeStatusRestResponse

`func NewRevokeStatusRestResponse() *RevokeStatusRestResponse`

NewRevokeStatusRestResponse instantiates a new RevokeStatusRestResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRevokeStatusRestResponseWithDefaults

`func NewRevokeStatusRestResponseWithDefaults() *RevokeStatusRestResponse`

NewRevokeStatusRestResponseWithDefaults instantiates a new RevokeStatusRestResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIssuerDn

`func (o *RevokeStatusRestResponse) GetIssuerDn() string`

GetIssuerDn returns the IssuerDn field if non-nil, zero value otherwise.

### GetIssuerDnOk

`func (o *RevokeStatusRestResponse) GetIssuerDnOk() (*string, bool)`

GetIssuerDnOk returns a tuple with the IssuerDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerDn

`func (o *RevokeStatusRestResponse) SetIssuerDn(v string)`

SetIssuerDn sets IssuerDn field to given value.

### HasIssuerDn

`func (o *RevokeStatusRestResponse) HasIssuerDn() bool`

HasIssuerDn returns a boolean if a field has been set.

### GetSerialNumber

`func (o *RevokeStatusRestResponse) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *RevokeStatusRestResponse) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *RevokeStatusRestResponse) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *RevokeStatusRestResponse) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

### GetRevocationReason

`func (o *RevokeStatusRestResponse) GetRevocationReason() string`

GetRevocationReason returns the RevocationReason field if non-nil, zero value otherwise.

### GetRevocationReasonOk

`func (o *RevokeStatusRestResponse) GetRevocationReasonOk() (*string, bool)`

GetRevocationReasonOk returns a tuple with the RevocationReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevocationReason

`func (o *RevokeStatusRestResponse) SetRevocationReason(v string)`

SetRevocationReason sets RevocationReason field to given value.

### HasRevocationReason

`func (o *RevokeStatusRestResponse) HasRevocationReason() bool`

HasRevocationReason returns a boolean if a field has been set.

### GetRevocationDate

`func (o *RevokeStatusRestResponse) GetRevocationDate() time.Time`

GetRevocationDate returns the RevocationDate field if non-nil, zero value otherwise.

### GetRevocationDateOk

`func (o *RevokeStatusRestResponse) GetRevocationDateOk() (*time.Time, bool)`

GetRevocationDateOk returns a tuple with the RevocationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevocationDate

`func (o *RevokeStatusRestResponse) SetRevocationDate(v time.Time)`

SetRevocationDate sets RevocationDate field to given value.

### HasRevocationDate

`func (o *RevokeStatusRestResponse) HasRevocationDate() bool`

HasRevocationDate returns a boolean if a field has been set.

### GetMessage

`func (o *RevokeStatusRestResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *RevokeStatusRestResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *RevokeStatusRestResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *RevokeStatusRestResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetRevoked

`func (o *RevokeStatusRestResponse) GetRevoked() bool`

GetRevoked returns the Revoked field if non-nil, zero value otherwise.

### GetRevokedOk

`func (o *RevokeStatusRestResponse) GetRevokedOk() (*bool, bool)`

GetRevokedOk returns a tuple with the Revoked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevoked

`func (o *RevokeStatusRestResponse) SetRevoked(v bool)`

SetRevoked sets Revoked field to given value.

### HasRevoked

`func (o *RevokeStatusRestResponse) HasRevoked() bool`

HasRevoked returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


