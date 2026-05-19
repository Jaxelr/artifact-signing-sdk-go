//go:build integration
// +build integration

package tests

import (
	"context"
	"crypto/sha256"
	"io"
	"testing"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/jaxelr/artifact-signing-sdk-go/codesigning"
)

const (
	testEndpoint    = "https://eus.codesigning.azure.net"
	testAccountName = "jaxelr"
	testProfileName = "jaxelr-pt"
)

func TestGetSignCertificateChain(t *testing.T) {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		t.Fatalf("failed to create credential: %v", err)
	}

	client, err := codesigning.NewCertificateProfileClient(testEndpoint, cred, nil)
	if err != nil {
		t.Fatalf("failed to create client: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	resp, err := client.GetSignCertificateChain(ctx, testAccountName, testProfileName, nil)
	if err != nil {
		t.Fatalf("GetSignCertificateChain failed: %v", err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("failed to read response body: %v", err)
	}

	t.Logf("Certificate chain retrieved successfully, size: %d bytes", len(data))
}

func TestGetSignCertificateRoot(t *testing.T) {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		t.Fatalf("failed to create credential: %v", err)
	}

	client, err := codesigning.NewCertificateProfileClient(testEndpoint, cred, nil)
	if err != nil {
		t.Fatalf("failed to create client: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	resp, err := client.GetSignCertificateRoot(ctx, testAccountName, testProfileName, nil)
	if err != nil {
		t.Fatalf("GetSignCertificateRoot failed: %v", err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("failed to read response body: %v", err)
	}

	t.Logf("Root certificate retrieved successfully, size: %d bytes", len(data))
}

func TestGetSignEku(t *testing.T) {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		t.Fatalf("failed to create credential: %v", err)
	}

	client, err := codesigning.NewCertificateProfileClient(testEndpoint, cred, nil)
	if err != nil {
		t.Fatalf("failed to create client: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	resp, err := client.GetSignEku(ctx, testAccountName, testProfileName, nil)
	if err != nil {
		t.Fatalf("GetSignEku failed: %v", err)
	}

	t.Logf("EKU values retrieved: %v", resp.StringArray)
}

func TestSignAndTimestamp(t *testing.T) {
	// Create Azure credential
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		t.Fatalf("failed to create credential: %v", err)
	}

	// Create the signing client
	client, err := codesigning.NewCertificateProfileClient(testEndpoint, cred, nil)
	if err != nil {
		t.Fatalf("failed to create client: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	// Create a sample payload and compute its SHA-256 digest
	payload := []byte("Hello, this is a test payload for signing and timestamping!")
	hash := sha256.Sum256(payload)
	digest := hash[:]

	t.Logf("Payload digest (SHA-256): %x", digest)

	// Step 1: Sign the digest using Artifact Signing
	poller, err := client.BeginSign(ctx, testAccountName, testProfileName, codesigning.SignRequest{
		Digest:             digest,
		SignatureAlgorithm: to.Ptr(codesigning.SignatureAlgorithmRS256),
	}, nil)
	if err != nil {
		t.Fatalf("failed to start signing: %v", err)
	}

	result, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		t.Fatalf("failed to complete signing: %v", err)
	}

	if result.Signature == nil || len(result.Signature) == 0 {
		t.Fatal("signing completed but no signature was returned")
	}

	t.Logf("Signing completed with status: %s", *result.Status)
	t.Logf("Signature length: %d bytes", len(result.Signature))

	// Step 2: Timestamp the signature 
	tsClient := codesigning.NewTimestampClient("", nil) // Uses DefaultMicrosoftTSAURL

	tsResult, err := tsClient.Timestamp(ctx, result.Signature, &codesigning.TimestampOptions{
		RequestCertificates: to.Ptr(true),
	})
	if err != nil {
		t.Fatalf("failed to timestamp signature: %v", err)
	}

	t.Logf("Signature successfully timestamped at: %s", tsResult.Time)
	t.Logf("Timestamp serial number: %s", tsResult.SerialNumber)
}
