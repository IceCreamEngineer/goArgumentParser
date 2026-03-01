package adapters

import "goArgumentParser/ports"

type NoArgumentMarshaler struct{}

func (m NoArgumentMarshaler) Set(currentArgument string) {}

func (m NoArgumentMarshaler) GetValueFrom(marshaler ports.ArgumentMarshaler) any {
	return nil
}
