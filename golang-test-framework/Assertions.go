package golang_test_framework

import "testing"

type Assertions interface {
	Log(...any)
	Equal(any, any)
	NotEqual(any, any)
	True(func() bool, ...any)
	False(func() bool, ...any)
}

type assertions struct {
	testName string
	*testing.T
}

func createAssertions(testName string, t *testing.T) Assertions {
	return &assertions{testName, t}
}

func (a *assertions) Equal(lhs any, rhs any) {
	a.Helper()
	if lhs != rhs {
		a.Log(a.testName, "Expected", lhs, "to be equal to", rhs)
		a.FailNow()
	}
}

func (a *assertions) NotEqual(lhs any, rhs any) {
	a.Helper()
	if lhs == rhs {
		a.Log(a.testName, "Expected", lhs, "not to be equal to", rhs)
		a.FailNow()
	}
}

func (a *assertions) True(test func() bool, message ...any) {
	a.Helper()
	if !test() {
		a.Log(a.testName, message)
		a.FailNow()
	}
}

func (a *assertions) False(test func() bool, message ...any) {
	a.Helper()
	if test() {
		a.Log(a.testName, message)
		a.FailNow()
	}
}
