package utils

type Stack[T any] []T

func (s *Stack[T]) PopN(n int) []T {
	l := len(*s) - n
	ret := (*s)[l:]
	*s = (*s)[:l]
	return ret
}

func (s *Stack[T]) Pop() T {
	return s.PopN(1)[0]
}

func (s *Stack[T]) PushN(n []T) {
	*s = append(*s, n...)
}

func (s *Stack[T]) Push(n T) {
	s.PushN([]T{n})
}

func (s Stack[T]) Peek() T {
	return s[len(s)-1]
}
