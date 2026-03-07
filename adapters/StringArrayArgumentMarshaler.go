package adapters

import (
	"goArgumentParser/entities"
	"goArgumentParser/ports"
	"iter"
)

var strings []string

type StringArrayArgumentMarshaler struct{}

func (m StringArrayArgumentMarshaler) Set(currentArgument iter.Seq[any]) error {
	next, stop := iter.Pull(currentArgument)
	defer stop()
	aValue, ok := next()
	strings = append(strings, aValue.(string))
	if !ok {
		return &entities.ArgumentError{ErrorCode: entities.MissingString}
	}
	return nil
}

func (m StringArrayArgumentMarshaler) GetValueFrom(marshaler ports.ArgumentMarshaler) any {
	return strings
}
