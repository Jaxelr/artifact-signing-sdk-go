package codesigning

// SignatureAlgorithm - The following algorithm identifiers are supported.
type SignatureAlgorithm string

const (
	SignatureAlgorithmES256 SignatureAlgorithm = "ES256"
	SignatureAlgorithmES256K SignatureAlgorithm = "ES256K"
	SignatureAlgorithmES384 SignatureAlgorithm = "ES384"
	SignatureAlgorithmES512 SignatureAlgorithm = "ES512"
	SignatureAlgorithmPS256 SignatureAlgorithm = "PS256"
	SignatureAlgorithmPS384 SignatureAlgorithm = "PS384"
	SignatureAlgorithmPS512 SignatureAlgorithm = "PS512"
	SignatureAlgorithmRS256 SignatureAlgorithm = "RS256"
	SignatureAlgorithmRS384 SignatureAlgorithm = "RS384"
	SignatureAlgorithmRS512 SignatureAlgorithm = "RS512"
)

// PossibleSignatureAlgorithmValues returns the possible values for the SignatureAlgorithm const type.
func PossibleSignatureAlgorithmValues() []SignatureAlgorithm {
	return []SignatureAlgorithm{	
		SignatureAlgorithmES256,
		SignatureAlgorithmES256K,
		SignatureAlgorithmES384,
		SignatureAlgorithmES512,
		SignatureAlgorithmPS256,
		SignatureAlgorithmPS384,
		SignatureAlgorithmPS512,
		SignatureAlgorithmRS256,
		SignatureAlgorithmRS384,
		SignatureAlgorithmRS512,
	}
}

// Status - Operation statuses, expanded to incorporate 2 sets. 1 - based on the ACS private preview version using the values:
// InProgress, Succeeded, Failed, Canceled, TimedOut, NotFound 2 - based on Azure's
// standard guidance see https://github.com/microsoft/api-guidelines/blob/vNext/azure/Guidelines.md?plain=1#L936: NotStarted,
// Running, Succeeded, Failed, Canceled
type Status string

const (
	StatusFailed Status = "Failed"
	StatusInProgress Status = "InProgress"
	StatusNotFound Status = "NotFound"
	StatusRunning Status = "Running"
	StatusSucceeded Status = "Succeeded"
	StatusTimedOut Status = "TimedOut"
)

// PossibleStatusValues returns the possible values for the Status const type.
func PossibleStatusValues() []Status {
	return []Status{	
		StatusFailed,
		StatusInProgress,
		StatusNotFound,
		StatusRunning,
		StatusSucceeded,
		StatusTimedOut,
	}
}

