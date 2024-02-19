package gotest

type TestRunner interface {
	Test(string, TestBody) TestRunner
	Scenarios() Parameterized
}

func Tests(t TWrapper) TestRunner {
	t.Parallel()
	return &runner{t, GetRealRunStrategy(t)}
}

type runner struct {
	TWrapper
	runTest RunStrategy
}

type TestBody func(Assertions)

func (i *runner) Test(name string, test TestBody) TestRunner {
	i.runTest(name, func(t TWrapper) {
		asserts := createAssertions(t)
		t.Parallel()
		t.Helper()
		test(asserts)
	})
	return i
}

func (i *runner) Scenarios() Parameterized {
	return createParameterized(i, i)
}

type TWrapper interface {
	Parallel()
	Helper()
	Log(...any)
	FailNow()
	Error(...any)
	Name() string
}
