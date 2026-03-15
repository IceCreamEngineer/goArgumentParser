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
		"usage: client.go [-h]\n" +
		"\n" +
		"My client\n" +
		"\n" +
		"optional arguments:\n" +
		"  -h, --help            show this help message and exit\n"
	h.Presenter.Present(helpMessage)
}
