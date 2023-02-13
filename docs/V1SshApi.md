# \V1SshApi

All URIs are relative to *http://localhost/ejbca/ejbca-rest-api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Pubkey**](V1SshApi.md#Pubkey) | **Get** /v1/ssh/{ca_name}/pubkey | Retrieves a CA&#39;s public key in SSH format.
[**Status8**](V1SshApi.md#Status8) | **Get** /v1/ssh/status | Get the status of this REST Resource



## Pubkey

> SshPublicKeyRestResponse Pubkey(ctx, caName).Execute()

Retrieves a CA's public key in SSH format.



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
    caName := "caName_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.V1SshApi.Pubkey(context.Background(), caName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `V1SshApi.Pubkey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Pubkey`: SshPublicKeyRestResponse
    fmt.Fprintf(os.Stdout, "Response from `V1SshApi.Pubkey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**caName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPubkeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SshPublicKeyRestResponse**](SshPublicKeyRestResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Status8

> RestResourceStatusRestResponse Status8(ctx).Execute()

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
    resp, r, err := apiClient.V1SshApi.Status8(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `V1SshApi.Status8``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Status8`: RestResourceStatusRestResponse
    fmt.Fprintf(os.Stdout, "Response from `V1SshApi.Status8`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiStatus8Request struct via the builder pattern


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

