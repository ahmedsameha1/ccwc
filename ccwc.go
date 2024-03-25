package main

import (
	"fmt"
	"os"

	"github.com/ahmedsameha1/ccwc/app"
)

func main() {
	fmt.Print(app.App(os.ReadFile, os.Args))
}
