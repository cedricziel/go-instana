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

// TagMatcherDto struct for TagMatcherDto
type TagMatcherDto struct {
	Type     string `json:"type,omitempty"`
	Key      string `json:"key"`
	Entity   string `json:"entity"`
	Value    string `json:"value,omitempty"`
	Operator string `json:"operator"`
}
