package stack_test

import (
	"testing"

	"github.com/Carter907/golang-dsa/stack"
)

func TestLint(t *testing.T) {
	testValue := "183()912397((()()())(())(1267))"
	ans := stack.Lint(testValue)
	if !ans {
		t.Fatal("test value is", testValue, "and should be correct, but was incorrect")
	}
}
