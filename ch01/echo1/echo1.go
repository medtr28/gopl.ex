// Echo1 : print command line arg
package main

import (
	"fmt"
	"os"
)

func main() {
	var s, seq string
	for i := 1; i < len(os.Args); i++ {
		s += seq + os.Args[i]
		seq = " "
	}
	fmt.Println(s)
}
