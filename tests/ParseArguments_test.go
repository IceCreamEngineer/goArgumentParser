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
	argumentParser := useCases.ArgumentParser{}
	if argumentParser.NextArgument() != 0 {
		t.Error("Should return 0")
	}
}

func TestNoSchemaButOneArgument(t *testing.T) {
	argumentParser := useCases.ArgumentParser{Arguments: []string{"-x"}}
	err := argumentParser.Parse()
	assertCorrectArgumentError(t, err, entities.UnexpectedArgument, "x")
}

func TestNoSchemaButMultipleArguments(t *testing.T) {
	argumentParser := useCases.ArgumentParser{Arguments: []string{"-x", "-y"}}
	err := argumentParser.Parse()
	assertCorrectArgumentError(t, err, entities.UnexpectedArgument, "x")
}

func TestNonLetterSchema(t *testing.T) {
	argumentParser := useCases.ArgumentParser{Schema: []entities.ArgumentSchemaElement{{Name: "*"}}}
	err := argumentParser.Parse()
	assertCorrectArgumentError(t, err, entities.InvalidArgumentName, "*")
}

func TestNonLetterSchemaLongName(t *testing.T) {
	argumentParser := useCases.ArgumentParser{Schema: []entities.ArgumentSchemaElement{{Name: "x", LongName: "**"}}}
	err := argumentParser.Parse()
	assertCorrectArgumentError(t, err, entities.InvalidArgumentName, "**")
}

func TestInvalidArgumentFormat(t *testing.T) {
	setup()
	argumentParser := useCases.ArgumentParser{Schema: []entities.ArgumentSchemaElement{{Name: "f", ArgumentType: "~"}},
		MarshalerFactory: argumentMarshalerFactory}
	err := argumentParser.Parse()
	assertCorrectArgumentError(t, err, entities.InvalidArgumentFormat, "f")
}

func TestMissingRequiredArgumentForNoArguments(t *testing.T) {
	argumentParser := useCases.ArgumentParser{Schema: []entities.ArgumentSchemaElement{{Name: "x", IsRequired: true}},
		MarshalerFactory: argumentMarshalerFactory}
	err := argumentParser.Parse()
	assertCorrectArgumentError(t, err, entities.MissingRequiredArgument, "")
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
