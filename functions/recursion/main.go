package main

import "fmt"

func main() {
	fac := factorial(5)

	fmt.Println(fac)
}

func factorial(number int) int {
	if number == 0 {
		return 1
	}
	return number * factorial(number-1)
}
