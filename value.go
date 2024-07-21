package null

type Value[T any] struct {
	value *T
	valid bool
}

func ValueFrom[T any](val T) Value[T] {
	return Value[T]{&val, true}
}

func ValueFromPtr[T any](val *T) Value[T] {
	return Value[T]{val, val != nil}
}

func (s Value[T]) ValueOrZero() T {
	if !s.valid {
		var zero T
		return zero
	}

	return *s.value
}

func (s Value[T]) ValueOrDefault(def T) T {
	if !s.valid {
		return def
	}

	return *s.value
}

func (s Value[T]) Ptr() *T {
	return s.value
}
