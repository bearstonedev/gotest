# gotest

A simple test framework for Go, intended for TDD beginners and novice developers.

## Usage

- Use `Tests` to wrap the `testing.T` object.
- `Test` will create an individual, parallelized test. It can be chained to create multiple tests at once.
- `Scenarios` sets up a parameterized test.
- `Assertions` provide methods for asserting post-conditions.

### Examples

See examples in [gotest/examples](examples/tests-example_test.go).
