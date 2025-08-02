package main

import (
	"fmt"
	"os"
	"schools42-level0/aff_a"
)

func main() {
	args := os.Args
	if len(args) == 2 {

		sampleString := args[1]
		aff_a.DisplayFirstA(sampleString)
	} else {
		fmt.Println("a")
	}
}
