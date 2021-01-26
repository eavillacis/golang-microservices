package errors

// ForbiddenCode text code
const ForbiddenCode = "Forbidden"

type errForbidden struct {
	error
}

func (e errForbidden) Cause() error {
	return e.error
}

func (e errForbidden) Code() string {
	return ForbiddenCode
}

func (e errForbidden) Message() string {
	return "Forbidden"
}

// IsForbidden ...
func IsForbidden(err error) bool {
	_, ok := err.(errForbidden)
	return ok
}

// Forbidden ...
func Forbidden(err error) error {
	if err == nil {
		return nil
	}

	return errForbidden{err}
}
