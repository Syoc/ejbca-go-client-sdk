# SearchEndEntitiesRestResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EndEntities** | Pointer to [**[]EndEntityRestResponse**](EndEntityRestResponse.md) |  | [optional] 
**MoreResults** | Pointer to **bool** |  | [optional] 

## Methods

### NewSearchEndEntitiesRestResponse

`func NewSearchEndEntitiesRestResponse() *SearchEndEntitiesRestResponse`

NewSearchEndEntitiesRestResponse instantiates a new SearchEndEntitiesRestResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchEndEntitiesRestResponseWithDefaults

`func NewSearchEndEntitiesRestResponseWithDefaults() *SearchEndEntitiesRestResponse`

NewSearchEndEntitiesRestResponseWithDefaults instantiates a new SearchEndEntitiesRestResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndEntities

`func (o *SearchEndEntitiesRestResponse) GetEndEntities() []EndEntityRestResponse`

GetEndEntities returns the EndEntities field if non-nil, zero value otherwise.

### GetEndEntitiesOk

`func (o *SearchEndEntitiesRestResponse) GetEndEntitiesOk() (*[]EndEntityRestResponse, bool)`

GetEndEntitiesOk returns a tuple with the EndEntities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndEntities

`func (o *SearchEndEntitiesRestResponse) SetEndEntities(v []EndEntityRestResponse)`

SetEndEntities sets EndEntities field to given value.

### HasEndEntities

`func (o *SearchEndEntitiesRestResponse) HasEndEntities() bool`

HasEndEntities returns a boolean if a field has been set.

### GetMoreResults

`func (o *SearchEndEntitiesRestResponse) GetMoreResults() bool`

GetMoreResults returns the MoreResults field if non-nil, zero value otherwise.

### GetMoreResultsOk

`func (o *SearchEndEntitiesRestResponse) GetMoreResultsOk() (*bool, bool)`

GetMoreResultsOk returns a tuple with the MoreResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoreResults

`func (o *SearchEndEntitiesRestResponse) SetMoreResults(v bool)`

SetMoreResults sets MoreResults field to given value.

### HasMoreResults

`func (o *SearchEndEntitiesRestResponse) HasMoreResults() bool`

HasMoreResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


