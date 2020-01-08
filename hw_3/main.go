package main

import (
	"fmt"
	"github.com/teploff/otus/hw_3/analyzer"
)

func main() {
	// Usage example
	frequencyAnalyzer, err := analyzer.NewFrequencyAnalyzer("Hello heLLo otus Otus otUs")
	if err != nil {
		panic(err)
	}
	fmt.Println(frequencyAnalyzer.Search())
}
