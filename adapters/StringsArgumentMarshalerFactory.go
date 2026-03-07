package adapters

import (
	"goArgumentParser/ports"
	"maps"
	"slices"
)

type StringsArgumentMarshalerFactory struct{}

var stringsMarshalers = map[string]ports.ArgumentMarshaler{
	"*": StringArgumentMarshaler{}, "[*]": StringArrayArgumentMarshaler{}, "": NoArgumentMarshaler{},
}

func (a StringsArgumentMarshalerFactory) ArgumentTypes() []string {
	return slices.Collect(maps.Keys(stringsMarshalers))
}
