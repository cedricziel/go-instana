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

// FullTrace struct for FullTrace
type FullTrace struct {
	Id              string `json:"id"`
	TotalErrorCount int32  `json:"totalErrorCount,omitempty"`
	RootSpan        Span   `json:"rootSpan"`
}
