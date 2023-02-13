# SearchEndEntityCriteriaRestRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Property** | Pointer to **string** | A search property | [optional] 
**Value** | Pointer to **string** | A search value. This could be string value, an appropriate string name of End Entity Profile or Certificate Profile or CA | [optional] 
**Operation** | Pointer to **string** | An operation for property on inserted value. &#39;EQUALS&#39; for string, &#39;LIKE&#39; for string value (&#39;QUERY&#39;) | [optional] 

## Methods

### NewSearchEndEntityCriteriaRestRequest

`func NewSearchEndEntityCriteriaRestRequest() *SearchEndEntityCriteriaRestRequest`

NewSearchEndEntityCriteriaRestRequest instantiates a new SearchEndEntityCriteriaRestRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchEndEntityCriteriaRestRequestWithDefaults

`func NewSearchEndEntityCriteriaRestRequestWithDefaults() *SearchEndEntityCriteriaRestRequest`

NewSearchEndEntityCriteriaRestRequestWithDefaults instantiates a new SearchEndEntityCriteriaRestRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProperty

`func (o *SearchEndEntityCriteriaRestRequest) GetProperty() string`

GetProperty returns the Property field if non-nil, zero value otherwise.

### GetPropertyOk

`func (o *SearchEndEntityCriteriaRestRequest) GetPropertyOk() (*string, bool)`

GetPropertyOk returns a tuple with the Property field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperty

`func (o *SearchEndEntityCriteriaRestRequest) SetProperty(v string)`

SetProperty sets Property field to given value.

### HasProperty

`func (o *SearchEndEntityCriteriaRestRequest) HasProperty() bool`

HasProperty returns a boolean if a field has been set.

### GetValue

`func (o *SearchEndEntityCriteriaRestRequest) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *SearchEndEntityCriteriaRestRequest) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *SearchEndEntityCriteriaRestRequest) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *SearchEndEntityCriteriaRestRequest) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetOperation

`func (o *SearchEndEntityCriteriaRestRequest) GetOperation() string`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *SearchEndEntityCriteriaRestRequest) GetOperationOk() (*string, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *SearchEndEntityCriteriaRestRequest) SetOperation(v string)`

SetOperation sets Operation field to given value.

### HasOperation

`func (o *SearchEndEntityCriteriaRestRequest) HasOperation() bool`

HasOperation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


