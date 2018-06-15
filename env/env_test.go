package env

import (
	"testing"
	"github.com/aos-stack/env"
	"github.com/stretchr/testify/assert"
)

func TestEnv(t *testing.T) {
	SetupEnv()
	assert.Equal(t, env.Develop, env.Get())
}
