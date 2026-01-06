package main

import "fmt"

func main() {
	var a [5]int
	fmt.Println("emp:", a)

	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])
	fmt.Println()

	fmt.Println("len:", len(a))
	fmt.Println()

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)
	fmt.Println()

	b = [...]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)
	fmt.Println()

	b = [...]int{100, 3: 400, 500}
	fmt.Println("idx: ", b)

	var twoD [2][3]int
	for i := range 2 {
		for j := range 3 {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)

	twoD = [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	fmt.Println("2d: ", twoD)
}
