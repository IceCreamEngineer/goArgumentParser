package ports

import "iter"

type ArgumentMarshaler interface {
	Set(currentArgument iter.Seq[any]) error
	GetValueFrom(marshaler ArgumentMarshaler) any
}
