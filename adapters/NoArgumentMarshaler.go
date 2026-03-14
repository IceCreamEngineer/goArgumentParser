package adapters

import (
	"goArgumentParser/ports"
)

type NoArgumentMarshaler struct{}

func (m NoArgumentMarshaler) Set(nextArgument func() (any, bool)) error {
	return nil
}

func (m NoArgumentMarshaler) GetValueFrom(marshaler ports.ArgumentMarshaler) any {
	return nil
}
