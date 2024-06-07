# \V1ConfigdumpApi

All URIs are relative to *http://localhost/ejbca/ejbca-rest-api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetJsonConfigdump**](V1ConfigdumpApi.md#GetJsonConfigdump) | **Get** /v1/configdump | Get the configuration in JSON.
[**GetJsonConfigdumpForType**](V1ConfigdumpApi.md#GetJsonConfigdumpForType) | **Get** /v1/configdump/{type} | Get the configuration for type in JSON.
[**GetJsonConfigdumpForTypeAndSetting**](V1ConfigdumpApi.md#GetJsonConfigdumpForTypeAndSetting) | **Get** /v1/configdump/{type}/{setting} | Get the configuration for a type and setting in JSON.
[**GetZipExport**](V1ConfigdumpApi.md#GetZipExport) | **Get** /v1/configdump/configdump.zip | Get the configuration as a ZIP file.
[**PostJsonImport**](V1ConfigdumpApi.md#PostJsonImport) | **Post** /v1/configdump | Put the configuration in JSON.
[**PostZipImport**](V1ConfigdumpApi.md#PostZipImport) | **Post** /v1/configdump/configdump.zip | Put the configuration as a ZIP file.
[**Status4**](V1ConfigdumpApi.md#Status4) | **Get** /v1/configdump/status | Get the status of this REST Resource



## GetJsonConfigdump

> []string GetJsonConfigdump(ctx).Ignoreerrors(ignoreerrors).Defaults(defaults).Externalcas(externalcas).Include(include).Exclude(exclude).Execute()

Get the configuration in JSON.



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
    ignoreerrors := true // bool | Print a warning instead of aborting and throwing an exception on errors. (optional) (default to false)
    defaults := true // bool | Also include fields having the default value. (optional) (default to false)
    externalcas := true // bool | Enables export of external CAs (i.e. CAs where there's only a certificate and nothing else) (optional) (default to false)
    include := []string{"Inner_example"} // []string | Names of items/types to include in the export. The syntax is identical to that of exclude. For items of types that aren't listed, everything is included. (optional)
    exclude := []string{"Inner_example"} // []string | Names of items/types to exclude in the export, separated by semicolon. Type and name is separated by a colon, and wildcards \"\\*\" are allowed. Both are case-insensitive. E.g. exclude=\"\\*:Example CA;cryptotoken:Example\\*;systemconfiguration:\\*\".  Supported types are: ACMECONFIG/acme-config, CA/certification-authorities,  CRYPTOTOKEN/crypto-tokens, PUBLISHER/publishers, APPROVALPROFILE/approval-profiles, CERTPROFILE/certificate-profiles, EEPROFILE/end-entity-profiles, SERVICE/services, ROLE/admin-roles, KEYBINDING/internal-key-bindings, ADMINPREFS/admin-preferences, OCSPCONFIG/ocsp-configuration, PEERCONNECTOR/peer-connectors, SCEPCONFIG/scep-config, CMPCONFIG/cmp-config, ESTCONFIG/est-config, VALIDATOR/validators, CTLOG/ct-logs, EXTENDEDKEYUSAGE/extended-key-usage, CERTEXTENSION/custom-certificate-extensions,  OAUTHKEY/trusted-oauth-providers, AVAILABLEPROTOCOLS/available-protocols (optional)

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
    resp, r, err := apiClient.V1ConfigdumpApi.GetJsonConfigdump(context.Background()).Ignoreerrors(ignoreerrors).Defaults(defaults).Externalcas(externalcas).Include(include).Exclude(exclude).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `V1ConfigdumpApi.GetJsonConfigdump``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetJsonConfigdump`: []string
    fmt.Fprintf(os.Stdout, "Response from `V1ConfigdumpApi.GetJsonConfigdump`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetJsonConfigdumpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ignoreerrors** | **bool** | Print a warning instead of aborting and throwing an exception on errors. | [default to false]
 **defaults** | **bool** | Also include fields having the default value. | [default to false]
 **externalcas** | **bool** | Enables export of external CAs (i.e. CAs where there&#39;s only a certificate and nothing else) | [default to false]
 **include** | **[]string** | Names of items/types to include in the export. The syntax is identical to that of exclude. For items of types that aren&#39;t listed, everything is included. | 
 **exclude** | **[]string** | Names of items/types to exclude in the export, separated by semicolon. Type and name is separated by a colon, and wildcards \&quot;\\*\&quot; are allowed. Both are case-insensitive. E.g. exclude&#x3D;\&quot;\\*:Example CA;cryptotoken:Example\\*;systemconfiguration:\\*\&quot;.  Supported types are: ACMECONFIG/acme-config, CA/certification-authorities,  CRYPTOTOKEN/crypto-tokens, PUBLISHER/publishers, APPROVALPROFILE/approval-profiles, CERTPROFILE/certificate-profiles, EEPROFILE/end-entity-profiles, SERVICE/services, ROLE/admin-roles, KEYBINDING/internal-key-bindings, ADMINPREFS/admin-preferences, OCSPCONFIG/ocsp-configuration, PEERCONNECTOR/peer-connectors, SCEPCONFIG/scep-config, CMPCONFIG/cmp-config, ESTCONFIG/est-config, VALIDATOR/validators, CTLOG/ct-logs, EXTENDEDKEYUSAGE/extended-key-usage, CERTEXTENSION/custom-certificate-extensions,  OAUTHKEY/trusted-oauth-providers, AVAILABLEPROTOCOLS/available-protocols | 

### Return type

**[]string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetJsonConfigdumpForType

> []string GetJsonConfigdumpForType(ctx, type_).Ignoreerrors(ignoreerrors).Defaults(defaults).Externalcas(externalcas).Execute()

Get the configuration for type in JSON.



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
    type_ := "type__example" // string | Configuration type to export.  Supported types are: ACMECONFIG/acme-config, CA/certification-authorities,  CRYPTOTOKEN/crypto-tokens, PUBLISHER/publishers, APPROVALPROFILE/approval-profiles, CERTPROFILE/certificate-profiles, EEPROFILE/end-entity-profiles, SERVICE/services, ROLE/admin-roles, KEYBINDING/internal-key-bindings, ADMINPREFS/admin-preferences, OCSPCONFIG/ocsp-configuration, PEERCONNECTOR/peer-connectors, SCEPCONFIG/scep-config, CMPCONFIG/cmp-config, ESTCONFIG/est-config, VALIDATOR/validators, CTLOG/ct-logs, EXTENDEDKEYUSAGE/extended-key-usage, CERTEXTENSION/custom-certificate-extensions,  OAUTHKEY/trusted-oauth-providers, AVAILABLEPROTOCOLS/available-protocols
    ignoreerrors := true // bool | Print a warning instead of aborting and throwing an exception on errors. (optional) (default to false)
    defaults := true // bool | Also include fields having the default value. (optional) (default to false)
    externalcas := true // bool | Enables export of external CAs (i.e. CAs where there's only a certificate and nothing else) (optional) (default to false)

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
    resp, r, err := apiClient.V1ConfigdumpApi.GetJsonConfigdumpForType(context.Background(), type_).Ignoreerrors(ignoreerrors).Defaults(defaults).Externalcas(externalcas).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `V1ConfigdumpApi.GetJsonConfigdumpForType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetJsonConfigdumpForType`: []string
    fmt.Fprintf(os.Stdout, "Response from `V1ConfigdumpApi.GetJsonConfigdumpForType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**type_** | **string** | Configuration type to export.  Supported types are: ACMECONFIG/acme-config, CA/certification-authorities,  CRYPTOTOKEN/crypto-tokens, PUBLISHER/publishers, APPROVALPROFILE/approval-profiles, CERTPROFILE/certificate-profiles, EEPROFILE/end-entity-profiles, SERVICE/services, ROLE/admin-roles, KEYBINDING/internal-key-bindings, ADMINPREFS/admin-preferences, OCSPCONFIG/ocsp-configuration, PEERCONNECTOR/peer-connectors, SCEPCONFIG/scep-config, CMPCONFIG/cmp-config, ESTCONFIG/est-config, VALIDATOR/validators, CTLOG/ct-logs, EXTENDEDKEYUSAGE/extended-key-usage, CERTEXTENSION/custom-certificate-extensions,  OAUTHKEY/trusted-oauth-providers, AVAILABLEPROTOCOLS/available-protocols | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetJsonConfigdumpForTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ignoreerrors** | **bool** | Print a warning instead of aborting and throwing an exception on errors. | [default to false]
 **defaults** | **bool** | Also include fields having the default value. | [default to false]
 **externalcas** | **bool** | Enables export of external CAs (i.e. CAs where there&#39;s only a certificate and nothing else) | [default to false]

### Return type

**[]string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetJsonConfigdumpForTypeAndSetting

> []string GetJsonConfigdumpForTypeAndSetting(ctx, type_, setting).Ignoreerrors(ignoreerrors).Defaults(defaults).Execute()

Get the configuration for a type and setting in JSON.



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
    type_ := "type__example" // string | Configuration type to export.  Supported types are: ACMECONFIG/acme-config, CA/certification-authorities,  CRYPTOTOKEN/crypto-tokens, PUBLISHER/publishers, APPROVALPROFILE/approval-profiles, CERTPROFILE/certificate-profiles, EEPROFILE/end-entity-profiles, SERVICE/services, ROLE/admin-roles, KEYBINDING/internal-key-bindings, ADMINPREFS/admin-preferences, OCSPCONFIG/ocsp-configuration, PEERCONNECTOR/peer-connectors, SCEPCONFIG/scep-config, CMPCONFIG/cmp-config, ESTCONFIG/est-config, VALIDATOR/validators, CTLOG/ct-logs, EXTENDEDKEYUSAGE/extended-key-usage, CERTEXTENSION/custom-certificate-extensions,  OAUTHKEY/trusted-oauth-providers, AVAILABLEPROTOCOLS/available-protocols
    setting := "setting_example" // string | Individual configuration name to export
    ignoreerrors := true // bool | Print a warning instead of aborting and throwing an exception on errors. (optional) (default to false)
    defaults := true // bool | Also include fields having the default value. (optional) (default to false)

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
    resp, r, err := apiClient.V1ConfigdumpApi.GetJsonConfigdumpForTypeAndSetting(context.Background(), type_, setting).Ignoreerrors(ignoreerrors).Defaults(defaults).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `V1ConfigdumpApi.GetJsonConfigdumpForTypeAndSetting``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetJsonConfigdumpForTypeAndSetting`: []string
    fmt.Fprintf(os.Stdout, "Response from `V1ConfigdumpApi.GetJsonConfigdumpForTypeAndSetting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**type_** | **string** | Configuration type to export.  Supported types are: ACMECONFIG/acme-config, CA/certification-authorities,  CRYPTOTOKEN/crypto-tokens, PUBLISHER/publishers, APPROVALPROFILE/approval-profiles, CERTPROFILE/certificate-profiles, EEPROFILE/end-entity-profiles, SERVICE/services, ROLE/admin-roles, KEYBINDING/internal-key-bindings, ADMINPREFS/admin-preferences, OCSPCONFIG/ocsp-configuration, PEERCONNECTOR/peer-connectors, SCEPCONFIG/scep-config, CMPCONFIG/cmp-config, ESTCONFIG/est-config, VALIDATOR/validators, CTLOG/ct-logs, EXTENDEDKEYUSAGE/extended-key-usage, CERTEXTENSION/custom-certificate-extensions,  OAUTHKEY/trusted-oauth-providers, AVAILABLEPROTOCOLS/available-protocols | 
**setting** | **string** | Individual configuration name to export | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetJsonConfigdumpForTypeAndSettingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **ignoreerrors** | **bool** | Print a warning instead of aborting and throwing an exception on errors. | [default to false]
 **defaults** | **bool** | Also include fields having the default value. | [default to false]

### Return type

**[]string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetZipExport

> []string GetZipExport(ctx).Ignoreerrors(ignoreerrors).Defaults(defaults).Externalcas(externalcas).Include(include).Exclude(exclude).Execute()

Get the configuration as a ZIP file.



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
    ignoreerrors := true // bool | Print a warning instead of aborting and throwing an exception on errors. (optional) (default to false)
    defaults := true // bool | Also include fields having the default value. (optional) (default to false)
    externalcas := true // bool | Enables export of external CAs (i.e. CAs where there's only a certificate and nothing else) (optional) (default to false)
    include := []string{"Inner_example"} // []string | Names of items/types to include in the export. The syntax is identical to that of exclude. For items of types that aren't listed, everything is included. (optional)
    exclude := []string{"Inner_example"} // []string | Names of items/types to exclude in the export, separated by semicolon. Type and name is separated by a colon, and wildcards \"\\*\" are allowed. Both are case-insensitive. E.g. exclude=\"\\*:Example CA;cryptotoken:Example\\*;systemconfiguration:\\*\".  Supported types are: ACMECONFIG/acme-config, CA/certification-authorities,  CRYPTOTOKEN/crypto-tokens, PUBLISHER/publishers, APPROVALPROFILE/approval-profiles, CERTPROFILE/certificate-profiles, EEPROFILE/end-entity-profiles, SERVICE/services, ROLE/admin-roles, KEYBINDING/internal-key-bindings, ADMINPREFS/admin-preferences, OCSPCONFIG/ocsp-configuration, PEERCONNECTOR/peer-connectors, SCEPCONFIG/scep-config, CMPCONFIG/cmp-config, ESTCONFIG/est-config, VALIDATOR/validators, CTLOG/ct-logs, EXTENDEDKEYUSAGE/extended-key-usage, CERTEXTENSION/custom-certificate-extensions,  OAUTHKEY/trusted-oauth-providers, AVAILABLEPROTOCOLS/available-protocols (optional)

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
    resp, r, err := apiClient.V1ConfigdumpApi.GetZipExport(context.Background()).Ignoreerrors(ignoreerrors).Defaults(defaults).Externalcas(externalcas).Include(include).Exclude(exclude).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `V1ConfigdumpApi.GetZipExport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetZipExport`: []string
    fmt.Fprintf(os.Stdout, "Response from `V1ConfigdumpApi.GetZipExport`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetZipExportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ignoreerrors** | **bool** | Print a warning instead of aborting and throwing an exception on errors. | [default to false]
 **defaults** | **bool** | Also include fields having the default value. | [default to false]
 **externalcas** | **bool** | Enables export of external CAs (i.e. CAs where there&#39;s only a certificate and nothing else) | [default to false]
 **include** | **[]string** | Names of items/types to include in the export. The syntax is identical to that of exclude. For items of types that aren&#39;t listed, everything is included. | 
 **exclude** | **[]string** | Names of items/types to exclude in the export, separated by semicolon. Type and name is separated by a colon, and wildcards \&quot;\\*\&quot; are allowed. Both are case-insensitive. E.g. exclude&#x3D;\&quot;\\*:Example CA;cryptotoken:Example\\*;systemconfiguration:\\*\&quot;.  Supported types are: ACMECONFIG/acme-config, CA/certification-authorities,  CRYPTOTOKEN/crypto-tokens, PUBLISHER/publishers, APPROVALPROFILE/approval-profiles, CERTPROFILE/certificate-profiles, EEPROFILE/end-entity-profiles, SERVICE/services, ROLE/admin-roles, KEYBINDING/internal-key-bindings, ADMINPREFS/admin-preferences, OCSPCONFIG/ocsp-configuration, PEERCONNECTOR/peer-connectors, SCEPCONFIG/scep-config, CMPCONFIG/cmp-config, ESTCONFIG/est-config, VALIDATOR/validators, CTLOG/ct-logs, EXTENDEDKEYUSAGE/extended-key-usage, CERTEXTENSION/custom-certificate-extensions,  OAUTHKEY/trusted-oauth-providers, AVAILABLEPROTOCOLS/available-protocols | 

### Return type

**[]string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/zip

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostJsonImport

> ConfigdumpResults PostJsonImport(ctx).Ignoreerrors(ignoreerrors).Initialize(initialize).Continue_(continue_).Overwrite(overwrite).Resolve(resolve).Body(body).Execute()

Put the configuration in JSON.



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
    ignoreerrors := true // bool | Add to warnings instead of aborting on errors. (optional) (default to false)
    initialize := true // bool | Generate initial certificate for CAs on import (optional) (default to false)
    continue_ := true // bool | Continue on errors. Default is to abort. (optional) (default to false)
    overwrite := "overwrite_example" // string | How to handle already existing configuration. Options are abort,skip,yes (optional) (default to "abort")
    resolve := "resolve_example" // string | How to resolve missing references. Options are abort,skip,default (optional) (default to "abort")
    body := "body_example" // string | JSON data in configdump format (optional)

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
    resp, r, err := apiClient.V1ConfigdumpApi.PostJsonImport(context.Background()).Ignoreerrors(ignoreerrors).Initialize(initialize).Continue_(continue_).Overwrite(overwrite).Resolve(resolve).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `V1ConfigdumpApi.PostJsonImport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostJsonImport`: ConfigdumpResults
    fmt.Fprintf(os.Stdout, "Response from `V1ConfigdumpApi.PostJsonImport`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostJsonImportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ignoreerrors** | **bool** | Add to warnings instead of aborting on errors. | [default to false]
 **initialize** | **bool** | Generate initial certificate for CAs on import | [default to false]
 **continue_** | **bool** | Continue on errors. Default is to abort. | [default to false]
 **overwrite** | **string** | How to handle already existing configuration. Options are abort,skip,yes | [default to &quot;abort&quot;]
 **resolve** | **string** | How to resolve missing references. Options are abort,skip,default | [default to &quot;abort&quot;]
 **body** | **string** | JSON data in configdump format | 

### Return type

[**ConfigdumpResults**](ConfigdumpResults.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostZipImport

> ConfigdumpResults PostZipImport(ctx).Zipfile(zipfile).Ignoreerrors(ignoreerrors).Initialize(initialize).Continue_(continue_).Overwrite(overwrite).Resolve(resolve).Execute()

Put the configuration as a ZIP file.



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
    zipfile := os.NewFile(1234, "some_file") // *os.File | A zipfile containing directories of YAML files. (optional)
    ignoreerrors := true // bool | Add to warnings instead of aborting on errors. (optional) (default to false)
    initialize := true // bool | Generate initial certificate for CAs on import (optional) (default to false)
    continue_ := true // bool | Continue on errors. Default is to abort. (optional) (default to false)
    overwrite := "overwrite_example" // string | How to handle already existing configuration. Options are abort,skip,yes (optional) (default to "abort")
    resolve := "resolve_example" // string | How to resolve missing references. Options are abort,skip,default (optional) (default to "abort")

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
    resp, r, err := apiClient.V1ConfigdumpApi.PostZipImport(context.Background()).Zipfile(zipfile).Ignoreerrors(ignoreerrors).Initialize(initialize).Continue_(continue_).Overwrite(overwrite).Resolve(resolve).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `V1ConfigdumpApi.PostZipImport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostZipImport`: ConfigdumpResults
    fmt.Fprintf(os.Stdout, "Response from `V1ConfigdumpApi.PostZipImport`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostZipImportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **zipfile** | ***os.File** | A zipfile containing directories of YAML files. | 
 **ignoreerrors** | **bool** | Add to warnings instead of aborting on errors. | [default to false]
 **initialize** | **bool** | Generate initial certificate for CAs on import | [default to false]
 **continue_** | **bool** | Continue on errors. Default is to abort. | [default to false]
 **overwrite** | **string** | How to handle already existing configuration. Options are abort,skip,yes | [default to &quot;abort&quot;]
 **resolve** | **string** | How to resolve missing references. Options are abort,skip,default | [default to &quot;abort&quot;]

### Return type

[**ConfigdumpResults**](ConfigdumpResults.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Status4

> RestResourceStatusRestResponse Status4(ctx).Execute()

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
    resp, r, err := apiClient.V1ConfigdumpApi.Status4(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `V1ConfigdumpApi.Status4``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Status4`: RestResourceStatusRestResponse
    fmt.Fprintf(os.Stdout, "Response from `V1ConfigdumpApi.Status4`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiStatus4Request struct via the builder pattern


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

