# \SLISettingsApi

All URIs are relative to *https://unit-tenant.instana.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSli**](SLISettingsApi.md#CreateSli) | **Post** /api/settings/sli | Create SLI Config
[**Delete**](SLISettingsApi.md#Delete) | **Delete** /api/settings/sli/{id} | Delete SLI Config
[**GetSli1**](SLISettingsApi.md#GetSli1) | **Get** /api/settings/sli/{id} | Get SLI Config
[**GetSli2**](SLISettingsApi.md#GetSli2) | **Get** /api/settings/sli | Get All SLI Configs
[**PutSli**](SLISettingsApi.md#PutSli) | **Put** /api/settings/sli/{id} | Update SLI Config



## CreateSli

> []SliConfigurationWithLastUpdated CreateSli(ctx).SliConfiguration(sliConfiguration).Execute()

Create SLI Config

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
    sliConfiguration := *openapiclient.NewSliConfiguration("Id_example", "SliName_example", *openapiclient.NewSliEntity("SliType_example")) // SliConfiguration | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SLISettingsApi.CreateSli(context.Background()).SliConfiguration(sliConfiguration).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SLISettingsApi.CreateSli``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSli`: []SliConfigurationWithLastUpdated
    fmt.Fprintf(os.Stdout, "Response from `SLISettingsApi.CreateSli`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSliRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sliConfiguration** | [**SliConfiguration**](SliConfiguration.md) |  | 

### Return type

[**[]SliConfigurationWithLastUpdated**](SliConfigurationWithLastUpdated.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Delete

> Delete(ctx, id).Execute()

Delete SLI Config

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SLISettingsApi.Delete(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SLISettingsApi.Delete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRequest struct via the builder pattern


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


## GetSli1

> SliConfigurationWithLastUpdated GetSli1(ctx, id).Execute()

Get SLI Config

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SLISettingsApi.GetSli1(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SLISettingsApi.GetSli1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSli1`: SliConfigurationWithLastUpdated
    fmt.Fprintf(os.Stdout, "Response from `SLISettingsApi.GetSli1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSli1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SliConfigurationWithLastUpdated**](SliConfigurationWithLastUpdated.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSli2

> []SliConfigurationWithLastUpdated GetSli2(ctx).Execute()

Get All SLI Configs

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
    resp, r, err := api_client.SLISettingsApi.GetSli2(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SLISettingsApi.GetSli2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSli2`: []SliConfigurationWithLastUpdated
    fmt.Fprintf(os.Stdout, "Response from `SLISettingsApi.GetSli2`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSli2Request struct via the builder pattern


### Return type

[**[]SliConfigurationWithLastUpdated**](SliConfigurationWithLastUpdated.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutSli

> []SliConfigurationWithLastUpdated PutSli(ctx, id).SliConfigurationWithLastUpdated(sliConfigurationWithLastUpdated).Execute()

Update SLI Config

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
    sliConfigurationWithLastUpdated := *openapiclient.NewSliConfigurationWithLastUpdated("Id_example", "SliName_example", *openapiclient.NewSliEntity("SliType_example")) // SliConfigurationWithLastUpdated | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SLISettingsApi.PutSli(context.Background(), id).SliConfigurationWithLastUpdated(sliConfigurationWithLastUpdated).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SLISettingsApi.PutSli``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutSli`: []SliConfigurationWithLastUpdated
    fmt.Fprintf(os.Stdout, "Response from `SLISettingsApi.PutSli`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutSliRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sliConfigurationWithLastUpdated** | [**SliConfigurationWithLastUpdated**](SliConfigurationWithLastUpdated.md) |  | 

### Return type

[**[]SliConfigurationWithLastUpdated**](SliConfigurationWithLastUpdated.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

