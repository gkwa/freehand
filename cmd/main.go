package main

import (
	"os"

	"github.com/taylormonacelli/freehand"
)

func main() {
	code := freehand.Execute()
	os.Exit(code)
}
