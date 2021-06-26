package expired

type StoryError struct {
	Msg string
}

func (e *StoryError) Error() string { return e.Msg }
