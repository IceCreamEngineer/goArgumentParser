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
		"  -h, --help  show this help message and exit\n")
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
		"  -h, --help  show this help message and exit\n"+
		"  -a          My arg\n")
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
		"  -h, --help  show this help message and exit\n"+
		"  -a, --arg   My arg\n")
}

func TestPresentHelpForRequiredArgs(t *testing.T) {
	helpSetup()
	argumentParser := useCases.ArgumentParser{Schema: []entities.ArgumentSchemaElement{{Name: "a", LongName: "arg",
		Description: "My arg"}, {Name: "b", LongName: "otherarg", Description: "My other arg"}}, Arguments: []string{"-h"},
		MarshalerFactory: simpleArgumentMarshalerFactory, HelpMessagePresenter: helpMessagePresenter}
	AssertThatThereWasNoError(t, argumentParser.Parse())
	AssertPresented(t, ""+
		"usage: client.go [-h] -a -b\n"+
		"\n"+
		"My client\n"+
		"\n"+
		"optional arguments:\n"+
		"  -h, --help      show this help message and exit\n"+
		"  -a, --arg       My arg\n"+
		"  -b, --otherarg  My other arg\n")
}

func TestWrappedLongDescription(t *testing.T) {
	helpSetup()
	helpMessagePresenter.Description = "" +
		"Two roads diverged in a yellow wood, And sorry I could not travel both And be one traveler, long I stood And looked down one as far as I could To where it bent in the undergrowth;\n" +
		"Then took the other, as just as fair, And having perhaps the better claim, Because it was grassy and wanted wear; Though as for that the passing there Had worn them really about the same,\n" +
		"And both that morning equally lay In leaves no step had trodden black. Oh, I kept the first for another day! Yet knowing how way leads on to way, I doubted if I should ever come back.\n" +
		"I shall be telling this with a sigh Somewhere ages and ages hence: Two roads diverged in a wood, and I— I took the one less traveled by, And that has made all the difference.\n"
	argumentParser := useCases.ArgumentParser{Arguments: []string{"-h"}, MarshalerFactory: simpleArgumentMarshalerFactory,
		HelpMessagePresenter: helpMessagePresenter}
	AssertThatThereWasNoError(t, argumentParser.Parse())
	AssertPresented(t, ""+
		"usage: client.go [-h]\n"+
		"\n"+
		"Two roads diverged in a yellow wood, And sorry I could not travel both\n"+
		"And be one traveler, long I stood And looked down one as far as I could\n"+
		"To where it bent in the undergrowth; Then took the other, as just as\n"+
		"fair, And having perhaps the better claim, Because it was grassy and\n"+
		"wanted wear; Though as for that the passing there Had worn them really\n"+
		"about the same, And both that morning equally lay In leaves no step had\n"+
		"trodden black. Oh, I kept the first for another day! Yet knowing how way\n"+
		"leads on to way, I doubted if I should ever come back. I shall be\n"+
		"telling this with a sigh Somewhere ages and ages hence: Two roads\n"+
		"diverged in a wood, and I— I took the one less traveled by, And that\n"+
		"has made all the difference.\n"+
		"\n"+
		"optional arguments:\n"+
		"  -h, --help  show this help message and exit\n")
}

func AssertPresented(t *testing.T, presentedMessage string) {
	if presenter.GetPresented() != presentedMessage {
		t.Errorf("Expected presented message '%s' but got '%s'", presentedMessage, presenter.GetPresented())
	}
}
