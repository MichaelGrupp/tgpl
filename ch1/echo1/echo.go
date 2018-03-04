// Simple "echo" implementation.
package main

import (
	"fmt"
	"os"
)

func main() {
	// initialized as zero value of its type, here: ""
	var s, sep string
	// deduce type with :=
	// only prefix ++, no postfix
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}
