package errors

// InvalidBodyCode ...
const InvalidBodyCode = "InvalidBody"

type invalidBodyError struct {
	error
}

func (e invalidBodyError) Cause() error {
	return e.error
}

func (e invalidBodyError) Code() string {
	return InvalidBodyCode
}

func (e invalidBodyError) Message() string {
	return "Invalid Body"
}

// IsInvalidBody ...
func IsInvalidBody(err error) bool {
	_, ok := err.(invalidBodyError)
	return ok
}

// InvalidBody ...
func InvalidBody(err error) error {
	if err == nil {
		return nil
	}

	return invalidBodyError{err}
}
