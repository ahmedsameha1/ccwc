package app

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func App(readFile func(name string) ([]byte, error), args []string) string {
	contentInBytes, _ := readFile(args[2])
	contentString := string(contentInBytes)
	if args[1] == "-l" {
		count := strings.Count(contentString, "\n")
		if !strings.HasSuffix(contentString, "\n") {
			return fmt.Sprintf("%d %s\n", count+1, args[2])
		} else {
			return fmt.Sprintf("%d %s\n", count, args[2])
		}
	} else if args[1] == "-w" {
		words := strings.Fields(contentString)
		return fmt.Sprintf("%d %s\n", len(words), args[2])
	} else if args[1] == "-m" {
		return fmt.Sprintf("%d %s\n", utf8.RuneCountInString(contentString), args[2])
	}
	return fmt.Sprintf("%d %s\n", len(contentInBytes), args[2])
}
