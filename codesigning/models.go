package codesigning

type ErrorDetail struct {
// One of a server-defined set of error codes.
	Code *string

// A human-readable representation of the error.
	Message *string

// The target of the error.
	Target *string
}

type ErrorResponse struct {
	ErrorDetail *ErrorDetail
}

type SignRequest struct {
// REQUIRED; Digest to sign.
	Digest []byte

// REQUIRED; The following algorithm identifiers are supported.
	SignatureAlgorithm *SignatureAlgorithm

// List of authenticode hash.
	AuthenticodeHashList [][]byte

// List of full file hash.
	FileHashList [][]byte
}

type SignStatus struct {
// REQUIRED; ID of the operation.
	OperationID *string

// REQUIRED; Operation statuses, expanded to incorporate 2 sets. 1 - based on the ACS private preview version using the values:
// InProgress, Succeeded, Failed, Canceled, TimedOut, NotFound 2 - based on Azure's
// standard guidance see https://github.com/microsoft/api-guidelines/blob/vNext/azure/Guidelines.md?plain=1#L936: NotStarted,
// Running, Succeeded, Failed, Canceled
	Status *Status

// Signature of the requested digest.
	Signature []byte

// Signing certificate corresponding to the private key used to sign the requested digest.
	SigningCertificate []byte
}

