package endtoendtests

import (
	"os"
	"os/exec"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidation(t *testing.T) {
	ccjsonparserCommand := exec.Command("./ccjsonparser", "tests/step1/valid.json")
	ccjsonparserCommand.Dir = "./.."
	var out strings.Builder
	ccjsonparserCommand.Stdout = &out
	err := ccjsonparserCommand.Run()
	assert.NoError(t, err)
	assert.Equal(t, "This is a valid JSON\n", out.String())

	ccjsonparserCommand = exec.Command("./ccjsonparser", "tests/step1/invalid.json")
	ccjsonparserCommand.Dir = "./.."
	out.Reset()
	ccjsonparserCommand.Stdout = &out
	err = ccjsonparserCommand.Run()
	assert.NoError(t, err)
	assert.Equal(t, "MUST be an object, array, number, or string, or false or null or true\n", out.String())

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
	ccjsonparserCommand.Stdout = &out
	err = ccjsonparserCommand.Run()
	assert.NoError(t, err)
	assert.Equal(t, "This is an invalid JSON\n", out.String())

	ccjsonparserCommand = exec.Command("./ccjsonparser", "tests/step2/invalid2.json")
	ccjsonparserCommand.Dir = "./.."
	out.Reset()
	ccjsonparserCommand.Stdout = &out
	err = ccjsonparserCommand.Run()
	assert.NoError(t, err)
	assert.Equal(t, "This is an invalid JSON\n", out.String())

	ccjsonparserCommand = exec.Command("./ccjsonparser", "tests/step2/invalid3.json")
	ccjsonparserCommand.Dir = "./.."
	out.Reset()
	ccjsonparserCommand.Stdout = &out
	err = ccjsonparserCommand.Run()
	assert.NoError(t, err)
	assert.Equal(t, "This is an invalid JSON\n", out.String())

	ccjsonparserCommand = exec.Command("./ccjsonparser", "tests/step2/invalid4.json")
	ccjsonparserCommand.Dir = "./.."
	out.Reset()
	ccjsonparserCommand.Stdout = &out
	err = ccjsonparserCommand.Run()
	assert.NoError(t, err)
	assert.Equal(t, "This is an invalid JSON\n", out.String())

	ccjsonparserCommand = exec.Command("./ccjsonparser", "tests/step2/invalid5.json")
	ccjsonparserCommand.Dir = "./.."
	out.Reset()
	ccjsonparserCommand.Stdout = &out
	err = ccjsonparserCommand.Run()
	assert.NoError(t, err)
	assert.Equal(t, "This is an invalid JSON\n", out.String())

	ccjsonparserCommand = exec.Command("./ccjsonparser", "tests/step2/invalid6.json")
	ccjsonparserCommand.Dir = "./.."
	out.Reset()
	ccjsonparserCommand.Stdout = &out
	err = ccjsonparserCommand.Run()
	assert.NoError(t, err)
	assert.Equal(t, "This is an invalid JSON\n", out.String())

	ccjsonparserCommand = exec.Command("./ccjsonparser", "tests/step2/invalid7.json")
	ccjsonparserCommand.Dir = "./.."
	out.Reset()
	ccjsonparserCommand.Stdout = &out
	err = ccjsonparserCommand.Run()
	assert.NoError(t, err)
	assert.Equal(t, "This is an invalid JSON\n", out.String())

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

	ccjsonparserCommand = exec.Command("./ccjsonparser", "tests/step3/invalid.json")
	ccjsonparserCommand.Dir = "./.."
	out.Reset()
	ccjsonparserCommand.Stdout = &out
	err = ccjsonparserCommand.Run()
	assert.NoError(t, err)
	assert.Equal(t, "This is an invalid JSON\n", out.String())

	ccjsonparserCommand = exec.Command("./ccjsonparser", "tests/step3/invalid2.json")
	ccjsonparserCommand.Dir = "./.."
	out.Reset()
	ccjsonparserCommand.Stdout = &out
	err = ccjsonparserCommand.Run()
	assert.NoError(t, err)
	assert.Equal(t, "This is an invalid JSON\n", out.String())

	ccjsonparserCommand = exec.Command("./ccjsonparser", "tests/step3/invalid3.json")
	ccjsonparserCommand.Dir = "./.."
	out.Reset()
	ccjsonparserCommand.Stdout = &out
	err = ccjsonparserCommand.Run()
	assert.NoError(t, err)
	assert.Equal(t, "This is an invalid JSON\n", out.String())

	ccjsonparserCommand = exec.Command("./ccjsonparser", "tests/step3/invalid4.json")
	ccjsonparserCommand.Dir = "./.."
	out.Reset()
	ccjsonparserCommand.Stdout = &out
	err = ccjsonparserCommand.Run()
	assert.NoError(t, err)
	assert.Equal(t, "This is an invalid JSON\n", out.String())

	ccjsonparserCommand = exec.Command("./ccjsonparser", "tests/step3/invalid5.json")
	ccjsonparserCommand.Dir = "./.."
	out.Reset()
	ccjsonparserCommand.Stdout = &out
	err = ccjsonparserCommand.Run()
	assert.NoError(t, err)
	assert.Equal(t, "This is an invalid JSON\n", out.String())

	ccjsonparserCommand = exec.Command("./ccjsonparser", "tests/step3/invalid6.json")
	ccjsonparserCommand.Dir = "./.."
	out.Reset()
	ccjsonparserCommand.Stdout = &out
	err = ccjsonparserCommand.Run()
	assert.NoError(t, err)
	assert.Equal(t, "This is an invalid JSON\n", out.String())

	ccjsonparserCommand = exec.Command("./ccjsonparser", "tests/step3/invalid7.json")
	ccjsonparserCommand.Dir = "./.."
	out.Reset()
	ccjsonparserCommand.Stdout = &out
	err = ccjsonparserCommand.Run()
	assert.NoError(t, err)
	assert.Equal(t, "This is an invalid JSON\n", out.String())

	ccjsonparserCommand = exec.Command("./ccjsonparser", "tests/step3/invalid8.json")
	ccjsonparserCommand.Dir = "./.."
	out.Reset()
	ccjsonparserCommand.Stdout = &out
	err = ccjsonparserCommand.Run()
	assert.NoError(t, err)
	assert.Equal(t, "This is an invalid JSON\n", out.String())

	ccjsonparserCommand = exec.Command("./ccjsonparser", "tests/step3/invalid9.json")
	ccjsonparserCommand.Dir = "./.."
	out.Reset()
	ccjsonparserCommand.Stdout = &out
	err = ccjsonparserCommand.Run()
	assert.NoError(t, err)
	assert.Equal(t, "This is an invalid JSON\n", out.String())

	ccjsonparserCommand = exec.Command("./ccjsonparser", "tests/step3/invalid10.json")
	ccjsonparserCommand.Dir = "./.."
	out.Reset()
	ccjsonparserCommand.Stdout = &out
	err = ccjsonparserCommand.Run()
	assert.NoError(t, err)
	assert.Equal(t, "This is an invalid JSON\n", out.String())

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
	ccjsonparserCommand.Stdout = &out
	err = ccjsonparserCommand.Run()
	assert.NoError(t, err)
	assert.Equal(t, "This is an invalid JSON\n", out.String())

	ccjsonparserCommand = exec.Command("./ccjsonparser", "tests/step4/invalid2.json")
	ccjsonparserCommand.Dir = "./.."
	out.Reset()
	ccjsonparserCommand.Stdout = &out
	err = ccjsonparserCommand.Run()
	assert.NoError(t, err)
	assert.Equal(t, "This is an invalid JSON\n", out.String())

	ccjsonparserCommand = exec.Command("./ccjsonparser", "tests/step4/invalid3.json")
	ccjsonparserCommand.Dir = "./.."
	out.Reset()
	ccjsonparserCommand.Stdout = &out
	err = ccjsonparserCommand.Run()
	assert.NoError(t, err)
	assert.Equal(t, "This is an invalid JSON\n", out.String())

	ccjsonparserCommand = exec.Command("./ccjsonparser", "tests/step4/invalid4.json")
	ccjsonparserCommand.Dir = "./.."
	out.Reset()
	ccjsonparserCommand.Stdout = &out
	err = ccjsonparserCommand.Run()
	assert.NoError(t, err)
	assert.Equal(t, "This is an invalid JSON\n", out.String())

	ccjsonparserCommand = exec.Command("./ccjsonparser", "tests/step4/invalid5.json")
	ccjsonparserCommand.Dir = "./.."
	out.Reset()
	ccjsonparserCommand.Stdout = &out
	err = ccjsonparserCommand.Run()
	assert.NoError(t, err)
	assert.Equal(t, "This is an invalid JSON\n", out.String())

	ccjsonparserCommand = exec.Command("./ccjsonparser", "tests/step4/invalid115.json")
	ccjsonparserCommand.Dir = "./.."
	var errOut strings.Builder
	ccjsonparserCommand.Stderr = &errOut
	out.Reset()
	ccjsonparserCommand.Stdout = &out
	err = ccjsonparserCommand.Run()
	assert.Error(t, err)
	assert.Contains(t, errOut.String(), "no such file")
	assert.Equal(t, "", out.String())
}

func TestJsonValidationOnJsonOrgTests(t *testing.T) {
	testFiles, err := os.ReadDir("./../tests/json_org_tests")
	if err != nil {
		t.Fatal(err)
	}
	for _, testFile := range testFiles {
		if strings.HasPrefix(testFile.Name(), "pass") {
			ccjsonparserCommand := exec.Command("./ccjsonparser", "tests/json_org_tests/"+testFile.Name())
			ccjsonparserCommand.Dir = "./.."
			var out strings.Builder
			ccjsonparserCommand.Stdout = &out
			err := ccjsonparserCommand.Run()
			assert.NoError(t, err)
			if !assert.Equal(t, "This is a valid JSON\n", out.String()) {
				t.Log(testFile.Name())
			}
		} else {
			ccjsonparserCommand := exec.Command("./ccjsonparser", "tests/json_org_tests/"+testFile.Name())
			ccjsonparserCommand.Dir = "./.."
			var out strings.Builder
			ccjsonparserCommand.Stdout = &out
			err := ccjsonparserCommand.Run()
			assert.NoError(t, err)
			if !assert.Equal(t, "This is an invalid JSON\n", out.String()) {
				t.Log(testFile.Name())
			}
		}
	}
}
