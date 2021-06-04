package exceptions

type UnauthorizedAccessError struct {
	Msg string
}

func (e *UnauthorizedAccessError) Error() string { return e.Msg }

