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

// MyApiService -
type MyApiService struct {
}

// NewMyApiService creates a default api service
func NewMyApiService() *MyApiService {
	return &MyApiService{}
}

// TestGet -
func (s *MyApiService) TestGet() (interface{}, error) {
	return TestPostResponse{
		Name: "my name",
		Tag:  "my tag",
		Id:   1,
	}, nil
}

// TestPost -
func (s *MyApiService) TestPost(testPost TestPost) (interface{}, error) {
	return TestPostResponse{
		Name: "my name",
		Tag:  "my tag",
		Id:   1,
	}, nil
}