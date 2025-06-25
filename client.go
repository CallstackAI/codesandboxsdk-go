package codesandboxsdk

import (
	"fmt"
	"net/http"

	"github.com/oapi-codegen/oapi-codegen/v2/pkg/securityprovider"

	csbapi "github.com/callstackai/codesandboxsdk-go/pkg/api"
	csbsetup "github.com/callstackai/codesandboxsdk-go/pkg/setup"
)

const (
	APIBaseURL = "https://api.codesandbox.io"
)

func NewAPIClient(baseURL, token string, opts ...csbapi.ClientOption) (*csbapi.ClientWithResponses, error) {
	httpClient := &http.Client{}
	bearerAuth, err := securityprovider.NewSecurityProviderBearerToken(token)
	if err != nil {
		return nil, fmt.Errorf("failed to create auth provider: %w", err)
	}
	defaultOpts := []csbapi.ClientOption{
		csbapi.WithHTTPClient(httpClient),
		csbapi.WithBaseURL(baseURL),
		csbapi.WithRequestEditorFn(bearerAuth.Intercept),
	}
	return csbapi.NewClientWithResponses(
		baseURL,
		append(defaultOpts, opts...)...,
	)
}

func NewSetupClient(baseURL, token string, opts ...csbsetup.ClientOption) (*csbsetup.ClientWithResponses, error) {
	httpClient := &http.Client{}
	bearerAuth, err := securityprovider.NewSecurityProviderBearerToken(token)
	if err != nil {
		return nil, fmt.Errorf("failed to create auth provider: %w", err)
	}
	defaultOpts := []csbsetup.ClientOption{
		csbsetup.WithHTTPClient(httpClient),
		csbsetup.WithBaseURL(baseURL),
		csbsetup.WithRequestEditorFn(bearerAuth.Intercept),
	}
	return csbsetup.NewClientWithResponses(
		baseURL,
		append(defaultOpts, opts...)...,
	)
}
