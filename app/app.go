package app

import (
	"errors"
	"regexp"
)

const (
	strinG                               string = `"([^"\n\t\\]*?(\\"|\\\t|\\\\|\\b|\\f|\\n|\\r|\\t|\\/|\\u)+[^"\n\t\\]*?)+"|"[^"\n\t\\]*"`
	number                               string = `-?\d{1}\.\d+([eE][-+]?)\d+|-?[1-9]\d+\.\d+([eE][-+]?)\d+|-?[1-9]\d*([eE][-+]?)\d+|-?\d{1}\.\d+|-?[1-9]\d+\.\d+|-?[1-9]\d*|-?0([eE][-+]?\d+){0,1}`
	innerBrackets                        string = `\[[^][]*\]|{[^}{]*}|\[.*\[.*\].*\]|\{.*\{.*\}.*\}`
	stringValues                         string = `|` + strinG + `|`
	innerElement                         string = `\s*(null|true|false|` + number + stringValues + innerBrackets + `){1}`
	lastElementInOuterSqurareBrackets    string = `(` + innerElement + `\s*)`
	multipleElmentsInOuterSquareBrackets string = `(` + innerElement + `\s*,\s*)*`
	outerSquareBrackets                  string = `\[\s*(` + multipleElmentsInOuterSquareBrackets + lastElementInOuterSqurareBrackets + `{1}){0,1}\]`
	objectKey                            string = `(` + strinG + `)`
	lastElementInOuterCurlyBrackets      string = `(\s*` + objectKey + `\s*:` + innerElement + `\s*)`
	multipleElmentsInOuterCurlyBrackets  string = `(\s*` + objectKey + `\s*:` + innerElement + `\s*,\s*)*`
	outerCurlyBrakets                    string = `{\s*(` + multipleElmentsInOuterCurlyBrackets + lastElementInOuterCurlyBrackets + `{1}){0,1}}`
	validJSONPattern                     string = `(?s)\A\s*(` + strinG + `|` + number + `|false|null|true|` + outerSquareBrackets + `|` +
		outerCurlyBrakets + `){1}\s*\z`
)

var validJSONregex *regexp.Regexp = regexp.MustCompile(validJSONPattern)

func App(readFile func(name string) ([]byte, error), args []string) (string, error) {
	fileContentInByteArray, err := readFile(args[1])
	if err != nil {
		return "", err
	}
	fileContentString := string(fileContentInByteArray)
	if !validJSONregex.MatchString(fileContentString) {
		return "", errors.New(produceAReasonForInvalidation(fileContentString))
	}
	bracketsIndices := getBracketsIndices(fileContentString)
	if len(bracketsIndices) > 2 { // Becuase we should not count start and end brackets of the json
		return handleInnerBrackets(fileContentString, bracketsIndices)
	}
	return "This is a valid JSON", nil
}

func getBracketsIndices(stringContent string) [][]int {
	bracketsRegex := regexp.MustCompile(`\{|\}|\[|\]`)
	bracketsIndices := bracketsRegex.FindAllStringIndex(stringContent, -1)
	if len(bracketsIndices) > 0 {
		bracketsIndices = removeBracketsThatAreInStrings(stringContent, bracketsIndices)
	}
	return bracketsIndices
}

func handleInnerBrackets(innerString string, bracketsIndices [][]int) (string, error) {
	OpenningBracketsIndexes := make([]int, 0)
	for k := range bracketsIndices {
		if k != 0 && k != len(bracketsIndices)-1 {
			value := innerString[bracketsIndices[k][0]:bracketsIndices[k][1]]
			if value == "[" || value == "{" {
				OpenningBracketsIndexes = append(OpenningBracketsIndexes, bracketsIndices[k][0])
			} else { // ] or }
				if len(OpenningBracketsIndexes) > 0 {
					starting := OpenningBracketsIndexes[len(OpenningBracketsIndexes)-1]
					ending := bracketsIndices[k][1]
					innerObjectOrArray := innerString[starting:ending]
					if !validJSONregex.MatchString(innerObjectOrArray) {
						return "", errors.New(produceAReasonForInvalidation(innerObjectOrArray))
					}
					OpenningBracketsIndexes = OpenningBracketsIndexes[:len(OpenningBracketsIndexes)-1]
				} else {
					return "", errors.New("This is an invalid JSON\nThis is an array that is surrounded by invalid \"][}{\"")
				}
			}
		}
	}
	if len(OpenningBracketsIndexes) > 0 {
		return "", errors.New("This is an invalid JSON\nThis is an array that is surrounded by invalid \"][}{\"")
	}
	return "This is a valid JSON", nil
}

func removeBracketsThatAreInStrings(innerString string, indices [][]int) [][]int {
	stringValuesRegex := regexp.MustCompile(strinG)
	stringValuesIndices := stringValuesRegex.FindAllStringIndex(innerString, -1)
	if len(stringValuesIndices) > 0 {
		if indices[len(indices)-1][0] < (stringValuesIndices[0][1]-1) ||
			indices[0][0] > (stringValuesIndices[len(stringValuesIndices)-1][1]) {
			return indices
		}
		var revisedIndices [][]int = make([][]int, 0)
		for _, v := range indices {
			found := false
			for _, v2 := range stringValuesIndices {
				if v[0] > v2[0] && v[1] < v2[1] {
					found = true
					break
				}
			}
			if !found {
				revisedIndices = append(revisedIndices, v)
			}
		}
		return revisedIndices
	}
	return indices
}

func produceAReasonForInvalidation(fileContentString string) string {
	var invalid string = "This is an invalid JSON"
	if isThereNoObjectOrArray(fileContentString) {
		return invalid + "\nMUST be an object, array, number, or string, or false or null or true"
	}
	if isALeadedZeroNumber(fileContentString) {
		return invalid + "\nAn invalid number, there is a leading zero"
	}
	if multipleValuesOutsidAnObjectOrArray(fileContentString) {
		return invalid + "\nMultiple values outside of an object or array"
	}
	if isString(fileContentString) {
		return invalid + "\nThis is an invalid string"
	}
	if isNull(fileContentString) {
		return invalid + "\nShould be \"null\""
	}
	if isFalse(fileContentString) {
		return invalid + "\nShould be \"false\""
	}
	if isTrue(fileContentString) {
		return invalid + "\nShould be \"true\""
	}
	if isALeadedPlusNumber(fileContentString) {
		return invalid + "\nAn invalid number, there is a leading +"
	}
	if isAnArrayThatSurroundedByInvalidBrackets(fileContentString) {
		return invalid + "\nThis is an array that is surrounded by invalid \"][}{\""
	}
	if isAnArrayThatSurroundedByInvalidCommas(fileContentString) {
		return invalid + "\nThis is an array that is surrounded by invalid commas"
	}
	if isAnArray(fileContentString) {
		if isAnUnclosedArray(fileContentString) {
			return invalid + "\nThis is an unclosed array"
		}
		if isAnArrayThatClosedAsAnObject(fileContentString) {
			return invalid + "\nThis is an array that is closed as an object"
		}
		if isAnArrayThatContainsExtraAdvancingCommas(fileContentString) {
			return invalid + "\nThis is an array that contains extra advancing comma(s)"
		}
		if isAnArrayThatContainsExtraTailCommas(fileContentString) {
			return invalid + "\nThis is an array that contains extra tail comma(s)"
		}
	}
	return invalid
}

func isThereNoObjectOrArray(fileContentString string) bool {
	regex := regexp.MustCompile(`(?s)\A\s*\z`)
	return regex.MatchString(fileContentString)
}

func multipleValuesOutsidAnObjectOrArray(fileContentString string) bool {
	regex := regexp.MustCompile(`(?s)\A\s*((` + strinG + `|` + number + `|false|null|true|` + innerBrackets + `)\s*(,\s*)*){2,}\s*\z`)
	return regex.MatchString(fileContentString)
}

func isString(fileContentString string) bool {
	regex := regexp.MustCompile(`(?s)\A\s*".*"\s*\z`)
	return regex.MatchString(fileContentString)
}

func isNull(fileContentString string) bool {
	regex := regexp.MustCompile(`(?si)\A\s*null\s*\z`)
	return regex.MatchString(fileContentString)
}

func isFalse(fileContentString string) bool {
	regex := regexp.MustCompile(`(?si)\A\s*false\s*\z`)
	return regex.MatchString(fileContentString)
}

func isTrue(fileContentString string) bool {
	regex := regexp.MustCompile(`(?si)\A\s*true\s*\z`)
	return regex.MatchString(fileContentString)
}

func isALeadedZeroNumber(fileContentString string) bool {
	regex := regexp.MustCompile(`(?s)\A\s*([0-]?\d{1}\.\d+([eE][-+]?)\d+|[0-]?[1-9]\d+\.\d+([eE][-+]?)\d+|[0-]?[1-9]\d*([eE][-+]?)\d+|[0-]?\d{1}\.\d+|[0-]?[1-9]\d+\.\d+|[0-]?[1-9]\d*)\s*\z`)
	return regex.MatchString(fileContentString)
}

func isALeadedPlusNumber(fileContentString string) bool {
	regex := regexp.MustCompile(`(?s)\A\s*([+-]?\d{1}\.\d+([eE][-+]?)\d+|[+-]?[1-9]\d+\.\d+([eE][-+]?)\d+|[+-]?[1-9]\d*([eE][-+]?)\d+|[+-]?\d{1}\.\d+|[+-]?[1-9]\d+\.\d+|[+-]?[1-9]\d*)\s*\z`)
	return regex.MatchString(fileContentString)
}

func isAnArray(fileContentString string) bool {
	regex := regexp.MustCompile(`(?s)\A\s*\[.*\s*\z`)
	return regex.MatchString(fileContentString)
}

func isAnUnclosedArray(fileContentString string) bool {
	regex := regexp.MustCompile(`(?s)\A\s*\[[^]}]*\s*\z`)
	return regex.MatchString(fileContentString)
}

func isAnArrayThatClosedAsAnObject(fileContentString string) bool {
	regex := regexp.MustCompile(`(?s)\A\s*\[.*}\s*\z`)
	return regex.MatchString(fileContentString)
}

func isAnArrayThatContainsExtraTailCommas(fileContentString string) bool {
	regex := regexp.MustCompile(`(?s)\A\s*\[\s*((\s*(null|true|false|-?\d{1}\.\d+([eE][-+]?)\d+|-?[1-9]\d+\.\d+([eE][-+]?)\d+|-?[1-9]\d*([eE][-+]?)\d+|-?\d{1}\.\d+|-?[1-9]\d+\.\d+|-?[1-9]\d*|"([^"\n\t\\]*?(\\"|\\\t|\\\\|\\b|\\f|\\n|\\r|\\t|\\/)+[^"\n\t\\]*?)+"|"[^"\n\t\\]*"|\[[^][]*\]|{[^}{]*}|\[.*\[.*\].*\]|\{.*\{.*\}.*\}){1}\s*(,\s*)+)*)\]\s*\z`)
	return regex.MatchString(fileContentString)
}

func isAnArrayThatContainsExtraAdvancingCommas(fileContentString string) bool {
	regex := regexp.MustCompile(`(?s)\A\s*\[\s*(,\s*)+((\s*(null|true|false|-?\d{1}\.\d+([eE][-+]?)\d+|-?[1-9]\d+\.\d+([eE][-+]?)\d+|-?[1-9]\d*([eE][-+]?)\d+|-?\d{1}\.\d+|-?[1-9]\d+\.\d+|-?[1-9]\d*|"([^"\n\t\\]*?(\\"|\\\t|\\\\|\\b|\\f|\\n|\\r|\\t|\\/)+[^"\n\t\\]*?)+"|"[^"\n\t\\]*"|\[[^][]*\]|{[^}{]*}|\[.*\[.*\].*\]|\{.*\{.*\}.*\}){1}\s*(,\s*)*)*)\]\s*\z`)
	return regex.MatchString(fileContentString)
}

func isAnArrayThatSurroundedByInvalidBrackets(fileContentString string) bool {
	regex := regexp.MustCompile(`(?s)\A\s*(([[\]{}]\s*)+` + outerSquareBrackets + `|` + outerSquareBrackets + `(\s*[[\]{}])+|([[\]{}]\s*)+` + outerSquareBrackets + `(\s*[[\]{}])+)\s*\z`)
	return regex.MatchString(fileContentString)
}

func isAnArrayThatSurroundedByInvalidCommas(fileContentString string) bool {
	regex := regexp.MustCompile(`(?s)\A\s*((,\s*)+` + outerSquareBrackets + `|` + outerSquareBrackets + `(\s*,)+|(,\s*)+` + outerSquareBrackets + `(\s*,)+)\s*\z`)
	return regex.MatchString(fileContentString)
}
