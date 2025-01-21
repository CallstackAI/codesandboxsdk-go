# \MetaAPI

All URIs are relative to *https://api.codesandbox.stream*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MetaInfo**](MetaAPI.md#MetaInfo) | **Get** /meta/info | Metadata about the API



## MetaInfo

> MetaInformation MetaInfo(ctx).Execute()

Metadata about the API

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MetaAPI.MetaInfo(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetaAPI.MetaInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MetaInfo`: MetaInformation
	fmt.Fprintf(os.Stdout, "Response from `MetaAPI.MetaInfo`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiMetaInfoRequest struct via the builder pattern


### Return type

[**MetaInformation**](MetaInformation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

