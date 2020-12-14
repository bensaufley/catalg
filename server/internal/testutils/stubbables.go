package testutils

import (
	"time"

	"github.com/bensaufley/catalg/server/internal/stubbables"
)

type CallLog struct {
	Calls [][]interface{}
}

func (c *CallLog) AddCall(call []interface{}) {
	c.Calls = append(c.Calls, call)
}

func StubTimeNow(t time.Time, log ...CallLog) func() {
	orig := stubbables.TimeNow

	stubbables.TimeNow = func() time.Time {
		if len(log) > 0 {
			log[0].AddCall([]interface{}{nil})
		}
		return t
	}

	return func() {
		stubbables.TimeNow = orig
	}
}

func StubUUIDv1(uuid string, log ...CallLog) func() {
	orig := stubbables.UUIDv1

	stubbables.UUIDv1 = func() string {
		if len(log) > 0 {
			log[0].AddCall([]interface{}{nil})
		}
		return uuid
	}

	return func() {
		stubbables.UUIDv1 = orig
	}
}

func StubArgon2IDKey(digest string, log ...CallLog) func() {
	orig := stubbables.Argon2IDKey

	stubbables.Argon2IDKey = func(password []byte, salt []byte, time uint32, memory uint32, threads uint8, keyLen uint32) []byte {
		if len(log) > 0 {
			log[0].AddCall([]interface{}{password, salt, time, memory, threads, keyLen})
		}
		return []byte(digest)
	}

	return func() {
		stubbables.Argon2IDKey = orig
	}
}

func StubGetEnvWithDefault(val string, log ...CallLog) func() {
	orig := stubbables.GetEnvWithDefault

	stubbables.GetEnvWithDefault = func (key string, dflt string) string {
		if len(log) > 0 {
			log[0].AddCall([]interface{}{key, dflt})
		}
		return val
	}

	return func() {
		stubbables.GetEnvWithDefault = orig
	}
}

func StubMustGetEnv(val *string, log ...CallLog) func() {
	orig := stubbables.MustGetEnv

	stubbables.MustGetEnv = func(key string) string {
		if len(log) > 0 {
			log[0].AddCall([]interface{}{key})
		}
		if val == nil {
			return ""
		} else {
			return *val
		}
	}

	return func() {
		stubbables.MustGetEnv = orig
	}
}
