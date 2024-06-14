package endtoendtests

import (
	"os"
	"os/exec"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidation(t *testing.T) {
	var out strings.Builder
	var errOut strings.Builder
	ccjsonparserCommand := exec.Command("./ccjsonparser", "tests/step1/valid.json")
	ccjsonparserCommand.Dir = "./.."
	ccjsonparserCommand.Stdout = &out
	err := ccjsonparserCommand.Run()
	assert.NoError(t, err)
	assert.Equal(t, "This is a valid JSON\n", out.String())

	ccjsonparserCommand = exec.Command("./ccjsonparser", "tests/step1/invalid.json")
	ccjsonparserCommand.Dir = "./.."
	out.Reset()
	errOut.Reset()
	ccjsonparserCommand.Stderr = &errOut
	ccjsonparserCommand.Stdout = &out
	err = ccjsonparserCommand.Run()
	assert.Error(t, err)
	assert.Contains(t, errOut.String(), "This is an invalid JSON\nMUST be an object, array, number, or string, or false or null or true\n")
	assert.Equal(t, "", out.String())

	ccjsonparserCommand = exec.Command("./ccjsonparser", "tests/step2/valid.json")
	ccjsonparserCommand.Dir = "./.."
	out.Reset()
	ccjsonparserCommand.Stdout = &out
	err = ccjsonparserCommand.Run()
	assert.NoError(t, err)
	assert.Equal(t, "This is a valid JSON\n", out.String())

	ccjsonparserCommand = exec.Command("./ccjsonparser", "tests/step2/valid2.json")
	ccjsonparserCommand.Dir = "./.."
	out.Reset()
	ccjsonparserCommand.Stdout = &out
	err = ccjsonparserCommand.Run()
	assert.NoError(t, err)
	assert.Equal(t, "This is a valid JSON\n", out.String())

	ccjsonparserCommand = exec.Command("./ccjsonparser", "tests/step2/valid3.json")
	ccjsonparserCommand.Dir = "./.."
	out.Reset()
	ccjsonparserCommand.Stdout = &out
	err = ccjsonparserCommand.Run()
	assert.NoError(t, err)
	assert.Equal(t, "This is a valid JSON\n", out.String())

	ccjsonparserCommand = exec.Command("./ccjsonparser", "tests/step2/valid4.json")
	ccjsonparserCommand.Dir = "./.."
	out.Reset()
	ccjsonparserCommand.Stdout = &out
	err = ccjsonparserCommand.Run()
	assert.NoError(t, err)
	assert.Equal(t, "This is a valid JSON\n", out.String())

	ccjsonparserCommand = exec.Command("./ccjsonparser", "tests/step2/invalid.json")
	ccjsonparserCommand.Dir = "./.."
	out.Reset()
	errOut.Reset()
	ccjsonparserCommand.Stderr = &errOut
	ccjsonparserCommand.Stdout = &out
	err = ccjsonparserCommand.Run()
	assert.Error(t, err)
	assert.Equal(t, errOut.String(), "This is an invalid JSON\n")
	assert.Equal(t, "", out.String())

	ccjsonparserCommand = exec.Command("./ccjsonparser", "tests/step2/invalid2.json")
	ccjsonparserCommand.Dir = "./.."
	out.Reset()
	errOut.Reset()
	ccjsonparserCommand.Stderr = &errOut
	ccjsonparserCommand.Stdout = &out
	err = ccjsonparserCommand.Run()
	assert.Error(t, err)
	assert.Equal(t, errOut.String(), "This is an invalid JSON\n")
	assert.Equal(t, "", out.String())

	ccjsonparserCommand = exec.Command("./ccjsonparser", "tests/step2/invalid3.json")
	ccjsonparserCommand.Dir = "./.."
	out.Reset()
	errOut.Reset()
	ccjsonparserCommand.Stderr = &errOut
	ccjsonparserCommand.Stdout = &out
	err = ccjsonparserCommand.Run()
	assert.Error(t, err)
	assert.Equal(t, errOut.String(), "This is an invalid JSON\n")
	assert.Equal(t, "", out.String())

	ccjsonparserCommand = exec.Command("./ccjsonparser", "tests/step2/invalid4.json")
	ccjsonparserCommand.Dir = "./.."
	out.Reset()
	errOut.Reset()
	ccjsonparserCommand.Stderr = &errOut
	ccjsonparserCommand.Stdout = &out
	err = ccjsonparserCommand.Run()
	assert.Error(t, err)
	assert.Equal(t, errOut.String(), "This is an invalid JSON\n")
	assert.Equal(t, "", out.String())

	ccjsonparserCommand = exec.Command("./ccjsonparser", "tests/step2/invalid5.json")
	ccjsonparserCommand.Dir = "./.."
	out.Reset()
	errOut.Reset()
	ccjsonparserCommand.Stderr = &errOut
	ccjsonparserCommand.Stdout = &out
	err = ccjsonparserCommand.Run()
	assert.Error(t, err)
	assert.Equal(t, errOut.String(), "This is an invalid JSON\n")
	assert.Equal(t, "", out.String())

	ccjsonparserCommand = exec.Command("./ccjsonparser", "tests/step2/invalid6.json")
	ccjsonparserCommand.Dir = "./.."
	out.Reset()
	errOut.Reset()
	ccjsonparserCommand.Stderr = &errOut
	ccjsonparserCommand.Stdout = &out
	err = ccjsonparserCommand.Run()
	assert.Error(t, err)
	assert.Equal(t, errOut.String(), "This is an invalid JSON\n")
	assert.Equal(t, "", out.String())

	ccjsonparserCommand = exec.Command("./ccjsonparser", "tests/step2/invalid7.json")
	ccjsonparserCommand.Dir = "./.."
	out.Reset()
	errOut.Reset()
	ccjsonparserCommand.Stderr = &errOut
	ccjsonparserCommand.Stdout = &out
	err = ccjsonparserCommand.Run()
	assert.Error(t, err)
	assert.Equal(t, errOut.String(), "This is an invalid JSON\n")
	assert.Equal(t, "", out.String())

	ccjsonparserCommand = exec.Command("./ccjsonparser", "tests/step2/invalid8.json")
	ccjsonparserCommand.Dir = "./.."
	out.Reset()
	errOut.Reset()
	ccjsonparserCommand.Stderr = &errOut
	ccjsonparserCommand.Stdout = &out
	err = ccjsonparserCommand.Run()
	assert.Error(t, err)
	assert.Equal(t, errOut.String(), "This is an invalid JSON\nMultiple values outside of an array\n")
	assert.Equal(t, "", out.String())

	ccjsonparserCommand = exec.Command("./ccjsonparser", "tests/step2/invalid9.json")
	ccjsonparserCommand.Dir = "./.."
	out.Reset()
	errOut.Reset()
	ccjsonparserCommand.Stderr = &errOut
	ccjsonparserCommand.Stdout = &out
	err = ccjsonparserCommand.Run()
	assert.Error(t, err)
	assert.Equal(t, errOut.String(), "This is an invalid JSON\nThis is an invalid string\n")
	assert.Equal(t, "", out.String())

	ccjsonparserCommand = exec.Command("./ccjsonparser", "tests/step2/invalid10.json")
	ccjsonparserCommand.Dir = "./.."
	out.Reset()
	errOut.Reset()
	ccjsonparserCommand.Stderr = &errOut
	ccjsonparserCommand.Stdout = &out
	err = ccjsonparserCommand.Run()
	assert.Error(t, err)
	assert.Equal(t, errOut.String(), "This is an invalid JSON\nThis is an invalid string\n")
	assert.Equal(t, "", out.String())

	ccjsonparserCommand = exec.Command("./ccjsonparser", "tests/step2/invalid11.json")
	ccjsonparserCommand.Dir = "./.."
	out.Reset()
	errOut.Reset()
	ccjsonparserCommand.Stderr = &errOut
	ccjsonparserCommand.Stdout = &out
	err = ccjsonparserCommand.Run()
	assert.Error(t, err)
	assert.Equal(t, errOut.String(), "This is an invalid JSON\nThis is an invalid string\n")
	assert.Equal(t, "", out.String())

	ccjsonparserCommand = exec.Command("./ccjsonparser", "tests/step2/invalid12.json")
	ccjsonparserCommand.Dir = "./.."
	out.Reset()
	errOut.Reset()
	ccjsonparserCommand.Stderr = &errOut
	ccjsonparserCommand.Stdout = &out
	err = ccjsonparserCommand.Run()
	assert.Error(t, err)
	assert.Equal(t, errOut.String(), "This is an invalid JSON\nThis is an invalid string\n")
	assert.Equal(t, "", out.String())

	ccjsonparserCommand = exec.Command("./ccjsonparser", "tests/step2/invalid13.json")
	ccjsonparserCommand.Dir = "./.."
	out.Reset()
	errOut.Reset()
	ccjsonparserCommand.Stderr = &errOut
	ccjsonparserCommand.Stdout = &out
	err = ccjsonparserCommand.Run()
	assert.Error(t, err)
	assert.Equal(t, errOut.String(), "This is an invalid JSON\nThis is an invalid string\n")
	assert.Equal(t, "", out.String())

	ccjsonparserCommand = exec.Command("./ccjsonparser", "tests/step2/invalid14.json")
	ccjsonparserCommand.Dir = "./.."
	out.Reset()
	errOut.Reset()
	ccjsonparserCommand.Stderr = &errOut
	ccjsonparserCommand.Stdout = &out
	err = ccjsonparserCommand.Run()
	assert.Error(t, err)
	assert.Equal(t, errOut.String(), "This is an invalid JSON\nThis is an invalid string\n")
	assert.Equal(t, "", out.String())

	ccjsonparserCommand = exec.Command("./ccjsonparser", "tests/step2/fail10.json")
	ccjsonparserCommand.Dir = "./.."
	out.Reset()
	errOut.Reset()
	ccjsonparserCommand.Stderr = &errOut
	ccjsonparserCommand.Stdout = &out
	err = ccjsonparserCommand.Run()
	assert.Error(t, err)
	assert.Contains(t, errOut.String(), "This is an invalid JSON\nMultiple values outside of an array\n")
	assert.Equal(t, "", out.String())

	ccjsonparserCommand = exec.Command("./ccjsonparser", "tests/step3/valid.json")
	ccjsonparserCommand.Dir = "./.."
	out.Reset()
	ccjsonparserCommand.Stdout = &out
	err = ccjsonparserCommand.Run()
	assert.NoError(t, err)
	assert.Equal(t, "This is a valid JSON\n", out.String())

	ccjsonparserCommand = exec.Command("./ccjsonparser", "tests/step3/valid2.json")
	ccjsonparserCommand.Dir = "./.."
	out.Reset()
	ccjsonparserCommand.Stdout = &out
	err = ccjsonparserCommand.Run()
	assert.NoError(t, err)
	assert.Equal(t, "This is a valid JSON\n", out.String())

	ccjsonparserCommand = exec.Command("./ccjsonparser", "tests/step3/valid3.json")
	ccjsonparserCommand.Dir = "./.."
	out.Reset()
	ccjsonparserCommand.Stdout = &out
	err = ccjsonparserCommand.Run()
	assert.NoError(t, err)
	assert.Equal(t, "This is a valid JSON\n", out.String())

	ccjsonparserCommand = exec.Command("./ccjsonparser", "tests/step3/valid4.json")
	ccjsonparserCommand.Dir = "./.."
	out.Reset()
	ccjsonparserCommand.Stdout = &out
	err = ccjsonparserCommand.Run()
	assert.NoError(t, err)
	assert.Equal(t, "This is a valid JSON\n", out.String())

	ccjsonparserCommand = exec.Command("./ccjsonparser", "tests/step3/valid5.json")
	ccjsonparserCommand.Dir = "./.."
	out.Reset()
	ccjsonparserCommand.Stdout = &out
	err = ccjsonparserCommand.Run()
	assert.NoError(t, err)
	assert.Equal(t, "This is a valid JSON\n", out.String())

	ccjsonparserCommand = exec.Command("./ccjsonparser", "tests/step3/valid6.json")
	ccjsonparserCommand.Dir = "./.."
	out.Reset()
	ccjsonparserCommand.Stdout = &out
	err = ccjsonparserCommand.Run()
	assert.NoError(t, err)
	assert.Equal(t, "This is a valid JSON\n", out.String())

	ccjsonparserCommand = exec.Command("./ccjsonparser", "tests/step3/valid7.json")
	ccjsonparserCommand.Dir = "./.."
	out.Reset()
	ccjsonparserCommand.Stdout = &out
	err = ccjsonparserCommand.Run()
	assert.NoError(t, err)
	assert.Equal(t, "This is a valid JSON\n", out.String())

	ccjsonparserCommand = exec.Command("./ccjsonparser", "tests/step3/valid8.json")
	ccjsonparserCommand.Dir = "./.."
	out.Reset()
	ccjsonparserCommand.Stdout = &out
	err = ccjsonparserCommand.Run()
	assert.NoError(t, err)
	assert.Equal(t, "This is a valid JSON\n", out.String())

	ccjsonparserCommand = exec.Command("./ccjsonparser", "tests/step3/valid9.json")
	ccjsonparserCommand.Dir = "./.."
	out.Reset()
	ccjsonparserCommand.Stdout = &out
	err = ccjsonparserCommand.Run()
	assert.NoError(t, err)
	assert.Equal(t, "This is a valid JSON\n", out.String())

	ccjsonparserCommand = exec.Command("./ccjsonparser", "tests/step3/valid10.json")
	ccjsonparserCommand.Dir = "./.."
	out.Reset()
	ccjsonparserCommand.Stdout = &out
	err = ccjsonparserCommand.Run()
	assert.NoError(t, err)
	assert.Equal(t, "This is a valid JSON\n", out.String())

	ccjsonparserCommand = exec.Command("./ccjsonparser", "tests/step3/invalid.json")
	ccjsonparserCommand.Dir = "./.."
	out.Reset()
	errOut.Reset()
	ccjsonparserCommand.Stderr = &errOut
	ccjsonparserCommand.Stdout = &out
	err = ccjsonparserCommand.Run()
	assert.Error(t, err)
	assert.Equal(t, errOut.String(), "This is an invalid JSON\n")
	assert.Equal(t, "", out.String())

	ccjsonparserCommand = exec.Command("./ccjsonparser", "tests/step3/invalid2.json")
	ccjsonparserCommand.Dir = "./.."
	out.Reset()
	errOut.Reset()
	ccjsonparserCommand.Stderr = &errOut
	ccjsonparserCommand.Stdout = &out
	err = ccjsonparserCommand.Run()
	assert.Error(t, err)
	assert.Equal(t, errOut.String(), "This is an invalid JSON\n")
	assert.Equal(t, "", out.String())

	ccjsonparserCommand = exec.Command("./ccjsonparser", "tests/step3/invalid3.json")
	ccjsonparserCommand.Dir = "./.."
	out.Reset()
	errOut.Reset()
	ccjsonparserCommand.Stderr = &errOut
	ccjsonparserCommand.Stdout = &out
	err = ccjsonparserCommand.Run()
	assert.Error(t, err)
	assert.Contains(t, errOut.String(), "This is an invalid JSON\n")
	assert.Equal(t, "", out.String())

	ccjsonparserCommand = exec.Command("./ccjsonparser", "tests/step3/invalid4.json")
	ccjsonparserCommand.Dir = "./.."
	out.Reset()
	errOut.Reset()
	ccjsonparserCommand.Stderr = &errOut
	ccjsonparserCommand.Stdout = &out
	err = ccjsonparserCommand.Run()
	assert.Error(t, err)
	assert.Contains(t, errOut.String(), "This is an invalid JSON\n")
	assert.Equal(t, "", out.String())

	ccjsonparserCommand = exec.Command("./ccjsonparser", "tests/step3/invalid5.json")
	ccjsonparserCommand.Dir = "./.."
	out.Reset()
	errOut.Reset()
	ccjsonparserCommand.Stderr = &errOut
	ccjsonparserCommand.Stdout = &out
	err = ccjsonparserCommand.Run()
	assert.Error(t, err)
	assert.Contains(t, errOut.String(), "This is an invalid JSON\n")
	assert.Equal(t, "", out.String())

	ccjsonparserCommand = exec.Command("./ccjsonparser", "tests/step3/invalid6.json")
	ccjsonparserCommand.Dir = "./.."
	out.Reset()
	errOut.Reset()
	ccjsonparserCommand.Stderr = &errOut
	ccjsonparserCommand.Stdout = &out
	err = ccjsonparserCommand.Run()
	assert.Error(t, err)
	assert.Contains(t, errOut.String(), "This is an invalid JSON\n")
	assert.Equal(t, "", out.String())

	ccjsonparserCommand = exec.Command("./ccjsonparser", "tests/step3/invalid7.json")
	ccjsonparserCommand.Dir = "./.."
	out.Reset()
	errOut.Reset()
	ccjsonparserCommand.Stderr = &errOut
	ccjsonparserCommand.Stdout = &out
	err = ccjsonparserCommand.Run()
	assert.Error(t, err)
	assert.Contains(t, errOut.String(), "This is an invalid JSON\n")
	assert.Equal(t, "", out.String())

	ccjsonparserCommand = exec.Command("./ccjsonparser", "tests/step3/invalid8.json")
	ccjsonparserCommand.Dir = "./.."
	out.Reset()
	errOut.Reset()
	ccjsonparserCommand.Stderr = &errOut
	ccjsonparserCommand.Stdout = &out
	err = ccjsonparserCommand.Run()
	assert.Error(t, err)
	assert.Contains(t, errOut.String(), "This is an invalid JSON\n")
	assert.Equal(t, "", out.String())

	ccjsonparserCommand = exec.Command("./ccjsonparser", "tests/step3/invalid9.json")
	ccjsonparserCommand.Dir = "./.."
	out.Reset()
	errOut.Reset()
	ccjsonparserCommand.Stderr = &errOut
	ccjsonparserCommand.Stdout = &out
	err = ccjsonparserCommand.Run()
	assert.Error(t, err)
	assert.Contains(t, errOut.String(), "This is an invalid JSON\n")
	assert.Equal(t, "", out.String())

	ccjsonparserCommand = exec.Command("./ccjsonparser", "tests/step3/invalid10.json")
	ccjsonparserCommand.Dir = "./.."
	out.Reset()
	errOut.Reset()
	ccjsonparserCommand.Stderr = &errOut
	ccjsonparserCommand.Stdout = &out
	err = ccjsonparserCommand.Run()
	assert.Error(t, err)
	assert.Contains(t, errOut.String(), "This is an invalid JSON\n")
	assert.Equal(t, "", out.String())

	ccjsonparserCommand = exec.Command("./ccjsonparser", "tests/step3/invalid11.json")
	ccjsonparserCommand.Dir = "./.."
	out.Reset()
	errOut.Reset()
	ccjsonparserCommand.Stderr = &errOut
	ccjsonparserCommand.Stdout = &out
	err = ccjsonparserCommand.Run()
	assert.Error(t, err)
	assert.Equal(t, errOut.String(), "This is an invalid JSON\nShould be \"null\"\n")
	assert.Equal(t, "", out.String())

	ccjsonparserCommand = exec.Command("./ccjsonparser", "tests/step3/invalid12.json")
	ccjsonparserCommand.Dir = "./.."
	out.Reset()
	errOut.Reset()
	ccjsonparserCommand.Stderr = &errOut
	ccjsonparserCommand.Stdout = &out
	err = ccjsonparserCommand.Run()
	assert.Error(t, err)
	assert.Equal(t, errOut.String(), "This is an invalid JSON\nShould be \"false\"\n")
	assert.Equal(t, "", out.String())

	ccjsonparserCommand = exec.Command("./ccjsonparser", "tests/step3/invalid13.json")
	ccjsonparserCommand.Dir = "./.."
	out.Reset()
	errOut.Reset()
	ccjsonparserCommand.Stderr = &errOut
	ccjsonparserCommand.Stdout = &out
	err = ccjsonparserCommand.Run()
	assert.Error(t, err)
	assert.Equal(t, "This is an invalid JSON\nShould be \"true\"\n", errOut.String())
	assert.Equal(t, "", out.String())

	ccjsonparserCommand = exec.Command("./ccjsonparser", "tests/step4/valid.json")
	ccjsonparserCommand.Dir = "./.."
	out.Reset()
	ccjsonparserCommand.Stdout = &out
	err = ccjsonparserCommand.Run()
	assert.NoError(t, err)
	assert.Equal(t, "This is a valid JSON\n", out.String())

	ccjsonparserCommand = exec.Command("./ccjsonparser", "tests/step4/valid2.json")
	ccjsonparserCommand.Dir = "./.."
	out.Reset()
	ccjsonparserCommand.Stdout = &out
	err = ccjsonparserCommand.Run()
	assert.NoError(t, err)
	assert.Equal(t, "This is a valid JSON\n", out.String())

	ccjsonparserCommand = exec.Command("./ccjsonparser", "tests/step4/valid3.json")
	ccjsonparserCommand.Dir = "./.."
	out.Reset()
	ccjsonparserCommand.Stdout = &out
	err = ccjsonparserCommand.Run()
	assert.NoError(t, err)
	assert.Equal(t, "This is a valid JSON\n", out.String())

	ccjsonparserCommand = exec.Command("./ccjsonparser", "tests/step4/valid4.json")
	ccjsonparserCommand.Dir = "./.."
	out.Reset()
	ccjsonparserCommand.Stdout = &out
	err = ccjsonparserCommand.Run()
	assert.NoError(t, err)
	assert.Equal(t, "This is a valid JSON\n", out.String())

	ccjsonparserCommand = exec.Command("./ccjsonparser", "tests/step4/valid5.json")
	ccjsonparserCommand.Dir = "./.."
	out.Reset()
	ccjsonparserCommand.Stdout = &out
	err = ccjsonparserCommand.Run()
	assert.NoError(t, err)
	assert.Equal(t, "This is a valid JSON\n", out.String())

	ccjsonparserCommand = exec.Command("./ccjsonparser", "tests/step4/valid6.json")
	ccjsonparserCommand.Dir = "./.."
	out.Reset()
	ccjsonparserCommand.Stdout = &out
	err = ccjsonparserCommand.Run()
	assert.NoError(t, err)
	assert.Equal(t, "This is a valid JSON\n", out.String())

	ccjsonparserCommand = exec.Command("./ccjsonparser", "tests/step4/invalid.json")
	ccjsonparserCommand.Dir = "./.."
	out.Reset()
	errOut.Reset()
	ccjsonparserCommand.Stderr = &errOut
	ccjsonparserCommand.Stdout = &out
	err = ccjsonparserCommand.Run()
	assert.Error(t, err)
	assert.Equal(t, errOut.String(), "This is an invalid JSON\n")
	assert.Equal(t, "", out.String())

	ccjsonparserCommand = exec.Command("./ccjsonparser", "tests/step4/invalid2.json")
	ccjsonparserCommand.Dir = "./.."
	out.Reset()
	errOut.Reset()
	ccjsonparserCommand.Stderr = &errOut
	ccjsonparserCommand.Stdout = &out
	err = ccjsonparserCommand.Run()
	assert.Error(t, err)
	assert.Equal(t, errOut.String(), "This is an invalid JSON\n")
	assert.Equal(t, "", out.String())

	ccjsonparserCommand = exec.Command("./ccjsonparser", "tests/step4/invalid3.json")
	ccjsonparserCommand.Dir = "./.."
	out.Reset()
	errOut.Reset()
	ccjsonparserCommand.Stderr = &errOut
	ccjsonparserCommand.Stdout = &out
	err = ccjsonparserCommand.Run()
	assert.Error(t, err)
	assert.Equal(t, errOut.String(), "This is an invalid JSON\n")
	assert.Equal(t, "", out.String())

	ccjsonparserCommand = exec.Command("./ccjsonparser", "tests/step4/invalid4.json")
	ccjsonparserCommand.Dir = "./.."
	out.Reset()
	errOut.Reset()
	ccjsonparserCommand.Stderr = &errOut
	ccjsonparserCommand.Stdout = &out
	err = ccjsonparserCommand.Run()
	assert.Error(t, err)
	assert.Equal(t, errOut.String(), "This is an invalid JSON\n")
	assert.Equal(t, "", out.String())

	ccjsonparserCommand = exec.Command("./ccjsonparser", "tests/step4/invalid5.json")
	ccjsonparserCommand.Dir = "./.."
	out.Reset()
	errOut.Reset()
	ccjsonparserCommand.Stderr = &errOut
	ccjsonparserCommand.Stdout = &out
	err = ccjsonparserCommand.Run()
	assert.Error(t, err)
	assert.Equal(t, errOut.String(), "This is an invalid JSON\n")
	assert.Equal(t, "", out.String())

	ccjsonparserCommand = exec.Command("./ccjsonparser", "tests/step4/invalid115.json")
	ccjsonparserCommand.Dir = "./.."
	out.Reset()
	errOut.Reset()
	ccjsonparserCommand.Stderr = &errOut
	ccjsonparserCommand.Stdout = &out
	err = ccjsonparserCommand.Run()
	assert.Error(t, err)
	assert.Contains(t, errOut.String(), "no such file")
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
