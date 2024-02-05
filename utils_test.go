package gotest_test

import (
	"github.com/bearstonedev/gotest"
	"slices"
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
	f(nil)
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

func (mock *tMock) shouldBeCalledTimes(times int, callName string) {
	mock.real.Helper()

	theCall := mock.calls[callName]
	if theCall == nil {
		mock.logAndFail(callName, "was never called.")
	}

	if theCall.count != times {
		mock.logAndFail("Incorrect call count for:", callName, "wanted:", times, "got:", theCall.count)
	}
}

func (mock *tMock) shouldBeCalled(callName string) {
	mock.shouldBeCalledTimes(1, callName)
}

func (mock *tMock) shouldBeCalledWith(callName string, args ...any) {
	mock.shouldBeCalled(callName)

	if len(mock.calls[callName].args) != len(args) {
		mock.logAndFail(callName, "was called with too few args", args, "Wanted:", len(args), "got:", len(mock.calls[callName].args))
	}

	for index, arg := range args {
		if mock.calls[callName].args[index] != arg {
			mock.logAndFail(callName, "was not called with", arg)
		}
	}
}

func (mock *tMock) shouldBeCalledWithSome(callName string, args ...any) {
	mock.shouldBeCalled(callName)

	for _, arg := range args {
		if !slices.Contains(mock.calls[callName].args, arg) {
			mock.logAndFail(callName, "was not called with", arg)
		}
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
