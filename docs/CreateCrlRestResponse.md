# CreateCrlRestResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IssuerDn** | Pointer to **string** | Issuer Distinguished Name | [optional] 
**LatestCrlVersion** | Pointer to **int32** | Latest base CRL version | [optional] 
**LatestDeltaCrlVersion** | Pointer to **int32** | Latest delta CRL version | [optional] 
**LatestPartitionCrlVersions** | Pointer to **map[string]int32** |  | [optional] 
**LatestPartitionDeltaCrlVersions** | Pointer to **map[string]int32** |  | [optional] 
**AllSuccess** | Pointer to **bool** |  | [optional] 

## Methods

### NewCreateCrlRestResponse

`func NewCreateCrlRestResponse() *CreateCrlRestResponse`

NewCreateCrlRestResponse instantiates a new CreateCrlRestResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateCrlRestResponseWithDefaults

`func NewCreateCrlRestResponseWithDefaults() *CreateCrlRestResponse`

NewCreateCrlRestResponseWithDefaults instantiates a new CreateCrlRestResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIssuerDn

`func (o *CreateCrlRestResponse) GetIssuerDn() string`

GetIssuerDn returns the IssuerDn field if non-nil, zero value otherwise.

### GetIssuerDnOk

`func (o *CreateCrlRestResponse) GetIssuerDnOk() (*string, bool)`

GetIssuerDnOk returns a tuple with the IssuerDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerDn

`func (o *CreateCrlRestResponse) SetIssuerDn(v string)`

SetIssuerDn sets IssuerDn field to given value.

### HasIssuerDn

`func (o *CreateCrlRestResponse) HasIssuerDn() bool`

HasIssuerDn returns a boolean if a field has been set.

### GetLatestCrlVersion

`func (o *CreateCrlRestResponse) GetLatestCrlVersion() int32`

GetLatestCrlVersion returns the LatestCrlVersion field if non-nil, zero value otherwise.

### GetLatestCrlVersionOk

`func (o *CreateCrlRestResponse) GetLatestCrlVersionOk() (*int32, bool)`

GetLatestCrlVersionOk returns a tuple with the LatestCrlVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestCrlVersion

`func (o *CreateCrlRestResponse) SetLatestCrlVersion(v int32)`

SetLatestCrlVersion sets LatestCrlVersion field to given value.

### HasLatestCrlVersion

`func (o *CreateCrlRestResponse) HasLatestCrlVersion() bool`

HasLatestCrlVersion returns a boolean if a field has been set.

### GetLatestDeltaCrlVersion

`func (o *CreateCrlRestResponse) GetLatestDeltaCrlVersion() int32`

GetLatestDeltaCrlVersion returns the LatestDeltaCrlVersion field if non-nil, zero value otherwise.

### GetLatestDeltaCrlVersionOk

`func (o *CreateCrlRestResponse) GetLatestDeltaCrlVersionOk() (*int32, bool)`

GetLatestDeltaCrlVersionOk returns a tuple with the LatestDeltaCrlVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestDeltaCrlVersion

`func (o *CreateCrlRestResponse) SetLatestDeltaCrlVersion(v int32)`

SetLatestDeltaCrlVersion sets LatestDeltaCrlVersion field to given value.

### HasLatestDeltaCrlVersion

`func (o *CreateCrlRestResponse) HasLatestDeltaCrlVersion() bool`

HasLatestDeltaCrlVersion returns a boolean if a field has been set.

### GetLatestPartitionCrlVersions

`func (o *CreateCrlRestResponse) GetLatestPartitionCrlVersions() map[string]int32`

GetLatestPartitionCrlVersions returns the LatestPartitionCrlVersions field if non-nil, zero value otherwise.

### GetLatestPartitionCrlVersionsOk

`func (o *CreateCrlRestResponse) GetLatestPartitionCrlVersionsOk() (*map[string]int32, bool)`

GetLatestPartitionCrlVersionsOk returns a tuple with the LatestPartitionCrlVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestPartitionCrlVersions

`func (o *CreateCrlRestResponse) SetLatestPartitionCrlVersions(v map[string]int32)`

SetLatestPartitionCrlVersions sets LatestPartitionCrlVersions field to given value.

### HasLatestPartitionCrlVersions

`func (o *CreateCrlRestResponse) HasLatestPartitionCrlVersions() bool`

HasLatestPartitionCrlVersions returns a boolean if a field has been set.

### GetLatestPartitionDeltaCrlVersions

`func (o *CreateCrlRestResponse) GetLatestPartitionDeltaCrlVersions() map[string]int32`

GetLatestPartitionDeltaCrlVersions returns the LatestPartitionDeltaCrlVersions field if non-nil, zero value otherwise.

### GetLatestPartitionDeltaCrlVersionsOk

`func (o *CreateCrlRestResponse) GetLatestPartitionDeltaCrlVersionsOk() (*map[string]int32, bool)`

GetLatestPartitionDeltaCrlVersionsOk returns a tuple with the LatestPartitionDeltaCrlVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestPartitionDeltaCrlVersions

`func (o *CreateCrlRestResponse) SetLatestPartitionDeltaCrlVersions(v map[string]int32)`

SetLatestPartitionDeltaCrlVersions sets LatestPartitionDeltaCrlVersions field to given value.

### HasLatestPartitionDeltaCrlVersions

`func (o *CreateCrlRestResponse) HasLatestPartitionDeltaCrlVersions() bool`

HasLatestPartitionDeltaCrlVersions returns a boolean if a field has been set.

### GetAllSuccess

`func (o *CreateCrlRestResponse) GetAllSuccess() bool`

GetAllSuccess returns the AllSuccess field if non-nil, zero value otherwise.

### GetAllSuccessOk

`func (o *CreateCrlRestResponse) GetAllSuccessOk() (*bool, bool)`

GetAllSuccessOk returns a tuple with the AllSuccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllSuccess

`func (o *CreateCrlRestResponse) SetAllSuccess(v bool)`

SetAllSuccess sets AllSuccess field to given value.

### HasAllSuccess

`func (o *CreateCrlRestResponse) HasAllSuccess() bool`

HasAllSuccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


