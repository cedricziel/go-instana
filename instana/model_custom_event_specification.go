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

// CustomEventSpecification struct for CustomEventSpecification
type CustomEventSpecification struct {
	Id             string         `json:"id"`
	Name           string         `json:"name"`
	EntityType     string         `json:"entityType"`
	Query          string         `json:"query,omitempty"`
	Triggering     bool           `json:"triggering,omitempty"`
	Description    string         `json:"description,omitempty"`
	ExpirationTime int64          `json:"expirationTime,omitempty"`
	Enabled        bool           `json:"enabled,omitempty"`
	Rules          []AbstractRule `json:"rules"`
}
