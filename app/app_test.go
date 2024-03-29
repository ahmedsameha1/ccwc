package app

import (
	"errors"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestApp(t *testing.T) {
	result, err := App(func(name string) ([]byte, error) {
		if name != "test.txt" {
			panic("error")
		}
		return []byte{4, 127, 128, 129, 255}, nil
	}, []string{"ccwc", "-c", "test.txt"})
	assert.NoError(t, err)
	assert.Equal(t, "5 test.txt", result)

	result, err = App(func(name string) ([]byte, error) {
		if name != "test.txt" {
			panic("error")
		}
		return []byte{4, 127, 128, 129, 255}, nil
	}, []string{"ccwc", "-l", "test.txt"})
	assert.NoError(t, err)
	assert.Equal(t, "0 test.txt", result)

	result, err = App(func(name string) ([]byte, error) {
		if name != "test.txt" {
			panic("error")
		}
		return []byte{4, 127, 128, 129, 255, 10, 66, 67, 68, 10}, nil
	}, []string{"ccwc", "-l", "test.txt"})
	assert.NoError(t, err)
	assert.Equal(t, "2 test.txt", result)

	result, err = App(func(name string) ([]byte, error) {
		if name != "test.txt" {
			panic("error")
		}
		return []byte{4, 127, 128, 129, 255, 10, 66, 67, 68}, nil
	}, []string{"ccwc", "-l", "test.txt"})
	assert.NoError(t, err)
	assert.Equal(t, "1 test.txt", result)

	result, err = App(func(name string) ([]byte, error) {
		if name != "test.txt" {
			panic("error")
		}
		return []byte("BCDEF\nBCD"), nil
	}, []string{"ccwc", "-w", "test.txt"})
	assert.NoError(t, err)
	assert.Equal(t, "2 test.txt", result)

	result, err = App(func(name string) ([]byte, error) {
		if name != "test.txt" {
			panic("error")
		}
		return []byte("BCDEF BCD"), nil
	}, []string{"ccwc", "-w", "test.txt"})
	assert.NoError(t, err)
	assert.Equal(t, "2 test.txt", result)

	result, err = App(func(name string) ([]byte, error) {
		if name != "test.txt" {
			panic("error")
		}
		return []byte("BCDEF\rBCD"), nil
	}, []string{"ccwc", "-w", "test.txt"})
	assert.NoError(t, err)
	assert.Equal(t, "2 test.txt", result)

	result, err = App(func(name string) ([]byte, error) {
		if name != "test.txt" {
			panic("error")
		}
		return []byte("BCDEF BCD "), nil
	}, []string{"ccwc", "-w", "test.txt"})
	assert.NoError(t, err)
	assert.Equal(t, "2 test.txt", result)

	result, err = App(func(name string) ([]byte, error) {
		if name != "test.txt" {
			panic("error")
		}
		return []byte(" BCDEF BCD "), nil
	}, []string{"ccwc", "-w", "test.txt"})
	assert.NoError(t, err)
	assert.Equal(t, "2 test.txt", result)

	result, err = App(func(name string) ([]byte, error) {
		if name != "test.txt" {
			panic("error")
		}
		return []byte(" BCDEF BCD"), nil
	}, []string{"ccwc", "-w", "test.txt"})
	assert.NoError(t, err)
	assert.Equal(t, "2 test.txt", result)

	result, err = App(func(name string) ([]byte, error) {
		if name != "test.txt" {
			panic("error")
		}
		return []byte("BCDEF BCD BCDE"), nil
	}, []string{"ccwc", "-w", "test.txt"})
	assert.NoError(t, err)
	assert.Equal(t, "3 test.txt", result)

	result, err = App(func(name string) ([]byte, error) {
		if name != "test.txt" {
			panic("error")
		}
		return []byte("BCDEF BCD BCDE"), nil
	}, []string{"ccwc", "-m", "test.txt"})
	assert.NoError(t, err)
	assert.Equal(t, "14 test.txt", result)

	result, err = App(func(name string) ([]byte, error) {
		if name != "test.txt" {
			panic("error")
		}
		return []byte("BCDEF BCD BCDE"), nil
	}, []string{"ccwc", "test.txt"})
	assert.NoError(t, err)
	assert.Equal(t, "     0      3     14 test.txt", result)

	result, err = App(func(name string) ([]byte, error) {
		if name != "test.txt" && name != "test2.txt" && name != "test3.txt" {
			panic("error")
		}
		if name == "test.txt" {
			return []byte("BCDEF BCD BCDE"), nil
		} else if name == "test2.txt" {
			return []byte("BCDEF BCD BCDE\nBCD"), nil
		} else {
			return []byte("BCDEF BCD BCDE\n" + strings.Repeat("n", 1000)), nil
		}
	}, []string{"ccwc", "test.txt", "test2.txt", "test3.txt"})
	assert.NoError(t, err)
	assert.Equal(t, "     0      3     14 test.txt\n     1      4     18 test2.txt\n     1      4   1015 test3.txt", result)
}

func TestAppValidation(t *testing.T) {
	result, err := App(func(name string) ([]byte, error) {
		if name != "test.txt" {
			panic("error")
		}
		return []byte("BCDEF BCD BCDE"), nil
	}, []string{"ccwc", "-w"})
	assert.Empty(t, result)
	assert.Equal(t, "there is an error with your options/arguments", err.Error())

	result, err = App(func(name string) ([]byte, error) {
		if name != "b.txt" {
			panic("error")
		}
		return nil, errors.New("There is no such file")
	}, []string{"ccwc", "b.txt"})
	assert.Empty(t, result)
	assert.Equal(t, "there is no such file", err.Error())

	result, err = App(func(name string) ([]byte, error) {
		if name != "b.txt" {
			panic("error")
		}
		return nil, errors.New("There is no such file")
	}, []string{"ccwc", "-w", "b.txt"})
	assert.Empty(t, result)
	assert.Equal(t, "there is no such file: b.txt", err.Error())

	result, err = App(func(name string) ([]byte, error) {
		return nil, nil
	}, []string{"ccwc"})
	assert.Empty(t, result)
	assert.Equal(t, "where is the file name?", err.Error())

	result, err = App(func(name string) ([]byte, error) {
		if name != "b.txt" && name != "-w" {
			panic("error")
		}
		return nil, errors.New("There is no such file")
	}, []string{"ccwc", "b.txt", "-w"})
	assert.Empty(t, result)
	assert.Equal(t, "there is no such file: b.txt\nthere is no such file: -w", err.Error())

	result, err = App(func(name string) ([]byte, error) {
		if name != "b.txt" && name != "c.txt" {
			panic("error")
		}
		return nil, errors.New("There is no such file")
	}, []string{"ccwc", "-w", "b.txt", "c.txt"})
	assert.Empty(t, result)
	assert.Equal(t, "there is no such file: b.txt\nthere is no such file: c.txt", err.Error())
}
