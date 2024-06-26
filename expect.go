package gotest

import "fmt"

type ExpectationBuilder struct {
	context testContext
}

type testContext interface {
	Fail()
	Parallel()
	CreateSubtest(name string, subtest func(context testContext))
}

func createExpectationBuilder(t testContext) ExpectationBuilder {
	t.Parallel()
	return ExpectationBuilder{
		t,
	}
}

type Actual struct {
	ExpectationBuilder
	actual any
	name   string
}

func (builder *ExpectationBuilder) Expect(actual any) *Actual {
	return &Actual{*builder, actual, ""}
}

type Expectation struct {
	ExpectationBuilder
	Actual
	context  testContext
	name     string
	expected any
}

func (a *Actual) ToBe(expected any) *Expectation {
	name := a.actual
	if a.name != "" {
		name = a.name
	}
	expectationName := fmt.Sprintf("I expect %v to be %v", name, expected)

	var context testContext
	a.context.CreateSubtest(expectationName, func(c testContext) {
		context = c
		context.Parallel()
		if a.actual != expected {
			context.Fail()
		}
	})

	return &Expectation{
		a.ExpectationBuilder,
		*a,
		context,
		expectationName,
		expected,
	}
}

func (a *Actual) As(name string) *Actual {
	a.name = name
	return a
}

func (e *Expectation) String() string {
	return e.name
}
