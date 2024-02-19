# gotest

A simple test framework for Go, intended for TDD beginners and novice developers.

## Goals

### :rabbit2: Make it easier to write tests first (TDD) by making tests easier to write.

- An `Assertions` object provides simple test assertions.
- A `Scenarios` interface helps provide a structured way to write parameterized test cases.

### :memo: Make it easier to write descriptive specifications of behaviour (BDD) by reducing boilerplate test code.

- All tests are run in parallel, both to reduce boilerplate and to help enforce independent tests.
- `t.FailNow`, `t.Log`, etc. are abstracted.

### :children_crossing: Maintain maximal compatibility with the existing Go toolset.

- Standard `go` commands should work as expected to automate testing.
- Example: `go test -run 'SomeTestSuite/specific_scenario' -v` should produce the same result as with native Go tests.

## Usage

### 1. Write a "parent" test for a feature or behaviour.

This will contain one or more _related_ tests focused on a behaviour or outcome, and it should adhere
to [standard Go test conventions](https://pkg.go.dev/testing):

```golang
package pizza_ordering

import (
	"testing"
)

func TestOrderingADeluxePizza(t *testing.T) {}

```

### 2. Use `Tests(t)` to wrap the `*testing.T` pointer.

This is the "starting point" for writing tests, and will be used to write one or more test cases; these tests should be
__cohesive__ (related to each other).

```golang
package pizza_ordering

import (
	. "github.com/bearstonedev/gotest"
	"testing"
)

func TestOrderingADeluxePizza(t *testing.T) {
	Tests(t)
}

```

### 3. Write the tests.

There are a few options for implementing test cases, depending on what you're testing:

#### a. `Test()` will create a standalone parallelized test.

`Test()` can be chained to create multiple tests, like this:

```go
package pizza_ordering

import (
	. "github.com/bearstonedev/gotest"
	"testing"
)

func TestOrderingADeluxePizza(t *testing.T) {
	order := "deluxe pizza"
	Tests(t).
		Test("the pizza should be fully-cooked", func(shouldBe Assertions) {
			pizza := systemUnderTest.orderPizza(order)
			shouldBe.True(pizza.IsCooked())
		}).
		Test("the pizza should have deluxe toppings", func(shouldBe Assertions) {
			pizza := systemUnderTest.orderPizza(order)
			shouldBe.Equal(pizza.Toppings(), []string{"green pepper", "mushroom", "tomato"})
		})
}

```

#### b. `Scenarios()` sets up a parameterized test.

One or more `Scenario()`s should be specified, followed by a `Test()` call to define the test case for each scenario,
like this:

```go
package pizza_ordering

import (
	. "github.com/bearstonedev/gotest"
	"testing"
)

func TestOrderingAPizzaWithCustomToppings(t *testing.T) {
	Tests(t).
		Scenarios().
		Scenario("one topping", []string{"green pepper"}).
		Scenario("two toppings", []string{"green pepper", "tomatoes"}).
		Test(func(shouldBe Assertions, parameters []any) {
			order := parameters[0].([]string)
			pizza := systemUnderTest.orderPizza(order)
			shouldBe.Equal(pizza.Toppings(), order)
		})
}

```

#### The `Assertions` object provides methods for asserting post-conditions.

`Assertions` is injected with test-specific information via the test case:

```go
package pizza_ordering

import (
	. "github.com/bearstonedev/gotest"
	"testing"
)

func TestOrderingAPizzaWithCustomCrust(t *testing.T) {
	order := "thin crust"
	Tests(t).
		Test("the pizza should have the expected crust", func(shouldBe Assertions) {
			pizza := systemUnderTest.orderPizza(order)
			shouldBe.Equal(pizza.Crust(), "thin")
		})
}

```

## Examples

Working examples are shown here: [gotest/examples](examples/tests-example_test.go).

(The examples are not bundled with the published module.)

## Version

The currently published version can be found [here](scripts/published.version).