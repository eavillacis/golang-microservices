package errors

// NotFoundCode ...
const NotFoundCode = "NotFound"

type notFoundError struct{ error }

func (e notFoundError) Cause() error {
	return e.error
}

func (e notFoundError) Code() string {
	return NotFoundCode
}

func (e notFoundError) Message() string {
	return "Resource not found"
}

// IsNotFound ...
func IsNotFound(err error) bool {
	_, ok := err.(notFoundError)
	return ok
}

// NotFound ...
func NotFound(err error) error {
	if err == nil {
		return nil
	}

	return notFoundError{err}
}
