package examples

import (
	"testing"
)

func TestShouldAddTwoNumbers(t *testing.T) {
	sut := createSystemUnderTest()
	Tests(t).Test("1 + 1 should be 2", func(shouldBe Assertions) {
		expectedOutput := 2
		actualOutput := sut.add(1, 1)
		shouldBe.Equal(actualOutput, expectedOutput)
	}).Test("1 + 2 should be 3", func(shouldBe Assertions) {
		expectedOutput := 3
		actualOutput := sut.add(1, 2)
		shouldBe.Equal(actualOutput, expectedOutput)
	}).Test("2 + 2 should be 4", func(shouldBe Assertions) {
		expectedOutput := 4
		actualOutput := sut.add(2, 2)
		shouldBe.Equal(actualOutput, expectedOutput)
	}).Test("-1 + -1 should be -2", func(shouldBe Assertions) {
		expectedOutput := -2
		actualOutput := sut.add(-1, -1)
		shouldBe.Equal(actualOutput, expectedOutput)
	}).Test("-1 + 1 should be 0", func(shouldBe Assertions) {
		expectedOutput := 0
		actualOutput := sut.add(-1, 1)
		shouldBe.Equal(actualOutput, expectedOutput)
	})
}

func TestShouldCompareTwoNumbers(t *testing.T) {
	sut := createSystemUnderTest()
	Tests(t).
		Test("2 should be greater than 1", func(shouldBe Assertions) {
			shouldBe.True(func() bool {
				return sut.isGreaterThan(2, 1)
			})
		}).
		Test("1 should not be greater than 2", func(shouldBe Assertions) {
			shouldBe.False(func() bool {
				return sut.isGreaterThan(1, 2)
			})
		})
}

func TestShouldSubtractTwoNumbers(t *testing.T) {
	sut := createSystemUnderTest()
	Tests(t).
		Scenarios().
		Scenario(1, 1, 0).
		NamedScenario("1 - 2 should be -1", 1, 2, -1).
		NamedScenario("2 - 2 should be 0", 2, 2, 0).
		NamedScenario("1 - 0 should be 1", 1, 0, 1).
		NamedScenario("0 - 1 should be -1", 0, 1, -1).
		Test(func(shouldBe Assertions, args []any) {
			expectedOutput := args[2].(int)
			actualOutput := sut.subtract(args[0].(int), args[1].(int))
			shouldBe.Equal(actualOutput, expectedOutput)
		})
}