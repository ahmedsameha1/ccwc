package endtoendtests

import (
	"os"
	"os/exec"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	buildCommand := exec.Command("go", "build", ".")
	buildCommand.Dir = "./.."
	err := buildCommand.Run()
	if err != nil {
		panic(err)
	}

	result := m.Run()

	removeCommand := exec.Command("rm", "ccwc")
	removeCommand.Dir = "./.."
	err = removeCommand.Run()
	if err != nil {
		panic(err)
	}
	os.Exit(result)
}

func TestCcwc(t *testing.T) {
	ccwcCommand := exec.Command("./ccwc", "-c", "test.txt")
	ccwcCommand.Dir = "./.."
	var out strings.Builder
	ccwcCommand.Stdout = &out
	err := ccwcCommand.Run()
	assert.NoError(t, err)
	assert.Equal(t, "342190 test.txt", out.String())
}
