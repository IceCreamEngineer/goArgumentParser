package adapters

import (
	"goArgumentParser/ports"
	"iter"
)

type NoArgumentMarshaler struct{}

func (m NoArgumentMarshaler) Set(currentArgument iter.Seq[any]) error {
	return nil
}

func (m NoArgumentMarshaler) GetValueFrom(marshaler ports.ArgumentMarshaler) any {
	return nil
}
