package tests

import (
	"errors"
	"goArgumentParser/adapters"
	"goArgumentParser/entities"
	"goArgumentParser/ports"
	"goArgumentParser/useCases"
	"testing"
)

var argumentMarshalerFactory ports.ArgumentMarshalerFactory

func setup() {
	argumentMarshalerFactory = adapters.NoArgumentMarshalerFactory{}
}

func TestNoSchemaOrArguments(t *testing.T) {
	argumentParser := &useCases.ArgumentParser{}
	assertNoNextArgument(t, argumentParser)
}

func TestNoSchemaButOneArgument(t *testing.T) {
	argumentParser := &useCases.ArgumentParser{Arguments: []string{"-x"}}
	assertCorrectArgumentError(t, argumentParser.Parse(), entities.UnexpectedArgument, "x")
}

func TestNoSchemaButMultipleArguments(t *testing.T) {
	argumentParser := &useCases.ArgumentParser{Arguments: []string{"-x", "-y"}}
	assertCorrectArgumentError(t, argumentParser.Parse(), entities.UnexpectedArgument, "x")
}

func TestNonLetterSchema(t *testing.T) {
	argumentParser := &useCases.ArgumentParser{Schema: []entities.ArgumentSchemaElement{{Name: "*"}}}
	assertCorrectArgumentError(t, argumentParser.Parse(), entities.InvalidArgumentName, "*")
}

func TestNonLetterSchemaLongName(t *testing.T) {
	argumentParser := &useCases.ArgumentParser{Schema: []entities.ArgumentSchemaElement{{Name: "x", LongName: "**"}}}
	assertCorrectArgumentError(t, argumentParser.Parse(), entities.InvalidArgumentName, "**")
}

func TestInvalidArgumentFormat(t *testing.T) {
	setup()
	argumentParser := &useCases.ArgumentParser{Schema: []entities.ArgumentSchemaElement{{Name: "f", ArgumentType: "~"}},
		MarshalerFactory: argumentMarshalerFactory}
	assertCorrectArgumentError(t, argumentParser.Parse(), entities.InvalidArgumentFormat, "f")
}

func TestMissingRequiredArgumentForNoArguments(t *testing.T) {
	argumentParser := &useCases.ArgumentParser{Schema: []entities.ArgumentSchemaElement{{Name: "x"}},
		MarshalerFactory: argumentMarshalerFactory}
	assertCorrectArgumentError(t, argumentParser.Parse(), entities.MissingRequiredArgument, "")
}

func TestMissingRequiredArgumentForSomeArgument(t *testing.T) {
	argumentParser := &useCases.ArgumentParser{Schema: []entities.ArgumentSchemaElement{{Name: "x"},
		{Name: "y"}}, Arguments: []string{"-x"}, MarshalerFactory: argumentMarshalerFactory}
	assertCorrectArgumentError(t, argumentParser.Parse(), entities.MissingRequiredArgument, "")
}

func TestMissingOptionalArgumentForNoArguments(t *testing.T) {
	required := false
	argumentParser := &useCases.ArgumentParser{Schema: []entities.ArgumentSchemaElement{{Name: "x",
		Required: &required}}, MarshalerFactory: argumentMarshalerFactory}
	assertThatThereWasNoError(t, argumentParser.Parse())
}

func TestMissingOptionalArgumentForSomeArgument(t *testing.T) {
	required := false
	argumentParser := &useCases.ArgumentParser{Schema: []entities.ArgumentSchemaElement{{Name: "x"}, {Name: "y",
		Required: &required}}, Arguments: []string{"-x"}, MarshalerFactory: argumentMarshalerFactory}
	assertThatThereWasNoError(t, argumentParser.Parse())
}

func TestExtraArgumentsThatLookLikeFlags(t *testing.T) {
	required := false
	argumentParser := &useCases.ArgumentParser{Schema: []entities.ArgumentSchemaElement{{Name: "x"}, {Name: "y",
		Required: &required}}, Arguments: []string{"-x", "alpha", "-y", "alpha"},
		MarshalerFactory: argumentMarshalerFactory}
	parseError := argumentParser.Parse()
	assertThatThereWasNoError(t, parseError)
	assertParsed(t, argumentParser, "x")
	assertParsed(t, argumentParser, "y")
	assertNoNextArgument(t, argumentParser)
}

func assertCorrectArgumentError(t *testing.T, err error, errorCode int, errorArgumentId string) {
	assertThatThereWasAnError(t, err)
	var aErr *entities.ArgumentError
	errors.As(err, &aErr)
	assertExpectedErrorField(t, aErr.ErrorCode, errorCode)
	assertExpectedErrorField(t, aErr.ErrorArgumentId, errorArgumentId)
}

func assertThatThereWasAnError(t *testing.T, err error) {
	if err == nil {
		t.Error("Should return err")
	}
}

func assertExpectedErrorField(t *testing.T, actualErrorField any, expectedErrorField any) {
	if actualErrorField != expectedErrorField {
		t.Error("Expected ", expectedErrorField, " but got ", actualErrorField)
	}
}

func assertThatThereWasNoError(t *testing.T, err error) {
	if err != nil {
		t.Error("Should not return an error")
	}
}

func assertParsed(t *testing.T, argumentParser *useCases.ArgumentParser, argument string) {
	if !argumentParser.Has(argument) {
		t.Error("Should have parsed ", argument)
	}
}

func assertNoNextArgument(t *testing.T, argumentParser *useCases.ArgumentParser) {
	if argumentParser.NextArgument() != 0 {
		t.Error("Should return 0")
	}
}
