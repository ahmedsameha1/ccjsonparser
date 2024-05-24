package app

import (
	"regexp"
	"strings"
)

func App(readFile func(name string) ([]byte, error), args []string) (string, error) {
	fileContentInByteArray, _ := readFile(args[1])
	fileContentString := string(fileContentInByteArray)
	string_keys_values := `"([^"\n\t\\]*?(\\"|\\\t|\\\\|\\b|\\f|\\n|\\r|\\t|\\/)+[^"\n\t\\]*?)+"|"[^"\n\t\\]*"`
	string_values := `|` + string_keys_values + `|`
	numbers := `-?\d{1}\.\d+([eE][-+]?)\d+|-?[1-9]\d+\.\d+([eE][-+]?)\d+|-?[1-9]\d*([eE][-+]?)\d+|-?\d{1}\.\d+|-?[1-9]\d+\.\d+|-?[1-9]\d*`
	inner_brackets := `\[[^][]*\]|{[^}{]*}|\[.*\[.*\].*\]|\{.*\{.*\}.*\}`
	inner_element := `\s*(null|true|false|` + numbers + string_values + inner_brackets + `){1}`
	last_element_in_outer_squrare_brackets := `(` + inner_element + `\s*)`
	multiple_elments_in_outer_square_brackets := `(` + inner_element + `\s*,\s*)*`
	outer_square_brackets := `\[\s*(` + multiple_elments_in_outer_square_brackets + last_element_in_outer_squrare_brackets + `{1}){0,1}\]`
	object_key := `(` + string_keys_values + `)`
	last_element_in_outer_curly_brackets := `(\s*` + object_key + `\s*:` + inner_element + `\s*)`
	multiple_elments_in_outer_curly_brackets := `(\s*` + object_key + `\s*:` + inner_element + `\s*,\s*)*`
	outer_curly_brakets := `{\s*(` + multiple_elments_in_outer_curly_brackets + last_element_in_outer_curly_brackets + `{1}){0,1}}`
	regex_pattern := `(?s)\A\s*(` + outer_square_brackets + `|` + outer_curly_brakets + `){1}\s*\z`
	regex := regexp.MustCompile(regex_pattern)
	if !regex.MatchString(fileContentString) {
		return "This is an invalid JSON", nil
	}
	if containsInnerListsOrObjects(fileContentString) {
		return handleJsonWithInnerListsOrObjects(fileContentString, regex)
	}
	return "This is a valid JSON", nil
}

func containsInnerListsOrObjects(stringContent string) bool {
	innerBracketCheckerPattern := `(?s){(\s*("([^"\n\t\\]*?(\\"|\\\t|\\\\|\\b|\\f|\\n|\\r|\\t|\\/)+[^"\n\t\\]*?)+"|"[^"\n\t\\]*")\s*:\s*(null|true|false|-?\d{1}\.\d+([eE][-+]?)\d+|-?[1-9]\d+\.\d+([eE][-+]?)\d+|-?[1-9]\d*([eE][-+]?)\d+|-?\d{1}\.\d+|-?[1-9]\d+\.\d+|-?[1-9]\d*|"([^"\n\t\\]*?(\\"|\\\t|\\\\|\\b|\\f|\\n|\\r|\\t|\\/)+[^"\n\t\\]*?)+"|"[^"\n\t\\]*"){1}\s*,\s*)*("([^"\n\t\\]*?(\\"|\\\t|\\\\|\\b|\\f|\\n|\\r|\\t|\\/)+[^"\n\t\\]*?)+"|"[^"\n\t\\]*")\s*:\s*[{\[]|\[(\s*(null|true|false|-?\d{1}\.\d+([eE][-+]?)\d+|-?[1-9]\d+\.\d+([eE][-+]?)\d+|-?[1-9]\d*([eE][-+]?)\d+|-?\d{1}\.\d+|-?[1-9]\d+\.\d+|-?[1-9]\d*|"([^"\n\t\\]*?(\\"|\\\t|\\\\|\\b|\\f|\\n|\\r|\\t|\\/)+[^"\n\t\\]*?)+"|"[^"\n\t\\]*"){1}\s*,\s*)*[{\[]`
	innerBracketCheckerRegex := regexp.MustCompile(innerBracketCheckerPattern)
	return innerBracketCheckerRegex.MatchString(stringContent)
}

func theWholeJSONstartsWithCurlyBracket(stringContent string) bool {
	startWithCurlyBracketPattern := `(?s){(\s*("([^"\n\t\\]*?(\\"|\\\t|\\\\|\\b|\\f|\\n|\\r|\\t|\\/)+[^"\n\t\\]*?)+"|"[^"\n\t\\]*")\s*:\s*(null|true|false|-?\d{1}\.\d+([eE][-+]?)\d+|-?[1-9]\d+\.\d+([eE][-+]?)\d+|-?[1-9]\d*([eE][-+]?)\d+|-?\d{1}\.\d+|-?[1-9]\d+\.\d+|-?[1-9]\d*|"([^"\n\t\\]*?(\\"|\\\t|\\\\|\\b|\\f|\\n|\\r|\\t|\\/)+[^"\n\t\\]*?)+"|"[^"\n\t\\]*"){1}\s*,\s*)*("([^"\n\t\\]*?(\\"|\\\t|\\\\|\\b|\\f|\\n|\\r|\\t|\\/)+[^"\n\t\\]*?)+"|"[^"\n\t\\]*")\s*:\s*[{\[]`
	startsWithCurlyBracketRegex := regexp.MustCompile(startWithCurlyBracketPattern)
	return startsWithCurlyBracketRegex.MatchString(stringContent)
}

func handleJsonWithInnerListsOrObjects(fileContentString string, regex *regexp.Regexp) (string, error) {
	// Starts with {
	if theWholeJSONstartsWithCurlyBracket(fileContentString) {
		innerString := removeOpenningBracketFromTheWholeJsonString(fileContentString, "{")
		innerObjectsIndices := getInnerObjects(innerString)
		innerObjectsIndices = removeObjectsInStringValues(innerString, innerObjectsIndices)
		for _, value := range innerObjectsIndices {
			if !validateInnerJson(innerString[value[0]+1:value[1]-1], regex) {
				return "This is an invalid JSON", nil
			}
		}
		// Starts with [
	} else {
		innerString := removeOpenningBracketFromTheWholeJsonString(fileContentString, "[")
		innerListsIndices := getInnerLists(innerString)
		innerListsIndices = removeListsInStringValues(innerString, innerListsIndices)
		for _, value := range innerListsIndices {
			if !validateInnerJson(innerString[value[0]+1:value[1]-1], regex) {
				return "This is an invalid JSON", nil
			}
		}
	}
	return "This is a valid JSON", nil
}

func removeOpenningBracketFromTheWholeJsonString(fileContentString, openning string) string {
	firstSquareBracketIndex := strings.Index(fileContentString, openning)
	return fileContentString[firstSquareBracketIndex+1:]
}

func getInnerObjects(innerString string) [][]int {
	innerObjectsPattern := `:\s*(\[[^][]*\]|{[^}{]*}|\[\s*".*"\s*\]|{\s*".*"\s*}|\[.*\[.*\].*\]|\{.*\{.*\}.*\})\s*[,}]`
	innerObjectsRegex := regexp.MustCompile(innerObjectsPattern)
	return innerObjectsRegex.FindAllIndex([]byte(innerString), -1)
}

func getInnerLists(innerString string) [][]int {
	innerListsPattern := `\s*(\[[^][]*\]|{[^}{]*}|\[\s*".*"\s*\]|{\s*".*"\s*}|\[.*\[.*\].*\]|\{.*\{.*\}.*\})\s*[,\]]`
	innerListsRegex := regexp.MustCompile(innerListsPattern)
	return innerListsRegex.FindAllIndex([]byte(innerString), -1)
}

func validateInnerJson(objectString string, regex *regexp.Regexp) bool {
	return regex.MatchString(objectString)
}

// Needs to be rewritten to remove the unneeded work
func removeObjectsInStringValues(innerString string, indices [][]int) [][]int {
	stringValuesPattern := `:\s*".*"\s*[,}]`
	stringValuesRegex := regexp.MustCompile(stringValuesPattern)
	stringValuesIndices := stringValuesRegex.FindAllIndex([]byte(innerString), -1)
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

// Needs to be rewritten to remove the unneeded work
func removeListsInStringValues(innerString string, indices [][]int) [][]int {
	stringValuesPattern := `\s*".*"\s*[,\]]`
	stringValuesRegex := regexp.MustCompile(stringValuesPattern)
	stringValuesIndices := stringValuesRegex.FindAllIndex([]byte(innerString), -1)
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
