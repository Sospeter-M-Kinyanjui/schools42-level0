package revAlpha

import "fmt"

func PrintAlphaReverse() {
	for i:=25; i >=0; i-- {
		if i % 2 != 0 {
			fmt.Print(string(i + 65))
		} else {
			fmt.Print(string(i + 97))
		}
	}
	fmt.Println()
}