package app

import (
	"regexp"
)

func App(readFile func(name string) ([]byte, error), args []string) (string, error) {
	fileContentInByteArray, _ := readFile(args[1])
	fileContentString := string(fileContentInByteArray)
	numbers := `-{0,1}\d+\.{0,1}\d+|-{0,1}\d+`
	inner_brackets := `\[[^][]*\]|{[^}{]*}|\[.*\[.*\].*\]|\{.*\{.*\}.*\}`
	inner_element := `\s*(null|true|false|` + numbers + `|"([^"\n]*?\\"[^"\n]*?)+"|"[^"\n]*"|` + inner_brackets + `){1}`
	last_element_in_outer_squrare_brackets := `(` + inner_element + `\s*)`
	multiple_elments_in_outer_square_brackets := `(` + inner_element + `\s*,\s*)*`
	outer_square_brackets := `\[\s*(` + multiple_elments_in_outer_square_brackets + last_element_in_outer_squrare_brackets + `{1}){0,1}\]`
	object_key := `("([^"\n]*?\\"[^"\n]*?)+"|"[^"\n]*")`
	last_element_in_outer_curly_brackets := `(\s*` + object_key + `\s*:` + inner_element + `\s*)`
	multiple_elments_in_outer_curly_brackets := `(\s*` + object_key + `\s*:` + inner_element + `\s*,\s*)*`
	outer_curly_brakets := `{\s*(` + multiple_elments_in_outer_curly_brackets + last_element_in_outer_curly_brackets + `{1}){0,1}}`
	regex_pattern := `(?s)\A\s*(` + outer_square_brackets + `|` + outer_curly_brakets + `){1}\s*\z`
	//fmt.Println(regex_pattern)
	regex := regexp.MustCompile(regex_pattern)
	if !regex.MatchString(fileContentString) {
		return "This is an invalid JSON", nil
	}
	return "This is a valid JSON", nil
}
/*
"(?s)\\A\\s*(\\[\\s*((\\s*(null|true|false|-{0,1}\\d+\\.{0,1}\\d+|-{0,1}\\d+|\"[^\"]*\"|\\[[^][]*\\]|{[^}{]*}|\\[.*\\[.*\\].*\\]|\\{.*\\{.*\\}.*\\}){1}\\s*,\\s*)*(\\s*(null|true|false|-{0,1}\\d+\\.{0,1}\\d+|-{0,1}\\d+|\"[^\"]*\"|\\[[^][]*\\]|{[^}{]*}|\\[.*\\[.*\\].*\\]|\\{.*\\{.*\\}.*\\}){1}\\s*){1}){0,1}\\]|{\\s*((\\s*\"[^\"]*\"\\s*:\\s*(null|true|false|-{0,1}\\d+\\.{0,1}\\d+|-{0,1}\\d+|\"[^\"]*\"|\\[[^][]*\\]|{[^}{]*}|\\[.*\\[.*\\].*\\]|\\{.*\\{.*\\}.*\\}){1}\\s*,\\s*)*(\\s*\"[^\"]*\"\\s*:\\s*(null|true|false|-{0,1}\\d+\\.{0,1}\\d+|-{0,1}\\d+|\"[^\"]*\"|\\[[^][]*\\]|{[^}{]*}|\\[.*\\[.*\\].*\\]|\\{.*\\{.*\\}.*\\}){1}\\s*){1}){0,1}}){1}\\s*\\z"
*/