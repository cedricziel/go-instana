# \InfrastructureMetricsApi

All URIs are relative to *https://unit-tenant.instana.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetInfrastructureMetrics**](InfrastructureMetricsApi.md#GetInfrastructureMetrics) | **Post** /api/infrastructure-monitoring/metrics | Get infrastructure metrics
[**GetSnapshot**](InfrastructureMetricsApi.md#GetSnapshot) | **Get** /api/infrastructure-monitoring/snapshots/{id} | Get snapshot details
[**GetSnapshots**](InfrastructureMetricsApi.md#GetSnapshots) | **Get** /api/infrastructure-monitoring/snapshots | Search snapshots



## GetInfrastructureMetrics

> InfrastructureMetricResult GetInfrastructureMetrics(ctx).Offline(offline).GetCombinedMetrics(getCombinedMetrics).Execute()

Get infrastructure metrics

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
    offline := true // bool |  (optional)
    getCombinedMetrics := *openapiclient.NewGetCombinedMetrics("Plugin_example", []string{"Metrics_example"}) // GetCombinedMetrics |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InfrastructureMetricsApi.GetInfrastructureMetrics(context.Background()).Offline(offline).GetCombinedMetrics(getCombinedMetrics).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InfrastructureMetricsApi.GetInfrastructureMetrics``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetInfrastructureMetrics`: InfrastructureMetricResult
    fmt.Fprintf(os.Stdout, "Response from `InfrastructureMetricsApi.GetInfrastructureMetrics`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetInfrastructureMetricsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **offline** | **bool** |  | 
 **getCombinedMetrics** | [**GetCombinedMetrics**](GetCombinedMetrics.md) |  | 

### Return type

[**InfrastructureMetricResult**](InfrastructureMetricResult.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSnapshot

> SnapshotItem GetSnapshot(ctx, id).To(to).WindowSize(windowSize).Execute()

Get snapshot details

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
    resp, r, err := api_client.InfrastructureMetricsApi.GetSnapshot(context.Background(), id).To(to).WindowSize(windowSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InfrastructureMetricsApi.GetSnapshot``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSnapshot`: SnapshotItem
    fmt.Fprintf(os.Stdout, "Response from `InfrastructureMetricsApi.GetSnapshot`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSnapshotRequest struct via the builder pattern


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


## GetSnapshots

> SnapshotResult GetSnapshots(ctx).Query(query).To(to).WindowSize(windowSize).Size(size).Plugin(plugin).Offline(offline).Execute()

Search snapshots

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
    plugin := "plugin_example" // string |  (optional)
    offline := true // bool |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InfrastructureMetricsApi.GetSnapshots(context.Background()).Query(query).To(to).WindowSize(windowSize).Size(size).Plugin(plugin).Offline(offline).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InfrastructureMetricsApi.GetSnapshots``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSnapshots`: SnapshotResult
    fmt.Fprintf(os.Stdout, "Response from `InfrastructureMetricsApi.GetSnapshots`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSnapshotsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **query** | **string** |  | 
 **to** | **int64** |  | 
 **windowSize** | **int64** |  | 
 **size** | **int32** |  | 
 **plugin** | **string** |  | 
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

