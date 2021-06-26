package denied

type MessageRequestDeniedError struct {
	Msg string
}

func (e *MessageRequestDeniedError) Error() string { return e.Msg }

