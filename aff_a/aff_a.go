package aff_a

import "fmt"

func DisplayFirstA(str string) {
	for index, char := range str {
		if char == 'a' {
			fmt.Println("found", char, "at index", index)
			return
		}
	}
	fmt.Println()
}