package main

import (
	"fmt"
	"github.com/beevik/ntp"
	"os"
	"time"
)

func main() {
	currentTime := time.Now()
	fmt.Println("Текущее время:", currentTime)

	exactTime, err := ntp.Time("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		code, _ := os.Stderr.WriteString(err.Error())
		os.Exit(code)
	} else {
		fmt.Println("Точное время:", exactTime)
	}
}
