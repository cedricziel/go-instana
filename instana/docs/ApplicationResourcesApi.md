# \ApplicationResourcesApi

All URIs are relative to *https://unit-tenant.instana.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApplicationResourcesEndpoints**](ApplicationResourcesApi.md#ApplicationResourcesEndpoints) | **Get** /api/application-monitoring/applications/services/endpoints | Get endpoints
[**GetApplicationServices**](ApplicationResourcesApi.md#GetApplicationServices) | **Get** /api/application-monitoring/applications/services | Get applications/services
[**GetApplications**](ApplicationResourcesApi.md#GetApplications) | **Get** /api/application-monitoring/applications | Get applications
[**GetServices**](ApplicationResourcesApi.md#GetServices) | **Get** /api/application-monitoring/services | Get services



## ApplicationResourcesEndpoints

> EndpointResult ApplicationResourcesEndpoints(ctx).NameFilter(nameFilter).Types(types).Technologies(technologies).WindowSize(windowSize).To(to).Page(page).PageSize(pageSize).ApplicationBoundaryScope(applicationBoundaryScope).Execute()

Get endpoints

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
    nameFilter := "nameFilter_example" // string |  (optional)
    types := []string{"Inner_example"} // []string |  (optional)
    technologies := []string{"Inner_example"} // []string |  (optional)
    windowSize := int64(789) // int64 |  (optional)
    to := int64(789) // int64 |  (optional)
    page := int32(56) // int32 |  (optional)
    pageSize := int32(56) // int32 |  (optional)
    applicationBoundaryScope := "applicationBoundaryScope_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ApplicationResourcesApi.ApplicationResourcesEndpoints(context.Background()).NameFilter(nameFilter).Types(types).Technologies(technologies).WindowSize(windowSize).To(to).Page(page).PageSize(pageSize).ApplicationBoundaryScope(applicationBoundaryScope).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationResourcesApi.ApplicationResourcesEndpoints``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApplicationResourcesEndpoints`: EndpointResult
    fmt.Fprintf(os.Stdout, "Response from `ApplicationResourcesApi.ApplicationResourcesEndpoints`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApplicationResourcesEndpointsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nameFilter** | **string** |  | 
 **types** | **[]string** |  | 
 **technologies** | **[]string** |  | 
 **windowSize** | **int64** |  | 
 **to** | **int64** |  | 
 **page** | **int32** |  | 
 **pageSize** | **int32** |  | 
 **applicationBoundaryScope** | **string** |  | 

### Return type

[**EndpointResult**](EndpointResult.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApplicationServices

> ServiceResult GetApplicationServices(ctx).NameFilter(nameFilter).WindowSize(windowSize).To(to).Page(page).PageSize(pageSize).ApplicationBoundaryScope(applicationBoundaryScope).Execute()

Get applications/services

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
    nameFilter := "nameFilter_example" // string |  (optional)
    windowSize := int64(789) // int64 |  (optional)
    to := int64(789) // int64 |  (optional)
    page := int32(56) // int32 |  (optional)
    pageSize := int32(56) // int32 |  (optional)
    applicationBoundaryScope := "applicationBoundaryScope_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ApplicationResourcesApi.GetApplicationServices(context.Background()).NameFilter(nameFilter).WindowSize(windowSize).To(to).Page(page).PageSize(pageSize).ApplicationBoundaryScope(applicationBoundaryScope).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationResourcesApi.GetApplicationServices``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApplicationServices`: ServiceResult
    fmt.Fprintf(os.Stdout, "Response from `ApplicationResourcesApi.GetApplicationServices`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationServicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nameFilter** | **string** |  | 
 **windowSize** | **int64** |  | 
 **to** | **int64** |  | 
 **page** | **int32** |  | 
 **pageSize** | **int32** |  | 
 **applicationBoundaryScope** | **string** |  | 

### Return type

[**ServiceResult**](ServiceResult.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApplications

> ApplicationResult GetApplications(ctx).NameFilter(nameFilter).WindowSize(windowSize).To(to).Page(page).PageSize(pageSize).ApplicationBoundaryScope(applicationBoundaryScope).Execute()

Get applications

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
    nameFilter := "nameFilter_example" // string |  (optional)
    windowSize := int64(789) // int64 |  (optional)
    to := int64(789) // int64 |  (optional)
    page := int32(56) // int32 |  (optional)
    pageSize := int32(56) // int32 |  (optional)
    applicationBoundaryScope := "applicationBoundaryScope_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ApplicationResourcesApi.GetApplications(context.Background()).NameFilter(nameFilter).WindowSize(windowSize).To(to).Page(page).PageSize(pageSize).ApplicationBoundaryScope(applicationBoundaryScope).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationResourcesApi.GetApplications``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApplications`: ApplicationResult
    fmt.Fprintf(os.Stdout, "Response from `ApplicationResourcesApi.GetApplications`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nameFilter** | **string** |  | 
 **windowSize** | **int64** |  | 
 **to** | **int64** |  | 
 **page** | **int32** |  | 
 **pageSize** | **int32** |  | 
 **applicationBoundaryScope** | **string** |  | 

### Return type

[**ApplicationResult**](ApplicationResult.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServices

> ServiceResult GetServices(ctx).NameFilter(nameFilter).WindowSize(windowSize).To(to).Page(page).PageSize(pageSize).Execute()

Get services

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
    nameFilter := "nameFilter_example" // string |  (optional)
    windowSize := int64(789) // int64 |  (optional)
    to := int64(789) // int64 |  (optional)
    page := int32(56) // int32 |  (optional)
    pageSize := int32(56) // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ApplicationResourcesApi.GetServices(context.Background()).NameFilter(nameFilter).WindowSize(windowSize).To(to).Page(page).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationResourcesApi.GetServices``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetServices`: ServiceResult
    fmt.Fprintf(os.Stdout, "Response from `ApplicationResourcesApi.GetServices`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetServicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nameFilter** | **string** |  | 
 **windowSize** | **int64** |  | 
 **to** | **int64** |  | 
 **page** | **int32** |  | 
 **pageSize** | **int32** |  | 

### Return type

[**ServiceResult**](ServiceResult.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

