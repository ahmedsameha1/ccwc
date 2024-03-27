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
	}, []string{"ccwc", "-c", "test.txt"})
	assert.Equal(t, "5 test.txt\n", result)

	result = App(func(name string) ([]byte, error) {
		if name != "test.txt" {
			panic("error")
		}
		return []byte{4, 127, 128, 129, 255}, nil
	}, []string{"ccwc", "-l", "test.txt"})
	assert.Equal(t, "1 test.txt\n", result)

	result = App(func(name string) ([]byte, error) {
		if name != "test.txt" {
			panic("error")
		}
		return []byte{4, 127, 128, 129, 255, 10, 66, 67, 68, 10}, nil
	}, []string{"ccwc", "-l", "test.txt"})
	assert.Equal(t, "2 test.txt\n", result)

	result = App(func(name string) ([]byte, error) {
		if name != "test.txt" {
			panic("error")
		}
		return []byte{4, 127, 128, 129, 255, 10, 66, 67, 68}, nil
	}, []string{"ccwc", "-l", "test.txt"})
	assert.Equal(t, "2 test.txt\n", result)

	result = App(func(name string) ([]byte, error) {
		if name != "test.txt" {
			panic("error")
		}
		return []byte("BCDEF\nBCD"), nil
	}, []string{"ccwc", "-w", "test.txt"})
	assert.Equal(t, "2 test.txt\n", result)

	result = App(func(name string) ([]byte, error) {
		if name != "test.txt" {
			panic("error")
		}
		return []byte("BCDEF BCD"), nil
	}, []string{"ccwc", "-w", "test.txt"})
	assert.Equal(t, "2 test.txt\n", result)

	result = App(func(name string) ([]byte, error) {
		if name != "test.txt" {
			panic("error")
		}
		return []byte("BCDEF\rBCD"), nil
	}, []string{"ccwc", "-w", "test.txt"})
	assert.Equal(t, "2 test.txt\n", result)

	result = App(func(name string) ([]byte, error) {
		if name != "test.txt" {
			panic("error")
		}
		return []byte("BCDEF BCD "), nil
	}, []string{"ccwc", "-w", "test.txt"})
	assert.Equal(t, "2 test.txt\n", result)

	result = App(func(name string) ([]byte, error) {
		if name != "test.txt" {
			panic("error")
		}
		return []byte(" BCDEF BCD "), nil
	}, []string{"ccwc", "-w", "test.txt"})
	assert.Equal(t, "2 test.txt\n", result)

	result = App(func(name string) ([]byte, error) {
		if name != "test.txt" {
			panic("error")
		}
		return []byte(" BCDEF BCD"), nil
	}, []string{"ccwc", "-w", "test.txt"})
	assert.Equal(t, "2 test.txt\n", result)

	result = App(func(name string) ([]byte, error) {
		if name != "test.txt" {
			panic("error")
		}
		return []byte("BCDEF BCD BCDE"), nil
	}, []string{"ccwc", "-w", "test.txt"})
	assert.Equal(t, "3 test.txt\n", result)

	result = App(func(name string) ([]byte, error) {
		if name != "test.txt" {
			panic("error")
		}
		return []byte("BCDEF BCD BCDE"), nil
	}, []string{"ccwc", "-m", "test.txt"})
	assert.Equal(t, "14 test.txt\n", result)
}
