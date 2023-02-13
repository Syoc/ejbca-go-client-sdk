# \V1CryptotokenApi

All URIs are relative to *http://localhost/ejbca/ejbca-rest-api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Activate1**](V1CryptotokenApi.md#Activate1) | **Put** /v1/cryptotoken/{cryptotoken_name}/activate | Activate a Crypto Token
[**Deactivate1**](V1CryptotokenApi.md#Deactivate1) | **Put** /v1/cryptotoken/{cryptotoken_name}/deactivate | Deactivate a Crypto Token
[**GenerateKeys**](V1CryptotokenApi.md#GenerateKeys) | **Post** /v1/cryptotoken/{cryptotoken_name}/generatekeys | Generate keys
[**RemoveKeys**](V1CryptotokenApi.md#RemoveKeys) | **Post** /v1/cryptotoken/{cryptotoken_name}/{key_pair_alias}/removekeys | Remove keys
[**Status5**](V1CryptotokenApi.md#Status5) | **Get** /v1/cryptotoken/status | Get the status of this REST Resource



## Activate1

> Activate1(ctx, cryptotokenName).CryptoTokenActivationRestRequest(cryptoTokenActivationRestRequest).Execute()

Activate a Crypto Token



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
    cryptotokenName := "cryptotokenName_example" // string | Name of the token to activate
    cryptoTokenActivationRestRequest := *openapiclient.NewCryptoTokenActivationRestRequest() // CryptoTokenActivationRestRequest | activation code (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.V1CryptotokenApi.Activate1(context.Background(), cryptotokenName).CryptoTokenActivationRestRequest(cryptoTokenActivationRestRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `V1CryptotokenApi.Activate1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cryptotokenName** | **string** | Name of the token to activate | 

### Other Parameters

Other parameters are passed through a pointer to a apiActivate1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cryptoTokenActivationRestRequest** | [**CryptoTokenActivationRestRequest**](CryptoTokenActivationRestRequest.md) | activation code | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Deactivate1

> Deactivate1(ctx, cryptotokenName).Execute()

Deactivate a Crypto Token



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
    cryptotokenName := "cryptotokenName_example" // string | Name of the token to deactivate

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.V1CryptotokenApi.Deactivate1(context.Background(), cryptotokenName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `V1CryptotokenApi.Deactivate1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cryptotokenName** | **string** | Name of the token to deactivate | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeactivate1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GenerateKeys

> GenerateKeys(ctx, cryptotokenName).CryptoTokenKeyGenerationRestRequest(cryptoTokenKeyGenerationRestRequest).Execute()

Generate keys



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
    cryptotokenName := "cryptotokenName_example" // string | Name of the token to generate keys for
    cryptoTokenKeyGenerationRestRequest := *openapiclient.NewCryptoTokenKeyGenerationRestRequest() // CryptoTokenKeyGenerationRestRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.V1CryptotokenApi.GenerateKeys(context.Background(), cryptotokenName).CryptoTokenKeyGenerationRestRequest(cryptoTokenKeyGenerationRestRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `V1CryptotokenApi.GenerateKeys``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cryptotokenName** | **string** | Name of the token to generate keys for | 

### Other Parameters

Other parameters are passed through a pointer to a apiGenerateKeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cryptoTokenKeyGenerationRestRequest** | [**CryptoTokenKeyGenerationRestRequest**](CryptoTokenKeyGenerationRestRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveKeys

> RemoveKeys(ctx, cryptotokenName, keyPairAlias).Execute()

Remove keys



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
    cryptotokenName := "cryptotokenName_example" // string | Name of the token to remove keys for.
    keyPairAlias := "keyPairAlias_example" // string | Alias for the key to be removed from the crypto token.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.V1CryptotokenApi.RemoveKeys(context.Background(), cryptotokenName, keyPairAlias).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `V1CryptotokenApi.RemoveKeys``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cryptotokenName** | **string** | Name of the token to remove keys for. | 
**keyPairAlias** | **string** | Alias for the key to be removed from the crypto token. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveKeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Status5

> RestResourceStatusRestResponse Status5(ctx).Execute()

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
    resp, r, err := apiClient.V1CryptotokenApi.Status5(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `V1CryptotokenApi.Status5``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Status5`: RestResourceStatusRestResponse
    fmt.Fprintf(os.Stdout, "Response from `V1CryptotokenApi.Status5`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiStatus5Request struct via the builder pattern


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

