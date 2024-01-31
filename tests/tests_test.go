package tests

import (
	"testing"
)

func TestAddition(t *testing.T) {
	Suite(t).Test("1 + 1 should be 2", func() bool {
		return 1+1 == 2
	}).Test("1 + 2 should not be 1", func() bool {
		return 1+2 != 1
	})
}

func TestSubtraction(t *testing.T) {
	Suite(t).Test("1 - 1 should be 0", func() bool {
		return 1-1 == 0
	}).Test("1 - 1 should not be 2", func() bool {
		return 1-1 != 2
	})
}
