package shared

import "reflect"

type ValueObject[T any] interface {
	Value() T
	Equals(other ValueObject[T]) bool
}

type valueObject[T any] struct {
	value T
}

func NewValueObject[T any](value T) ValueObject[T] {
	return &valueObject[T]{value: value}
}

func (vo *valueObject[T]) Value() T {
	return vo.value
}

func (vo *valueObject[T]) Equals(other ValueObject[T]) bool {
	return reflect.DeepEqual(vo.Value(), other.Value())
}
