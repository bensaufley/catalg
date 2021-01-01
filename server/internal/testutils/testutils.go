package testutils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func AssertError(test *testing.T, want interface{}, got error) bool {
	if w, ok := want.(bool); ok {
		if w {
			return assert.Error(test, got)
		}
		return assert.NoError(test, got)
	}
	w := want.(error)
	if w == nil {
		return assert.NoError(test, got)
	}
	return assert.EqualError(test, got, w.Error())
}
