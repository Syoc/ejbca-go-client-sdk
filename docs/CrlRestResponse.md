# CrlRestResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Crl** | Pointer to **string** | Certificate Revokation List (CRL) | [optional] 
**ResponseFormat** | Pointer to **string** | Response format | [optional] 

## Methods

### NewCrlRestResponse

`func NewCrlRestResponse() *CrlRestResponse`

NewCrlRestResponse instantiates a new CrlRestResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCrlRestResponseWithDefaults

`func NewCrlRestResponseWithDefaults() *CrlRestResponse`

NewCrlRestResponseWithDefaults instantiates a new CrlRestResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCrl

`func (o *CrlRestResponse) GetCrl() string`

GetCrl returns the Crl field if non-nil, zero value otherwise.

### GetCrlOk

`func (o *CrlRestResponse) GetCrlOk() (*string, bool)`

GetCrlOk returns a tuple with the Crl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrl

`func (o *CrlRestResponse) SetCrl(v string)`

SetCrl sets Crl field to given value.

### HasCrl

`func (o *CrlRestResponse) HasCrl() bool`

HasCrl returns a boolean if a field has been set.

### GetResponseFormat

`func (o *CrlRestResponse) GetResponseFormat() string`

GetResponseFormat returns the ResponseFormat field if non-nil, zero value otherwise.

### GetResponseFormatOk

`func (o *CrlRestResponse) GetResponseFormatOk() (*string, bool)`

GetResponseFormatOk returns a tuple with the ResponseFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseFormat

`func (o *CrlRestResponse) SetResponseFormat(v string)`

SetResponseFormat sets ResponseFormat field to given value.

### HasResponseFormat

`func (o *CrlRestResponse) HasResponseFormat() bool`

HasResponseFormat returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


