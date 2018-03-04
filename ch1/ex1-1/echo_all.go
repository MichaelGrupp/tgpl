package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// Print all elements of os.Args (explicit: os.Args[:])
	fmt.Println(strings.Join(os.Args, " "))
}
