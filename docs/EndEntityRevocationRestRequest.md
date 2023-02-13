# EndEntityRevocationRestRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReasonCode** | Pointer to **int32** | Reason code | [optional] 
**Delete** | Pointer to **bool** | Delete | [optional] 

## Methods

### NewEndEntityRevocationRestRequest

`func NewEndEntityRevocationRestRequest() *EndEntityRevocationRestRequest`

NewEndEntityRevocationRestRequest instantiates a new EndEntityRevocationRestRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEndEntityRevocationRestRequestWithDefaults

`func NewEndEntityRevocationRestRequestWithDefaults() *EndEntityRevocationRestRequest`

NewEndEntityRevocationRestRequestWithDefaults instantiates a new EndEntityRevocationRestRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReasonCode

`func (o *EndEntityRevocationRestRequest) GetReasonCode() int32`

GetReasonCode returns the ReasonCode field if non-nil, zero value otherwise.

### GetReasonCodeOk

`func (o *EndEntityRevocationRestRequest) GetReasonCodeOk() (*int32, bool)`

GetReasonCodeOk returns a tuple with the ReasonCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasonCode

`func (o *EndEntityRevocationRestRequest) SetReasonCode(v int32)`

SetReasonCode sets ReasonCode field to given value.

### HasReasonCode

`func (o *EndEntityRevocationRestRequest) HasReasonCode() bool`

HasReasonCode returns a boolean if a field has been set.

### GetDelete

`func (o *EndEntityRevocationRestRequest) GetDelete() bool`

GetDelete returns the Delete field if non-nil, zero value otherwise.

### GetDeleteOk

`func (o *EndEntityRevocationRestRequest) GetDeleteOk() (*bool, bool)`

GetDeleteOk returns a tuple with the Delete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelete

`func (o *EndEntityRevocationRestRequest) SetDelete(v bool)`

SetDelete sets Delete field to given value.

### HasDelete

`func (o *EndEntityRevocationRestRequest) HasDelete() bool`

HasDelete returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


