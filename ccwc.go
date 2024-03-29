package main

import (
	"fmt"
	"os"

	"github.com/ahmedsameha1/ccwc/app"
)

func main() {
	result, err := app.App(os.ReadFile, os.Args)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	} else {
		fmt.Println(result)
	}
}
