package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/leftslash/config"
)

// set environment before use:
// export KEY="welcome"

func main() {

	env := flag.String("env", "dev", "environment (e.g. dev, prod)")
	flag.Parse()

	conf, err := config.NewConfig()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	fmt.Printf("%s\n", conf.Get(*env, "db.file"))
	fmt.Printf("%s\n", conf.Get(*env, "db.passwd"))
	fmt.Printf("%d\n", conf.GetInt(*env, "net.port"))
}
