package codesigning

// CertificateProfileClientBeginSignOptions contains the optional parameters for the CertificateProfileClient.BeginSign method.
type CertificateProfileClientBeginSignOptions struct {
// Optional certificate thumbprint retrieve from existing leaf
	CertificateThumbprint *string

// Optional client version information from the client calling this request.
	ClientVersion *string

// Resumes the long-running operation from the provided token.
	ResumeToken string

// Optional correlation ID provided by the caller.
	XCorrelationID *string
}

// CertificateProfileClientGetSignCertificateChainOptions contains the optional parameters for the CertificateProfileClient.GetSignCertificateChain
// method.
type CertificateProfileClientGetSignCertificateChainOptions struct {
	// placeholder for future optional parameters
}

// CertificateProfileClientGetSignCertificateRootOptions contains the optional parameters for the CertificateProfileClient.GetSignCertificateRoot
// method.
type CertificateProfileClientGetSignCertificateRootOptions struct {
	// placeholder for future optional parameters
}

// CertificateProfileClientGetSignEkuOptions contains the optional parameters for the CertificateProfileClient.GetSignEku
// method.
type CertificateProfileClientGetSignEkuOptions struct {
	// placeholder for future optional parameters
}

// CertificateProfileClientGetSignOperationOptions contains the optional parameters for the CertificateProfileClient.GetSignOperation
// method.
type CertificateProfileClientGetSignOperationOptions struct {
	// placeholder for future optional parameters
}

