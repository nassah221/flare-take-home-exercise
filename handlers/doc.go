// Package classification of CheckUsername API
//
// Documentation for CheckUsername API
//
//	Schemes: http
//	BasePath: /
//	Version: 1.0.0
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
// swagger:meta
package handlers

// Query parameter for the checkUsername endpoint
// swagger:parameters checkUsernameQuery
type usernameQueryParameterWrapper struct { //nolint
	// in: query
	// required: true
	Query string
}

// Response of the checkUsername endpoint
// swagger:response usernameResponse
type usernameResponseWrapper struct { //nolint
	// in: body
	Body UsernameResponse
}

// Response of the checkHealth endpoint
// swagger:response healthResponse
type healthResponseWrapper struct { //nolint
	// in: body
	Body HealthResponse
}

// Generic error message returned as a string
// swagger:response errorResponse
type errResponseWrapper struct { //nolint
	// Description of the error
	// in: body
	Body GenericError
}
