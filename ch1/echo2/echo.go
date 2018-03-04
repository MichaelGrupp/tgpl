package main

import (
	"fmt"
	"os"
)

func main() {
	// Allowed initialization forms:
	// s := ""  // short form with type deduction
	// var s string  // default initialization with zero value
	// var s = ""  // rarely used
	// var s string = ""  // strict typing useful when types don't match
	s, sep := "", ""
	// for i, x := range Iterable
	// Unused variables are not permitted - ignore them (here: i).
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}
