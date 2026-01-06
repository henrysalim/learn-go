package main

import "fmt"

func calculateStats(numbers []int) (int, int, float64) {
	if len(numbers) == 0 {
		return 0, 0, 0.0
	}

	sum := 0
	for _, num := range numbers {
		sum += num
	}

	avg := float64(sum) / float64(len(numbers))
	max := numbers[0]

	for _, num := range numbers {
		if num > max {
			max = num
		}
	}

	return sum, max, avg
}

func divideNumbers(a, b int) (result int, err error) {
	if b == 0 {
		err = fmt.Errorf("division by zero")
		return
	}
	result = a / b
	return
}

func main() {
	numbers := []int{10, 20, 30, 40, 50}
	total, maximum, average := calculateStats(numbers)
	fmt.Printf("Total: %d, Maximum: %d, Average: %.2f\n", total, maximum, average)

	res, err := divideNumbers(20, 10)
	fmt.Println(res)
	fmt.Println(err)
}