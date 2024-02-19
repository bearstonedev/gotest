package gotest

import (
	"testing"
)

func TestRunner_Orchestration(t *testing.T) {
	t.Parallel()
	t.Run("should run tests automatically and in parallel", func(tt *testing.T) {
		tt.Parallel()
		sut, mock := injectMock(tt)
		asserts := createAssertions(tt)
		mock.shouldBeCalled("Parallel")

		testName := "some test"
		wasTestCalled := false
		testBody := func(Assertions) {
			wasTestCalled = true
		}
		sut.Test(testName, testBody)
		mock.shouldBeCalledWithSome("Run", testName)
		asserts.True(wasTestCalled)
	})
}

func injectMock(t *testing.T) (TestRunner, *tMock) {
	m := mockTesting(t)
	sut := Tests(m).(*runner)
	sut.runTest = func(testName string, testBody func(TWrapper)) {
		m.Run(testName, func(*testing.T) {
			testBody(m)
		})
	}

	return sut, m
}
