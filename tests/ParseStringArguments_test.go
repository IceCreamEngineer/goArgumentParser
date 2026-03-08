package tests

import (
	"goArgumentParser/adapters"
	"goArgumentParser/entities"
	"goArgumentParser/ports"
	"goArgumentParser/useCases"
	"testing"
)

var _argumentMarshalerFactory ports.ArgumentMarshalerFactory

func _setup() {
	_argumentMarshalerFactory = adapters.StringsArgumentMarshalerFactory{}
}

func TestSimpleStringPresent(t *testing.T) {
	_setup()
	argumentParser := useCases.ArgumentParser{Schema: []entities.ArgumentSchemaElement{{Name: "x",
		ArgumentType: "*"}}, Arguments: []string{"-x", "param"}, MarshalerFactory: _argumentMarshalerFactory}
	parseError := argumentParser.Parse()
	if parseError != nil {
		t.Error(parseError)
	}
	if !argumentParser.Has("x") {
		t.Error("Expected x")
	}
}
