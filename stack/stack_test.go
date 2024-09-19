package stack_test

import (
	"testing"

	"github.com/Carter907/golang-dsa/stack"
)

func TestStack(t *testing.T) {
	stack := stack.Stack[int]{
		Data: []int{1, 2, 6, 2},
	}
	t.Log("capacity before pop", cap(stack.Data))
	t.Log(stack)
	stack.Pop()
	t.Log(stack)

	t.Log(stack)
	stack.Push(7)
	t.Log(stack)
	t.Log("capacity after push", cap(stack.Data))
}
