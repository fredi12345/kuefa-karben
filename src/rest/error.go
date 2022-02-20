package rest

type (
	// swagger:response ErrorResponse
	ErrorResponse struct {
		// in: body
		Body ErrorResponseData
	}

	ErrorResponseData struct {
		// Generic code to detect kind of error
		// Required: true
		// Example: validation.parameter.required
		ErrorCode string `json:"errorCode"`

		// Human-readable description of the error
		// Required: true
		// Example: required field missing in request
		Description string `json:"description"`

		// Additional data to describe the error
		// Required: false
		// Example: {"field": "firstname"}
		AdditionalAttributes map[string]interface{} `json:"additionalAttributes,omitempty"`
	}
)
