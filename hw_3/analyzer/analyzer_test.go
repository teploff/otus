package analyzer

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// TestCase when an empty string is passed to the frequency analyzer
func TestEmptyInputString(t *testing.T) {
	emptyString := ""
	fa, err := NewFrequencyAnalyzer(emptyString)
	if err != nil {
		t.Fatal(err.Error())
	}

	actualWords := fa.Search()
	assert.Nil(t, actualWords)
}

// TestCase when a string, which doesn't contain the words, is passed to the frequency analyzer
func TestInputStringWithoutWords(t *testing.T) {
	inputString := "± § > < 1 ! 2 @ 3 # № 5 % : 6 ^ , 7 & . 8 * ; 9 ( 0 ) - _ = + ] [ ` ~  "
	fa, err := NewFrequencyAnalyzer(inputString)
	if err != nil {
		t.Fatal(err.Error())
	}

	actualWords := fa.Search()
	assert.Nil(t, actualWords)
}

// TestCase when a string, which contains unique words, is passed to the frequency analyzer
func TestAllWordsUnique(t *testing.T) {
	inputString := "Do you know what this word means?"
	fa, err := NewFrequencyAnalyzer(inputString)
	if err != nil {
		t.Fatal(err.Error())
	}

	actualWords := fa.Search()
	assert.Empty(t, actualWords)
}

// TestCase when a string, which contains only one word occurring twice, is passed to the frequency analyzer
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

// TestCase when a string, which contains repeated words that differ in case, is passed to the frequency analyzer
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

// TestCase when a string, which contains more than 10 repeated words that differ in case, is passed to the frequency
// analyzer
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

// TestCase when a string, which contains frequently encountered words, is passed to the frequency analyzer
func TestSortingByFrequentlyOccurringWords(t *testing.T) {
	inputString := "UnusUaL UNUsUaL UnUsUaL UnUsUaL UnUsUaL UnUsUaL UnUsUaL UNUsUaL UnUsUal UnUsUaL UnUsUal." +
		"iTs iTs iTs ITS iTs iTs iTs iTs iTs iTS." +
		"oF Of of Of OF Of oF Of of." +
		"bEcAuSe bEcAuSe bEcAuSe bEcAuSe bEcAuSe bEcAuSe bEcAuSe bEcAuSe." +
		"uniQue unIque uniquE Unique uniquE unIQue unIQue." +
		"Word Word Word Word Word WoRd." +
		"THIS THIS THIS THIS THIS." +
		"FinD fIND finD finD." +
		"MaY mAy MAy." +
		"You YOu"
	fa, err := NewFrequencyAnalyzer(inputString)
	if err != nil {
		t.Fatal(err.Error())
	}

	actualWords := fa.Search()
	expectedWords := []string{"unusual", "its", "of", "because", "unique", "word", "this", "find", "may", "you"}
	assert.Equal(t, expectedWords, actualWords)
}
