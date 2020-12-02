# \HealthApi

All URIs are relative to *https://unit-tenant.instana.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetHealthState**](HealthApi.md#GetHealthState) | **Get** /api/instana/health | Basic health traffic light
[**GetVersion**](HealthApi.md#GetVersion) | **Get** /api/instana/version | API version information



## GetHealthState

> HealthState GetHealthState(ctx, )

Basic health traffic light

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**HealthState**](HealthState.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVersion

> InstanaVersionInfo GetVersion(ctx, )

API version information

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**InstanaVersionInfo**](InstanaVersionInfo.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

