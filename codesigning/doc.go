// Package codesigning provides a client for Artifact Signing.
//
// The CertificateProfileClient allows you to sign artifacts using Artifact Signing,
// retrieve certificate chains, and manage signing operations.
//
// # Getting Started
//
// To create a client, use the NewCertificateProfileClient function with your
// Azure credentials and endpoint:
//
//	cred, err := azidentity.NewDefaultAzureCredential(nil)
//	if err != nil {
//		// handle error
//	}
//	client, err := codesigning.NewCertificateProfileClient("https://your-endpoint.codesigning.azure.net", cred, nil)
//	if err != nil {
//		// handle error
//	}
//
// # Signing Operations
//
// Use BeginSign to initiate an asynchronous signing operation:
//
//	poller, err := client.BeginSign(ctx, accountName, profileName, codesigning.SignRequest{
//		Digest:             digestBytes,
//		SignatureAlgorithm: to.Ptr(codesigning.SignatureAlgorithmRS256),
//	}, nil)
//	if err != nil {
//		// handle error
//	}
//	result, err := poller.PollUntilDone(ctx, nil)
//
// # Certificate Operations
//
// Retrieve certificate chain and root certificates:
//
//	chainResp, err := client.GetSignCertificateChain(ctx, accountName, profileName, nil)
//	rootResp, err := client.GetSignCertificateRoot(ctx, accountName, profileName, nil)
package codesigning
