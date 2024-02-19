package gotest

import "github.com/google/go-cmp/cmp"

type Assertions interface {
	Log(...any)
	Equal(any, any)
	NotEqual(any, any)
	True(bool, ...any)
	False(bool, ...any)
}

type assertions struct {
	TWrapper
}

func createAssertions(t TWrapper) Assertions {
	return &assertions{t}
}

func (a *assertions) Equal(lhs any, rhs any) {
	a.Helper()
	if !cmp.Equal(lhs, rhs) {
		a.Log(a.Name(), "Expected", lhs, "to be equal to", rhs)
		a.FailNow()
	}
}

func (a *assertions) NotEqual(lhs any, rhs any) {
	a.Helper()
	if cmp.Equal(lhs, rhs) {
		a.Log(a.Name(), "Expected", lhs, "not to be equal to", rhs)
		a.FailNow()
	}
}

func (a *assertions) True(test bool, message ...any) {
	a.Helper()
	if !test {
		messageWithPrefix := a.prependTestName(&message)
		a.Log(*messageWithPrefix...)
		a.FailNow()
	}
}

func (a *assertions) False(test bool, message ...any) {
	a.True(!test, message...)
}

func (a *assertions) prependTestName(message *[]any) *[]any {
	prefix := &[]any{"Test", "\"" + a.Name() + "\"", "failed:"}
	messageWithPrefix := append(*prefix, *message...)
	return &messageWithPrefix
}
