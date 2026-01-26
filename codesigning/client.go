package codesigning

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// CertificateProfileClientOptions contains the optional parameters for NewCertificateProfileClient.
type CertificateProfileClientOptions struct {
	azcore.ClientOptions
}

// CertificateProfileClient contains the methods for the CertificateProfile group.
// Don't use this type directly, use a constructor function instead.
type CertificateProfileClient struct {
	internal *azcore.Client
	endpoint string
}

// NewCertificateProfileClient creates a new instance of CertificateProfileClient with the specified values.
//   - endpoint - The Azure Artifact Signing endpoint (e.g., "https://scus.codesigning.azure.net").
//   - credential - Used to authorize requests. Usually a credential from azidentity.
//   - options - Optional client settings; pass nil to accept defaults.
func NewCertificateProfileClient(endpoint string, credential azcore.TokenCredential, options *CertificateProfileClientOptions) (*CertificateProfileClient, error) {
	if options == nil {
		options = &CertificateProfileClientOptions{}
	}
	azcoreClient, err := azcore.NewClient(moduleName, moduleVersion, runtime.PipelineOptions{
		PerRetry: []policy.Policy{
			runtime.NewBearerTokenPolicy(credential, []string{"https://codesigning.azure.net/.default"}, nil),
		},
	}, &options.ClientOptions)
	if err != nil {
		return nil, err
	}
	return &CertificateProfileClient{
		internal: azcoreClient,
		endpoint: endpoint,
	}, nil
}

// GetSignCertificateChain - Gets the certificate chain for that account and profile.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-06-15-preview
//   - codeSigningAccountName - Artifact Signing account name.
//   - certificateProfileName - Certificate profile name.
//   - options - CertificateProfileClientGetSignCertificateChainOptions contains the optional parameters for the CertificateProfileClient.GetSignCertificateChain
//     method.
func (client *CertificateProfileClient) GetSignCertificateChain(ctx context.Context, codeSigningAccountName string, certificateProfileName string, options *CertificateProfileClientGetSignCertificateChainOptions) (CertificateProfileClientGetSignCertificateChainResponse, error) {
	var err error
	req, err := client.getSignCertificateChainCreateRequest(ctx, codeSigningAccountName, certificateProfileName, options)
	if err != nil {
		return CertificateProfileClientGetSignCertificateChainResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return CertificateProfileClientGetSignCertificateChainResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return CertificateProfileClientGetSignCertificateChainResponse{}, err
	}
	return CertificateProfileClientGetSignCertificateChainResponse{Body: httpResp.Body}, nil
}

// getSignCertificateChainCreateRequest creates the GetSignCertificateChain request.
func (client *CertificateProfileClient) getSignCertificateChainCreateRequest(ctx context.Context, codeSigningAccountName string, certificateProfileName string, _ *CertificateProfileClientGetSignCertificateChainOptions) (*policy.Request, error) {
	urlPath := "/codesigningaccounts/{codeSigningAccountName}/certificateprofiles/{certificateProfileName}/sign/certchain"
	if codeSigningAccountName == "" {
		return nil, errors.New("parameter codeSigningAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{codeSigningAccountName}", url.PathEscape(codeSigningAccountName))
	if certificateProfileName == "" {
		return nil, errors.New("parameter certificateProfileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{certificateProfileName}", url.PathEscape(certificateProfileName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-06-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	runtime.SkipBodyDownload(req)
	req.Raw().Header["Accept"] = []string{"application/pkcs7-mime, application/x-x509-ca-cert, application/json"}
	return req, nil
}

// GetSignCertificateRoot - Gets the root certificate for that account and profile.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-06-15-preview
//   - codeSigningAccountName - Artifact Signing account name.
//   - certificateProfileName - Certificate profile name.
//   - options - CertificateProfileClientGetSignCertificateRootOptions contains the optional parameters for the CertificateProfileClient.GetSignCertificateRoot
//     method.
func (client *CertificateProfileClient) GetSignCertificateRoot(ctx context.Context, codeSigningAccountName string, certificateProfileName string, options *CertificateProfileClientGetSignCertificateRootOptions) (CertificateProfileClientGetSignCertificateRootResponse, error) {
	var err error
	req, err := client.getSignCertificateRootCreateRequest(ctx, codeSigningAccountName, certificateProfileName, options)
	if err != nil {
		return CertificateProfileClientGetSignCertificateRootResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return CertificateProfileClientGetSignCertificateRootResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return CertificateProfileClientGetSignCertificateRootResponse{}, err
	}
	return CertificateProfileClientGetSignCertificateRootResponse{Body: httpResp.Body}, nil
}

// getSignCertificateRootCreateRequest creates the GetSignCertificateRoot request.
func (client *CertificateProfileClient) getSignCertificateRootCreateRequest(ctx context.Context, codeSigningAccountName string, certificateProfileName string, _ *CertificateProfileClientGetSignCertificateRootOptions) (*policy.Request, error) {
	urlPath := "/codesigningaccounts/{codeSigningAccountName}/certificateprofiles/{certificateProfileName}/sign/rootcert"
	if codeSigningAccountName == "" {
		return nil, errors.New("parameter codeSigningAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{codeSigningAccountName}", url.PathEscape(codeSigningAccountName))
	if certificateProfileName == "" {
		return nil, errors.New("parameter certificateProfileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{certificateProfileName}", url.PathEscape(certificateProfileName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-06-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	runtime.SkipBodyDownload(req)
	req.Raw().Header["Accept"] = []string{"application/x-x509-ca-cert, application/json"}
	return req, nil
}

// GetSignEku - Gets the ekus defined for that account and profile.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-06-15-preview
//   - codeSigningAccountName - Artifact Signing account name.
//   - certificateProfileName - Certificate profile name.
//   - options - CertificateProfileClientGetSignEkuOptions contains the optional parameters for the CertificateProfileClient.GetSignEku
//     method.
func (client *CertificateProfileClient) GetSignEku(ctx context.Context, codeSigningAccountName string, certificateProfileName string, options *CertificateProfileClientGetSignEkuOptions) (CertificateProfileClientGetSignEkuResponse, error) {
	var err error
	req, err := client.getSignEkuCreateRequest(ctx, codeSigningAccountName, certificateProfileName, options)
	if err != nil {
		return CertificateProfileClientGetSignEkuResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return CertificateProfileClientGetSignEkuResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return CertificateProfileClientGetSignEkuResponse{}, err
	}
	resp, err := client.getSignEkuHandleResponse(httpResp)
	return resp, err
}

// getSignEkuCreateRequest creates the GetSignEku request.
func (client *CertificateProfileClient) getSignEkuCreateRequest(ctx context.Context, codeSigningAccountName string, certificateProfileName string, _ *CertificateProfileClientGetSignEkuOptions) (*policy.Request, error) {
	urlPath := "/codesigningaccounts/{codeSigningAccountName}/certificateprofiles/{certificateProfileName}/sign/eku"
	if codeSigningAccountName == "" {
		return nil, errors.New("parameter codeSigningAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{codeSigningAccountName}", url.PathEscape(codeSigningAccountName))
	if certificateProfileName == "" {
		return nil, errors.New("parameter certificateProfileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{certificateProfileName}", url.PathEscape(certificateProfileName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-06-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getSignEkuHandleResponse handles the GetSignEku response.
func (client *CertificateProfileClient) getSignEkuHandleResponse(resp *http.Response) (CertificateProfileClientGetSignEkuResponse, error) {
	result := CertificateProfileClientGetSignEkuResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.StringArray); err != nil {
		return CertificateProfileClientGetSignEkuResponse{}, err
	}
	return result, nil
}

// GetSignOperation - Gets the status of a sign operation.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-06-15-preview
//   - codeSigningAccountName - Artifact Signing account name.
//   - certificateProfileName - Certificate profile name.
//   - operationID - ID of the operation.
//   - options - CertificateProfileClientGetSignOperationOptions contains the optional parameters for the CertificateProfileClient.GetSignOperation
//     method.
func (client *CertificateProfileClient) GetSignOperation(ctx context.Context, codeSigningAccountName string, certificateProfileName string, operationID string, options *CertificateProfileClientGetSignOperationOptions) (CertificateProfileClientGetSignOperationResponse, error) {
	var err error
	req, err := client.getSignOperationCreateRequest(ctx, codeSigningAccountName, certificateProfileName, operationID, options)
	if err != nil {
		return CertificateProfileClientGetSignOperationResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return CertificateProfileClientGetSignOperationResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return CertificateProfileClientGetSignOperationResponse{}, err
	}
	resp, err := client.getSignOperationHandleResponse(httpResp)
	return resp, err
}

// getSignOperationCreateRequest creates the GetSignOperation request.
func (client *CertificateProfileClient) getSignOperationCreateRequest(ctx context.Context, codeSigningAccountName string, certificateProfileName string, operationID string, _ *CertificateProfileClientGetSignOperationOptions) (*policy.Request, error) {
	urlPath := "/codesigningaccounts/{codeSigningAccountName}/certificateprofiles/{certificateProfileName}/sign/{operationId}"
	if codeSigningAccountName == "" {
		return nil, errors.New("parameter codeSigningAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{codeSigningAccountName}", url.PathEscape(codeSigningAccountName))
	if certificateProfileName == "" {
		return nil, errors.New("parameter certificateProfileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{certificateProfileName}", url.PathEscape(certificateProfileName))
	if operationID == "" {
		return nil, errors.New("parameter operationID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{operationId}", url.PathEscape(operationID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-06-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getSignOperationHandleResponse handles the GetSignOperation response.
func (client *CertificateProfileClient) getSignOperationHandleResponse(resp *http.Response) (CertificateProfileClientGetSignOperationResponse, error) {
	result := CertificateProfileClientGetSignOperationResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SignStatus); err != nil {
		return CertificateProfileClientGetSignOperationResponse{}, err
	}
	return result, nil
}

// BeginSign - Initiates a sign operation.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-06-15-preview
//   - codeSigningAccountName - Artifact Signing account name.
//   - certificateProfileName - Certificate profile name.
//   - body - Sign request details.
//   - options - CertificateProfileClientBeginSignOptions contains the optional parameters for the CertificateProfileClient.BeginSign
//     method.
func (client *CertificateProfileClient) BeginSign(ctx context.Context, codeSigningAccountName string, certificateProfileName string, body SignRequest, options *CertificateProfileClientBeginSignOptions) (*runtime.Poller[CertificateProfileClientSignResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.sign(ctx, codeSigningAccountName, certificateProfileName, body, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[CertificateProfileClientSignResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[CertificateProfileClientSignResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Sign - Initiates a sign operation.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-06-15-preview
func (client *CertificateProfileClient) sign(ctx context.Context, codeSigningAccountName string, certificateProfileName string, body SignRequest, options *CertificateProfileClientBeginSignOptions) (*http.Response, error) {
	var err error
	req, err := client.signCreateRequest(ctx, codeSigningAccountName, certificateProfileName, body, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// signCreateRequest creates the Sign request.
func (client *CertificateProfileClient) signCreateRequest(ctx context.Context, codeSigningAccountName string, certificateProfileName string, body SignRequest, options *CertificateProfileClientBeginSignOptions) (*policy.Request, error) {
	urlPath := "/codesigningaccounts/{codeSigningAccountName}/certificateprofiles/{certificateProfileName}/sign"
	if codeSigningAccountName == "" {
		return nil, errors.New("parameter codeSigningAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{codeSigningAccountName}", url.PathEscape(codeSigningAccountName))
	if certificateProfileName == "" {
		return nil, errors.New("parameter certificateProfileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{certificateProfileName}", url.PathEscape(certificateProfileName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-06-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if options != nil && options.CertificateThumbprint != nil {
		req.Raw().Header["certificate-thumbprint"] = []string{*options.CertificateThumbprint}
	}
	if options != nil && options.ClientVersion != nil {
		req.Raw().Header["client-version"] = []string{*options.ClientVersion}
	}
	if options != nil && options.XCorrelationID != nil {
		req.Raw().Header["x-correlation-id"] = []string{*options.XCorrelationID}
	}
	if err := runtime.MarshalAsJSON(req, body); err != nil {
	return nil, err
}
;	return req, nil
}

