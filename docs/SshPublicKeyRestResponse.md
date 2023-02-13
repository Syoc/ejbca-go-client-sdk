# SshPublicKeyRestResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CaName** | Pointer to **string** | Certificate Authority (CA) name | [optional] 
**Response** | Pointer to **string** | CA’s public key | [optional] 

## Methods

### NewSshPublicKeyRestResponse

`func NewSshPublicKeyRestResponse() *SshPublicKeyRestResponse`

NewSshPublicKeyRestResponse instantiates a new SshPublicKeyRestResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSshPublicKeyRestResponseWithDefaults

`func NewSshPublicKeyRestResponseWithDefaults() *SshPublicKeyRestResponse`

NewSshPublicKeyRestResponseWithDefaults instantiates a new SshPublicKeyRestResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCaName

`func (o *SshPublicKeyRestResponse) GetCaName() string`

GetCaName returns the CaName field if non-nil, zero value otherwise.

### GetCaNameOk

`func (o *SshPublicKeyRestResponse) GetCaNameOk() (*string, bool)`

GetCaNameOk returns a tuple with the CaName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaName

`func (o *SshPublicKeyRestResponse) SetCaName(v string)`

SetCaName sets CaName field to given value.

### HasCaName

`func (o *SshPublicKeyRestResponse) HasCaName() bool`

HasCaName returns a boolean if a field has been set.

### GetResponse

`func (o *SshPublicKeyRestResponse) GetResponse() string`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *SshPublicKeyRestResponse) GetResponseOk() (*string, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *SshPublicKeyRestResponse) SetResponse(v string)`

SetResponse sets Response field to given value.

### HasResponse

`func (o *SshPublicKeyRestResponse) HasResponse() bool`

HasResponse returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


