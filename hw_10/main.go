package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/teploff/otus/hw_10/client"
)

var timeOut = flag.Duration("timeout", 10*time.Second, "reactive power frequency")

func main() {
	flag.Parse()

	if len(flag.Args()) < 2 {
		log.Fatal("not enough cli arguments: ip & port")
	}

	addr := fmt.Sprintf("%s:%s", flag.Args()[0], flag.Args()[1])
	tn, err := client.NewTelnetClient(addr, *timeOut)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(tn.Run())
}
