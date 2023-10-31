package main

import "fmt"

func main() {
	var title string
	title = "For the Love of Go"
	var author string
	author = "John Arundel"
	var copies int
	copies = 99
	var inStock bool
	inStock = true
	var royaltyPercentage float64
	royaltyPercentage = 12.5
	var isSpecialOffer bool
	isSpecialOffer = true

	printTitle(title)
	fmt.Println(author)
	fmt.Println(copies)
	fmt.Println("In Stock: ", inStock)
	fmt.Println(royaltyPercentage)
	fmt.Println("Special Offer: ", isSpecialOffer)
}

func printTitle(title string) {
	fmt.Println("Title: ", title)
}
