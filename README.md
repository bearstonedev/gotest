# gotest

A simple test framework for Go, intended for TDD beginners and novice developers.

## Goals

### :rabbit2: Make it easier to write tests first (TDD) by making tests easier to write.

This is a test:

```go
package calculator

import (
	"github.com/bearstonedev/gotest"
	"testing"
)

func Test_ShouldCompareTwoNumbers(t *testing.T) {
	calculator := MyCalculator{}
	I := gotest.Expec(t)
	I.Expect(calculator.Add(1, 1)).ToBe(2)
}
```

Note the readability; the test **documents** the expected result. Here's how the same test would be written in standard
Go:

```go
package calculator

import "testing"

func Test_ShouldCompareTwoNumbers(t *testing.T) {
	calculator := MyCalculator{}
	t.Parallel()
	t.Run("", func(tt *testing.T) {
		tt.Parallel()
		expected := 2
		actual := calculator.Add(1, 1)
		if actual != expected {
			tt.Fatalf("Expected %v to be %v", expected, actual)
		}
	})
}
```

You can write multiple _expectations_ in one test ...

```go
package calculator

import (
	"github.com/bearstonedev/gotest"
	"testing"
)

func Test_ShouldCompareTwoNumbers(t *testing.T) {
	calc := CreateCalculator()
	I := gotest.Expec(t)
	I.Expect(calc.IsGreaterThan(2, 1)).ToBe(true)
	I.Expect(calc.IsGreaterThan(1, 2)).ToBe(false)
}
```

...and you can name them, if you want to:

```go
package calculator

import (
	"github.com/bearstonedev/gotest"
	"testing"
)

func Test_ShouldCompareTwoNumbers(t *testing.T) {
	calc := CreateCalculator()
	I := gotest.Expec(t)
	I.Expect(calc.IsGreaterThan(2, 1)).As("2 > 1").ToBe(true)  // Test will be named: I_expect_2_>_1_to_be_true
	I.Expect(calc.IsGreaterThan(1, 2)).As("1 > 2").ToBe(false) // Test will be named: I_expect_1_>_2_to_be_false
}
```

Each expectation is run as a [subtest](https://go.dev/blog/subtests) of the containing test.

### :memo: Make it easier to write descriptive specifications of behaviour (BDD) by reducing boilerplate test code.

- All tests are run in parallel, both to reduce boilerplate and to help enforce independent tests.
- `t.FailNow`, `t.Log`, etc. are abstracted.

### :dancers: Make it easier to write parallel tests.

Each "expectation" is run as a distinct, parallel test. This helps ensure:

- Tests are **temporally decoupled**, meaning they don't need to be run in a specific order;
- Tests are **distinct**, making it easier to localize a specific failure; and,
- Tests are **safe**
  because [low-level test implementation details](https://gist.github.com/posener/92a55c4cd441fc5e5e85f27bca008721) are
  managed on your behalf.

### :children_crossing: Maintain compatibility with the existing Go toolset.

- Standard `go` commands should work as expected to automate testing.
- Example: `go test -run 'SomeTestSuite/specific_scenario' -v` should produce the same result as with native Go tests.

## Detailed Usage for Beginners

### What are tests?

Test are written to **specify expected behaviour**; we use them to measure and evaluate the software we write. Tests
should be written *in advance* to help guide design and implementation; this is known as **[test-driven development
(TDD)](https://martinfowler.com/bliki/TestDrivenDevelopment.html)**.

A test is a function which invokes a part of the software we're writing; this can be a function or method, a module, or
even an entire application. The test specifies the expected inputs (values we provide to our software) and outputs (
values we receive from our software) for a specific behaviour. A collection of tests, known as a **test suite**, helps
us identify problems with our design and/or implementation *during development* when those problems are easiest to
address.

### 1. Write a test function for a feature or behaviour.

This function should follow [standard Go test conventions](https://pkg.go.dev/testing).

```go
package pizza_ordering

import "testing"

func TestOrderingADeluxePizza(t *testing.T) {}
```

### 2. Use `Expec(t)` to wrap the `*testing.T` pointer.

This is the "starting point" for writing tests, and will be used to write one or more test cases; these tests should be
__cohesive__ (related to each other).

```go
package pizza_ordering

import (
	. "github.com/bearstonedev/gotest"
	"testing"
)

func TestOrderingADeluxePizza(t *testing.T) {
	I := Expec(t)
}
```

### 3. Set up the system-under-test (SUT).

Depending on what you're testing, you might have a type or module to be tested; this is colloquially known as the *
*system-under-test** (SUT). You should set it up with any relevant preconditions for the scenario you're testing.

```go
package pizza_ordering

import (
	. "github.com/bearstonedev/gotest"
	"testing"
)

func TestOrderingADeluxePizza(t *testing.T) {
	I := Expec(t)

	sut := &PizzaOrderMaker{location: "New York"}
}
```

### 4. Write one or more expectations (tests).

This will compare the expected output to the output produced by your program.

```go
package pizza_ordering

import (
	. "github.com/bearstonedev/gotest"
	"testing"
)

func TestOrderingADeluxePizza(t *testing.T) {
	I := Expec(t)

	sut := &PizzaOrderMaker{location: "New York"}
	expected := &PizzaOrder{
		toppings: []string{"tomato", "green olives", "red peppers"},
		locale:   "New York",
	}
	actual := sut.OrderPizza("tomato", "green olives", "red peppers")

	I.Expect(actual.toppings).ToBe(expected.toppings)
	I.Expect(actual.locale).ToBe(expected.locale)
}
```

### 5. Write the implementation to make the test pass.

Now that you have a failing test, you should implement the changes needed to make that test pass. Try to avoid writing
many tests upfront, or writing tests after-the-fact; writing tests and implementation iteratively will "drive" the
design of the software you're building in small increments.

### 6. Refactor.

> Refactoring (noun): a change made to the internal structure of software to make it easier to understand and cheaper to
> modify without changing its observable behavior.  
> Refactoring (verb): to restructure software by applying a series of refactorings without changing its observable
> behavior.  
https://martinfowler.com/bliki/DefinitionOfRefactoring.html

Refactoring must be done frequently to minimize accidental complexity. When building software, try working in cycles
of "Red-Green-Refactor":

1. Write a failing test;
2. Make that test pass;
3. Refactor the test and the implementation.

## Examples

Working examples are shown here: [gotest/examples](examples/tests-example_test.go).

(The examples are not bundled with the published module.)

## Version

The currently published version can be found [here](publish-update/published.version).