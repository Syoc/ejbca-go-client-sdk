# KeyStoreRestRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Username** | Pointer to **string** | Username | [optional] 
**Password** | Pointer to **string** | Password | [optional] 
**KeyAlg** | Pointer to **string** | Key algorithm used for enrollment | [optional] 
**KeySpec** | Pointer to **string** | Key specification to use | [optional] 

## Methods

### NewKeyStoreRestRequest

`func NewKeyStoreRestRequest() *KeyStoreRestRequest`

NewKeyStoreRestRequest instantiates a new KeyStoreRestRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeyStoreRestRequestWithDefaults

`func NewKeyStoreRestRequestWithDefaults() *KeyStoreRestRequest`

NewKeyStoreRestRequestWithDefaults instantiates a new KeyStoreRestRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsername

`func (o *KeyStoreRestRequest) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *KeyStoreRestRequest) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *KeyStoreRestRequest) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *KeyStoreRestRequest) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetPassword

`func (o *KeyStoreRestRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *KeyStoreRestRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *KeyStoreRestRequest) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *KeyStoreRestRequest) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetKeyAlg

`func (o *KeyStoreRestRequest) GetKeyAlg() string`

GetKeyAlg returns the KeyAlg field if non-nil, zero value otherwise.

### GetKeyAlgOk

`func (o *KeyStoreRestRequest) GetKeyAlgOk() (*string, bool)`

GetKeyAlgOk returns a tuple with the KeyAlg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyAlg

`func (o *KeyStoreRestRequest) SetKeyAlg(v string)`

SetKeyAlg sets KeyAlg field to given value.

### HasKeyAlg

`func (o *KeyStoreRestRequest) HasKeyAlg() bool`

HasKeyAlg returns a boolean if a field has been set.

### GetKeySpec

`func (o *KeyStoreRestRequest) GetKeySpec() string`

GetKeySpec returns the KeySpec field if non-nil, zero value otherwise.

### GetKeySpecOk

`func (o *KeyStoreRestRequest) GetKeySpecOk() (*string, bool)`

GetKeySpecOk returns a tuple with the KeySpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeySpec

`func (o *KeyStoreRestRequest) SetKeySpec(v string)`

SetKeySpec sets KeySpec field to given value.

### HasKeySpec

`func (o *KeyStoreRestRequest) HasKeySpec() bool`

HasKeySpec returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


