package validation

import (
	"strings"
)

var letters = func() string {
	str := strings.Builder{}
	for c := 'A'; c <= 'Z'; c++ {
		str.WriteRune(c)
		str.WriteRune(c + 'a' - 'A')
	}
	return str.String()
}()

// Password returns an error if a password is not valid
func Password(password string) *Error {
	if password == "" {
		return NewError("password", "cannot be blank")
	}

	if len(password) < 8 {
		return NewError("password", "cannot be shorter than 8 characters")
	}

	if len(password) > 128 {
		return NewError("password", "cannot be longer than 128 characters")
	}

	if !strings.ContainsAny(password, letters) {
		return NewError("password", "must contain letters")
	}

	if !strings.ContainsAny(password, "1234567890`~!@#$%^&*()_+-=,./;'[]\\<>?:\"{}| ") {
		return NewError("password", "must contain numbers or symbols")
	}

	return nil
}
