package main

import (
	"context"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"github.com/Keyfactor/ejbca-go-client-sdk/api/ejbca"
	"io"
	"log"
	"strings"
)

func main() {
	configuration := ejbca.NewConfiguration()
	configuration.Debug = true

	apiClient, err := ejbca.NewAPIClient(configuration)
	if err != nil {
		log.Fatal(err)
	}

	dn := "CN=ManagementCA,OU=Certification Authorities,O=Internal Test,C=US"

	chain, err := getCaChain(context.Background(), apiClient, dn)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("CA with DN \"%s\" has a chain with %d certificates\n", dn, len(chain))
	fmt.Printf("Chain:\n%s", compileCertificatesToPemString(chain))
}

func getCaChain(ctx context.Context, client *ejbca.APIClient, issuerDn string) ([]*x509.Certificate, error) {
	caResp, err := client.V1CaApi.GetCertificateAsPem(ctx, issuerDn).Execute()
	if err != nil {
		return nil, err
	}

	encodedBytes, err := io.ReadAll(caResp.Body) // EJBCA returns CA chain as a single PEM file
	if err != nil {
		return nil, err
	}

	// Decode PEM file into a slice of der bytes
	var block *pem.Block
	var derBytes []byte
	for {
		block, encodedBytes = pem.Decode(encodedBytes)
		if block == nil {
			break
		}
		derBytes = append(derBytes, block.Bytes...)
	}

	certificates, err := x509.ParseCertificates(derBytes)
	if err != nil {
		return nil, err
	}

	return certificates, nil
}

// compileCertificatesToPemString takes a slice of x509 certificates and returns a string containing the certificates in PEM format
// If an error occurred, the function logs the error and continues to parse the remaining objects.
func compileCertificatesToPemString(certificates []*x509.Certificate) string {
	var pemBuilder strings.Builder

	for _, certificate := range certificates {
		err := pem.Encode(&pemBuilder, &pem.Block{
			Type:  "CERTIFICATE",
			Bytes: certificate.Raw,
		})
		if err != nil {
			fmt.Printf("Failed to encode certificate with serial number %s to PEM format. Continuing anyway (%s)", certificate.SerialNumber.String(), err.Error())
		}
	}

	return pemBuilder.String()
}
