package gotest

import (
	"testing"
)

var _ Assertions = (*assertions)(nil)

func TestAssertions_Inequality(t *testing.T) {
	Tests(t).Scenarios().
		Scenario("Two unequal numbers should be unequal", 1, 2).
		Scenario("Two unequal strings should be unequal", "same", "different").
		Scenario("Two unequal slices should be unequal", []string{"ラメン"}, []string{"寿司"}).
		Test(func(sut Assertions, scenario ...any) {
			mockT := setUpSystemUnderTest(sut)
			lhs := scenario[0]
			rhs := scenario[1]

			sut.NotEqual(lhs, rhs)
			mockT.shouldNotFailTest()

			sut.Equal(lhs, rhs)
			mockT.shouldFailTest(mockT.real.Name(), "Expected", lhs, "to be equal to", rhs)
		})
}

func TestAssertions_Equality(t *testing.T) {
	Tests(t).Scenarios().
		Scenario("Two equal numbers should be equal", 1, 1).
		Scenario("Two equal strings should be equal", "same", "same").
		Scenario("Two equal slices should be equal", []string{"ラメン"}, []string{"ラメン"}).
		Test(func(sut Assertions, scenario ...any) {
			mockT := setUpSystemUnderTest(sut)
			lhs := scenario[0]
			rhs := scenario[1]

			sut.Equal(lhs, rhs)
			mockT.shouldNotFailTest()

			sut.NotEqual(lhs, rhs)
			mockT.shouldFailTest(mockT.real.Name(), "Expected", lhs, "not to be equal to", rhs)
		})
}

func TestAssertions_TrueAndFalse(t *testing.T) {
	Tests(t).Scenarios().
		Scenario("true should be true", Assertions.True, true, "Expected true to be true").
		Scenario("false should be false", Assertions.False, false, "Expected false to be false").
		Test(func(sut Assertions, scenario ...any) {
			mockT := setUpSystemUnderTest(sut)

			assertionUnderTest := scenario[0].(func(Assertions, bool, ...any))
			expected := scenario[1].(bool)
			message := scenario[2].(string)

			assertionUnderTest(sut, expected, message)
			mockT.shouldBeCalledTimes(1, "Helper")
			mockT.shouldNotFailTest()

			assertionUnderTest(sut, !expected, message)
			mockT.shouldBeCalledTimes(2, "Helper")
			mockT.shouldFailTest("Test", "\""+mockT.real.Name()+"\"", "failed:", message)
		})
}

func TestAssertions_Log(t *testing.T) {
	Tests(t).Test("should pass logging calls", func(sut Assertions) {
		mockT := setUpSystemUnderTest(sut)
		args := []any{"some", "stuff", "to", "log"}
		sut.Log(args...)
		mockT.shouldBeCalledWith("Log", args...)
	})
}

func setUpSystemUnderTest(asserts Assertions) *tMock {
	sut := asserts.(*assertions)
	mockT := mockTesting(sut.TWrapper)
	sut.TWrapper = mockT
	return mockT
}
