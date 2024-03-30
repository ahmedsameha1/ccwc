package endtoendtests

import (
	"os/exec"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGettingTheNumberOfBytesInTheFile(t *testing.T) {
	ccwcCommand := exec.Command("./ccwc", "-c", "test.txt")
	ccwcCommand.Dir = "./.."
	var out strings.Builder
	ccwcCommand.Stdout = &out
	err := ccwcCommand.Run()
	assert.NoError(t, err)
	assert.Equal(t, "342190 test.txt\n", out.String())

	ccwcCommand = exec.Command("./ccwc", "-c", "test.txt", "test2.txt", "test3.txt")
	ccwcCommand.Dir = "./.."
	out.Reset()
	ccwcCommand.Stdout = &out
	err = ccwcCommand.Run()
	assert.NoError(t, err)
	assert.Equal(t, "342190 test.txt\n   165 test2.txt\n   911 test3.txt\n", out.String())
}
