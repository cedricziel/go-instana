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

// IntegrationOverview struct for IntegrationOverview
type IntegrationOverview struct {
	Id         string            `json:"id"`
	Kind       string            `json:"kind"`
	Name       string            `json:"name"`
	Properties map[string]string `json:"properties,omitempty"`
}
