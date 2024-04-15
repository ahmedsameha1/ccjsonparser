package endtoendtests

import (
	"os/exec"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseTheSimplestJsonObject(t *testing.T) {
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
	assert.Equal(t, "This is an invalid JSON\n", out.String())

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
}
