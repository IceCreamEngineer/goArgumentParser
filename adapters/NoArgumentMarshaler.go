package adapters

type NoArgumentMarshaler struct{}

func (m NoArgumentMarshaler) Set(nextArgument func() (any, bool)) error {
	return nil
}

func (m NoArgumentMarshaler) GetValue() any {
	return nil
}
