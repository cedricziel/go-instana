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

// TagFilter struct for TagFilter
type TagFilter struct {
	Name     string `json:"name"`
	Value    string `json:"value"`
	Operator string `json:"operator"`
	Entity   string `json:"entity,omitempty"`
}
