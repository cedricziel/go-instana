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

// UsageResult struct for UsageResult
type UsageResult struct {
	Time  int64              `json:"time,omitempty"`
	Items []UsageResultItems `json:"items,omitempty"`
}
