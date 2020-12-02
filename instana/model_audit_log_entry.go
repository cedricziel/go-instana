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

// AuditLogEntry struct for AuditLogEntry
type AuditLogEntry struct {
	Id        string                 `json:"id"`
	Action    string                 `json:"action"`
	Message   string                 `json:"message"`
	Actor     LogEntryActor          `json:"actor"`
	Timestamp int64                  `json:"timestamp,omitempty"`
	Meta      map[string]interface{} `json:"meta"`
}
