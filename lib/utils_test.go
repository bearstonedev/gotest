package lib_test

import (
	gotest "github.com/bearstonedev/gotest/lib"
	"testing"
)

type call struct {
	count int
	args  []any
}

type tMock struct {
	real  *testing.T
	calls map[string]*call
}

var _ gotest.T = (*tMock)(nil)

func (mock *tMock) Helper() {
	mock.incrementCallCount("Helper")
}

func (mock *tMock) Log(a ...any) {
	mock.trackCall("Log", a)
}

func (mock *tMock) FailNow() {
	mock.incrementCallCount("FailNow")
}

func (mock *tMock) Run(name string, f func(t *testing.T)) bool {
	mock.trackCall("Run", name, f)
	return true
}

func (mock *tMock) Parallel() {
	mock.incrementCallCount("Parallel")
}

func (mock *tMock) incrementCallCount(calledName string) {
	mock.trackCall(calledName, nil)
}

func (mock *tMock) trackCall(calledName string, args ...any) {
	maybe := mock.calls[calledName]
	if maybe == nil {
		maybe = &call{0, nil}
	}

	newCall := *maybe
	newCall.count++
	if args != nil {
		newCall.args = args
	}

	mock.calls[calledName] = &newCall
}

func mockTesting(t *testing.T) *tMock {
	return &tMock{t, make(map[string]*call)}
}

func (mock *tMock) verifyCallCount(callName string) {
	mock.real.Helper()

	theCall := mock.calls[callName]
	if theCall == nil {
		mock.logAndFail(callName, "was never called.")
	}

	if theCall.count != 1 {
		mock.logAndFail("Incorrect count for:", callName, "count:", theCall.count)
	}
}

func logAndFail(t *testing.T, args ...any) {
	t.Helper()
	t.Error(args...)
	t.FailNow()
}

func (mock *tMock) logAndFail(args ...any) {
	mock.real.Helper()
	logAndFail(mock.real, args...)
}
