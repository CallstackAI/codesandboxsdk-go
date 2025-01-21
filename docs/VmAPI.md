# \VmAPI

All URIs are relative to *https://api.codesandbox.stream*

Method | HTTP request | Description
------------- | ------------- | -------------
[**VmHibernate**](VmAPI.md#VmHibernate) | **Post** /vm/{id}/hibernate | Hibernate a VM
[**VmShutdown**](VmAPI.md#VmShutdown) | **Post** /vm/{id}/shutdown | Shutdown a VM
[**VmStart**](VmAPI.md#VmStart) | **Post** /vm/{id}/start | Start a VM
[**VmUpdateHibernationTimeout**](VmAPI.md#VmUpdateHibernationTimeout) | **Put** /vm/{id}/hibernation_timeout | Update VM Hibernation Timeout
[**VmUpdateSpecs**](VmAPI.md#VmUpdateSpecs) | **Put** /vm/{id}/specs | Update VM Specs
[**VmUpdateSpecs2**](VmAPI.md#VmUpdateSpecs2) | **Post** /vm/{id}/update_specs | Update VM Specs



## VmHibernate

> VMHibernateResponse VmHibernate(ctx, id).Body(body).Execute()

Hibernate a VM



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
	id := "new" // string | Sandbox ID
	body := interface{}(987) // interface{} | VM Hibernate Request (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VmAPI.VmHibernate(context.Background(), id).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VmAPI.VmHibernate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VmHibernate`: VMHibernateResponse
	fmt.Fprintf(os.Stdout, "Response from `VmAPI.VmHibernate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Sandbox ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiVmHibernateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **interface{}** | VM Hibernate Request | 

### Return type

[**VMHibernateResponse**](VMHibernateResponse.md)

### Authorization

[authorization](../README.md#authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VmShutdown

> VMShutdownResponse VmShutdown(ctx, id).Body(body).Execute()

Shutdown a VM



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
	id := "new" // string | Sandbox ID
	body := interface{}(987) // interface{} | VM Shutdown Request (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VmAPI.VmShutdown(context.Background(), id).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VmAPI.VmShutdown``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VmShutdown`: VMShutdownResponse
	fmt.Fprintf(os.Stdout, "Response from `VmAPI.VmShutdown`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Sandbox ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiVmShutdownRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **interface{}** | VM Shutdown Request | 

### Return type

[**VMShutdownResponse**](VMShutdownResponse.md)

### Authorization

[authorization](../README.md#authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VmStart

> VMStartResponse VmStart(ctx, id).VMStartRequest(vMStartRequest).Execute()

Start a VM



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
	id := "new" // string | Sandbox ID
	vMStartRequest := *openapiclient.NewVMStartRequest() // VMStartRequest | VM Start Request (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VmAPI.VmStart(context.Background(), id).VMStartRequest(vMStartRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VmAPI.VmStart``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VmStart`: VMStartResponse
	fmt.Fprintf(os.Stdout, "Response from `VmAPI.VmStart`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Sandbox ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiVmStartRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **vMStartRequest** | [**VMStartRequest**](VMStartRequest.md) | VM Start Request | 

### Return type

[**VMStartResponse**](VMStartResponse.md)

### Authorization

[authorization](../README.md#authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VmUpdateHibernationTimeout

> VMUpdateHibernationTimeoutResponse VmUpdateHibernationTimeout(ctx, id).VMUpdateHibernationTimeoutRequest(vMUpdateHibernationTimeoutRequest).Execute()

Update VM Hibernation Timeout



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
	id := "new" // string | Sandbox ID
	vMUpdateHibernationTimeoutRequest := *openapiclient.NewVMUpdateHibernationTimeoutRequest(int32(300)) // VMUpdateHibernationTimeoutRequest | VM Update Hibernation Timeout Request (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VmAPI.VmUpdateHibernationTimeout(context.Background(), id).VMUpdateHibernationTimeoutRequest(vMUpdateHibernationTimeoutRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VmAPI.VmUpdateHibernationTimeout``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VmUpdateHibernationTimeout`: VMUpdateHibernationTimeoutResponse
	fmt.Fprintf(os.Stdout, "Response from `VmAPI.VmUpdateHibernationTimeout`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Sandbox ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiVmUpdateHibernationTimeoutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **vMUpdateHibernationTimeoutRequest** | [**VMUpdateHibernationTimeoutRequest**](VMUpdateHibernationTimeoutRequest.md) | VM Update Hibernation Timeout Request | 

### Return type

[**VMUpdateHibernationTimeoutResponse**](VMUpdateHibernationTimeoutResponse.md)

### Authorization

[authorization](../README.md#authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VmUpdateSpecs

> VMUpdateSpecsResponse VmUpdateSpecs(ctx, id).VMUpdateSpecsRequest(vMUpdateSpecsRequest).Execute()

Update VM Specs



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
	id := "new" // string | Sandbox ID
	vMUpdateSpecsRequest := *openapiclient.NewVMUpdateSpecsRequest("Micro") // VMUpdateSpecsRequest | VM Update Specs Request (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VmAPI.VmUpdateSpecs(context.Background(), id).VMUpdateSpecsRequest(vMUpdateSpecsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VmAPI.VmUpdateSpecs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VmUpdateSpecs`: VMUpdateSpecsResponse
	fmt.Fprintf(os.Stdout, "Response from `VmAPI.VmUpdateSpecs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Sandbox ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiVmUpdateSpecsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **vMUpdateSpecsRequest** | [**VMUpdateSpecsRequest**](VMUpdateSpecsRequest.md) | VM Update Specs Request | 

### Return type

[**VMUpdateSpecsResponse**](VMUpdateSpecsResponse.md)

### Authorization

[authorization](../README.md#authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VmUpdateSpecs2

> VMUpdateSpecsResponse VmUpdateSpecs2(ctx, id).VMUpdateSpecsRequest(vMUpdateSpecsRequest).Execute()

Update VM Specs



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
	id := "new" // string | Sandbox ID
	vMUpdateSpecsRequest := *openapiclient.NewVMUpdateSpecsRequest("Micro") // VMUpdateSpecsRequest | VM Update Specs Request (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VmAPI.VmUpdateSpecs2(context.Background(), id).VMUpdateSpecsRequest(vMUpdateSpecsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VmAPI.VmUpdateSpecs2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VmUpdateSpecs2`: VMUpdateSpecsResponse
	fmt.Fprintf(os.Stdout, "Response from `VmAPI.VmUpdateSpecs2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Sandbox ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiVmUpdateSpecs2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **vMUpdateSpecsRequest** | [**VMUpdateSpecsRequest**](VMUpdateSpecsRequest.md) | VM Update Specs Request | 

### Return type

[**VMUpdateSpecsResponse**](VMUpdateSpecsResponse.md)

### Authorization

[authorization](../README.md#authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

