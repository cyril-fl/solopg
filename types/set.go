package types

type Set[T comparable] map[T]struct{}

func (s Set[T]) Add(value T) {
	s[value] = struct{}{}
}

func (s Set[T]) Has(value T) bool {
	_, exists := s[value]
	return exists
}

func (s Set[T]) Remove(value T) {
	delete(s, value)
}

func (s Set[T]) Size() int {
	return len(s)
}

func (s Set[T]) ToSlice() []T {
	slice := make([]T, 0, len(s))
	for value := range s {
		slice = append(slice, value)
	}
	return slice
}

func NewSet[T comparable]() Set[T] {
	return make(Set[T])
}