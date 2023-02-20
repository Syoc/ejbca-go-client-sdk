package main

import (
    "context"
    "fmt"
    "github.com/Keyfactor/ejbca-go-client-sdk/api/ejbca"
    "log"
)

func main() {
    // environemnt variables for username password host
    // https://it-ca01.pkihosted-dev.c2company.com/ejbca/

    configuration := ejbca.NewConfiguration()
    configuration.Debug = true

    apiClient, err := ejbca.NewAPIClient(configuration)
    if err != nil {
        log.Fatal(err)
    }
    resp, err := apiClient.V1CaApi.GetCertificateAsPem(context.Background(), "CN=ManagementCA,OU=Certification Authorities,O=Internal Test,C=US").Execute()
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Response: %v", resp)
}
