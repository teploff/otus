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
		inputRunes:       []rune(input),
		validationRegExp: validationRegExp,
		groupRegExp:      groupRegExp,
	}
}

// StringConverter performs a primitive string unpacking containing repeated characters / runes
type StringConverter struct {
	inputString      string
	inputRunes       []rune
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

// group grouping complex chars, for example a4 or \4 or \\ in input string
// Or if there is no special char in input string, all string in group
func (sc StringConverter) group() []string {
	groups := make([]string, 0)

	// if special chars not found return original string as single group
	matches := sc.groupRegExp.FindAllIndex([]byte(sc.inputString), -1)
	if matches == nil {
		groups = append(groups, sc.inputString)
		return groups
	}

	index := 0
	var indexInGroup bool
	for index < len(sc.inputRunes) {
		for _, group := range matches {
			if group[0] == index {
				groups = append(groups, sc.inputString[group[0]:group[1]])
				index = group[1]
				indexInGroup = true
				break
			}
		}

		if !indexInGroup {
			groups = append(groups, string(sc.inputRunes[index]))
			index++
		} else {
			indexInGroup = false
		}
	}

	return groups
}

// unpack extracting groups
func (sc StringConverter) unpack(groups []string) string {
	result := make([]string, 0)

	// if only string
	if len(groups) == 1 {
		return groups[0]
	}

	for _, group := range groups {
		runes := []rune(group)
		if len(runes) == 1 {
			result = append(result, string(runes))
		} else if (unicode.IsLetter(runes[0]) || runes[0] == '\\') && unicode.IsDigit(runes[1]) {
			tmp := ""
			digits, _ := strconv.Atoi(string(runes[1:]))
			for i := 0; i < digits; i++ {
				tmp += string(runes[0])
			}
			result = append(result, tmp)
		} else if len(runes) == 2 && runes[0] == '\\' && runes[1] == '\\' {
			result = append(result, string(runes[0]))
		} else if runes[0] == '\\' && runes[1] == '\\' && unicode.IsDigit(runes[2]) {
			tmp := ""
			digits, _ := strconv.Atoi(string(runes[2:]))
			for i := 0; i < digits; i++ {
				tmp += string(runes[0])
			}
			result = append(result, tmp)
		} else {
			result = append(result, string(runes))
		}
	}
	return strings.Join(result, "")
}

// Do launch string converter
func (sc StringConverter) Do() (string, error) {
	if err := sc.validate(); err != nil {
		return "", err
	}

	groups := sc.group()

	result := sc.unpack(groups)

	return result, nil
}
