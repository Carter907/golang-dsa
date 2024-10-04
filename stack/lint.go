package stack

func Lint(s string) bool {
	stack := Stack[rune]{
		[]rune{},
	}
	for _, r := range []rune(s) {
		switch r {
		case ')':
			if c, e := stack.Pop(); !c || e != '(' {
				return false
			}
		case ']':
			if c, e := stack.Pop(); !c || e != '[' {
				return false
			}
		case '}':

			if c, e := stack.Pop(); !c || e != '{' {
				return false
			}
		case '(', '{', '[':
			stack.Push(r)
		}
	}
	return len(stack.Data) == 0
}
