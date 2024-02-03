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

		mockT.verifyCallCount("Parallel")
	})
}