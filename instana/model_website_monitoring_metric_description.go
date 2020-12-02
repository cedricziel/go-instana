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

// WebsiteMonitoringMetricDescription struct for WebsiteMonitoringMetricDescription
type WebsiteMonitoringMetricDescription struct {
	MetricId     string   `json:"metricId,omitempty"`
	Label        string   `json:"label,omitempty"`
	Formatter    string   `json:"formatter,omitempty"`
	Description  string   `json:"description,omitempty"`
	Aggregations []string `json:"aggregations,omitempty"`
	BeaconTypes  []string `json:"beaconTypes,omitempty"`
}
