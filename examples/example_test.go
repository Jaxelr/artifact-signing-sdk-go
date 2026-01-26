package examples_test

import (
	"context"
	"fmt"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/jaxelr/artifact-signing-sdk-go/codesigning"
)

func ExampleCertificateProfileClient_GetSignCertificateChain() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to create credential: %v", err)
	}

	client, err := codesigning.NewCertificateProfileClient(
		"https://your-endpoint.codesigning.azure.net",
		cred,
		nil,
	)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}

	ctx := context.Background()
	resp, err := client.GetSignCertificateChain(ctx, "account-name", "profile-name", nil)
	if err != nil {
		log.Fatalf("failed to get certificate chain: %v", err)
	}
	defer resp.Body.Close()

	fmt.Println("Certificate chain retrieved successfully")
}

func ExampleCertificateProfileClient_BeginSign() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to create credential: %v", err)
	}

	client, err := codesigning.NewCertificateProfileClient(
		"https://your-endpoint.codesigning.azure.net",
		cred,
		nil,
	)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}

	ctx := context.Background()

	// Example digest (in practice, this would be a real SHA-256 hash)
	digest := []byte{0x01, 0x02, 0x03, 0x04}

	poller, err := client.BeginSign(ctx, "account-name", "profile-name", codesigning.SignRequest{
		Digest:             digest,
		SignatureAlgorithm: to.Ptr(codesigning.SignatureAlgorithmRS256),
	}, nil)
	if err != nil {
		log.Fatalf("failed to start signing: %v", err)
	}

	result, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to complete signing: %v", err)
	}

	fmt.Printf("Signing completed with status: %s\n", *result.Status)
}
