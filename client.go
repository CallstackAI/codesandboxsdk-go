package codesandboxsdk

import (
	"fmt"
	"net/http"

	"github.com/oapi-codegen/oapi-codegen/v2/pkg/securityprovider"
)

const (
	APIBaseURL = "https://api.codesandbox.io"
)

type APIClient struct {
	*ClientWithResponses
}

func NewSDKClient(baseURL, token string) (*APIClient, error) {
	httpClient := &http.Client{}
	bearerAuth, err := securityprovider.NewSecurityProviderBearerToken(token)
	if err != nil {
		return nil, fmt.Errorf("failed to create auth provider: %w", err)
	}
	client, err := NewClientWithResponses(
		baseURL,
		WithHTTPClient(httpClient),
		WithBaseURL(baseURL),
		WithRequestEditorFn(bearerAuth.Intercept),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create codesandbox SDK client: %w", err)
	}
	return &APIClient{client}, nil
}
