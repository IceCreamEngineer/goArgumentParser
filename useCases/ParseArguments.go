package useCases

import (
	"errors"
	"goArgumentParser/entities"
	"goArgumentParser/ports"
	"iter"
	"slices"
	"strings"
	"unicode"
)

var (
	currentArgument iter.Seq[any]
	marshalers      map[*entities.ArgumentNames]ports.ArgumentMarshaler
	argumentsFound  entities.ArgumentSet
)

type ArgumentParser struct {
	Schema               []entities.ArgumentSchemaElement
	Arguments            []string
	MarshalerFactory     ports.ArgumentMarshalerFactory
	HelpMessagePresenter HelpMessagePresenter
}

func (a *ArgumentParser) Parse() error {
	parseError := a.tryToParse()
	if parseError != nil {
		if errors.Is(parseError, &PresentHelp{}) {
			a.HelpMessagePresenter.PresentHelpMessage(a.Schema)
			return nil
		}
		return parseError
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
	missingRequiredArgumentError := a.checkForRequiredArguments()
	if missingRequiredArgumentError != nil {
		return missingRequiredArgumentError
	}
	return nil
}

func (a *ArgumentParser) parseSchema() error {
	marshalers = make(map[*entities.ArgumentNames]ports.ArgumentMarshaler)
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
	for _, character := range elementName {
		if !unicode.IsLetter(character) {
			return &entities.ArgumentError{ErrorCode: entities.InvalidArgumentName, ErrorArgumentId: elementName}
		}
	}
	return nil
}

func (a *ArgumentParser) setMarshalerFor(schemaElement *entities.ArgumentSchemaElement) error {
	var marshaler ports.ArgumentMarshaler
	if slices.Contains(a.MarshalerFactory.ArgumentTypes(), schemaElement.ArgumentType) {
		marshaler = a.MarshalerFactory.CreateFrom(schemaElement.ArgumentType)
		marshalers[&entities.ArgumentNames{schemaElement.Name, schemaElement.LongName}] = marshaler
		return nil
	}
	return &entities.ArgumentError{ErrorCode: entities.InvalidArgumentFormat, ErrorArgumentId: schemaElement.Name}
}

func (a *ArgumentParser) parseArguments() error {
	matchingNames, argument := a.initializeArgumentTracking()
	next, stop := iter.Pull(currentArgument)
	defer stop()
	hasMoreArguments := true
	for hasMoreArguments {
		hasMoreArguments, argument = a.iterateArgument(next)
		if a.isArgumentValue(argument) {
			continue
		}
		argumentParseError := a.parseArgument(matchingNames, argument, next)
		if argumentParseError != nil {
			return argumentParseError
		}
	}
	return nil
}

func (a *ArgumentParser) initializeArgumentTracking() (*entities.ArgumentNames, string) {
	argumentsFound = entities.NewArgumentSet()
	var matchingNames *entities.ArgumentNames
	currentArgument = createIteratorFrom(a.Arguments)
	var argument string
	return matchingNames, argument
}

func createIteratorFrom(s []string) iter.Seq[any] {
	return func(yield func(any) bool) {
		for _, v := range s {
			if !yield(v) {
				return
			}
		}
	}
}

func (a *ArgumentParser) iterateArgument(next func() (any, bool)) (bool, string) {
	nextArgument, ok := next()
	if !ok {
		return false, ""
	}
	argument := nextArgument.(string)
	return ok, argument
}

func (a *ArgumentParser) isArgumentValue(argument string) bool {
	return !strings.HasPrefix(argument, "--") && !strings.HasPrefix(argument, "-")
}

func (a *ArgumentParser) parseArgument(matchingNames *entities.ArgumentNames, argument string, next func() (any, bool)) error {
	matchingNames = a.matchingNamesFor(argument)
	expectedArgumentError := a.checkForExpected(argument, matchingNames)
	if expectedArgumentError != nil {
		return expectedArgumentError
	}
	argumentsFound.Add(a.parseArgumentNameFrom(argument))
	marshaler := marshalers[matchingNames]
	marshalError := marshaler.Set(next)
	if marshalError != nil {
		return marshalError
	}
	return nil
}

func (a *ArgumentParser) matchingNamesFor(argument string) *entities.ArgumentNames {
	for names := range marshalers {
		if strings.Split(argument, "-")[1] == names.Name || a.isLongName(argument, names) {
			return names
		}
	}
	return nil
}

func (a *ArgumentParser) checkForExpected(argument string, matchingNames *entities.ArgumentNames) error {
	if argument == "-h" || argument == "--help" {
		return &PresentHelp{}
	}
	if matchingNames == nil {
		return &entities.ArgumentError{ErrorCode: entities.UnexpectedArgument,
			ErrorArgumentId: strings.Split(a.Arguments[0], "-")[1]}
	}
	return nil
}

func (a *ArgumentParser) isLongName(argument string, names *entities.ArgumentNames) bool {
	return strings.HasPrefix(argument, "--") && strings.Split(argument, "--")[1] == names.LongName
}

func (a *ArgumentParser) parseArgumentNameFrom(argument string) string {
	var argumentName string
	if strings.HasPrefix(argument, "--") {
		argumentName = strings.Split(argument, "--")[1]
	} else {
		argumentName = strings.Split(argument, "-")[1]
	}
	return argumentName
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
	return argumentsFound.Has(argument)
}

func (a *ArgumentParser) NextArgument() int {
	if argumentsFound.Length() == 0 {
		return 0
	}
	return argumentsFound.Length() - 1
}

func (a *ArgumentParser) GetValueOf(names entities.ArgumentNames) any {
	for _names, marshaler := range marshalers {
		if _names.Name == names.Name || _names.LongName == names.LongName {
			return marshaler.GetValue()
		}
	}
	return nil
}

type PresentHelp struct{}

func (p PresentHelp) Error() string {
	return "help"
}
