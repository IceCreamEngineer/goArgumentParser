package useCases

import (
	"goArgumentParser/entities"
	"strings"
)

type ArgumentParser struct {
	Schema    []entities.ArgumentSchemaElement
	Arguments []string
}

func (a ArgumentParser) Parse() entities.ArgumentError {
	return entities.ArgumentError{ErrorCode: entities.UnexpectedArgument,
		ErrorArgumentId: strings.Split(a.Arguments[0], "-")[1]}
}

func (a ArgumentParser) NextArgument() int {
	return 0
}
