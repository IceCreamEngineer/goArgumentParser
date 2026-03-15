package tests

import (
	"errors"
	"goArgumentParser/entities"
	"goArgumentParser/useCases"
	"testing"
)

func AssertCorrectArgumentError(t *testing.T, err error, errorCode int, errorArgumentId string) {
	AssertThatThereWasAnError(t, err)
	var aErr *entities.ArgumentError
	errors.As(err, &aErr)
	assertExpectedErrorField(t, aErr.ErrorCode, errorCode)
	assertExpectedErrorField(t, aErr.ErrorArgumentId, errorArgumentId)
}

func AssertThatThereWasAnError(t *testing.T, err error) {
	if err == nil {
		t.Error("Should return err")
	}
}

func assertExpectedErrorField(t *testing.T, actualErrorField any, expectedErrorField any) {
	if actualErrorField != expectedErrorField {
		t.Error("Expected", expectedErrorField, " but got", actualErrorField)
	}
}

func AssertThatThereWasNoError(t *testing.T, err error) {
	if err != nil {
		t.Error("Should not return an error")
	}
}

func AssertParsed(t *testing.T, argumentParser *useCases.ArgumentParser, argument string) {
	if !argumentParser.Has(argument) {
		t.Error("Should have parsed", argument)
	}
}

func AssertArgumentValue(t *testing.T, argumentParser useCases.ArgumentParser, argumentNames entities.ArgumentNames,
	argumentValue string) {
	if argumentParser.GetValueOf(argumentNames) != argumentValue {
		t.Error("Should have parsed", argumentValue)
	}
}

func AssertNextArgument(t *testing.T, argumentParser *useCases.ArgumentParser, nextArgument int) {
	if argumentParser.NextArgument() != nextArgument {
		t.Error("Should return", nextArgument)
	}
}
