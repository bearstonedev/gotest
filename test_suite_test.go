package gotest_test

import (
	"github.com/bearstonedev/gotest"
	"testing"
)

func TestRunnerOrchestration(t *testing.T) {
	t.Parallel()
	t.Run("should run tests automatically and in parallel", func(tt *testing.T) {
		tt.Parallel()
		m := mockTesting(tt)
		configurableRunner := gotest.Tests(m).(gotest.ConfigurableTestRunner)
		m.shouldBeCalled("Parallel")

		var mockList []*tMock
		testsQueuedCount := 0
		configurableRunner.ChangeTestRunnerStrategy(func(_ gotest.T, name string, test gotest.TestBody) {
			mm := mockTesting(tt)
			mockList = append(mockList, mm)
			testsQueuedCount++
			gotest.DefaultRunnerStrategy(mm, name, test)
		})

		sampleTest := func(assert gotest.Assertions) {
			assert.Equal("yes", "no")
		}
		configurableRunner.
			Test("a test", sampleTest).
			Test("another test", sampleTest)

		m.shouldBeCalledTimes(testsQueuedCount, "Run")
		for _, mm := range mockList {
			mm.shouldBeCalledTimes(2, "Helper")
			mm.shouldBeCalled("Parallel")
			mm.shouldBeCalled("Log")
			mm.shouldBeCalled("FailNow")
		}
	})
}
