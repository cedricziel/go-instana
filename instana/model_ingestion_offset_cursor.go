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

// IngestionOffsetCursor struct for IngestionOffsetCursor
type IngestionOffsetCursor struct {
	IngestionTime int64 `json:"ingestionTime,omitempty"`
	Offset        int32 `json:"offset,omitempty"`
}
