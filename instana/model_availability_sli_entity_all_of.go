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

// AvailabilitySliEntityAllOf struct for AvailabilitySliEntityAllOf
type AvailabilitySliEntityAllOf struct {
	ApplicationId    string      `json:"applicationId,omitempty"`
	ServiceId        string      `json:"serviceId,omitempty"`
	EndpointId       string      `json:"endpointId,omitempty"`
	BoundaryScope    string      `json:"boundaryScope,omitempty"`
	GoodEventFilters []TagFilter `json:"goodEventFilters,omitempty"`
	BadEventFilters  []TagFilter `json:"badEventFilters,omitempty"`
}
