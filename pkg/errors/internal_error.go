package errors

// InternalErrorCode text code
const InternalErrorCode = "InternalError"

type internalError struct {
	error
}

func (e internalError) Cause() error {
	return e.error
}

func (e internalError) Code() string {
	return InternalErrorCode
}

func (e internalError) Message() string {
	return "Internal Error"
}

// IsInternalError ...
func IsInternalError(err error) bool {
	_, ok := err.(internalError)
	return ok
}

// InternalError ...
func InternalError(err error) error {
	if err == nil {
		return nil
	}

	return internalError{err}
}
