// Package rest Kuefa-API
//
// This is the Rest-API specification for server-side API.
//
//     Schemes: http, https
//     Host: localhost:3001
//     BasePath: /api
//     Version: 1.0.0
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
// swagger:meta
package rest

//go:generate env SWAGGER_GENERATE_EXTENSION=false swagger generate spec -o ../../res/kuefa-api.yaml
