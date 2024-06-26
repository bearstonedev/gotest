package gotest

import (
	"fmt"
	"testing"
)

func person(name string, age int) struct {
	name string
	age  int
} {
	return struct {
		name string
		age  int
	}{
		name,
		age,
	}
}

func Test_ShouldFailSimpleTest(t *testing.T) {
	t.Parallel()
	shouldFail := func(expect *Expectation) {
		t.Run(fmt.Sprintf("%v should fail", expect), func(tt *testing.T) {
			tt.Parallel()
			spy := getSpy(expect)
			if !spy.wasCalled("fail") {
				tt.Fatalf("Didn't fail when expected")
			}
			spy.reset()
		})
	}

	I := createExpectationBuilder(createContextSpy())
	shouldFail(I.Expect(1 + 1).ToBe(3))
	shouldFail(I.Expect(true).ToBe(false))
	shouldFail(I.Expect("cat").ToBe("dog"))
	shouldFail(I.Expect(nil).ToBe("not nil"))
	shouldFail(I.Expect(person("Frank", 43)).ToBe(person("Bob", 42)))

	if I.context.(*fakeTestContext).failWasCalled {
		t.Fatalf("Only subtests are expected to fail.")
	}
}

func Test_ShouldPassSimpleTest(t *testing.T) {
	t.Parallel()
	shouldPass := func(expect *Expectation) {
		t.Run(expect.String(), func(tt *testing.T) {
			tt.Parallel()
			spy := getSpy(expect)
			if spy.wasCalled("fail") {
				t.Fatalf("Test was expected to pass.")
			}
		})
	}

	I := createExpectationBuilder(createContextSpy())
	shouldPass(I.Expect(1 + 1).ToBe(2))
	shouldPass(I.Expect(true).ToBe(true))
	shouldPass(I.Expect("cat").ToBe("cat"))
	shouldPass(I.Expect(nil).ToBe(nil))
	shouldPass(I.Expect(person("Bob", 42)).ToBe(person("Bob", 42)))

	if I.context.(*fakeTestContext).failWasCalled {
		t.Fatalf("All tests expected to pass.")
	}
}

func Test_ShouldRunInParallel(t *testing.T) {
	t.Parallel()
	spy := createContextSpy()
	createExpectationBuilder(spy.context())
	if !spy.wasCalled("parallel") {
		t.Errorf("Test did not run in parallel.")
	}
}

func Test_ShouldCreateSubtestForExpectation(t *testing.T) {
	t.Parallel()
	spy := createContextSpy()
	builder := createExpectationBuilder(spy.context())
	expectation := builder.Expect(true).ToBe(true)

	if expectation.context == spy.context() {
		t.Fatalf("Subtest context is the same as the test context.")
	}
}

func Test_ShouldCreateSubtestsInParallel(t *testing.T) {
	t.Parallel()
	builder := createExpectationBuilder(createContextSpy())
	expectation := builder.Expect(true).ToBe(true)
	if !getSpy(expectation).wasCalled("parallel") {
		t.Fatalf("Subtest was not run in parallel.")
	}
}

func Test_ShouldDistinguishBetweenSubtests(t *testing.T) {
	t.Parallel()
	builder := createExpectationBuilder(createContextSpy())

	first := builder.Expect(1 + 1).ToBe(3)
	subSpyOne := getSpy(first)
	if !subSpyOne.wasCalled("fail") {
		t.Fatalf("First test didn't fail")
	}
	subSpyOne.reset()

	second := builder.Expect(2 + 2).ToBe(5)
	subSpyTwo := getSpy(second)
	if !subSpyTwo.wasCalled("fail") {
		t.Fatalf("Second test didn't fail")
	}
	if subSpyOne.wasCalled("fail") {
		t.Fatalf("First test failed when it shouldn't have")
	}
}

func Test_ShouldNameExpectations(t *testing.T) {
	t.Parallel()
	I := createExpectationBuilder(createContextSpy())

	name := "I expect 2 to be 2"
	expectation := I.Expect(1 + 1).ToBe(2)
	if expectation.String() != name {
		t.Fatalf("Expectation was not named properly: expected \"%v\" to be \"%v\"", expectation, name)
	}
}

func Test_ShouldSupportCustomNamesForExpectations(t *testing.T) {
	t.Parallel()
	I := createExpectationBuilder(createContextSpy())

	name := "I expect 1 + 1 to be 2"
	expectation := I.Expect(1 + 1).As("1 + 1").ToBe(2)
	if expectation.String() != name {
		t.Fatalf("Expectation was not named properly: expected \"%v\" to be \"%v\"", expectation, name)
	}
}

func Test_ShouldUseExpectationNameWhenCreatingSubtest(t *testing.T) {
	t.Parallel()
	I := createExpectationBuilder(createContextSpy())
	name := "I expect 1 + 1 to be 2"
	expectation := I.Expect(1 + 1).As("1 + 1").ToBe(2)
	nameInContext := getSpy(expectation).name
	if nameInContext != name {
		t.Fatalf("Subtest was not invoked with expected name: expected \"%v\" to be \"%v\"", nameInContext, name)
	}
}
