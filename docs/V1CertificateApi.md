# \V1CertificateApi

All URIs are relative to *http://localhost/ejbca/ejbca-rest-api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CertificateRequest**](V1CertificateApi.md#CertificateRequest) | **Post** /v1/certificate/certificaterequest | Enrollment with client generated keys for an existing End Entity
[**EnrollKeystore**](V1CertificateApi.md#EnrollKeystore) | **Post** /v1/certificate/enrollkeystore | Keystore enrollment
[**EnrollPkcs10Certificate**](V1CertificateApi.md#EnrollPkcs10Certificate) | **Post** /v1/certificate/pkcs10enroll | Enrollment with client generated keys, using CSR subject
[**FinalizeEnrollment**](V1CertificateApi.md#FinalizeEnrollment) | **Post** /v1/certificate/{request_id}/finalize | Finalize enrollment
[**GetCertificatesAboutToExpire**](V1CertificateApi.md#GetCertificatesAboutToExpire) | **Get** /v1/certificate/expire | Get a list of certificates that are about to expire
[**RevocationStatus**](V1CertificateApi.md#RevocationStatus) | **Get** /v1/certificate/{issuer_dn}/{certificate_serial_number}/revocationstatus | Checks revocation status of the specified certificate
[**RevokeCertificate**](V1CertificateApi.md#RevokeCertificate) | **Put** /v1/certificate/{issuer_dn}/{certificate_serial_number}/revoke | Revokes the specified certificate
[**SearchCertificates**](V1CertificateApi.md#SearchCertificates) | **Post** /v1/certificate/search | Searches for certificates confirming given criteria.
[**Status2**](V1CertificateApi.md#Status2) | **Get** /v1/certificate/status | Get the status of this REST Resource



## CertificateRequest

> CertificateRestResponse CertificateRequest(ctx).CertificateRequestRestRequest(certificateRequestRestRequest).Execute()

Enrollment with client generated keys for an existing End Entity



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
    certificateRequestRestRequest := *openapiclient.NewCertificateRequestRestRequest() // CertificateRequestRestRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.V1CertificateApi.CertificateRequest(context.Background()).CertificateRequestRestRequest(certificateRequestRestRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `V1CertificateApi.CertificateRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CertificateRequest`: CertificateRestResponse
    fmt.Fprintf(os.Stdout, "Response from `V1CertificateApi.CertificateRequest`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCertificateRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **certificateRequestRestRequest** | [**CertificateRequestRestRequest**](CertificateRequestRestRequest.md) |  | 

### Return type

[**CertificateRestResponse**](CertificateRestResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnrollKeystore

> CertificateRestResponse EnrollKeystore(ctx).KeyStoreRestRequest(keyStoreRestRequest).Execute()

Keystore enrollment



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
    keyStoreRestRequest := *openapiclient.NewKeyStoreRestRequest() // KeyStoreRestRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.V1CertificateApi.EnrollKeystore(context.Background()).KeyStoreRestRequest(keyStoreRestRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `V1CertificateApi.EnrollKeystore``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EnrollKeystore`: CertificateRestResponse
    fmt.Fprintf(os.Stdout, "Response from `V1CertificateApi.EnrollKeystore`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEnrollKeystoreRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **keyStoreRestRequest** | [**KeyStoreRestRequest**](KeyStoreRestRequest.md) |  | 

### Return type

[**CertificateRestResponse**](CertificateRestResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnrollPkcs10Certificate

> CertificateRestResponse EnrollPkcs10Certificate(ctx).EnrollCertificateRestRequest(enrollCertificateRestRequest).Execute()

Enrollment with client generated keys, using CSR subject



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
    enrollCertificateRestRequest := *openapiclient.NewEnrollCertificateRestRequest() // EnrollCertificateRestRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.V1CertificateApi.EnrollPkcs10Certificate(context.Background()).EnrollCertificateRestRequest(enrollCertificateRestRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `V1CertificateApi.EnrollPkcs10Certificate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EnrollPkcs10Certificate`: CertificateRestResponse
    fmt.Fprintf(os.Stdout, "Response from `V1CertificateApi.EnrollPkcs10Certificate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEnrollPkcs10CertificateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **enrollCertificateRestRequest** | [**EnrollCertificateRestRequest**](EnrollCertificateRestRequest.md) |  | 

### Return type

[**CertificateRestResponse**](CertificateRestResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FinalizeEnrollment

> CertificateRestResponse FinalizeEnrollment(ctx, requestId).FinalizeRestRequest(finalizeRestRequest).Execute()

Finalize enrollment



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
    requestId := int32(56) // int32 | Approval request id
    finalizeRestRequest := *openapiclient.NewFinalizeRestRequest() // FinalizeRestRequest | responseFormat must be one of 'P12', 'BCFKS', 'JKS', 'DER' (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.V1CertificateApi.FinalizeEnrollment(context.Background(), requestId).FinalizeRestRequest(finalizeRestRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `V1CertificateApi.FinalizeEnrollment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FinalizeEnrollment`: CertificateRestResponse
    fmt.Fprintf(os.Stdout, "Response from `V1CertificateApi.FinalizeEnrollment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**requestId** | **int32** | Approval request id | 

### Other Parameters

Other parameters are passed through a pointer to a apiFinalizeEnrollmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **finalizeRestRequest** | [**FinalizeRestRequest**](FinalizeRestRequest.md) | responseFormat must be one of &#39;P12&#39;, &#39;BCFKS&#39;, &#39;JKS&#39;, &#39;DER&#39; | 

### Return type

[**CertificateRestResponse**](CertificateRestResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCertificatesAboutToExpire

> ExpiringCertificatesRestResponse GetCertificatesAboutToExpire(ctx).Days(days).Offset(offset).MaxNumberOfResults(maxNumberOfResults).Execute()

Get a list of certificates that are about to expire



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
    days := int64(789) // int64 | Request certificates expiring within this number of days (optional)
    offset := int32(56) // int32 | Next offset to display results of, if maxNumberOfResults is exceeded. Starts from 0. (optional)
    maxNumberOfResults := int32(56) // int32 | Maximum number of certificates to display. If result exceeds this value. Modify 'offset' to retrieve more results (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.V1CertificateApi.GetCertificatesAboutToExpire(context.Background()).Days(days).Offset(offset).MaxNumberOfResults(maxNumberOfResults).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `V1CertificateApi.GetCertificatesAboutToExpire``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCertificatesAboutToExpire`: ExpiringCertificatesRestResponse
    fmt.Fprintf(os.Stdout, "Response from `V1CertificateApi.GetCertificatesAboutToExpire`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCertificatesAboutToExpireRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **days** | **int64** | Request certificates expiring within this number of days | 
 **offset** | **int32** | Next offset to display results of, if maxNumberOfResults is exceeded. Starts from 0. | 
 **maxNumberOfResults** | **int32** | Maximum number of certificates to display. If result exceeds this value. Modify &#39;offset&#39; to retrieve more results | 

### Return type

[**ExpiringCertificatesRestResponse**](ExpiringCertificatesRestResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RevocationStatus

> RevokeStatusRestResponse RevocationStatus(ctx, issuerDn, certificateSerialNumber).Execute()

Checks revocation status of the specified certificate



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
    issuerDn := "issuerDn_example" // string | Subject DN of the issuing CA
    certificateSerialNumber := "certificateSerialNumber_example" // string | hex serial number (without prefix, e.g. '00')

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.V1CertificateApi.RevocationStatus(context.Background(), issuerDn, certificateSerialNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `V1CertificateApi.RevocationStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RevocationStatus`: RevokeStatusRestResponse
    fmt.Fprintf(os.Stdout, "Response from `V1CertificateApi.RevocationStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**issuerDn** | **string** | Subject DN of the issuing CA | 
**certificateSerialNumber** | **string** | hex serial number (without prefix, e.g. &#39;00&#39;) | 

### Other Parameters

Other parameters are passed through a pointer to a apiRevocationStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**RevokeStatusRestResponse**](RevokeStatusRestResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RevokeCertificate

> RevokeStatusRestResponse RevokeCertificate(ctx, issuerDn, certificateSerialNumber).Reason(reason).Date(date).Execute()

Revokes the specified certificate



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    issuerDn := "issuerDn_example" // string | Subject DN of the issuing CA
    certificateSerialNumber := "certificateSerialNumber_example" // string | hex serial number (without prefix, e.g. '00')
    reason := "reason_example" // string | Must be valid RFC5280 reason. One of  NOT_REVOKED, UNSPECIFIED ,KEY_COMPROMISE,  CA_COMPROMISE, AFFILIATION_CHANGED, SUPERSEDED, CESSATION_OF_OPERATION,  CERTIFICATE_HOLD, REMOVE_FROM_CRL, PRIVILEGES_WITHDRAWN, AA_COMPROMISE (optional)
    date := time.Now() // time.Time | ISO 8601 Date string, eg. '2018-06-15T14:07:09Z' (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.V1CertificateApi.RevokeCertificate(context.Background(), issuerDn, certificateSerialNumber).Reason(reason).Date(date).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `V1CertificateApi.RevokeCertificate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RevokeCertificate`: RevokeStatusRestResponse
    fmt.Fprintf(os.Stdout, "Response from `V1CertificateApi.RevokeCertificate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**issuerDn** | **string** | Subject DN of the issuing CA | 
**certificateSerialNumber** | **string** | hex serial number (without prefix, e.g. &#39;00&#39;) | 

### Other Parameters

Other parameters are passed through a pointer to a apiRevokeCertificateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **reason** | **string** | Must be valid RFC5280 reason. One of  NOT_REVOKED, UNSPECIFIED ,KEY_COMPROMISE,  CA_COMPROMISE, AFFILIATION_CHANGED, SUPERSEDED, CESSATION_OF_OPERATION,  CERTIFICATE_HOLD, REMOVE_FROM_CRL, PRIVILEGES_WITHDRAWN, AA_COMPROMISE | 
 **date** | **time.Time** | ISO 8601 Date string, eg. &#39;2018-06-15T14:07:09Z&#39; | 

### Return type

[**RevokeStatusRestResponse**](RevokeStatusRestResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchCertificates

> SearchCertificatesRestResponse SearchCertificates(ctx).SearchCertificatesRestRequest(searchCertificatesRestRequest).Execute()

Searches for certificates confirming given criteria.



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
    searchCertificatesRestRequest := *openapiclient.NewSearchCertificatesRestRequest() // SearchCertificatesRestRequest | Maximum number of results and collection of search criterias. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.V1CertificateApi.SearchCertificates(context.Background()).SearchCertificatesRestRequest(searchCertificatesRestRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `V1CertificateApi.SearchCertificates``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchCertificates`: SearchCertificatesRestResponse
    fmt.Fprintf(os.Stdout, "Response from `V1CertificateApi.SearchCertificates`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchCertificatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **searchCertificatesRestRequest** | [**SearchCertificatesRestRequest**](SearchCertificatesRestRequest.md) | Maximum number of results and collection of search criterias. | 

### Return type

[**SearchCertificatesRestResponse**](SearchCertificatesRestResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Status2

> RestResourceStatusRestResponse Status2(ctx).Execute()

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
    resp, r, err := apiClient.V1CertificateApi.Status2(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `V1CertificateApi.Status2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Status2`: RestResourceStatusRestResponse
    fmt.Fprintf(os.Stdout, "Response from `V1CertificateApi.Status2`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiStatus2Request struct via the builder pattern


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

