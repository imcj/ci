package env

import (
	"github.com/aos-stack/env"
)

type RealEnvProvider struct {

}

func (e RealEnvProvider)Get() int {
	return 1
}

func (e RealEnvProvider)GetLabel() string {
	return "develop"
}

func SetupEnv() {
	env.SetProvider(RealEnvProvider{})
}