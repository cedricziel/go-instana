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

// MaintenanceWindow struct for MaintenanceWindow
type MaintenanceWindow struct {
	Id    string `json:"id"`
	Start int64  `json:"start,omitempty"`
	End   int64  `json:"end,omitempty"`
}
