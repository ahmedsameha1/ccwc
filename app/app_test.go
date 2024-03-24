package app

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestApp(t *testing.T) {
	result := App(func(name string) ([]byte, error) {
		if name != "test.txt" {
			panic("error")
		}
		return []byte{4, 127, 128, 129, 255}, nil
	}, "test.txt")
	assert.Equal(t, 5, result)
}
