package stack

type Stack[T comparable] struct {
	Data []T
}

func New[T comparable](data []T) Stack[T] {
	return Stack[T]{
		Data: data,
	}
}

func (s *Stack[T]) Push(e T) {
	s.Data = append(s.Data, e)
}

func (s *Stack[T]) Peek() (e T) {
	e = s.Data[len(s.Data)-1]
	return
}

func (s *Stack[T]) Pop() (c bool, e T) {
	c = true
	if len(s.Data) < 1 {
		c = false
		return
	}
	e = s.Data[len(s.Data)-1]
	s.Data = s.Data[:len(s.Data)-1]
	return
}
