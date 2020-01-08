package analyzer

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEmptyInputString(t *testing.T) {
	emptyString := ""
	fa, err := NewFrequencyAnalyzer(emptyString)
	if err != nil {
		t.Fatal(err.Error())
	}

	actualWords := fa.Search()
	assert.Nil(t, actualWords)
}

func TestInputStringWithoutWords(t *testing.T) {
	inputString := "± § > < 1 ! 2 @ 3 # № 5 % : 6 ^ , 7 & . 8 * ; 9 ( 0 ) - _ = + ] [ ` ~  "
	fa, err := NewFrequencyAnalyzer(inputString)
	if err != nil {
		t.Fatal(err.Error())
	}

	actualWords := fa.Search()
	assert.Nil(t, actualWords)
}

func TestAllWordsUnique(t *testing.T) {
	inputString := "Do you know what this word means?"
	fa, err := NewFrequencyAnalyzer(inputString)
	if err != nil {
		t.Fatal(err.Error())
	}

	actualWords := fa.Search()
	assert.Empty(t, actualWords)
}

func TestOnlyOneWordOccursTwice(t *testing.T) {
	inputString := "Do you know what this word means? Do"
	fa, err := NewFrequencyAnalyzer(inputString)
	if err != nil {
		t.Fatal(err.Error())
	}

	actualWords := fa.Search()
	expectedWords := []string{"do"}
	assert.Equal(t, expectedWords, actualWords)
}

func TestWordsWithLowerAndUpperCases(t *testing.T) {
	inputString := "Do you know what this word means? dO You kNow whaT thIs Word meanS?"
	fa, err := NewFrequencyAnalyzer(inputString)
	if err != nil {
		t.Fatal(err.Error())
	}

	actualWords := fa.Search()
	expectedWords := []string{"do", "you", "know", "what", "this", "word", "means"}
	assert.Equal(t, expectedWords, actualWords)
}

func TestMoreThanTenFrequencyWords(t *testing.T) {
	inputString := "You may find this word unique because of its unusual spelling. " +
		"YOU MAY FIND THIS WORD UNIQUE BECAUSE OF ITS UNUSUAL SPELLING. " +
		"yOu mAy fInD ThIs wOrD UnIqUe bEcAuSe oF ItS UnUsUaL SpElLiNg"
	fa, err := NewFrequencyAnalyzer(inputString)
	if err != nil {
		t.Fatal(err.Error())
	}

	actualWords := fa.Search()
	expectedWords := []string{"you", "may", "find", "this", "word", "unique", "because", "of", "its", "unusual"}
	assert.Equal(t, expectedWords, actualWords)
}
