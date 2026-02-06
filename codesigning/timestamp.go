package codesigning

import (
	"bytes"
	"context"
	"crypto"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/digitorus/timestamp"
)

const DefaultMicrosoftTSAURL = "http://timestamp.acs.microsoft.com"

// TimestampClient is a client for requesting RFC 3161 timestamps from a Time Stamping Authority (TSA).
type TimestampClient struct {
	tsaURL     string
	httpClient *http.Client
}

// TimestampClientOptions defines options for configuring the TimestampClient.
type TimestampClientOptions struct {
	HTTPClient *http.Client
}

// TimestampResult represents the result of a timestamp request, including the timestamp time, serial number, and raw token.
type TimestampResult struct {
	Time time.Time
	SerialNumber string
	RawToken []byte
}

// TimestampOptions defines options for requesting a timestamp, such as the hash algorithm to use and whether to request certificates in the response.
type TimestampOptions struct {
	Hash crypto.Hash
	RequestCertificates bool
}

// NewTimestampClient creates a new TimestampClient with the specified TSA URL.
// If tsaURL is empty, DefaultMicrosoftTSAURL is used.
func NewTimestampClient(tsaURL string, options *TimestampClientOptions) *TimestampClient {
	if tsaURL == "" {
		tsaURL = DefaultMicrosoftTSAURL
	}

	httpClient := &http.Client{Timeout: 10 * time.Second}
	if options != nil && options.HTTPClient != nil {
		httpClient = options.HTTPClient
	}

	return &TimestampClient{
		tsaURL:     tsaURL,
		httpClient: httpClient,
	}
}

// Timestamp requests an RFC 3161 timestamp for the given data.
func (c *TimestampClient) Timestamp(ctx context.Context, data []byte, options *TimestampOptions) (*TimestampResult, error) {
	if len(data) == 0 {
		return nil, fmt.Errorf("data cannot be empty")
	}

	hash := crypto.SHA256
	requestCerts := true

	if options != nil {
		if options.Hash != 0 {
			hash = options.Hash
		}

		if options.RequestCertificates == false {
			requestCerts = false
		}
	}

	tsReq, err := timestamp.CreateRequest(bytes.NewReader(data), &timestamp.RequestOptions{
		Hash:         hash,
		Certificates: requestCerts,
	})

	if err != nil {
		return nil, fmt.Errorf("failed to create timestamp request: %w", err)
	}

	httpReq, err := http.NewRequestWithContext(ctx, "POST", c.tsaURL, bytes.NewReader(tsReq))

	if err != nil {
		return nil, fmt.Errorf("failed to create HTTP request: %w", err)
	}

	httpReq.Header.Set("Content-Type", "application/timestamp-query")

	resp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("failed to send timestamp request: %w", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("timestamp request failed with status: %s", resp.Status)
	}

	respBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read timestamp response: %w", err)
	}

	tsResp, err := timestamp.ParseResponse(respBytes)
	if err != nil {
		return nil, fmt.Errorf("failed to parse timestamp response: %w", err)
	}

	return &TimestampResult{
		Time:         tsResp.Time,
		SerialNumber: tsResp.SerialNumber.String(),
		RawToken:     respBytes,
	}, nil
}
