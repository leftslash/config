package main

import (
	"fmt"

	"github.com/leftslash/config"
)

// set environment before use:
// export KEY="welcome"

func main() {

	conf := config.NewConfig()
	conf.Flag("config", "configuration file")
	conf.Flag("env", "environment (e.g. dev, prod)")
	conf.Load()

	env := conf.Get("env")
	fmt.Printf("%s\n", conf.Get(env, "db.file"))
	fmt.Printf("%s\n", conf.Get(env, "db.passwd"))
	fmt.Printf("%d\n", conf.GetInt(env, "net.port"))
}
