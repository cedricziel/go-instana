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

// CustomDashboard struct for CustomDashboard
type CustomDashboard struct {
	Id          string       `json:"id"`
	Title       string       `json:"title"`
	AccessRules []AccessRule `json:"accessRules"`
	Widgets     []Widget     `json:"widgets"`
}
