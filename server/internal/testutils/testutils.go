package testutils

import (
	"os"
	"testing"
)

// TestWrapper is a type for middleware-style wrappers to
// be used in TestMain
type TestWrapper func(func() int) func() int

// PrepareDB is a TestWrapper sets up and tears
// down the test database
func PrepareDB(cb func() int) func() int {
	return func() int {
		code := cb()

		return code
	}
}

// WrapTests receives a *testing.M and a list of wrappers
// to apply from innermost to outermost
func WrapTests(m *testing.M, wrappers ...TestWrapper) {
	cb := m.Run
	for _, wrapper := range wrappers {
		cb = wrapper(cb)
	}
	code := cb()
	os.Exit(code)
}
