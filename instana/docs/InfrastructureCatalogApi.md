# \InfrastructureCatalogApi

All URIs are relative to *https://unit-tenant.instana.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetInfrastructureCatalogMetrics**](InfrastructureCatalogApi.md#GetInfrastructureCatalogMetrics) | **Get** /api/infrastructure-monitoring/catalog/metrics/{plugin} | Get metric catalog
[**GetInfrastructureCatalogPlugins**](InfrastructureCatalogApi.md#GetInfrastructureCatalogPlugins) | **Get** /api/infrastructure-monitoring/catalog/plugins | Get plugin catalog
[**GetInfrastructureCatalogPluginsWithCustomMetrics**](InfrastructureCatalogApi.md#GetInfrastructureCatalogPluginsWithCustomMetrics) | **Get** /api/infrastructure-monitoring/catalog/plugins-with-custom-metrics | Get all plugins with custom metrics catalog
[**GetInfrastructureCatalogSearchFields**](InfrastructureCatalogApi.md#GetInfrastructureCatalogSearchFields) | **Get** /api/infrastructure-monitoring/catalog/search | get search field catalog



## GetInfrastructureCatalogMetrics

> []MetricInstance GetInfrastructureCatalogMetrics(ctx, plugin, optional)

Get metric catalog

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**plugin** | **string**|  | 
 **optional** | ***GetInfrastructureCatalogMetricsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetInfrastructureCatalogMetricsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filter** | **optional.String**|  | 

### Return type

[**[]MetricInstance**](MetricInstance.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInfrastructureCatalogPlugins

> []PluginResult GetInfrastructureCatalogPlugins(ctx, )

Get plugin catalog

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]PluginResult**](PluginResult.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInfrastructureCatalogPluginsWithCustomMetrics

> []PluginResult GetInfrastructureCatalogPluginsWithCustomMetrics(ctx, )

Get all plugins with custom metrics catalog

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]PluginResult**](PluginResult.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInfrastructureCatalogSearchFields

> []SearchFieldResult GetInfrastructureCatalogSearchFields(ctx, )

get search field catalog

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]SearchFieldResult**](SearchFieldResult.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

