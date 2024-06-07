/*
Copyright 2024 Keyfactor

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

EJBCA REST Interface

API reference documentation.

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ejbca

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
)

// Authenticator is used to build HTTP clients with a built-in authentication mechanism. It's expected that
// the authenticator will handle the authentication of the client and set the necessary headers.
type Authenticator interface {
	GetHttpClient() (*http.Client, error)
}

// OAuth Authenticator

var _ Authenticator = &OAuthAuthenticator{}

// OAuthAuthenticator is an Authenticator that uses OAuth2 for authentication.
type OAuthAuthenticator struct {
	client *http.Client
}

type OAuthAuthenticatorBuilder struct {
	clientId          string
	clientSecret      string
	tokenUrl          string
	audience          string
	scopes            []string
	caCertificatePath string
	caCertificates    []*x509.Certificate
}

func NewOAuthAuthenticatorBuilder() *OAuthAuthenticatorBuilder {
	return &OAuthAuthenticatorBuilder{}
}

func (b *OAuthAuthenticatorBuilder) WithClientId(clientId string) *OAuthAuthenticatorBuilder {
	b.clientId = clientId
	return b
}

func (b *OAuthAuthenticatorBuilder) WithClientSecret(clientSecret string) *OAuthAuthenticatorBuilder {
	b.clientSecret = clientSecret
	return b
}

func (b *OAuthAuthenticatorBuilder) WithTokenUrl(tokenUrl string) *OAuthAuthenticatorBuilder {
	b.tokenUrl = tokenUrl
	return b
}

func (b *OAuthAuthenticatorBuilder) WithScopes(scopes []string) *OAuthAuthenticatorBuilder {
	b.scopes = scopes
	return b
}

func (b *OAuthAuthenticatorBuilder) WithAudience(audience string) *OAuthAuthenticatorBuilder {
	b.audience = audience
	return b
}

func (b *OAuthAuthenticatorBuilder) WithCaCertificatePath(caCertificatePath string) *OAuthAuthenticatorBuilder {
	b.caCertificatePath = caCertificatePath
	return b
}

func (b *OAuthAuthenticatorBuilder) WithCaCertificates(caCertificates []*x509.Certificate) *OAuthAuthenticatorBuilder {
	b.caCertificates = caCertificates
	return b
}

func (b *OAuthAuthenticatorBuilder) Build() (Authenticator, error) {
	config := &clientcredentials.Config{
		ClientID:     b.clientId,
		ClientSecret: b.clientSecret,
		TokenURL:     b.tokenUrl,
		Scopes:       b.scopes,
	}

	if b.audience != "" {
		config.EndpointParams = map[string][]string{
			"audience": {
				b.audience,
			},
		}
	}

	tokenSource := config.TokenSource(context.Background())
	oauthTransport := &oauth2.Transport{
		Base:   http.DefaultTransport,
		Source: tokenSource,
	}

	if b.caCertificates == nil {
		var err error
		b.caCertificates, err = findCaCertificate(b.caCertificatePath)
		if err != nil {
			return nil, fmt.Errorf("failed to find CA certificates: %w", err)
		}
	}

	if len(b.caCertificates) > 0 {
		tlsConfig := &tls.Config{
			Renegotiation: tls.RenegotiateOnceAsClient,
		}

		tlsConfig.RootCAs = x509.NewCertPool()
		for _, caCert := range b.caCertificates {
			tlsConfig.RootCAs.AddCert(caCert)
		}

		customTransport := http.DefaultTransport.(*http.Transport).Clone()
		customTransport.TLSClientConfig = tlsConfig
		customTransport.TLSHandshakeTimeout = 10 * time.Second

		// Wrap the custom transport with the oauth2.Transport
		oauthTransport.Base = customTransport
	}

	client := &http.Client{
		Transport: oauthTransport,
	}

	return &OAuthAuthenticator{client: client}, nil
}

func (a *OAuthAuthenticator) GetHttpClient() (*http.Client, error) {
	return a.client, nil
}

// mTLS Authenticator

var _ Authenticator = &MTLSAuthenticator{}

type MTLSAuthenticator struct {
	client *http.Client
}

type MTLSAuthenticatorBuilder struct {
	clientCertificate        *tls.Certificate
	clientCertificatePath    string
	clientCertificateKeyPath string
	caCertificatePath        string
	caCertificates           []*x509.Certificate
}

func NewMTLSAuthenticatorBuilder() *MTLSAuthenticatorBuilder {
	return &MTLSAuthenticatorBuilder{}
}

func (b *MTLSAuthenticatorBuilder) WithClientCertificatePath(clientCertificatePath string) *MTLSAuthenticatorBuilder {
	b.clientCertificatePath = clientCertificatePath
	return b
}

func (b *MTLSAuthenticatorBuilder) WithClientCertificateKeyPath(clientCertificateKeyPath string) *MTLSAuthenticatorBuilder {
	b.clientCertificateKeyPath = clientCertificateKeyPath
	return b
}

func (b *MTLSAuthenticatorBuilder) WithCaCertificatePath(caCertificatePath string) *MTLSAuthenticatorBuilder {
	b.caCertificatePath = caCertificatePath
	return b
}

func (b *MTLSAuthenticatorBuilder) WithClientCertificate(clientCertificate *tls.Certificate) *MTLSAuthenticatorBuilder {
	b.clientCertificate = clientCertificate
	return b
}

func (b *MTLSAuthenticatorBuilder) WithCaCertificates(caCertificates []*x509.Certificate) *MTLSAuthenticatorBuilder {
	b.caCertificates = caCertificates
	return b
}

func (b *MTLSAuthenticatorBuilder) findClientCert() error {
	if b.clientCertificatePath == "" {
		return fmt.Errorf("no client certificate path specified")
	}
	// Read and parse the passed certificate file which could contain the certificate and private key
	buf, err := os.ReadFile(b.clientCertificatePath)
	if err != nil {
		return err
	}
	certificates, privKey, err := decodePEMBytes(buf)
	if err != nil {
		return fmt.Errorf("failed to decode PEM bytes: %w", err)
	}
	if len(privKey) <= 0 {
		// if no private key was found, see if a path was specified to a private key
		if b.clientCertificateKeyPath == "" {
			return fmt.Errorf("no private key found in %s and no path to a private key specified", b.clientCertificatePath)
		}
		buf, err = os.ReadFile(b.clientCertificateKeyPath)
		if err != nil {
			return err
		}
		_, privKey, err = decodePEMBytes(buf)
		if err != nil {
			return err
		}
		if len(privKey) <= 0 {
			return fmt.Errorf("didn't find private key in path %s", b.clientCertificateKeyPath)
		}
	}
	if len(certificates) <= 0 {
		return fmt.Errorf("didn't find certificate in file at path %s", b.clientCertificateKeyPath)
	}
	cert, err := tls.X509KeyPair(pem.EncodeToMemory(certificates[0]), privKey)
	if err != nil {
		return err
	}

	b.clientCertificate = &cert
	return nil
}

func (b *MTLSAuthenticatorBuilder) Build() (Authenticator, error) {
	var err error
	if b.clientCertificate == nil {
		err = b.findClientCert()
		if err != nil {
			return nil, fmt.Errorf("failed to find client certificate: %w", err)
		}
	}

	if len(b.caCertificates) == 0 {
		b.caCertificates, err = findCaCertificate(b.caCertificatePath)
		if err != nil {
			return nil, fmt.Errorf("failed to find CA certificates: %w", err)
		}
	}

	tlsConfig := &tls.Config{
		Certificates:  []tls.Certificate{*b.clientCertificate},
		Renegotiation: tls.RenegotiateOnceAsClient,
	}

	if len(b.caCertificates) > 0 {
		tlsConfig.RootCAs = x509.NewCertPool()
		for _, caCert := range b.caCertificates {
			tlsConfig.RootCAs.AddCert(caCert)
		}
	}

	customTransport := http.DefaultTransport.(*http.Transport).Clone()
	customTransport.TLSClientConfig = tlsConfig
	customTransport.TLSHandshakeTimeout = 10 * time.Second

	httpClient := &http.Client{
		Transport: customTransport,
		Timeout:   10 * time.Second,
	}

	return &MTLSAuthenticator{client: httpClient}, nil
}

func (a *MTLSAuthenticator) GetHttpClient() (*http.Client, error) {
	return a.client, nil
}

func findCaCertificate(caCertificatePath string) ([]*x509.Certificate, error) {
	if caCertificatePath == "" {
		return nil, nil
	}

	buf, err := os.ReadFile(caCertificatePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read CA certificate file at path %s: %w", caCertificatePath, err)
	}
	// Decode the PEM encoded certificates into a slice of PEM blocks
	chainBlocks, _, err := decodePEMBytes(buf)
	if err != nil {
		return nil, err
	}
	if len(chainBlocks) <= 0 {
		return nil, fmt.Errorf("didn't find certificate in file at path %s", caCertificatePath)
	}

	caChain := []*x509.Certificate{}
	for _, block := range chainBlocks {
		// Parse the PEM block into an x509 certificate
		cert, err := x509.ParseCertificate(block.Bytes)
		if err != nil {
			return nil, fmt.Errorf("failed to parse CA certificate: %w", err)
		}

		caChain = append(caChain, cert)
	}

	return caChain, nil
}

func decodePEMBytes(buf []byte) ([]*pem.Block, []byte, error) {
	var privKey []byte
	var certificates []*pem.Block
	var block *pem.Block
	for {
		block, buf = pem.Decode(buf)
		if block == nil {
			break
		} else if strings.Contains(block.Type, "PRIVATE KEY") {
			privKey = pem.EncodeToMemory(block)
		} else {
			certificates = append(certificates, block)
		}
	}
	return certificates, privKey, nil
}
