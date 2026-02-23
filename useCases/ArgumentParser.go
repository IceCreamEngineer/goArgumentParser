package useCases

type ArgumentParser struct {
	schema, arguments string
}

func (a ArgumentParser) NextArgument() int {
	return 0
}
