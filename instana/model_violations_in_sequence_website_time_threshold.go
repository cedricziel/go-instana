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

// ViolationsInSequenceWebsiteTimeThreshold struct for ViolationsInSequenceWebsiteTimeThreshold
type ViolationsInSequenceWebsiteTimeThreshold struct {
	Type       string `json:"type"`
	TimeWindow int64  `json:"timeWindow,omitempty"`
}
