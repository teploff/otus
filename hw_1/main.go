package main

import (
	"fmt"
	"github.com/beevik/ntp"
	"os"
	"time"
)

const hostName = "0.beevik-ntp.pool.ntp.org"

func main() {
	currentTime := time.Now()
	fmt.Println("Текущее время:", currentTime)

	exactTime, err := ntp.Time(hostName)
	if err != nil {
		code, _ := os.Stderr.WriteString(err.Error())
		os.Exit(code)
	} else {
		fmt.Println("Точное время:", exactTime)
	}
}
