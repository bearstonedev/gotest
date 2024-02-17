package gotest

import (
	"testing"
)

func TestRunner_Orchestration(t *testing.T) {
	t.Parallel()
	t.Run("should run tests automatically and in parallel", func(tt *testing.T) {
		tt.Parallel()
		m := mockTesting(tt)
		configurableRunner := Tests(m).(ConfigurableTestRunner)
		m.shouldBeCalled("Parallel")

		var mockList []*tMock
		testsQueuedCount := 0
		configurableRunner.ChangeTestRunnerStrategy(func(_ T, name string, test TestBody) {
			mm := mockTesting(tt)
			mockList = append(mockList, mm)
			testsQueuedCount++
			DefaultRunnerStrategy(mm, name, test)
		})
		sut := configurableRunner.(TestRunner)

		sampleTest := func(assert Assertions) {
			assert.Equal("yes", "no")
		}
		sut.
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
