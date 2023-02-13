# SearchCertificatesRestRequestV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pagination** | Pointer to [**Pagination**](Pagination.md) |  | [optional] 
**Sort** | Pointer to [**SearchCertificateSortRestRequest**](SearchCertificateSortRestRequest.md) |  | [optional] 
**Criteria** | Pointer to [**[]SearchCertificateCriteriaRestRequest**](SearchCertificateCriteriaRestRequest.md) | A List of search criteria. | [optional] 

## Methods

### NewSearchCertificatesRestRequestV2

`func NewSearchCertificatesRestRequestV2() *SearchCertificatesRestRequestV2`

NewSearchCertificatesRestRequestV2 instantiates a new SearchCertificatesRestRequestV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchCertificatesRestRequestV2WithDefaults

`func NewSearchCertificatesRestRequestV2WithDefaults() *SearchCertificatesRestRequestV2`

NewSearchCertificatesRestRequestV2WithDefaults instantiates a new SearchCertificatesRestRequestV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPagination

`func (o *SearchCertificatesRestRequestV2) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *SearchCertificatesRestRequestV2) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *SearchCertificatesRestRequestV2) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *SearchCertificatesRestRequestV2) HasPagination() bool`

HasPagination returns a boolean if a field has been set.

### GetSort

`func (o *SearchCertificatesRestRequestV2) GetSort() SearchCertificateSortRestRequest`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *SearchCertificatesRestRequestV2) GetSortOk() (*SearchCertificateSortRestRequest, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *SearchCertificatesRestRequestV2) SetSort(v SearchCertificateSortRestRequest)`

SetSort sets Sort field to given value.

### HasSort

`func (o *SearchCertificatesRestRequestV2) HasSort() bool`

HasSort returns a boolean if a field has been set.

### GetCriteria

`func (o *SearchCertificatesRestRequestV2) GetCriteria() []SearchCertificateCriteriaRestRequest`

GetCriteria returns the Criteria field if non-nil, zero value otherwise.

### GetCriteriaOk

`func (o *SearchCertificatesRestRequestV2) GetCriteriaOk() (*[]SearchCertificateCriteriaRestRequest, bool)`

GetCriteriaOk returns a tuple with the Criteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCriteria

`func (o *SearchCertificatesRestRequestV2) SetCriteria(v []SearchCertificateCriteriaRestRequest)`

SetCriteria sets Criteria field to given value.

### HasCriteria

`func (o *SearchCertificatesRestRequestV2) HasCriteria() bool`

HasCriteria returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


