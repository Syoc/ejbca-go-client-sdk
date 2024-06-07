# \V1CaManagementApi

All URIs are relative to *http://localhost/ejbca/ejbca-rest-api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Activate**](V1CaManagementApi.md#Activate) | **Put** /v1/ca_management/{ca_name}/activate | Activate a CA
[**Deactivate**](V1CaManagementApi.md#Deactivate) | **Put** /v1/ca_management/{ca_name}/deactivate | Deactivate a CA
[**Status**](V1CaManagementApi.md#Status) | **Get** /v1/ca_management/status | Get the status of this REST Resource



## Activate

> Activate(ctx, caName).Execute()

Activate a CA



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    openapiclient "github.com/Keyfactor/ejbca-go-client-sdk/api/ejbca"
)

func main() {
    caName := "caName_example" // string | Name of the CA to activate

    authenticator, err := openapiclient.NewMTLSAuthenticatorBuilder().
        WithClientCertificatePath("<path to client certificate>").
        WithClientCertificateKeyPath("<path to client key>").
        WithCaCertificatePath("<path to ca certificate>").
        Build()
    if err != nil {
        panic(err)
    }

    configuration := openapiclient.NewConfiguration()
    configuration.Host = "<hostname>:<optional port>"
    configuration.SetAuthenticator(authenticator)

    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.V1CaManagementApi.Activate(context.Background(), caName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `V1CaManagementApi.Activate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**caName** | **string** | Name of the CA to activate | 

### Other Parameters

Other parameters are passed through a pointer to a apiActivateRequest struct via the builder pattern


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


## Deactivate

> Deactivate(ctx, caName).Execute()

Deactivate a CA



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    openapiclient "github.com/Keyfactor/ejbca-go-client-sdk/api/ejbca"
)

func main() {
    caName := "caName_example" // string | Name of the CA to deactivate

    authenticator, err := openapiclient.NewMTLSAuthenticatorBuilder().
        WithClientCertificatePath("<path to client certificate>").
        WithClientCertificateKeyPath("<path to client key>").
        WithCaCertificatePath("<path to ca certificate>").
        Build()
    if err != nil {
        panic(err)
    }

    configuration := openapiclient.NewConfiguration()
    configuration.Host = "<hostname>:<optional port>"
    configuration.SetAuthenticator(authenticator)

    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.V1CaManagementApi.Deactivate(context.Background(), caName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `V1CaManagementApi.Deactivate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**caName** | **string** | Name of the CA to deactivate | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeactivateRequest struct via the builder pattern


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


## Status

> RestResourceStatusRestResponse Status(ctx).Execute()

Get the status of this REST Resource



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    openapiclient "github.com/Keyfactor/ejbca-go-client-sdk/api/ejbca"
)

func main() {

    authenticator, err := openapiclient.NewMTLSAuthenticatorBuilder().
        WithClientCertificatePath("<path to client certificate>").
        WithClientCertificateKeyPath("<path to client key>").
        WithCaCertificatePath("<path to ca certificate>").
        Build()
    if err != nil {
        panic(err)
    }

    configuration := openapiclient.NewConfiguration()
    configuration.Host = "<hostname>:<optional port>"
    configuration.SetAuthenticator(authenticator)

    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.V1CaManagementApi.Status(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `V1CaManagementApi.Status``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Status`: RestResourceStatusRestResponse
    fmt.Fprintf(os.Stdout, "Response from `V1CaManagementApi.Status`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiStatusRequest struct via the builder pattern


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

