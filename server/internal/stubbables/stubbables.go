package stubbables

import (
	"math/rand"
	"os"
	"strings"
	"time"

	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/argon2"

	"github.com/bensaufley/catalg/server/internal/log"
)

// stubbables provides methods that can be easily stubbed for testing

// TimeNow wraps time.Now
var TimeNow = time.Now

// UUIDv1 wraps uuid.NewV1().String()
var UUIDv1 = func() string {
	return uuid.NewV1().String()
}

var characters = func() []rune {
	chars := make([]rune, 126-33+1)
	for c := 33; c <= 126; c++ {
		chars[c-33] = rune(c)
	}
	return chars
}()

var charLen = len(characters)

// RandomChars returns a random set of characters of length ln
var RandomChars = func(ln int) string {
	rnd := rand.New(rand.NewSource(TimeNow().Unix()))
	str := &strings.Builder{}
	for i := 0; i < ln; i++ {
		str.WriteRune(characters[rnd.Intn(charLen)])
	}
	return str.String()
}

// Argon2IDKey wraps argon2.IDKey
var Argon2IDKey = argon2.IDKey

// GetEnvWithDefault wraps os.Getenv and accepts a default fallback
// if the environment variable is blank or unset
var GetEnvWithDefault = func(key string, dflt string) string {
	val := os.Getenv(key)
	if val == "" {
		log.Warnf("environment variable %s was empty; defaulting to %s", key, dflt)
		return dflt
	}
	return val
}

// MustGetEnv wraps os.Getenv and logs fatal (which exits) if the
// environment variable is blank or unset
var MustGetEnv = func(key string) string {
	val := os.Getenv(key)
	if val != "" {
		return val
	}
	log.Fatalf("environment variable %s was empty", key)
	return ""
}
