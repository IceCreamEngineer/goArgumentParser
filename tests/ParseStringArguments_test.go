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

func TestLongStringName(t *testing.T) {
	stringsSetup()
	argumentParser := useCases.ArgumentParser{Schema: []entities.ArgumentSchemaElement{{Name: "x",
		LongName: "excelsior", ArgumentType: "*"}}, Arguments: []string{"--excelsior", "alpha"},
		MarshalerFactory: stringsArgumentMarshalerFactory}
	parseError := argumentParser.Parse()
	AssertThatThereWasNoError(t, parseError)
	AssertParsed(t, &argumentParser, "excelsior")
}

func TestMissingStringArgument(t *testing.T) {
	stringsSetup()
	argumentParser := useCases.ArgumentParser{Schema: []entities.ArgumentSchemaElement{{Name: "x",
		ArgumentType: "*"}}, Arguments: []string{"-x"}, MarshalerFactory: stringsArgumentMarshalerFactory}
	AssertCorrectArgumentError(t, argumentParser.Parse(), entities.MissingString, "")
}

func TestExtraArguments(t *testing.T) {
	stringsSetup()
	argumentParser := useCases.ArgumentParser{Schema: []entities.ArgumentSchemaElement{{Name: "y",
		ArgumentType: "*"}}, Arguments: []string{"-y", "alpha", "beta"}, MarshalerFactory: stringsArgumentMarshalerFactory}
	AssertThatThereWasNoError(t, argumentParser.Parse())
	AssertArgumentValue(t, argumentParser, useCases.Names{Name: "y"}, "alpha")
	AssertNextArgument(t, &argumentParser, 0)
}
