package codesandboxsdk

import (
	"fmt"
	"net/http"

	"github.com/oapi-codegen/oapi-codegen/v2/pkg/securityprovider"
)

const (
	APIBaseURL = "https://api.codesandbox.io"
)

type SDKClient struct {
	*ClientWithResponses
}

func NewSDKClient(baseURL, token string, opts ...ClientOption) (*SDKClient, error) {
	httpClient := &http.Client{}
	bearerAuth, err := securityprovider.NewSecurityProviderBearerToken(token)
	if err != nil {
		return nil, fmt.Errorf("failed to create auth provider: %w", err)
	}
	defaultOpts := []ClientOption{
		WithHTTPClient(httpClient),
		WithBaseURL(baseURL),
		WithRequestEditorFn(bearerAuth.Intercept),
	}
	client, err := NewClientWithResponses(
		baseURL,
		append(defaultOpts, opts...)...,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create codesandbox SDK client: %w", err)
	}
	return &SDKClient{client}, nil
}
