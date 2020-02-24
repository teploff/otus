package main

import (
	"github.com/teploff/hw_7/envdir"
	"log"
	"os"
)

func main() {

	args := os.Args
	if len(args) < 3 {
		log.Fatalln("not enough passing args")
	}

	env, err := envdir.ReadDir(args[1])
	if err != nil {
		log.Fatalln(err)
	}

	if exitCode := envdir.RunCmd(args[2:], env); exitCode != 0 {
		log.Fatalln("fail to exec command")
	}
}
