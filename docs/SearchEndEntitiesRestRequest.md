# SearchEndEntitiesRestRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaxNumberOfResults** | Pointer to **int32** | Maximum number of results | [optional] 
**CurrentPage** | Pointer to **int32** | Current page number | [optional] 
**Criteria** | Pointer to [**[]SearchEndEntityCriteriaRestRequest**](SearchEndEntityCriteriaRestRequest.md) | A List of search criteria. | [optional] 

## Methods

### NewSearchEndEntitiesRestRequest

`func NewSearchEndEntitiesRestRequest() *SearchEndEntitiesRestRequest`

NewSearchEndEntitiesRestRequest instantiates a new SearchEndEntitiesRestRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchEndEntitiesRestRequestWithDefaults

`func NewSearchEndEntitiesRestRequestWithDefaults() *SearchEndEntitiesRestRequest`

NewSearchEndEntitiesRestRequestWithDefaults instantiates a new SearchEndEntitiesRestRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaxNumberOfResults

`func (o *SearchEndEntitiesRestRequest) GetMaxNumberOfResults() int32`

GetMaxNumberOfResults returns the MaxNumberOfResults field if non-nil, zero value otherwise.

### GetMaxNumberOfResultsOk

`func (o *SearchEndEntitiesRestRequest) GetMaxNumberOfResultsOk() (*int32, bool)`

GetMaxNumberOfResultsOk returns a tuple with the MaxNumberOfResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxNumberOfResults

`func (o *SearchEndEntitiesRestRequest) SetMaxNumberOfResults(v int32)`

SetMaxNumberOfResults sets MaxNumberOfResults field to given value.

### HasMaxNumberOfResults

`func (o *SearchEndEntitiesRestRequest) HasMaxNumberOfResults() bool`

HasMaxNumberOfResults returns a boolean if a field has been set.

### GetCurrentPage

`func (o *SearchEndEntitiesRestRequest) GetCurrentPage() int32`

GetCurrentPage returns the CurrentPage field if non-nil, zero value otherwise.

### GetCurrentPageOk

`func (o *SearchEndEntitiesRestRequest) GetCurrentPageOk() (*int32, bool)`

GetCurrentPageOk returns a tuple with the CurrentPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPage

`func (o *SearchEndEntitiesRestRequest) SetCurrentPage(v int32)`

SetCurrentPage sets CurrentPage field to given value.

### HasCurrentPage

`func (o *SearchEndEntitiesRestRequest) HasCurrentPage() bool`

HasCurrentPage returns a boolean if a field has been set.

### GetCriteria

`func (o *SearchEndEntitiesRestRequest) GetCriteria() []SearchEndEntityCriteriaRestRequest`

GetCriteria returns the Criteria field if non-nil, zero value otherwise.

### GetCriteriaOk

`func (o *SearchEndEntitiesRestRequest) GetCriteriaOk() (*[]SearchEndEntityCriteriaRestRequest, bool)`

GetCriteriaOk returns a tuple with the Criteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCriteria

`func (o *SearchEndEntitiesRestRequest) SetCriteria(v []SearchEndEntityCriteriaRestRequest)`

SetCriteria sets Criteria field to given value.

### HasCriteria

`func (o *SearchEndEntitiesRestRequest) HasCriteria() bool`

HasCriteria returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


