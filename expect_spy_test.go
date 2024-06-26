package gotest

var _ testContext = &fakeTestContext{}

type fakeTestContext struct {
	name              string
	failWasCalled     bool
	parallelWasCalled bool
}

func (s *fakeTestContext) Fail() {
	s.failWasCalled = true
}

func (s *fakeTestContext) Parallel() {
	s.parallelWasCalled = true
}

func (s *fakeTestContext) CreateSubtest(name string, subtest func(context testContext)) {
	spy := createContextSpy()
	subtest(spy)
	spy.name = name
}

func createContextSpy() *fakeTestContext {
	return &fakeTestContext{}
}

func getSpy(e *Expectation) *fakeTestContext {
	return e.context.(*fakeTestContext)
}

func (s *fakeTestContext) context() testContext {
	var stub testContext = s
	return stub
}

func (s *fakeTestContext) wasCalled(expected string) bool {
	switch expected {
	case "fail":
		return s.failWasCalled
	case "parallel":
		return s.parallelWasCalled
	default:
		return false
	}
}

func (s *fakeTestContext) reset() {
	s.failWasCalled = false
	s.parallelWasCalled = false
}
