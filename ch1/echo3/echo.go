package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// Reduce the complexity by using Join() on the os.Args slice.
	fmt.Println(strings.Join(os.Args[1:], " "))
	// Shortest version if formatting is not needed:
	// fmt.Println(os.Args[1:])
}
