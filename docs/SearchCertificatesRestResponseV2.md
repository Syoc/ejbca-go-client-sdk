# SearchCertificatesRestResponseV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Certificates** | Pointer to [**[]CertificateRestResponseV2**](CertificateRestResponseV2.md) |  | [optional] 
**PaginationSummary** | Pointer to [**PaginationSummary**](PaginationSummary.md) |  | [optional] 

## Methods

### NewSearchCertificatesRestResponseV2

`func NewSearchCertificatesRestResponseV2() *SearchCertificatesRestResponseV2`

NewSearchCertificatesRestResponseV2 instantiates a new SearchCertificatesRestResponseV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchCertificatesRestResponseV2WithDefaults

`func NewSearchCertificatesRestResponseV2WithDefaults() *SearchCertificatesRestResponseV2`

NewSearchCertificatesRestResponseV2WithDefaults instantiates a new SearchCertificatesRestResponseV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCertificates

`func (o *SearchCertificatesRestResponseV2) GetCertificates() []CertificateRestResponseV2`

GetCertificates returns the Certificates field if non-nil, zero value otherwise.

### GetCertificatesOk

`func (o *SearchCertificatesRestResponseV2) GetCertificatesOk() (*[]CertificateRestResponseV2, bool)`

GetCertificatesOk returns a tuple with the Certificates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificates

`func (o *SearchCertificatesRestResponseV2) SetCertificates(v []CertificateRestResponseV2)`

SetCertificates sets Certificates field to given value.

### HasCertificates

`func (o *SearchCertificatesRestResponseV2) HasCertificates() bool`

HasCertificates returns a boolean if a field has been set.

### GetPaginationSummary

`func (o *SearchCertificatesRestResponseV2) GetPaginationSummary() PaginationSummary`

GetPaginationSummary returns the PaginationSummary field if non-nil, zero value otherwise.

### GetPaginationSummaryOk

`func (o *SearchCertificatesRestResponseV2) GetPaginationSummaryOk() (*PaginationSummary, bool)`

GetPaginationSummaryOk returns a tuple with the PaginationSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaginationSummary

`func (o *SearchCertificatesRestResponseV2) SetPaginationSummary(v PaginationSummary)`

SetPaginationSummary sets PaginationSummary field to given value.

### HasPaginationSummary

`func (o *SearchCertificatesRestResponseV2) HasPaginationSummary() bool`

HasPaginationSummary returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


