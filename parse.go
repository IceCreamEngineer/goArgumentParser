// usr/bin/true; exec /usr/bin/env go run "$0" "$@"
package main

import (
	"fmt"
	"goArgumentParser/adapters"
	"goArgumentParser/entities"
	"goArgumentParser/useCases"
	"os"
)

func main() {
	presenter := adapters.Printer{}
	helpMessagePresenter := useCases.HelpMessagePresenter{Presenter: presenter}
	argumentMarshalerFactory := adapters.StringsArgumentMarshalerFactory{}
	arrayRequired := false
	schema := []entities.ArgumentSchemaElement{{Name: "s", LongName: "string", ArgumentType: "*", Description: "A string argument."},
		{Name: "a", LongName: "stringArray", ArgumentType: "[*]", Description: "A string array argument.", Required: &arrayRequired}}
	arguments := os.Args[1:]
	argumentParser := useCases.ArgumentParser{Arguments: arguments, Schema: schema, MarshalerFactory: argumentMarshalerFactory, HelpMessagePresenter: helpMessagePresenter}
	err := argumentParser.Parse()
	if err != nil {
		panic(err)
	}
	presenter.Present("My string argument is " + fmt.Sprint(argumentParser.GetValueOf(entities.ArgumentNames{Name: "s", LongName: "string"})))
}
