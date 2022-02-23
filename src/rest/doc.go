// Package rest contains the API handlers for Kuefa-API.
//
// @Version 1.0.0
// @Title Kuefa-API
// @Description This is the Rest-API specification for server-side API.
// @Server http://localhost:3001/api Server-1
package rest

//go:generate goas --module-path ../../. --output ../../res/kuefa-api.json --main-file-path doc.go --omit-packages
