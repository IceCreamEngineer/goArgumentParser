package useCases

import (
	"goArgumentParser/entities"
)

type ArgumentParser struct {
	Schema, Arguments string
}

func (a ArgumentParser) Parse() entities.ArgumentError {
	return entities.ArgumentError{ErrorCode: 1}
}

func (a ArgumentParser) NextArgument() int {
	return 0
}
