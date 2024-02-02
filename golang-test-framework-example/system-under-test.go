package golang_test_framework_example

type systemUnderTest struct{}

func (sut *systemUnderTest) add(lhs int, rhs int) int {
	return lhs + rhs
}

func (sut *systemUnderTest) subtract(lhs int, rhs int) int {
	return lhs - rhs
}

func (sut *systemUnderTest) isGreaterThan(lhs int, rhs int) bool {
	return lhs > rhs
}

func createSystemUnderTest() *systemUnderTest {
	return &systemUnderTest{}
}
