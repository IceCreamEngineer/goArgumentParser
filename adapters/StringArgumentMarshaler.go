package adapters

import (
	"goArgumentParser/entities"
	"goArgumentParser/ports"
	"iter"
)

var stringValue string

type StringArgumentMarshaler struct{}

func (m StringArgumentMarshaler) Set(currentArgument iter.Seq[any]) error {
	next, stop := iter.Pull(currentArgument)
	defer stop()
	aValue, ok := next()
	stringValue = aValue.(string)
	if !ok {
		return &entities.ArgumentError{ErrorCode: entities.MissingString}
	}
	return nil
}

func (m StringArgumentMarshaler) GetValueFrom(marshaler ports.ArgumentMarshaler) any {
	return stringValue
}
