# \DefaultAPI

All URIs are relative to */signserver/rest/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**WorkersIdDelete**](DefaultAPI.md#WorkersIdDelete) | **Delete** /workers/{id} | Removing worker
[**WorkersIdOrNameProcessPost**](DefaultAPI.md#WorkersIdOrNameProcessPost) | **Post** /workers/{idOrName}/process | Submit data for processing
[**WorkersIdPatch**](DefaultAPI.md#WorkersIdPatch) | **Patch** /workers/{id} | Submit data for update and delete worker properties
[**WorkersIdPost**](DefaultAPI.md#WorkersIdPost) | **Post** /workers/{id} | Submit data for adding a new worker from multiple properties
[**WorkersIdPut**](DefaultAPI.md#WorkersIdPut) | **Put** /workers/{id} | Submit data for replace worker properties with the new properties
[**WorkersPost**](DefaultAPI.md#WorkersPost) | **Post** /workers | Submit data for adding a new worker from multiple properties



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
    openapiclient "./openapi"
)

func main() {
    id := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.WorkersIdDelete(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.WorkersIdDelete``: %v\n", err)
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


## WorkersIdOrNameProcessPost

> ProcessResponse WorkersIdOrNameProcessPost(ctx, idOrName).ProcessRequest(processRequest).Execute()

Submit data for processing



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    idOrName := *openapiclient.NewWorkersIdOrNameProcessPostIdOrNameParameter() // WorkersIdOrNameProcessPostIdOrNameParameter | Worker Id or name of the worker
    processRequest := *openapiclient.NewProcessRequest("Data_example") // ProcessRequest | The request (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.WorkersIdOrNameProcessPost(context.Background(), idOrName).ProcessRequest(processRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.WorkersIdOrNameProcessPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WorkersIdOrNameProcessPost`: ProcessResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.WorkersIdOrNameProcessPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**idOrName** | [**WorkersIdOrNameProcessPostIdOrNameParameter**](.md) | Worker Id or name of the worker | 

### Other Parameters

Other parameters are passed through a pointer to a apiWorkersIdOrNameProcessPostRequest struct via the builder pattern


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
    openapiclient "./openapi"
)

func main() {
    id := int32(56) // int32 | 
    workerRequest := *openapiclient.NewWorkerRequest() // WorkerRequest | The request (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.WorkersIdPatch(context.Background(), id).WorkerRequest(workerRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.WorkersIdPatch``: %v\n", err)
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
    openapiclient "./openapi"
)

func main() {
    id := int32(56) // int32 | 
    workerRequest := *openapiclient.NewWorkerRequest() // WorkerRequest | The request (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.WorkersIdPost(context.Background(), id).WorkerRequest(workerRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.WorkersIdPost``: %v\n", err)
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
    openapiclient "./openapi"
)

func main() {
    id := int32(56) // int32 | 
    workerRequest := *openapiclient.NewWorkerRequest() // WorkerRequest | The request (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.WorkersIdPut(context.Background(), id).WorkerRequest(workerRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.WorkersIdPut``: %v\n", err)
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


## WorkersPost

> WorkersPost(ctx).WorkerRequest(workerRequest).Execute()

Submit data for adding a new worker from multiple properties



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    workerRequest := *openapiclient.NewWorkerRequest() // WorkerRequest | The request (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.WorkersPost(context.Background()).WorkerRequest(workerRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.WorkersPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWorkersPostRequest struct via the builder pattern


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

