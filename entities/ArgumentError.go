package entities

const UnexpectedArgument = 1
const InvalidArgumentName = 2
const InvalidArgumentFormat = 3
const MissingRequiredArgument = 4

type ArgumentError struct {
	ErrorCode       int
	ErrorArgumentId string
	ErrorMessage    string
}

func (err ArgumentError) Error() string {
	return err.ErrorMessage
}
