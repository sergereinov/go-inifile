package main

import (
	"fmt"

	"github.com/sergereinov/go-inifile"
)

func main() {
	// Create ini file accessor
	file := inifile.New("./custom.ini")

	// Read some values
	dir := file.String("Logs", "Dir", "./logs")
	tags := file.Strings("Summary", "Tags", []string{"one", "five", "ten"})

	// Save values back with defaults for missing keys
	file.Save()

	// Print results
	fmt.Println("ini file is in", file.Path())
	fmt.Println("dir", dir)
	fmt.Println("tags", tags)
}
