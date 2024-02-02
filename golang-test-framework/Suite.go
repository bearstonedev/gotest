package golang_test_framework

import (
	"strconv"
	"testing"
)

type suite struct {
	*testing.T
}

func Tests(t *testing.T) Suite {
	t.Parallel()
	return &suite{t}
}

type Suite interface {
	Test(string, func(Assertions)) Suite
	Scenarios() Parameterized
}

func (s *suite) Test(name string, test func(Assertions)) Suite {
	s.Run(name, func(t *testing.T) {
		asserts := createAssertions(name, t)
		t.Parallel()
		t.Helper()
		test(asserts)
	})
	return s
}

func (s *suite) Scenarios() Parameterized {
	return &parameterized{s.T, make([]scenario, 0)}
}

type scenario struct {
	name *string
	args []any
}

type parameterized struct {
	*testing.T
	scenarios []scenario
}

type Parameterized interface {
	Scenario(...any) Parameterized
	NamedScenario(string, ...any) Parameterized
	Test(func(Assertions, []any))
}

func (p *parameterized) Scenario(args ...any) Parameterized {
	p.scenarios = append(p.scenarios, scenario{nil, args})
	return p
}

func (p *parameterized) NamedScenario(name string, args ...any) Parameterized {
	p.scenarios = append(p.scenarios, scenario{&name, args})
	return p
}

func (p *parameterized) Test(body func(Assertions, []any)) {
	p.Helper()

	parameterCount := 0
	if len(p.scenarios) >= 1 {
		parameterCount = len(p.scenarios[0].args)
	} else {
		p.Log("No scenarios to run.")
		p.FailNow()
	}

	for testNumber, testScenario := range p.scenarios {
		var name string
		if testScenario.name == nil {
			name = "Test #" + strconv.Itoa(testNumber)
		} else {
			name = *testScenario.name
		}

		p.Run(name, func(t *testing.T) {
			testScenarioCopy := testScenario
			asserts := createAssertions(name, t)
			t.Parallel()
			t.Helper()
			asserts.True(func() bool {
				return len(testScenarioCopy.args) == parameterCount
			}, "Parameter count for scenario doesn't match prior scenarios.")
			body(asserts, testScenarioCopy.args)
		})
	}
}
