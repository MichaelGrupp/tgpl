// Run e.g. as:
// $ echo -e "hello\nhello\nhi" | go run dup.go
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Map keys can be of any type that has an '==' operator.
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	// Each call to scan fetches the next line and strips the '\n'
	for input.Scan() {
		counts[input.Text()]++
	}

	fmt.Println("Duplicated lines:")
	// Range for loop over key, value pairs.
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
