package justSayZ

import (
	"fmt"
	"os"
)

func SayIt() {
	fmt.Fprint(os.Stdout, string('z'))
}