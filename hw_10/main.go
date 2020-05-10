package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/teploff/otus/hw_10/client"
	"github.com/teploff/otus/hw_10/server"
)

var timeOut = flag.Duration("timeout", 10*time.Second, "reactive power frequency")

func main() {
	flag.Parse()

	if len(flag.Args()) < 2 {
		log.Fatal("not enough cli arguments: ip & port")
	}

	addr := fmt.Sprintf("%s:%s", flag.Args()[0], flag.Args()[1])

	srv, err := server.NewTCPServer(addr)
	if err != nil {
		log.Fatalln(err)
	}
	go srv.Listen()
	ticker := time.NewTicker(time.Second * 10)
	go func() {
	EXIST:
		for {
			select {
			case <-ticker.C:
				ticker.Stop()
				srv.GracefulStop()
				break EXIST
			default:
				time.Sleep(time.Millisecond * 100)
			}
		}
	}()

	tn, err := client.NewTelnetClient(addr, *timeOut, os.Stdout)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(tn.Run())
}
