package tests

import (
	"goArgumentParser/adapters"
	"goArgumentParser/entities"
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
	AssertThatThereWasNoError(t, argumentParser.Parse())
	AssertPresented(t, ""+
		"usage: client.go [-h]\n"+
		"\n"+
		"My client\n"+
		"\n"+
		"optional arguments:\n"+
		"  -h, --help            show this help message and exit\n")
}

func TestPresentHelpWithSchema(t *testing.T) {
	helpSetup()
	required := false
	argumentParser := useCases.ArgumentParser{Schema: []entities.ArgumentSchemaElement{{Name: "a", Description: "My arg",
		Required: &required}}, Arguments: []string{"-h"}, MarshalerFactory: simpleArgumentMarshalerFactory,
		HelpMessagePresenter: helpMessagePresenter}
	AssertThatThereWasNoError(t, argumentParser.Parse())
	AssertPresented(t, ""+
		"usage: client.go [-h] [-a]\n"+
		"\n"+
		"My client\n"+
		"\n"+
		"optional arguments:\n"+
		"  -h, --help            show this help message and exit\n"+
		"  -a                    My arg\n")
}

func TestPresentHelpWithSchemaAndLongName(t *testing.T) {
	helpSetup()
	required := false
	argumentParser := useCases.ArgumentParser{Schema: []entities.ArgumentSchemaElement{{Name: "a", LongName: "arg",
		Description: "My arg", Required: &required}}, Arguments: []string{"-h"},
		MarshalerFactory: simpleArgumentMarshalerFactory, HelpMessagePresenter: helpMessagePresenter}
	AssertThatThereWasNoError(t, argumentParser.Parse())
	AssertPresented(t, ""+
		"usage: client.go [-h] [-a]\n"+
		"\n"+
		"My client\n"+
		"\n"+
		"optional arguments:\n"+
		"  -h, --help            show this help message and exit\n"+
		"  -a, --arg             My arg\n")
}

func AssertPresented(t *testing.T, presentedMessage string) {
	if presenter.GetPresented() != presentedMessage {
		t.Errorf("Expected presented message '%s' but got '%s'", presentedMessage, presenter.GetPresented())
	}
}
