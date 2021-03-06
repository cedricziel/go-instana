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

// ApplicationConfig struct for ApplicationConfig
type ApplicationConfig struct {
	Id                 string             `json:"id"`
	Label              string             `json:"label"`
	MatchSpecification MatchExpressionDto `json:"matchSpecification"`
	Scope              string             `json:"scope"`
	BoundaryScope      string             `json:"boundaryScope"`
}
