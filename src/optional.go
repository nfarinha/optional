package optional

type Optional[T any] struct {
	defined bool
	value   T
}

func New[T any](value T) Optional[T] {
	return Optional[T]{defined: true, value: value}
}

func (o *Optional[T]) Nil() *Optional[T] {
	o.defined = false
	return o
}

func (o *Optional[T]) IsDefined() bool {
	return o.defined
}
