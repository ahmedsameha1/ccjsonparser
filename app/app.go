package app

import (
	"regexp"
	"strings"
)

func App(readFile func(name string) ([]byte, error), args []string) (string, error) {
	fileContentInByteArray, err := readFile(args[1])
	if err != nil {
		return "", err
	}
	fileContentString := string(fileContentInByteArray)
	strinG := `"([^"\n\t\\]*?(\\"|\\\t|\\\\|\\b|\\f|\\n|\\r|\\t|\\/)+[^"\n\t\\]*?)+"|"[^"\n\t\\]*"`
	string_values := `|` + strinG + `|`
	number := `-?\d{1}\.\d+([eE][-+]?)\d+|-?[1-9]\d+\.\d+([eE][-+]?)\d+|-?[1-9]\d*([eE][-+]?)\d+|-?\d{1}\.\d+|-?[1-9]\d+\.\d+|-?[1-9]\d*`
	inner_brackets := `\[[^][]*\]|{[^}{]*}|\[.*\[.*\].*\]|\{.*\{.*\}.*\}`
	inner_element := `\s*(null|true|false|` + number + string_values + inner_brackets + `){1}`
	last_element_in_outer_squrare_brackets := `(` + inner_element + `\s*)`
	multiple_elments_in_outer_square_brackets := `(` + inner_element + `\s*,\s*)*`
	outer_square_brackets := `\[\s*(` + multiple_elments_in_outer_square_brackets + last_element_in_outer_squrare_brackets + `{1}){0,1}\]`
	object_key := `(` + strinG + `)`
	last_element_in_outer_curly_brackets := `(\s*` + object_key + `\s*:` + inner_element + `\s*)`
	multiple_elments_in_outer_curly_brackets := `(\s*` + object_key + `\s*:` + inner_element + `\s*,\s*)*`
	outer_curly_brakets := `{\s*(` + multiple_elments_in_outer_curly_brackets + last_element_in_outer_curly_brackets + `{1}){0,1}}`
	regex_pattern := `(?s)\A\s*(` + strinG + `|` + number + `|` + outer_square_brackets + `|` +
		outer_curly_brakets + `){1}\s*\z`
	regex := regexp.MustCompile(regex_pattern)
	if !validate(fileContentString, regex, 0) {
		return produceAReasonForInvalidation(fileContentString), nil
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
	if IsThereNoObjectOrArray(fileContentString) {
		return "MUST be an object, array, number, or string, or false or null or true"
	}
	return "This is an invalid JSON"
}

func IsThereNoObjectOrArray(fileContentString string) bool {
	regex := regexp.MustCompile(`(?s)\A\s*\z`)
	return regex.MatchString(fileContentString)
}
