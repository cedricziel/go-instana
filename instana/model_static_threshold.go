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

// StaticThreshold struct for StaticThreshold
type StaticThreshold struct {
	Type        string  `json:"type"`
	Operator    string  `json:"operator"`
	LastUpdated int64   `json:"lastUpdated,omitempty"`
	Value       float64 `json:"value,omitempty"`
}
