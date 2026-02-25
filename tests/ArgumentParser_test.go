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
	argumentParser := useCases.ArgumentParser{Arguments: "-x"}
	err := argumentParser.Parse()
	assertThatThereWasAnError(t, err)
	assertCorrectErrorCode(t, err, 1)
}

func assertCorrectErrorCode(t *testing.T, err entities.ArgumentError, errorCode int) {
	if err.ErrorCode != errorCode {
		t.Error("Should return error code ", errorCode)
	}
}

func assertThatThereWasAnError(t *testing.T, err error) {
	if err == nil {
		t.Error("Should return err")
	}
}
