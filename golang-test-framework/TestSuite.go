package golang_test_framework

import (
	"testing"
)

type suite struct {
	*testing.T
}

func Tests(t *testing.T) TestSuite {
	t.Parallel()
	return &suite{t}
}

type TestSuite interface {
	Test(string, func(Assertions)) TestSuite
	Scenarios() Parameterized
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
