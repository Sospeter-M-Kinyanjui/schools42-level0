package alpha

import "fmt"

func PrintAlphaLetters() {
	for i:=0; i<26; i++ {
		if i % 2 != 0 {
			asciiPos := 65 + i 
			fmt.Print(string(asciiPos))
		} else {
			asciiPos := 97 + i 
			fmt.Print(string(asciiPos))
		}
	}
	fmt.Println()
}