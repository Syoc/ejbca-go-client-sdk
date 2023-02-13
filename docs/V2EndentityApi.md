# \V2EndentityApi

All URIs are relative to *http://localhost/ejbca/ejbca-rest-api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAuthorizedEndEntityProfiles**](V2EndentityApi.md#GetAuthorizedEndEntityProfiles) | **Get** /v2/endentity/profiles/authorized | List of authorized end entity profiles for the current admin.
[**Profile**](V2EndentityApi.md#Profile) | **Get** /v2/endentity/profile/{endentity_profile_name} | Get End Entity Profile content
[**SortedSearch**](V2EndentityApi.md#SortedSearch) | **Post** /v2/endentity/search | Searches and sorts for end entity conforming given criteria.
[**Status7**](V2EndentityApi.md#Status7) | **Get** /v2/endentity/status | Get the status of this REST Resource



## GetAuthorizedEndEntityProfiles

> AuthorizedEEPsRestResponse GetAuthorizedEndEntityProfiles(ctx).Execute()

List of authorized end entity profiles for the current admin.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.V2EndentityApi.GetAuthorizedEndEntityProfiles(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `V2EndentityApi.GetAuthorizedEndEntityProfiles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAuthorizedEndEntityProfiles`: AuthorizedEEPsRestResponse
    fmt.Fprintf(os.Stdout, "Response from `V2EndentityApi.GetAuthorizedEndEntityProfiles`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAuthorizedEndEntityProfilesRequest struct via the builder pattern


### Return type

[**AuthorizedEEPsRestResponse**](AuthorizedEEPsRestResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Profile

> EndEntityProfileResponse Profile(ctx, endentityProfileName).Execute()

Get End Entity Profile content



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    endentityProfileName := "endentityProfileName_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.V2EndentityApi.Profile(context.Background(), endentityProfileName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `V2EndentityApi.Profile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Profile`: EndEntityProfileResponse
    fmt.Fprintf(os.Stdout, "Response from `V2EndentityApi.Profile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**endentityProfileName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EndEntityProfileResponse**](EndEntityProfileResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SortedSearch

> SearchEndEntitiesRestResponse SortedSearch(ctx).SearchEndEntitiesRestRequestV2(searchEndEntitiesRestRequestV2).Execute()

Searches and sorts for end entity conforming given criteria.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    searchEndEntitiesRestRequestV2 := *openapiclient.NewSearchEndEntitiesRestRequestV2() // SearchEndEntitiesRestRequestV2 | Maximum number of results and collection of search criterias. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.V2EndentityApi.SortedSearch(context.Background()).SearchEndEntitiesRestRequestV2(searchEndEntitiesRestRequestV2).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `V2EndentityApi.SortedSearch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SortedSearch`: SearchEndEntitiesRestResponse
    fmt.Fprintf(os.Stdout, "Response from `V2EndentityApi.SortedSearch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSortedSearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **searchEndEntitiesRestRequestV2** | [**SearchEndEntitiesRestRequestV2**](SearchEndEntitiesRestRequestV2.md) | Maximum number of results and collection of search criterias. | 

### Return type

[**SearchEndEntitiesRestResponse**](SearchEndEntitiesRestResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Status7

> RestResourceStatusRestResponse Status7(ctx).Execute()

Get the status of this REST Resource



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.V2EndentityApi.Status7(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `V2EndentityApi.Status7``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Status7`: RestResourceStatusRestResponse
    fmt.Fprintf(os.Stdout, "Response from `V2EndentityApi.Status7`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiStatus7Request struct via the builder pattern


### Return type

[**RestResourceStatusRestResponse**](RestResourceStatusRestResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

