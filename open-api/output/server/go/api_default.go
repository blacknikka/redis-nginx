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

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

// A DefaultApiController binds http requests to an api service and writes the service results to the http response
type DefaultApiController struct {
	service DefaultApiServicer
}

// NewDefaultApiController creates a default api controller
func NewDefaultApiController(s DefaultApiServicer) Router {
	return &DefaultApiController{ service: s }
}

// Routes returns all of the api route for the DefaultApiController
func (c *DefaultApiController) Routes() Routes {
	return Routes{ 
		{
			"TestPost",
			strings.ToUpper("Post"),
			"/hello/test",
			c.TestPost,
		},
	}
}

// TestPost - 
func (c *DefaultApiController) TestPost(w http.ResponseWriter, r *http.Request) { 
	testPost := &TestPost{}
	if err := json.NewDecoder(r.Body).Decode(&testPost); err != nil {
		w.WriteHeader(500)
		return
	}
	
	result, err := c.service.TestPost(*testPost)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	
	EncodeJSONResponse(result, nil, w)
}