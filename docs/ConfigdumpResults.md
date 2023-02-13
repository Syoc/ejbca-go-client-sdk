# ConfigdumpResults

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | Pointer to **bool** |  | [optional] 
**Errors** | Pointer to **[]string** |  | [optional] 
**Warnings** | Pointer to **[]string** |  | [optional] 

## Methods

### NewConfigdumpResults

`func NewConfigdumpResults() *ConfigdumpResults`

NewConfigdumpResults instantiates a new ConfigdumpResults object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigdumpResultsWithDefaults

`func NewConfigdumpResultsWithDefaults() *ConfigdumpResults`

NewConfigdumpResultsWithDefaults instantiates a new ConfigdumpResults object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *ConfigdumpResults) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *ConfigdumpResults) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *ConfigdumpResults) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *ConfigdumpResults) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetErrors

`func (o *ConfigdumpResults) GetErrors() []string`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *ConfigdumpResults) GetErrorsOk() (*[]string, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *ConfigdumpResults) SetErrors(v []string)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *ConfigdumpResults) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetWarnings

`func (o *ConfigdumpResults) GetWarnings() []string`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *ConfigdumpResults) GetWarningsOk() (*[]string, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *ConfigdumpResults) SetWarnings(v []string)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *ConfigdumpResults) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


