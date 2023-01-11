package wmi

import "github.com/vela-security/vela-public/assert"

func WithEnv(env assert.Environment) {
	env.Warn("not support wmi with linux")
}
