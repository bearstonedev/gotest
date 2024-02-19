package gotest

import (
	"testing"
)

func TestParameterized_Scenario(t *testing.T) {
	r := Tests(t)
	r.Test("should set up scenarios", func(shouldBe Assertions) {
		p, _ := createParameterizedWithMock(r)
		wasCalled := false
		p.Scenario("", "first arg", "second arg").
			Test(func(innerShould Assertions, parameters ...any) {
				shouldBe.Equal(parameters[0], "first arg")
				shouldBe.Equal(parameters[1], "second arg")
				wasCalled = true
			})
		shouldBe.True(wasCalled)
	})
	r.Test("should fail without scenarios", func(Assertions) {
		p, m := createParameterizedWithMock(r)
		p.Test(func(Assertions, ...any) {
		})
		m.shouldBeCalled("FailNow")
	})
}

func createParameterizedWithMock(testRunner TestRunner) (Parameterized, *tMock) {
	exposedRunner := testRunner.(*runner)
	m := mockTesting(exposedRunner)
	exposedRunner.runTest = func(name string, testBody func(TWrapper)) {
		m.Run(name, func(t *testing.T) {
			testBody(m)
		})
	}
	p := createParameterized(testRunner, m)
	return p, m
}
