// Run e.g. as:
// $ echo -e "hello\nhello\nhi" | go run dup.go
// or:
// $ go run dup.go foo.txt bar.txt
package main

import (
	"bufio"
	"fmt"
	"os"
)

type count struct {
	total int
	files map[string]bool // a "set"
}

func countLines(f *os.File, counts map[string]count) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		str := input.Text()
		// if counts[str] doesn't exist yet, we get the "zero" value
		// existance could be checked with: c, exists := counts[str]
		c := counts[str]
		c.total++
		if c.files == nil {
			c.files = make(map[string]bool)
		}
		if !c.files[f.Name()] {
			c.files[f.Name()] = true
		}
		counts[str] = c
	}
}

func main() {
	counts := make(map[string]count)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}
	for line, c := range counts {
		if c.total > 1 {
			// see docs:"when printing structs, the plus flag (%+v) adds field names"
			fmt.Printf("%s\t%+v\n", line, c)
		}
	}
}
