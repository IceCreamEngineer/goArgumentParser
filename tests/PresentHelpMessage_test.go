package tests

import (
	"goArgumentParser/adapters"
	"goArgumentParser/ports"
	"goArgumentParser/tests/testDoubles"
	"goArgumentParser/useCases"
	"testing"
)

var (
	simpleArgumentMarshalerFactory ports.ArgumentMarshalerFactory
	presenter                      testDoubles.PresenterSpy
	helpMessagePresenter           useCases.HelpMessagePresenter
)

func helpSetup() {
	simpleArgumentMarshalerFactory = adapters.NoArgumentMarshalerFactory{}
	presenter = testDoubles.NewPresenterSpy()
	helpMessagePresenter = useCases.HelpMessagePresenter{ProgramFileName: "client.go", Description: "My client",
		Presenter: presenter}
}

func TestPresentHelp(t *testing.T) {
	helpSetup()
	argumentParser := useCases.ArgumentParser{Arguments: []string{"-h"},
		MarshalerFactory: simpleArgumentMarshalerFactory, HelpMessagePresenter: helpMessagePresenter}
	parseError := argumentParser.Parse()
	AssertThatThereWasNoError(t, parseError)
	AssertPresented(t, ""+
		"usage: client.go [-h]\n"+
		"\n"+
		"My client\n"+
		"\n"+
		"optional arguments:\n"+
		"  -h, --help            show this help message and exit\n")
}

func AssertPresented(t *testing.T, presentedMessage string) {
	if presenter.GetPresented() != presentedMessage {
		t.Errorf("Expected presented message '%s' but got '%s'", presentedMessage, presenter.GetPresented())
	}
}
