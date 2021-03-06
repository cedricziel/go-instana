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

// GetWebsiteMetrics struct for GetWebsiteMetrics
type GetWebsiteMetrics struct {
	TimeFrame  TimeFrame                               `json:"timeFrame,omitempty"`
	Metrics    []WebsiteMonitoringMetricsConfiguration `json:"metrics"`
	Type       string                                  `json:"type"`
	TagFilters []TagFilter                             `json:"tagFilters,omitempty"`
}
