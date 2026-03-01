package adapters

import (
	"goArgumentParser/ports"
)

type NoArgumentMarshalerFactory struct{}

func (a NoArgumentMarshalerFactory) ArgumentTypes() []string {
	return []string{}
}

func (a NoArgumentMarshalerFactory) CreateFrom(argumentType string) ports.ArgumentMarshaler {
	return NoArgumentMarshaler{}
}
