package adapters

import (
	"goArgumentParser/entities"
)

var stringValue string

type StringArgumentMarshaler struct{}

func (m StringArgumentMarshaler) Set(nextArgument func() (any, bool)) error {
	aValue, ok := nextArgument()
	if !ok {
		return &entities.ArgumentError{ErrorCode: entities.MissingString}
	}
	stringValue = aValue.(string)
	return nil
}

func (m StringArgumentMarshaler) GetValue() any {
	return stringValue
}
