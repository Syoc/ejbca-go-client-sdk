# SearchCertificateCriteriaRestRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Property** | Pointer to **string** | A search property | [optional] 
**Value** | Pointer to **string** | A search value. This could be sting value, ISO 8601 Date string, an appropriate string name of End Entity Profile or Certificate Profile or CA | [optional] 
**Operation** | Pointer to **string** | An operation for property on inserted value. &#39;EQUAL&#39; for string, &#39;LIKE&#39; for string value (&#39;QUERY&#39;), &#39;BEFORE&#39; or &#39;AFTER&#39; for date values | [optional] 

## Methods

### NewSearchCertificateCriteriaRestRequest

`func NewSearchCertificateCriteriaRestRequest() *SearchCertificateCriteriaRestRequest`

NewSearchCertificateCriteriaRestRequest instantiates a new SearchCertificateCriteriaRestRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchCertificateCriteriaRestRequestWithDefaults

`func NewSearchCertificateCriteriaRestRequestWithDefaults() *SearchCertificateCriteriaRestRequest`

NewSearchCertificateCriteriaRestRequestWithDefaults instantiates a new SearchCertificateCriteriaRestRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProperty

`func (o *SearchCertificateCriteriaRestRequest) GetProperty() string`

GetProperty returns the Property field if non-nil, zero value otherwise.

### GetPropertyOk

`func (o *SearchCertificateCriteriaRestRequest) GetPropertyOk() (*string, bool)`

GetPropertyOk returns a tuple with the Property field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperty

`func (o *SearchCertificateCriteriaRestRequest) SetProperty(v string)`

SetProperty sets Property field to given value.

### HasProperty

`func (o *SearchCertificateCriteriaRestRequest) HasProperty() bool`

HasProperty returns a boolean if a field has been set.

### GetValue

`func (o *SearchCertificateCriteriaRestRequest) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *SearchCertificateCriteriaRestRequest) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *SearchCertificateCriteriaRestRequest) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *SearchCertificateCriteriaRestRequest) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetOperation

`func (o *SearchCertificateCriteriaRestRequest) GetOperation() string`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *SearchCertificateCriteriaRestRequest) GetOperationOk() (*string, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *SearchCertificateCriteriaRestRequest) SetOperation(v string)`

SetOperation sets Operation field to given value.

### HasOperation

`func (o *SearchCertificateCriteriaRestRequest) HasOperation() bool`

HasOperation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


