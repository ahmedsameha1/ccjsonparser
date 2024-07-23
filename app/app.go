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
		return invalid + "\nThere is an invalid number, there is a leading zero"
	}
	if isALeadedPlusNumber(fileContentString) {
		return invalid + "\nAn invalid number, there is a leading +"
	}
	if hasAHexadecimalNumber(fileContentString) {
		return invalid + "\nThere is an invalid number, hexadecimal numbers are not allowed"
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
	if isAnArrayThatSurroundedByInvalidBrackets(fileContentString) {
		return invalid + "\nThis is an array that is surrounded by invalid \"][}{\""
	}
	if isAnArrayThatSurroundedByInvalidCommas(fileContentString) {
		return invalid + "\nThis is an array that is surrounded by invalid commas"
	}
	if isAnObject(fileContentString) {
		if isAnObjectThatIsClosedWithCommas(fileContentString) {
			return invalid + "\nThere is an object that is closed with a comma(s)"
		}
		if isAnUnclosedObject(fileContentString) {
			return invalid + "\nThere is an unclosed object"
		}
		if isAnObjectThatClosedAsAnArray(fileContentString) {
			return invalid + "\nThere is an object that is closed as an array"
		}
		if isAnObjectThatContainsExtraAdvancingCommas(fileContentString) {
			return invalid + "\nThere is an object that contains an extra advancing comma(s)"
		}
		if isAnObjectThatContainsExtraTailCommas(fileContentString) {
			return invalid + "\nThere is an object that contains an extra tail comma(s)"
		}
		if isAnObjectThatHasAMissingColon(fileContentString) {
			return invalid + "\nThere is an object that has a missing (:)"
		}
		if isAnObjectThatHasAnInvalidColons(fileContentString) {
			return invalid + "\nThere is an object that has an invalid (:)s"
		}
	}
	if isAnArray(fileContentString) {
		if isAnArrayThatIsClosedWithCommas(fileContentString) {
			return invalid + "\nThere is an array that is closed with a comma(s)"
		}
		if isAnUnclosedArray(fileContentString) {
			return invalid + "\nThere is an unclosed array"
		}
		if isAnArrayThatClosedAsAnObject(fileContentString) {
			return invalid + "\nThere is an array that is closed as an object"
		}
		if isAnArrayThatContainsExtraAdvancingCommas(fileContentString) {
			return invalid + "\nThere is an array that contains an extra advancing comma(s)"
		}
		if isAnArrayThatContainsExtraTailCommas(fileContentString) {
			return invalid + "\nThere is an array that contains an extra tail comma(s)"
		}
		if isAnArrayThatHasAColonInsteadOfAComma(fileContentString) {
			return invalid + "\nThere is an array that has a (:) instead of a (,)"
		}
	}
	if isThereAStringThatIsNotSurroundedCorrectlyWithDoubleQuotes(fileContentString) {
		return invalid + "\nThere is a string that is not surrounded correctly by (\"\")"
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
	regex := regexp.MustCompile(`[^"]*([\[\{]\s*|\A\s*|\s+)-?0\d*[e+-.]*\d*[eE+-.]*\d*([\]\}]\s*|\s*\z|\s+)[^"]*`)
	return regex.MatchString(fileContentString)
}

func isALeadedPlusNumber(fileContentString string) bool {
	regex := regexp.MustCompile(`(?s)\A\s*([+-]?\d{1}\.\d+([eE][-+]?)\d+|[+-]?[1-9]\d+\.\d+([eE][-+]?)\d+|[+-]?[1-9]\d*([eE][-+]?)\d+|[+-]?\d{1}\.\d+|[+-]?[1-9]\d+\.\d+|[+-]?[1-9]\d*)\s*\z`)
	return regex.MatchString(fileContentString)
}

func hasAHexadecimalNumber(fileContentString string) bool {
	regex := regexp.MustCompile(`[^"]*([\[\{]\s*|\A\s*|\s+)0[xX][0-9a-fA-F]+([\]\}]\s*|\s*\z|\s+)[^"]*`)
	return regex.MatchString(fileContentString)
}

func isAnArray(fileContentString string) bool {
	regex := regexp.MustCompile(`(?s)\A\s*\[.*\s*\z`)
	return regex.MatchString(fileContentString)
}

func isAnObject(fileContentString string) bool {
	regex := regexp.MustCompile(`(?s)\A\s*\{.*\s*\z`)
	return regex.MatchString(fileContentString)
}

func isAnUnclosedArray(fileContentString string) bool {
	regex := regexp.MustCompile(`(?s)\A\s*\[[^]}]*\s*\z`)
	return regex.MatchString(fileContentString)
}

func isAnUnclosedObject(fileContentString string) bool {
	regex := regexp.MustCompile(`(?s)\A\s*\{[^]}]*\s*\z`)
	return regex.MatchString(fileContentString)
}

func isAnObjectThatIsClosedWithCommas(fileContentString string) bool {
	regex := regexp.MustCompile(`(?s)\A\s*\{[^]}]*(,\s*)+\s*\z`)
	return regex.MatchString(fileContentString)
}

func isAnArrayThatIsClosedWithCommas(fileContentString string) bool {
	regex := regexp.MustCompile(`(?s)\A\s*\[[^]}]*(,\s*)+\s*\z`)
	return regex.MatchString(fileContentString)
}

func isAnArrayThatClosedAsAnObject(fileContentString string) bool {
	regex := regexp.MustCompile(`(?s)\A\s*\[.*}\s*\z`)
	return regex.MatchString(fileContentString)
}

func isAnObjectThatClosedAsAnArray(fileContentString string) bool {
	regex := regexp.MustCompile(`(?s)\A\s*\{.*\]\s*\z`)
	return regex.MatchString(fileContentString)
}

func isAnArrayThatContainsExtraTailCommas(fileContentString string) bool {
	regex := regexp.MustCompile(`(?s)\A\s*\[\s*((\s*(null|true|false|-?\d{1}\.\d+([eE][-+]?)\d+|-?[1-9]\d+\.\d+([eE][-+]?)\d+|-?[1-9]\d*([eE][-+]?)\d+|-?\d{1}\.\d+|-?[1-9]\d+\.\d+|-?[1-9]\d*|"([^"\n\t\\]*?(\\"|\\\t|\\\\|\\b|\\f|\\n|\\r|\\t|\\/)+[^"\n\t\\]*?)+"|"[^"\n\t\\]*"|\[[^][]*\]|{[^}{]*}|\[.*\[.*\].*\]|\{.*\{.*\}.*\}){1}\s*(,\s*)+)*)\]\s*\z`)
	return regex.MatchString(fileContentString)
}

func isAnObjectThatContainsExtraTailCommas(fileContentString string) bool {
	regex := regexp.MustCompile(`(?s)\A\s*{\s*((\s*("([^"\n\t\\]*?(\\"|\\\t|\\\\|\\b|\\f|\\n|\\r|\\t|\\/|\\u)+[^"\n\t\\]*?)+"|"[^"\n\t\\]*")\s*:\s*(null|true|false|-?\d{1}\.\d+([eE][-+]?)\d+|-?[1-9]\d+\.\d+([eE][-+]?)\d+|-?[1-9]\d*([eE][-+]?)\d+|-?\d{1}\.\d+|-?[1-9]\d+\.\d+|-?[1-9]\d*|-?0([eE][-+]?\d+){0,1}|"([^"\n\t\\]*?(\\"|\\\t|\\\\|\\b|\\f|\\n|\\r|\\t|\\/|\\u)+[^"\n\t\\]*?)+"|"[^"\n\t\\]*"|\[[^][]*\]|{[^}{]*}|\[.*\[.*\].*\]|\{.*\{.*\}.*\}){1}\s*,\s*)*(\s*("([^"\n\t\\]*?(\\"|\\\t|\\\\|\\b|\\f|\\n|\\r|\\t|\\/|\\u)+[^"\n\t\\]*?)+"|"[^"\n\t\\]*")\s*:\s*(null|true|false|-?\d{1}\.\d+([eE][-+]?)\d+|-?[1-9]\d+\.\d+([eE][-+]?)\d+|-?[1-9]\d*([eE][-+]?)\d+|-?\d{1}\.\d+|-?[1-9]\d+\.\d+|-?[1-9]\d*|-?0([eE][-+]?\d+){0,1}|"([^"\n\t\\]*?(\\"|\\\t|\\\\|\\b|\\f|\\n|\\r|\\t|\\/|\\u)+[^"\n\t\\]*?)+"|"[^"\n\t\\]*"|\[[^][]*\]|{[^}{]*}|\[.*\[.*\].*\]|\{.*\{.*\}.*\}){1}\s*(,\s*)+){1}){0,1}}\s*\z`)
	return regex.MatchString(fileContentString)
}

func isAnObjectThatContainsExtraAdvancingCommas(fileContentString string) bool {
	regex := regexp.MustCompile(`(?s)\A\s*{\s*(,\s*)+((\s*("([^"\n\t\\]*?(\\"|\\\t|\\\\|\\b|\\f|\\n|\\r|\\t|\\/|\\u)+[^"\n\t\\]*?)+"|"[^"\n\t\\]*")\s*:\s*(null|true|false|-?\d{1}\.\d+([eE][-+]?)\d+|-?[1-9]\d+\.\d+([eE][-+]?)\d+|-?[1-9]\d*([eE][-+]?)\d+|-?\d{1}\.\d+|-?[1-9]\d+\.\d+|-?[1-9]\d*|-?0([eE][-+]?\d+){0,1}|"([^"\n\t\\]*?(\\"|\\\t|\\\\|\\b|\\f|\\n|\\r|\\t|\\/|\\u)+[^"\n\t\\]*?)+"|"[^"\n\t\\]*"|\[[^][]*\]|{[^}{]*}|\[.*\[.*\].*\]|\{.*\{.*\}.*\}){1}\s*,\s*)*(\s*("([^"\n\t\\]*?(\\"|\\\t|\\\\|\\b|\\f|\\n|\\r|\\t|\\/|\\u)+[^"\n\t\\]*?)+"|"[^"\n\t\\]*")\s*:\s*(null|true|false|-?\d{1}\.\d+([eE][-+]?)\d+|-?[1-9]\d+\.\d+([eE][-+]?)\d+|-?[1-9]\d*([eE][-+]?)\d+|-?\d{1}\.\d+|-?[1-9]\d+\.\d+|-?[1-9]\d*|-?0([eE][-+]?\d+){0,1}|"([^"\n\t\\]*?(\\"|\\\t|\\\\|\\b|\\f|\\n|\\r|\\t|\\/|\\u)+[^"\n\t\\]*?)+"|"[^"\n\t\\]*"|\[[^][]*\]|{[^}{]*}|\[.*\[.*\].*\]|\{.*\{.*\}.*\}){1}\s*(,\s*)*){1}){0,1}}\s*\z`)
	return regex.MatchString(fileContentString)
}

func isAnObjectThatHasAMissingColon(fileContentString string) bool {
	regex := regexp.MustCompile(`(?s)\A\s*{\s*((\s*("([^"\n\t\\]*?(\\"|\\\t|\\\\|\\b|\\f|\\n|\\r|\\t|\\/|\\u)+[^"\n\t\\]*?)+"|"[^"\n\t\\]*")\s*:?\s*(null|true|false|-?\d{1}\.\d+([eE][-+]?)\d+|-?[1-9]\d+\.\d+([eE][-+]?)\d+|-?[1-9]\d*([eE][-+]?)\d+|-?\d{1}\.\d+|-?[1-9]\d+\.\d+|-?[1-9]\d*|-?0([eE][-+]?\d+){0,1}|"([^"\n\t\\]*?(\\"|\\\t|\\\\|\\b|\\f|\\n|\\r|\\t|\\/|\\u)+[^"\n\t\\]*?)+"|"[^"\n\t\\]*"|\[[^][]*\]|{[^}{]*}|\[.*\[.*\].*\]|\{.*\{.*\}.*\}){1}\s*,\s*)*(\s*("([^"\n\t\\]*?(\\"|\\\t|\\\\|\\b|\\f|\\n|\\r|\\t|\\/|\\u)+[^"\n\t\\]*?)+"|"[^"\n\t\\]*")\s*:?\s*(null|true|false|-?\d{1}\.\d+([eE][-+]?)\d+|-?[1-9]\d+\.\d+([eE][-+]?)\d+|-?[1-9]\d*([eE][-+]?)\d+|-?\d{1}\.\d+|-?[1-9]\d+\.\d+|-?[1-9]\d*|-?0([eE][-+]?\d+){0,1}|"([^"\n\t\\]*?(\\"|\\\t|\\\\|\\b|\\f|\\n|\\r|\\t|\\/|\\u)+[^"\n\t\\]*?)+"|"[^"\n\t\\]*"|\[[^][]*\]|{[^}{]*}|\[.*\[.*\].*\]|\{.*\{.*\}.*\}){1}\s*){1}){0,1}}\s*\z`)
	return regex.MatchString(fileContentString)
}

func isAnObjectThatHasAnInvalidColons(fileContentString string) bool {
	regex := regexp.MustCompile(`(?s)\A\s*{\s*(:\s*)*((\s*("([^"\n\t\\]*?(\\"|\\\t|\\\\|\\b|\\f|\\n|\\r|\\t|\\/|\\u)+[^"\n\t\\]*?)+"|"[^"\n\t\\]*")\s*(:\s*)+\s*(null|true|false|-?\d{1}\.\d+([eE][-+]?)\d+|-?[1-9]\d+\.\d+([eE][-+]?)\d+|-?[1-9]\d*([eE][-+]?)\d+|-?\d{1}\.\d+|-?[1-9]\d+\.\d+|-?[1-9]\d*|-?0([eE][-+]?\d+){0,1}|"([^"\n\t\\]*?(\\"|\\\t|\\\\|\\b|\\f|\\n|\\r|\\t|\\/|\\u)+[^"\n\t\\]*?)+"|"[^"\n\t\\]*"|\[[^][]*\]|{[^}{]*}|\[.*\[.*\].*\]|\{.*\{.*\}.*\}){1}\s*(:\s*)*,(\s*:)*\s*)*(\s*("([^"\n\t\\]*?(\\"|\\\t|\\\\|\\b|\\f|\\n|\\r|\\t|\\/|\\u)+[^"\n\t\\]*?)+"|"[^"\n\t\\]*")\s*(:\s*)+\s*(null|true|false|-?\d{1}\.\d+([eE][-+]?)\d+|-?[1-9]\d+\.\d+([eE][-+]?)\d+|-?[1-9]\d*([eE][-+]?)\d+|-?\d{1}\.\d+|-?[1-9]\d+\.\d+|-?[1-9]\d*|-?0([eE][-+]?\d+){0,1}|"([^"\n\t\\]*?(\\"|\\\t|\\\\|\\b|\\f|\\n|\\r|\\t|\\/|\\u)+[^"\n\t\\]*?)+"|"[^"\n\t\\]*"|\[[^][]*\]|{[^}{]*}|\[.*\[.*\].*\]|\{.*\{.*\}.*\}){1}\s*){1}){0,1}(:\s*)*}\s*\z`)
	return regex.MatchString(fileContentString)
}

func isAnArrayThatContainsExtraAdvancingCommas(fileContentString string) bool {
	regex := regexp.MustCompile(`(?s)\A\s*\[\s*(,\s*)+((\s*(null|true|false|-?\d{1}\.\d+([eE][-+]?)\d+|-?[1-9]\d+\.\d+([eE][-+]?)\d+|-?[1-9]\d*([eE][-+]?)\d+|-?\d{1}\.\d+|-?[1-9]\d+\.\d+|-?[1-9]\d*|"([^"\n\t\\]*?(\\"|\\\t|\\\\|\\b|\\f|\\n|\\r|\\t|\\/)+[^"\n\t\\]*?)+"|"[^"\n\t\\]*"|\[[^][]*\]|{[^}{]*}|\[.*\[.*\].*\]|\{.*\{.*\}.*\}){1}\s*(,\s*)*)*)\]\s*\z`)
	return regex.MatchString(fileContentString)
}

func isAnArrayThatHasAColonInsteadOfAComma(fileContentString string) bool {
	regex := regexp.MustCompile(`(?s)\A\s*\[\s*((\s*(null|true|false|-?\d{1}\.\d+([eE][-+]?)\d+|-?[1-9]\d+\.\d+([eE][-+]?)\d+|-?[1-9]\d*([eE][-+]?)\d+|-?\d{1}\.\d+|-?[1-9]\d+\.\d+|-?[1-9]\d*|-?0([eE][-+]?\d+){0,1}|"([^"\n\t\\]*?(\\"|\\\t|\\\\|\\b|\\f|\\n|\\r|\\t|\\/|\\u)+[^"\n\t\\]*?)+"|"[^"\n\t\\]*"|\[[^][]*\]|{[^}{]*}|\[.*\[.*\].*\]|\{.*\{.*\}.*\}){1}\s*[:,]\s*)*(\s*(null|true|false|-?\d{1}\.\d+([eE][-+]?)\d+|-?[1-9]\d+\.\d+([eE][-+]?)\d+|-?[1-9]\d*([eE][-+]?)\d+|-?\d{1}\.\d+|-?[1-9]\d+\.\d+|-?[1-9]\d*|-?0([eE][-+]?\d+){0,1}|"([^"\n\t\\]*?(\\"|\\\t|\\\\|\\b|\\f|\\n|\\r|\\t|\\/|\\u)+[^"\n\t\\]*?)+"|"[^"\n\t\\]*"|\[[^][]*\]|{[^}{]*}|\[.*\[.*\].*\]|\{.*\{.*\}.*\}){1}\s*){1}){0,1}\]\s*\z`)
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

func isThereAStringThatIsNotSurroundedCorrectlyWithDoubleQuotes(fileContentString string) bool {
	doubleQuotesRegex := regexp.MustCompile(`"`)
	doubleQuotesIndices := doubleQuotesRegex.FindAllStringIndex(fileContentString, -1)
	escapedDoubleQuotesRegex := regexp.MustCompile(`\\"`)
	escapedDoubleQuotesIndices := escapedDoubleQuotesRegex.FindAllStringIndex(fileContentString, -1)
	var revisedIndices [][]int = make([][]int, 0)
	for _, v1 := range doubleQuotesIndices {
		found := false
		for _, v2 := range escapedDoubleQuotesIndices {
			if v2[0]+1 == v1[0] {
				found = true
				break
			}
		}
		if !found {
			revisedIndices = append(revisedIndices, v1)
		}
	}
	if len(revisedIndices)%2 == 0 {
		regex1 := regexp.MustCompile(`[,:\[\{]\s*(\d+)?[^",0-9]+\s*(\d+)?[,:\]\}]`)
		unquotedStrings := regex1.FindAllString(fileContentString, -1)
		regex2 := regexp.MustCompile(`(?i)\b(null|true|false|-?\d{1}\.\d+([e][-+]?)\d*|-?\d+\.\d+([e][-+]?)\d*|-?\d+([e][-+]?)\d*|-?\d{1}\.\d*|-?\d+\.\d*|-?\d+|-?0([e][-+]?\d*){0,1}|0x[0-9a-f]+)\b|\s*:\s*:\s*`)
		revised := make([]string, 0)
		for _, v := range unquotedStrings {
			if !regex2.MatchString(v) {
				revised = append(revised, v)
			}
		}
		return len(revised) > 0
	}
	return len(revisedIndices)%2 == 1
}
