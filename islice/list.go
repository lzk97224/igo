package islice

type list[T any] []T

func NewList[T any]() list[T] {
	return list[T]{}
}

func (l *list[T]) Add(t T) {
	*l = append(*l, t)
}

func (l *list[T]) Get(index int) T {
	return (*l)[index]
}

func (l *list[T]) Del(index int) {
	if index < 0 || index >= len(*l) {
		return
	}
	*l = append((*l)[0:index], (*l)[index+1:]...)
}

func (l *list[T]) DelIf(fun func(T) bool) {
	*l = Filter[T](*l, fun)
}

func (l *list[T]) Len() int {
	return len(*l)
}
