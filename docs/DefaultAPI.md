# \DefaultAPI

All URIs are relative to *https://api.codesandbox.stream*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TokenCreate**](DefaultAPI.md#TokenCreate) | **Post** /org/workspace/{team_id}/tokens | Create an API Token
[**TokenUpdate**](DefaultAPI.md#TokenUpdate) | **Patch** /org/workspace/{team_id}/tokens/{token_id} | Update an API Token
[**WorkspaceCreate**](DefaultAPI.md#WorkspaceCreate) | **Post** /org/workspace | Create a Workspace



## TokenCreate

> TokenCreateResponse TokenCreate(ctx, teamId).TokenCreateRequest(tokenCreateRequest).Execute()

Create an API Token



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/callstackai/codesandboxsdk"
)

func main() {
	teamId := "teamId_example" // string | 
	tokenCreateRequest := *openapiclient.NewTokenCreateRequest() // TokenCreateRequest | Token Create Request (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.TokenCreate(context.Background(), teamId).TokenCreateRequest(tokenCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.TokenCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TokenCreate`: TokenCreateResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.TokenCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTokenCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tokenCreateRequest** | [**TokenCreateRequest**](TokenCreateRequest.md) | Token Create Request | 

### Return type

[**TokenCreateResponse**](TokenCreateResponse.md)

### Authorization

[authorization](../README.md#authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TokenUpdate

> TokenUpdateResponse TokenUpdate(ctx, teamId, tokenId).TokenUpdateRequest(tokenUpdateRequest).Execute()

Update an API Token



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/callstackai/codesandboxsdk"
)

func main() {
	teamId := "teamId_example" // string | 
	tokenId := "tokenId_example" // string | 
	tokenUpdateRequest := *openapiclient.NewTokenUpdateRequest() // TokenUpdateRequest | Token Update Request (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.TokenUpdate(context.Background(), teamId, tokenId).TokenUpdateRequest(tokenUpdateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.TokenUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TokenUpdate`: TokenUpdateResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.TokenUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **string** |  | 
**tokenId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTokenUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **tokenUpdateRequest** | [**TokenUpdateRequest**](TokenUpdateRequest.md) | Token Update Request | 

### Return type

[**TokenUpdateResponse**](TokenUpdateResponse.md)

### Authorization

[authorization](../README.md#authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WorkspaceCreate

> WorkspaceCreateResponse WorkspaceCreate(ctx).WorkspaceCreateRequest(workspaceCreateRequest).Execute()

Create a Workspace



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/callstackai/codesandboxsdk"
)

func main() {
	workspaceCreateRequest := *openapiclient.NewWorkspaceCreateRequest("Name_example") // WorkspaceCreateRequest | Workspace Create Request (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.WorkspaceCreate(context.Background()).WorkspaceCreateRequest(workspaceCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.WorkspaceCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WorkspaceCreate`: WorkspaceCreateResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.WorkspaceCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWorkspaceCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **workspaceCreateRequest** | [**WorkspaceCreateRequest**](WorkspaceCreateRequest.md) | Workspace Create Request | 

### Return type

[**WorkspaceCreateResponse**](WorkspaceCreateResponse.md)

### Authorization

[authorization](../README.md#authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

