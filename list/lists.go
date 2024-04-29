package main

import "fmt"

// Time to practice what you learned!

// 1) Create a new array (!) that contains three hobbies you have
// 		Output (print) that array in the command line.
// 2) Also output more data about that array:
//		- The first element (standalone)
//		- The second and third element combined as a new list
// 3) Create a slice based on the first element that contains
//		the first and second elements.
//		Create that slice in two different ways (i.e. create two slices in the end)
// 4) Re-slice the slice from (3) and change it to contain the second
//		and last element of the original array.
// 5) Create a "dynamic array" that contains your course goals (at least 2 goals)
// 6) Set the second goal to a different one AND then add a third goal to that existing dynamic array
// 7) Bonus: Create a "Product" struct with title, id, price and create a
//		dynamic list of products (at least 2 products).
//		Then add a third product to the existing list of products.

// type Product struct {
// 	id    int
// 	title string
// 	price float64
// }

// func main() {
// 	products := []Product{
// 		{
// 			id:    1,
// 			title: "Product 1",
// 			price: 10.99,
// 		},
// 		{
// 			id:    2,
// 			title: "Product 2",
// 			price: 20.01,
// 		},
// 	}

// 	newProduct := Product{
// 		id:    3,
// 		title: "Product 3",
// 		price: 30.50,
// 	}

// 	products = append(products, newProduct)

// 	fmt.Println(products)
// }

// func main() {
// 	prices := []float64{10.99, 8.99}
// 	prices = append(prices, 11.99)
// 	fmt.Println(prices)
// 	prices = prices[1:]
// 	fmt.Println(prices)
// }

// func main() {
// 	prices := [4]float64{10.99, 9.99, 45.99, 20.00}
// 	prices[3] = 20.01
// 	fmt.Println(prices[3])

// 	// Parts of array

// 	featuredPrices := prices[1:3]
// 	fmt.Println(featuredPrices)

// 	featuredPrices = prices[:3]
// 	fmt.Println(featuredPrices)

// 	featuredPrices = prices[1:]
// 	// Slices just a shalow copy of original array,
// 	// when modifi it whill change the original array
// 	featuredPrices[0] = 199.99
// 	fmt.Println(featuredPrices)
// 	highlightedPrices := featuredPrices[:1]
// 	fmt.Println(highlightedPrices)

// 	fmt.Println("Length:", len(featuredPrices))
// 	fmt.Println("Cap:", cap(featuredPrices))
// }

func main() {
	prices := []float64{10.99, 9.99, 45.99, 20.00}
	prices[3] = 20.01
	fmt.Println(prices[3])

	// Parts of array
	prices = append(prices, 5.99, 12.99, 29.99)
	fmt.Println(prices)

	discountPrices := []float64{101.99, 80.99, 20.59}
	prices = append(prices, discountPrices...)
	fmt.Println(prices)

}
