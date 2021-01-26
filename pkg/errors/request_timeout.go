package errors

// RequestTimeoutCode ...
const RequestTimeoutCode = "RequestTimeout"

type errRequestTimeout struct {
	error
}

func (e errRequestTimeout) Cause() error {
	return e.error
}

func (e errRequestTimeout) Code() string {
	return RequestTimeoutCode
}

func (e errRequestTimeout) Message() string {
	return "Request Timeout"
}

// IsRequestTimeout ...
func IsRequestTimeout(err error) bool {
	_, ok := err.(errRequestTimeout)
	return ok
}

// RequestTimeout ...
func RequestTimeout(err error) error {
	if err == nil {
		return nil
	}

	return errRequestTimeout{err}
}
