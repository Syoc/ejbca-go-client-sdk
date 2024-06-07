# \V1EndentityApi

All URIs are relative to *http://localhost/ejbca/ejbca-rest-api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Add**](V1EndentityApi.md#Add) | **Post** /v1/endentity | Add new end entity, if it does not exist
[**Delete**](V1EndentityApi.md#Delete) | **Delete** /v1/endentity/{endentity_name} | Deletes end entity
[**Revoke**](V1EndentityApi.md#Revoke) | **Put** /v1/endentity/{endentity_name}/revoke | Revokes all end entity certificates
[**Search**](V1EndentityApi.md#Search) | **Post** /v1/endentity/search | Searches for end entity confirming given criteria.
[**Setstatus**](V1EndentityApi.md#Setstatus) | **Post** /v1/endentity/{endentity_name}/setstatus | Edits end entity setting new status
[**Status6**](V1EndentityApi.md#Status6) | **Get** /v1/endentity/status | Get the status of this REST Resource



## Add

> Add(ctx).AddEndEntityRestRequest(addEndEntityRestRequest).Execute()

Add new end entity, if it does not exist



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
    addEndEntityRestRequest := *openapiclient.NewAddEndEntityRestRequest() // AddEndEntityRestRequest | request (optional)

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
    resp, r, err := apiClient.V1EndentityApi.Add(context.Background()).AddEndEntityRestRequest(addEndEntityRestRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `V1EndentityApi.Add``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **addEndEntityRestRequest** | [**AddEndEntityRestRequest**](AddEndEntityRestRequest.md) | request | 

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


## Delete

> Delete(ctx, endentityName).Execute()

Deletes end entity



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
    endentityName := "endentityName_example" // string | Name of the end entity

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
    resp, r, err := apiClient.V1EndentityApi.Delete(context.Background(), endentityName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `V1EndentityApi.Delete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**endentityName** | **string** | Name of the end entity | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRequest struct via the builder pattern


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


## Revoke

> Revoke(ctx, endentityName).EndEntityRevocationRestRequest(endEntityRevocationRestRequest).Execute()

Revokes all end entity certificates



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
    endentityName := "endentityName_example" // string | Name of the end entity
    endEntityRevocationRestRequest := *openapiclient.NewEndEntityRevocationRestRequest() // EndEntityRevocationRestRequest | request (optional)

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
    resp, r, err := apiClient.V1EndentityApi.Revoke(context.Background(), endentityName).EndEntityRevocationRestRequest(endEntityRevocationRestRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `V1EndentityApi.Revoke``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**endentityName** | **string** | Name of the end entity | 

### Other Parameters

Other parameters are passed through a pointer to a apiRevokeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **endEntityRevocationRestRequest** | [**EndEntityRevocationRestRequest**](EndEntityRevocationRestRequest.md) | request | 

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


## Search

> SearchEndEntitiesRestResponse Search(ctx).SearchEndEntitiesRestRequest(searchEndEntitiesRestRequest).Execute()

Searches for end entity confirming given criteria.



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
    searchEndEntitiesRestRequest := *openapiclient.NewSearchEndEntitiesRestRequest() // SearchEndEntitiesRestRequest | Maximum number of results and collection of search criterias. (optional)

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
    resp, r, err := apiClient.V1EndentityApi.Search(context.Background()).SearchEndEntitiesRestRequest(searchEndEntitiesRestRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `V1EndentityApi.Search``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Search`: SearchEndEntitiesRestResponse
    fmt.Fprintf(os.Stdout, "Response from `V1EndentityApi.Search`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **searchEndEntitiesRestRequest** | [**SearchEndEntitiesRestRequest**](SearchEndEntitiesRestRequest.md) | Maximum number of results and collection of search criterias. | 

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


## Setstatus

> Setstatus(ctx, endentityName).SetEndEntityStatusRestRequest(setEndEntityStatusRestRequest).Execute()

Edits end entity setting new status



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
    endentityName := "endentityName_example" // string | Name of the end entity to edit status for
    setEndEntityStatusRestRequest := *openapiclient.NewSetEndEntityStatusRestRequest() // SetEndEntityStatusRestRequest | request (optional)

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
    resp, r, err := apiClient.V1EndentityApi.Setstatus(context.Background(), endentityName).SetEndEntityStatusRestRequest(setEndEntityStatusRestRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `V1EndentityApi.Setstatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**endentityName** | **string** | Name of the end entity to edit status for | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetstatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **setEndEntityStatusRestRequest** | [**SetEndEntityStatusRestRequest**](SetEndEntityStatusRestRequest.md) | request | 

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


## Status6

> RestResourceStatusRestResponse Status6(ctx).Execute()

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
    resp, r, err := apiClient.V1EndentityApi.Status6(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `V1EndentityApi.Status6``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Status6`: RestResourceStatusRestResponse
    fmt.Fprintf(os.Stdout, "Response from `V1EndentityApi.Status6`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiStatus6Request struct via the builder pattern


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

