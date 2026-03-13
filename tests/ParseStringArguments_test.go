package tests

import (
	"goArgumentParser/adapters"
	"goArgumentParser/entities"
	"goArgumentParser/ports"
	"goArgumentParser/useCases"
	"testing"
)

var stringsArgumentMarshalerFactory ports.ArgumentMarshalerFactory

func stringsSetup() {
	stringsArgumentMarshalerFactory = adapters.StringsArgumentMarshalerFactory{}
}

func TestSimpleStringPresent(t *testing.T) {
	stringsSetup()
	argumentParser := useCases.ArgumentParser{Schema: []entities.ArgumentSchemaElement{{Name: "x",
		ArgumentType: "*"}}, Arguments: []string{"-x", "param"}, MarshalerFactory: stringsArgumentMarshalerFactory}
	parseError := argumentParser.Parse()
	AssertThatThereWasNoError(t, parseError)
	AssertParsed(t, &argumentParser, "x")
}
