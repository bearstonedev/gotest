package tests

import "testing"

type Runner struct {
	t *testing.T
}

func Suite(t *testing.T) *Runner {
	return &Runner{t}
}

func (r *Runner) Test(name string, test func() bool) *Runner {
	r.t.Helper()
	r.t.Run(name, func(t *testing.T) {
		t.Helper()
		t.Parallel()
		if !test() {
			t.Error("Test failed: " + name)
		}
	})
	return r
}
