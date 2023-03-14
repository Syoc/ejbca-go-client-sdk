# Go Client SDK for Keyfactor EJBCA

The Go Client SDK for Keyfactor EJBCA enables management of EJBCA resources utilizing the Go programming language.

# Support for the Keyfactor EJBCA Go Client SDK
We welcome contributions.

The Keyfactor EJBCA Go Client SDK is open source and community supported, meaning that there is **no SLA** applicable for these tools.

###### To report a problem or suggest a new feature, use the **[Issues](../../issues)** tab. If you want to contribute actual bug fixes or proposed enhancements, use the **[Pull requests](../../pulls)** tab.

## Installation

Install the Go Client SDK for Keyfactor EJBCA using the `go get` command:

```shell
go get github.com/Keyfactor/ejbca-go-client-sdk
```

Put the package under your project folder and add the following in import:

```golang
import "github.com/Keyfactor/ejbca-go-client-sdk/api/ejbca"
```

## Configuration

Communication with the EJBCA REST API is authenticated using a client certificate. The client certificate
must be a PEM encoded X509v3 certificate with an unencrypted private key in PKCS#8 format. These fields can
be configured either using environment variables or with an `ejbca.Configuration` struct. Configure
the configuration struct as shown below:

```go
configuration := ejbca.NewConfiguration()
configuration.Host = "example.com"
configuration.ClientCertificatePath = "auth_cert.pem" // Path to client certificate. The private key can be in the same file or in a file specified by the ClientCertificateKeyPath
configuration.ClientCertificateKeyPath = "auth_key.pem"
```

The following environment variables can be used to configure the client as well:
```shell
export EJBCA_HOSTNAME="example.com"
export EJBCA_CLIENT_CERT_PATH="auth_cert.pem"
export EJBCA_CLIENT_CERT_KEY_PATH="auth_key.key"
```

Configuration of the EJBCA client via the `ejbca.Configuration` struct will override values set in environment variables.

If the EJBCA REST API uses a port other than 443, it can be configured with the `EJBCA_HOSTNAME` or the `Host` field in the configuration struct by
adding `:port` to the end of the hostname.

## Documentation for API Endpoints

All URIs are relative to *http://localhost/ejbca/ejbca-rest-api*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*V1CaApi* | [**CreateCrl**](docs/V1CaApi.md#createcrl) | **Post** /v1/ca/{issuer_dn}/createcrl | Create CRL(main, partition and delta) issued by this CA
*V1CaApi* | [**GetCertificateAsPem**](docs/V1CaApi.md#getcertificateaspem) | **Get** /v1/ca/{subject_dn}/certificate/download | Get PEM file with the active CA certificate chain
*V1CaApi* | [**GetLatestCrl**](docs/V1CaApi.md#getlatestcrl) | **Get** /v1/ca/{issuer_dn}/getLatestCrl | Returns the latest CRL issued by this CA
*V1CaApi* | [**ImportCrl**](docs/V1CaApi.md#importcrl) | **Post** /v1/ca/{issuer_dn}/importcrl | Import a certificate revocation list (CRL) for a CA
*V1CaApi* | [**ListCas**](docs/V1CaApi.md#listcas) | **Get** /v1/ca | Returns the Response containing the list of CAs with general information per CA as Json
*V1CaApi* | [**Status1**](docs/V1CaApi.md#status1) | **Get** /v1/ca/status | Get the status of this REST Resource
*V1CaManagementApi* | [**Activate**](docs/V1CaManagementApi.md#activate) | **Put** /v1/ca_management/{ca_name}/activate | Activate a CA
*V1CaManagementApi* | [**Deactivate**](docs/V1CaManagementApi.md#deactivate) | **Put** /v1/ca_management/{ca_name}/deactivate | Deactivate a CA
*V1CaManagementApi* | [**Status**](docs/V1CaManagementApi.md#status) | **Get** /v1/ca_management/status | Get the status of this REST Resource
*V1CertificateApi* | [**CertificateRequest**](docs/V1CertificateApi.md#certificaterequest) | **Post** /v1/certificate/certificaterequest | Enrollment with client generated keys for an existing End Entity
*V1CertificateApi* | [**EnrollKeystore**](docs/V1CertificateApi.md#enrollkeystore) | **Post** /v1/certificate/enrollkeystore | Keystore enrollment
*V1CertificateApi* | [**EnrollPkcs10Certificate**](docs/V1CertificateApi.md#enrollpkcs10certificate) | **Post** /v1/certificate/pkcs10enroll | Enrollment with client generated keys, using CSR subject
*V1CertificateApi* | [**FinalizeEnrollment**](docs/V1CertificateApi.md#finalizeenrollment) | **Post** /v1/certificate/{request_id}/finalize | Finalize enrollment
*V1CertificateApi* | [**GetCertificatesAboutToExpire**](docs/V1CertificateApi.md#getcertificatesabouttoexpire) | **Get** /v1/certificate/expire | Get a list of certificates that are about to expire
*V1CertificateApi* | [**RevocationStatus**](docs/V1CertificateApi.md#revocationstatus) | **Get** /v1/certificate/{issuer_dn}/{certificate_serial_number}/revocationstatus | Checks revocation status of the specified certificate
*V1CertificateApi* | [**RevokeCertificate**](docs/V1CertificateApi.md#revokecertificate) | **Put** /v1/certificate/{issuer_dn}/{certificate_serial_number}/revoke | Revokes the specified certificate
*V1CertificateApi* | [**SearchCertificates**](docs/V1CertificateApi.md#searchcertificates) | **Post** /v1/certificate/search | Searches for certificates confirming given criteria.
*V1CertificateApi* | [**Status2**](docs/V1CertificateApi.md#status2) | **Get** /v1/certificate/status | Get the status of this REST Resource
*V1ConfigdumpApi* | [**GetJsonConfigdump**](docs/V1ConfigdumpApi.md#getjsonconfigdump) | **Get** /v1/configdump | Get the configuration in JSON.
*V1ConfigdumpApi* | [**GetJsonConfigdumpForType**](docs/V1ConfigdumpApi.md#getjsonconfigdumpfortype) | **Get** /v1/configdump/{type} | Get the configuration for type in JSON.
*V1ConfigdumpApi* | [**GetJsonConfigdumpForTypeAndSetting**](docs/V1ConfigdumpApi.md#getjsonconfigdumpfortypeandsetting) | **Get** /v1/configdump/{type}/{setting} | Get the configuration for a type and setting in JSON.
*V1ConfigdumpApi* | [**GetZipExport**](docs/V1ConfigdumpApi.md#getzipexport) | **Get** /v1/configdump/configdump.zip | Get the configuration as a ZIP file.
*V1ConfigdumpApi* | [**PostJsonImport**](docs/V1ConfigdumpApi.md#postjsonimport) | **Post** /v1/configdump | Put the configuration in JSON.
*V1ConfigdumpApi* | [**PostZipImport**](docs/V1ConfigdumpApi.md#postzipimport) | **Post** /v1/configdump/configdump.zip | Put the configuration as a ZIP file.
*V1ConfigdumpApi* | [**Status4**](docs/V1ConfigdumpApi.md#status4) | **Get** /v1/configdump/status | Get the status of this REST Resource
*V1CryptotokenApi* | [**Activate1**](docs/V1CryptotokenApi.md#activate1) | **Put** /v1/cryptotoken/{cryptotoken_name}/activate | Activate a Crypto Token
*V1CryptotokenApi* | [**Deactivate1**](docs/V1CryptotokenApi.md#deactivate1) | **Put** /v1/cryptotoken/{cryptotoken_name}/deactivate | Deactivate a Crypto Token
*V1CryptotokenApi* | [**GenerateKeys**](docs/V1CryptotokenApi.md#generatekeys) | **Post** /v1/cryptotoken/{cryptotoken_name}/generatekeys | Generate keys
*V1CryptotokenApi* | [**RemoveKeys**](docs/V1CryptotokenApi.md#removekeys) | **Post** /v1/cryptotoken/{cryptotoken_name}/{key_pair_alias}/removekeys | Remove keys
*V1CryptotokenApi* | [**Status5**](docs/V1CryptotokenApi.md#status5) | **Get** /v1/cryptotoken/status | Get the status of this REST Resource
*V1EndentityApi* | [**Add**](docs/V1EndentityApi.md#add) | **Post** /v1/endentity | Add new end entity, if it does not exist
*V1EndentityApi* | [**Delete**](docs/V1EndentityApi.md#delete) | **Delete** /v1/endentity/{endentity_name} | Deletes end entity
*V1EndentityApi* | [**Revoke**](docs/V1EndentityApi.md#revoke) | **Put** /v1/endentity/{endentity_name}/revoke | Revokes all end entity certificates
*V1EndentityApi* | [**Search**](docs/V1EndentityApi.md#search) | **Post** /v1/endentity/search | Searches for end entity confirming given criteria.
*V1EndentityApi* | [**Setstatus**](docs/V1EndentityApi.md#setstatus) | **Post** /v1/endentity/{endentity_name}/setstatus | Edits end entity setting new status
*V1EndentityApi* | [**Status6**](docs/V1EndentityApi.md#status6) | **Get** /v1/endentity/status | Get the status of this REST Resource
*V1SshApi* | [**Pubkey**](docs/V1SshApi.md#pubkey) | **Get** /v1/ssh/{ca_name}/pubkey | Retrieves a CA&#39;s public key in SSH format.
*V1SshApi* | [**Status8**](docs/V1SshApi.md#status8) | **Get** /v1/ssh/status | Get the status of this REST Resource
*V2CertificateApi* | [**GetCertificateProfileInfo**](docs/V2CertificateApi.md#getcertificateprofileinfo) | **Get** /v2/certificate/profile/{profile_name} | Get Certificate Profile Info.
*V2CertificateApi* | [**SearchCertificates1**](docs/V2CertificateApi.md#searchcertificates1) | **Post** /v2/certificate/search | Searches for certificates confirming given criteria and pagination.
*V2CertificateApi* | [**Status3**](docs/V2CertificateApi.md#status3) | **Get** /v2/certificate/status | Get the status of this REST Resource
*V2EndentityApi* | [**GetAuthorizedEndEntityProfiles**](docs/V2EndentityApi.md#getauthorizedendentityprofiles) | **Get** /v2/endentity/profiles/authorized | List of authorized end entity profiles for the current admin.
*V2EndentityApi* | [**Profile**](docs/V2EndentityApi.md#profile) | **Get** /v2/endentity/profile/{endentity_profile_name} | Get End Entity Profile content
*V2EndentityApi* | [**SortedSearch**](docs/V2EndentityApi.md#sortedsearch) | **Post** /v2/endentity/search | Searches and sorts for end entity conforming given criteria.
*V2EndentityApi* | [**Status7**](docs/V2EndentityApi.md#status7) | **Get** /v2/endentity/status | Get the status of this REST Resource


## Documentation For Models

 - [AddEndEntityRestRequest](docs/AddEndEntityRestRequest.md)
 - [AuthorizedEEPsRestResponse](docs/AuthorizedEEPsRestResponse.md)
 - [CaInfoRestResponse](docs/CaInfoRestResponse.md)
 - [CaInfosRestResponse](docs/CaInfosRestResponse.md)
 - [CertificateProfileInfoRestResponseV2](docs/CertificateProfileInfoRestResponseV2.md)
 - [CertificateRequestRestRequest](docs/CertificateRequestRestRequest.md)
 - [CertificateRestResponse](docs/CertificateRestResponse.md)
 - [CertificateRestResponseV2](docs/CertificateRestResponseV2.md)
 - [CertificatesRestResponse](docs/CertificatesRestResponse.md)
 - [ConfigdumpResults](docs/ConfigdumpResults.md)
 - [CreateCrlRestResponse](docs/CreateCrlRestResponse.md)
 - [CrlRestResponse](docs/CrlRestResponse.md)
 - [CryptoTokenActivationRestRequest](docs/CryptoTokenActivationRestRequest.md)
 - [CryptoTokenKeyGenerationRestRequest](docs/CryptoTokenKeyGenerationRestRequest.md)
 - [EndEntityProfileResponse](docs/EndEntityProfileResponse.md)
 - [EndEntityProfileRestResponse](docs/EndEntityProfileRestResponse.md)
 - [EndEntityRestResponse](docs/EndEntityRestResponse.md)
 - [EndEntityRevocationRestRequest](docs/EndEntityRevocationRestRequest.md)
 - [EnrollCertificateRestRequest](docs/EnrollCertificateRestRequest.md)
 - [ExpiringCertificatesRestResponse](docs/ExpiringCertificatesRestResponse.md)
 - [ExtendedInformationRestRequestComponent](docs/ExtendedInformationRestRequestComponent.md)
 - [ExtendedInformationRestResponseComponent](docs/ExtendedInformationRestResponseComponent.md)
 - [FinalizeRestRequest](docs/FinalizeRestRequest.md)
 - [KeyStoreRestRequest](docs/KeyStoreRestRequest.md)
 - [Pagination](docs/Pagination.md)
 - [PaginationRestResponseComponent](docs/PaginationRestResponseComponent.md)
 - [PaginationSummary](docs/PaginationSummary.md)
 - [RestResourceStatusRestResponse](docs/RestResourceStatusRestResponse.md)
 - [RevokeStatusRestResponse](docs/RevokeStatusRestResponse.md)
 - [SearchCertificateCriteriaRestRequest](docs/SearchCertificateCriteriaRestRequest.md)
 - [SearchCertificateSortRestRequest](docs/SearchCertificateSortRestRequest.md)
 - [SearchCertificatesRestRequest](docs/SearchCertificatesRestRequest.md)
 - [SearchCertificatesRestRequestV2](docs/SearchCertificatesRestRequestV2.md)
 - [SearchCertificatesRestResponse](docs/SearchCertificatesRestResponse.md)
 - [SearchCertificatesRestResponseV2](docs/SearchCertificatesRestResponseV2.md)
 - [SearchEndEntitiesRestRequest](docs/SearchEndEntitiesRestRequest.md)
 - [SearchEndEntitiesRestRequestV2](docs/SearchEndEntitiesRestRequestV2.md)
 - [SearchEndEntitiesRestResponse](docs/SearchEndEntitiesRestResponse.md)
 - [SearchEndEntitiesSortRestRequest](docs/SearchEndEntitiesSortRestRequest.md)
 - [SearchEndEntityCriteriaRestRequest](docs/SearchEndEntityCriteriaRestRequest.md)
 - [SetEndEntityStatusRestRequest](docs/SetEndEntityStatusRestRequest.md)
 - [SshPublicKeyRestResponse](docs/SshPublicKeyRestResponse.md)


## Application Notes
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 1.0
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen
