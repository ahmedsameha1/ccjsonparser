package app

import (
	"errors"
	"regexp"
	"strings"
)

const strinG string = `"([^"\n\t\\]*?(\\"|\\\t|\\\\|\\b|\\f|\\n|\\r|\\t|\\/)+[^"\n\t\\]*?)+"|"[^"\n\t\\]*"`
const number string = `-?\d{1}\.\d+([eE][-+]?)\d+|-?[1-9]\d+\.\d+([eE][-+]?)\d+|-?[1-9]\d*([eE][-+]?)\d+|-?\d{1}\.\d+|-?[1-9]\d+\.\d+|-?[1-9]\d*`
const innerBrackets string = `\[[^][]*\]|{[^}{]*}|\[.*\[.*\].*\]|\{.*\{.*\}.*\}`
const stringValues string = `|` + strinG + `|`
const innerElement string = `\s*(null|true|false|` + number + stringValues + innerBrackets + `){1}`
const lastElementInOuterSqurareBrackets string = `(` + innerElement + `\s*)`
const multipleElmentsInOuterSquareBrackets string = `(` + innerElement + `\s*,\s*)*`
const outerSquareBrackets string = `\[\s*(` + multipleElmentsInOuterSquareBrackets + lastElementInOuterSqurareBrackets + `{1}){0,1}\]`
const objectKey string = `(` + strinG + `)`
const lastElementInOuterCurlyBrackets string = `(\s*` + objectKey + `\s*:` + innerElement + `\s*)`
const multipleElmentsInOuterCurlyBrackets string = `(\s*` + objectKey + `\s*:` + innerElement + `\s*,\s*)*`
const outerCurlyBrakets string = `{\s*(` + multipleElmentsInOuterCurlyBrackets + lastElementInOuterCurlyBrackets + `{1}){0,1}}`
const validJSONPattern string = `(?s)\A\s*(` + strinG + `|` + number + `|false|null|true|` + outerSquareBrackets + `|` +
	outerCurlyBrakets + `){1}\s*\z`

func App(readFile func(name string) ([]byte, error), args []string) (string, error) {
	fileContentInByteArray, err := readFile(args[1])
	if err != nil {
		return "", err
	}
	fileContentString := string(fileContentInByteArray)
	validJSONregex := regexp.MustCompile(validJSONPattern)
	if !validate(fileContentString, validJSONregex, 0) {
		return "", errors.New(produceAReasonForInvalidation(fileContentString))
	}
	return "This is a valid JSON", nil
}

func validate(underValidationJson string, regex *regexp.Regexp, recursionCounter int) bool {
	if recursionCounter > 18 {
		return false
	}
	if !regex.MatchString(underValidationJson) {
		return false
	}
	if containsInnerListsOrObjects(underValidationJson) {
		return handleJsonWithInnerListsOrObjects(underValidationJson, regex, recursionCounter)
	}
	return true
}

func containsInnerListsOrObjects(stringContent string) bool {
	innerBracketCheckerPattern := `(?s){(\s*("([^"\n\t\\]*?(\\"|\\\t|\\\\|\\b|\\f|\\n|\\r|\\t|\\/)+[^"\n\t\\]*?)+"|"[^"\n\t\\]*")\s*:\s*(null|true|false|-?\d{1}\.\d+([eE][-+]?)\d+|-?[1-9]\d+\.\d+([eE][-+]?)\d+|-?[1-9]\d*([eE][-+]?)\d+|-?\d{1}\.\d+|-?[1-9]\d+\.\d+|-?[1-9]\d*|"([^"\n\t\\]*?(\\"|\\\t|\\\\|\\b|\\f|\\n|\\r|\\t|\\/)+[^"\n\t\\]*?)+"|"[^"\n\t\\]*"){1}\s*,\s*)*("([^"\n\t\\]*?(\\"|\\\t|\\\\|\\b|\\f|\\n|\\r|\\t|\\/)+[^"\n\t\\]*?)+"|"[^"\n\t\\]*")\s*:\s*[{\[]|\[(\s*(null|true|false|-?\d{1}\.\d+([eE][-+]?)\d+|-?[1-9]\d+\.\d+([eE][-+]?)\d+|-?[1-9]\d*([eE][-+]?)\d+|-?\d{1}\.\d+|-?[1-9]\d+\.\d+|-?[1-9]\d*|"([^"\n\t\\]*?(\\"|\\\t|\\\\|\\b|\\f|\\n|\\r|\\t|\\/)+[^"\n\t\\]*?)+"|"[^"\n\t\\]*"){1}\s*,\s*)*[{\[]`
	innerBracketCheckerRegex := regexp.MustCompile(innerBracketCheckerPattern)
	return innerBracketCheckerRegex.MatchString(stringContent)
}

func theWholeJsonIsAnObject(stringContent string) bool {
	startWithCurlyBracketPattern := `(?s){(\s*("([^"\n\t\\]*?(\\"|\\\t|\\\\|\\b|\\f|\\n|\\r|\\t|\\/)+[^"\n\t\\]*?)+"|"[^"\n\t\\]*")\s*:\s*(null|true|false|-?\d{1}\.\d+([eE][-+]?)\d+|-?[1-9]\d+\.\d+([eE][-+]?)\d+|-?[1-9]\d*([eE][-+]?)\d+|-?\d{1}\.\d+|-?[1-9]\d+\.\d+|-?[1-9]\d*|"([^"\n\t\\]*?(\\"|\\\t|\\\\|\\b|\\f|\\n|\\r|\\t|\\/)+[^"\n\t\\]*?)+"|"[^"\n\t\\]*"){1}\s*,\s*)*("([^"\n\t\\]*?(\\"|\\\t|\\\\|\\b|\\f|\\n|\\r|\\t|\\/)+[^"\n\t\\]*?)+"|"[^"\n\t\\]*")\s*:\s*[{\[]`
	startsWithCurlyBracketRegex := regexp.MustCompile(startWithCurlyBracketPattern)
	return startsWithCurlyBracketRegex.MatchString(stringContent)
}

func handleJsonWithInnerListsOrObjects(underValidationJson string, regex *regexp.Regexp, recursionCounter int) bool {
	// Starts with {
	if theWholeJsonIsAnObject(underValidationJson) {
		innerString := removeTheOpenningBracketFromTheWholeJsonString(underValidationJson, "{")
		innerObjectsIndices := getInnerObjects(innerString)
		for _, value := range innerObjectsIndices {
			if !validate(innerString[value[0]+1:value[1]-1], regex, recursionCounter+1) {
				return false
			}
		}
		// Starts with [
	} else {
		innerString := removeTheOpenningBracketFromTheWholeJsonString(underValidationJson, "[")
		innerListsIndices := getInnerLists(innerString)
		for _, value := range innerListsIndices {
			if !validate(innerString[value[0]:value[1]-1], regex, recursionCounter+1) {
				return false
			}
		}
	}
	return true
}

func removeTheOpenningBracketFromTheWholeJsonString(fileContentString, openning string) string {
	firstSquareBracketIndex := strings.Index(fileContentString, openning)
	return fileContentString[firstSquareBracketIndex+1:]
}

func getInnerObjects(innerString string) [][]int {
	innerObjectsPattern := `:\s*(\[[^][]*\]|{[^}{]*}|\[\s*".*"\s*\]|{\s*".*"\s*}|\[.*\[.*\].*\]|\{.*\{.*\}.*\})\s*[,}]`
	innerObjectsRegex := regexp.MustCompile(innerObjectsPattern)
	innerObjects := innerObjectsRegex.FindAllIndex([]byte(innerString), -1)
	innerObjects = removeObjectsInStringValues(innerString, innerObjects)
	return innerObjects
}

func getInnerLists(innerString string) [][]int {
	innerListsPattern := `\s*(\[[^][]*\]|{[^}{]*}|\[\s*".*"\s*\]|{\s*".*"\s*}|\[.*\[.*\].*\]|\{.*\{.*\}.*\})\s*[,\]]`
	innerListsRegex := regexp.MustCompile(innerListsPattern)
	innerLists := innerListsRegex.FindAllIndex([]byte(innerString), -1)
	innerLists = removeListsInStringValues(innerString, innerLists)
	return innerLists
}

func removeObjectsInStringValues(innerString string, indices [][]int) [][]int {
	stringValuesPattern := `:\s*".*"\s*[,}]`
	stringValuesRegex := regexp.MustCompile(stringValuesPattern)
	stringValuesIndices := stringValuesRegex.FindAllIndex([]byte(innerString), -1)
	if indices[len(indices)-1][0] < (stringValuesIndices[0][1]-1) ||
		indices[0][0] > (stringValuesIndices[len(stringValuesIndices)-1][1]) {
		return indices
	}
	var revisedIndices [][]int = make([][]int, 0)
	for _, v := range indices {
		found := false
		for _, v2 := range stringValuesIndices {
			if v[0] > v2[0] && v[1] < v2[1]-1 {
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

func removeListsInStringValues(innerString string, indices [][]int) [][]int {
	stringValuesPattern := `\s*".*"\s*[,\]]`
	stringValuesRegex := regexp.MustCompile(stringValuesPattern)
	stringValuesIndices := stringValuesRegex.FindAllIndex([]byte(innerString), -1)
	if indices[len(indices)-1][0] < (stringValuesIndices[0][1]-1) ||
		indices[0][0] > (stringValuesIndices[len(stringValuesIndices)-1][1]) {
		return indices
	}
	var revisedIndices [][]int = make([][]int, 0)
	for _, v := range indices {
		found := false
		for _, v2 := range stringValuesIndices {
			if v[0] > v2[0] && v[1] < v2[1]-1 {
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

func produceAReasonForInvalidation(fileContentString string) string {
	var invalid string = "This is an invalid JSON"
	if isThereNoObjectOrArray(fileContentString) {
		return invalid + "\nMUST be an object, array, number, or string, or false or null or true"
	}
	if multipleValuesOutsidAnArray(fileContentString) {
		return invalid + "\nMultiple values outside of an array"
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
	if isALeadedZeroNumber(fileContentString) {
		return invalid + "\nAn invalid number, there is a leading zero"
	}
	if isALeadedPlusNumber(fileContentString) {
		return invalid + "\nAn invalid number, there is a leading +"
	}
	if isAnArray(fileContentString) {
		if isAnUnclosedArray(fileContentString) {
			return invalid + "\nThis is an unclosed array"
		}
		if isAnArrayThatClosedAsAnObject(fileContentString) {
			return invalid + "\nThis is an array that is closed as an object"
		}
	}
	return invalid
}

func isThereNoObjectOrArray(fileContentString string) bool {
	regex := regexp.MustCompile(`(?s)\A\s*\z`)
	return regex.MatchString(fileContentString)
}

func multipleValuesOutsidAnArray(fileContentString string) bool {
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