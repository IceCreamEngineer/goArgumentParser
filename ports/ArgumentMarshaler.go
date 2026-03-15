package ports

type ArgumentMarshaler interface {
	Set(nextArgument func() (any, bool)) error
	GetValue() any
}
