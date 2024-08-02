package endtoendtests

import (
	"os"
	"os/exec"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	buildCommand := exec.Command("go", "build", "cmd/cli/ccjsonparser.go")
	buildCommand.Dir = "./.."
	err := buildCommand.Run()
	if err != nil {
		panic(err)
	}

	result := m.Run()

	removeCommand := exec.Command("rm", "ccjsonparser")
	removeCommand.Dir = "./.."
	err = removeCommand.Run()
	if err != nil {
		panic(err)
	}
	os.Exit(result)
}

func TestFileReadingFailure(t *testing.T) {
	var out strings.Builder
	var errOut strings.Builder
	ccjsonparserCommand := exec.Command("./ccjsonparser", "tests/step4/invalid115.json")
	ccjsonparserCommand.Dir = "./.."
	ccjsonparserCommand.Stderr = &errOut
	ccjsonparserCommand.Stdout = &out
	err := ccjsonparserCommand.Run()
	assert.Error(t, err)
	assert.Contains(t, errOut.String(), "no such file")
	assert.Equal(t, "", out.String())
}

func TestNoFileNameProvided(t *testing.T) {
	var out strings.Builder
	var errOut strings.Builder
	ccjsonparserCommand := exec.Command("./ccjsonparser")
	ccjsonparserCommand.Dir = "./.."
	ccjsonparserCommand.Stderr = &errOut
	ccjsonparserCommand.Stdout = &out
	err := ccjsonparserCommand.Run()
	assert.Error(t, err)
	assert.Equal(t, errOut.String(), "Provide a file name\n")
	assert.Equal(t, "", out.String())
}

func TestMoreThanOneFileNameProvided(t *testing.T) {
	var out strings.Builder
	var errOut strings.Builder
	ccjsonparserCommand := exec.Command("./ccjsonparser", "filename1", "filename2")
	ccjsonparserCommand.Dir = "./.."
	ccjsonparserCommand.Stderr = &errOut
	ccjsonparserCommand.Stdout = &out
	err := ccjsonparserCommand.Run()
	assert.Error(t, err)
	assert.Equal(t, errOut.String(), "Provide just one file name\n")
	assert.Equal(t, "", out.String())
}

func TestJsonValidationOnJsonOrgTests(t *testing.T) {
	var out strings.Builder
	var errOut strings.Builder
	testFiles, err := os.ReadDir("./../tests/json_org_tests")
	if err != nil {
		t.Fatal(err)
	}
	for _, testFile := range testFiles {
		out.Reset()
		errOut.Reset()
		if strings.HasPrefix(testFile.Name(), "pass") {
			ccjsonparserCommand := exec.Command("./ccjsonparser", "tests/json_org_tests/"+testFile.Name())
			ccjsonparserCommand.Dir = "./.."
			ccjsonparserCommand.Stdout = &out
			ccjsonparserCommand.Stderr = &errOut
			err := ccjsonparserCommand.Run()
			assert.NoError(t, err)
			if !assert.Equal(t, "This is a valid JSON\n", out.String()) {
				t.Log(testFile.Name())
			}
		} else {
			ccjsonparserCommand := exec.Command("./ccjsonparser", "tests/json_org_tests/"+testFile.Name())
			ccjsonparserCommand.Dir = "./.."
			ccjsonparserCommand.Stdout = &out
			ccjsonparserCommand.Stderr = &errOut
			err := ccjsonparserCommand.Run()
			assert.Error(t, err)
			if !assert.Equal(t, "This is an invalid JSON\n", errOut.String()) {
				t.Log(testFile.Name())
			}
			assert.Equal(t, "", out.String())
		}
	}
}

func TestValidCases(t *testing.T) {
	var out strings.Builder
	var errOut strings.Builder

	for _, filePath := range validJSONTests {
		ccjsonparserCommand := exec.Command("./ccjsonparser", filePath)
		ccjsonparserCommand.Dir = "./.."
		out.Reset()
		errOut.Reset()
		ccjsonparserCommand.Stdout = &out
		ccjsonparserCommand.Stderr = &errOut
		err := ccjsonparserCommand.Run()
		if !assert.NoError(t, err) ||
			!assert.Equal(t, "This is a valid JSON\n", out.String()) ||
			!assert.Equal(t, "", errOut.String()) {
			t.Log(filePath)
		}
	}
}

func TestInvalidCases(t *testing.T) {
	var out strings.Builder
	var errOut strings.Builder

	for _, test := range invalidJSONTests {
		ccjsonparserCommand := exec.Command("./ccjsonparser", test.filePath)
		ccjsonparserCommand.Dir = "./.."
		out.Reset()
		errOut.Reset()
		ccjsonparserCommand.Stdout = &out
		ccjsonparserCommand.Stderr = &errOut
		err := ccjsonparserCommand.Run()
		if !assert.Equal(t, test.errResult, errOut.String()) ||
			!assert.Equal(t, "", out.String()) {
			t.Log(test)
		}
		assert.Equal(t, 1, err.(*exec.ExitError).ExitCode())
	}
}
