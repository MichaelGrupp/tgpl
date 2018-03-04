// Run with: go test -bench=. | grep s/op
package echo

import (
	"fmt"
	"os"
	"testing"
)

// BenchmarkPlainLoopEcho runs PlainLoopEcho b.N times.
func BenchmarkPlainLoopEcho(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fmt.Println(b.N)
		// os.Args is a slice, passing a slice is by reference
		PlainLoopEcho(os.Args[1:])
	}
}

// BenchmarkRangeEcho runs RangeEcho b.N times.
func BenchmarkRangeEcho(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RangeEcho(os.Args[1:])
	}
}

// BenchmarkJoinPrintlnEcho runs JoinPrintlnEcho b.N times.
func BenchmarkJoinPrintlnEcho(b *testing.B) {
	for i := 0; i < b.N; i++ {
		JoinPrintlnEcho(os.Args[1:])
	}
}

// BenchmarkPrintlnEcho runs PrintlnEcho b.N times.
func BenchmarkPrintlnEcho(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PrintlnEcho(os.Args[1:])
	}
}
