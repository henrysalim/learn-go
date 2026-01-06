package main

import (
	"fmt"
	"iter"
	"slices"
)

func checkMAS(seq iter.Seq2[int, string]) bool {
	var XMAS = []string{"X", "M", "A", "S"}
	var idx = 0

	for _, val := range seq {
		if val != XMAS[idx] {
			return false
		}

		idx++
	}

	return true
}

func main() {
	var workingExampleForward = []string{"X", "M", "A", "S"}
	var workingExampleBackward = []string{"S", "A", "M", "X"}
	var nonWorkingExample = []string{"X", "M", "S", "A"}

	fmt.Println(checkMAS(slices.All(workingExampleForward)) || checkMAS(slices.Backward(workingExampleForward)))
	fmt.Println(checkMAS(slices.All(workingExampleBackward)) || checkMAS(slices.Backward(workingExampleBackward)))
	fmt.Println(checkMAS(slices.All(nonWorkingExample)) || checkMAS(slices.Backward(nonWorkingExample)))
}
