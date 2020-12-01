/*
 * Introduction to Instana public APIs
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.192.86
 * Contact: support@instana.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package instana

import (
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
	"strings"
)

// Linger please
var (
	_ _context.Context
)

// InfrastructureMetricsApiService InfrastructureMetricsApi service
type InfrastructureMetricsApiService service

type ApiGetInfrastructureMetricsRequest struct {
	ctx                _context.Context
	ApiService         *InfrastructureMetricsApiService
	offline            *bool
	getCombinedMetrics *GetCombinedMetrics
}

func (r ApiGetInfrastructureMetricsRequest) Offline(offline bool) ApiGetInfrastructureMetricsRequest {
	r.offline = &offline
	return r
}
func (r ApiGetInfrastructureMetricsRequest) GetCombinedMetrics(getCombinedMetrics GetCombinedMetrics) ApiGetInfrastructureMetricsRequest {
	r.getCombinedMetrics = &getCombinedMetrics
	return r
}

func (r ApiGetInfrastructureMetricsRequest) Execute() (InfrastructureMetricResult, *_nethttp.Response, error) {
	return r.ApiService.GetInfrastructureMetricsExecute(r)
}

/*
 * GetInfrastructureMetrics Get infrastructure metrics
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiGetInfrastructureMetricsRequest
 */
func (a *InfrastructureMetricsApiService) GetInfrastructureMetrics(ctx _context.Context) ApiGetInfrastructureMetricsRequest {
	return ApiGetInfrastructureMetricsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

/*
 * Execute executes the request
 * @return InfrastructureMetricResult
 */
func (a *InfrastructureMetricsApiService) GetInfrastructureMetricsExecute(r ApiGetInfrastructureMetricsRequest) (InfrastructureMetricResult, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  InfrastructureMetricResult
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InfrastructureMetricsApiService.GetInfrastructureMetrics")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/infrastructure-monitoring/metrics"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.offline != nil {
		localVarQueryParams.Add("offline", parameterToString(*r.offline, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.getCombinedMetrics
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ApiKeyAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v InfrastructureMetricResult
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetSnapshotRequest struct {
	ctx        _context.Context
	ApiService *InfrastructureMetricsApiService
	id         string
	to         *int64
	windowSize *int64
}

func (r ApiGetSnapshotRequest) To(to int64) ApiGetSnapshotRequest {
	r.to = &to
	return r
}
func (r ApiGetSnapshotRequest) WindowSize(windowSize int64) ApiGetSnapshotRequest {
	r.windowSize = &windowSize
	return r
}

func (r ApiGetSnapshotRequest) Execute() (SnapshotItem, *_nethttp.Response, error) {
	return r.ApiService.GetSnapshotExecute(r)
}

/*
 * GetSnapshot Get snapshot details
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param id
 * @return ApiGetSnapshotRequest
 */
func (a *InfrastructureMetricsApiService) GetSnapshot(ctx _context.Context, id string) ApiGetSnapshotRequest {
	return ApiGetSnapshotRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

/*
 * Execute executes the request
 * @return SnapshotItem
 */
func (a *InfrastructureMetricsApiService) GetSnapshotExecute(r ApiGetSnapshotRequest) (SnapshotItem, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  SnapshotItem
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InfrastructureMetricsApiService.GetSnapshot")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/infrastructure-monitoring/snapshots/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", _neturl.PathEscape(parameterToString(r.id, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.to != nil {
		localVarQueryParams.Add("to", parameterToString(*r.to, ""))
	}
	if r.windowSize != nil {
		localVarQueryParams.Add("windowSize", parameterToString(*r.windowSize, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ApiKeyAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v SnapshotItem
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetSnapshotsRequest struct {
	ctx        _context.Context
	ApiService *InfrastructureMetricsApiService
	query      *string
	to         *int64
	windowSize *int64
	size       *int32
	plugin     *string
	offline    *bool
}

func (r ApiGetSnapshotsRequest) Query(query string) ApiGetSnapshotsRequest {
	r.query = &query
	return r
}
func (r ApiGetSnapshotsRequest) To(to int64) ApiGetSnapshotsRequest {
	r.to = &to
	return r
}
func (r ApiGetSnapshotsRequest) WindowSize(windowSize int64) ApiGetSnapshotsRequest {
	r.windowSize = &windowSize
	return r
}
func (r ApiGetSnapshotsRequest) Size(size int32) ApiGetSnapshotsRequest {
	r.size = &size
	return r
}
func (r ApiGetSnapshotsRequest) Plugin(plugin string) ApiGetSnapshotsRequest {
	r.plugin = &plugin
	return r
}
func (r ApiGetSnapshotsRequest) Offline(offline bool) ApiGetSnapshotsRequest {
	r.offline = &offline
	return r
}

func (r ApiGetSnapshotsRequest) Execute() (SnapshotResult, *_nethttp.Response, error) {
	return r.ApiService.GetSnapshotsExecute(r)
}

/*
 * GetSnapshots Search snapshots
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiGetSnapshotsRequest
 */
func (a *InfrastructureMetricsApiService) GetSnapshots(ctx _context.Context) ApiGetSnapshotsRequest {
	return ApiGetSnapshotsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

/*
 * Execute executes the request
 * @return SnapshotResult
 */
func (a *InfrastructureMetricsApiService) GetSnapshotsExecute(r ApiGetSnapshotsRequest) (SnapshotResult, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  SnapshotResult
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InfrastructureMetricsApiService.GetSnapshots")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/infrastructure-monitoring/snapshots"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.query != nil {
		localVarQueryParams.Add("query", parameterToString(*r.query, ""))
	}
	if r.to != nil {
		localVarQueryParams.Add("to", parameterToString(*r.to, ""))
	}
	if r.windowSize != nil {
		localVarQueryParams.Add("windowSize", parameterToString(*r.windowSize, ""))
	}
	if r.size != nil {
		localVarQueryParams.Add("size", parameterToString(*r.size, ""))
	}
	if r.plugin != nil {
		localVarQueryParams.Add("plugin", parameterToString(*r.plugin, ""))
	}
	if r.offline != nil {
		localVarQueryParams.Add("offline", parameterToString(*r.offline, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ApiKeyAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v SnapshotResult
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
