# golang-test-framework

A simple test framework for Go.

## Usage

Use the `Suite` object to wrap the `testing.T` object.

Tests are created from the `Suite` object.

### Example

```go
func TestSomeTests(t *testing.T) {
	Suite(t).Test("1 + 1 should be 2", func() bool {
		return 1+1 == 2
	}).Test("1 + 2 should not be 1", func() bool {
		return 1+2 != 1
	})
}
```

### Support

This is a work-in-progress and will not be actively supported.