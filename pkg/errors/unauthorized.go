package errors

// UnauthorizedCode ...
const UnauthorizedCode = "Unauthorized"

type unathorizedError struct{ error }

func (e unathorizedError) Cause() error {
	return e.error
}

func (e unathorizedError) Code() string {
	return UnauthorizedCode
}

func (e unathorizedError) Message() string {
	return "Unauthorized"
}

// IsUnauthorized ...
func IsUnauthorized(err error) bool {
	_, ok := err.(unathorizedError)
	return ok
}

// Unauthorized ...
func Unauthorized(err error) error {
	if err == nil {
		return nil
	}

	return unathorizedError{err}
}
