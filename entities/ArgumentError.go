package entities

const UnexpectedArgument = 1

type ArgumentError struct {
	ErrorCode       int
	ErrorArgumentId string
	ErrorMessage    string
}

func (err ArgumentError) Error() string {
	return err.ErrorMessage
}
