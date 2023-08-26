<p align="center">
</p>

# GetChar

The `GetChar` function is a utility to read a single character from the standard input (stdin) without requiring the user to press the Enter key. It's particularly useful when you want to capture a single character input quickly, such as in interactive command-line applications or games.

---

## Usage

### Installation:
```bash
go get -u github.com/anotherhadi/getchar@latest
```

### Importing:
```go
import (
    "github.com/anotherhadi/getchar"
)
```

### Basic Usage:
```go
func main() {
  fmt.Print("Type any key:")
	ascii, arrow, err := getchar.GetChar()

	fmt.Println("\n", ascii, arrow, err)  
}
```
Example available in "test/test.go"

---

<img src="https://img.buymeacoffee.com/button-api/?text=Buy me a cookie&emoji=🍪&slug=anotherhadi&button_colour=eed2cc&font_colour=000000&font_family=Inter&outline_colour=ffffff&coffee_colour=ff0000" />

