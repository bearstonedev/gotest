package examples

type Calculator interface {
	Add(int, int) int
	Subtract(int, int) int
	IsGreaterThan(int, int) bool
}

type calculator struct{}

func (sut *calculator) Add(lhs int, rhs int) int {
	return lhs + rhs
}

func (sut *calculator) Subtract(lhs int, rhs int) int {
	return lhs - rhs
}

func (sut *calculator) IsGreaterThan(lhs int, rhs int) bool {
	return lhs > rhs
}

func CreateCalculator() Calculator {
	return &calculator{}
}
