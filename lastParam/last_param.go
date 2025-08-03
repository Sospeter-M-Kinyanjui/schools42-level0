package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args 
	if len(args) == 1 {
		fmt.Println()
	} else {
		fmt.Println(args[len(args)-1])
	}
}