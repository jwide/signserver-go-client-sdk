# \WorkersAPI

All URIs are relative to */signserver/rest/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateWorker**](WorkersAPI.md#CreateWorker) | **Post** /workers | Create a new worker given a list of properties
[**CreateWorkerWithId**](WorkersAPI.md#CreateWorkerWithId) | **Post** /workers/{id} | Submit data for adding a new worker from multiple properties
[**DeleteWorkerWithId**](WorkersAPI.md#DeleteWorkerWithId) | **Delete** /workers/{id} | Removing worker
[**PatchWorkerWithId**](WorkersAPI.md#PatchWorkerWithId) | **Patch** /workers/{id} | Submit data for update and delete worker properties
[**Sign**](WorkersAPI.md#Sign) | **Post** /workers/{workerName}/process | Submit data for processing
[**UpdateWorkerWithId**](WorkersAPI.md#UpdateWorkerWithId) | **Put** /workers/{id} | Submit data for replace worker properties with the new properties



## CreateWorker

> CreateWorker(ctx).WorkerRequest(workerRequest).Execute()

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
    resp, r, err := apiClient.WorkersAPI.CreateWorker(context.Background()).WorkerRequest(workerRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkersAPI.CreateWorker``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

 |

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWorkerRequest struct via the builder pattern


| Name | Type | Description  | Notes |
| ------------- | ------------- | ------------- | ------------- |
| **workerRequest** | [**WorkerRequest**](WorkerRequest.md) | Properties of the worker to be created |  |

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


## CreateWorkerWithId

> CreateWorkerWithId(ctx, id).WorkerRequest(workerRequest).Execute()

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
    resp, r, err := apiClient.WorkersAPI.CreateWorkerWithId(context.Background(), id).WorkerRequest(workerRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkersAPI.CreateWorkerWithId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


| Name | Type | Description  | Notes |
| ------------- | ------------- | ------------- | ------------- |
| **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. | 
| **id** | **int32** |  |  |

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWorkerWithIdRequest struct via the builder pattern


| Name | Type | Description  | Notes |
| ------------- | ------------- | ------------- | ------------- |
| **workerRequest** | [**WorkerRequest**](WorkerRequest.md) | The request |  |

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


## DeleteWorkerWithId

> DeleteWorkerWithId(ctx, id).Execute()

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
    resp, r, err := apiClient.WorkersAPI.DeleteWorkerWithId(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkersAPI.DeleteWorkerWithId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


| Name | Type | Description  | Notes |
| ------------- | ------------- | ------------- | ------------- |
| **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. | 
| **id** | **int32** |  |  |

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteWorkerWithIdRequest struct via the builder pattern


| Name | Type | Description  | Notes |
| ------------- | ------------- | ------------- | ------------- |
 |

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


## PatchWorkerWithId

> PatchWorkerWithId(ctx, id).WorkerRequest(workerRequest).Execute()

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
    resp, r, err := apiClient.WorkersAPI.PatchWorkerWithId(context.Background(), id).WorkerRequest(workerRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkersAPI.PatchWorkerWithId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


| Name | Type | Description  | Notes |
| ------------- | ------------- | ------------- | ------------- |
| **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. | 
| **id** | **int32** |  |  |

### Other Parameters

Other parameters are passed through a pointer to a apiPatchWorkerWithIdRequest struct via the builder pattern


| Name | Type | Description  | Notes |
| ------------- | ------------- | ------------- | ------------- |
| **workerRequest** | [**WorkerRequest**](WorkerRequest.md) | The request |  |

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


## Sign

> ProcessResponse Sign(ctx, workerName).ProcessRequest(processRequest).Execute()

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
    workerName := "TimeStampSigner" // string | The name of the worker to use for signing
    processRequest := *openapiclient.NewProcessRequest("Data_example") // ProcessRequest | The request (optional)

    configuration := signserver.NewConfiguration()
    apiClient := signserver.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkersAPI.Sign(context.Background(), workerName).ProcessRequest(processRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkersAPI.Sign``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Sign`: ProcessResponse
    fmt.Fprintf(os.Stdout, "Response from `WorkersAPI.Sign`: %v\n", resp)
}
```

### Path Parameters


| Name | Type | Description  | Notes |
| ------------- | ------------- | ------------- | ------------- |
| **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. | 
| **workerName** | **string** | The name of the worker to use for signing |  |

### Other Parameters

Other parameters are passed through a pointer to a apiSignRequest struct via the builder pattern


| Name | Type | Description  | Notes |
| ------------- | ------------- | ------------- | ------------- |
| **processRequest** | [**ProcessRequest**](ProcessRequest.md) | The request |  |

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


## UpdateWorkerWithId

> UpdateWorkerWithId(ctx, id).WorkerRequest(workerRequest).Execute()

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
    resp, r, err := apiClient.WorkersAPI.UpdateWorkerWithId(context.Background(), id).WorkerRequest(workerRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkersAPI.UpdateWorkerWithId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


| Name | Type | Description  | Notes |
| ------------- | ------------- | ------------- | ------------- |
| **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. | 
| **id** | **int32** |  |  |

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWorkerWithIdRequest struct via the builder pattern


| Name | Type | Description  | Notes |
| ------------- | ------------- | ------------- | ------------- |
| **workerRequest** | [**WorkerRequest**](WorkerRequest.md) | The request |  |

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

