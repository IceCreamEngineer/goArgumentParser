package useCases

import (
	"goArgumentParser/entities"
	"goArgumentParser/ports"
	"slices"
	"strings"
)

const IndentLength = 2
const LineLength = 72
const PaddingPlaceholder = "|"

var indent = strings.Repeat(" ", IndentLength)

type HelpMessagePresenter struct {
	ProgramFileName string
	Description     string
	Presenter       ports.Presenter
}

func (h HelpMessagePresenter) PresentHelpMessage(schema []entities.ArgumentSchemaElement) {
	programTitle, argumentSpecificationMessages := h.buildSchemaDependentMessages(schema)
	h.Presenter.Present("" +
		programTitle +
		"\n" +
		h.textWrap(h.Description) + "\n" +
		"\n" +
		"optional arguments:\n" +
		argumentSpecificationMessages)
}

func (h HelpMessagePresenter) buildSchemaDependentMessages(schema []entities.ArgumentSchemaElement) (string, string) {
	programTitle := "usage: " + h.ProgramFileName + " [-h]"
	argumentSpecifications := []string{"-h, --help|show this help message and exit\n"}
	argumentSpecificationLabelLengths := []int{len("-h, --help")}
	for _, schemaElement := range schema {
		programTitle += h.buildProgramTitleFrom(schemaElement)
		argumentSpecification := h.buildArgumentSpecificationFrom(schemaElement)
		argumentSpecifications = append(argumentSpecifications, argumentSpecification)
		argumentSpecificationLabelLengths = append(argumentSpecificationLabelLengths,
			len(h.argumentSpecificationLabel(argumentSpecification)))
	}
	return programTitle + "\n", h.buildArgumentSpecificationMessages(argumentSpecificationLabelLengths, argumentSpecifications)
}

func (h HelpMessagePresenter) buildProgramTitleFrom(schemaElement entities.ArgumentSchemaElement) string {
	if schemaElement.IsRequired() {
		return " -" + schemaElement.Name
	}
	return " [-" + schemaElement.Name + "]"
}

func (h HelpMessagePresenter) buildArgumentSpecificationFrom(schemaElement entities.ArgumentSchemaElement) string {
	longNameAddition := h.checkToAddLongName(schemaElement)
	return "-" + schemaElement.Name + longNameAddition + PaddingPlaceholder + schemaElement.Description + "\n"
}

func (h HelpMessagePresenter) checkToAddLongName(schemaElement entities.ArgumentSchemaElement) string {
	if schemaElement.LongName != "" {
		return ", --" + schemaElement.LongName
	}
	return ""
}

func (h HelpMessagePresenter) argumentSpecificationLabel(argumentSpecification string) string {
	return strings.Split(argumentSpecification, PaddingPlaceholder)[0]
}

func (h HelpMessagePresenter) buildArgumentSpecificationMessages(argumentSpecificationLabelLengths []int,
	argumentSpecifications []string) string {
	maxArgumentSpecificationLabelLength := slices.Max(argumentSpecificationLabelLengths)
	nextWrappedLinePadding := strings.Repeat(" ", IndentLength+maxArgumentSpecificationLabelLength+IndentLength)
	argumentSpecificationMessages := ""
	for i, specification := range argumentSpecifications {
		unwrappedMessage := indent + strings.Replace(specification, PaddingPlaceholder,
			h.calculatePaddingFrom(maxArgumentSpecificationLabelLength, argumentSpecificationLabelLengths[i]), 1)
		argumentSpecificationMessages += h.textWrapWithWrappedLinePadding(unwrappedMessage, nextWrappedLinePadding)
	}
	return argumentSpecificationMessages
}

func (h HelpMessagePresenter) calculatePaddingFrom(maxArgumentSpecificationLabelLength int,
	currentArgumentSpecificationLabelLength int) string {
	return strings.Repeat(" ", h.absoluteValueOf(
		(maxArgumentSpecificationLabelLength+IndentLength)-currentArgumentSpecificationLabelLength))
}

func (h HelpMessagePresenter) absoluteValueOf(value int) int {
	return max(value, -value)
}

func (h HelpMessagePresenter) textWrapWithWrappedLinePadding(text string, wrappedLinePadding string) string {
	words := strings.Split(text, " ")
	if len(words) == 0 {
		return text
	}
	return h.wrap(words, wrappedLinePadding)
}

func (h HelpMessagePresenter) textWrap(text string) string {
	words := strings.Fields(strings.TrimSpace(text))
	if len(words) == 0 {
		return text
	}
	return h.wrap(words, "")
}

func (h HelpMessagePresenter) wrap(words []string, wrappedLinePadding string) string {
	wrapped := words[0]
	spaceLeft := LineLength - len(wrapped)
	for _, word := range words[1:] {
		if len(word)+1 > spaceLeft {
			wrapped += "\n" + wrappedLinePadding + word
			spaceLeft = LineLength - len(wrappedLinePadding+word)
		} else {
			wrapped += " " + word
			spaceLeft -= 1 + len(word)
		}
	}
	return wrapped
}
