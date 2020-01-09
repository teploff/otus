package main

import (
	"fmt"
	"github.com/teploff/otus/hw_3/analyzer"
)

func main() {
	// Usage example
	frequencyAnalyzer := analyzer.NewFrequencyAnalyzer("hello hello my my world")
	fmt.Println(frequencyAnalyzer.Search())
}
