package converter

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

var validationRegExp *regexp.Regexp
var groupRegExp *regexp.Regexp

func init() {
	validationRegExp = regexp.MustCompile(`((^|[^\\])\\([^\\\d]|$))|^\d+`)
	groupRegExp = regexp.MustCompile(`(\\{2}\d+)|(\\{2})|(\\\d{2,})|(\D\d+)`)
}

// NewStringConverter returns StringConverter instance
func NewStringConverter(input string) StringConverter {
	return StringConverter{
		inputString:      input,
		validationRegExp: validationRegExp,
		groupRegExp:      groupRegExp,
	}
}

// StringConverter performs a primitive string unpacking containing repeated characters / runes
type StringConverter struct {
	inputString      string
	validationRegExp *regexp.Regexp
	groupRegExp      *regexp.Regexp
}

// validate validates input string
func (sc StringConverter) validate() error {
	if match := sc.validationRegExp.MatchString(sc.inputString); match {
		return fmt.Errorf("invalid string")
	}

	return nil
}

// do grouping complex chars, for example a4 or \4 or \\ in input string
// Or if there is no special char in input string, all string in do
func (sc StringConverter) do() string {
	result := make([]string, 0)

	// if special chars not found return original string as single do
	matches := sc.groupRegExp.FindAllIndex([]byte(sc.inputString), -1)
	if matches == nil {
		return sc.inputString
	}

	// else find all special do by indexes are find by regExp from input string
	currentIndex := 0
	for _, match := range matches {
		if currentIndex < match[0] {
			result = append(result, sc.inputString[currentIndex:match[0]])
		}
		unpackedString := sc.unpack(sc.inputString[match[0]:match[1]])
		result = append(result, unpackedString)
		currentIndex = match[1]
	}

	if currentIndex != len(sc.inputString) {
		result = append(result, sc.inputString[currentIndex:])
	}

	return strings.Join(result, "")
}

// unpack extracting groups
func (sc StringConverter) unpack(group string) string {
	unpackedString := make([]string, 0)

	runes := []rune(group)
	if unicode.IsDigit(runes[1]) {
		digits, _ := strconv.Atoi(string(runes[1:]))
		tmp := sc.extendLine(string(runes[0]), digits)
		unpackedString = append(unpackedString, tmp)
	} else if len(runes) == 2 {
		unpackedString = append(unpackedString, string(runes[0]))
	} else {
		digits, _ := strconv.Atoi(string(runes[2:]))
		tmp := sc.extendLine(string(runes[0]), digits)
		unpackedString = append(unpackedString, tmp)
	}

	return strings.Join(unpackedString, "")
}

// extendLine returns characters char count times
func (sc StringConverter) extendLine(char string, count int) string {
	out := ""
	for i := 0; i < count; i++ {
		out += char
	}

	return out
}

// Do launch string converter
func (sc StringConverter) Do() (string, error) {
	if err := sc.validate(); err != nil {
		return "", err
	}

	result := sc.do()

	return result, nil
}
