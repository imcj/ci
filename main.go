package main

import (
	env_setup "github.com/aos-stack/ci/env"
	"github.com/aos-stack/env"
	"fmt"
)

func main() {
	env_setup.SetupEnv()
	fmt.Println(env.GetLabel())
}