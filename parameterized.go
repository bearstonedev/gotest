package gotest

import (
	"strconv"
	"strings"
)

type Parameterized interface {
	Scenario(string, ...any) Parameterized
	Test(func(Assertions, ...any)) TestRunner
}

func createParameterized(runner TestRunner, t TWrapper) *parameterized {
	return &parameterized{t, runner, make([]scenario, 0)}
}

type parameterized struct {
	TWrapper
	runner    TestRunner
	scenarios []scenario
}

type scenario struct {
	name *string
	args []any
}

func (p *parameterized) Scenario(name string, args ...any) Parameterized {
	p.scenarios = append(p.scenarios, scenario{&name, args})
	return p
}

func (p *parameterized) Test(body func(Assertions, ...any)) TestRunner {
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
		if len(strings.TrimSpace(*testScenario.name)) == 0 {
			name = "Test #" + strconv.Itoa(testNumber)
		} else {
			name = *testScenario.name
		}

		p.runner.Test(name, func(a Assertions) {
			testScenarioCopy := testScenario
			a.True(len(testScenarioCopy.args) == parameterCount, "Parameter count for scenario doesn't match prior scenarios.")
			body(a, testScenarioCopy.args...)
		})
	}

	return p.runner
}
