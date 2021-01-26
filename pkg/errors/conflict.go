package errors

const conflictCode = "Conflict"

type conflictError struct {
	error
}

func (e conflictError) Cause() error {
	return e.error
}

func (e conflictError) Code() string {
	return conflictCode
}

func (e conflictError) Message() string {
	return "Conflict in action"
}

// IsConflict ...
func IsConflict(err error) bool {
	_, ok := err.(conflictError)
	return ok
}

// Conflict ...
func Conflict(err error) error {
	if err == nil {
		return nil
	}

	return conflictError{err}
}
