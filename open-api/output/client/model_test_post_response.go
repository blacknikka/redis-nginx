/*
 * my-swagger
 *
 * This is my first sample for Swagger (Open-API)
 *
 * API version: 1.0.0
 * Contact: example@example.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// TestPostResponse struct for TestPostResponse
type TestPostResponse struct {
	Name string `json:"name"`
	Tag string `json:"tag,omitempty"`
	Id int64 `json:"id"`
}
