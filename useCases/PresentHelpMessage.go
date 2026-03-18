package useCases

import (
	"goArgumentParser/entities"
	"goArgumentParser/ports"
)

type HelpMessagePresenter struct {
	ProgramFileName string
	Description     string
	Presenter       ports.Presenter
}

var helpMessage string

func (h HelpMessagePresenter) PresentHelpMessage(schema []entities.ArgumentSchemaElement) {
	helpMessage = "" +
		h.buildProgramTitleFrom(schema) +
		"\n" +
		h.Description + "\n" +
		"\n" +
		"optional arguments:\n" +
		"  -h, --help            show this help message and exit\n" +
		h.buildArgumentSpecificationsFrom(schema)
	h.Presenter.Present(helpMessage)
}

func (h HelpMessagePresenter) buildProgramTitleFrom(schema []entities.ArgumentSchemaElement) string {
	programTitle := "usage: " + h.ProgramFileName + " [-h]"
	for _, schemaElement := range schema {
		if schemaElement.IsRequired() {
			programTitle += " -" + schemaElement.Name
		} else {
			programTitle += " [-" + schemaElement.Name + "]"
		}
	}
	return programTitle + "\n"
}

func (h HelpMessagePresenter) buildArgumentSpecificationsFrom(schema []entities.ArgumentSchemaElement) string {
	argumentSpecifications := ""
	for _, schemaElement := range schema {
		argumentSpecifications += "  -" + schemaElement.Name + "                    " + schemaElement.Description + "\n"
	}
	return argumentSpecifications
}
