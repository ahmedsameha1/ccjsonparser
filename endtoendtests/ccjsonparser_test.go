package endtoendtests

import (
	"os"
	"os/exec"
	"testing"
)

func TestMain(m *testing.M) {
	buildCommand := exec.Command("go", "build", ".")
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
