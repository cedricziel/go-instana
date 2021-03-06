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

// Release struct for Release
type Release struct {
	Name         string             `json:"name"`
	Start        int64              `json:"start,omitempty"`
	Services     []ServiceScope     `json:"services,omitempty"`
	Applications []ApplicationScope `json:"applications,omitempty"`
}
