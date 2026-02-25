package entities

type ArgumentError struct {
	ErrorCode    int
	ErrorMessage string
}

func (err ArgumentError) Error() string {
	return err.ErrorMessage
}
