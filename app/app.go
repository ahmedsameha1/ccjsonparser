package app

import (
	"regexp"
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
	return "This is a valid JSON", nil
}
