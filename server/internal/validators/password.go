package validators

import (
	"errors"
	"strings"
)

var letters = func()string {
	str := strings.Builder{}
	for c := 'A'; c <= 'Z'; c++ {
		str.WriteRune(c)
		str.WriteRune(c + 'a' - 'A')
	}
	return str.String()
}()

// Password returns an error if a password is not valid
func Password(password string) error {
	if password == "" {
		return errors.New("password is blank")
	}

	if len(password) < 8 {
		return errors.New("password is too short")
	}

	if len(password) > 128 {
		return errors.New("password is too long")
	}

	if !strings.ContainsAny(password, letters) {
		return errors.New("password contains no letters")
	}

	if !strings.ContainsAny(password, "1234567890`~!@#$%^&*()_+-=,./;'[]\\<>?:\"{}| ") {
		return errors.New("password contains no numbers or symbols")
	}

	return nil
}
