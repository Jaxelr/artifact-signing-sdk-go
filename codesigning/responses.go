package codesigning

import "io"

// CertificateProfileClientGetSignCertificateChainResponse contains the response from method CertificateProfileClient.GetSignCertificateChain.
type CertificateProfileClientGetSignCertificateChainResponse struct {
// Body contains the streaming response.
	Body io.ReadCloser
}

// CertificateProfileClientGetSignCertificateRootResponse contains the response from method CertificateProfileClient.GetSignCertificateRoot.
type CertificateProfileClientGetSignCertificateRootResponse struct {
// Body contains the streaming response.
	Body io.ReadCloser
}

// CertificateProfileClientGetSignEkuResponse contains the response from method CertificateProfileClient.GetSignEku.
type CertificateProfileClientGetSignEkuResponse struct {
// Array of Get200ApplicationJsonItemsItem
	StringArray []*string
}

// CertificateProfileClientGetSignOperationResponse contains the response from method CertificateProfileClient.GetSignOperation.
type CertificateProfileClientGetSignOperationResponse struct {
	SignStatus
}

// CertificateProfileClientSignResponse contains the response from method CertificateProfileClient.BeginSign.
type CertificateProfileClientSignResponse struct {
	SignStatus
}

