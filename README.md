# CodeSandboxSDK Go

This repository contains Codesandbox SDK for Go.

SDK is generated from the OpenAPI specification of the Codesandbox API.

## Usage

```go
package main

import (
	"context"
	"fmt"

	"github.com/callstackai/codesandboxsdk-go"
)

func main() {
	client, err := codesandboxsdk.NewSDKClient(codesandboxsdk.APIBaseURL, "<api_token>")
	if err != nil {
		panic(err)
	}
	ctx := context.TODO()

	sandbox, err := client.SandboxgetWithResponse(ctx, "<sandbox_id>")
	if err != nil {
		panic(err)
	}
	// SDK generates response type with Data as []bytes for manual unmarshalling and JSON200 property as a struct with the response data 
	data := sandbox.JSON200.Data
	fmt.Println(data.Id)
}
```

## Generating the SDK
```shell
go install github.com/deepmap/oapi-codegen/v2/cmd/oapi-codegen@v2.2.0
oapi-codegen -config cfg.yaml openapi-time.yaml
```

## OpenAPI Specification

Some modifications to specification were made to make it compatible with the generator.

- Custom Time type is used to parse time without timezone.
- Items arrays which include oneOf are replaced with and object.
- Sandbox and VMStartData definitions were created and replaced in the definition where they are used
to create types that can be passed around when using the SDK.