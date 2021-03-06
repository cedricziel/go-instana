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

// ServiceScopeWithMetadata struct for ServiceScopeWithMetadata
type ServiceScopeWithMetadata struct {
	Id       string                      `json:"id"`
	Name     string                      `json:"name,omitempty"`
	ScopedTo ServiceScopedToWithMetadata `json:"scopedTo,omitempty"`
}
