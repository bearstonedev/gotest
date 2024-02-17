package gotest

import (
	"testing"
)

type TestableAssertions interface {
	Assertions
	createMockT() *tMock
	getTestName() string
}

func (a *assertions) createMockT() *tMock {
	mock := mockTesting(a.T)
	a.T = mock
	return mock
}

func (a *assertions) getTestName() string {
	return a.testName
}

func TestAssertions_Inequality(t *testing.T) {
	Tests(t).Scenarios().
		Scenario("Two unequal numbers should be unequal", 1, 2).
		Scenario("Two unequal strings should be unequal", "same", "different").
		Scenario("Two unequal slices should be unequal", []string{"ラメン"}, []string{"寿司"}).
		Test(func(asserts Assertions, scenario []any) {
			sut := asserts.(TestableAssertions)
			mm := sut.createMockT()
			lhs := scenario[0]
			rhs := scenario[1]
			sut.NotEqual(lhs, rhs)
			mm.shouldNotFailTest()
			sut.Equal(lhs, rhs)
			mm.shouldFailTest(sut.getTestName(), "Expected", lhs, "to be equal to", rhs)
		})
}

func TestAssertions_Equality(t *testing.T) {
	Tests(t).Scenarios().
		Scenario("Two equal numbers should be equal", 1, 1).
		Scenario("Two equal strings should be equal", "same", "same").
		Scenario("Two equal slices should be equal", []string{"ラメン"}, []string{"ラメン"}).
		Test(func(asserts Assertions, scenario []any) {
			sut := asserts.(TestableAssertions)
			mm := sut.createMockT()
			lhs := scenario[0]
			rhs := scenario[1]
			sut.Equal(lhs, rhs)
			mm.shouldNotFailTest()
			sut.NotEqual(lhs, rhs)
			mm.shouldFailTest(sut.getTestName(), "Expected", lhs, "not to be equal to", rhs)
		})
}

func TestAssertions_True(t *testing.T) {
	Tests(t).Test("should pass when true", func(asserts Assertions) {
		sut := asserts.(TestableAssertions)
		mm := sut.createMockT()
		message := "Expected true to be true"
		sut.True(true, message)
		mm.shouldBeCalledTimes(1, "Helper")
		mm.shouldNotFailTest()
		sut.True(false, message)
		mm.shouldBeCalledTimes(2, "Helper")
		mm.shouldFailTest("Test", "\""+sut.getTestName()+"\"", "failed:", message)
	})
}
