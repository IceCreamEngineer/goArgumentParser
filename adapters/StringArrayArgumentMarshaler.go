package adapters

import (
	"goArgumentParser/entities"
)

var strings []string

type StringArrayArgumentMarshaler struct{}

func (m StringArrayArgumentMarshaler) Set(nextArgument func() (any, bool)) error {
	aValue, ok := nextArgument()
	strings = append(strings, aValue.(string))
	if !ok {
		return &entities.ArgumentError{ErrorCode: entities.MissingString}
	}
	return nil
}

func (m StringArrayArgumentMarshaler) GetValue() any {
	return strings
}
