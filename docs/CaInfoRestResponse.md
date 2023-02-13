# CaInfoRestResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | CA identifier | [optional] 
**Name** | Pointer to **string** | Certificate Authority (CA) name | [optional] 
**SubjectDn** | Pointer to **string** | Subject Distinguished Name | [optional] 
**IssuerDn** | Pointer to **string** | Issuer Distinguished Name | [optional] 
**ExpirationDate** | Pointer to **time.Time** | Expiration date | [optional] 

## Methods

### NewCaInfoRestResponse

`func NewCaInfoRestResponse() *CaInfoRestResponse`

NewCaInfoRestResponse instantiates a new CaInfoRestResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCaInfoRestResponseWithDefaults

`func NewCaInfoRestResponseWithDefaults() *CaInfoRestResponse`

NewCaInfoRestResponseWithDefaults instantiates a new CaInfoRestResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CaInfoRestResponse) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CaInfoRestResponse) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CaInfoRestResponse) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *CaInfoRestResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *CaInfoRestResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CaInfoRestResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CaInfoRestResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CaInfoRestResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSubjectDn

`func (o *CaInfoRestResponse) GetSubjectDn() string`

GetSubjectDn returns the SubjectDn field if non-nil, zero value otherwise.

### GetSubjectDnOk

`func (o *CaInfoRestResponse) GetSubjectDnOk() (*string, bool)`

GetSubjectDnOk returns a tuple with the SubjectDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectDn

`func (o *CaInfoRestResponse) SetSubjectDn(v string)`

SetSubjectDn sets SubjectDn field to given value.

### HasSubjectDn

`func (o *CaInfoRestResponse) HasSubjectDn() bool`

HasSubjectDn returns a boolean if a field has been set.

### GetIssuerDn

`func (o *CaInfoRestResponse) GetIssuerDn() string`

GetIssuerDn returns the IssuerDn field if non-nil, zero value otherwise.

### GetIssuerDnOk

`func (o *CaInfoRestResponse) GetIssuerDnOk() (*string, bool)`

GetIssuerDnOk returns a tuple with the IssuerDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerDn

`func (o *CaInfoRestResponse) SetIssuerDn(v string)`

SetIssuerDn sets IssuerDn field to given value.

### HasIssuerDn

`func (o *CaInfoRestResponse) HasIssuerDn() bool`

HasIssuerDn returns a boolean if a field has been set.

### GetExpirationDate

`func (o *CaInfoRestResponse) GetExpirationDate() time.Time`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *CaInfoRestResponse) GetExpirationDateOk() (*time.Time, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *CaInfoRestResponse) SetExpirationDate(v time.Time)`

SetExpirationDate sets ExpirationDate field to given value.

### HasExpirationDate

`func (o *CaInfoRestResponse) HasExpirationDate() bool`

HasExpirationDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


