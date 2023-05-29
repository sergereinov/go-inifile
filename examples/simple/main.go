package main

import (
	"fmt"

	"github.com/sergereinov/go-inifile"
)

var (
	file        = inifile.New()
	dir         = file.String("Logs", "Dir", "./logs")
	maxFileSize = file.Int("Logs", "MaxFileSizeMB", 10)
	enable      = file.Bool("Logs", "Enable", true)
	tags        = file.Strings("Summary", "Tags", []string{"one", "five", "ten"})
	intervals   = file.Ints("Summary", "Intervals", []int{1, 5, 10})
)

func main() {
	fmt.Println("ini file is in", file.Path())
	fmt.Println("dir", dir)
	fmt.Println("maxFileSize", maxFileSize)
	fmt.Println("enable", enable)
	fmt.Println("tags", tags)
	fmt.Println("intervals", intervals)
}
