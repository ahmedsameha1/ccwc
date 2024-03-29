package app

import (
	"errors"
	"fmt"
	"strings"
	"unicode/utf8"
)

func App(readFile func(name string) ([]byte, error), args []string) (string, error) {
	err := validate(readFile, args)
	if err != nil {
		return "", err
	}
	contentInBytes, _ := readFile(args[2])
	contentString := string(contentInBytes)
	if args[1] == "-l" {
		count := strings.Count(contentString, "\n")
		if !strings.HasSuffix(contentString, "\n") {
			return fmt.Sprintf("%d %s", count+1, args[2]), nil
		} else {
			return fmt.Sprintf("%d %s", count, args[2]), nil
		}
	} else if args[1] == "-w" {
		words := strings.Fields(contentString)
		return fmt.Sprintf("%d %s", len(words), args[2]), nil
	} else if args[1] == "-m" {
		return fmt.Sprintf("%d %s", utf8.RuneCountInString(contentString), args[2]), nil
	}
	return fmt.Sprintf("%d %s", len(contentInBytes), args[2]), nil
}

func validate(readFile func(name string) ([]byte, error), args []string) error {
	lengthOfArgs := len(args)
	if lengthOfArgs > 2 {
		if strings.HasPrefix(args[1], "-") {
			return checkFilesExistance(readFile, args[2:])
		} else {
			return checkFilesExistance(readFile, args[1:])
		}
	} else if lengthOfArgs == 2 {
		if strings.HasPrefix(args[1], "-") {
			return errors.New("there is an error with your options/arguments")
		} else {
			_, err := readFile(args[1])
			if err != nil {
				if strings.Contains(err.Error(), "no such file") {
					return errors.New("there is no such file")
				}
			}
		}
	} else if lengthOfArgs == 1 {
		return errors.New("where is the file name?")
	}
	return nil
}

func checkFilesExistance(readFile func(name string) ([]byte, error), args []string) error {
	var errorMessage string
	for i := 0; i < len(args); i++ {
		_, err := readFile(args[i])
		if err != nil {
			if i == (len(args) - 1) {
				if strings.Contains(err.Error(), "no such file") {
					errorMessage = errorMessage + "there is no such file: " + args[i]
				}
			} else {
				if strings.Contains(err.Error(), "no such file") {
					errorMessage = errorMessage + "there is no such file: " + args[i] + "\n"
				}
			}
		}
	}
	if len(errorMessage) > 0 {
		return errors.New(errorMessage)
	}
	return nil
}
