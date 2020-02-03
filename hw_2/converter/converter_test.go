package converter

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// Test case with incorrect input string with all digit cases
func TestInvalidInputStringWithDigits(t *testing.T) {
	inputStrings := []string{
		"4abc",
		"4563",
	}

	for _, input := range inputStrings {
		converter := NewStringConverter(input)
		out, err := converter.Do()
		assert.Empty(t, out)
		assert.Error(t, err)
	}
}

// Test case with incorrect input string with all Backslash cases
func TestInvalidInputStringWithBackslash(t *testing.T) {
	inputStrings := []string{
		"\\",
		"\\abc",
		"abc\\",
		"344\\",
	}

	for _, input := range inputStrings {
		converter := NewStringConverter(input)
		out, err := converter.Do()
		assert.Empty(t, out)
		assert.Error(t, err)
	}
}

// Test case with correct input string with all digit cases
func TestValidInputStringWithDigits(t *testing.T) {
	testData := []struct {
		in  string
		out string
	}{
		{"abed", "abed"},
		{"a4bc2d5e", "aaaabccddddde"},
	}

	for _, data := range testData {
		converter := NewStringConverter(data.in)
		out, err := converter.Do()
		assert.Nil(t, err)
		assert.Equal(t, out, data.out)
	}
}

// Test case with correct input string with all Backslash cases
func TestValidInputStringWithBackslash(t *testing.T) {
	testData := []struct {
		in  string
		out string
	}{
		{"qwe\\4\\5", "qwe\\\\\\\\\\\\\\\\\\"},
		{"qwe\\4", "qwe\\\\\\\\"},
		{"\\\\", "\\\\"},
	}

	for _, data := range testData {
		converter := NewStringConverter(data.in)
		out, err := converter.Do()
		assert.Nil(t, err)
		assert.Equal(t, out, data.out)
	}
}
