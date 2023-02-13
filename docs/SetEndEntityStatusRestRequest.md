# SetEndEntityStatusRestRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Password** | Pointer to **string** | Password | [optional] 
**Token** | Pointer to **string** | Token type property | [optional] 
**Status** | Pointer to **string** | End entity status property | [optional] 

## Methods

### NewSetEndEntityStatusRestRequest

`func NewSetEndEntityStatusRestRequest() *SetEndEntityStatusRestRequest`

NewSetEndEntityStatusRestRequest instantiates a new SetEndEntityStatusRestRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSetEndEntityStatusRestRequestWithDefaults

`func NewSetEndEntityStatusRestRequestWithDefaults() *SetEndEntityStatusRestRequest`

NewSetEndEntityStatusRestRequestWithDefaults instantiates a new SetEndEntityStatusRestRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPassword

`func (o *SetEndEntityStatusRestRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *SetEndEntityStatusRestRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *SetEndEntityStatusRestRequest) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *SetEndEntityStatusRestRequest) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetToken

`func (o *SetEndEntityStatusRestRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *SetEndEntityStatusRestRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *SetEndEntityStatusRestRequest) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *SetEndEntityStatusRestRequest) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetStatus

`func (o *SetEndEntityStatusRestRequest) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SetEndEntityStatusRestRequest) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SetEndEntityStatusRestRequest) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SetEndEntityStatusRestRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


