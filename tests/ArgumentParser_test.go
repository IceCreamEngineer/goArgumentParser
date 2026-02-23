package tests

import (
	"goArgumentParser/useCases"
	"testing"
)

func TestNoSchemaOrArguments(t *testing.T) {
	argumentParser := useCases.ArgumentParser{}
	if argumentParser.NextArgument() != 0 {
		t.Error("Should return 0")
	}
}
