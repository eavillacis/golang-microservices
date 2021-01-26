package httputils

import "github.com/eavillacis/velociraptor/pkg/errors"

// Error ...
type Error struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Details string `json:"details"`
}

// ErrorResponse ...
func ErrorResponse(err error) *Error {
	apiError, ok := err.(errors.APIError)

	if ok {
		return &Error{
			Code:    apiError.Code(),
			Message: apiError.Message(),
			Details: err.Error(),
		}
	}

	return &Error{
		Code:    "0",
		Message: "Internal Server Error",
	}
}
