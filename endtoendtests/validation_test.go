package endtoendtests

import (
	"os/exec"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidation(t *testing.T) {
	ccwcCommand := exec.Command("./ccwc", "-w")
	ccwcCommand.Dir = "./.."
	var errOut strings.Builder
	ccwcCommand.Stderr = &errOut
	err := ccwcCommand.Run()
	assert.Error(t, err)
	assert.Equal(t, "there is an error with your options/arguments\n", errOut.String())

	ccwcCommand = exec.Command("./ccwc", "b.txt")
	ccwcCommand.Dir = "./.."
	errOut.Reset()
	ccwcCommand.Stderr = &errOut
	err = ccwcCommand.Run()
	assert.Error(t, err)
	assert.Equal(t, "there is no such file\n", errOut.String())

	ccwcCommand = exec.Command("./ccwc", "-w", "b.txt")
	ccwcCommand.Dir = "./.."
	errOut.Reset()
	ccwcCommand.Stderr = &errOut
	err = ccwcCommand.Run()
	assert.Error(t, err)
	assert.Equal(t, "there is no such file: b.txt\n", errOut.String())

	ccwcCommand = exec.Command("./ccwc")
	ccwcCommand.Dir = "./.."
	errOut.Reset()
	ccwcCommand.Stderr = &errOut
	err = ccwcCommand.Run()
	assert.Error(t, err)
	assert.Equal(t, "where is the file name?\n", errOut.String())

	ccwcCommand = exec.Command("./ccwc", "b.txt", "-w")
	ccwcCommand.Dir = "./.."
	errOut.Reset()
	ccwcCommand.Stderr = &errOut
	err = ccwcCommand.Run()
	assert.Error(t, err)
	assert.Equal(t, "there is no such file: b.txt\nthere is no such file: -w\n", errOut.String())

	ccwcCommand = exec.Command("./ccwc", "-w", "b.txt", "c.txt")
	ccwcCommand.Dir = "./.."
	errOut.Reset()
	ccwcCommand.Stderr = &errOut
	err = ccwcCommand.Run()
	assert.Error(t, err)
	assert.Equal(t, "there is no such file: b.txt\nthere is no such file: c.txt\n", errOut.String())
}
