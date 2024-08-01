package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/ahmedsameha1/ccjsonparser/internal/app"
)

func main() {
	fileContent, err := readFileContent(os.ReadFile, os.Args)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	err = app.Validate(fileContent)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Println("This is a valid JSON")
}

func readFileContent(readFile func(string) ([]byte, error), args []string) (string, error) {
	if len(args) < 2 {
		return "", errors.New("Provide a file name")
	}
	if len(args) > 2 {
		return "", errors.New("Provide just one file name")
	}
	fileContent, err := readFile(args[1])
	return string(fileContent), err
}
