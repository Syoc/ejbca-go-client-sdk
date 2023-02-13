# SearchCertificatesRestResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Certificates** | Pointer to [**[]CertificateRestResponse**](CertificateRestResponse.md) |  | [optional] 
**MoreResults** | Pointer to **bool** |  | [optional] 

## Methods

### NewSearchCertificatesRestResponse

`func NewSearchCertificatesRestResponse() *SearchCertificatesRestResponse`

NewSearchCertificatesRestResponse instantiates a new SearchCertificatesRestResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchCertificatesRestResponseWithDefaults

`func NewSearchCertificatesRestResponseWithDefaults() *SearchCertificatesRestResponse`

NewSearchCertificatesRestResponseWithDefaults instantiates a new SearchCertificatesRestResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCertificates

`func (o *SearchCertificatesRestResponse) GetCertificates() []CertificateRestResponse`

GetCertificates returns the Certificates field if non-nil, zero value otherwise.

### GetCertificatesOk

`func (o *SearchCertificatesRestResponse) GetCertificatesOk() (*[]CertificateRestResponse, bool)`

GetCertificatesOk returns a tuple with the Certificates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificates

`func (o *SearchCertificatesRestResponse) SetCertificates(v []CertificateRestResponse)`

SetCertificates sets Certificates field to given value.

### HasCertificates

`func (o *SearchCertificatesRestResponse) HasCertificates() bool`

HasCertificates returns a boolean if a field has been set.

### GetMoreResults

`func (o *SearchCertificatesRestResponse) GetMoreResults() bool`

GetMoreResults returns the MoreResults field if non-nil, zero value otherwise.

### GetMoreResultsOk

`func (o *SearchCertificatesRestResponse) GetMoreResultsOk() (*bool, bool)`

GetMoreResultsOk returns a tuple with the MoreResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoreResults

`func (o *SearchCertificatesRestResponse) SetMoreResults(v bool)`

SetMoreResults sets MoreResults field to given value.

### HasMoreResults

`func (o *SearchCertificatesRestResponse) HasMoreResults() bool`

HasMoreResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


