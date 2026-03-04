package useCases

import (
	"goArgumentParser/entities"
	"goArgumentParser/ports"
	"slices"
	"strings"
	"unicode"
)

type void struct{}

var argumentsFound = make(map[string]void)

var entry void

type ArgumentParser struct {
	Schema           []entities.ArgumentSchemaElement
	Arguments        []string
	MarshalerFactory ports.ArgumentMarshalerFactory
}

func (a ArgumentParser) Parse() error {
	argumentsFound = make(map[string]void)
	schemaError := a.parseSchema()
	if schemaError != nil {
		return schemaError
	}
	a.parseArguments()
	missingRequiredArgumentError := a.checkForRequiredArguments()
	if missingRequiredArgumentError != nil {
		return missingRequiredArgumentError
	}
	return &entities.ArgumentError{ErrorCode: entities.UnexpectedArgument,
		ErrorArgumentId: strings.Split(a.Arguments[0], "-")[1]}
}

func (a ArgumentParser) parseArguments() {
	for _, argument := range a.Arguments {
		a.checkToParseArgumentName(argument, "-")
		a.checkToParseArgumentName(argument, "--")
	}
}

func (a ArgumentParser) checkToParseArgumentName(argument string, prefix string) {
	isArgumentName := strings.HasPrefix(argument, prefix)
	if isArgumentName {
		argumentName := strings.Split(argument, prefix)[len(prefix)]
		argumentsFound[argumentName] = entry
	}
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
	if !slices.Contains(a.MarshalerFactory.ArgumentTypes(), schemaElement.ArgumentType) {
		return &entities.ArgumentError{ErrorCode: entities.InvalidArgumentFormat, ErrorArgumentId: schemaElement.Name}
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

func (a ArgumentParser) checkForRequiredArguments() error {
	for _, element := range a.Schema {
		if !a.Has(element.Name) && !a.Has(element.LongName) && element.IsRequired {
			return &entities.ArgumentError{ErrorCode: entities.MissingRequiredArgument}
		}
	}
	return nil
}

func (a ArgumentParser) Has(argument string) bool {
	_, found := argumentsFound[argument]
	return found
}

func (a ArgumentParser) NextArgument() int {
	return 0
}
