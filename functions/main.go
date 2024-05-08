package main

import "fmt"

func main() {
	// numbers := []int{1, 10, 15}
	// sum := sumup(numbers)

	numbers := []int{1, 10, 15}
	sum := sumup(1, 10, 15)
	anotherSum := sumup(numbers...)

	fmt.Println(sum)
	fmt.Println(anotherSum)
}

func sumup(numbers ...int) int {
	sum := 0

	for _, v := range numbers {
		sum += v
	}

	return sum
}

// func sumup(numbers []int) int {
// 	sum := 0

// 	for _, v := range numbers {
// 		sum += v
// 	}

// 	return sum
// }
