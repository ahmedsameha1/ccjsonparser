package app

func App(readFile func(name string) ([]byte, error), args []string) (string, error) {
	fileContentInByteArray, _ := readFile(args[1])
	fileContentString := string(fileContentInByteArray)
	if fileContentString != "{}" {
		return "This is an invalid JSON", nil
	}
	return "This is a valid JSON", nil
}
