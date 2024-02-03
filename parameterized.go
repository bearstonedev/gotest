package gotest

import (
	"strconv"
	"testing"
)

func createParameterized(t T) *parameterized {
	return &parameterized{t, make([]scenario, 0)}
}

type parameterized struct {
	T
	scenarios []scenario
}

type scenario struct {
	name *string
	args []any
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
