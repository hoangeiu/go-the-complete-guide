package main

import "fmt"

type floatMap map[string]float64

func (m floatMap) output() {
	fmt.Println(m)
}

func main() {
	// Default null
	userNames := make([]string, 2, 5)

	userNames[0] = "Anh"
	userNames[1] = "Thu"

	userNames = append(userNames, "Hoang")
	userNames = append(userNames, "Minh")

	fmt.Println(userNames)

	// for range userNames {

	// }

	for index, value := range userNames {
		fmt.Println("Index:", index)
		fmt.Println("Value:", value)
	}

	courseRatings := make(floatMap, 3)

	courseRatings["go"] = 4.7
	courseRatings["react"] = 4.8
	courseRatings["angular"] = 3.8
	courseRatings["node"] = 2.8

	courseRatings.output()

	fmt.Println(courseRatings)

	for key, value := range courseRatings {
		fmt.Println("Index:", key)
		fmt.Println("Value:", value)
	}
}
