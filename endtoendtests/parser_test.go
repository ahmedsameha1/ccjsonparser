package endtoendtests

import (
	"os"
	"os/exec"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidCases(t *testing.T) {
	var out strings.Builder
	var errOut strings.Builder
	var validJSONTests = []string{
		"tests/step1/valid.json",
		"tests/step2/valid.json",
		"tests/step2/valid2.json",
		"tests/step2/valid3.json",
		"tests/step2/valid4.json",
		"tests/step2/valid5.json",
		"tests/step2/valid6.json",
		"tests/step2/valid7.json",
		"tests/step2/valid8.json",
		"tests/step2/valid9.json",
		"tests/step2/valid10.json",
		"tests/step2/valid11.json",
		"tests/step2/valid12.json",
		"tests/step2/valid13.json",
		"tests/step2/valid14.json",
		"tests/step2/valid15.json",
		"tests/step2/valid16.json",
		"tests/step2/valid17.json",
		"tests/step3/valid.json",
		"tests/step3/valid2.json",
		"tests/step3/valid3.json",
		"tests/step3/valid4.json",
		"tests/step3/valid5.json",
		"tests/step3/valid6.json",
		"tests/step3/valid7.json",
		"tests/step3/valid8.json",
		"tests/step3/valid9.json",
		"tests/step3/valid10.json",
		"tests/step3/valid11.json",
		"tests/step3/valid12.json",
		"tests/step3/valid13.json",
		"tests/step4/valid.json",
		"tests/step4/valid2.json",
		"tests/step4/valid3.json",
		"tests/step4/valid4.json",
		"tests/step4/valid5.json",
		"tests/step4/valid6.json",
		"tests/step4/valid7.json",
		"tests/step4/valid8.json",
		"tests/step4/valid9.json",
		"tests/step4/valid10.json",
		"tests/step4/valid11.json",
		"tests/step4/valid12.json",
		"tests/huge_json/valid.json",
	}

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
	var invalidJSONTests = []struct {
		filePath  string
		errResult string
	}{
		{filePath: "tests/step1/invalid.json",
			errResult: "This is an invalid JSON\nMUST be an object, array, number, or string, or false or null or true\n"},
		{filePath: "tests/step2/invalid.json",
			errResult: "This is an invalid JSON\nThere is an object that contains an extra tail comma(s)\n"},
		{filePath: "tests/step2/invalid2.json",
			errResult: "This is an invalid JSON\nThere is a string that is not surrounded correctly by (\"\")\n"},
		{filePath: "tests/step2/invalid3.json",
			errResult: "This is an invalid JSON\nThere is a string that is not surrounded correctly by (\"\")\n"},
		{filePath: "tests/step2/invalid4.json",
			errResult: "This is an invalid JSON\nThere is a string that is not surrounded correctly by (\"\")\n"},
		{filePath: "tests/step2/invalid5.json",
			errResult: "This is an invalid JSON\nThere is a string that contains tabs or new lines or Illegal backslash escapes\n"},
		{filePath: "tests/step2/invalid6.json",
			errResult: "This is an invalid JSON\nThere is a string that is not surrounded correctly by (\"\")\n"},
		{filePath: "tests/step2/invalid8.json",
			errResult: "This is an invalid JSON\nMultiple values outside of an object or array\n"},
		{filePath: "tests/step2/invalid9.json",
			errResult: "This is an invalid JSON\nThere is a string that contains tabs or new lines or Illegal backslash escapes\n"},
		{filePath: "tests/step2/invalid10.json",
			errResult: "This is an invalid JSON\nThere is a string that contains tabs or new lines or Illegal backslash escapes\n"},
		{filePath: "tests/step2/invalid11.json",
			errResult: "This is an invalid JSON\nThere is a string that contains tabs or new lines or Illegal backslash escapes\n"},
		{filePath: "tests/step2/invalid12.json",
			errResult: "This is an invalid JSON\nThere is a string that contains tabs or new lines or Illegal backslash escapes\n"},
		{filePath: "tests/step2/invalid13.json",
			errResult: "This is an invalid JSON\nThere is a string that contains tabs or new lines or Illegal backslash escapes\n"},
		{filePath: "tests/step2/invalid14.json",
			errResult: "This is an invalid JSON\nThere is a string that contains tabs or new lines or Illegal backslash escapes\n"},
		{filePath: "tests/step2/invalid15.json",
			errResult: "This is an invalid JSON\nThis is an array that is surrounded by invalid \"][}{\"\n"},
		{filePath: "tests/step2/invalid16.json",
			errResult: "This is an invalid JSON\nThis is an array that is surrounded by invalid \"][}{\"\n"},
		{filePath: "tests/step2/invalid17.json",
			errResult: "This is an invalid JSON\nThis is an array that is surrounded by invalid \"][}{\"\n"},
		{filePath: "tests/step2/invalid18.json",
			errResult: "This is an invalid JSON\nThis is an array that is surrounded by invalid \"][}{\"\n"},
		{filePath: "tests/step2/invalid19.json",
			errResult: "This is an invalid JSON\nThis is an array that is surrounded by invalid \"][}{\"\n"},
		{filePath: "tests/step2/invalid20.json",
			errResult: "This is an invalid JSON\nThis is an array that is surrounded by invalid \"][}{\"\n"},
		{filePath: "tests/step2/invalid21.json",
			errResult: "This is an invalid JSON\nThis is an array that is surrounded by invalid \"][}{\"\n"},
		{filePath: "tests/step2/invalid22.json",
			errResult: "This is an invalid JSON\nThis is an array that is surrounded by invalid \"][}{\"\n"},
		{filePath: "tests/step2/invalid23.json",
			errResult: "This is an invalid JSON\nThis is an array that is surrounded by invalid \"][}{\"\n"},
		{filePath: "tests/step2/invalid24.json",
			errResult: "This is an invalid JSON\nThis is an array that is surrounded by invalid \"][}{\"\n"},
		{filePath: "tests/step2/invalid25.json",
			errResult: "This is an invalid JSON\nThis is an array that is surrounded by invalid \"][}{\"\n"},
		{filePath: "tests/step2/invalid26.json",
			errResult: "This is an invalid JSON\nThis is an array that is surrounded by invalid commas\n"},
		{filePath: "tests/step2/invalid27.json",
			errResult: "This is an invalid JSON\nThis is an array that is surrounded by invalid commas\n"},
		{filePath: "tests/step2/invalid28.json",
			errResult: "This is an invalid JSON\nThis is an array that is surrounded by invalid commas\n"},
		{filePath: "tests/step2/invalid29.json",
			errResult: "This is an invalid JSON\nThis is an array that is surrounded by invalid commas\n"},
		{filePath: "tests/step2/invalid30.json",
			errResult: "This is an invalid JSON\nThis is an array that is surrounded by invalid commas\n"},
		{filePath: "tests/step2/invalid31.json",
			errResult: "This is an invalid JSON\nThis is an array that is surrounded by invalid commas\n"},
		{filePath: "tests/step2/invalid32.json",
			errResult: "This is an invalid JSON\nThis is an array that is surrounded by invalid commas\n"},
		{filePath: "tests/step2/invalid33.json",
			errResult: "This is an invalid JSON\nThis is an array that is surrounded by invalid commas\n"},
		{filePath: "tests/step2/invalid34.json",
			errResult: "This is an invalid JSON\nThere is an unclosed object\n"},
		{filePath: "tests/step2/invalid35.json",
			errResult: "This is an invalid JSON\nThere is an object that is closed as an array\n"},
		{filePath: "tests/step2/fail2.json",
			errResult: "This is an invalid JSON\nThere is an unclosed array\n"},
		{filePath: "tests/step2/fail3.json",
			errResult: "This is an invalid JSON\nThere is a string that is not surrounded correctly by (\"\")\n"},
		{filePath: "tests/step2/fail4.json",
			errResult: "This is an invalid JSON\nThere is an array that contains an extra tail comma(s)\n"},
		{filePath: "tests/step2/fail5.json",
			errResult: "This is an invalid JSON\nThere is an array that contains an extra tail comma(s)\n"},
		{filePath: "tests/step2/fail6.json",
			errResult: "This is an invalid JSON\nThere is an array that contains an extra advancing comma(s)\n"},
		{filePath: "tests/step2/fail7.json",
			errResult: "This is an invalid JSON\nThis is an array that is surrounded by invalid commas\n"},
		{filePath: "tests/step2/fail8.json",
			errResult: "This is an invalid JSON\nThis is an array that is surrounded by invalid \"][}{\"\n"},
		{filePath: "tests/step2/fail10.json",
			errResult: "This is an invalid JSON\nMultiple values outside of an object or array\n"},
		{filePath: "tests/step2/fail12.json",
			errResult: "This is an invalid JSON\nThere is a string that is not surrounded correctly by (\"\")\n"},
		{filePath: "tests/step2/fail16.json",
			errResult: "This is an invalid JSON\nThere is a string that is not surrounded correctly by (\"\")\n"},
		{filePath: "tests/step2/fail23.json",
			errResult: "This is an invalid JSON\nThere is a string that is not surrounded correctly by (\"\")\n"},
		{filePath: "tests/step2/fail24.json",
			errResult: "This is an invalid JSON\nThere is a string that is not surrounded correctly by (\"\")\n"},
		{filePath: "tests/step2/fail33.json",
			errResult: "This is an invalid JSON\nThere is an array that is closed as an object\n"},
		{filePath: "tests/step3/fail9.json",
			errResult: "This is an invalid JSON\nThere is an object that contains an extra tail comma(s)\n"},
		{filePath: "tests/step3/fail13.json",
			errResult: "This is an invalid JSON\nThere is an invalid number, there is a leading zero\n"},
		{filePath: "tests/step3/fail14.json",
			errResult: "This is an invalid JSON\nThere is an invalid number, hexadecimal numbers are not allowed\n"},
		{filePath: "tests/step3/fail15.json",
			errResult: "This is an invalid JSON\nThere is a string that contains tabs or new lines or Illegal backslash escapes\n"},
		{filePath: "tests/step3/fail17.json",
			errResult: "This is an invalid JSON\nThere is a string that contains tabs or new lines or Illegal backslash escapes\n"},
		{filePath: "tests/step3/fail19.json",
			errResult: "This is an invalid JSON\nThere is an object that has a missing (:)\n"},
		{filePath: "tests/step3/fail20.json",
			errResult: "This is an invalid JSON\nThere is an object that has an invalid (:)s\n"},
		{filePath: "tests/step3/fail21.json",
			errResult: "This is an invalid JSON\nThere is an object that has a comma instead of a colon\n"},
		{filePath: "tests/step3/fail22.json",
			errResult: "This is an invalid JSON\nThere is an array that has a (:) instead of a (,)\n"},
		{filePath: "tests/step3/fail25.json",
			errResult: "This is an invalid JSON\nThere is a string that contains tabs or new lines or Illegal backslash escapes\n"},
		{filePath: "tests/step3/fail26.json",
			errResult: "This is an invalid JSON\nThere is a string that contains tabs or new lines or Illegal backslash escapes\n"},
		{filePath: "tests/step3/fail27.json",
			errResult: "This is an invalid JSON\nThere is a string that contains tabs or new lines or Illegal backslash escapes\n"},
		{filePath: "tests/step3/fail28.json",
			errResult: "This is an invalid JSON\nThere is a string that contains tabs or new lines or Illegal backslash escapes\n"},
		{filePath: "tests/step3/fail29.json",
			errResult: "This is an invalid JSON\nThere is an invalid number, there is a leading zero\n"},
		{filePath: "tests/step3/fail30.json",
			errResult: "This is an invalid JSON\nThere is an invalid number, there is a leading zero\n"},
		{filePath: "tests/step3/fail31.json",
			errResult: "This is an invalid JSON\nThere is an invalid number, there is a leading zero\n"},
		{filePath: "tests/step3/fail32.json",
			errResult: "This is an invalid JSON\nThere is an object that is closed with a comma(s)\n"},
		{filePath: "tests/step3/invalid.json",
			errResult: "This is an invalid JSON\n"},
		{filePath: "tests/step3/invalid2.json",
			errResult: "This is an invalid JSON\nThere is an invalid number, there is a leading zero\n"},
		{filePath: "tests/step3/invalid3.json",
			errResult: "This is an invalid JSON\nThere is an invalid number, there is a leading zero\n"},
		{filePath: "tests/step3/invalid4.json",
			errResult: "This is an invalid JSON\nThere is an invalid number, there is a leading zero\n"},
		{filePath: "tests/step3/invalid5.json",
			errResult: "This is an invalid JSON\nThere is an invalid number, there is a leading zero\n"},
		{filePath: "tests/step3/invalid6.json",
			errResult: "This is an invalid JSON\nThere is an invalid number, there is a leading zero\n"},
		{filePath: "tests/step3/invalid7.json",
			errResult: "This is an invalid JSON\nThere is an invalid number, there is a leading zero\n"},
		{filePath: "tests/step3/invalid8.json",
			errResult: "This is an invalid JSON\nThere is an invalid number, there is a leading zero\n"},
		{filePath: "tests/step3/invalid9.json",
			errResult: "This is an invalid JSON\nThere is an invalid number, there is a leading zero\n"},
		{filePath: "tests/step3/invalid10.json",
			errResult: "This is an invalid JSON\nThere is an invalid number, there is a leading +\n"},
		{filePath: "tests/step3/invalid11.json",
			errResult: "This is an invalid JSON\nShould be \"null\"\n"},
		{filePath: "tests/step3/invalid12.json",
			errResult: "This is an invalid JSON\nShould be \"false\"\n"},
		{filePath: "tests/step3/invalid13.json",
			errResult: "This is an invalid JSON\nShould be \"true\"\n"},
		{filePath: "tests/step3/invalid14.json",
			errResult: "This is an invalid JSON\nThere is an invalid number, there is a leading zero\n"},
		{filePath: "tests/step3/invalid15.json",
			errResult: "This is an invalid JSON\nThere is an object that contains an extra tail comma(s)\n"},
		{filePath: "tests/step3/invalid16.json",
			errResult: "This is an invalid JSON\nThere is an array that contains an extra tail comma(s)\n"},
		{filePath: "tests/step3/invalid17.json",
			errResult: "This is an invalid JSON\nThere is a string that is not surrounded correctly by (\"\")\n"},
		{filePath: "tests/step3/invalid18.json",
			errResult: "This is an invalid JSON\nThere is a string that is not surrounded correctly by (\"\")\n"},
		{filePath: "tests/step3/invalid19.json",
			errResult: "This is an invalid JSON\n"},
		{filePath: "tests/step3/invalid20.json",
			errResult: "This is an invalid JSON\nThere is a string that is not surrounded correctly by (\"\")\n"},
		{filePath: "tests/step3/invalid21.json",
			errResult: "This is an invalid JSON\nThere is a string that is not surrounded correctly by (\"\")\n"},
		{filePath: "tests/step3/invalid22.json",
			errResult: "This is an invalid JSON\nThere is an invalid number, hexadecimal numbers are not allowed\n"},
		{filePath: "tests/step3/invalid23.json",
			errResult: "This is an invalid JSON\nThere is an invalid number, there is a leading +\n"},
		{filePath: "tests/step3/invalid24.json",
			errResult: "This is an invalid JSON\n"},
		{filePath: "tests/step3/invalid25.json",
			errResult: "This is an invalid JSON\nThere is an object that has a missing (:)\n"},
		{filePath: "tests/step3/invalid26.json",
			errResult: "This is an invalid JSON\nThere is a string that is not surrounded correctly by (\"\")\n"},
		{filePath: "tests/step3/invalid27.json",
			errResult: "This is an invalid JSON\nThere is an array that is closed with a comma(s)\n"},
		{filePath: "tests/step3/invalid28.json",
			errResult: "This is an invalid JSON\nThere is an object that contains an extra advancing comma(s)\n"},
		{filePath: "tests/step4/invalid.json",
			errResult: "This is an invalid JSON\nThere is a string that is not surrounded correctly by (\"\")\n"},
		{filePath: "tests/step4/invalid2.json",
			errResult: "This is an invalid JSON\nThere is a string that is not surrounded correctly by (\"\")\n"},
		{filePath: "tests/step4/invalid3.json",
			errResult: "This is an invalid JSON\n"},
		{filePath: "tests/step4/invalid4.json",
			errResult: "This is an invalid JSON\n"},
		{filePath: "tests/step4/invalid5.json",
			errResult: "This is an invalid JSON\n"},
		{filePath: "tests/step4/invalid6.json",
			errResult: "This is an invalid JSON\nThere is an invalid number, there is a leading zero\n"},
		{filePath: "tests/step4/invalid7.json",
			errResult: "This is an invalid JSON\nThere is an invalid number, there is a leading zero\n"},
		{filePath: "tests/step4/invalid8.json",
			errResult: "This is an invalid JSON\nThere is a string that is not surrounded correctly by (\"\")\n"},
		{filePath: "tests/step4/invalid9.json",
			errResult: "This is an invalid JSON\nThere is a string that is not surrounded correctly by (\"\")\n"},
		{filePath: "tests/step4/invalid10.json",
			errResult: "This is an invalid JSON\nThere is a string that is not surrounded correctly by (\"\")\n"},
		{filePath: "tests/step4/invalid11.json",
			errResult: "This is an invalid JSON\nThere is a string that is not surrounded correctly by (\"\")\n"},
		{filePath: "tests/step4/invalid12.json",
			errResult: "This is an invalid JSON\nThis is an array that is surrounded by invalid \"][}{\"\n"},
		{filePath: "tests/step4/invalid13.json",
			errResult: "This is an invalid JSON\nThis is an array that is surrounded by invalid \"][}{\"\n"},
		{filePath: "tests/step4/invalid14.json",
			errResult: "This is an invalid JSON\n"},
		{filePath: "tests/step4/invalid15.json",
			errResult: "This is an invalid JSON\nMultiple values outside of an object or array\n"},
		{filePath: "tests/step4/invalid16.json",
			errResult: "This is an invalid JSON\nThere is an object that has a missing (:)\n"},
		{filePath: "tests/step4/invalid17.json",
			errResult: "This is an invalid JSON\nThere is an object that has an invalid (:)s\n"},
		{filePath: "tests/step4/invalid18.json",
			errResult: "This is an invalid JSON\nThere is an object that has an invalid (:)s\n"},
		{filePath: "tests/huge_json/invalid.json",
			errResult: "This is an invalid JSON\nThere is an object that has a missing (:)\n"},
		{filePath: "tests/huge_json/invalid2.json",
			errResult: "This is an invalid JSON\nThere is a string that is not surrounded correctly by (\"\")\n"},
	}

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

func TestFileReadingFailure(t *testing.T) {
	var out strings.Builder
	var errOut strings.Builder
	ccjsonparserCommand := exec.Command("./ccjsonparser", "tests/step4/invalid115.json")
	ccjsonparserCommand.Dir = "./.."
	out.Reset()
	errOut.Reset()
	ccjsonparserCommand.Stderr = &errOut
	ccjsonparserCommand.Stdout = &out
	err := ccjsonparserCommand.Run()
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
