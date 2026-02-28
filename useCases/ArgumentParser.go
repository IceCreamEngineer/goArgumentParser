package useCases

import (
	"goArgumentParser/entities"
	"strings"
	"unicode"
)

type ArgumentParser struct {
	Schema    []entities.ArgumentSchemaElement
	Arguments []string
}

func (a ArgumentParser) Parse() error {
	err := a.validateSchema()
	if err != nil {
		return err
	}
	return &entities.ArgumentError{ErrorCode: entities.UnexpectedArgument,
		ErrorArgumentId: strings.Split(a.Arguments[0], "-")[1]}
}

func (a ArgumentParser) validateSchema() error {
	for _, schemaElement := range a.Schema {
		err := validateElement(schemaElement)
		if err != nil {
			return err
		}
	}
	return nil
}

func validateElement(schemaElement entities.ArgumentSchemaElement) error {
	for _, r := range schemaElement.Name {
		if !unicode.IsLetter(r) {
			return &entities.ArgumentError{ErrorCode: entities.InvalidArgumentName, ErrorArgumentId: schemaElement.Name}
		}
	}
	return nil
}

func (a ArgumentParser) NextArgument() int {
	return 0
}
