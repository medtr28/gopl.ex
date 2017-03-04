// echo : print command line Arg with times
package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	masrure(echo1)
	masrure(echo2)
}

func echo1() {
	s, seq := "", ""
	for _, arg := range os.Args[1:] {
		s += seq + arg
		seq = " "
	}
	fmt.Println(s)
}

func echo2() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}

func masrure(f func()) {
	start := time.Now()
	f()
	secs := time.Since(start).Seconds()
	fmt.Println(secs) // TODO : print func name
}
