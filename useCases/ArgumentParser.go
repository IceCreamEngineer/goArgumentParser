package useCases

import (
	"goArgumentParser/entities"
	"goArgumentParser/ports"
	"strings"
	"unicode"
)

type ArgumentParser struct {
	Schema           []entities.ArgumentSchemaElement
	Arguments        []string
	MarshalerFactory ports.ArgumentMarshalerFactory
}

func (a ArgumentParser) Parse() error {
	err := a.parseSchema()
	if err != nil {
		return err
	}
	return &entities.ArgumentError{ErrorCode: entities.UnexpectedArgument,
		ErrorArgumentId: strings.Split(a.Arguments[0], "-")[1]}
}

func (a ArgumentParser) parseSchema() error {
	for _, schemaElement := range a.Schema {
		err := a.parseSchemaElement(schemaElement)
		if err != nil {
			return err
		}
	}
	return nil
}

func (a ArgumentParser) parseSchemaElement(schemaElement entities.ArgumentSchemaElement) error {
	validationError := validate(schemaElement)
	if validationError != nil {
		return validationError
	}
	return nil
}

func validate(schemaElement entities.ArgumentSchemaElement) error {
	isAlphaNumericName := isAlphaNumeric(schemaElement.Name)
	if isAlphaNumericName != nil {
		return isAlphaNumericName
	}
	isAlphaNumericLongName := isAlphaNumeric(schemaElement.LongName)
	if isAlphaNumericLongName != nil {
		return isAlphaNumericLongName
	}
	return nil
}

func isAlphaNumeric(elementName string) error {
	for _, r := range elementName {
		if !unicode.IsLetter(r) {
			return &entities.ArgumentError{ErrorCode: entities.InvalidArgumentName, ErrorArgumentId: elementName}
		}
	}
	return nil
}

func (a ArgumentParser) NextArgument() int {
	return 0
}
