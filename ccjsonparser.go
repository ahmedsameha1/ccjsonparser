package main

import (
	"fmt"
	"os"

	"github.com/ahmedsameha1/ccjsonparser/app"
)

func main() {
	result, _ := app.App(os.ReadFile, os.Args)
	fmt.Println(result)
}
