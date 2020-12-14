package auth

import (
	"crypto/subtle"
	"os"

	"github.com/bensaufley/catalg/server/internal/log"
	"github.com/bensaufley/catalg/server/internal/stubbables"
)

var pepper string

func init() {
	pepper = os.Getenv("PASSWORD_PEPPER")
	if pepper == "" {
		log.Fatal("could not initialize package auth: PASSWORD_PEPPER is empty")
	}
}

func HashPassword(password string) (string, string) {
	salt := stubbables.RandomChars(32)
	hash := stubbables.Argon2IDKey([]byte(pepper+password), []byte(salt), 1, 64*1024, 4, 72)
	return string(hash), salt
}

func ComparePassword(input string, digest string, salt string) bool {
	hash := stubbables.Argon2IDKey([]byte(pepper +input), []byte(salt), 1, 64 *1024, 4, 72)
	return subtle.ConstantTimeCompare([]byte(digest), hash) == 1
}
