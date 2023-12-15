package iset

type set[T comparable] struct {
	data map[T]struct{}
}

func NewSet[T comparable]() *set[T] {
	return &set[T]{data: make(map[T]struct{})}
}

func (s set[T]) Add(t T) {
	s.data[t] = struct{}{}
}

func (s set[T]) Del(t T) {
	delete(s.data, t)
}
func (s set[T]) Exists(t T) bool {
	_, ok := s.data[t]
	return ok
}

func (s set[T]) Range(fun func(t T)) {
	for k := range s.data {
		fun(k)
	}
}
func (s set[T]) Len() int {
	return len(s.data)
}
func (s set[T]) Clear() {
	clear(s.data)
}
