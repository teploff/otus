package main

import (
	"fmt"
	"regexp"
	"sort"
	"strings"
)

func main() {
	_, _ = Top10("привет как дела деЛа тВои твОи Саша саша ")
	//_, _ = Top10("Privet privet kaKa kak dela tvoi sasha")
}

func Top10(input string) ([]string, error) {
	wordRegExp, err := regexp.Compile(`[a-zA-Zа-яА-я]+`)
	if err != nil {
		return []string{}, err
	}

	words := wordRegExp.FindAllString(input, -1)
	if words == nil {
		return []string{}, fmt.Errorf("not found")
	}

	uniqueWords := make(Pairs, 0, len(words))
	// To Lower case
	for _, word := range words {
		lowerCaseWord := strings.ToLower(word)
		uniqueWords.Append(lowerCaseWord)
	}

	sort.Sort(sort.Reverse(uniqueWords))
	for _, pair := range uniqueWords {
		fmt.Println(pair)
	}

	return []string{}, nil
}
