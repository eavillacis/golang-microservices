package errors

// FileTooBigCode text code
const FileTooBigCode = "FileTooBig"

type errFileTooBig struct {
	error
}

func (e errFileTooBig) Cause() error {
	return e.error
}

func (e errFileTooBig) Code() string {
	return FileTooBigCode
}

func (e errFileTooBig) Message() string {
	return "File Too Big"
}

// IsFileTooBig ...
func IsFileTooBig(err error) bool {
	_, ok := err.(errFileTooBig)
	return ok
}

// FileTooBig ...
func FileTooBig(err error) error {
	if err == nil {
		return nil
	}

	return errFileTooBig{err}
}
