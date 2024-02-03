# golang-test-framework

A simple test framework for Go, intended for TDD beginners and novice developers.

## Usage

- Use `Tests` to wrap the `testing.T` object.
- `Test` will create an individual, parallelized test.
- `Scenarios` sets up a parameterized test.
- `Assertions` provide methods for asserting post-conditions.

### Examples

See examples in [golang-test-framework-example](golang-test-framework-example/tests-example_test.go).
