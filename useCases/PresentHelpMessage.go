package useCases

import (
	"goArgumentParser/entities"
	"goArgumentParser/ports"
	"strings"
	"unicode/utf8"
)

type HelpMessagePresenter struct {
	ProgramFileName string
	Description     string
	Presenter       ports.Presenter
}

var helpMessage string

const MessageCharacterWidth = 30

func (h HelpMessagePresenter) PresentHelpMessage(schema []entities.ArgumentSchemaElement) {
	programTitle, argumentSpecifications := h.buildSchemaDependentMessages(schema)
	helpMessage = "" +
		programTitle +
		"\n" +
		h.Description + "\n" +
		"\n" +
		"optional arguments:\n" +
		"  -h, --help            show this help message and exit\n" +
		argumentSpecifications
	h.Presenter.Present(helpMessage)
}

func (h HelpMessagePresenter) buildSchemaDependentMessages(schema []entities.ArgumentSchemaElement) (string, string) {
	programTitle := "usage: " + h.ProgramFileName + " [-h]"
	argumentSpecifications := ""
	for _, schemaElement := range schema {
		programTitle += h.buildProgramTitleFrom(schemaElement)
		argumentSpecifications += h.buildArgumentSpecificationFrom(schemaElement)
	}
	return programTitle + "\n", argumentSpecifications
}

func (h HelpMessagePresenter) buildProgramTitleFrom(schemaElement entities.ArgumentSchemaElement) string {
	if schemaElement.IsRequired() {
		return " -" + schemaElement.Name
	}
	return " [-" + schemaElement.Name + "]"
}

func (h HelpMessagePresenter) buildArgumentSpecificationFrom(schemaElement entities.ArgumentSchemaElement) string {
	longNameAddition := h.checkToAddLongName(schemaElement)
	spacer := h.calculateSpacerFrom(schemaElement, longNameAddition)
	return "  -" + schemaElement.Name + longNameAddition + spacer + schemaElement.Description + "\n"
}

func (h HelpMessagePresenter) checkToAddLongName(schemaElement entities.ArgumentSchemaElement) string {
	if schemaElement.LongName != "" {
		return ", --" + schemaElement.LongName
	}
	return ""
}

func (h HelpMessagePresenter) calculateSpacerFrom(schemaElement entities.ArgumentSchemaElement,
	longNameAddition string) string {
	argumentSpecificationWithoutSpacer := "  -" + schemaElement.Name + longNameAddition + schemaElement.Description
	withoutSpacerWidth := utf8.RuneCountInString(argumentSpecificationWithoutSpacer)
	return strings.Repeat(" ", MessageCharacterWidth-withoutSpacerWidth)
}
