package examples

import (
	. "github.com/bearstonedev/gotest"
	"testing"
)

func TestShouldAddTwoNumbers(t *testing.T) {
	systemUnderTest := CreateCalculator()
	Tests(t).Test("1 + 1 should be 2", func(shouldBe Assertions) {
		expectedOutput := 2
		actualOutput := systemUnderTest.Add(1, 1)
		shouldBe.Equal(actualOutput, expectedOutput)
	}).Test("1 + 2 should be 3", func(shouldBe Assertions) {
		expectedOutput := 3
		actualOutput := systemUnderTest.Add(1, 2)
		shouldBe.Equal(actualOutput, expectedOutput)
	}).Test("2 + 2 should be 4", func(shouldBe Assertions) {
		expectedOutput := 4
		actualOutput := systemUnderTest.Add(2, 2)
		shouldBe.Equal(actualOutput, expectedOutput)
	}).Test("-1 + -1 should be -2", func(shouldBe Assertions) {
		expectedOutput := -2
		actualOutput := systemUnderTest.Add(-1, -1)
		shouldBe.Equal(actualOutput, expectedOutput)
	}).Test("-1 + 1 should be 0", func(shouldBe Assertions) {
		expectedOutput := 0
		actualOutput := systemUnderTest.Add(-1, 1)
		shouldBe.Equal(actualOutput, expectedOutput)
	})
}

func TestShouldCompareTwoNumbers(t *testing.T) {
	sut := CreateCalculator()
	Tests(t).
		Test("2 should be greater than 1", func(shouldBe Assertions) {
			shouldBe.True(func() bool {
				return sut.IsGreaterThan(2, 1)
			})
		}).
		Test("1 should not be greater than 2", func(shouldBe Assertions) {
			shouldBe.False(func() bool {
				return sut.IsGreaterThan(1, 2)
			})
		})
}

func TestShouldSubtractTwoNumbers(t *testing.T) {
	sut := CreateCalculator()
	Tests(t).
		Scenarios().
		Scenario(1, 1, 0).
		NamedScenario("1 - 2 should be -1", 1, 2, -1).
		NamedScenario("2 - 2 should be 0", 2, 2, 0).
		NamedScenario("1 - 0 should be 1", 1, 0, 1).
		NamedScenario("0 - 1 should be -1", 0, 1, -1).
		Test(func(shouldBe Assertions, args []any) {
			expectedOutput := args[2].(int)
			actualOutput := sut.Subtract(args[0].(int), args[1].(int))
			shouldBe.Equal(actualOutput, expectedOutput)
		})
}
