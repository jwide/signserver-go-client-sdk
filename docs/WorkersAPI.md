# \WorkersAPI

All URIs are relative to */signserver/rest/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**WorkersIdDelete**](WorkersAPI.md#WorkersIdDelete) | **Delete** /workers/{id} | Removing worker
[**WorkersIdPatch**](WorkersAPI.md#WorkersIdPatch) | **Patch** /workers/{id} | Submit data for update and delete worker properties
[**WorkersIdPost**](WorkersAPI.md#WorkersIdPost) | **Post** /workers/{id} | Submit data for adding a new worker from multiple properties
[**WorkersIdPut**](WorkersAPI.md#WorkersIdPut) | **Put** /workers/{id} | Submit data for replace worker properties with the new properties
[**WorkersNameProcessPost**](WorkersAPI.md#WorkersNameProcessPost) | **Post** /workers/{name}/process | Submit data for processing
[**WorkersPost**](WorkersAPI.md#WorkersPost) | **Post** /workers | Create a new worker given a list of properties



## WorkersIdDelete

> WorkersIdDelete(ctx, id).Execute()

Removing worker



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "github.com/Keyfactor/signserver-go-client-sdk/api/signserver"
)

func main() {
    id := int32(56) // int32 | 

    configuration := signserver.NewConfiguration()
    apiClient := signserver.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkersAPI.WorkersIdDelete(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkersAPI.WorkersIdDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiWorkersIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WorkersIdPatch

> WorkersIdPatch(ctx, id).WorkerRequest(workerRequest).Execute()

Submit data for update and delete worker properties



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "github.com/Keyfactor/signserver-go-client-sdk/api/signserver"
)

func main() {
    id := int32(56) // int32 | 
    workerRequest := *openapiclient.NewWorkerRequest() // WorkerRequest | The request (optional)

    configuration := signserver.NewConfiguration()
    apiClient := signserver.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkersAPI.WorkersIdPatch(context.Background(), id).WorkerRequest(workerRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkersAPI.WorkersIdPatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiWorkersIdPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **workerRequest** | [**WorkerRequest**](WorkerRequest.md) | The request | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WorkersIdPost

> WorkersIdPost(ctx, id).WorkerRequest(workerRequest).Execute()

Submit data for adding a new worker from multiple properties



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "github.com/Keyfactor/signserver-go-client-sdk/api/signserver"
)

func main() {
    id := int32(56) // int32 | 
    workerRequest := *openapiclient.NewWorkerRequest() // WorkerRequest | The request (optional)

    configuration := signserver.NewConfiguration()
    apiClient := signserver.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkersAPI.WorkersIdPost(context.Background(), id).WorkerRequest(workerRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkersAPI.WorkersIdPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiWorkersIdPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **workerRequest** | [**WorkerRequest**](WorkerRequest.md) | The request | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WorkersIdPut

> WorkersIdPut(ctx, id).WorkerRequest(workerRequest).Execute()

Submit data for replace worker properties with the new properties



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "github.com/Keyfactor/signserver-go-client-sdk/api/signserver"
)

func main() {
    id := int32(56) // int32 | 
    workerRequest := *openapiclient.NewWorkerRequest() // WorkerRequest | The request (optional)

    configuration := signserver.NewConfiguration()
    apiClient := signserver.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkersAPI.WorkersIdPut(context.Background(), id).WorkerRequest(workerRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkersAPI.WorkersIdPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiWorkersIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **workerRequest** | [**WorkerRequest**](WorkerRequest.md) | The request | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WorkersNameProcessPost

> ProcessResponse WorkersNameProcessPost(ctx, name).ProcessRequest(processRequest).Execute()

Submit data for processing



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "github.com/Keyfactor/signserver-go-client-sdk/api/signserver"
)

func main() {
    name := "TimeStampSigner" // string | The name of the worker to use for signing
    processRequest := *openapiclient.NewProcessRequest("Data_example") // ProcessRequest | The request (optional)

    configuration := signserver.NewConfiguration()
    apiClient := signserver.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkersAPI.WorkersNameProcessPost(context.Background(), name).ProcessRequest(processRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkersAPI.WorkersNameProcessPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WorkersNameProcessPost`: ProcessResponse
    fmt.Fprintf(os.Stdout, "Response from `WorkersAPI.WorkersNameProcessPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | The name of the worker to use for signing | 

### Other Parameters

Other parameters are passed through a pointer to a apiWorkersNameProcessPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **processRequest** | [**ProcessRequest**](ProcessRequest.md) | The request | 

### Return type

[**ProcessResponse**](ProcessResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WorkersPost

> WorkersPost(ctx).WorkerRequest(workerRequest).Execute()

Create a new worker given a list of properties



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "github.com/Keyfactor/signserver-go-client-sdk/api/signserver"
)

func main() {
    workerRequest := *openapiclient.NewWorkerRequest() // WorkerRequest | Properties of the worker to be created (optional)

    configuration := signserver.NewConfiguration()
    apiClient := signserver.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkersAPI.WorkersPost(context.Background()).WorkerRequest(workerRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkersAPI.WorkersPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWorkersPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **workerRequest** | [**WorkerRequest**](WorkerRequest.md) | Properties of the worker to be created | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

