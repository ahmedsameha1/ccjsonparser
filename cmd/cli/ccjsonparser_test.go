package main

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadFileContent(t *testing.T) {
	_, err := readFileContent(func(name string) ([]byte, error) {
		if name != "doesntexist.json" {
			panic("error")
		}
		return nil, errors.New("no such file")
	}, []string{"ccjsonparser", "doesntexist.json"})
	assert.Contains(t, err.Error(), "no such file")

	_,err = readFileContent(func(name string) ([]byte, error) {
		return nil, nil
	}, []string{"ccjsonparser"})
	assert.Equal(t, err.Error(), "Provide a file name")

	_,err = readFileContent(func(name string) ([]byte, error) {
		return nil, nil
	}, []string{"ccjsonparser", "filename1", "filename2"})
	assert.Equal(t, err.Error(), "Provide just one file name")

	fileContent, err := readFileContent(func(name string) ([]byte, error) {
		if name != "filename.json" {
			panic("error")
		}
		return []byte("[]"), nil
	}, []string{"ccjsonparser", "filename.json"})
	assert.Equal(t, err, nil)
	assert.Equal(t, fileContent, "[]")
}
