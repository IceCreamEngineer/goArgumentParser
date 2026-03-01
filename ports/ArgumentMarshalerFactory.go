package ports

type ArgumentMarshalerFactory interface {
	ArgumentTypes() []string
	CreateFrom(argumentType string) ArgumentMarshaler
}
