# EndEntityRestResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Username** | Pointer to **string** | Username | [optional] 
**Dn** | Pointer to **string** | Subject Distinguished Name | [optional] 
**SubjectAltName** | Pointer to **string** | Subject Alternative Name (SAN) | [optional] 
**Email** | Pointer to **string** | Email | [optional] 
**Status** | Pointer to **string** | End Entity status | [optional] 
**Token** | Pointer to **string** | Token type | [optional] 
**ExtensionData** | Pointer to [**[]ExtendedInformationRestResponseComponent**](ExtendedInformationRestResponseComponent.md) | Extended Information | [optional] 

## Methods

### NewEndEntityRestResponse

`func NewEndEntityRestResponse() *EndEntityRestResponse`

NewEndEntityRestResponse instantiates a new EndEntityRestResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEndEntityRestResponseWithDefaults

`func NewEndEntityRestResponseWithDefaults() *EndEntityRestResponse`

NewEndEntityRestResponseWithDefaults instantiates a new EndEntityRestResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsername

`func (o *EndEntityRestResponse) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *EndEntityRestResponse) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *EndEntityRestResponse) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *EndEntityRestResponse) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetDn

`func (o *EndEntityRestResponse) GetDn() string`

GetDn returns the Dn field if non-nil, zero value otherwise.

### GetDnOk

`func (o *EndEntityRestResponse) GetDnOk() (*string, bool)`

GetDnOk returns a tuple with the Dn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDn

`func (o *EndEntityRestResponse) SetDn(v string)`

SetDn sets Dn field to given value.

### HasDn

`func (o *EndEntityRestResponse) HasDn() bool`

HasDn returns a boolean if a field has been set.

### GetSubjectAltName

`func (o *EndEntityRestResponse) GetSubjectAltName() string`

GetSubjectAltName returns the SubjectAltName field if non-nil, zero value otherwise.

### GetSubjectAltNameOk

`func (o *EndEntityRestResponse) GetSubjectAltNameOk() (*string, bool)`

GetSubjectAltNameOk returns a tuple with the SubjectAltName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectAltName

`func (o *EndEntityRestResponse) SetSubjectAltName(v string)`

SetSubjectAltName sets SubjectAltName field to given value.

### HasSubjectAltName

`func (o *EndEntityRestResponse) HasSubjectAltName() bool`

HasSubjectAltName returns a boolean if a field has been set.

### GetEmail

`func (o *EndEntityRestResponse) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *EndEntityRestResponse) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *EndEntityRestResponse) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *EndEntityRestResponse) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetStatus

`func (o *EndEntityRestResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *EndEntityRestResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *EndEntityRestResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *EndEntityRestResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetToken

`func (o *EndEntityRestResponse) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *EndEntityRestResponse) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *EndEntityRestResponse) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *EndEntityRestResponse) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetExtensionData

`func (o *EndEntityRestResponse) GetExtensionData() []ExtendedInformationRestResponseComponent`

GetExtensionData returns the ExtensionData field if non-nil, zero value otherwise.

### GetExtensionDataOk

`func (o *EndEntityRestResponse) GetExtensionDataOk() (*[]ExtendedInformationRestResponseComponent, bool)`

GetExtensionDataOk returns a tuple with the ExtensionData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionData

`func (o *EndEntityRestResponse) SetExtensionData(v []ExtendedInformationRestResponseComponent)`

SetExtensionData sets ExtensionData field to given value.

### HasExtensionData

`func (o *EndEntityRestResponse) HasExtensionData() bool`

HasExtensionData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


