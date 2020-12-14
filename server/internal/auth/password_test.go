package auth_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/bensaufley/catalg/server/internal/auth"
)

func TestPasswordBehavior(test *testing.T) {
	digest, salt := auth.HashPassword("test password 1")

	assert.Len(test, digest, 72)
	assert.Len(test, salt, 32)

	got := auth.ComparePassword("test password 2", digest, salt)

	assert.False(test, got)

	got = auth.ComparePassword("test password 1", digest, salt)

	assert.True(test, got)
}
