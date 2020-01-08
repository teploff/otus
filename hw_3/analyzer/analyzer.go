package analyzer

import (
	"regexp"
	"sort"
	"strings"
)

func NewAnalyzer(inputText string) (Analyzer, error) {
	wordRegExp, err := regexp.Compile(`[a-zA-Zа-яА-я]+`)
	if err != nil {
		return Analyzer{}, err
	}

	return Analyzer{
		inputText:  inputText,
		wordRegExp: wordRegExp,
	}, nil
}

type Analyzer struct {
	inputText  string
	wordRegExp *regexp.Regexp
}

func (a Analyzer) Do() []string {
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

	result := make([]string, 0, len(uniqueWords))
	for i := 0; i < len(uniqueWords)-1 || i < 10; i++ {
		result = append(result, uniqueWords[i].Word)
	}

	return result
}
