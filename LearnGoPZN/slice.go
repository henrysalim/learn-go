package main

import "fmt"

func main() {
	nums := make([]int, 0, 2)

	for i := 0; i < 10; i++ {
		nums = append(nums, i)
		fmt.Printf("After appending %d: len = %d, cap = %d\n", i, len(nums), cap(nums))
	}
	fmt.Println("Slice's content:", nums)
	fmt.Println(nums[:9])
	fmt.Println(nums[1:4])

	/*
		Best practices:
		1. Preallocate Capacity: If you know how many items you’ll need, use make to preallocate the slice.
		2. Avoid Overgrowing: Be mindful of large slices that could waste memory during unnecessary growth.
		3. Profile and Optimize: Use Go’s profiling tools (pprof) to analyze memory usage and optimize your slices.
	*/
}
