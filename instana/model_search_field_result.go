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

// SearchFieldResult struct for SearchFieldResult
type SearchFieldResult struct {
	Keyword     string   `json:"keyword,omitempty"`
	Description string   `json:"description,omitempty"`
	Context     string   `json:"context,omitempty"`
	TermType    string   `json:"termType,omitempty"`
	FixedValues []string `json:"fixedValues,omitempty"`
}
