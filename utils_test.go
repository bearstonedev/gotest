package gotest

import (
	"github.com/google/go-cmp/cmp"
	"slices"
	"testing"
)

type call struct {
	count int
	args  *[]any
}

type tMock struct {
	real     TWrapper
	allCalls map[string]*call
}

var _ TWrapper = (*tMock)(nil)

func (m *tMock) Error(args ...any) {
	m.trackCall("Error", args)
}

func (m *tMock) Helper() {
	m.incrementCallCount("Helper")
}

func (m *tMock) Log(a ...any) {
	m.trackCall("Log", a...)
}

func (m *tMock) FailNow() {
	m.incrementCallCount("FailNow")
}

func (m *tMock) Run(name string, f func(t *testing.T)) bool {
	m.trackCall("Run", name, f)
	f(nil)
	return true
}

func (m *tMock) Parallel() {
	m.incrementCallCount("Parallel")
}

func (m *tMock) Name() string {
	m.incrementCallCount("Name")
	return m.real.Name()
}

func (m *tMock) incrementCallCount(calledName string) {
	m.trackCall(calledName, nil)
}

func (m *tMock) trackCall(calledName string, args ...any) {
	calls := m.allCalls[calledName]
	if calls == nil {
		calls = &call{0, nil}
	}

	calls.count++
	if args != nil {
		calls.args = &args
	}

	m.allCalls[calledName] = calls
}

func mockTesting(t TWrapper) *tMock {
	return &tMock{t, make(map[string]*call)}
}

func (m *tMock) shouldBeCalledTimes(times int, callName string) {
	m.real.Helper()

	theCall := m.allCalls[callName]
	if theCall == nil && times > 0 {
		m.logAndFail(callName, "was never called.")
	}

	if theCall != nil && theCall.count != times {
		m.logAndFail("Incorrect call count for:", callName, "wanted:", times, "got:", theCall.count)
	}
}

func (m *tMock) shouldBeCalled(callName string) {
	m.real.Helper()
	m.shouldBeCalledTimes(1, callName)
}

func (m *tMock) shouldNotBeCalled(callName string) {
	m.real.Helper()
	m.shouldBeCalledTimes(0, callName)
}

func (m *tMock) shouldBeCalledWith(callName string, args ...any) {
	m.real.Helper()
	m.shouldBeCalled(callName)
	if !cmp.Equal(*m.allCalls[callName].args, args) {
		m.logAndFail(callName, "was not called with", args, "; it was called with", *m.allCalls[callName].args)
	}
}

func (m *tMock) shouldBeCalledWithSome(callName string, args ...any) {
	m.shouldBeCalled(callName)

	for _, arg := range args {
		if !slices.Contains(*m.allCalls[callName].args, arg) {
			m.logAndFail(callName, "was not called with", arg)
		}
	}
}

func (m *tMock) logAndFail(args ...any) {
	m.real.Helper()
	m.real.Error(args...)
	m.real.FailNow()
}

func (m *tMock) shouldFailTest(failureMessage ...any) {
	m.real.Helper()
	m.shouldBeCalled("FailNow")
	m.shouldBeCalledWith("Log", failureMessage...)
}

func (m *tMock) shouldNotFailTest() {
	m.real.Helper()
	m.shouldNotBeCalled("Fail")
	m.shouldNotBeCalled("FailNow")
	m.shouldNotBeCalled("Failed")
	m.shouldNotBeCalled("Error")
	m.shouldNotBeCalled("Errorf")
	m.shouldNotBeCalled("Fatal")
	m.shouldNotBeCalled("Fatalf")
	m.shouldNotBeCalled("Skip")
	m.shouldNotBeCalled("SkipNow")
	m.shouldNotBeCalled("Skipped")
	m.shouldNotBeCalled("Skipf")
}
