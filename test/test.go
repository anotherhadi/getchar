package main

import (
	"fmt"
	"github.com/anotherhadi/getchar"
)

func main() {
	fmt.Print("Type any key:")
	ascii, arrow, err := getchar.GetChar()

	fmt.Println("\n", ascii, arrow, err)
}
