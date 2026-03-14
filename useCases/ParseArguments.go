package useCases

import (
	"goArgumentParser/entities"
	"goArgumentParser/ports"
	"iter"
	"slices"
	"strings"
	"unicode"
)

type Names struct {
	Name, LongName string
}

type void struct{}

var (
	currentArgument iter.Seq[any]
	marshalers      map[*Names]ports.ArgumentMarshaler
	argumentsFound  map[string]void
	entry           void
)

type ArgumentParser struct {
	Schema           []entities.ArgumentSchemaElement
	Arguments        []string
	MarshalerFactory ports.ArgumentMarshalerFactory
}

func (a *ArgumentParser) Parse() error {
	parseError := a.tryToParse()
	if parseError != nil {
		return parseError
	}
	missingRequiredArgumentError := a.checkForRequiredArguments()
	if missingRequiredArgumentError != nil {
		return missingRequiredArgumentError
	}
	return nil
}

func (a *ArgumentParser) tryToParse() error {
	schemaParsingError := a.parseSchema()
	if schemaParsingError != nil {
		return schemaParsingError
	}
	argumentParsingError := a.parseArguments()
	if argumentParsingError != nil {
		return argumentParsingError
	}
	return nil
}

func (a *ArgumentParser) parseSchema() error {
	marshalers = make(map[*Names]ports.ArgumentMarshaler)
	for _, schemaElement := range a.Schema {
		schemaElementParseError := a.parseSchemaElement(&schemaElement)
		if schemaElementParseError != nil {
			return schemaElementParseError
		}
	}
	return nil
}

func (a *ArgumentParser) parseSchemaElement(schemaElement *entities.ArgumentSchemaElement) error {
	validationError := validate(*schemaElement)
	if validationError != nil {
		return validationError
	}
	marshalerError := a.setMarshalerFor(schemaElement)
	if marshalerError != nil {
		return marshalerError
	}
	return nil
}

func (a *ArgumentParser) setMarshalerFor(schemaElement *entities.ArgumentSchemaElement) error {
	var marshaler ports.ArgumentMarshaler
	if slices.Contains(a.MarshalerFactory.ArgumentTypes(), schemaElement.ArgumentType) {
		marshaler = a.MarshalerFactory.CreateFrom(schemaElement.ArgumentType)
		marshalers[&Names{schemaElement.Name, schemaElement.LongName}] = marshaler
		return nil
	}
	return &entities.ArgumentError{ErrorCode: entities.InvalidArgumentFormat, ErrorArgumentId: schemaElement.Name}
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

func (a *ArgumentParser) parseArguments() error {
	unexpectedArgumentError := a.checkForUnexpectedArgument()
	if unexpectedArgumentError != nil {
		return unexpectedArgumentError
	}
	argumentsFound = make(map[string]void)
	var matchingNames *Names = nil
	currentArgument = seq(a.Arguments)
	next, stop := iter.Pull(currentArgument)
	defer stop()
	ok := true
	var argument string
	for ok {
		ok, argument = a.nextArgument(next)

		if strings.HasPrefix(argument, "--") {
			argumentName := strings.Split(argument, "--")[1]
			argumentsFound[argumentName] = entry
		} else if strings.HasPrefix(argument, "-") {
			argumentName := strings.Split(argument, "-")[1]
			argumentsFound[argumentName] = entry
		} else {
			continue
		}

		if strings.HasPrefix(argument, "--") || strings.HasPrefix(argument, "-") {
			matchingNames = a.matchingNamesFor(argument)
		}

		if strings.HasPrefix(argument, "-") && matchingNames != nil {
			marshalError := marshalers[matchingNames].Set(next)
			if marshalError != nil {
				return marshalError
			}
		}
	}
	return nil
}

func (a *ArgumentParser) nextArgument(next func() (any, bool)) (bool, string) {
	nextArgument, ok := next()
	if !ok {
		return false, ""
	}
	argument := nextArgument.(string)
	return ok, argument
}

func seq(s []string) iter.Seq[any] {
	return func(yield func(any) bool) {
		for _, v := range s {
			if !yield(v) {
				return
			}
		}
	}
}

func (a *ArgumentParser) isArgumentValue(argument string) bool {
	return !strings.HasPrefix(argument, "--") && !strings.HasPrefix(argument, "-")
}

func (a *ArgumentParser) matchingNamesFor(argument string) *Names {
	for names := range marshalers {
		if strings.Split(argument, "-")[1] == names.Name || a.isLongName(argument, names) {
			return names
		}
	}
	return nil
}

func (a *ArgumentParser) isLongName(argument string, names *Names) bool {
	return strings.HasPrefix(argument, "--") && strings.Split(argument, "--")[1] == names.LongName
}

func (a *ArgumentParser) checkForUnexpectedArgument() error {
	if len(a.Schema) == 0 && len(a.Arguments) != 0 {
		return &entities.ArgumentError{ErrorCode: entities.UnexpectedArgument,
			ErrorArgumentId: strings.Split(a.Arguments[0], "-")[1]}
	}
	return nil
}

func (a *ArgumentParser) checkForRequiredArguments() error {
	for _, element := range a.Schema {
		if !a.Has(element.Name) && !a.Has(element.LongName) && element.IsRequired() {
			return &entities.ArgumentError{ErrorCode: entities.MissingRequiredArgument}
		}
	}
	return nil
}

func (a *ArgumentParser) Has(argument string) bool {
	_, found := argumentsFound[argument]
	return found
}

func (a *ArgumentParser) NextArgument() int {
	if len(argumentsFound) == 0 {
		return 0
	}
	return len(argumentsFound) - 1
}
