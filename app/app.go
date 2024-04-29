package app

import "regexp"

func App(readFile func(name string) ([]byte, error), args []string) (string, error) {
	fileContentInByteArray, _ := readFile(args[1])
	fileContentString := string(fileContentInByteArray)
	regex := regexp.MustCompile(`(?s)\A\s*(\[\s*((\s*(null|true|false|-{0,1}\d+\.{0,1}\d+|-{0,1}\d+|"[^\s"]+"|\[[^][]*\]|{[^}{]*}|\[.*\[.*\].*\]|\{.*\{.*\}.*\}){1}\s*,\s*)*(\s*(null|true|false|-{0,1}\d+\.{0,1}\d+|-{0,1}\d+|"[^\s"]+"|\[[^][]*\]|{[^}{]*}|\[.*\[.*\].*\]|\{.*\{.*\}.*\}){1}\s*){1}){0,1}\]|{\s*((\s*"[^\s"]+"\s*:\s*(null|true|false|-{0,1}\d+\.{0,1}\d+|-{0,1}\d+|"[^\s"]+"|\[[^][]*\]|{[^}{]*}|\[.*\[.*\].*\]|\{.*\{.*\}.*\}){1}\s*,\s*)*(\s*"[^\s"]+"\s*:\s*(null|true|false|-{0,1}\d+\.{0,1}\d+|-{0,1}\d+|"[^\s"]+"|\[[^][]*\]|{[^}{]*}|\[.*\[.*\].*\]|\{.*\{.*\}.*\}){1}\s*){1}){0,1}}){1}\s*\z`)
	if !regex.MatchString(fileContentString) {
		return "This is an invalid JSON", nil
	}
	return "This is a valid JSON", nil
}
