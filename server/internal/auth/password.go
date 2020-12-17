package auth

import (
	"crypto/subtle"
	"encoding/base64"
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
	return base64.URLEncoding.EncodeToString(hash), salt
}

func ComparePassword(input string, digest string, salt string) bool {
	hash := stubbables.Argon2IDKey([]byte(pepper+input), []byte(salt), 1, 64*1024, 4, 72)
	bts, err := base64.URLEncoding.DecodeString(digest)
	if err != nil {
		log.WithError(err).Error("error decoding password digest from base64")
		return false
	}
	return subtle.ConstantTimeCompare(bts, hash) == 1
}
