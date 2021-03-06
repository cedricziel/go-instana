/*
 * Introduction to Instana public APIs
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.192.86
 * Contact: support@instana.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package instana

// ApplicationMetricResult struct for ApplicationMetricResult
type ApplicationMetricResult struct {
	Items     []ApplicationItem `json:"items"`
	Page      int32             `json:"page,omitempty"`
	PageSize  int32             `json:"pageSize,omitempty"`
	TotalHits int64             `json:"totalHits,omitempty"`
}
