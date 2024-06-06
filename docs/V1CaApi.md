# \V1CaApi

All URIs are relative to *http://localhost/ejbca/ejbca-rest-api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCrl**](V1CaApi.md#CreateCrl) | **Post** /v1/ca/{issuer_dn}/createcrl | Create CRL(main, partition and delta) issued by this CA
[**GetCertificateAsPem**](V1CaApi.md#GetCertificateAsPem) | **Get** /v1/ca/{subject_dn}/certificate/download | Get PEM file with the active CA certificate chain
[**GetLatestCrl**](V1CaApi.md#GetLatestCrl) | **Get** /v1/ca/{issuer_dn}/getLatestCrl | Returns the latest CRL issued by this CA
[**ImportCrl**](V1CaApi.md#ImportCrl) | **Post** /v1/ca/{issuer_dn}/importcrl | Import a certificate revocation list (CRL) for a CA
[**ListCas**](V1CaApi.md#ListCas) | **Get** /v1/ca | Returns the Response containing the list of CAs with general information per CA as Json
[**Status1**](V1CaApi.md#Status1) | **Get** /v1/ca/status | Get the status of this REST Resource



## CreateCrl

> CreateCrlRestResponse CreateCrl(ctx, issuerDn).Deltacrl(deltacrl).Execute()

Create CRL(main, partition and delta) issued by this CA



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    openapiclient "github.com/Keyfactor/ejbca-go-client-sdk/ejbca"
)

func main() {
    issuerDn := "issuerDn_example" // string | the CRL issuers DN (CAs subject DN)
    deltacrl := true // bool | true to also create the deltaCRL, false to only create the base CRL (optional) (default to false)

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
    resp, r, err := apiClient.V1CaApi.CreateCrl(context.Background(), issuerDn).Deltacrl(deltacrl).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `V1CaApi.CreateCrl``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCrl`: CreateCrlRestResponse
    fmt.Fprintf(os.Stdout, "Response from `V1CaApi.CreateCrl`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**issuerDn** | **string** | the CRL issuers DN (CAs subject DN) | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCrlRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **deltacrl** | **bool** | true to also create the deltaCRL, false to only create the base CRL | [default to false]

### Return type

[**CreateCrlRestResponse**](CreateCrlRestResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCertificateAsPem

> GetCertificateAsPem(ctx, subjectDn).Execute()

Get PEM file with the active CA certificate chain



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    openapiclient "github.com/Keyfactor/ejbca-go-client-sdk/ejbca"
)

func main() {
    subjectDn := "subjectDn_example" // string | CAs subject DN

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
    resp, r, err := apiClient.V1CaApi.GetCertificateAsPem(context.Background(), subjectDn).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `V1CaApi.GetCertificateAsPem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subjectDn** | **string** | CAs subject DN | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCertificateAsPemRequest struct via the builder pattern


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


## GetLatestCrl

> CrlRestResponse GetLatestCrl(ctx, issuerDn).DeltaCrl(deltaCrl).CrlPartitionIndex(crlPartitionIndex).Execute()

Returns the latest CRL issued by this CA



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    openapiclient "github.com/Keyfactor/ejbca-go-client-sdk/ejbca"
)

func main() {
    issuerDn := "issuerDn_example" // string | the CRL issuers DN (CAs subject DN)
    deltaCrl := true // bool | true to get the latest deltaCRL, false to get the latest complete CRL (optional) (default to false)
    crlPartitionIndex := int32(56) // int32 | the CRL partition index (optional) (default to 0)

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
    resp, r, err := apiClient.V1CaApi.GetLatestCrl(context.Background(), issuerDn).DeltaCrl(deltaCrl).CrlPartitionIndex(crlPartitionIndex).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `V1CaApi.GetLatestCrl``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLatestCrl`: CrlRestResponse
    fmt.Fprintf(os.Stdout, "Response from `V1CaApi.GetLatestCrl`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**issuerDn** | **string** | the CRL issuers DN (CAs subject DN) | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLatestCrlRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **deltaCrl** | **bool** | true to get the latest deltaCRL, false to get the latest complete CRL | [default to false]
 **crlPartitionIndex** | **int32** | the CRL partition index | [default to 0]

### Return type

[**CrlRestResponse**](CrlRestResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImportCrl

> ImportCrl(ctx, issuerDn).CrlFile(crlFile).CrlPartitionIndex(crlPartitionIndex).Execute()

Import a certificate revocation list (CRL) for a CA



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    openapiclient "github.com/Keyfactor/ejbca-go-client-sdk/ejbca"
)

func main() {
    issuerDn := "issuerDn_example" // string | the CRL issuers DN (CAs subject DN)
    crlFile := os.NewFile(1234, "some_file") // *os.File | CRL file in DER format (optional)
    crlPartitionIndex := int32(56) // int32 | CRL partition index (optional) (default to 0)

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
    resp, r, err := apiClient.V1CaApi.ImportCrl(context.Background(), issuerDn).CrlFile(crlFile).CrlPartitionIndex(crlPartitionIndex).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `V1CaApi.ImportCrl``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**issuerDn** | **string** | the CRL issuers DN (CAs subject DN) | 

### Other Parameters

Other parameters are passed through a pointer to a apiImportCrlRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **crlFile** | ***os.File** | CRL file in DER format | 
 **crlPartitionIndex** | **int32** | CRL partition index | [default to 0]

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCas

> CaInfosRestResponse ListCas(ctx).Execute()

Returns the Response containing the list of CAs with general information per CA as Json



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    openapiclient "github.com/Keyfactor/ejbca-go-client-sdk/ejbca"
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
    resp, r, err := apiClient.V1CaApi.ListCas(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `V1CaApi.ListCas``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListCas`: CaInfosRestResponse
    fmt.Fprintf(os.Stdout, "Response from `V1CaApi.ListCas`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListCasRequest struct via the builder pattern


### Return type

[**CaInfosRestResponse**](CaInfosRestResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Status1

> RestResourceStatusRestResponse Status1(ctx).Execute()

Get the status of this REST Resource



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    openapiclient "github.com/Keyfactor/ejbca-go-client-sdk/ejbca"
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
    resp, r, err := apiClient.V1CaApi.Status1(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `V1CaApi.Status1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Status1`: RestResourceStatusRestResponse
    fmt.Fprintf(os.Stdout, "Response from `V1CaApi.Status1`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiStatus1Request struct via the builder pattern


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

