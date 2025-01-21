# \SandboxAPI

All URIs are relative to *https://api.codesandbox.stream*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SandboxCreate**](SandboxAPI.md#SandboxCreate) | **Post** /sandbox | Create a Sandbox
[**SandboxFork**](SandboxAPI.md#SandboxFork) | **Post** /sandbox/{id}/fork | Fork a Sandbox
[**SandboxGet**](SandboxAPI.md#SandboxGet) | **Get** /sandbox/{id} | Get a Sandbox
[**SandboxList**](SandboxAPI.md#SandboxList) | **Get** /sandbox | List Sandboxes



## SandboxCreate

> SandboxCreateResponse SandboxCreate(ctx).SandboxCreateRequest(sandboxCreateRequest).Execute()

Create a Sandbox



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
	sandboxCreateRequest := *openapiclient.NewSandboxCreateRequest(map[string]SandboxCreateRequestFilesValue{"key": *openapiclient.NewSandboxCreateRequestFilesValue()}) // SandboxCreateRequest | Sandbox Create Request (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SandboxAPI.SandboxCreate(context.Background()).SandboxCreateRequest(sandboxCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SandboxAPI.SandboxCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SandboxCreate`: SandboxCreateResponse
	fmt.Fprintf(os.Stdout, "Response from `SandboxAPI.SandboxCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSandboxCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sandboxCreateRequest** | [**SandboxCreateRequest**](SandboxCreateRequest.md) | Sandbox Create Request | 

### Return type

[**SandboxCreateResponse**](SandboxCreateResponse.md)

### Authorization

[authorization](../README.md#authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SandboxFork

> SandboxForkResponse SandboxFork(ctx, id).SandboxForkRequest(sandboxForkRequest).Execute()

Fork a Sandbox



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
	id := "new" // string | Short ID of the sandbox to fork
	sandboxForkRequest := *openapiclient.NewSandboxForkRequest() // SandboxForkRequest | Sandbox Fork Request (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SandboxAPI.SandboxFork(context.Background(), id).SandboxForkRequest(sandboxForkRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SandboxAPI.SandboxFork``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SandboxFork`: SandboxForkResponse
	fmt.Fprintf(os.Stdout, "Response from `SandboxAPI.SandboxFork`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Short ID of the sandbox to fork | 

### Other Parameters

Other parameters are passed through a pointer to a apiSandboxForkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sandboxForkRequest** | [**SandboxForkRequest**](SandboxForkRequest.md) | Sandbox Fork Request | 

### Return type

[**SandboxForkResponse**](SandboxForkResponse.md)

### Authorization

[authorization](../README.md#authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SandboxGet

> SandboxGetResponse SandboxGet(ctx, id).Execute()

Get a Sandbox



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
	id := "new" // string | Short ID of the sandbox to retrieve

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SandboxAPI.SandboxGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SandboxAPI.SandboxGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SandboxGet`: SandboxGetResponse
	fmt.Fprintf(os.Stdout, "Response from `SandboxAPI.SandboxGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Short ID of the sandbox to retrieve | 

### Other Parameters

Other parameters are passed through a pointer to a apiSandboxGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SandboxGetResponse**](SandboxGetResponse.md)

### Authorization

[authorization](../README.md#authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SandboxList

> SandboxListResponse SandboxList(ctx).Tags(tags).OrderBy(orderBy).Direction(direction).PageSize(pageSize).Page(page).Execute()

List Sandboxes



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
	tags := "sha-123,feature-x" // string | Comma-separated list of tags to filter by (optional)
	orderBy := "updated_at" // string | Field to order results by (optional) (default to "updated_at")
	direction := "desc" // string | Order direction (optional) (default to "desc")
	pageSize := int32(56) // int32 | Maximum number of sandboxes to return in a single response (optional) (default to 20)
	page := int32(56) // int32 | Page number to return (optional) (default to 1)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SandboxAPI.SandboxList(context.Background()).Tags(tags).OrderBy(orderBy).Direction(direction).PageSize(pageSize).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SandboxAPI.SandboxList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SandboxList`: SandboxListResponse
	fmt.Fprintf(os.Stdout, "Response from `SandboxAPI.SandboxList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSandboxListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tags** | **string** | Comma-separated list of tags to filter by | 
 **orderBy** | **string** | Field to order results by | [default to &quot;updated_at&quot;]
 **direction** | **string** | Order direction | [default to &quot;desc&quot;]
 **pageSize** | **int32** | Maximum number of sandboxes to return in a single response | [default to 20]
 **page** | **int32** | Page number to return | [default to 1]

### Return type

[**SandboxListResponse**](SandboxListResponse.md)

### Authorization

[authorization](../README.md#authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

