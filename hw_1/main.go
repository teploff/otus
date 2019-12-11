package main

import (
	"fmt"
	"github.com/beevik/ntp"
	"log"
	"time"
)

const hostName = "0.beevik-ntp.pool.ntp.org"

func main() {
	currentTime := time.Now()
	fmt.Println("Текущее время:", currentTime)

	exactTime, err := ntp.Time(hostName)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println("Точное время:", exactTime)
}
