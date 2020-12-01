# \HostAgentApi

All URIs are relative to *https://unit-tenant.instana.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAgentSnapshot**](HostAgentApi.md#GetAgentSnapshot) | **Get** /api/host-agent/{id} | Get host agent snapshot details
[**GetLogs**](HostAgentApi.md#GetLogs) | **Get** /api/host-agent/{hostId}/logs | Agent download logs
[**Search**](HostAgentApi.md#Search) | **Get** /api/host-agent | Query host agent snapshots
[**Update**](HostAgentApi.md#Update) | **Post** /api/host-agent/{hostId}/update | Agent update
[**UpdateConfigurationByHost**](HostAgentApi.md#UpdateConfigurationByHost) | **Post** /api/host-agent/{hostId}/configuration | Update agent configuration by host
[**UpdateConfigurationByQuery**](HostAgentApi.md#UpdateConfigurationByQuery) | **Post** /api/host-agent/configuration | Update agent configuration by query



## GetAgentSnapshot

> SnapshotItem GetAgentSnapshot(ctx, id).To(to).WindowSize(windowSize).Execute()

Get host agent snapshot details

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
    id := "id_example" // string | 
    to := int64(789) // int64 |  (optional)
    windowSize := int64(789) // int64 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.HostAgentApi.GetAgentSnapshot(context.Background(), id).To(to).WindowSize(windowSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HostAgentApi.GetAgentSnapshot``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAgentSnapshot`: SnapshotItem
    fmt.Fprintf(os.Stdout, "Response from `HostAgentApi.GetAgentSnapshot`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAgentSnapshotRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **to** | **int64** |  | 
 **windowSize** | **int64** |  | 

### Return type

[**SnapshotItem**](SnapshotItem.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLogs

> GetLogs(ctx, hostId).File(file).Download(download).Execute()

Agent download logs

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
    hostId := "hostId_example" // string | 
    file := []string{"Inner_example"} // []string | 
    download := true // bool |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.HostAgentApi.GetLogs(context.Background(), hostId).File(file).Download(download).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HostAgentApi.GetLogs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hostId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **file** | **[]string** |  | 
 **download** | **bool** |  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Search

> SnapshotResult Search(ctx).Query(query).To(to).WindowSize(windowSize).Size(size).Offline(offline).Execute()

Query host agent snapshots

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
    query := "query_example" // string |  (optional)
    to := int64(789) // int64 |  (optional)
    windowSize := int64(789) // int64 |  (optional)
    size := int32(56) // int32 |  (optional)
    offline := true // bool |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.HostAgentApi.Search(context.Background()).Query(query).To(to).WindowSize(windowSize).Size(size).Offline(offline).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HostAgentApi.Search``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Search`: SnapshotResult
    fmt.Fprintf(os.Stdout, "Response from `HostAgentApi.Search`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **query** | **string** |  | 
 **to** | **int64** |  | 
 **windowSize** | **int64** |  | 
 **size** | **int32** |  | 
 **offline** | **bool** |  | 

### Return type

[**SnapshotResult**](SnapshotResult.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Update

> Update(ctx, hostId).Execute()

Agent update

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
    hostId := "hostId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.HostAgentApi.Update(context.Background(), hostId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HostAgentApi.Update``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hostId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateConfigurationByHost

> UpdateConfigurationByHost(ctx, hostId).AgentConfigurationUpdate(agentConfigurationUpdate).Execute()

Update agent configuration by host

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
    hostId := "hostId_example" // string | 
    agentConfigurationUpdate := *openapiclient.NewAgentConfigurationUpdate() // AgentConfigurationUpdate |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.HostAgentApi.UpdateConfigurationByHost(context.Background(), hostId).AgentConfigurationUpdate(agentConfigurationUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HostAgentApi.UpdateConfigurationByHost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hostId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateConfigurationByHostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **agentConfigurationUpdate** | [**AgentConfigurationUpdate**](AgentConfigurationUpdate.md) |  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateConfigurationByQuery

> UpdateConfigurationByQuery(ctx).Query(query).To(to).WindowSize(windowSize).Size(size).Offline(offline).AgentConfigurationUpdate(agentConfigurationUpdate).Execute()

Update agent configuration by query

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
    query := "query_example" // string |  (optional)
    to := int64(789) // int64 |  (optional)
    windowSize := int64(789) // int64 |  (optional)
    size := int32(56) // int32 |  (optional)
    offline := true // bool |  (optional)
    agentConfigurationUpdate := *openapiclient.NewAgentConfigurationUpdate() // AgentConfigurationUpdate |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.HostAgentApi.UpdateConfigurationByQuery(context.Background()).Query(query).To(to).WindowSize(windowSize).Size(size).Offline(offline).AgentConfigurationUpdate(agentConfigurationUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HostAgentApi.UpdateConfigurationByQuery``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateConfigurationByQueryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **query** | **string** |  | 
 **to** | **int64** |  | 
 **windowSize** | **int64** |  | 
 **size** | **int32** |  | 
 **offline** | **bool** |  | 
 **agentConfigurationUpdate** | [**AgentConfigurationUpdate**](AgentConfigurationUpdate.md) |  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

