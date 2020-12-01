# \WebsiteAnalyzeApi

All URIs are relative to *https://unit-tenant.instana.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetBeaconGroups**](WebsiteAnalyzeApi.md#GetBeaconGroups) | **Post** /api/website-monitoring/analyze/beacon-groups | Get grouped beacon metrics
[**GetBeacons**](WebsiteAnalyzeApi.md#GetBeacons) | **Post** /api/website-monitoring/analyze/beacons | Get all beacons



## GetBeaconGroups

> BeaconGroupsResult GetBeaconGroups(ctx).FillTimeSeries(fillTimeSeries).GetWebsiteBeaconGroups(getWebsiteBeaconGroups).Execute()

Get grouped beacon metrics

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
    fillTimeSeries := true // bool |  (optional)
    getWebsiteBeaconGroups := *openapiclient.NewGetWebsiteBeaconGroups(*openapiclient.NewWebsiteBeaconTagGroup("GroupbyTag_example", "GroupbyTagEntity_example"), "Type_example", []openapiclient.WebsiteMonitoringMetricsConfiguration{*openapiclient.NewWebsiteMonitoringMetricsConfiguration("Metric_example", "Aggregation_example")}) // GetWebsiteBeaconGroups |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WebsiteAnalyzeApi.GetBeaconGroups(context.Background()).FillTimeSeries(fillTimeSeries).GetWebsiteBeaconGroups(getWebsiteBeaconGroups).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebsiteAnalyzeApi.GetBeaconGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBeaconGroups`: BeaconGroupsResult
    fmt.Fprintf(os.Stdout, "Response from `WebsiteAnalyzeApi.GetBeaconGroups`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetBeaconGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fillTimeSeries** | **bool** |  | 
 **getWebsiteBeaconGroups** | [**GetWebsiteBeaconGroups**](GetWebsiteBeaconGroups.md) |  | 

### Return type

[**BeaconGroupsResult**](BeaconGroupsResult.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBeacons

> BeaconResult GetBeacons(ctx).GetWebsiteBeacons(getWebsiteBeacons).Execute()

Get all beacons

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
    getWebsiteBeacons := *openapiclient.NewGetWebsiteBeacons("Type_example") // GetWebsiteBeacons |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WebsiteAnalyzeApi.GetBeacons(context.Background()).GetWebsiteBeacons(getWebsiteBeacons).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebsiteAnalyzeApi.GetBeacons``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBeacons`: BeaconResult
    fmt.Fprintf(os.Stdout, "Response from `WebsiteAnalyzeApi.GetBeacons`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetBeaconsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getWebsiteBeacons** | [**GetWebsiteBeacons**](GetWebsiteBeacons.md) |  | 

### Return type

[**BeaconResult**](BeaconResult.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

