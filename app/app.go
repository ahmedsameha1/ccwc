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
	lengthOfArgs := len(args)
	args1 := args[1]
	if strings.HasPrefix(args[1], "-") {
		if lengthOfArgs <= 3 { // less than 3 is handled be validate(), so it should never be less than 3
			contentInBytes, _ := readFile(args[2])
			contentString := string(contentInBytes)
			if args1 == "-l" {
				return fmt.Sprintf("%d %s", linesOption(contentString), args[2]), nil
			} else if args1 == "-w" {
				return fmt.Sprintf("%d %s", wordsOption(contentString), args[2]), nil
			} else if args1 == "-m" {
				return fmt.Sprintf("%d %s", utf8.RuneCountInString(contentString), args[2]), nil
			} else {
				return fmt.Sprintf("%d %s", len(contentInBytes), args[2]), nil
			}
		} else {
			var result string
			fileNames := args[2:]
			if args1 == "-l" {
				for i := 0; i < len(fileNames); i++ {
					contentInBytes, _ := readFile(fileNames[i])
					contentString := string(contentInBytes)
					if i == len(fileNames)-1 {
						result = result + fmt.Sprintf("%6d %s", linesOption(contentString), fileNames[i])
					} else {
						result = result + fmt.Sprintf("%6d %s\n", linesOption(contentString), fileNames[i])
					}
				}
				return result, nil
			} else if args1 == "-w" {
				for i := 0; i < len(fileNames); i++ {
					contentInBytes, _ := readFile(fileNames[i])
					contentString := string(contentInBytes)
					if i == len(fileNames)-1 {
						result = result + fmt.Sprintf("%6d %s", wordsOption(contentString), fileNames[i])
					} else {
						result = result + fmt.Sprintf("%6d %s\n", wordsOption(contentString), fileNames[i])
					}
				}
				return result, nil
			} else if args1 == "-m" {
				for i := 0; i < len(fileNames); i++ {
					contentInBytes, _ := readFile(fileNames[i])
					contentString := string(contentInBytes)
					if i == len(fileNames)-1 {
						result = result + fmt.Sprintf("%6d %s", utf8.RuneCountInString(contentString), fileNames[i])
					} else {
						result = result + fmt.Sprintf("%6d %s\n", utf8.RuneCountInString(contentString), fileNames[i])
					}
				}
				return result, nil
			} else {
				for i := 0; i < len(fileNames); i++ {
					contentInBytes, _ := readFile(fileNames[i])
					if i == len(fileNames)-1 {
						result = result + fmt.Sprintf("%6d %s", len(contentInBytes), fileNames[i])
					} else {
						result = result + fmt.Sprintf("%6d %s\n", len(contentInBytes), fileNames[i])
					}
				}
				return result, nil
			}
		}
	} else {
		var result string
		fileNames := args[1:]
		for i := 0; i < len(fileNames); i++ {
			if i == len(fileNames)-1 {
				contentInBytes, _ := readFile(fileNames[i])
				contentString := string(contentInBytes)
				result = result + fmt.Sprintf("%6d %6d %6d %s", linesOption(contentString), wordsOption(contentString),
					len(contentInBytes), fileNames[i])
			} else {
				contentInBytes, _ := readFile(fileNames[i])
				contentString := string(contentInBytes)
				result = result + fmt.Sprintf("%6d %6d %6d %s\n", linesOption(contentString), wordsOption(contentString),
					len(contentInBytes), fileNames[i])
			}
		}
		return result, nil
	}
}

func validate(readFile func(name string) ([]byte, error), args []string) error {
	lengthOfArgs := len(args)
	if lengthOfArgs > 2 {
		if strings.HasPrefix(args[1], "-") {
			if args[1] != "-l" && args[1] != "-w" && args[1] != "-c" && args[1] != "-m" {
				return errors.New("there is an error with your options/arguments")
			}
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

func linesOption(contentString string) int {
	count := strings.Count(contentString, "\n")
	return count
}

func wordsOption(contentString string) int {
	return len(strings.Fields(contentString))
}
