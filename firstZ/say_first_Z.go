package firstZ

import "os"

func SayIt(str string) {
	for i := 0; i < len(str); i++ {
		if str[i] == 'z' {
			os.Stdout.Write([]byte{'z','\n'})
			return
		}
	}
	os.Stdout.Write([]byte{'z', '\n'})
}
