# \V2CertificateApi

All URIs are relative to *http://localhost/ejbca/ejbca-rest-api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCertificateProfileInfo**](V2CertificateApi.md#GetCertificateProfileInfo) | **Get** /v2/certificate/profile/{profile_name} | Get Certificate Profile Info.
[**SearchCertificates1**](V2CertificateApi.md#SearchCertificates1) | **Post** /v2/certificate/search | Searches for certificates confirming given criteria and pagination.
[**Status3**](V2CertificateApi.md#Status3) | **Get** /v2/certificate/status | Get the status of this REST Resource



## GetCertificateProfileInfo

> CertificateProfileInfoRestResponseV2 GetCertificateProfileInfo(ctx, profileName).Execute()

Get Certificate Profile Info.



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
    profileName := "profileName_example" // string | 

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
    resp, r, err := apiClient.V2CertificateApi.GetCertificateProfileInfo(context.Background(), profileName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `V2CertificateApi.GetCertificateProfileInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCertificateProfileInfo`: CertificateProfileInfoRestResponseV2
    fmt.Fprintf(os.Stdout, "Response from `V2CertificateApi.GetCertificateProfileInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**profileName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCertificateProfileInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CertificateProfileInfoRestResponseV2**](CertificateProfileInfoRestResponseV2.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchCertificates1

> SearchCertificatesRestResponseV2 SearchCertificates1(ctx).SearchCertificatesRestRequestV2(searchCertificatesRestRequestV2).Execute()

Searches for certificates confirming given criteria and pagination.



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
    searchCertificatesRestRequestV2 := *openapiclient.NewSearchCertificatesRestRequestV2() // SearchCertificatesRestRequestV2 | Collection of search criterias and pagination information. (optional)

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
    resp, r, err := apiClient.V2CertificateApi.SearchCertificates1(context.Background()).SearchCertificatesRestRequestV2(searchCertificatesRestRequestV2).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `V2CertificateApi.SearchCertificates1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchCertificates1`: SearchCertificatesRestResponseV2
    fmt.Fprintf(os.Stdout, "Response from `V2CertificateApi.SearchCertificates1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchCertificates1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **searchCertificatesRestRequestV2** | [**SearchCertificatesRestRequestV2**](SearchCertificatesRestRequestV2.md) | Collection of search criterias and pagination information. | 

### Return type

[**SearchCertificatesRestResponseV2**](SearchCertificatesRestResponseV2.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Status3

> RestResourceStatusRestResponse Status3(ctx).Execute()

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
    resp, r, err := apiClient.V2CertificateApi.Status3(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `V2CertificateApi.Status3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Status3`: RestResourceStatusRestResponse
    fmt.Fprintf(os.Stdout, "Response from `V2CertificateApi.Status3`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiStatus3Request struct via the builder pattern


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

