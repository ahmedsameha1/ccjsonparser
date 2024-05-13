package app

import (
	"log"
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
	if containsInnerBracket(fileContentString) {
		// Starts with {
		if startsWithCurlyBracket(fileContentString) {
			firstCurlyBracketIndex := strings.Index(fileContentString, "{")
			lastCurlyBracketIndex := strings.LastIndex(fileContentString, "}")
			innerString := fileContentString[firstCurlyBracketIndex+1 : lastCurlyBracketIndex]
			if isLeftMostBracketCurly(innerString) {
				firstCurlyBracketIndex = strings.Index(innerString, "{")
				firstClosingCurlyBracketAfterTheOpenningOnePattern := `(?s)}\s*[,}\]]?\s*`
				firstClosingCurlyBracketAfterTheOpenningOneRegex :=
					regexp.MustCompile(firstClosingCurlyBracketAfterTheOpenningOnePattern)
				startIndexAndEndIndexOfTheClosingCurlyBracket :=
					firstClosingCurlyBracketAfterTheOpenningOneRegex.FindStringIndex(innerString)
				if startIndexAndEndIndexOfTheClosingCurlyBracket != nil {
					startIndexOfTheClosingCurlyBracket := startIndexAndEndIndexOfTheClosingCurlyBracket[0]
					supposed_inner_json := innerString[firstCurlyBracketIndex : startIndexOfTheClosingCurlyBracket+1]
					log.Printf(`"""%s"""\n`, supposed_inner_json)
					if !regex.MatchString(supposed_inner_json) {
						return "This is an invalid JSON", nil
					}
				} else {
					return "This is an invalid JSON", nil
				}
				// Leftmost is [
			} else {
				firstSquareBracketIndex := strings.Index(fileContentString, "[")
				lastSquareBracketIndex := strings.LastIndex(fileContentString, "]")
				innerString := fileContentString[firstSquareBracketIndex+1 : lastSquareBracketIndex]
				firstSquareBracketIndex = strings.Index(innerString, "[")
				firstClosingSquareBracketAfterTheOpenningOnePattern := `(?s)]\s*[,}\]]?\s*`
				firstClosingSquareBracketAfterTheOpenningOneRegex :=
					regexp.MustCompile(firstClosingSquareBracketAfterTheOpenningOnePattern)
				startIndexAndEndIndexOfTheClosingSquareBracket :=
					firstClosingSquareBracketAfterTheOpenningOneRegex.FindStringIndex(innerString)
				if startIndexAndEndIndexOfTheClosingSquareBracket != nil {
					startIndexOfTheClosingSquareBracket := startIndexAndEndIndexOfTheClosingSquareBracket[0]
					supposed_inner_json := innerString[firstSquareBracketIndex : startIndexOfTheClosingSquareBracket+1]
					log.Printf(`"""%s"""\n`, supposed_inner_json)
					if !regex.MatchString(supposed_inner_json) {
						return "This is an invalid JSON", nil
					}
				} else {
					return "This is an invalid JSON", nil
				}
			}
			// Starts with [
		} else {
			firstSquareBracketIndex := strings.Index(fileContentString, "[")
			lastSquareBracketIndex := strings.LastIndex(fileContentString, "]")
			innerString := fileContentString[firstSquareBracketIndex+1 : lastSquareBracketIndex]
			if isLeftMostBracketCurly(innerString) {
				firstCurlyBracketIndex := strings.Index(innerString, "{")
				firstClosingCurlyBracketAfterTheOpenningOnePattern := `(?s)}\s*[,}\]]?\s*`
				firstClosingCurlyBracketAfterTheOpenningOneRegex :=
					regexp.MustCompile(firstClosingCurlyBracketAfterTheOpenningOnePattern)
				startIndexAndEndIndexOfTheClosingCurlyBracket :=
					firstClosingCurlyBracketAfterTheOpenningOneRegex.FindStringIndex(innerString)
				if startIndexAndEndIndexOfTheClosingCurlyBracket != nil {
					startIndexOfTheClosingCurlyBracket := startIndexAndEndIndexOfTheClosingCurlyBracket[0]
					supposed_inner_json := innerString[firstCurlyBracketIndex : startIndexOfTheClosingCurlyBracket+1]
					log.Printf(`"""%s"""\n`, supposed_inner_json)
					if !regex.MatchString(supposed_inner_json) {
						return "This is an invalid JSON", nil
					}
				} else {
					return "This is an invalid JSON", nil
				}
				// Leftmost is [
			} else {
				firstSquareBracketIndex := strings.Index(fileContentString, "[")
				lastSquareBracketIndex := strings.LastIndex(fileContentString, "]")
				innerString := fileContentString[firstSquareBracketIndex+1 : lastSquareBracketIndex]
				firstSquareBracketIndex = strings.Index(innerString, "[")
				firstClosingSquareBracketAfterTheOpenningOnePattern := `(?s)]\s*[,}\]]?\s*`
				firstClosingSquareBracketAfterTheOpenningOneRegex :=
					regexp.MustCompile(firstClosingSquareBracketAfterTheOpenningOnePattern)
				startIndexAndEndIndexOfTheClosingSquareBracket :=
					firstClosingSquareBracketAfterTheOpenningOneRegex.FindStringIndex(innerString)
				if startIndexAndEndIndexOfTheClosingSquareBracket != nil {
					startIndexOfTheClosingSquareBracket := startIndexAndEndIndexOfTheClosingSquareBracket[0]
					supposed_inner_json := innerString[firstSquareBracketIndex : startIndexOfTheClosingSquareBracket+1]
					log.Printf(`"""%s"""\n`, supposed_inner_json)
					if !regex.MatchString(supposed_inner_json) {
						return "This is an invalid JSON", nil
					}
				} else {
					return "This is an invalid JSON", nil
				}
			}
			return "This is a valid JSON", nil
		}
	}
	return "This is a valid JSON", nil
}

func containsInnerBracket(stringContent string) bool {
	//innerBracketCheckerPattern := `(?s){\s*(.*,)*\s*.*[{\[]|\[\s*(.*,)*\s*[{\[]`
	innerBracketCheckerPattern := `(?s){(\s*("([^"\n\t\\]*?(\\"|\\\t|\\\\|\\b|\\f|\\n|\\r|\\t|\\/)+[^"\n\t\\]*?)+"|"[^"\n\t\\]*")\s*:\s*(null|true|false|-?\d{1}\.\d+([eE][-+]?)\d+|-?[1-9]\d+\.\d+([eE][-+]?)\d+|-?[1-9]\d*([eE][-+]?)\d+|-?\d{1}\.\d+|-?[1-9]\d+\.\d+|-?[1-9]\d*|"([^"\n\t\\]*?(\\"|\\\t|\\\\|\\b|\\f|\\n|\\r|\\t|\\/)+[^"\n\t\\]*?)+"|"[^"\n\t\\]*"){1}\s*,\s*)*("([^"\n\t\\]*?(\\"|\\\t|\\\\|\\b|\\f|\\n|\\r|\\t|\\/)+[^"\n\t\\]*?)+"|"[^"\n\t\\]*")\s*:\s*[{\[]|\[(\s*(null|true|false|-?\d{1}\.\d+([eE][-+]?)\d+|-?[1-9]\d+\.\d+([eE][-+]?)\d+|-?[1-9]\d*([eE][-+]?)\d+|-?\d{1}\.\d+|-?[1-9]\d+\.\d+|-?[1-9]\d*|"([^"\n\t\\]*?(\\"|\\\t|\\\\|\\b|\\f|\\n|\\r|\\t|\\/)+[^"\n\t\\]*?)+"|"[^"\n\t\\]*"){1}\s*,\s*)*[{\[]`
	innerBracketCheckerRegex := regexp.MustCompile(innerBracketCheckerPattern)
	return innerBracketCheckerRegex.MatchString(stringContent)
}

func startsWithCurlyBracket(stringContent string) bool {
	startWithCurlyBracketPattern := `(?s){(\s*("([^"\n\t\\]*?(\\"|\\\t|\\\\|\\b|\\f|\\n|\\r|\\t|\\/)+[^"\n\t\\]*?)+"|"[^"\n\t\\]*")\s*:\s*(null|true|false|-?\d{1}\.\d+([eE][-+]?)\d+|-?[1-9]\d+\.\d+([eE][-+]?)\d+|-?[1-9]\d*([eE][-+]?)\d+|-?\d{1}\.\d+|-?[1-9]\d+\.\d+|-?[1-9]\d*|"([^"\n\t\\]*?(\\"|\\\t|\\\\|\\b|\\f|\\n|\\r|\\t|\\/)+[^"\n\t\\]*?)+"|"[^"\n\t\\]*"){1}\s*,\s*)*("([^"\n\t\\]*?(\\"|\\\t|\\\\|\\b|\\f|\\n|\\r|\\t|\\/)+[^"\n\t\\]*?)+"|"[^"\n\t\\]*")\s*:\s*[{\[]`
	startsWithCurlyBracketRegex := regexp.MustCompile(startWithCurlyBracketPattern)
	return startsWithCurlyBracketRegex.MatchString(stringContent)
}

func isLeftMostBracketCurly(innerString string) bool {
	innerBracketCheckerPattern := `(?s)(\s*("([^"\n\t\\]*?(\\"|\\\t|\\\\|\\b|\\f|\\n|\\r|\\t|\\/)+[^"\n\t\\]*?)+"|"[^"\n\t\\]*")\s*:\s*(null|true|false|-?\d{1}\.\d+([eE][-+]?)\d+|-?[1-9]\d+\.\d+([eE][-+]?)\d+|-?[1-9]\d*([eE][-+]?)\d+|-?\d{1}\.\d+|-?[1-9]\d+\.\d+|-?[1-9]\d*|"([^"\n\t\\]*?(\\"|\\\t|\\\\|\\b|\\f|\\n|\\r|\\t|\\/)+[^"\n\t\\]*?)+"|"[^"\n\t\\]*"){1}\s*,\s*)*("([^"\n\t\\]*?(\\"|\\\t|\\\\|\\b|\\f|\\n|\\r|\\t|\\/)+[^"\n\t\\]*?)+"|"[^"\n\t\\]*")\s*:\s*[{\[]`
	innerBracketCheckerRegex := regexp.MustCompile(innerBracketCheckerPattern)
	stringEndingWithABracket := innerBracketCheckerRegex.FindString(innerString)
	if len(stringEndingWithABracket) > 0 {
		return string(stringEndingWithABracket[len(stringEndingWithABracket)-1]) == "{"
	}
	return false
}
