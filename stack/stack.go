package stack

type Stack[T comparable] struct {
	Data []T
}

func (s *Stack[T]) Push(e T) {
	s.Data = append(s.Data, e)
}

func (s *Stack[T]) Pop() (e T) {
	e = s.Data[len(s.Data)-1]
	s.Data = s.Data[:len(s.Data)-1]
	return
}
