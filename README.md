# Artifact Signing SDK for Go

This module provides a Go client library for signing and retrieving relevant signing data using Artifact Signing.

> [!IMPORTANT]
> This is not an official Azure SDK. This library is developed and maintained by the community.
> For dataplane SDKs, there isnt currently an official Azure SDK for Artifact Signing.

## Installation

```bash
go get github.com/jaxelr/artifact-signing-sdk-go/codesigning
```

## Usage

> [!NOTE]
> This example is missing timestamping for simplicity. In production scenarios, it is recommended to include timestamping when signing artifacts, since these are short lived keys.

```go
package main

import (
	"context"
	"crypto/sha256"
	"fmt"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/jaxelr/artifact-signing-sdk-go/codesigning"
)

func main() {
	// Create a credential using Azure Identity
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to create credential: %v", err)
	}

	// Create the client
	client, err := codesigning.NewCertificateProfileClient(
		"https://your-endpoint.codesigning.azure.net",
		cred,
		nil,
	)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}

	ctx := context.Background()
	accountName := "your-account-name"
	profileName := "your-profile-name"

	// Create a SHA256 digest of the content to sign
	content := []byte("content to sign")
	digest := sha256.Sum256(content)

	// Create the sign request
	signRequest := codesigning.SignRequest{
		Digest:             digest[:],
		SignatureAlgorithm: to.Ptr(codesigning.SignatureAlgorithmRS256),
	}

	// Initiate the signing operation
	poller, err := client.BeginSign(ctx, accountName, profileName, signRequest, nil)
	if err != nil {
		log.Fatalf("failed to start signing operation: %v", err)
	}

	// Wait for the signing operation to complete
	resp, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to complete signing operation: %v", err)
	}

	fmt.Printf("Signing operation completed. Signature length: %d bytes\n", len(resp.Signature))
}
```

## Features

- **Sign artifacts** - Sign digests using Artifact Signing
- **Get certificate chain** - Retrieves the full certificate chain for a signing profile
- **Get root certificate** - Retrieves the root certificate
- **Get EKU** - Retrieve Extended Key Usage values for a profile

## Requirements

- Go 1.23 or later
- Azure subscription with Artifact Signing resource

## References

For more information on what is Artifact Signing, please see [here](https://learn.microsoft.com/en-us/azure/artifact-signing/overview)

## License

MIT License - see [LICENSE](LICENSE) for details.