package main

import (
	"fmt"
	"os"
	"schools42-level0/firstZ"
)

func main() {
	args := os.Args
	if len(args) == 2 {
		firstZ.SayIt(args[1])
	} else {
		fmt.Println("z")
	}
}
