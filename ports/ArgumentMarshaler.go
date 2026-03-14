package ports

type ArgumentMarshaler interface {
	Set(nextArgument func() (any, bool)) error
	GetValueFrom(marshaler ArgumentMarshaler) any
}
