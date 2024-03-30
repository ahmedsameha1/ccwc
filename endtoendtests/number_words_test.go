package endtoendtests

import (
	"os/exec"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGettingTheNumberOfWordsInTheFile(t *testing.T) {
	ccwcCommand := exec.Command("./ccwc", "-w", "test.txt")
	ccwcCommand.Dir = "./.."
	var out strings.Builder
	ccwcCommand.Stdout = &out
	err := ccwcCommand.Run()
	assert.NoError(t, err)
	assert.Equal(t, "58164 test.txt\n", out.String())

	ccwcCommand = exec.Command("./ccwc", "-w", "test.txt", "test2.txt", "test3.txt")
	ccwcCommand.Dir = "./.."
	out.Reset()
	ccwcCommand.Stdout = &out
	err = ccwcCommand.Run()
	assert.NoError(t, err)
	assert.Equal(t, " 58164 test.txt\n    32 test2.txt\n   152 test3.txt\n", out.String())
}
