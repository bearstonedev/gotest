package examples

import (
	. "github.com/bearstonedev/gotest"
	"testing"
)

func Test_ShouldAddTwoNumbers(t *testing.T) {
	calc := CreateCalculator()
	I := Expec(t)

	I.Expect(calc.Add(1, 1)).ToBe(2)
	I.Expect(calc.Add(1, 2)).ToBe(3)
	I.Expect(calc.Add(2, 2)).ToBe(4)
	I.Expect(calc.Add(-1, -1)).ToBe(-2)
	I.Expect(calc.Add(-1, 1)).ToBe(0)
}

func Test_ShouldCompareTwoNumbers(t *testing.T) {
	calc := CreateCalculator()
	I := Expec(t)
	I.Expect(calc.IsGreaterThan(2, 1)).As("2 > 1").ToBe(true)
	I.Expect(calc.IsGreaterThan(1, 2)).As("1 > 2").ToBe(false)
}

func Test_ShouldSubtractTwoNumbers(t *testing.T) {
	sut := CreateCalculator()
	I := Expec(t)
	I.Expect(sut.Subtract(1, 2)).ToBe(-1).
		Expect(sut.Subtract(1, 2)).ToBe(-1).
		Expect(sut.Subtract(2, 2)).ToBe(0).
		Expect(sut.Subtract(1, 0)).ToBe(1).
		Expect(sut.Subtract(0, 1)).ToBe(-1)
}
