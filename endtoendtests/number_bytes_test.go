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
}
