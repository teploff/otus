package analyzer

import (
	"regexp"
	"sort"
	"strings"
)

const maxWordAmount = 10

// NewFrequencyAnalyzer get FrequencyAnalyzer instance
func NewFrequencyAnalyzer(inputText string) (FrequencyAnalyzer, error) {
	wordRegExp, err := regexp.Compile(`[a-zA-Zа-яА-я]+`)
	if err != nil {
		return FrequencyAnalyzer{}, err
	}

	return FrequencyAnalyzer{
		inputText:  inputText,
		wordRegExp: wordRegExp,
	}, nil
}

// FrequencyAnalyzer accepts a text string as input (inputText) and returns a slice with the 10 most frequently
// encountered words in the text (When call method Search).
type FrequencyAnalyzer struct {
	inputText  string
	wordRegExp *regexp.Regexp
}

// Search method returns a slice with the 10 most frequently encountered words in the text
func (a FrequencyAnalyzer) Search() []string {
	words := a.wordRegExp.FindAllString(a.inputText, -1)
	if words == nil {
		return nil
	}

	uniqueWords := make(Pairs, 0, len(words))
	// To Lower case
	for _, word := range words {
		lowerCaseWord := strings.ToLower(word)
		uniqueWords.Append(lowerCaseWord)
	}

	sort.Sort(sort.Reverse(uniqueWords))
	uniqueWords.DeleteUniqueWords()

	result := make([]string, 0, len(uniqueWords))
	for i := 0; i < len(uniqueWords) && i < maxWordAmount; i++ {
		result = append(result, uniqueWords[i].Word)
	}

	return result
}
