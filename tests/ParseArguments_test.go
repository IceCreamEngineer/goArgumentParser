package tests

import (
	"goArgumentParser/adapters"
	"goArgumentParser/entities"
	"goArgumentParser/ports"
	"goArgumentParser/useCases"
	"testing"
)

var noArgumentMarshalerFactory ports.ArgumentMarshalerFactory

func noSetup() {
	noArgumentMarshalerFactory = adapters.NoArgumentMarshalerFactory{}
}

func TestNoSchemaOrArguments(t *testing.T) {
	argumentParser := &useCases.ArgumentParser{}
	assertNoNextArgument(t, argumentParser)
}

func TestNoSchemaButOneArgument(t *testing.T) {
	argumentParser := &useCases.ArgumentParser{Arguments: []string{"-x"}}
	AssertCorrectArgumentError(t, argumentParser.Parse(), entities.UnexpectedArgument, "x")
}

func TestNoSchemaButMultipleArguments(t *testing.T) {
	argumentParser := &useCases.ArgumentParser{Arguments: []string{"-x", "-y"}}
	AssertCorrectArgumentError(t, argumentParser.Parse(), entities.UnexpectedArgument, "x")
}

func TestNonLetterSchema(t *testing.T) {
	argumentParser := &useCases.ArgumentParser{Schema: []entities.ArgumentSchemaElement{{Name: "*"}}}
	AssertCorrectArgumentError(t, argumentParser.Parse(), entities.InvalidArgumentName, "*")
}

func TestNonLetterSchemaLongName(t *testing.T) {
	argumentParser := &useCases.ArgumentParser{Schema: []entities.ArgumentSchemaElement{{Name: "x", LongName: "**"}}}
	AssertCorrectArgumentError(t, argumentParser.Parse(), entities.InvalidArgumentName, "**")
}

func TestInvalidArgumentFormat(t *testing.T) {
	noSetup()
	argumentParser := &useCases.ArgumentParser{Schema: []entities.ArgumentSchemaElement{{Name: "f", ArgumentType: "~"}},
		MarshalerFactory: noArgumentMarshalerFactory}
	AssertCorrectArgumentError(t, argumentParser.Parse(), entities.InvalidArgumentFormat, "f")
}

func TestMissingRequiredArgumentForNoArguments(t *testing.T) {
	argumentParser := &useCases.ArgumentParser{Schema: []entities.ArgumentSchemaElement{{Name: "x"}},
		MarshalerFactory: noArgumentMarshalerFactory}
	AssertCorrectArgumentError(t, argumentParser.Parse(), entities.MissingRequiredArgument, "")
}

func TestMissingRequiredArgumentForSomeArgument(t *testing.T) {
	argumentParser := &useCases.ArgumentParser{Schema: []entities.ArgumentSchemaElement{{Name: "x"},
		{Name: "y"}}, Arguments: []string{"-x"}, MarshalerFactory: noArgumentMarshalerFactory}
	AssertCorrectArgumentError(t, argumentParser.Parse(), entities.MissingRequiredArgument, "")
}

func TestMissingOptionalArgumentForNoArguments(t *testing.T) {
	required := false
	argumentParser := &useCases.ArgumentParser{Schema: []entities.ArgumentSchemaElement{{Name: "x",
		Required: &required}}, MarshalerFactory: noArgumentMarshalerFactory}
	AssertThatThereWasNoError(t, argumentParser.Parse())
}

func TestMissingOptionalArgumentForSomeArgument(t *testing.T) {
	required := false
	argumentParser := &useCases.ArgumentParser{Schema: []entities.ArgumentSchemaElement{{Name: "x"}, {Name: "y",
		Required: &required}}, Arguments: []string{"-x"}, MarshalerFactory: noArgumentMarshalerFactory}
	AssertThatThereWasNoError(t, argumentParser.Parse())
}

func TestExtraArgumentsThatLookLikeFlags(t *testing.T) {
	required := false
	argumentParser := &useCases.ArgumentParser{Schema: []entities.ArgumentSchemaElement{{Name: "x"}, {Name: "y",
		Required: &required}}, Arguments: []string{"-x", "alpha", "-y", "alpha"},
		MarshalerFactory: noArgumentMarshalerFactory}
	AssertThatThereWasNoError(t, argumentParser.Parse())
	AssertParsed(t, argumentParser, "x")
	AssertParsed(t, argumentParser, "y")
	assertNoNextArgument(t, argumentParser)
}

func assertNoNextArgument(t *testing.T, argumentParser *useCases.ArgumentParser) {
	if argumentParser.NextArgument() != 0 {
		t.Error("Should return 0")
	}
}
