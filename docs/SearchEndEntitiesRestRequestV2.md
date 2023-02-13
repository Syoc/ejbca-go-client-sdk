# SearchEndEntitiesRestRequestV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaxNumberOfResults** | Pointer to **int32** | Maximum number of results | [optional] 
**CurrentPage** | Pointer to **int32** | Current page number | [optional] 
**Criteria** | Pointer to [**[]SearchEndEntityCriteriaRestRequest**](SearchEndEntityCriteriaRestRequest.md) | A List of search criteria. | [optional] 
**SortOperation** | Pointer to [**SearchEndEntitiesSortRestRequest**](SearchEndEntitiesSortRestRequest.md) |  | [optional] 

## Methods

### NewSearchEndEntitiesRestRequestV2

`func NewSearchEndEntitiesRestRequestV2() *SearchEndEntitiesRestRequestV2`

NewSearchEndEntitiesRestRequestV2 instantiates a new SearchEndEntitiesRestRequestV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchEndEntitiesRestRequestV2WithDefaults

`func NewSearchEndEntitiesRestRequestV2WithDefaults() *SearchEndEntitiesRestRequestV2`

NewSearchEndEntitiesRestRequestV2WithDefaults instantiates a new SearchEndEntitiesRestRequestV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaxNumberOfResults

`func (o *SearchEndEntitiesRestRequestV2) GetMaxNumberOfResults() int32`

GetMaxNumberOfResults returns the MaxNumberOfResults field if non-nil, zero value otherwise.

### GetMaxNumberOfResultsOk

`func (o *SearchEndEntitiesRestRequestV2) GetMaxNumberOfResultsOk() (*int32, bool)`

GetMaxNumberOfResultsOk returns a tuple with the MaxNumberOfResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxNumberOfResults

`func (o *SearchEndEntitiesRestRequestV2) SetMaxNumberOfResults(v int32)`

SetMaxNumberOfResults sets MaxNumberOfResults field to given value.

### HasMaxNumberOfResults

`func (o *SearchEndEntitiesRestRequestV2) HasMaxNumberOfResults() bool`

HasMaxNumberOfResults returns a boolean if a field has been set.

### GetCurrentPage

`func (o *SearchEndEntitiesRestRequestV2) GetCurrentPage() int32`

GetCurrentPage returns the CurrentPage field if non-nil, zero value otherwise.

### GetCurrentPageOk

`func (o *SearchEndEntitiesRestRequestV2) GetCurrentPageOk() (*int32, bool)`

GetCurrentPageOk returns a tuple with the CurrentPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPage

`func (o *SearchEndEntitiesRestRequestV2) SetCurrentPage(v int32)`

SetCurrentPage sets CurrentPage field to given value.

### HasCurrentPage

`func (o *SearchEndEntitiesRestRequestV2) HasCurrentPage() bool`

HasCurrentPage returns a boolean if a field has been set.

### GetCriteria

`func (o *SearchEndEntitiesRestRequestV2) GetCriteria() []SearchEndEntityCriteriaRestRequest`

GetCriteria returns the Criteria field if non-nil, zero value otherwise.

### GetCriteriaOk

`func (o *SearchEndEntitiesRestRequestV2) GetCriteriaOk() (*[]SearchEndEntityCriteriaRestRequest, bool)`

GetCriteriaOk returns a tuple with the Criteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCriteria

`func (o *SearchEndEntitiesRestRequestV2) SetCriteria(v []SearchEndEntityCriteriaRestRequest)`

SetCriteria sets Criteria field to given value.

### HasCriteria

`func (o *SearchEndEntitiesRestRequestV2) HasCriteria() bool`

HasCriteria returns a boolean if a field has been set.

### GetSortOperation

`func (o *SearchEndEntitiesRestRequestV2) GetSortOperation() SearchEndEntitiesSortRestRequest`

GetSortOperation returns the SortOperation field if non-nil, zero value otherwise.

### GetSortOperationOk

`func (o *SearchEndEntitiesRestRequestV2) GetSortOperationOk() (*SearchEndEntitiesSortRestRequest, bool)`

GetSortOperationOk returns a tuple with the SortOperation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortOperation

`func (o *SearchEndEntitiesRestRequestV2) SetSortOperation(v SearchEndEntitiesSortRestRequest)`

SetSortOperation sets SortOperation field to given value.

### HasSortOperation

`func (o *SearchEndEntitiesRestRequestV2) HasSortOperation() bool`

HasSortOperation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


