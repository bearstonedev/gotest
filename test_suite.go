package gotest

import (
	"testing"
)

type TestRunner interface {
	Test(string, TestBody) TestRunner
	Scenarios() Parameterized
}

func Tests(t T) TestRunner {
	t.Parallel()
	return &runner{t, DefaultRunnerStrategy}
}

type runner struct {
	T
	runTest TestRunnerStrategy
}

type TestBody func(Assertions)

func (i *runner) Test(name string, test TestBody) TestRunner {
	i.Run(name, func(t *testing.T) {
		i.runTest(t, name, test)
	})
	return i
}

func (i *runner) Scenarios() Parameterized {
	return createParameterized(i)
}

type ConfigurableTestRunner interface {
	TestRunner
	ChangeTestRunnerStrategy(TestRunnerStrategy)
}

func (i *runner) ChangeTestRunnerStrategy(strategy TestRunnerStrategy) {
	i.runTest = strategy
}

type TestRunnerStrategy func(T, string, TestBody)

func DefaultRunnerStrategy(t T, name string, test TestBody) {
	asserts := createAssertions(name, t)
	t.Parallel()
	t.Helper()
	test(asserts)
}

type T interface {
	Parallel()
	Helper()
	Log(...any)
	FailNow()
	Error(...any)
	Run(name string, f func(t *testing.T)) bool
}
