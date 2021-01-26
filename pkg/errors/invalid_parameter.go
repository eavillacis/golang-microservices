package errors

// InvalidParameterCode ...
const InvalidParameterCode = "InvalidParameter"

type invalidParameterError struct {
	error
}

func (e invalidParameterError) Cause() error {
	return e.error
}

func (e invalidParameterError) Code() string {
	return InvalidParameterCode
}

func (e invalidParameterError) Message() string {
	return "Invalid Parameters"
}

// IsInvalidParameter ...
func IsInvalidParameter(err error) bool {
	_, ok := err.(invalidParameterError)
	return ok
}

// InvalidParameter ...
func InvalidParameter(err error) error {
	if err == nil {
		return nil
	}

	return invalidParameterError{err}
}
