package goArgumentParser

import (
	"goArgumentParser/adapters"
	"goArgumentParser/entities"
	"goArgumentParser/useCases"
)

func main() {
	presenter := adapters.Printer{}
	helpMessagePresenter := useCases.HelpMessagePresenter{Presenter: presenter}
	argumentMarshalerFactory := adapters.StringsArgumentMarshalerFactory{}
	arrayRequired := false
	schema := []entities.ArgumentSchemaElement{{Name: "s", LongName: "string", ArgumentType: "*", Description: "A string argument."},
		{Name: "a", LongName: "stringArray", ArgumentType: "[*]", Description: "A string array argument.", Required: &arrayRequired}}
	argumentParser := useCases.ArgumentParser{Schema: schema, MarshalerFactory: argumentMarshalerFactory, HelpMessagePresenter: helpMessagePresenter}
	err := argumentParser.Parse()
	if err != nil {
		panic(err)
	}
}
