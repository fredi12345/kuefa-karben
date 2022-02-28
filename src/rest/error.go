package rest

const (
	RequestValidationFailed = "request.validation.failed"
	RequestBindFailed       = "request.bind.failed"
)

type (
	// ErrorResponse is the generic response for all failed API-calls.
	ErrorResponse struct {
		// Generic code to detect kind of error
		// @Required
		ErrorCode string `json:"errorCode"`

		// Human-readable description of the error
		// @Required
		Description string `json:"description"`

		// Additional data to describe the error
		AdditionalAttributes map[string]interface{} `json:"additionalAttributes,omitempty"`
	}
)
