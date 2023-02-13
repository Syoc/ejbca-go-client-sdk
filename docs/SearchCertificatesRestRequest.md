# SearchCertificatesRestRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaxNumberOfResults** | Pointer to **int32** | Maximum number of results | [optional] 
**Criteria** | Pointer to [**[]SearchCertificateCriteriaRestRequest**](SearchCertificateCriteriaRestRequest.md) | A List of search criteria. | [optional] 

## Methods

### NewSearchCertificatesRestRequest

`func NewSearchCertificatesRestRequest() *SearchCertificatesRestRequest`

NewSearchCertificatesRestRequest instantiates a new SearchCertificatesRestRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchCertificatesRestRequestWithDefaults

`func NewSearchCertificatesRestRequestWithDefaults() *SearchCertificatesRestRequest`

NewSearchCertificatesRestRequestWithDefaults instantiates a new SearchCertificatesRestRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaxNumberOfResults

`func (o *SearchCertificatesRestRequest) GetMaxNumberOfResults() int32`

GetMaxNumberOfResults returns the MaxNumberOfResults field if non-nil, zero value otherwise.

### GetMaxNumberOfResultsOk

`func (o *SearchCertificatesRestRequest) GetMaxNumberOfResultsOk() (*int32, bool)`

GetMaxNumberOfResultsOk returns a tuple with the MaxNumberOfResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxNumberOfResults

`func (o *SearchCertificatesRestRequest) SetMaxNumberOfResults(v int32)`

SetMaxNumberOfResults sets MaxNumberOfResults field to given value.

### HasMaxNumberOfResults

`func (o *SearchCertificatesRestRequest) HasMaxNumberOfResults() bool`

HasMaxNumberOfResults returns a boolean if a field has been set.

### GetCriteria

`func (o *SearchCertificatesRestRequest) GetCriteria() []SearchCertificateCriteriaRestRequest`

GetCriteria returns the Criteria field if non-nil, zero value otherwise.

### GetCriteriaOk

`func (o *SearchCertificatesRestRequest) GetCriteriaOk() (*[]SearchCertificateCriteriaRestRequest, bool)`

GetCriteriaOk returns a tuple with the Criteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCriteria

`func (o *SearchCertificatesRestRequest) SetCriteria(v []SearchCertificateCriteriaRestRequest)`

SetCriteria sets Criteria field to given value.

### HasCriteria

`func (o *SearchCertificatesRestRequest) HasCriteria() bool`

HasCriteria returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


