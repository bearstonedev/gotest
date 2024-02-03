package gotest

import (
	"testing"
)

type T interface {
	Parallel()
	Helper()
	Log(...any)
	FailNow()
	Run(name string, f func(t *testing.T)) bool
}

type suite struct {
	T
}

type TestSuite interface {
	Test(string, func(Assertions)) TestSuite
	Scenarios() Parameterized
}

func Tests(t T) TestSuite {
	t.Parallel()
	return &suite{t}
}

func (s *suite) Test(name string, test func(Assertions)) TestSuite {
	s.Run(name, func(t *testing.T) {
		asserts := createAssertions(name, t)
		t.Parallel()
		t.Helper()
		test(asserts)
	})
	return s
}

func (s *suite) Scenarios() Parameterized {
	return createParameterized(s.T)
}
