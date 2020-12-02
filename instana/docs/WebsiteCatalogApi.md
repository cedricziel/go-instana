# \WebsiteCatalogApi

All URIs are relative to *https://unit-tenant.instana.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetWebsiteCatalogMetrics**](WebsiteCatalogApi.md#GetWebsiteCatalogMetrics) | **Get** /api/website-monitoring/catalog/metrics | Metric catalog
[**GetWebsiteCatalogTags**](WebsiteCatalogApi.md#GetWebsiteCatalogTags) | **Get** /api/website-monitoring/catalog/tags | Get all existing website tags
[**GetWebsiteTagCatalog**](WebsiteCatalogApi.md#GetWebsiteTagCatalog) | **Get** /api/website-monitoring/catalog | Get website tag catalog



## GetWebsiteCatalogMetrics

> []WebsiteMonitoringMetricDescription GetWebsiteCatalogMetrics(ctx, )

Metric catalog

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]WebsiteMonitoringMetricDescription**](WebsiteMonitoringMetricDescription.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWebsiteCatalogTags

> []Tag GetWebsiteCatalogTags(ctx, )

Get all existing website tags

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]Tag**](Tag.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWebsiteTagCatalog

> TagCatalog GetWebsiteTagCatalog(ctx, beaconType, useCase)

Get website tag catalog

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**beaconType** | **string**|  | 
**useCase** | **string**|  | 

### Return type

[**TagCatalog**](TagCatalog.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

