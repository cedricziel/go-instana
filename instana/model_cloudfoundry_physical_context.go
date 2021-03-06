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

// CloudfoundryPhysicalContext struct for CloudfoundryPhysicalContext
type CloudfoundryPhysicalContext struct {
	Application     SnapshotPreview `json:"application,omitempty"`
	Space           SnapshotPreview `json:"space,omitempty"`
	Organization    SnapshotPreview `json:"organization,omitempty"`
	CfInstanceIndex string          `json:"cfInstanceIndex,omitempty"`
}
