# \InfrastructureResourcesApi

All URIs are relative to *https://unit-tenant.instana.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetInfrastructureViewTree**](InfrastructureResourcesApi.md#GetInfrastructureViewTree) | **Get** /api/infrastructure-monitoring/graph/views | Get view tree
[**GetMonitoringState**](InfrastructureResourcesApi.md#GetMonitoringState) | **Get** /api/infrastructure-monitoring/monitoring-state | Monitored host count
[**GetRelatedHosts**](InfrastructureResourcesApi.md#GetRelatedHosts) | **Get** /api/infrastructure-monitoring/graph/related-hosts/{snapshotId} | Related hosts
[**SoftwareVersions**](InfrastructureResourcesApi.md#SoftwareVersions) | **Get** /api/infrastructure-monitoring/software/versions | Get installed software



## GetInfrastructureViewTree

> TreeNodeResult GetInfrastructureViewTree(ctx).Execute()

Get view tree

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InfrastructureResourcesApi.GetInfrastructureViewTree(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InfrastructureResourcesApi.GetInfrastructureViewTree``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetInfrastructureViewTree`: TreeNodeResult
    fmt.Fprintf(os.Stdout, "Response from `InfrastructureResourcesApi.GetInfrastructureViewTree`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetInfrastructureViewTreeRequest struct via the builder pattern


### Return type

[**TreeNodeResult**](TreeNodeResult.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMonitoringState

> MonitoringState GetMonitoringState(ctx).Execute()

Monitored host count

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InfrastructureResourcesApi.GetMonitoringState(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InfrastructureResourcesApi.GetMonitoringState``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMonitoringState`: MonitoringState
    fmt.Fprintf(os.Stdout, "Response from `InfrastructureResourcesApi.GetMonitoringState`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetMonitoringStateRequest struct via the builder pattern


### Return type

[**MonitoringState**](MonitoringState.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRelatedHosts

> []string GetRelatedHosts(ctx, snapshotId).Execute()

Related hosts

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
    snapshotId := "snapshotId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InfrastructureResourcesApi.GetRelatedHosts(context.Background(), snapshotId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InfrastructureResourcesApi.GetRelatedHosts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRelatedHosts`: []string
    fmt.Fprintf(os.Stdout, "Response from `InfrastructureResourcesApi.GetRelatedHosts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**snapshotId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRelatedHostsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**[]string**

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SoftwareVersions

> []SoftwareVersion SoftwareVersions(ctx).Time(time).Origin(origin).Type_(type_).Name(name).Version(version).Execute()

Get installed software



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
    time := int64(789) // int64 |  (optional)
    origin := "origin_example" // string |  (optional)
    type_ := "type__example" // string |  (optional)
    name := "name_example" // string |  (optional)
    version := "version_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InfrastructureResourcesApi.SoftwareVersions(context.Background()).Time(time).Origin(origin).Type_(type_).Name(name).Version(version).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InfrastructureResourcesApi.SoftwareVersions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SoftwareVersions`: []SoftwareVersion
    fmt.Fprintf(os.Stdout, "Response from `InfrastructureResourcesApi.SoftwareVersions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSoftwareVersionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **time** | **int64** |  | 
 **origin** | **string** |  | 
 **type_** | **string** |  | 
 **name** | **string** |  | 
 **version** | **string** |  | 

### Return type

[**[]SoftwareVersion**](SoftwareVersion.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

