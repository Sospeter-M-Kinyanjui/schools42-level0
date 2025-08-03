package justSayA

import "os"

func SayIt() {
	os.Stdout.Write([]byte{'a'})
}