// echo2 : print command line Arg #2
package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("index : arg")
	for i, arg := range os.Args[1:] {
		i_s := fmt.Sprintf("%d", i)
		fmt.Println(i_s, arg)
	}
}
