package ports

type ArgumentMarshaler interface {
	Set(currentArgument string)
	GetValueFrom(marshaler ArgumentMarshaler)
}
