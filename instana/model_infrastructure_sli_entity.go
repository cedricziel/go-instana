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

// InfrastructureSliEntity struct for InfrastructureSliEntity
type InfrastructureSliEntity struct {
	SliType             string `json:"sliType"`
	InstanceAggregation string `json:"instanceAggregation"`
	Query               string `json:"query"`
}
