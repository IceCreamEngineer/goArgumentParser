package tests

import (
	"goArgumentParser/adapters"
	"goArgumentParser/entities"
	"goArgumentParser/ports"
	"goArgumentParser/useCases"
	"testing"
)

var argumentMarshalerFactory ports.ArgumentMarshalerFactory

func setup() {
	argumentMarshalerFactory = adapters.StringsArgumentMarshalerFactory{}
}

func TestSimpleStringPresent(t *testing.T) {
	argumentParser := useCases.ArgumentParser{Schema: []entities.ArgumentSchemaElement{{Name: "x",
		ArgumentType: "*"}}, Arguments: []string{"-x", "param"}}
}
