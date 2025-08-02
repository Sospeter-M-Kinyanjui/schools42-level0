package aff_a

import "fmt"

func DisplayFirstA(str string) {
	for _, char := range str {
		if char == 'a' {
			fmt.Println(char)
			return
		}
	}
	fmt.Println()
}