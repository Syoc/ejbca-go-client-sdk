package main

import (
	"context"
	"fmt"
	"github.com/Keyfactor/ejbca-go-client-sdk/api/ejbca"
	"log"
)

func main() {
	configuration := ejbca.NewConfiguration()
	configuration.Host = "example.com"
	configuration.ClientCertificatePath = "auth_cert.pem" // Path to client certificate. The private key can be in the same file or in a file specified by the ClientCertificateKeyPath
	configuration.ClientCertificateKeyPath = "auth_key.pem"

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
