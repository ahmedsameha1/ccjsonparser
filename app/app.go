package app

import "regexp"

func App(readFile func(name string) ([]byte, error), args []string) (string, error) {
	fileContentInByteArray, _ := readFile(args[1])
	fileContentString := string(fileContentInByteArray)
	regex, _ := regexp.Compile(`^\s*{((\s*"\S+"\s*:\s*"\S+"\s*,\s*)*(\s*"\S+"\s*:\s*"\S+"\s*){1}){0,1}}\s*$`)
	if !regex.MatchString(fileContentString) {
		return "This is an invalid JSON", nil
	}
	return "This is a valid JSON", nil
}
