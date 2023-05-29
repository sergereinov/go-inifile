# go-inifile
Wrapper around `gopkg.in/ini.v1` to make it look more like `flag.String()` etc.

## Examples

1. *examples/simple/main.go*

```go
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
```
Prints:
```
ini file is in D:\Progs\Go\go-inifile\simple.ini
dir ./logs
maxFileSize 10
enable true
tags [one five ten]
intervals [1 5 10]
```

2. *examples/named/main.go*

```go
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
```
It prints:
```
ini file is in D:\Progs\Go\go-inifile\custom.ini
dir ./logs
tags [one five ten]
```
And saves *custom.ini*:
```ini
[Logs]
Dir = ./logs

[Summary]
Tags = one, five, ten
```
