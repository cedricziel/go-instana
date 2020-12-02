# \ReleasesApi

All URIs are relative to *https://unit-tenant.instana.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteRelease**](ReleasesApi.md#DeleteRelease) | **Delete** /api/releases/{releaseId} | Delete release
[**GetAllReleases**](ReleasesApi.md#GetAllReleases) | **Get** /api/releases | Get all releases
[**GetRelease**](ReleasesApi.md#GetRelease) | **Get** /api/releases/{releaseId} | Get release
[**PostRelease**](ReleasesApi.md#PostRelease) | **Post** /api/releases | Create release
[**PutRelease**](ReleasesApi.md#PutRelease) | **Put** /api/releases/{releaseId} | Update release



## DeleteRelease

> DeleteRelease(ctx, releaseId)

Delete release

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**releaseId** | **string**|  | 

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


## GetAllReleases

> []ReleaseWithMetadata GetAllReleases(ctx, optional)

Get all releases

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetAllReleasesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetAllReleasesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **from** | **optional.Int64**|  | 
 **to** | **optional.Int64**|  | 
 **maxResults** | **optional.Int32**|  | 

### Return type

[**[]ReleaseWithMetadata**](ReleaseWithMetadata.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRelease

> ReleaseWithMetadata GetRelease(ctx, releaseId)

Get release

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**releaseId** | **string**|  | 

### Return type

[**ReleaseWithMetadata**](ReleaseWithMetadata.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostRelease

> ReleaseWithMetadata PostRelease(ctx, release)

Create release

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**release** | [**Release**](Release.md)|  | 

### Return type

[**ReleaseWithMetadata**](ReleaseWithMetadata.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutRelease

> ReleaseWithMetadata PutRelease(ctx, releaseId, release)

Update release

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**releaseId** | **string**|  | 
**release** | [**Release**](Release.md)|  | 

### Return type

[**ReleaseWithMetadata**](ReleaseWithMetadata.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

