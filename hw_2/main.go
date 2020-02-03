package main

import (
	"fmt"
	"github.com/teploff/otus/hw_2/converter"
	"log"
)

func main() {
	var input string
	if _, err := fmt.Scanln(&input); err != nil {
		log.Fatal(err)
	}

	stringConverter := converter.NewStringConverter(input)
	result, err := stringConverter.Do()
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(result)
}
