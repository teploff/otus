package main

import (
	"fmt"
	"otus/hw_3/analyzer"
)

func main() {
	analyz, err := analyzer.NewFrequencyAnalyzer("привет как дела деЛа тВои твОи Саша саша ")
	//analyz, err := analyzer.NewFrequencyAnalyzer("rivet privet kaKa kak dela tvoi sasha")
	if err != nil {
		panic(err)
	}
	fmt.Println(analyz.Search())
	//_, _ = Top10("привет как дела деЛа тВои твОи Саша саша ")
	//_, _ = Top10("Privet privet kaKa kak dela tvoi sasha")
}
