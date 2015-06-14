package common

type StrError struct {
	Str string
}

func (e *StrError) Error() string {
	return e.Str
}
