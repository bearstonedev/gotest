package gotest

import "testing"

func Expec(t *testing.T) ExpectationBuilder {
	return createExpectationBuilder(wrapTestContext(t))
}

type nativeTestContext struct {
	*testing.T
}

func wrapTestContext(t *testing.T) testContext {
	return &nativeTestContext{t}
}

func (n *nativeTestContext) CreateSubtest(name string, subtest func(context testContext)) {
	n.Run(name, func(t *testing.T) {
		subtest(wrapTestContext(t))
	})
}
