package tests

import (
	"goArgumentParser/entities"
	"goArgumentParser/useCases"
	"testing"
)

func TestNoSchemaOrArguments(t *testing.T) {
	argumentParser := useCases.ArgumentParser{}
	if argumentParser.NextArgument() != 0 {
		t.Error("Should return 0")
	}
}

func TestNoSchemaButOneArgument(t *testing.T) {
	argumentParser := useCases.ArgumentParser{Arguments: []string{"-x"}}
	err := argumentParser.Parse()
	assertThatThereWasAnError(t, err)
	assertCorrectErrorCode(t, err, entities.UnexpectedArgument)
	assertCorrectErrorArgumentId(t, err, "x")
}

func TestNoSchemaButMultipleArguments(t *testing.T) {
	argumentParser := useCases.ArgumentParser{Arguments: []string{"-x", "-y"}}
	err := argumentParser.Parse()
	assertThatThereWasAnError(t, err)
	assertCorrectErrorCode(t, err, entities.UnexpectedArgument)
	assertCorrectErrorArgumentId(t, err, "x")
}

func TestNonLetterSchema(t *testing.T) {
	argumentParser := useCases.ArgumentParser{Schema: []entities.ArgumentSchemaElement{entities.NewArgumentSchemaElement(Name: "*")}}
	err := argumentParser.Parse()
	assertThatThereWasAnError(t, err)
	assertCorrectErrorCode(t, err, entities.InvalidArgumentName)
}

func assertThatThereWasAnError(t *testing.T, err error) {
	if err == nil {
		t.Error("Should return err")
	}
}

func assertCorrectErrorCode(t *testing.T, err entities.ArgumentError, errorCode int) {
	if err.ErrorCode != errorCode {
		t.Error("Should return error code ", errorCode)
	}
}

func assertCorrectErrorArgumentId(t *testing.T, err entities.ArgumentError, errorArgumentId string) {
	if err.ErrorArgumentId != errorArgumentId {
		t.Error("Should return error argument id ", errorArgumentId)
	}
}
