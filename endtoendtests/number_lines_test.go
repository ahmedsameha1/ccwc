package endtoendtests

import (
	"os/exec"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGettingTheNumberOfLinesInTheFile(t *testing.T) {
	ccwcCommand := exec.Command("./ccwc", "-l", "test.txt")
	ccwcCommand.Dir = "./.."
	var out strings.Builder
	ccwcCommand.Stdout = &out
	err := ccwcCommand.Run()
	assert.NoError(t, err)
	assert.Equal(t, "7145 test.txt\n", out.String())
}
