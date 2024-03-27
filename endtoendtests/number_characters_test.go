package endtoendtests

import (
	"os/exec"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGettingTheNumberOfCharactersInTheFile(t *testing.T) {
	ccwcCommand := exec.Command("./ccwc", "-m", "test.txt")
	ccwcCommand.Dir = "./.."
	var out strings.Builder
	ccwcCommand.Stdout = &out
	err := ccwcCommand.Run()
	assert.NoError(t, err)
	assert.Equal(t, "339292 test.txt\n", out.String())
}
