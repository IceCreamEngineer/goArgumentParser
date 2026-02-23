package useCases

import "errors"

type ArgumentParser struct {
	Schema, Arguments string
}

func (a ArgumentParser) Parse() error {
	return errors.New("invalid arguments")
}

func (a ArgumentParser) NextArgument() int {
	return 0
}
