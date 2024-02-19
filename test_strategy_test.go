package gotest

import (
	"fmt"
	"testing"
)

var _ TWrapper = (*testing.T)(nil)

func TestRealStrategyOnce(t *testing.T) {
	called := false
	strategy := GetRealRunStrategy(t)

	strategy("should invoke Run from testing object", func(TWrapper) {
		called = true
	})

	if !called {
		t.Log("Testing object was not invoked.")
		t.FailNow()
	}
}

func TestRealStrategySequential(t *testing.T) {
	called := [3]bool{false, false, false}
	strategy := GetRealRunStrategy(t)

	for i := range [3]int{0, 1, 2} {
		testName := fmt.Sprintf("should invoke Run from testing object %v", i)
		strategy(testName, func(tt TWrapper) {
			if called[i] {
				tt.FailNow()
			}
			called[i] = true
		})
	}

	for i, c := range called {
		if !c {
			t.Log("Testing object", i, "was not invoked.")
			t.Fail()
		}
	}
}

func TestRealStrategyParallel(t *testing.T) {
	t.Parallel()
	strategy := GetRealRunStrategy(t)
	called := [3]bool{false, false, false}

	for i := range [3]int{0, 1, 2} {
		testName := fmt.Sprintf("should invoke Run from testing object %v", i)
		strategy(testName, func(tt TWrapper) {
			tt.Parallel()
			if called[i] {
				tt.FailNow()
			}
			called[i] = true
		})
	}

	t.Cleanup(func() {
		for i, c := range called {
			if !c {
				t.Log("Testing object", i, "was not invoked.")
				t.Fail()
			}
		}
	})
}
