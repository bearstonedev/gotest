package gotest_test

import (
	"github.com/bearstonedev/gotest"
	"testing"
)

func TestTests(realTOuter *testing.T) {
	realTOuter.Parallel()
	realTOuter.Run("should create suite", func(realTInner *testing.T) {
		realTInner.Parallel()
		mockT := mockTesting(realTInner)
		testSuite := gotest.Tests(mockT)
		if testSuite == nil {
			logAndFail(realTInner, "testObj is nil")
		}

		mockT.shouldBeCalled("Parallel")
	})
	realTOuter.Run("should create a test", func(realTInner *testing.T) {
		realTInner.Parallel()
		mockTOuter := mockTesting(realTInner)
		testName := "this is a test name"
		timesTestFuncCalled := 0
		testFunc := func(assertions gotest.Assertions) {
			timesTestFuncCalled++
		}
		gotest.Tests(mockTOuter).Test(testName, testFunc)
		mockTOuter.shouldBeCalledWithSome("Run", testName)
	})
}
