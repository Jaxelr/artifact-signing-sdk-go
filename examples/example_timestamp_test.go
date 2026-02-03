package examples_test

import (
	"context"
	"fmt"
	"log"

	"github.com/jaxelr/artifact-signing-sdk-go/codesigning"
)

func Example_TimestampSignature() {
	// Example signature bytes (in practice, this would be the actual signature from BeginSign)
	signature := []byte{0x01, 0x02, 0x03, 0x04}

	// Create a timestamp client using Microsoft's default TSA URL
	tsClient := codesigning.NewTimestampClient("", nil) // Empty string uses DefaultMicrosoftTSAURL

	ctx := context.Background()
	result, err := tsClient.Timestamp(ctx, signature, &codesigning.TimestampOptions{
		RequestCertificates: true,
	})
	if err != nil {
		log.Fatalf("failed to timestamp signature: %v", err)
	}

	fmt.Printf("Signature timestamped at: %s\n", result.Time)
}
