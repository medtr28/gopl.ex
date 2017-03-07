// Dup print text and count
// scan from stdin or filename
package main

import (
	"bufio"
	"fmt"
	"os"
)

type Tally struct {
	Files map[string]int
	Total int
}

func main() {
	counts := make(map[string]*Tally)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, "none")
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts, arg)
			f.Close()
		}
	}
	for line, tally := range counts {
		if tally.Total > 1 {
			fmt.Printf("%q : %d times\n", line, tally.Total)
			for file, count := range tally.Files {
				fmt.Printf("  %s: %d times\n", file, count)
			}
		}
	}
}

func countLines(f *os.File, counts map[string]*Tally, arg string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		line := input.Text()
		t, ok := counts[line]
		if !ok {
			t = &Tally{
				Files: make(map[string]int),
				Total: 0,
			}
			counts[line] = t
		}
		t.Files[arg]++
		t.Total++
	}
	// ignore error from input.Err()
}
