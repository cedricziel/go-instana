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

// SyntheticCallWithDefaultsConfig struct for SyntheticCallWithDefaultsConfig
type SyntheticCallWithDefaultsConfig struct {
	DefaultRulesEnabled bool                `json:"defaultRulesEnabled,omitempty"`
	CustomRules         []SyntheticCallRule `json:"customRules"`
	DefaultRules        []SyntheticCallRule `json:"defaultRules"`
}
