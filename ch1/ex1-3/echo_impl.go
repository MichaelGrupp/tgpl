package echo

import (
	"fmt"
	"strings"
)

// PlainLoopEcho is an echo implementation with a plain old for-loop.
func PlainLoopEcho(args []string) {
	s, sep := "", ""
	for i := 0; i < len(args); i++ {
		s += sep + args[i]
		sep = " "
	}
	fmt.Println(s)
}

// RangeEcho is an echo implementation with a range for-loop.
func RangeEcho(args []string) {
	s, sep := "", ""
	for _, arg := range args {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}

// JoinPrintlnEcho is an echo implementation with strings.Join().
func JoinPrintlnEcho(args []string) {
	fmt.Println(strings.Join(args, " "))
}

// PrintlnEcho is an echo implementation with fmt.Println() only.
func PrintlnEcho(args []string) {
	fmt.Println(args)
}
