package dlib

import (
	"crypto/tls"
	"crypto/x509"
	"net/http"
	"time"

	"google.golang.org/grpc/credentials"
)

// GetHTTPSClient creates a new client for connecting to HTTPS servers.
func GetHTTPSClient(cert string) (*http.Client, error) {
	roots := x509.NewCertPool()
	ok := roots.AppendCertsFromPEM([]byte(cert))
	if !ok {
		return nil, NewError(500, "unable to parse root certificate")
	}

	tc := tls.Config{RootCAs: roots}
	tc.BuildNameToCertificate()
	tr := &http.Transport{TLSClientConfig: &tc}
	return &http.Client{
		Transport: tr,
		Timeout:   time.Second * 30,
	}, nil
}

// GetGRPCServerCredentials returns a set of server LS credentials for gRPC.
func GetGRPCServerCredentials(cert, key string) (credentials.TransportCredentials, error) {
	crt, err := tls.X509KeyPair([]byte(cert), []byte(key))
	if err != nil {
		return nil, err
	}

	creds := credentials.NewServerTLSFromCert(&crt)
	return creds, nil
}

// GetGRPCClientCredentials returns a set of client TLS credentials for gRPC.
func GetGRPCClientCredentials(cert string) (credentials.TransportCredentials, error) {
	pool := x509.NewCertPool()
	ok := pool.AppendCertsFromPEM([]byte(cert))
	if !ok {
		return nil, NewError(500, "unable to parse certificate")
	}

	creds := credentials.NewClientTLSFromCert(pool, "")
	return creds, nil
}
